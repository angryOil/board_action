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

func NewRepository(db bun.IDB) Repository {
	return Repository{db: db}
}
