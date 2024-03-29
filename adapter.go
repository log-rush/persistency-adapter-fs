package storageAdapterFs

import (
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

type Adapter struct {
	fileManager handlesManager
	cronHandler *cron.Cron
}

type Config struct {
	BasePath                string
	OpenHandleTimeout       time.Duration
	ForceUpdateOnMidnight   bool
	DateFormat              string
	Timezone                string
	GroupStreamsIntoFolders bool
	FilePermissions         os.FileMode
}

func NewFSStorageAdapter(config Config) (*Adapter, error) {

	if config.FilePermissions == 0 {
		config.FilePermissions = 0744
	}

	// ensure logs directory
	if _, err := os.Stat(config.BasePath); os.IsNotExist(err) {
		if err := os.MkdirAll(config.BasePath, config.FilePermissions); err != nil {
			return &Adapter{}, err
		}
	}

	adapter := Adapter{
		fileManager: newHandlesManager(config),
	}

	if config.ForceUpdateOnMidnight {
		customLocation, err := time.LoadLocation(config.Timezone)
		if err != nil {
			return &Adapter{}, err
		}

		cronHandler := cron.New(cron.WithLocation(customLocation))
		cronHandler.AddFunc("0 0 * * * *", func() {
			adapter.fileManager.CloseAll()
		})
		cronHandler.Start()
		adapter.cronHandler = cronHandler
	}

	return &adapter, nil
}

func (a *Adapter) AppendLogs(key string, logs string) {
	a.fileManager.Write(key, logs)
}

func (a *Adapter) ListLogFiles(key string) []string {
	return a.fileManager.ListLogFiles(key)
}

func (a *Adapter) GetLogs(key string, logFile string) ([]byte, error) {
	return a.fileManager.GetLogs(key, logFile)
}

func (a *Adapter) Shutdown() {
	a.fileManager.CloseAll()
	a.cronHandler.Stop()
}
