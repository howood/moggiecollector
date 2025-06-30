package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/howood/moggiecollector/di/dbcluster"
	"github.com/howood/moggiecollector/di/svcluster"
	"github.com/howood/moggiecollector/domain/dto"
	"github.com/howood/moggiecollector/domain/entity"
	"github.com/howood/moggiecollector/domain/model"
	"gorm.io/gorm"
)

type UserMfaUsecase struct {
	DataStore dbcluster.DataStore
	SvCluster *svcluster.ServiceCluster
}

func NewUserMfaUsecase(dataStore dbcluster.DataStore, svCluster *svcluster.ServiceCluster) *UserMfaUsecase {
	return &UserMfaUsecase{
		DataStore: dataStore,
		SvCluster: svCluster,
	}
}

func (uc *UserMfaUsecase) UpsertAuthenticator(ctx context.Context, req *dto.UserMfaTotpDto) error {
	ok, err := uc.SvCluster.AuthenticatorSV.ValidateBySecretString(ctx, req.Passcode, req.Secret)
	if err != nil {
		return fmt.Errorf("failed to validate passcode: %w", err)
	}
	if !ok {
		return errors.New("invalid passcode")
	}
	userMfa, err := uc.DataStore.DSRepository().UserMfaRepository.Get(uc.DataStore.DBInstanceClient(ctx), req.UserID, model.MfaTypeTOTP)
	if err != nil {
		if !uc.DataStore.RecordNotFoundError(err) {
			return fmt.Errorf("failed to get user MFA authenticator: %w", err)
		}
		userMfa = &model.UserMfa{
			UserID:  req.UserID,
			MfaType: model.MfaTypeTOTP,
		}
	}
	userMfa.Secret = req.Secret
	userMfa.IsDefault = req.IsDefault

	return uc.DataStore.DBInstanceClient(ctx).Transaction(func(tx *gorm.DB) error {
		if err := uc.DataStore.DSRepository().UserMfaRepository.Upsert(tx, userMfa); err != nil {
			return fmt.Errorf("failed to create user MFA authenticator: %w", err)
		}
		if userMfa.IsDefault {
			if err := uc.DataStore.DSRepository().UserMfaRepository.UnsetDefault(tx, userMfa.UserID, model.MfaTypeTOTP); err != nil {
				return fmt.Errorf("failed to unset default user MFA authenticator: %w", err)
			}
		}
		return nil
	})
}

func (uc *UserMfaUsecase) GetAuthenticatorSecret(ctx context.Context, userID uuid.UUID) (string, error) {
	return uc.SvCluster.AuthenticatorSV.GenerateSecret(ctx, userID)
}

func (uc *UserMfaUsecase) GetDefaultMfa(ctx context.Context, userID uuid.UUID) (*entity.MfaType, error) {
	accountMfa, err := uc.DataStore.DSRepository().UserMfaRepository.GetDefault(uc.DataStore.DBInstanceClient(ctx), userID)
	if err != nil && !uc.DataStore.RecordNotFoundError(err) {
		return nil, err
	}
	if accountMfa == nil {
		return nil, nil
	}
	mfaType := entity.MfaType(accountMfa.MfaType)
	return &mfaType, nil
}

func (uc *UserMfaUsecase) ValidateAuthenticatorCode(ctx context.Context, verifyMfaAuthenticator dto.VerifyMfaAuthenticator) (bool, error) {
	return uc.SvCluster.AuthenticatorSV.Validate(ctx, verifyMfaAuthenticator.UserID, verifyMfaAuthenticator.Passcode)
}
