package repository

import (
	"board_action/internal/domain"
	"board_action/internal/domain/vo"
	"board_action/internal/repository/model"
	"board_action/internal/repository/req"
	"context"
	"errors"
	"github.com/uptrace/bun"
	"log"
	"strings"
)

type Repository struct {
	db bun.IDB
}

func NewRepository(db bun.IDB) Repository {
	return Repository{db: db}
}

const (
	InternalServerError = "internal server error"
)

func (r Repository) Create(ctx context.Context, c req.Create) error {
	m := model.ToCreateModel(c)
	_, err := r.db.NewInsert().Model(&m).Exec(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return err
		}
		log.Println("Create NewInsert err: ", err)
		return errors.New(InternalServerError)
	}
	return nil
}

func (r Repository) GetByCafeIdTypeId(ctx context.Context, cafeId int, typeId int) (domain.BoardAction, error) {
	var models []model.BoardAction
	err := r.db.NewSelect().Model(&models).Where("cafe_id = ? and board_type_id = ?", cafeId, typeId).Scan(ctx)
	if err != nil {
		log.Println("GetByCafeIdTypeId NewSelect err: ", err)
		return domain.NewBoardActionBuilder().Build(), errors.New(InternalServerError)
	}
	if len(models) == 0 {
		return domain.NewBoardActionBuilder().Build(), nil
	}

	return models[0].ToDomain(), nil
}

func (r Repository) Patch(ctx context.Context, id int,
	validFunc func(domains []domain.BoardAction) (domain.BoardAction, error),
	mergeFunc func(oldD domain.BoardAction) (vo.Update, error)) error {

	db, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Println("Patch BeginTx err: ", err)
		return errors.New(InternalServerError)
	}
	var models []model.BoardAction
	err = db.NewSelect().Model(&models).Where("id = ?", id).Scan(ctx)
	if err != nil {
		log.Println("Patch NewSelect err: ", err)
		return errors.New("internal server error")
	}

	validDomain, err := validFunc(model.ToDomainList(models))
	if err != nil {
		return err
	}

	mergedV, err := mergeFunc(validDomain)
	if err != nil {
		return err
	}

	mergedModel := model.ToUpdateModel(req.Update{
		Id:          mergedV.Id,
		CafeId:      mergedV.CafeId,
		BoardTypeId: mergedV.BoardTypeId,
		ReadRoles:   mergedV.ReadRoles,
		CreateRoles: mergedV.CreateRoles,
		UpdateRoles: mergedV.UpdateRoles,
		UpdateAble:  mergedV.UpdateAble,
		DeleteRoles: mergedV.DeleteRoles,
		CreatedAt:   mergedV.CreatedAt,
	})

	_, err = db.NewInsert().Model(&mergedModel).
		On("CONFLICT (id) DO UPDATE").
		On("conflict (cafe_id,board_type_id) do update").
		Exec(ctx)

	if err != nil {
		log.Println("Patch NewInsert err: ", err)
		return errors.New(InternalServerError)
	}

	err = db.Commit()
	if err != nil {
		log.Println("Patch Commit() err: ", err)
		return errors.New(InternalServerError)
	}
	return nil
}

func (r Repository) Delete(ctx context.Context, cafeId int, typeId int, id int) error {
	var m model.BoardAction
	_, err := r.db.NewDelete().Model(&m).Where("cafe_id = ? and board_type_id = ? and id = ?", cafeId, typeId, id).Exec(ctx)
	if err != nil {
		log.Println("Delete NewDelete err: ", err)
		return errors.New("internal server error")
	}
	return err
}
