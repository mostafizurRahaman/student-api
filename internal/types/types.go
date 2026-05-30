package types

type Student struct {
	Id        int    `json:"id"`
	Name      string `json:"name" validate:"required,min=2,max=20"`
	Email     string `json:"email" validte:"required,email"`
	Role      string `json:"role" validate:"required,oneof=user admin super_admin"`
	HasIdCard bool   `json:"hasIdCard" validate:"required"`
}
