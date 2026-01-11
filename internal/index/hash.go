package index

import (
	"errors"
	"sync"
)

type HashIndex struct {
	mu     sync.RWMutex
	unique bool
	data   map[any][]int
}

func NewHashIndex(unique bool) *HashIndex {
	return &HashIndex{
		unique: unique,
		data:   make(map[any][]int),
	}
}

func (h *HashIndex) Insert(v any, id int) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.unique && len(h.data[v]) > 0 {
		return ErrUniqueViolation
	}
	h.data[v] = append(h.data[v], id)
	return nil
}

func (h *HashIndex) Search(v any) []int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.data[v]
}

var ErrUniqueViolation = errors.New("unique constraint violation")
