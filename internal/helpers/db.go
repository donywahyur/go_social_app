package helpers

import (
	"fmt"

	"gorm.io/gorm"
)

func RunDBTransaction(db *gorm.DB, fn func(tx *gorm.DB) error) error {
	tx := db.Begin()

	if tx.Error != nil {
		return fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}

	defer func() {
		if tx.Error != nil {
			if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
				tx.Error = fmt.Errorf("rollback failed: %v, original error: %w", rollbackErr, tx.Error)
			}
		}
	}()

	err := fn(tx)
	if err != nil {
		tx.Error = err
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
