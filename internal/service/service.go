package service

import (
	"board_action/internal/domain"
	"board_action/internal/repository"
	"context"
	"errors"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return Service{repo: repo}
}

func (s Service) Create(ctx context.Context, d domain.BoardAction) error {
	err := validFiled(d)
	if err != nil {
		return err
	}
	err = s.repo.Create(ctx, d)
	return err
}

func (s Service) GetInfo(ctx context.Context, cafeId int, typeId int) (domain.BoardAction, error) {
	d, err := s.repo.GetByCafeIdTypeId(ctx, cafeId, typeId)
	return d, err
}

func validFiled(d domain.BoardAction) error {
	if d.BoardTypeId == 0 {
		return errors.New("invalid cafe type id")
	}
	if d.CafeId == 0 {
		return errors.New("invalid cafe id ")
	}
	return nil
}
