package storageAdapterFs

import (
	"os"
	"time"

	logRush "github.com/log-rush/server-devkit"
)

type Adapter struct {
	fileManager handlesManager
}

type Config struct {
	BasePath          string
	OpenHandleTimeout time.Duration
}

func NewFSStorageAdapter(config Config) (*Adapter, error) {

	// ensure logs directory
	if _, err := os.Stat(config.BasePath); os.IsNotExist(err) {
		if err := os.MkdirAll(config.BasePath, 0777); err != nil {
			return &Adapter{}, err
		}
	}

	return &Adapter{
		fileManager: newHandlesManager(config),
	}, nil
}

func (a *Adapter) HandleLog(log logRush.Log) {
	a.fileManager.Write(log.Stream, log.Message)
}

func (a *Adapter) Shutdown() {
	a.fileManager.CloseAll()
}
