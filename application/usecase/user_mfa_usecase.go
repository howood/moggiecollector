package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/howood/moggiecollector/config"
	"github.com/howood/moggiecollector/di/dbcluster"
	"github.com/howood/moggiecollector/di/svcluster"
	"github.com/howood/moggiecollector/domain/dto"
	"github.com/howood/moggiecollector/domain/entity"
	"github.com/howood/moggiecollector/domain/model"
	"github.com/howood/moggiecollector/library/utils"
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
	ok, err := uc.SvCluster.AuthenticatorSV.ValidateBySecretString(req.Passcode, req.Secret)
	if err != nil {
		return fmt.Errorf("failed to validate passcode: %w", err)
	}
	if !ok {
		//nolint:err113
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

func (uc *UserMfaUsecase) GetAuthenticatorSecret(_ context.Context, userID uuid.UUID) (string, error) {
	return uc.SvCluster.AuthenticatorSV.GenerateSecret(userID)
}

func (uc *UserMfaUsecase) GetDefaultMfa(ctx context.Context, userID uuid.UUID) (*entity.MfaType, error) {
	accountMfa, err := uc.DataStore.DSRepository().UserMfaRepository.GetDefault(uc.DataStore.DBInstanceClient(ctx), userID)
	if err != nil && !uc.DataStore.RecordNotFoundError(err) {
		return nil, err
	}
	if accountMfa == nil {
		//nolint:nilnil
		return nil, nil
	}
	mfaType := entity.MfaType(accountMfa.MfaType)
	return &mfaType, nil
}

func (uc *UserMfaUsecase) ValidateAuthenticatorCode(ctx context.Context, verifyMfaAuthenticator *dto.VerifyMfaAuthenticator) (bool, *entity.User, error) {
	val, ok, err := uc.SvCluster.AuthCacheSV.Get(ctx, verifyMfaAuthenticator.Identifier)
	if err != nil {
		return false, nil, err
	}
	if !ok {
		//nolint:err113
		return false, nil, errors.New("invalid identifier")
	}

	switch v := val.(type) {
	case string:
		var user entity.User
		if err := utils.ByteToJSONStruct([]byte(v), &user); err != nil {
			return false, nil, fmt.Errorf("failed to unmarshal user from cache: %w", err)
		}
		ok, err := uc.SvCluster.AuthenticatorSV.Validate(ctx, user.ID, verifyMfaAuthenticator.Passcode)
		if err != nil {
			return false, nil, fmt.Errorf("failed to validate passcode: %w", err)
		}
		if !ok {
			//nolint:err113
			return false, nil, errors.New("invalid passcode")
		}
		if err := uc.SvCluster.AuthCacheSV.Del(ctx, verifyMfaAuthenticator.Identifier); err != nil {
			return false, nil, fmt.Errorf("failed to delete cache: %w", err)
		}
		return true, &user, nil
	default:
		//nolint:err113
		return false, nil, fmt.Errorf("unexpected type %T for identifier", val)
	}
}

func (uc *UserMfaUsecase) IsUseMfa(ctx context.Context, user *entity.User) (*entity.VerifyMfa, entity.MfaType, error) {
	userMfa, err := uc.DataStore.DSRepository().UserMfaRepository.GetDefault(uc.DataStore.DBInstanceClient(ctx), user.ID)
	if err != nil {
		if !uc.DataStore.RecordNotFoundError(err) {
			return nil, "", err
		}
		return nil, "", nil
	}
	verifyMfa := entity.NewVerifyMfa()
	if err := uc.SvCluster.AuthCacheSV.Set(ctx, verifyMfa.Identifier, *user, time.Duration(config.TotpRedisExpired)*time.Second); err != nil {
		return nil, "", err
	}
	return verifyMfa, entity.MfaType(userMfa.MfaType), nil
}
