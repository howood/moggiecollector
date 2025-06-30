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

type UserUsecase struct {
	DataStore dbcluster.DataStore
}

func NewUserUsecase(dataStore dbcluster.DataStore) *UserUsecase {
	return &UserUsecase{
		DataStore: dataStore,
	}
}

func (uu *UserUsecase) GetUsers(ctx context.Context, withinactive string) ([]*entity.User, error) {
	if withinactive == "true" {
		users, err := uu.DataStore.DSRepository().UserRepository.GetAllWithInActive(uu.DataStore.DBInstanceClient(ctx))
		return uu.convertToEntityUsers(users), err
	}
	users, err := uu.DataStore.DSRepository().UserRepository.GetAll(uu.DataStore.DBInstanceClient(ctx))
	return uu.convertToEntityUsers(users), err
}

func (uu *UserUsecase) GetUser(ctx context.Context, userid uuid.UUID) (*entity.User, error) {
	user, err := uu.DataStore.DSRepository().UserRepository.Get(uu.DataStore.DBInstanceClient(ctx), userid)
	if err != nil {
		return &entity.User{}, err
	}
	return entity.NewUser(&user), nil
}

func (uu *UserUsecase) CreateUser(ctx context.Context, userDto *dto.UserDto) error {
	user, err := uu.createUser(ctx, userDto)
	if err != nil {
		return err
	}
	return uu.DataStore.DBInstanceClient(ctx).Transaction(func(tx *gorm.DB) error {
		_, err := uu.DataStore.DSRepository().UserRepository.GetByEmail(tx, userDto.Email)
		if err != nil && !uu.DataStore.RecordNotFoundError(err) {
			return err
		}
		if err == nil {
			//nolint:err113
			return errors.New("exist user with requested email")
		}
		return uu.DataStore.DSRepository().UserRepository.Create(tx, &user)
	})
}

func (uu *UserUsecase) UpdateUser(ctx context.Context, userid uuid.UUID, userDto *dto.UserDto) error {
	user, err := uu.createUser(ctx, userDto)
	if err != nil {
		return err
	}
	user.ID = userid
	return uu.DataStore.DBInstanceClient(ctx).Transaction(func(tx *gorm.DB) error {
		return uu.DataStore.DSRepository().UserRepository.Update(tx, &user)
	})
}

func (uu *UserUsecase) InActiveUser(ctx context.Context, userid uuid.UUID) error {
	return uu.DataStore.DBInstanceClient(ctx).Transaction(func(tx *gorm.DB) error {
		return uu.DataStore.DSRepository().UserRepository.InActive(tx, userid)
	})
}

func (uu *UserUsecase) GetUserByToken(ctx context.Context, claims *entity.JwtClaims) (*entity.User, error) {
	user, err := uu.DataStore.DSRepository().UserRepository.GetByIDAndEmail(uu.DataStore.DBInstanceClient(ctx), claims.UserID, claims.Name)
	if err != nil {
		return &entity.User{}, err
	}
	return entity.NewUser(&user), nil
}

func (uu *UserUsecase) createUser(ctx context.Context, userDto *dto.UserDto) (model.User, error) {
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

func (uu *UserUsecase) convertToEntityUsers(users []model.User) []*entity.User {
	entityUsers := make([]*entity.User, 0)

	for _, user := range users {
		entityUsers = append(entityUsers, &entity.User{
			ID:     user.ID,
			Name:   user.Name,
			Email:  user.Email,
			Status: user.Status,
		})
	}

	return entityUsers
}
