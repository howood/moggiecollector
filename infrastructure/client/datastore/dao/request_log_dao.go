package dao

import (
	"github.com/howood/moggiecollector/domain/model"
	"github.com/howood/moggiecollector/domain/repository"
	"gorm.io/gorm"
)

// SNidkTransactionIdDao struct.
type RequestLogDao struct{}

// NewAuthorityGroupDao creates a new AuthorityGroupRepository.
//
//nolint:ireturn,nolintlint
func NewRequestLogDao() repository.RequestLogRepository {
	return &RequestLogDao{}
}

func (u *RequestLogDao) Create(db *gorm.DB, requestLog model.RequestLog) error {
	return db.Create(&requestLog).Error
}
