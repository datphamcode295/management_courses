package repo

import (
	"context"

	"gorm.io/gorm"
)

// NewTestStore .
func NewTestStore() DBRepo {
	return &testStore{}
}

type testStore struct {
}

func (*testStore) DB() *gorm.DB {
	return nil
}

func (s *testStore) NewTransaction() (DBRepo, FinallyFunc) {
	return s, func(err error) error { return err }
}

func (s *testStore) NewTransactionWithContext(ctx context.Context) (DBRepo, FinallyFunc) {
	return s, func(err error) error { return err }
}
