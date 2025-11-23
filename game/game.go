package game

import (
	"cellular_automata/game/internal"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	automata *internal.Automata
	renderer *internal.Renderer
	surf     *ebiten.Image
	tickFreq float64
}

func (g *Game) Init() {
	initCells := make([][]internal.Cell, 100)
	for i := range initCells {
		initCells[i] = make([]internal.Cell, 100)
	}
	// Find a better way to do this
	initCells[0][0] = internal.NewCell(&internal.ImmortalAction{}, internal.CellState{Alive: true}, 0, 0)
	initCells[0][1] = internal.NewCell(&internal.ImmortalAction{}, internal.CellState{Alive: true}, 0, 1)
	initCells[0][2] = internal.NewCell(&internal.ImmortalAction{}, internal.CellState{Alive: true}, 0, 2)

	g.automata.Seed(initCells)
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.automata.Paused = !g.automata.Paused
		if !g.automata.Paused {
			g.automata.Play()
		}
	}

	if g.automata.Poll() {
		g.automata.Tick()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.surf.Fill(color.RGBA{0, 0, 0, 255})

	grid := g.automata.GetGridState()
	g.renderer.DrawGrid(g.surf, grid)

	screen.DrawImage(g.surf, &ebiten.DrawImageOptions{})

	if g.automata.Paused {
		ebitenutil.DebugPrint(screen, "Paused")
	} else {
		ebitenutil.DebugPrint(screen, "Playing")
	}
}

func (g *Game) Layout(outsideWith, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}

func NewGame() *Game {
	return &Game{
		automata: internal.NewAutomata(1),
		renderer: internal.NewRenderer(64, 64),
		surf:     ebiten.NewImage(1920, 1080),
	}
}
