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

func (s Service) Patch(ctx context.Context, reqD domain.BoardAction) error {
	err := validFiled(reqD)
	if err != nil {
		return err
	}
	err = s.repo.Patch(ctx, reqD.CafeId, reqD.BoardTypeId,
		func(domains []domain.BoardAction) (domain.BoardAction, error) {
			if len(domains) == 0 {
				return domain.BoardAction{}, errors.New("no rows")
			}
			return domains[0], nil
		},
		func(oldD domain.BoardAction) domain.BoardAction {
			return domain.BoardAction{
				Id:          oldD.Id,
				CafeId:      oldD.CafeId,
				BoardTypeId: oldD.BoardTypeId,
				ReadRoles:   reqD.ReadRoles,
				CreateRoles: reqD.CreateRoles,
				UpdateRoles: reqD.UpdateRoles,
				UpdateAble:  reqD.UpdateAble,
				DeleteRoles: reqD.DeleteRoles,
				CreatedAt:   oldD.CreatedAt,
			}
		},
	)
	return err
}

func (s Service) Delete(ctx context.Context, cafeId int, typeId int, id int) error {
	err := s.repo.Delete(ctx, cafeId, typeId, id)
	return err
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
