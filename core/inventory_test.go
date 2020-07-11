package core

import (
	"testing"
)

func TestReadInventory(t *testing.T) {

	t.Run("tt.name", func(t *testing.T) {
		got, err := ReadInventory("TT", "../files/inventory.yml")

		t.Errorf("Inventory = %v,%v", got, err)

	})

}
