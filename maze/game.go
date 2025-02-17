package maze

import (
	"fmt"
	"github.com/rivo/tview"
	"strings"
)

type Game struct {
	content *tview.TextView
	world   [][]string
	posX    int
	posY    int
}

func NewGame(content string, view *tview.TextView) *Game {
	trows := strings.Split(content, "\n")
	rows := make([][]string, len(trows))
	for r := 0; r < len(rows); r++ {
		line := make([]string, len(trows[r]))
		trimmedLine := strings.Trim(trows[r], " ")
		for pos, char := range trimmedLine {
			line[pos] = fmt.Sprintf("%c", char)
		}
		rows[r] = line
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

func (g *Game) Ymax() int {
	return len(g.world) - 1
}

func (g *Game) Xmax() int {
	return len(g.world[0]) - 1
}

func (g *Game) Draw() {
	strs := make([]string, len(g.world))
	for i, rs := range g.world {
		strs[i] = strings.Join(rs, "")
	}
	maze := strings.Join(strs, "\n")
	g.content.SetText(maze)
}

func (g *Game) SetPosition(x, y int) {
	g.posX = x
	g.posY = y

	g.world[g.Y()][g.X()] = "X"
	g.Draw()
}

func (g *Game) GetPosition(x, y int) string {
	return g.world[y][x]
}
