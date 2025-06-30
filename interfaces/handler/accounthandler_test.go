package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/howood/moggiecollector/di/dbcluster"
	"github.com/howood/moggiecollector/di/uccluster"
	"github.com/howood/moggiecollector/domain/model"
	"github.com/howood/moggiecollector/interfaces/handler"
	"github.com/howood/moggiecollector/interfaces/handler/response"
	"github.com/howood/moggiecollector/library/testtools"
	"github.com/howood/moggiecollector/library/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

//nolint:funlen,tparallel
func TestAccountHandler_GetUser(t *testing.T) {
	t.Parallel()

	type args struct {
		id uuid.UUID
	}

	tests := []struct {
		name       string
		args       args
		wantStatus int
		want       response.UserResponse
	}{
		{
			name: "正常系: データが取得できる",
			args: args{
				id: uuid.MustParse("dc059ab8-5569-492f-8229-939b7de055dc"),
			},
			wantStatus: http.StatusOK,
			want: response.UserResponse{
				ID:     uuid.MustParse("dc059ab8-5569-492f-8229-939b7de055dc"),
				Name:   "xxxxxxx",
				Email:  "xxxxxxx",
				Status: 0,
			},
		},
		{
			name: "正常系: 存在しないデータは取得できない",
			args: args{
				id: uuid.MustParse("6b165cc6-7580-4a6d-9080-f2dc4e9db154"),
			},
			wantStatus: http.StatusNotFound,
			want:       response.UserResponse{},
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
			dataStore := dbcluster.NewDatastoreForTest(tx)
			uccluster := uccluster.NewUsecaseCluster(dataStore)
			baseHandler := handler.BaseHandler{UcCluster: uccluster}

			if err := tx.Create(initialData).Error; err != nil {
				t.Fatal(err)
			}

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users/" + tt.args.id.String())
			c.SetParamNames("id")
			c.SetParamValues(tt.args.id.String())

			h := &handler.AccountHandler{BaseHandler: baseHandler}

			if assert.NoError(t, h.GetUser(c)) {
				assert.Equal(t, tt.wantStatus, rec.Code)

				response := response.UserResponse{}
				if err := utils.ByteToJSONStruct(rec.Body.Bytes(), &response); err != nil {
					t.Error(err)
				}

				assert.Equal(t, tt.want, response)
			}
		})
	}
}
