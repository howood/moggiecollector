package testtools

import (
	"testing"

	"github.com/howood/moggiecollector/infrastructure/client"
	"gorm.io/gorm"
)

func DBTx(t *testing.T) *gorm.DB {
	t.Helper()

	dataaccessor := client.NewDataStoreAccesser()

	tx := dataaccessor.Instance.GetClient().WithContext(t.Context()).Begin()
	if tx.Error != nil {
		t.Fatalf("Failed to start a transaction: %v", tx.Error)
	}

	t.Cleanup(func() {
		if r := tx.Rollback(); r.Error != nil {
			t.Fatalf("Failed to rollback the transaction: %v", r.Error)
		}
	})

	return tx
}
