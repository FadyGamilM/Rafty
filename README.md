# Rafty
Rafty is a Distributed Cache built in Go, utilizing the Raft consensus algorithm

# How Cache Eviction is handled (via TTL strategy) ? 
```go
func (rc *RaftyCache) Set(k, v []byte, ttl time.Duration) error {
	// lcok
	rc.lock.Lock()
	// unlock
	defer rc.lock.Unlock()

	// emsure this key never exists before
	key := string(k)
	_, ok := rc.Data[key]
	if ok {
		return fmt.Errorf("➜ key already exists")
	}
	rc.Data[key] = v

	// cache eviction
	go rc.EvictMsg(k, ttl)
	return nil
}

// handle cache eviction
func (rc *RaftyCache) EvictMsg(key []byte, ttl time.Duration) {
	ticker := time.NewTicker(ttl)
	<-ticker.C
	delete(rc.Data, string(key))
}
```
- after setting the message in the in-mem cache.
- we spawn a go routine that takes control to delete the key from the cache after the specified time duration by opening a ticker and block untill the ticker is triggered.