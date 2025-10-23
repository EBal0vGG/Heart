package game

import (
	"heart/internal/models"
	"heart/internal/physics"
	"heart/internal/renderer"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game implements the ebiten.Game interface
type Game struct {
	state    *models.GameState
	config   models.Config
	renderer *renderer.Renderer
}

// NewGame creates a new game instance
func NewGame(config models.Config) *Game {
	return &Game{
		state: &models.GameState{
			Points: physics.GenerateHeartPoints(config.HeartPointCount),
		},
		config:   config,
		renderer: renderer.NewRenderer(config.ScreenWidth, config.ScreenHeight),
	}
}

// Update handles game logic updates
func (g *Game) Update() error {
	// Handle mouse input
	g.handleMouseInput()

	// Update time
	dt := 1.0 / 60.0
	g.state.Time += dt

	// Handle auto-spawning balls
	g.handleAutoSpawn(dt)

	// Update balls
	g.updateBalls(dt)

	// Update heart animation
	physics.UpdateHeartAnimation(g.state.Points, g.state.Time, g.state.Balls)

	return nil
}

// Draw renders the game
func (g *Game) Draw(screen *ebiten.Image) {
	g.renderer.Render(screen, g.state)
}

// Layout returns the game layout dimensions
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.config.ScreenWidth, g.config.ScreenHeight
}

// handleMouseInput processes mouse clicks
func (g *Game) handleMouseInput() {
	mx, my := ebiten.CursorPosition()
	mousePressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	if mousePressed && !g.state.PrevMousePressed {
		for i := 0; i < g.config.BallSpawnCount; i++ {
			ball := physics.SpawnBall(float64(mx), float64(my), g.config)
			g.state.Balls = append(g.state.Balls, ball)
		}
	}
	g.state.PrevMousePressed = mousePressed
}

// handleAutoSpawn manages automatic ball spawning
func (g *Game) handleAutoSpawn(dt float64) {
	g.state.AutoSpawnTimer += dt
	if g.state.AutoSpawnTimer > g.config.AutoSpawnDelay {
		g.state.AutoSpawnTimer = 0
		ball := physics.SpawnFlyingBall(g.config)
		g.state.Balls = append(g.state.Balls, ball)
	}
}

// updateBalls updates all balls and removes expired ones
func (g *Game) updateBalls(dt float64) {
	activeBalls := g.state.Balls[:0]
	for _, b := range g.state.Balls {
		ball := b // Create a copy to modify
		if physics.UpdateBall(&ball, g.config, dt) {
			activeBalls = append(activeBalls, ball)
		}
	}
	g.state.Balls = activeBalls
}
