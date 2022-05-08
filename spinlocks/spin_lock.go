package main

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type SpinLock int32

func NewSpinLock() sync.Locker {
	var lock SpinLock
	return &lock
}

func (l *SpinLock) Lock() {
	for !atomic.CompareAndSwapInt32((*int32)(l), 0, 1) {
		runtime.Gosched()
	}
}

func (l *SpinLock) Unlock() {
	atomic.StoreInt32((*int32)(l), 0)
}
