package requestid_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/howood/moggiecollector/infrastructure/requestid"
	"github.com/labstack/echo/v4"
)

func Test_RequestIDHandler(t *testing.T) {
	t.Parallel()

	checkvalue := "azxswedcvfr"
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set("X-Request-ID", checkvalue)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	e := echo.New()
	c := e.NewContext(req, httptest.NewRecorder())

	if xrequestid := requestid.GetRequestID(c.Request()); xrequestid != checkvalue {
		t.Fatal("failed test RequestIDHandler")
	}

	t.Log("success RequestIDHandler")
}

func Test_RequestIDHandlerCreateNew(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	e := echo.New()
	c := e.NewContext(req, httptest.NewRecorder())

	if xrequestid := requestid.GetRequestID(c.Request()); xrequestid == "" {
		t.Fatal("failed test RequestIDHandler Create New")
	}

	t.Log("success RequestIDHandler Create New")
}
