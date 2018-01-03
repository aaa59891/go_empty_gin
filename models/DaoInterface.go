package models

import (
	"github.com/aaa59891/go_empty_gin/constants"
	"github.com/jinzhu/gorm"
)

type Dao interface {
	Insert(tx *gorm.DB) error
	Update(tx *gorm.DB) error
	Delete(tx *gorm.DB) error
}

func InsertAndRecord(d Dao, user, from string) Transaction {
	return func(tx *gorm.DB) error {
		if err := d.Insert(tx); err != nil {
			return err
		}
		return nil
	}
}

func UpdateAndRecord(d Dao, user, from string, actions ...string) Transaction {
	return func(tx *gorm.DB) error {
		if tx.NewRecord(d) {
			return gorm.ErrRecordNotFound
		}

		if err := d.Update(tx); err != nil {
			return err
		}
		return nil
	}
}

func DeleteAndRecord(d Dao, user, from string) Transaction {
	return func(tx *gorm.DB) error {
		if err := d.Delete(tx); err != nil {
			return err
		}
		return nil
	}
}

/* Special -- not use dao interface */
func DeleteById(model, id interface{}, user, from string) Transaction {
	return func(tx *gorm.DB) error {
		if tx.NewRecord(model) {
			return constants.ErrNoId
		}

		if err := tx.Delete(model, "id = ?", id).Error; err != nil {
			return err
		}
		return nil
	}
}
