package usecase

import (
	"context"

	"github.com/howood/moggiecollector/di/dbcluster"
	"github.com/howood/moggiecollector/domain/entity"
)

type ClientUsecase struct {
	DataStore dbcluster.DataStore
}

func NewClientUsecase(dataStore dbcluster.DataStore) *ClientUsecase {
	return &ClientUsecase{
		DataStore: dataStore,
	}
}

func (cu *ClientUsecase) GetUserByToken(ctx context.Context, claims *entity.JwtClaims) (*entity.User, error) {
	user, err := cu.DataStore.DSRepository().UserRepository.GetByIDAndEmail(cu.DataStore.DBInstanceClient(ctx), claims.UserID, claims.Name)
	if err != nil {
		return &entity.User{}, err
	}
	return entity.NewUser(&user), nil
}
