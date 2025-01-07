package form

// CreateUserForm entity
type CreateUserForm struct {
	Name     string `form:"name"     validate:"required,max=255"`
	Email    string `form:"email"    validate:"required,max=255"`
	Password string `form:"password" validate:"required,min=8,max=255"`
}
