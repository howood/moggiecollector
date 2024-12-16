package usecase

import (
	"github.com/howood/moggiecollector/di"
	"github.com/howood/moggiecollector/domain/entity"
)

type ClientUsecase struct {
}

func (au ClientUsecase) GetUserByToken(claims *entity.JwtClaims) (entity.User, error) {
	return di.GetDataStore().User.GetByIDAndEmail(claims.UserID, claims.Name)
}
