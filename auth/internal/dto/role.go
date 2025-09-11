package dto

type UpdateRoleParams struct {
	ID   string
	Name string
}

type ListRolesParams struct {
	RoleIDs []string
}
