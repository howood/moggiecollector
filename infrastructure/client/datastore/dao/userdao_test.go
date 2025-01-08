package dao_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/howood/moggiecollector/domain/model"
	"github.com/howood/moggiecollector/infrastructure/client/datastore/dao"
	"github.com/howood/moggiecollector/library/testtools"
	"github.com/stretchr/testify/assert"
)

//nolint:funlen
func TestUserDao_GetAll(t *testing.T) {
	t.Parallel()

	now := time.Now()
	tests := []struct {
		name    string
		want    []*model.User
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "正常系: 同じデータが取得できる",
			want: []*model.User{
				{
					UserID:   uuid.MustParse("dc059ab8-5569-492f-8229-939b7de055dc"),
					Name:     "xxxxxxx",
					Email:    "xxxxxxx",
					Password: "xxxxxxx",
					Salt:     "xxxxxxx",
					Status:   0,
				},
				{
					UserID:   uuid.MustParse("64d9eee6-69b6-4a44-8980-55470a424434"),
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
			UserID:   uuid.MustParse("dc059ab8-5569-492f-8229-939b7de055dc"),
			Name:     "xxxxxxx",
			Email:    "xxxxxxx",
			Password: "xxxxxxx",
			Salt:     "xxxxxxx",
			Status:   0,
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
			assert.Equal(t, tt.want[0].UserID, got[0].UserID)
			assert.Equal(t, tt.want[0].Name, got[0].Name)
			assert.Equal(t, tt.want[0].Email, got[0].Email)
			assert.Equal(t, tt.want[0].Password, got[0].Password)
			assert.Equal(t, tt.want[0].Salt, got[0].Salt)
			assert.Equal(t, tt.want[0].Status, got[0].Status)

			assert.Equal(t, tt.want[1].UserID, got[1].UserID)
			assert.Equal(t, tt.want[1].Name, got[1].Name)
			assert.Equal(t, tt.want[1].Email, got[1].Email)
			assert.Equal(t, tt.want[1].Password, got[1].Password)
			assert.Equal(t, tt.want[1].Salt, got[1].Salt)
			assert.Equal(t, tt.want[1].Status, got[1].Status)
		})
	}
}
