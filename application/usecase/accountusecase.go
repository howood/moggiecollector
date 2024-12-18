package usecase

import (
	"context"
	"errors"

	"github.com/howood/moggiecollector/di"
	"github.com/howood/moggiecollector/domain/entity"
)

type AccountUsecase struct {
	Ctx context.Context
}

func (au AccountUsecase) GetUsers(withinactive string) ([]entity.User, error) {
	if withinactive == "true" {
		return di.GetDataStore().User.GetAllWithInActive(au.Ctx)
	} else {
		return di.GetDataStore().User.GetAll(au.Ctx)
	}
}

func (au AccountUsecase) GetUser(userid int) (entity.User, error) {
	return di.GetDataStore().User.Get(au.Ctx, uint64(userid))
}

func (au AccountUsecase) CreateUser(form entity.CreateUserForm) error {
	_, err := di.GetDataStore().User.GetByEmail(au.Ctx, form.Email)
	if err != nil && !di.GetDataStore().User.RecordNotFoundError(err) {
		return err
	}
	if err == nil {
		return errors.New("exist user with requested email")
	}
	return di.GetDataStore().User.Create(au.Ctx, form.Name, form.Email, form.Password)
}

func (au AccountUsecase) UpdateUser(userid int, form entity.CreateUserForm) error {
	return di.GetDataStore().User.Update(au.Ctx, uint64(userid), form.Name, form.Email, form.Password)
}

func (au AccountUsecase) InActiveUser(userid int) error {
	return di.GetDataStore().User.InActive(au.Ctx, uint64(userid))
}

func (au AccountUsecase) AuthUser(form entity.LoginUserForm) (entity.User, error) {
	return di.GetDataStore().User.Auth(au.Ctx, form.Email, form.Password)
}
