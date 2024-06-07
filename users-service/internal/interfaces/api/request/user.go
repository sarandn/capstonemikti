package request

type UserServiceRequest struct {
	RoleIDFK int    `validate:"required" json:"role_id_fk"`
	FullName string `validate:"required" json:"full_name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
	PhoneNum string `validate:"required,gte=10,lte=15,numeric" json:"phone_num"`
	Address  string `validate:"required" json:"address"`
}
