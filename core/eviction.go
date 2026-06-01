package core

import (
	"rediss/config"
	"time"
)

func evictFirst() {
	for k := range store {
		delete(store, k)
		return
	}
}

func getCurrentClock() uint32 {
	return uint32(time.Now().UnixMilli()) & 0x00FFFFFF
}

func evictALlKeysRandom() {
	evictCount := int64(config.EvictionRatio * float64(config.KeysLimit))

	for key := range store {
		evictCount--
		delete(store, key)
		if evictCount <= 0 {
			break
		}
	}
}

func getIdleTime(lastAccessedAt uint32) uint32 {
	c := getCurrentClock()
	if c >= lastAccessedAt {
		return c - lastAccessedAt // circle back
	}
	return ((0x00FFFFFF) - lastAccessedAt) + c
}

func populateEvictionPool() {
	sampleSize := 5
	for k := range store {
		ePool.Push(k, store[k].LastAccessedAt)
		sampleSize--
		if sampleSize == 0 {
			break
		}
	}
}

func evictAllkeysLRU() {
	populateEvictionPool()
	evictCount := int16(config.EvictionRatio * float64(config.KeysLimit))
	for i := 0; i < int(evictCount) && len(ePool.pool) > 0; i++ {
		item := ePool.Pop()
		if item == nil {
			return
		}
		Del(item.key)
	}
}

func evict() {
	// evictFirst()
	// evictALlKeysRandom()
	evictAllkeysLRU()
}
