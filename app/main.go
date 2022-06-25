package main

import (
	"fmt"
	"time"

	storageAdapterFs "github.com/log-rush/persistency-adapter-fs"
)

func main() {
	a, err := storageAdapterFs.NewFSStorageAdapter(storageAdapterFs.Config{
		BasePath:              "./abc/efg",
		OpenHandleTimeout:     time.Second * 5,
		ForceUpdateOnMidnight: true,
		DateFormat:            "05--02_01_06",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("writing")
	a.AppendLogs("test", "a")
	time.Sleep(time.Second * 10)
	fmt.Println("writing")
	a.AppendLogs("test", "a")
	time.Sleep(time.Second * 2)
	fmt.Println("writing")
	a.AppendLogs("test", "a")
	time.Sleep(time.Second * 3)
	fmt.Println("should not be closed")
	fmt.Println("shutdown")
	a.Shutdown()
	time.Sleep(time.Second * 2)
	fmt.Println("done")
}
