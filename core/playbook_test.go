package core

import (
	"fmt"
	//"reflect"
	"testing"
)

func Test_loadrawbook(t *testing.T) {
	t.Run("test3-rawbook.yml", func(t *testing.T) {
		book, err := loadrawbook("..\\files\\main.yml")
		fmt.Println(book, book["tasks"], len(book))
		if err != nil {
			t.Errorf("LoadPlaybook() = %v", err)
		}
		t.Errorf("LoadPlaybook() = %v", err)

	})
}

func Test_loader(t *testing.T) {
	t.Run("test3-rawbook.yml", func(t *testing.T) {
		var pbyml PlaybookYML
		pbyml.Loader("..\\files\\main.yml", "..\\files\\inventory.yml")
		fmt.Println(pbyml.tasklist)

		t.Errorf("-----LoadPlaybook() = %v", pbyml)

	})
}
