package storageAdapterFs

import (
	"path/filepath"
)

type handlesManager struct {
	config Config
	files  map[string]struct {
		handle fileHandle
	}
}

func newHandlesManager(config Config) handlesManager {
	return handlesManager{
		config: config,
	}
}

func (m handlesManager) Write(stream string, content string) {
	handle, ok := m.files[stream]
	if ok {
		handle.handle.Append([]byte(content))
	} else {
		handle, err := m.CreateOutput(stream)
		if err != nil {
			// error
		} else {
			handle.Append([]byte(content))
		}
	}
}

func (m handlesManager) CreateOutput(stream string) (fileHandle, error) {
	path := filepath.Join(m.config.BasePath, stream+".log")
	return newFileHandle(path)
}
