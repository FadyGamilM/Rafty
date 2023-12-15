package raftycache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	raftyCache := New()
	err := raftyCache.Set([]byte("1"), []byte("fadyGamil"), 10)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	raftyCache := New()
	err := raftyCache.Set([]byte("1"), []byte("fadyGamil"), 10)
	assert.NoError(t, err)

	val, err := raftyCache.Get([]byte("1"))
	assert.NoError(t, err)
	assert.Equal(t, []byte("fadyGamil"), val)
}

func TestDel(t *testing.T) {
	raftyCache := New()
	err := raftyCache.Set([]byte("1"), []byte("fadyGamil"), 10)
	assert.NoError(t, err)

	val, err := raftyCache.Get([]byte("1"))
	assert.NoError(t, err)
	assert.Equal(t, []byte("fadyGamil"), val)

	ok, err := raftyCache.Del([]byte("1"))
	assert.NoError(t, err)
	assert.Equal(t, true, ok)
}
