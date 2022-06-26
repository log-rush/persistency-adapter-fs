package storageAdapterFs

import (
	"os"
	"path/filepath"
)

type fileHandle struct {
	path string
	f    *os.File
}

func newFileHandle(path string, permissions os.FileMode) (*fileHandle, error) {
	dir := filepath.Dir(path)
	// ensure stream directory
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, permissions); err != nil {
			return &fileHandle{}, err
		}
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, permissions)
	return &fileHandle{
		path: path,
		f:    f,
	}, err
}

func (h *fileHandle) Append(content []byte) error {
	_, err := h.f.Write(content)
	return err
}

func (h *fileHandle) Close() error {
	return h.f.Close()
}
