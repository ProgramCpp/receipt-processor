package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccess(t *testing.T){
	db := NewMemDb()

	id := db.Put("hello")
	assert.NotEmpty(t, id)

	v, found := db.Get(id)
	assert.True(t, found)
	assert.Equal(t, "hello", v)
}

func TestFailure(t *testing.T){
	db := NewMemDb()

	id := db.Put("hello")
	assert.NotEmpty(t, id)

	v, found := db.Get("x")
	assert.False(t, found)
	assert.NotEqual(t, "hello", v)
}