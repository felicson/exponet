package exponet

import (
	"fmt"

	"github.com/felicson/exponet/expo"
)

//Storage need for store data somewhere
type Storage interface {
	Insert([]expo.Expo) error
}

//Store func for storing items
func Store(storage Storage, exhs []expo.Expo) error {
	if err := storage.Insert(exhs); err != nil {
		return fmt.Errorf("on store err: %w", err)
	}
	return nil
}
