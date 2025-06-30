package request

type VerifyAuthenticatorForm struct {
	Identifier string `form:"identifier" validate:"required"`
	Passcode   string `form:"passcode" validate:"required"`
}
