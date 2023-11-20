package domain

import "time"

var _ BoardActionBuilder = (*boardActionBuilder)(nil)

func NewBoardActionBuilder() BoardActionBuilder {
	return &boardActionBuilder{}
}

type BoardActionBuilder interface {
	Id(id int) BoardActionBuilder
	CafeId(cafeId int) BoardActionBuilder
	BoardTypeId(boardTypeId int) BoardActionBuilder
	ReadRoles(readRoles string) BoardActionBuilder
	CreateRoles(createRoles string) BoardActionBuilder
	CreatedAt(createdAt time.Time) BoardActionBuilder
	Build() BoardAction
}

type boardActionBuilder struct {
	id          int
	cafeId      int
	boardTypeId int
	readRoles   string
	createRoles string
	createdAt   time.Time
}

func (b *boardActionBuilder) Id(id int) BoardActionBuilder {
	b.id = id
	return b
}

func (b *boardActionBuilder) CafeId(cafeId int) BoardActionBuilder {
	b.cafeId = cafeId
	return b
}

func (b *boardActionBuilder) BoardTypeId(boardTypeId int) BoardActionBuilder {
	b.boardTypeId = boardTypeId
	return b
}

func (b *boardActionBuilder) ReadRoles(readRoles string) BoardActionBuilder {
	b.readRoles = readRoles
	return b
}

func (b *boardActionBuilder) CreateRoles(createRoles string) BoardActionBuilder {
	b.createRoles = createRoles
	return b
}

func (b *boardActionBuilder) CreatedAt(createdAt time.Time) BoardActionBuilder {
	b.createdAt = createdAt
	return b
}

func (b *boardActionBuilder) Build() BoardAction {
	return &boardAction{
		id:          b.id,
		cafeId:      b.cafeId,
		boardTypeId: b.boardTypeId,
		readRoles:   b.readRoles,
		createRoles: b.createRoles,
		createdAt:   b.createdAt,
	}
}
