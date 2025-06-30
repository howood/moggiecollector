package dao_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/howood/moggiecollector/di/dbcluster"
	"github.com/howood/moggiecollector/domain/model"
	"github.com/howood/moggiecollector/infrastructure/client/datastore/dao"
	"github.com/howood/moggiecollector/library/testtools"
	"github.com/stretchr/testify/assert"
)

//nolint:funlen,tparallel
func TestUserDao_GetAll(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		want    []*model.User
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "正常系: 同じデータが取得できる",
			want: []*model.User{
				{
					BaseModel: model.BaseModel{
						ID: uuid.MustParse("dc059ab8-5569-492f-8229-939b7de055dc"),
					},
					Name:     "xxxxxxx",
					Email:    "xxxxxxx",
					Password: "xxxxxxx",
					Salt:     "xxxxxxx",
					Status:   0,
				},
				{
					BaseModel: model.BaseModel{
						ID: uuid.MustParse("64d9eee6-69b6-4a44-8980-55470a424434"),
					},
					Name:     "xxxxxxx2",
					Email:    "xxxxxxx2",
					Password: "xxxxxxx2",
					Salt:     "xxxxxxx2",
					Status:   0,
				},
			},
			wantErr: assert.NoError,
		},
	}

	initialData := []*model.User{
		{
			BaseModel: model.BaseModel{
				ID: uuid.MustParse("dc059ab8-5569-492f-8229-939b7de055dc"),
			},
			Name:     "xxxxxxx",
			Email:    "xxxxxxx",
			Password: "xxxxxxx",
			Salt:     "xxxxxxx",
			Status:   0,
		},
		{
			BaseModel: model.BaseModel{
				ID: uuid.MustParse("64d9eee6-69b6-4a44-8980-55470a424434"),
			},
			Name:     "xxxxxxx2",
			Email:    "xxxxxxx2",
			Password: "xxxxxxx2",
			Salt:     "xxxxxxx2",
			Status:   0,
		},
	}

	//nolint:paralleltest
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := testtools.DBTx(t)
			if err := tx.Create(initialData).Error; err != nil {
				t.Fatal(err)
			}

			ud := dao.NewUserDao()

			got, err := ud.GetAll(tx)
			if !tt.wantErr(t, err) {
				t.Fatal(t, err)

				return
			}
			for _, gotuser := range got {
				for _, wantuser := range tt.want {
					if gotuser.ID == wantuser.ID {
						assert.Equal(t, wantuser.Name, gotuser.Name)
						assert.Equal(t, wantuser.Email, gotuser.Email)
						assert.Equal(t, wantuser.Password, gotuser.Password)
						assert.Equal(t, wantuser.Salt, gotuser.Salt)
						assert.Equal(t, wantuser.Status, gotuser.Status)
					}
				}
			}
		})
	}
}

//nolint:funlen,tparallel
func TestUserDao_Get(t *testing.T) {
	t.Parallel()

	type args struct {
		id uuid.UUID
	}

	tests := []struct {
		name    string
		args    args
		want    *model.User
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "正常系: データが取得できる",
			args: args{
				id: uuid.MustParse("dc059ab8-5569-492f-8229-939b7de055dc"),
			},
			want: &model.User{
				BaseModel: model.BaseModel{
					ID: uuid.MustParse("dc059ab8-5569-492f-8229-939b7de055dc"),
				},
				Name:     "xxxxxxx",
				Email:    "xxxxxxx",
				Password: "xxxxxxx",
				Salt:     "xxxxxxx",
				Status:   0,
			},
			wantErr: assert.NoError,
		},
		{
			name: "正常系: データが取得できない",
			args: args{
				id: uuid.MustParse("dc059ab8-5569-492f-8229-939b7de055de"),
			},
			want: &model.User{},
			wantErr: func(t assert.TestingT, err error, _ ...interface{}) bool {
				return assert.EqualError(t, err, dbcluster.RecordNotFoundMsg)
			},
		},
	}

	initialData := []*model.User{
		{
			BaseModel: model.BaseModel{
				ID: uuid.MustParse("dc059ab8-5569-492f-8229-939b7de055dc"),
			},
			Name:     "xxxxxxx",
			Email:    "xxxxxxx",
			Password: "xxxxxxx",
			Salt:     "xxxxxxx",
			Status:   0,
		},
		{
			BaseModel: model.BaseModel{
				ID: uuid.MustParse("64d9eee6-69b6-4a44-8980-55470a424434"),
			},
			Name:     "xxxxxxx2",
			Email:    "xxxxxxx2",
			Password: "xxxxxxx2",
			Salt:     "xxxxxxx2",
			Status:   0,
		},
	}

	//nolint:paralleltest
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := testtools.DBTx(t)
			if err := tx.Create(initialData).Error; err != nil {
				t.Fatal(err)
			}

			ud := dao.NewUserDao()

			got, err := ud.Get(tx, tt.args.id)
			if !tt.wantErr(t, err) || err != nil {
				return
			}
			assert.Equal(t, tt.want.ID, got.ID)
			assert.Equal(t, tt.want.Name, got.Name)
			assert.Equal(t, tt.want.Email, got.Email)
			assert.Equal(t, tt.want.Password, got.Password)
			assert.Equal(t, tt.want.Salt, got.Salt)
			assert.Equal(t, tt.want.Status, got.Status)
		})
	}
}
