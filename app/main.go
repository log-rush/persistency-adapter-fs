package main

import storageAdapterFs "github.com/log-rush/persistency-adapter-fs"

func main() {
	a := storageAdapterFs.NewFSStorageAdapter(storageAdapterFs.Config{"./abc/efg"})
	a.HandleLog("a", "test")
}
