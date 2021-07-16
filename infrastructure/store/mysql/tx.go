package mysql

import (
	"context"

	"gorm.io/gorm"
)

type txKey struct{}

type Transaction struct {
	db *gorm.DB
}

// GetTx Contextからtxを取得する
func GetTx(ctx context.Context) (*gorm.DB, bool) {
	tx, ok := ctx.Value(txKey{}).(*gorm.DB)
	return tx, ok
}

func ProvideTransaction(db *gorm.DB) *Transaction {
	return &Transaction{db: db}
}

func (t *Transaction) Transaction(ctx context.Context, f func(ctx context.Context) error) error {
	return t.db.Transaction(func(tx *gorm.DB) error {
		return f(ctx)
	})
}
