package usecase

import (
	"errors"

	"github.com/howood/moggiecollector/domain/entity"
	"github.com/howood/moggiecollector/interfaces/config"
)

type AccountUsecase struct {
}

func (au AccountUsecase) GetUsers(withinactive string) ([]entity.User, error) {
	if withinactive == "true" {
		return config.GetDataStore().User.GetAllWithInActive()
	} else {
		return config.GetDataStore().User.GetAll()
	}
}

func (au AccountUsecase) GetUser(userid int) (entity.User, error) {
	return config.GetDataStore().User.Get(uint64(userid))
}

func (au AccountUsecase) CreateUser(form entity.CreateUserForm) error {
	_, err := config.GetDataStore().User.GetByEmail(form.Email)
	if err != nil && !config.GetDataStore().User.RecordNotFoundError(err) {
		return err
	}
	if err == nil {
		return errors.New("exist user with requested email")
	}
	return config.GetDataStore().User.Create(form.Name, form.Email, form.Password)
}

func (au AccountUsecase) UpdateUser(userid int, form entity.CreateUserForm) error {
	return config.GetDataStore().User.Update(uint64(userid), form.Name, form.Email, form.Password)
}

func (au AccountUsecase) InActiveUser(userid int) error {
	return config.GetDataStore().User.InActive(uint64(userid))
}

func (au AccountUsecase) AuthUser(form entity.LoginUserForm) (entity.User, error) {
	return config.GetDataStore().User.Auth(form.Email, form.Password)
}
