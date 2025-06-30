package model

const (
	HTTPTypeRequest  = "request"
	HTTPTypeResponse = "response"
)

type RequestLog struct {
	BaseModel
	XRequestID string
	Endpoint   string
	Method     string
	HTTPType   string
	URLQuery   *string
	Body       *string
	Header     string
}
