package models

import (
	"strings"
	"time"

	"github.com/aaa59891/go_empty_gin/constants"
	"github.com/aaa59891/go_empty_gin/db"
	"github.com/jinzhu/gorm"
)

type Transaction func(tx *gorm.DB) error

func Transactional(ts ...func(db2 *gorm.DB) error) error {
	if ts == nil || len(ts) == 0 {
		return constants.ErrNoTransaction
	}
	tx := db.DB.Begin()

	for _, t := range ts {
		if err := t(tx); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func SqlLikeString(str string) string {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return ""
	}
	return "%" + str + "%"
}

func GetFilterDateValue(startDate, endDate time.Time) (time.Time, time.Time) {
	if startDate.IsZero() {
		startDate = endDate
	}
	if endDate.IsZero() {
		endDate = startDate
	}
	endDate = endDate.AddDate(0, 0, 1)
	return startDate, endDate
}
