package storageAdapterFs

import "os"

type fileHandle struct {
	path string
	f    *os.File
}

func newFileHandle(path string) (fileHandle, error) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	return fileHandle{
		path: path,
		f:    f,
	}, err
}

func (h fileHandle) Append(content []byte) error {
	_, err := h.f.Write(content)
	return err
}

func (h fileHandle) Close() error {
	return h.f.Close()
}
