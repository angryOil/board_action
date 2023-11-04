package req

import (
	"board_action/internal/domain"
	"time"
)

type CreateDto struct {
	ReadRoles   string `json:"read_roles"`
	CreateRoles string `json:"create_roles"`
	UpdateRoles string `json:"update_roles"`
	UpdateAble  bool   `json:"update_able"`
	DeleteRoles string `json:"delete_roles"`
}

func (d CreateDto) ToDomain(cafeId, boardTypeId int) domain.BoardAction {
	return domain.BoardAction{
		CafeId:      cafeId,
		BoardTypeId: boardTypeId,
		ReadRoles:   d.ReadRoles,
		CreateRoles: d.CreateRoles,
		UpdateRoles: d.UpdateRoles,
		UpdateAble:  d.UpdateAble,
		DeleteRoles: d.DeleteRoles,
		CreatedAt:   time.Now(),
	}
}

type PatchDto struct {
	ReadRoles   string `json:"read_roles"`
	CreateRoles string `json:"create_roles"`
	UpdateRoles string `json:"update_roles"`
	UpdateAble  bool   `json:"update_able"`
	DeleteRoles string `json:"delete_roles"`
}

func (d PatchDto) ToDomain(cafeId, boardTypeId int) domain.BoardAction {
	return domain.BoardAction{
		CafeId:      cafeId,
		BoardTypeId: boardTypeId,
		ReadRoles:   d.ReadRoles,
		CreateRoles: d.CreateRoles,
		UpdateRoles: d.UpdateRoles,
		UpdateAble:  d.UpdateAble,
		DeleteRoles: d.DeleteRoles,
	}
}
