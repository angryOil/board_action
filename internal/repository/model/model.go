package model

import (
	"board_action/internal/domain"
	"board_action/internal/repository/req"
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
	CreatedAt   time.Time `bun:"created_at"`
}

func ToCreateModel(c req.Create) BoardAction {
	return BoardAction{
		CafeId:      c.CafeId,
		BoardTypeId: c.BoardTypeId,
		ReadRoles:   c.ReadRoles,
		CreateRoles: c.CreateRoles,
		CreatedAt:   c.CreatedAt,
	}
}

func ToUpdateModel(u req.Update) BoardAction {
	return BoardAction{
		CafeId:      u.CafeId,
		BoardTypeId: u.BoardTypeId,
		ReadRoles:   u.ReadRoles,
		CreateRoles: u.CreateRoles,
		CreatedAt:   u.CreatedAt,
	}
}

func (b BoardAction) ToDomain() domain.BoardAction {
	return domain.NewBoardActionBuilder().
		Id(b.Id).
		CafeId(b.CafeId).
		BoardTypeId(b.BoardTypeId).
		ReadRoles(b.ReadRoles).
		CreateRoles(b.CreateRoles).
		CreatedAt(b.CreatedAt).
		Build()
}

func ToDomainList(models []BoardAction) []domain.BoardAction {
	result := make([]domain.BoardAction, len(models))
	for i, m := range models {
		result[i] = m.ToDomain()
	}
	return result
}
