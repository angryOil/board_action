package controller

import (
	"board_action/internal/controller/req"
	"board_action/internal/controller/res"
	"board_action/internal/service"
	req2 "board_action/internal/service/req"
	"context"
)

type Controller struct {
	s service.Service
}

func (c Controller) Create(ctx context.Context, cafeId int, boardTypeId int, d req.CreateDto) error {
	err := c.s.Create(ctx, req2.Create{
		CafeId:      cafeId,
		BoardTypeId: boardTypeId,
		ReadRoles:   d.ReadRoles,
		CreateRoles: d.CreateRoles,
		UpdateRoles: d.UpdateRoles,
		UpdateAble:  d.UpdateAble,
		DeleteRoles: d.DeleteRoles,
	})
	return err
}

func (c Controller) GetInfo(ctx context.Context, cafeId int, boardTypeId int) (res.BoardActionDto, error) {
	d, err := c.s.GetInfo(ctx, cafeId, boardTypeId)
	if err != nil {
		return res.BoardActionDto{}, err
	}
	return res.BoardActionDto{
		Id:          d.Id,
		CafeId:      d.CafeId,
		BoardTypeId: d.BoardTypeId,
		ReadRoles:   d.ReadRoles,
		CreateRoles: d.CreateRoles,
		UpdateRoles: d.UpdateRoles,
		UpdateAble:  d.UpdateAble,
		DeleteRoles: d.DeleteRoles,
	}, err
}

func (c Controller) Patch(ctx context.Context, id, cafeId int, typeId int, d req.PatchDto) error {
	err := c.s.Patch(ctx, req2.Update{
		Id:          id,
		CafeId:      cafeId,
		BoardTypeId: typeId,
		ReadRoles:   d.ReadRoles,
		CreateRoles: d.CreateRoles,
		UpdateRoles: d.UpdateRoles,
		UpdateAble:  d.UpdateAble,
		DeleteRoles: d.DeleteRoles,
	})
	return err
}

func (c Controller) Delete(ctx context.Context, cafeId int, typeId int, id int) error {
	err := c.s.Delete(ctx, cafeId, typeId, id)
	return err
}

func NewController(s service.Service) Controller {
	return Controller{s: s}
}
