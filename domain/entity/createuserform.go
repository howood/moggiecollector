package entity

// CreateUserForm entity
type CreateUserForm struct {
	Name     string `validate:"required,max=255"`
	Email    string `validate:"required,max=255"`
	Password string `validate:"required,min=8,max=255"`
}
