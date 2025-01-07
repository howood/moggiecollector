package usecase

import (
	"context"
	"errors"

	"github.com/howood/moggiecollector/application/actor"
	"github.com/howood/moggiecollector/di/dbcluster"
	"github.com/howood/moggiecollector/domain/form"
	"github.com/howood/moggiecollector/domain/model"
	"gorm.io/gorm"
)

type AccountUsecase struct {
	DataStore dbcluster.DataStore
}

func NewAccountUsecase(dataStore dbcluster.DataStore) *AccountUsecase {
	return &AccountUsecase{
		DataStore: dataStore,
	}
}

func (au *AccountUsecase) GetUsers(ctx context.Context, withinactive string) ([]model.User, error) {
	if withinactive == "true" {
		return au.DataStore.DSRepository().UserRepository.GetAllWithInActive(au.DataStore.DBInstanceClient(ctx))
	}
	return au.DataStore.DSRepository().UserRepository.GetAll(au.DataStore.DBInstanceClient(ctx))
}

func (au *AccountUsecase) GetUser(ctx context.Context, userid int) (model.User, error) {
	//nolint:gosec
	return au.DataStore.DSRepository().UserRepository.Get(au.DataStore.DBInstanceClient(ctx), uint64(userid))
}

func (au *AccountUsecase) CreateUser(ctx context.Context, form form.CreateUserForm) error {
	user, err := au.createUser(ctx, form)
	if err != nil {
		return err
	}
	return au.DataStore.DBInstanceClient(ctx).Transaction(func(tx *gorm.DB) error {
		if _, err := au.DataStore.DSRepository().UserRepository.GetByEmail(tx, form.Email); err != nil && !au.DataStore.RecordNotFoundError(err) {
			return err
		}
		if err == nil {
			//nolint:err113
			return errors.New("exist user with requested email")
		}
		return au.DataStore.DSRepository().UserRepository.Create(tx, &user)
	})
}

func (au *AccountUsecase) UpdateUser(ctx context.Context, userid int, form form.CreateUserForm) error {
	user, err := au.createUser(ctx, form)
	if err != nil {
		return err
	}
	//nolint:gosec
	user.UserID = uint64(userid)
	return au.DataStore.DBInstanceClient(ctx).Transaction(func(tx *gorm.DB) error {
		return au.DataStore.DSRepository().UserRepository.Update(tx, &user)
	})
}

func (au *AccountUsecase) InActiveUser(ctx context.Context, userid int) error {
	return au.DataStore.DBInstanceClient(ctx).Transaction(func(tx *gorm.DB) error {
		//nolint:gosec
		return au.DataStore.DSRepository().UserRepository.InActive(tx, uint64(userid))
	})
}

func (au *AccountUsecase) AuthUser(ctx context.Context, form form.LoginUserForm) (model.User, error) {
	user, err := au.DataStore.DSRepository().UserRepository.GetByEmail(au.DataStore.DBInstanceClient(ctx), form.Email)
	if err != nil {
		return model.User{}, err
	}
	if err := au.comparePassword(user, form.Password); err != nil {
		return model.User{}, err
	}
	return user, nil
}

// ComparePassword compares input password to roomdata password
func (au *AccountUsecase) comparePassword(user model.User, password string) error {
	return actor.PasswordOperator{}.ComparePassword(user.Password, password, user.Salt)
}

func (au *AccountUsecase) createUser(ctx context.Context, form form.CreateUserForm) (model.User, error) {
	hashedpassword, salt, err := actor.PasswordOperator{}.GetHashedPassword(ctx, form.Password)
	if err != nil {
		return model.User{}, err
	}
	return model.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: hashedpassword,
		Salt:     salt,
	}, nil
}
