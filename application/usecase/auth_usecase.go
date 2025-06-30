package usecase

import (
	"context"

	"github.com/howood/moggiecollector/application/actor"
	"github.com/howood/moggiecollector/di/dbcluster"
	"github.com/howood/moggiecollector/domain/dto"
	"github.com/howood/moggiecollector/domain/entity"
	"github.com/howood/moggiecollector/domain/model"
)

type AuthUsecase struct {
	DataStore dbcluster.DataStore
}

func NewAuthUsecase(dataStore dbcluster.DataStore) *AuthUsecase {
	return &AuthUsecase{
		DataStore: dataStore,
	}
}

func (au *AuthUsecase) AuthUser(ctx context.Context, loginDto *dto.LoginDto) (*entity.User, error) {
	user, err := au.DataStore.DSRepository().UserRepository.GetByEmail(au.DataStore.DBInstanceClient(ctx), loginDto.Email)
	if err != nil {
		return &entity.User{}, err
	}
	if err := au.comparePassword(user, loginDto.Password); err != nil {
		return &entity.User{}, err
	}
	return entity.NewUser(&user), nil
}

// ComparePassword compares input password to roomdata password
func (au *AuthUsecase) comparePassword(user model.User, password string) error {
	return actor.PasswordOperator{}.ComparePassword(user.Password, password, user.Salt)
}
