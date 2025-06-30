package mfa

import (
	"crypto/rand"
	"time"

	"github.com/google/uuid"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

const (
	issuerName = "TokiumID"
	secretSize = 20
)

type Authenticator struct {
	issuer string
}

func NewAuthenticator() *Authenticator {
	return &Authenticator{
		issuer: issuerName,
	}
}

func (a *Authenticator) initializeKey(userID uuid.UUID, period uint) (*otp.Key, error) {
	return totp.Generate(totp.GenerateOpts{
		Issuer:      a.issuer,
		AccountName: userID.String(),
		Period:      period,
		SecretSize:  secretSize,
		Secret:      nil,
		Digits:      otp.DigitsSix,
		Algorithm:   otp.AlgorithmSHA1,
		Rand:        rand.Reader,
	})
}

func (a *Authenticator) GenerateKey(userID uuid.UUID, period uint) (string, error) {
	key, err := a.initializeKey(userID, period)
	if err != nil {
		return "", err
	}

	return key.Secret(), nil
}

func (a *Authenticator) Validate(passcode, secret string, period uint) (bool, error) {
	return totp.ValidateCustom(passcode, secret, time.Now(), totp.ValidateOpts{
		Period:    period,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
}
