package controller

import (
	"board_action/internal/controller/req"
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

func NewController(s service.Service) Controller {
	return Controller{s: s}
}
