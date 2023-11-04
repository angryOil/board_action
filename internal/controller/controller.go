package controller

import (
	"board_action/internal/controller/req"
	"board_action/internal/controller/res"
	"board_action/internal/service"
	"context"
)

type Controller struct {
	s service.Service
}

func (c Controller) Create(ctx context.Context, cafeId int, boardTypeId int, rDto req.CreateDto) error {
	d := rDto.ToDomain(cafeId, boardTypeId)
	err := c.s.Create(ctx, d)
	return err
}

func (c Controller) GetInfo(ctx context.Context, cafeId int, boardTypeId int) (res.BoardActionDto, error) {
	d, err := c.s.GetInfo(ctx, cafeId, boardTypeId)
	if err != nil {
		return res.BoardActionDto{}, err
	}
	return res.ToDto(d), err
}

func (c Controller) Patch(ctx context.Context, cafeId int, typeId int, rDto req.PatchDto) error {
	d := rDto.ToDomain(cafeId, typeId)
	err := c.s.Patch(ctx, d)
	return err
}

func NewController(s service.Service) Controller {
	return Controller{s: s}
}
