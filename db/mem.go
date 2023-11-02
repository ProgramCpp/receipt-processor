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

func (m *MemDb)Put(i interface{}) string{
	pk := uuid.New().String()

	m.store[pk] = i

	return pk
}

func (m *MemDb)Get(k string) (interface{}, bool){
	v, found := m.store[k]
	return v, found
}