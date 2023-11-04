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

func NewRepository(db bun.IDB) Repository {
	return Repository{db: db}
}
