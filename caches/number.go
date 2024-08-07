package caches

import (
	"fmt"
	"sync/atomic"
)

var (
	Number atomic.Uint32
)

func InitNumber(num int64) {
	Number.Store(uint32(num))
}

func GetNumber() string {
	return fmt.Sprintf("#%d", Number.Load())
}
