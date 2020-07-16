package main

import (
	"exponet"
	"exponet/expo"
	"exponet/storage"
	"fmt"
	"os"
)

func main() {
	var (
		stor exponet.Storage
		err  error
		exhs []expo.Expo
	)
	if stor, err = storage.NewStorage(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if exhs, err = exponet.GetExhibitions(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err = exponet.Store(stor, exhs); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
