package core

import (
	"fmt"
	//"reflect"
	"testing"
)

func TestLoadPlaybook_a(t *testing.T) {
	t.Run("test2-main.yml", func(t *testing.T) {
		var pl Playebook

		var pby PlaybookYML
		pl = &pby

		err := pl.Loader("..\\files\\main.yml")
		tasks := pby.tasklist
		fmt.Println(pby, len(tasks))
		fmt.Println("tasks---:", tasks, len(tasks))
		if err != nil {
			t.Errorf("ERR: LoadPlaybook() = %v", err)
		}
		t.Errorf("No err -- LoadPlaybook() = %v", err)

	})
}

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
