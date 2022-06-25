package storageAdapterFs

import (
	"log"
	"path/filepath"
	"sync"
	"time"
)

type openFileHandle struct {
	handle  *fileHandle
	timeout *time.Timer
}

type handlesManager struct {
	config Config
	files  map[string]*openFileHandle
	mutex  *sync.RWMutex
}

func newHandlesManager(config Config) handlesManager {
	return handlesManager{
		config: config,
		files:  map[string]*openFileHandle{},
		mutex:  &sync.RWMutex{},
	}
}

func (m *handlesManager) Write(stream string, content string) {
	m.mutex.RLock()
	handle, ok := m.files[stream]
	m.mutex.RUnlock()
	if ok {
		handle.handle.Append([]byte(content))
		handle.timeout.Reset(m.config.OpenHandleTimeout)
	} else {
		handle, err := m.CreateOutput(stream)
		if err != nil {
			log.Printf("cant create log file handle: %s", err)
		} else {
			handle.handle.Append([]byte(content))
		}
	}
}

func (m *handlesManager) CreateOutput(stream string) (openFileHandle, error) {
	path := filepath.Join(m.config.BasePath, stream+".log")
	handle, err := newFileHandle(path)
	if err != nil {
		return openFileHandle{}, err
	}

	timer := time.NewTimer(m.config.OpenHandleTimeout)
	fileHandle := openFileHandle{
		handle:  handle,
		timeout: timer,
	}

	m.mutex.Lock()
	m.files[stream] = &fileHandle
	m.mutex.Unlock()

	go func(stream string) {
		<-timer.C
		err := handle.Close()
		if err != nil {
			log.Printf("error while closing log file: %s", err)
		}

		m.mutex.Lock()
		delete(m.files, stream)
		m.mutex.Unlock()
	}(stream)

	return fileHandle, err
}

func (m *handlesManager) CloseAll() {
	m.mutex.Lock()
	for _, handle := range m.files {
		handle.timeout.Stop()
		err := handle.handle.Close()
		if err != nil {
			log.Printf("error while closing log file: %s", err)
		}
	}
}
