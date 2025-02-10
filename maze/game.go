package maze

import (
	"github.com/rivo/tview"
	"strings"
)

type Game struct {
	content *tview.TextView
	world   [][]rune
	posX    int
	posY    int
}

func NewGame(content string, view *tview.TextView) *Game {
	trows := strings.Split(content, "\n")
	rows := make([][]rune, len(trows))
	for r := 0; r < len(rows); r++ {
		rows[r] = []rune(strings.Trim(trows[r], " "))
	}

	return &Game{
		content: view,
		world:   rows,
		posX:    0,
		posY:    1,
	}
}

func (g *Game) X() int {
	return g.posX
}

func (g *Game) Y() int {
	return g.posY
}

func (g *Game) Draw() {
	strs := make([]string, len(g.world))
	for i, rs := range g.world {
		strs[i] = string(rs)
	}
	maze := strings.Join(strs, "\n")
	g.content.SetText(maze)
}

func (g *Game) SetPosition(x, y int) {
	g.posX = x
	g.posY = y

	g.world[g.Y()][g.X()] = 'X'
	g.Draw()
}
