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
		DateFormat:              "02_01_06",
		GroupStreamsIntoFolders: true,
	})
	if err != nil {
		fmt.Println(err)
	}

	a.AppendLogs("test", "ABC\n")
	a.AppendLogs("test", "EFG\n")
	a.AppendLogs("test", "HIJ\n")

	fmt.Println(a.GetLogs("test", "test_21_08_22.log"))
}
