package model

import (
	"time"

	"github.com/google/uuid"
)

const (
	HTTPTypeRequest  = "request"
	HTTPTypeResponse = "response"
)

type RequestLog struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;size:255;default:uuid_generate_v4()"`
	XRequestID string
	Endpoint   string
	Method     string
	HTTPType   string
	URLQuery   *string
	Body       *string
	Header     string
	CreatedAt  time.Time
}
