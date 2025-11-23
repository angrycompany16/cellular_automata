package internal

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Renderer struct {
	gridWidth, gridHeight int
}

func (r *Renderer) DrawGrid(dst *ebiten.Image, grid [][]Cell) {
	for y := range grid {
		for x := range grid[0] {
			if grid[y][x].state.Alive {
				vector.FillRect(dst, float32(x*r.gridWidth), float32(y*r.gridHeight), float32(r.gridWidth), float32(r.gridHeight), color.White, false)
			}
		}
	}
}

func NewRenderer(w, h int) *Renderer {
	return &Renderer{
		gridWidth:  w,
		gridHeight: h,
	}
}
