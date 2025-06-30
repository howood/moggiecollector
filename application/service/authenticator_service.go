package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/howood/moggiecollector/config"
	"github.com/howood/moggiecollector/di/dbcluster"
	"github.com/howood/moggiecollector/domain/model"
	"github.com/howood/moggiecollector/infrastructure/mfa"
)

type AuthenticatorService interface {
	GenerateSecret(ctx context.Context, userID uuid.UUID) (string, error)
	Validate(ctx context.Context, userID uuid.UUID, passcode string) (bool, error)
	ValidateBySecretString(ctx context.Context, passcode string, secret string) (bool, error)
}

type authenticatorSv struct {
	authenticator *mfa.Authenticator
	DataStore     dbcluster.DataStore
}

// NewAuthenticatorService creates a AuthenticatorService.
func NewAuthenticatorService(dataStore dbcluster.DataStore) AuthenticatorService {
	return &authenticatorSv{
		authenticator: mfa.NewAuthenticator(),
		DataStore:     dataStore,
	}
}

func (as *authenticatorSv) GenerateSecret(ctx context.Context, userID uuid.UUID) (string, error) {
	return as.authenticator.GenerateKey(userID, config.TotpPeriodD)
}

func (as *authenticatorSv) Validate(ctx context.Context, userID uuid.UUID, passcode string) (bool, error) {
	userMfa, err := as.DataStore.DSRepository().UserMfaRepository.Get(as.DataStore.DBInstanceClient(ctx).WithContext(ctx), userID, model.MfaTypeTOTP)
	if err != nil {
		return false, err
	}
	return as.authenticator.Validate(passcode, userMfa.Secret, config.TotpPeriodD)
}

func (as *authenticatorSv) ValidateBySecretString(ctx context.Context, passcode string, secret string) (bool, error) {
	return as.authenticator.Validate(passcode, secret, config.TotpPeriodD)
}
