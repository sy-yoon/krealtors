package utils

import (
	"github.com/sy-yoon/krealtors/gms"
	"gorm.io/gorm"
)

func CheckError(tx *gorm.DB) error {
	if tx.Error != nil {
		gms.Logger.Error("DB", "SQL", tx.Error)
		return tx.Error
	}
	return nil
}