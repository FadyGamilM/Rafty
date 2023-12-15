package cache

import "time"

type Cacher interface {
	Set(k, v []byte, ttl time.Duration) error
	Get(k []byte) ([]byte, error)
	Del(k []byte) (bool, error)
	Has(k []byte) (bool, error)
}
