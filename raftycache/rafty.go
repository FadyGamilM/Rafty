package raftycache

import (
	"fmt"
	"sync"
	"time"
)

type MsgKey string
type MsgVal []byte

type RaftyCache struct {
	// the in-mem data storage
	Data map[MsgKey]MsgVal
	// the lock around Data to ensure isolation
	lock *sync.RWMutex
}

func New() *RaftyCache {
	return &RaftyCache{
		Data: make(map[MsgKey]MsgVal, 0),
		lock: &sync.RWMutex{},
	}
}

// implement the interface
func (rc *RaftyCache) Set(k, v []byte, ttl time.Duration) error {
	// lcok
	rc.lock.Lock()
	// unlock
	defer rc.lock.Unlock()

	// emsure this key never exists before
	key := MsgKey(string(k))
	_, ok := rc.Data[key]
	if ok {
		return fmt.Errorf("➜ key already exists")
	}
	rc.Data[key] = v
	return nil
}
func (rc *RaftyCache) Get(k []byte) ([]byte, error) {
	// lcok
	rc.lock.RLock()
	// unlock
	defer rc.lock.RUnlock()

	key := MsgKey(string(k))
	val, ok := rc.Data[key]
	if !ok {
		return nil, fmt.Errorf("➜ key does not exists")
	}
	return val, nil
}
func (rc *RaftyCache) Del(k []byte) (bool, error) {
	// lcok
	rc.lock.Lock()
	// unlock
	defer rc.lock.Unlock()

	key := MsgKey(string(k))
	_, ok := rc.Data[key]
	if !ok {
		return false, fmt.Errorf("➜ key does not exists")
	}
	delete(rc.Data, key)
	return true, nil
}
func (rc *RaftyCache) Has(k []byte) (bool, error) {
	// lcok
	rc.lock.RLock()
	// unlock
	defer rc.lock.RUnlock()

	key := MsgKey(string(k))
	_, ok := rc.Data[key]
	if ok {
		return false, fmt.Errorf("➜ key already exists")
	}

	return true, nil
}
