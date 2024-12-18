package usecase

import (
	"context"

	"github.com/howood/moggiecollector/di"
	"github.com/howood/moggiecollector/domain/entity"
)

type ClientUsecase struct {
	Ctx context.Context
}

func (au ClientUsecase) GetUserByToken(claims *entity.JwtClaims) (entity.User, error) {
	return di.GetDataStore().User.GetByIDAndEmail(au.Ctx, claims.UserID, claims.Name)
}
