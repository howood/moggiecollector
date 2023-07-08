package usecase

import (
	"github.com/howood/moggiecollector/domain/entity"
	"github.com/howood/moggiecollector/interfaces/config"
)

type ClientUsecase struct {
}

func (au ClientUsecase) GetUserByToken(claims *entity.JwtClaims) (entity.User, error) {
	return config.GetDataStore().User.GetByIDAndEmail(claims.UserID, claims.Name)
}
