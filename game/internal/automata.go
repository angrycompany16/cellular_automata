package internal

import (
	"cellular_automata/game/internal/utils"
	"time"
)

type Automata struct {
	Paused     bool
	Ticker     *time.Ticker
	gridHeight int
	gridWidth  int
	grid       [][]Cell
	tickRate   float64
}

func (a *Automata) Seed(cells [][]Cell) {
	a.grid = cells
	// a.grid = make([][]Cell, len(cells))
	// for i := range a.grid {
	// 	a.grid[i] = make([]Cell, len(cells[0]))
	// }
	// for y := range a.grid {
	// 	for x := range a.grid[y] {
	// 		a.grid[y][x] = cells[y][x]
	// 	}
	// }
}

func (a *Automata) Tick() {
	if a.Paused {
		return
	}

	for _, row := range a.grid {
		for _, cell := range row {
			cell.Update(a.grid)
		}
	}

	for _, row := range a.grid {
		for _, cell := range row {
			cell.Propagate()
		}
	}
}

func (a *Automata) GetGridState() [][]Cell {
	return a.grid
}

func (a *Automata) Play() {
	a.Ticker.Reset(time.Duration((1 / a.tickRate) * (float64(time.Second))))
}

func (a *Automata) Poll() bool {
	if a.Paused {
		return false
	}
	_, tick := utils.PollThread(a.Ticker.C)
	return tick
}

func NewAutomata(tickRate float64) *Automata {
	return &Automata{
		Paused:   true,
		tickRate: tickRate,
		Ticker:   time.NewTicker(time.Duration((1 / tickRate) * (float64(time.Second)))),
	}
}
