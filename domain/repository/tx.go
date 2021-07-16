package repository

import (
	"context"
)

type Tx interface {
	Transaction(ctx context.Context, f func(ctx context.Context) error) error
}
