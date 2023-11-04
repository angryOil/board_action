package model

import (
	"board_action/internal/domain"
	"github.com/uptrace/bun"
	"time"
)

type BoardAction struct {
	bun.BaseModel `bun:"table:board_action,alias:ba"`

	Id          int       `bun:"id,pk,autoincrement"`
	CafeId      int       `bun:"cafe_id,notnull"`
	BoardTypeId int       `bun:"board_type_id,notnull"`
	ReadRoles   string    `bun:"read_roles,notnull"`
	CreateRoles string    `bun:"create_roles,notnull"`
	UpdateRoles string    `bun:"update_roles,notnull"`
	UpdateAble  bool      `bun:"update_able,notnull"`
	DeleteRoles string    `bun:"delete_roles,notnull"`
	CreatedAt   time.Time `bun:"created_at"`
}

func ToModel(d domain.BoardAction) BoardAction {
	return BoardAction{
		Id:          d.Id,
		CafeId:      d.CafeId,
		BoardTypeId: d.BoardTypeId,
		ReadRoles:   d.ReadRoles,
		CreateRoles: d.CreateRoles,
		UpdateRoles: d.UpdateRoles,
		UpdateAble:  d.UpdateAble,
		DeleteRoles: d.DeleteRoles,
		CreatedAt:   d.CreatedAt,
	}
}

func (m BoardAction) ToDomain() domain.BoardAction {
	return domain.BoardAction{
		Id:          m.Id,
		CafeId:      m.CafeId,
		BoardTypeId: m.BoardTypeId,
		ReadRoles:   m.ReadRoles,
		CreateRoles: m.CreateRoles,
		UpdateRoles: m.UpdateRoles,
		UpdateAble:  m.UpdateAble,
		DeleteRoles: m.DeleteRoles,
		CreatedAt:   m.CreatedAt,
	}
}

func ToDomainList(models []BoardAction) []domain.BoardAction {
	results := make([]domain.BoardAction, len(models))
	for i, m := range models {
		results[i] = m.ToDomain()
	}
	return results
}
