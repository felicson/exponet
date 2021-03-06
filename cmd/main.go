package main

import (
	"github.com/felicson/exponet/expo"

	"github.com/felicson/exponet"

	"log"

	storage "github.com/felicson/exponet/storage/mysql"
)

var (
	//define it via -X ldflag
	dsn      = ""
	indexURL = "https://www.exponet.ru/exhibitions/countries/rus/topics/promexpo/dates/future/p1l10000.ru.html"
)

func main() {
	var (
		stor exponet.Storage
		err  error
		exhs []expo.Expo
	)

	if stor, err = storage.NewStorage(dsn); err != nil {
		log.Fatal(err)
	}
	if exhs, err = exponet.GetExhibitions(indexURL); err != nil {
		log.Fatal(err)
	}
	if err = exponet.Store(stor, exhs); err != nil {
		log.Fatal(err)
	}
}
