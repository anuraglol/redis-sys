package core

import "time"

var store map[string]*Obj
var MAX_KEYS_LIM = 20000

func init() {
	store = make(map[string]*Obj)
}

func NewObj(value interface{}, durationMs int64, oType uint8, oEnc uint8) *Obj {
	var expiresAt int64 = -1
	if durationMs > 0 {
		expiresAt = time.Now().UnixMilli() + durationMs
	}

	return &Obj{
		Value:        value,
		ExpiresAt:    expiresAt,
		TypeEncoding: oType | oEnc,
	}
}

func Put(key string, value *Obj) {
	if len(store) >= MAX_KEYS_LIM {
		evict()
	}
	store[key] = value
}

func Get(key string) *Obj {
	v := store[key]
	if v != nil {
		if v.ExpiresAt != -1 && v.ExpiresAt <= time.Now().UnixMilli() {
			delete(store, key)
			return nil
		}
	}
	return v
}

func Del(key string) bool {
	if _, ok := store[key]; ok {
		delete(store, key)
		return true
	}
	return false
}
