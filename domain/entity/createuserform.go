package entity

// CreateUserForm entity
type CreateUserForm struct {
	Name     string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required,min=8"`
}
