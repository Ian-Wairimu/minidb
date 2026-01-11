package storage

import (
	"os"
	"sync"
)

type WAL struct {
	mu   sync.Mutex
	file *os.File
}

func Open(path string) (*WAL, error) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	return &WAL{file: f}, err
}

func (w *WAL) Append(line string) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	_, err := w.file.WriteString(line + "\n")
	return err
}
