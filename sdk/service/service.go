package service

import (
	"fmt"

	"git.tz.com/devops/gin-core/logger"
	"git.tz.com/devops/gin-core/storage"
	"gorm.io/gorm"
)

type Service struct {
	Orm   *gorm.DB
	Msg   string
	MsgID string
	Log   *logger.Helper
	Error error
	Cache storage.AdapterCache
}

func (db *Service) AddError(err error) error {
	if db.Error == nil {
		db.Error = err
	} else if err != nil {
		db.Error = fmt.Errorf("%v; %w", db.Error, err)
	}
	return db.Error
}