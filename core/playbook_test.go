package core

import (
	"fmt"
	//"reflect"
	"testing"
)


func TestLoadPlaybook_a(t *testing.T) {
	t.Run("test2-main.yml", func(t *testing.T) {
		var pl player

		var pby playbookYML
		pl = &pby

		err := pl.Loader("..\\files\\main.yml")
		fmt.Println(pby.tasks, len(pby.tasks))
		if err != nil {
			t.Errorf("LoadPlaybook() = %v", err)
		}
		t.Errorf("LoadPlaybook() = %v", err)

	})
}

func Test_loadrawbook(t *testing.T) {
	t.Run("test3-rawbook.yml", func(t *testing.T) {
		book,err := loadrawbook("..\\files\\main.yml")
		fmt.Println(book["tasks"], len(book))
		if err != nil {
			t.Errorf("LoadPlaybook() = %v", err)
		}
		t.Errorf("LoadPlaybook() = %v", err)

	})
}
