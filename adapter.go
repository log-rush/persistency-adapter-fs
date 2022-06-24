package storageAdapterFs

import (
	"log"
	"os"
)

type Adapter struct {
	fileManager handlesManager
}

type Config struct {
	BasePath string
}

func NewFSStorageAdapter(config Config) Adapter {

	// ensure logs directory
	if _, err := os.Stat(config.BasePath); os.IsNotExist(err) {
		if err := os.MkdirAll(config.BasePath, 0777); err != nil {
			log.Fatalf("cant create logs directory: %s\n", err)
			return Adapter{}
		}
	}

	return Adapter{
		fileManager: newHandlesManager(config),
	}
}

func (a Adapter) HandleLog(stream string, message string) {
	a.fileManager.Write(stream, message)
}
