package request

type RoleServiceRequest struct {
	RoleID   int    `validate:"required" json:"role_id"`
	RoleName string `validate:"required" json:"role_name"`
}

type RoleUpdateServiceRequest struct {
	RoleName string `json:"role_name"`
}
