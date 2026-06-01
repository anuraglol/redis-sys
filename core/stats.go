package core

import (
	"sync/atomic"
)

type Stats struct {
	TotalCommands int64
	SetCmds       int64
	GetCmds       int64
	PingCmds      int64
	ExpireCmds    int64
}

var stats = &Stats{}

func IncrTotalCommands() {
	atomic.AddInt64(&stats.TotalCommands, 1)
}

func IncrSetCmds() {
	atomic.AddInt64(&stats.SetCmds, 1)
}

func IncrGetCmds() {
	atomic.AddInt64(&stats.GetCmds, 1)
}

func IncrPingCmds() {
	atomic.AddInt64(&stats.PingCmds, 1)
}

func IncrExpireCmds() {
	atomic.AddInt64(&stats.ExpireCmds, 1)
}

func GetStats() *Stats {
	return stats
}
