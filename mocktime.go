package mocktime

import (
	"sync"
	"time"
)

// Mock provides mock for the standard function time.Now.
type Mock struct {
	mutex   sync.Mutex
	curtime time.Time
}

// New creates new mock for time.Now.
func New() *Mock {
	return &Mock{
		mutex:   sync.Mutex{},
		curtime: time.Now(),
	}
}

// Now returns the current mock time.
func (m *Mock) Now() time.Time {
	return m.curtime
}

// Set sets the current time.
func (m *Mock) Set(v time.Time) {
	m.mutex.Lock()
	m.curtime = v
	m.mutex.Unlock()
}
