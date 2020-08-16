package entity

// LoginUserForm entity
type LoginUserForm struct {
	Email    string `validate:"required"`
	Password string `validate:"required,min=8"`
}
