package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/howood/moggiecollector/application/usecase"
	"github.com/howood/moggiecollector/di/dbcluster"
	"github.com/howood/moggiecollector/domain/dto"
	"github.com/howood/moggiecollector/domain/entity"
	"github.com/howood/moggiecollector/domain/model"
	"github.com/howood/moggiecollector/library/testtools"
	"github.com/stretchr/testify/assert"
)

//nolint:funlen
func TestAccountUsecase_GetUsers(t *testing.T) {
	t.Parallel()

	now := time.Now()
	tests := []struct {
		name    string
		want    []*entity.User
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "正常系: 同じデータが取得できる",
			want: []*entity.User{
				{
					UserID: uuid.MustParse("dc059ab8-5569-492f-8229-939b7de055dc"),
					Name:   "xxxxxxx",
					Email:  "xxxxxxx",
					Status: 0,
				},
				{
					UserID: uuid.MustParse("64d9eee6-69b6-4a44-8980-55470a424434"),
					Name:   "xxxxxxx2",
					Email:  "xxxxxxx2",
					Status: 0,
				},
			},
			wantErr: assert.NoError,
		},
	}

	initialData := []*model.User{
		{
			UserID:    uuid.MustParse("dc059ab8-5569-492f-8229-939b7de055dc"),
			Name:      "xxxxxxx",
			Email:     "xxxxxxx",
			Password:  "xxxxxxx",
			Salt:      "xxxxxxx",
			Status:    0,
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			UserID:    uuid.MustParse("64d9eee6-69b6-4a44-8980-55470a424434"),
			Name:      "xxxxxxx2",
			Email:     "xxxxxxx2",
			Password:  "xxxxxxx2",
			Salt:      "xxxxxxx2",
			Status:    0,
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			tx := testtools.DBTx(t)
			if err := tx.Create(initialData).Error; err != nil {
				t.Fatal(err)
			}
			dataStore := dbcluster.NewDatastoreForTest(tx)

			au := usecase.NewAccountUsecase(dataStore)

			got, err := au.GetUsers(ctx, "true")
			if !tt.wantErr(t, err) {
				t.Fatal(t, err)

				return
			}
			assert.Equal(t, tt.want[0].UserID, got[0].UserID)
			assert.Equal(t, tt.want[0].Name, got[0].Name)
			assert.Equal(t, tt.want[0].Email, got[0].Email)
			assert.Equal(t, tt.want[0].Status, got[0].Status)

			assert.Equal(t, tt.want[1].UserID, got[1].UserID)
			assert.Equal(t, tt.want[1].Name, got[1].Name)
			assert.Equal(t, tt.want[1].Email, got[1].Email)
			assert.Equal(t, tt.want[1].Status, got[1].Status)
		})
	}
}

//nolint:funlen
func TestAccountUsecase_GetUser(t *testing.T) {
	t.Parallel()

	type args struct {
		id uuid.UUID
	}

	now := time.Now()
	tests := []struct {
		name    string
		args    args
		want    *entity.User
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "正常系: データが取得できる",
			args: args{
				id: uuid.MustParse("dc059ab8-5569-492f-8229-939b7de055dc"),
			},
			want: &entity.User{
				UserID: uuid.MustParse("dc059ab8-5569-492f-8229-939b7de055dc"),
				Name:   "xxxxxxx",
				Email:  "xxxxxxx",
				Status: 0,
			},
			wantErr: assert.NoError,
		},
		{
			name: "正常系: データが取得できない",
			args: args{
				id: uuid.MustParse("dc059ab8-5569-492f-8229-939b7de055de"),
			},
			want: &entity.User{},
			wantErr: func(t assert.TestingT, err error, _ ...interface{}) bool {
				return assert.EqualError(t, err, dbcluster.RecordNotFoundMsg)
			},
		},
	}

	initialData := []*model.User{
		{
			UserID:    uuid.MustParse("dc059ab8-5569-492f-8229-939b7de055dc"),
			Name:      "xxxxxxx",
			Email:     "xxxxxxx",
			Password:  "xxxxxxx",
			Salt:      "xxxxxxx",
			Status:    0,
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			UserID:    uuid.MustParse("64d9eee6-69b6-4a44-8980-55470a424434"),
			Name:      "xxxxxxx2",
			Email:     "xxxxxxx2",
			Password:  "xxxxxxx2",
			Salt:      "xxxxxxx2",
			Status:    0,
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			tx := testtools.DBTx(t)
			if err := tx.Create(initialData).Error; err != nil {
				t.Fatal(err)
			}
			dataStore := dbcluster.NewDatastoreForTest(tx)

			au := usecase.NewAccountUsecase(dataStore)

			got, err := au.GetUser(ctx, tt.args.id)
			if !tt.wantErr(t, err) || err != nil {
				return
			}

			assert.Equal(t, tt.want.UserID, got.UserID)
			assert.Equal(t, tt.want.Name, got.Name)
			assert.Equal(t, tt.want.Email, got.Email)
			assert.Equal(t, tt.want.Status, got.Status)
		})
	}
}

func TestAccountUsecase_CreateUser(t *testing.T) {
	t.Parallel()

	type args struct {
		input *dto.UserDto
	}

	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "正常系: データが登録できる",
			args: args{
				input: &dto.UserDto{
					Name:     "xxxxxxx",
					Email:    "xxxxxxx",
					Password: "xxxxxxx",
				},
			},

			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			tx := testtools.DBTx(t)
			dataStore := dbcluster.NewDatastoreForTest(tx)
			au := usecase.NewAccountUsecase(dataStore)

			err := au.CreateUser(ctx, tt.args.input)
			if !tt.wantErr(t, err) || err != nil {
				return
			}
		})
	}
}
