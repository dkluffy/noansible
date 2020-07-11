package core

import (
	"fmt"
	//"reflect"
	"testing"
)

func TestLoadPlaybook_a(t *testing.T) {
	t.Run("test2-main.yml", func(t *testing.T) {

		var pby PlaybookYML
		var pl Playbook
		pl = &pby

		pl.Loader("..\\files\\main.yml", "../files/inventory.yml")
		var tklogs TaskLogs
		pl.Player(&tklogs)
		fmt.Println("----tklogs:", tklogs)

		// tasks := pby.tasklist
		// fmt.Println(pby, len(tasks), tklogs)
		// fmt.Println("tasks---:", tasks, len(tasks))

		t.Errorf("No err -- LoadPlaybook() = ")

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
