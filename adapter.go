package storageAdapterFs

type Adapter struct {
	fileManager handlesManager
}

func NewFSStorageAdapter() Adapter {
	return Adapter{
		fileManager: newHandlesManager(),
	}
}

func (a Adapter) HandleLog(stream string, message string) {
	a.fileManager.Write(stream, message)
}
