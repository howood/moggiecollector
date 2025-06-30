package request

// UpsertUserAuthenticatorForm entity
type UpsertUserAuthenticatorForm struct {
	Secret    string `form:"secret"    validate:"required"`
	Passcode  string `form:"passcode"  validate:"required"`
	IsDefault bool   `form:"is_default"`
}
