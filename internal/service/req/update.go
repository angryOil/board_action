package req

type Update struct {
	Id          int
	CafeId      int
	BoardTypeId int
	ReadRoles   string
	CreateRoles string
	UpdateRoles string
	UpdateAble  bool
	DeleteRoles string
}
