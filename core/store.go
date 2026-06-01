package core

import (
	"rediss/config"
	"time"
)

var store map[string]*Obj
var expires map[*Obj]uint64

func init() {
	store = make(map[string]*Obj)
	expires = make(map[*Obj]uint64)
}

func setExpiry(obj *Obj, exDurationMs int64) {
	expires[obj] = uint64(time.Now().UnixMilli()) + uint64(exDurationMs)
}

func NewObj(value interface{}, expDurationMs int64, oType uint8, oEnc uint8) *Obj {
	obj := &Obj{
		Value:          value,
		TypeEncoding:   oType | oEnc,
		LastAccessedAt: getCurrentClock(),
	}
	if expDurationMs > 0 {
		setExpiry(obj, expDurationMs)
	}
	return obj
}

func Put(key string, obj *Obj) {
	if len(store) >= config.KeysLimit {
		evict()
	}
	obj.LastAccessedAt = getCurrentClock()
	store[key] = obj
}

func Get(key string) *Obj {
	v := store[key]
	if v == nil {
		return nil
	}
	if v != nil {
		if hasExpired(v) {
			Del(key)
			return nil
		}
	}
	v.LastAccessedAt = getCurrentClock()
	return v
}

func Del(key string) bool {
	if obj, ok := store[key]; ok {
		delete(store, key)
		delete(expires, obj)
		return true
	}
	return false
}
