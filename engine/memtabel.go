package engine

import (
	"sync"

	apperr "github.com/bimal009/db/pkg/err"
)

type MemTabel struct {
	mu        sync.RWMutex
	entries   map[string]entry
	sizeBytes int // total bytes of all keys and values stored
}

type entry struct {
	value   []byte
	deleted bool
}

func NewMemtable() *MemTabel {
	return &MemTabel{entries: make(map[string]entry)}
}

// Put stores value directly without copying.
// The caller must not modify value after this call.
func (m *MemTabel) Put(key string, value []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if old, ok := m.entries[key]; ok {
		m.sizeBytes -= len(key) + len(old.value)
	}

	m.entries[key] = entry{value: value, deleted: false}
	m.sizeBytes += len(key) + len(value)
	return nil
}

// Get returns the stored slice directly. Callers must not modify it.
func (m *MemTabel) Get(key string) ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	e, ok := m.entries[key]
	if !ok {
		return nil, apperr.ErrKeyNotFound
	}
	if e.deleted {
		return nil, apperr.ErrKeyDeleted
	}
	return e.value, nil
}

func (m *MemTabel) Delete(key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if old, ok := m.entries[key]; ok {
		m.sizeBytes -= len(key) + len(old.value)
	}

	m.entries[key] = entry{value: nil, deleted: true}
	m.sizeBytes += len(key)
	return nil
}

func (m *MemTabel) Size() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.sizeBytes
}
