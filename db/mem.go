package db

import (
	"github.com/google/uuid"
)

type MemDb struct {
	store map[string]interface{}
}

func NewMemDb() *MemDb {
	return &MemDb{
		store: make(map[string]interface{}),
	}
}

func (m *MemDb)Insert(i interface{}) string{
	pk := uuid.New().String()

	m.store[pk] = i

	return pk
}