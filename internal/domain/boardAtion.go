package domain

import "time"

type BoardAction struct {
	Id          int
	CafeId      int
	CafeTypeId  int
	ReadRoles   string
	CreateRoles string
	UpdateRoles string
	UpdateAble  bool
	DeleteRoles string
	CreatedAt   time.Time
}
