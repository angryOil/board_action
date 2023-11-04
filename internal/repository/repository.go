package repository

import (
	"board_action/internal/domain"
	"board_action/internal/repository/model"
	"context"
	"errors"
	"github.com/uptrace/bun"
	"log"
	"strings"
)

type Repository struct {
	db bun.IDB
}

func (r Repository) Create(ctx context.Context, d domain.BoardAction) error {
	m := model.ToModel(d)
	_, err := r.db.NewInsert().Model(&m).Exec(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return err
		}
		log.Println("Create NewInsert err: ", err)
		return errors.New("internal server error")
	}
	return nil
}

func (r Repository) GetByCafeIdTypeId(ctx context.Context, cafeId int, typeId int) (domain.BoardAction, error) {
	var models []model.BoardAction
	err := r.db.NewSelect().Model(&models).Where("cafe_id = ? and board_type_id = ?", cafeId, typeId).Scan(ctx)
	if err != nil {
		log.Println("GetByCafeIdTypeId NewSelect err: ", err)
		return domain.BoardAction{}, errors.New("internal server error")
	}
	if len(models) == 0 {
		return domain.BoardAction{}, nil
	}
	return models[0].ToDomain(), nil
}

func (r Repository) Patch(ctx context.Context, cafeId int, typeId int, validFunc func(domains []domain.BoardAction) (domain.BoardAction, error), mergeFunc func(oldD domain.BoardAction) domain.BoardAction) error {
	var models []model.BoardAction
	err := r.db.NewSelect().Model(&models).Where("cafe_id = ? and board_type_id =?", cafeId, typeId).Scan(ctx)
	if err != nil {
		log.Println("Patch NewSelect err: ", err)
		return errors.New("internal server error")
	}
	validDomain, err := validFunc(model.ToDomainList(models))
	if err != nil {
		return err
	}
	mergedDomain := mergeFunc(validDomain)
	mergedModel := model.ToModel(mergedDomain)

	_, err = r.db.NewInsert().Model(&mergedModel).
		On("CONFLICT (id) DO UPDATE").Exec(ctx)

	if err != nil {
		log.Println("Patch NewInsert err: ", err)
		return errors.New("internal server error")
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

func NewRepository(db bun.IDB) Repository {
	return Repository{db: db}
}
