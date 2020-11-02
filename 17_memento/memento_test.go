package memento

import (
	"bufio"
	"fmt"
	"os"
)

func ExampleGame() {
	game := &Game{
		hp: 10,
		mp: 10,
	}

	game.Status()
	progress := game.Save()

	game.Play(-2, -3)
	game.Status()

	game.Load(progress)
	game.Status()

	// Output:
	// Current HP:10, MP:10
	// Current HP:7, MP:8
	// Current HP:10, MP:10
}

func ExampleInputText() {
	inputText := NewInputText()
	snapshotHolder := NewSnapshotHolder()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == ":list" {
			fmt.Println(inputText.GetText())
		} else if input == ":undo" {
			snapshot := snapshotHolder.PopSnapshot()
			if snapshot != nil {
				inputText.RestoreSnapshot(snapshot)
			}
		} else {
			snapshotHolder.PushSnapshot(inputText.CreateSnapshot())
			inputText.Append(input)
		}
	}
}
