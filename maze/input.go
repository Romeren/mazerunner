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
		if game.GetPosition (x, y-1) == ' '{
			 //(game.GetPosition(x, y-1))
			y--
		}
	}
	if event.Key() == tcell.KeyDown {
		y++
	} 

	if x <0 {
		x=0
	}
	if y <0 {
		y=0
	}
	if x>game.Xmax(){
		x=game.Xmax()
	}
	if y>game.Ymax(){
		y=game.Ymax ()
	}
	game.SetPosition(x, y)
	return event
}
