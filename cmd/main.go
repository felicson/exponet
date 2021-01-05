package main

import (
	"exponet"
	"exponet/expo"

	//storage "exponet/storage/mock"

	storage "exponet/storage/mysql"
	"log"
)

//define via -X ldflag
var dsn = ""

func main() {
	var (
		stor exponet.Storage
		err  error
		exhs []expo.Expo
	)

	if stor, err = storage.NewStorage(dsn); err != nil {
		log.Fatal(err)
	}
	if exhs, err = exponet.GetExhibitions(); err != nil {
		log.Fatal(err)
	}
	if err = exponet.Store(stor, exhs); err != nil {
		log.Fatal(err)
	}
}
