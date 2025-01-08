package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/howood/moggiecollector/application/actor"
	"github.com/howood/moggiecollector/di/dbcluster"
	"github.com/howood/moggiecollector/domain/dto"
	"github.com/howood/moggiecollector/domain/entity"
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

func (au *AccountUsecase) GetUsers(ctx context.Context, withinactive string) ([]*entity.User, error) {
	if withinactive == "true" {
		users, err := au.DataStore.DSRepository().UserRepository.GetAllWithInActive(au.DataStore.DBInstanceClient(ctx))
		return au.convertToEntityUsers(users), err
	}
	users, err := au.DataStore.DSRepository().UserRepository.GetAll(au.DataStore.DBInstanceClient(ctx))
	return au.convertToEntityUsers(users), err
}

func (au *AccountUsecase) GetUser(ctx context.Context, userid uuid.UUID) (*entity.User, error) {
	user, err := au.DataStore.DSRepository().UserRepository.Get(au.DataStore.DBInstanceClient(ctx), userid)
	if err != nil {
		return &entity.User{}, err
	}
	return entity.NewUser(&user), nil
}

func (au *AccountUsecase) CreateUser(ctx context.Context, userDto *dto.UserDto) error {
	user, err := au.createUser(ctx, userDto)
	if err != nil {
		return err
	}
	return au.DataStore.DBInstanceClient(ctx).Transaction(func(tx *gorm.DB) error {
		if _, err := au.DataStore.DSRepository().UserRepository.GetByEmail(tx, userDto.Email); err != nil && !au.DataStore.RecordNotFoundError(err) {
			return err
		}
		if err == nil {
			//nolint:err113
			return errors.New("exist user with requested email")
		}
		return au.DataStore.DSRepository().UserRepository.Create(tx, &user)
	})
}

func (au *AccountUsecase) UpdateUser(ctx context.Context, userid uuid.UUID, userDto *dto.UserDto) error {
	user, err := au.createUser(ctx, userDto)
	if err != nil {
		return err
	}
	user.UserID = userid
	return au.DataStore.DBInstanceClient(ctx).Transaction(func(tx *gorm.DB) error {
		return au.DataStore.DSRepository().UserRepository.Update(tx, &user)
	})
}

func (au *AccountUsecase) InActiveUser(ctx context.Context, userid uuid.UUID) error {
	return au.DataStore.DBInstanceClient(ctx).Transaction(func(tx *gorm.DB) error {
		return au.DataStore.DSRepository().UserRepository.InActive(tx, userid)
	})
}

func (au *AccountUsecase) AuthUser(ctx context.Context, loginDto *dto.LoginDto) (*entity.User, error) {
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
func (au *AccountUsecase) comparePassword(user model.User, password string) error {
	return actor.PasswordOperator{}.ComparePassword(user.Password, password, user.Salt)
}

func (au *AccountUsecase) createUser(ctx context.Context, userDto *dto.UserDto) (model.User, error) {
	hashedpassword, salt, err := actor.PasswordOperator{}.GetHashedPassword(ctx, userDto.Password)
	if err != nil {
		return model.User{}, err
	}
	return model.User{
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: hashedpassword,
		Salt:     salt,
	}, nil
}

func (au *AccountUsecase) convertToEntityUsers(users []model.User) []*entity.User {
	entityUsers := make([]*entity.User, 0)

	for _, user := range users {
		entityUsers = append(entityUsers, &entity.User{
			UserID: user.UserID,
			Name:   user.Name,
			Email:  user.Email,
			Status: user.Status,
		})
	}

	return entityUsers
}
