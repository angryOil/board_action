package vo

import "time"

type Update struct {
	Id          int
	CafeId      int
	BoardTypeId int
	ReadRoles   string
	CreateRoles string
	CreatedAt   time.Time
}
