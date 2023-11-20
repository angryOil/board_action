package service

import (
	"board_action/internal/domain"
	"board_action/internal/domain/vo"
	"board_action/internal/repository"
	req2 "board_action/internal/repository/req"
	"board_action/internal/service/req"
	"board_action/internal/service/res"
	"context"
	"errors"
	"time"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return Service{repo: repo}
}

func (s Service) Create(ctx context.Context, d req.Create) error {
	cafeId, boardTypeId := d.CafeId, d.BoardTypeId
	readRoles, createRoles := d.ReadRoles, d.CreateRoles
	createdAt := time.Now()
	err := domain.NewBoardActionBuilder().
		CafeId(cafeId).
		BoardTypeId(boardTypeId).
		ReadRoles(readRoles).
		CreateRoles(createRoles).
		CreatedAt(createdAt).
		Build().ValidCreate()
	if err != nil {
		return err
	}

	err = s.repo.Create(ctx, req2.Create{
		CafeId:      cafeId,
		BoardTypeId: boardTypeId,
		ReadRoles:   readRoles,
		CreateRoles: createRoles,
		CreatedAt:   createdAt,
	})
	return err
}

func (s Service) GetInfo(ctx context.Context, cafeId int, typeId int) (res.GetByCafeIdTypeId, error) {
	d, err := s.repo.GetByCafeIdTypeId(ctx, cafeId, typeId)
	v := d.ToInfo()
	return res.GetByCafeIdTypeId{
		Id:          v.Id,
		CafeId:      v.CafeId,
		BoardTypeId: v.BoardTypeId,
		ReadRoles:   v.ReadRoles,
		CreateRoles: v.CreateRoles,
	}, err
}

func (s Service) Patch(ctx context.Context, u req.Update) error {
	id, cafeId, boardTypeId := u.Id, u.CafeId, u.BoardTypeId
	readRoles, createRoles, updateRoles, deleteRoles := u.ReadRoles, u.CreateRoles, u.UpdateRoles, u.DeleteRoles
	updateAble := u.UpdateAble
	err := domain.NewBoardActionBuilder().
		Id(id).
		CafeId(cafeId).
		BoardTypeId(boardTypeId).
		ReadRoles(readRoles).
		CreateRoles(createRoles).
		Build().ValidUpdate()
	if err != nil {
		return err
	}

	err = s.repo.Patch(ctx, id,
		func(domains []domain.BoardAction) (domain.BoardAction, error) {
			if len(domains) != 1 {
				return domain.NewBoardActionBuilder().Build(), errors.New("no rows")
			}
			return domains[0], nil
		},
		func(oldD domain.BoardAction) (vo.Update, error) {
			u := oldD.Update(readRoles, createRoles, updateRoles, deleteRoles, updateAble)
			err := u.ValidUpdate()
			if err != nil {
				return vo.Update{}, err
			}
			return u.ToUpdate(), nil
		},
	)
	return err
}

func (s Service) Delete(ctx context.Context, cafeId int, typeId int, id int) error {
	err := s.repo.Delete(ctx, cafeId, typeId, id)
	return err
}
