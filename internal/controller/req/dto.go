package req

type CreateDto struct {
	ReadRoles   string `json:"read_roles"`
	CreateRoles string `json:"create_roles"`
}

type PatchDto struct {
	ReadRoles   string `json:"read_roles"`
	CreateRoles string `json:"create_roles"`
}
