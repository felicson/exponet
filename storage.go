package exponet

import (
	"exponet/expo"
	"fmt"
)

//Storage need for store data somewhere
type Storage interface {
	Insert([]expo.Expo) error
}

//Store func for storing items
func Store(storage Storage, exhs []expo.Expo) error {
	if err := storage.Insert(exhs); err != nil {
		return fmt.Errorf("On store: %v", err)
	}
	return nil
}
