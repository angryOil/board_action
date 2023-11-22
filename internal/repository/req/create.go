package req

import "time"

type Create struct {
	CafeId      int
	BoardTypeId int
	ReadRoles   string
	CreateRoles string
	CreatedAt   time.Time
}
