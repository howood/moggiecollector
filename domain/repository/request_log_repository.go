package repository

import (
	"github.com/howood/moggiecollector/domain/model"
	"gorm.io/gorm"
)

// AuthorityGroupRepository interface.
type RequestLogRepository interface {
	Create(db *gorm.DB, requestLog model.RequestLog) error
}
