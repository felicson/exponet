package mock

import (
	"fmt"

	"github.com/felicson/exponet/expo"
)

type mock int

//NewStorage return new mock
func NewStorage(dsn string) (*mock, error) {
	return new(mock), nil
}

func (m mock) Insert(exs []expo.Expo) error {

	for _, ex := range exs {
		fmt.Println(ex)
	}

	return nil
}
