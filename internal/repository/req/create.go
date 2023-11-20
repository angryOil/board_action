package req

import "time"

type Create struct {
	CafeId      int
	BoardTypeId int
	ReadRoles   string
	CreateRoles string
	UpdateRoles string
	UpdateAble  bool
	DeleteRoles string
	CreatedAt   time.Time
}
