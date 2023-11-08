package res

type BoardActionDto struct {
	Id          int    `json:"id,omitempty"`
	CafeId      int    `json:"cafe_id,omitempty"`
	BoardTypeId int    `json:"board_type_id,omitempty"`
	ReadRoles   string `json:"read_roles,omitempty"`
	CreateRoles string `json:"create_roles,omitempty"`
	UpdateRoles string `json:"update_roles,omitempty"`
	UpdateAble  bool   `json:"update_able,omitempty"`
	DeleteRoles string `json:"delete_roles,omitempty"`
}
