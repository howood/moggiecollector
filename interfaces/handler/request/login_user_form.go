package request

// LoginUserForm struct.
type LoginUserForm struct {
	Email    string `form:"email"    validate:"required"`
	Password string `form:"password" validate:"required,min=8"`
}
