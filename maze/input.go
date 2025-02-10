package maze

import (
	"github.com/gdamore/tcell/v2"
)

func InputCapture(game *Game, event *tcell.EventKey) *tcell.EventKey {
	x := game.X()
	y := game.Y()

	if event.Key() == tcell.KeyRight {
		x++
	}
	if event.Key() == tcell.KeyLeft {
		x--
	}
	if event.Key() == tcell.KeyUp {
		y--
	}
	if event.Key() == tcell.KeyDown {
		y++
	}
	game.SetPosition(x, y)
	return event
}
