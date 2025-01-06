package usecase

import (
	"context"

	"github.com/howood/moggiecollector/di/dbcluster"
	"github.com/howood/moggiecollector/domain/entity"
	"github.com/howood/moggiecollector/domain/model"
)

type ClientUsecase struct {
	DataStore dbcluster.DataStore
}

func NewClientUsecase(dataStore dbcluster.DataStore) *ClientUsecase {
	return &ClientUsecase{
		DataStore: dataStore,
	}
}

func (cu *ClientUsecase) GetUserByToken(ctx context.Context, claims *entity.JwtClaims) (model.User, error) {
	return cu.DataStore.DSRepository().UserRepository.GetByIDAndEmail(cu.DataStore.DBInstanceClient(ctx), claims.UserID, claims.Name)
}
