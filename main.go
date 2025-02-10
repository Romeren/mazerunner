package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"mazerunner/maze"
)

func main() {
	level := `==================================================.
S     |           |       |           |   |     | |
.===. | .== .==== '== .=. | ====. ====' | | | | | |
|   | | |   |         | | |     |       |   | | | |
|== | |=' .===========' | |==== |===. .=====' | | |
|   | |   |       | |     |     |   | |       | | |
| ==| | .=' .==== | | ==========' | '=' | .===| | |
|   | | |   |     |               |     | |   |   |
|== | | | .=| ==. |===================. |=' | |=. |
|   | | | | |   | |           |       | |   | | | |
| ==| | | | '=. '=' ========= | ====. '=' .=' | | |
|   | |       |         |     |     |     |   | | |
|== | '=================| .======== '=====| ==| | |
|   |       |           | |               |   |   |
| ========. | .=======. | | .=============| | |===|
|         |   |       | |   |             | | |   |
| .=. .== |===' .===. | |===| .=========. | | | | |
| | | |   |     |   | | |   | |         | | |   | |
| | | '===' | ==' | '=' | | | | | ======' '=====' |
|   |       |     |       |   | |                 E
'==================================================`

	view := tview.NewTextView()
	game := maze.NewGame(level, view)
	app := tview.NewApplication()
	app.SetRoot(view, true).EnableMouse(false)
	app.SetInputCapture(
		func(event *tcell.EventKey) *tcell.EventKey {
			return maze.InputCapture(game, event)
		},
	)
	game.Draw()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
