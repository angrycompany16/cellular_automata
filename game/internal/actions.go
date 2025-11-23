package internal

import "cellular_automata/game/internal/utils"

type CellAction interface {
	Execute(cells [][]Cell, x, y int) CellState
}

// Simple cell that follows the game of life rules
type LifeAction struct{}

func (l *LifeAction) Execute(cells [][]Cell, x, y int) CellState {
	state := cells[x][y].state
	liveNeighbours := 0
	w, h := len(cells), len(cells[0])
	for i := utils.ClampInt(x-1, 0, w); i <= utils.ClampInt(x+1, 0, w); i++ {
		for j := utils.ClampInt(y-1, 0, h); j <= utils.ClampInt(y+1, 0, h); j++ {
			if cells[i][j].state.Alive {
				liveNeighbours += 1
			}
		}
	}
	if liveNeighbours == 3 && !state.Alive {
		state.Alive = true
	}
	if liveNeighbours < 2 && state.Alive {
		state.Alive = false
	}
	if liveNeighbours > 3 && state.Alive {
		state.Alive = false
	}
	return state
}

// Cell that is always alive
type ImmortalAction struct{}

func (i *ImmortalAction) Execute(cells [][]Cell, x, y int) CellState {
	state := CellState{Alive: true}
	return state
}
