package usecase

import (
	"errors"

	"github.com/howood/moggiecollector/di"
	"github.com/howood/moggiecollector/domain/entity"
)

type AccountUsecase struct {
}

func (au AccountUsecase) GetUsers(withinactive string) ([]entity.User, error) {
	if withinactive == "true" {
		return di.GetDataStore().User.GetAllWithInActive()
	} else {
		return di.GetDataStore().User.GetAll()
	}
}

func (au AccountUsecase) GetUser(userid int) (entity.User, error) {
	return di.GetDataStore().User.Get(uint64(userid))
}

func (au AccountUsecase) CreateUser(form entity.CreateUserForm) error {
	_, err := di.GetDataStore().User.GetByEmail(form.Email)
	if err != nil && !di.GetDataStore().User.RecordNotFoundError(err) {
		return err
	}
	if err == nil {
		return errors.New("exist user with requested email")
	}
	return di.GetDataStore().User.Create(form.Name, form.Email, form.Password)
}

func (au AccountUsecase) UpdateUser(userid int, form entity.CreateUserForm) error {
	return di.GetDataStore().User.Update(uint64(userid), form.Name, form.Email, form.Password)
}

func (au AccountUsecase) InActiveUser(userid int) error {
	return di.GetDataStore().User.InActive(uint64(userid))
}

func (au AccountUsecase) AuthUser(form entity.LoginUserForm) (entity.User, error) {
	return di.GetDataStore().User.Auth(form.Email, form.Password)
}
