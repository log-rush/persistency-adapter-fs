package storageAdapterFs

type handlesManager struct {
	files map[string]struct {
		handle fileHandle
	}
}

func newHandlesManager() handlesManager {
	return handlesManager{}
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
	return newFileHandle(stream)
}
