package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/howood/moggiecollector/di/dbcluster"
	"github.com/howood/moggiecollector/domain/model"
	"github.com/labstack/echo/v4"
)

type RequestLogService struct {
	DataStore dbcluster.DataStore
}

func NewRequestLogService(dataStore dbcluster.DataStore) *RequestLogService {
	return &RequestLogService{
		DataStore: dataStore,
	}
}

func (rs *RequestLogService) CreateRequest(ctx context.Context, c echo.Context) error {
	requestLog := model.RequestLog{
		XRequestID: fmt.Sprintf("%v", c.Get(echo.HeaderXRequestID)),
		Endpoint:   c.Request().URL.RequestURI(),
		Method:     c.Request().Method,
		HTTPType:   model.HTTPTypeRequest,
		URLQuery:   &c.Request().URL.RawQuery,
		Body:       rs.readRequestBodyPrev(c),
		Header:     fmt.Sprintf("%v", c.Request().Header),
	}
	//nolint:wrapcheck
	cli := rs.DataStore.DBInstanceClient(ctx)
	return rs.DataStore.DSRepository().RequestLogRepository.Create(cli, requestLog)
}

func (rs *RequestLogService) CreateResponse(ctx context.Context, c echo.Context, response interface{}) error {
	res := fmt.Sprintf("%v", response)
	requestLog := model.RequestLog{
		XRequestID: fmt.Sprintf("%v", c.Get(echo.HeaderXRequestID)),
		Endpoint:   c.Request().URL.RequestURI(),
		Method:     c.Request().Method,
		HTTPType:   model.HTTPTypeResponse,
		URLQuery:   &c.Request().URL.RawQuery,
		Body:       &res,
		Header:     fmt.Sprintf("%v", c.Request().Header),
	}
	//nolint:wrapcheck
	return rs.DataStore.DSRepository().RequestLogRepository.Create(rs.DataStore.DBInstanceClient(ctx), requestLog)
}

func (rs *RequestLogService) readRequestBodyPrev(c echo.Context) *string {
	if c.Request().Body == nil || c.Request().Body == http.NoBody {
		return nil
	}

	var b bytes.Buffer
	_, _ = b.ReadFrom(c.Request().Body)
	c.Request().Body = io.NopCloser(&b)

	body := b.String()

	return &body
}
