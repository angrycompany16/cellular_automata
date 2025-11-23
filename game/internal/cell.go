package internal

type CellState struct {
	Alive bool
}

type Cell struct {
	action    CellAction
	state     CellState
	nextState CellState
	x, y      int
}

func (c *Cell) Update(grid [][]Cell) {
	if c.action == nil {
		return
	}
	c.nextState = c.action.Execute(grid, c.x, c.y)
}

func (c *Cell) Propagate() {
	c.state = c.nextState
}

func NewCell(action CellAction, initState CellState, y, x int) Cell {
	return Cell{
		action:    action,
		state:     initState,
		nextState: initState,
		x:         x,
		y:         y,
	}
}
