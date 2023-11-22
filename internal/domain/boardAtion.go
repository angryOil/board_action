package domain

import (
	"board_action/internal/domain/vo"
	"errors"
	"time"
)

var _ BoardAction = (*boardAction)(nil)

type BoardAction interface {
	ValidCreate() error
	ValidUpdate() error

	Update(readRoles, createRoles, updateRoles, deleteRoles string, updateAble bool) BoardAction
	ToUpdate() vo.Update
	ToInfo() vo.Info
}

type boardAction struct {
	id          int
	cafeId      int
	boardTypeId int
	readRoles   string
	createRoles string
	createdAt   time.Time
}

func (b *boardAction) ToUpdate() vo.Update {
	return vo.Update{
		Id:          b.id,
		CafeId:      b.cafeId,
		BoardTypeId: b.boardTypeId,
		ReadRoles:   b.readRoles,
		CreateRoles: b.createRoles,
		CreatedAt:   b.createdAt,
	}
}

const (
	InvalidId          = "invalid id"
	InvalidCafeId      = "invalid cafe id"
	InvalidBoardTypeId = "invalid board type id"
)

func (b *boardAction) ValidCreate() error {
	if b.cafeId < 1 {
		return errors.New(InvalidCafeId)
	}
	if b.boardTypeId < 1 {
		return errors.New(InvalidBoardTypeId)
	}
	return nil
}

func (b *boardAction) ValidUpdate() error {
	if b.id < 1 {
		return errors.New(InvalidId)
	}
	if b.cafeId < 1 {
		return errors.New(InvalidCafeId)
	}
	if b.boardTypeId < 1 {
		return errors.New(InvalidBoardTypeId)
	}
	return nil
}

func (b *boardAction) Update(readRoles, createRoles, updateRoles, deleteRoles string, updateAble bool) BoardAction {
	b.readRoles = readRoles
	b.createRoles = createRoles
	return b
}

func (b *boardAction) ToInfo() vo.Info {
	return vo.Info{
		Id:          b.id,
		CafeId:      b.cafeId,
		BoardTypeId: b.boardTypeId,
		ReadRoles:   b.readRoles,
		CreateRoles: b.createRoles,
	}
}
