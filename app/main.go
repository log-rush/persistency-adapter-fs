package main

import (
	"fmt"
	"time"

	storageAdapterFs "github.com/log-rush/persistency-adapter-fs"
)

func main() {
	a, err := storageAdapterFs.NewFSStorageAdapter(storageAdapterFs.Config{
		BasePath:                "./_logs",
		OpenHandleTimeout:       time.Second * 5,
		ForceUpdateOnMidnight:   true,
		DateFormat:              "_02_01_06",
		GroupStreamsIntoFolders: true,
	})
	if err != nil {
		fmt.Println(err)
	}

	a.AppendLogs("test", "")

	fmt.Println(a.ListLogFiles("test"))
}
