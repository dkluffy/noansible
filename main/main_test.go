package main

import (
	"fmt"
	"noansible/core"
	"testing"
)

func Test_play(t *testing.T) {
	t.Run("tt.name", func(t *testing.T) {
		var pbyml core.PlaybookYML
		var tklogs core.TaskLogs

		play(&pbyml, "../files/main.yml", "../files/inventory.yml", &tklogs)
		fmt.Println("----tklogs:", tklogs)
		t.Errorf("aaa")

	})
}
