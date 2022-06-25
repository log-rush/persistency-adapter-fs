package main

import (
	"fmt"
	"time"

	storageAdapterFs "github.com/log-rush/persistency-adapter-fs"
	logRush "github.com/log-rush/server-devkit"
)

func main() {
	a, _ := storageAdapterFs.NewFSStorageAdapter(storageAdapterFs.Config{BasePath: "./abc/efg", OpenHandleTimeout: time.Second * 5})
	fmt.Println("writing")
	a.HandleLog(logRush.Log{Message: "test", Stream: "a"})
	time.Sleep(time.Second * 10)
	fmt.Println("writing")
	a.HandleLog(logRush.Log{Message: "test", Stream: "a"})
	time.Sleep(time.Second * 2)
	fmt.Println("writing")
	a.HandleLog(logRush.Log{Message: "test", Stream: "a"})
	time.Sleep(time.Second * 3)
	fmt.Println("should not be closed")
	fmt.Println("shutdown")
	a.Shutdown()
	time.Sleep(time.Second * 2)
	fmt.Println("done")
}
