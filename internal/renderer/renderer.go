package renderer

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"heart/internal/models"
)

// Renderer handles all rendering operations
type Renderer struct {
	trail *ebiten.Image
	fade  *ebiten.Image
}

// NewRenderer creates a new renderer instance
func NewRenderer(width, height int) *Renderer {
	return &Renderer{
		trail: ebiten.NewImage(width, height),
		fade:  ebiten.NewImage(width, height),
	}
}

// Render draws the entire game scene
func (r *Renderer) Render(screen *ebiten.Image, state *models.GameState) {
	// Initialize images if needed
	if r.trail == nil {
		r.trail = ebiten.NewImage(800, 600)
		r.fade = ebiten.NewImage(800, 600)
	}

	// Update trail fade effect
	r.updateTrailFade()

	// Draw balls onto trail
	r.drawBallsToTrail(state.Balls)

	// Draw background
	screen.Fill(color.RGBA{12, 8, 18, 255})

	// Draw trail over background
	r.drawTrail(screen)

	// Draw heart glow
	r.drawHeartGlow(screen, state.Points)

	// Draw main heart points
	r.drawHeartPoints(screen, state.Points)

	// Draw balls on top
	r.drawBalls(screen, state.Balls)
}

// updateTrailFade handles the trail fade effect
func (r *Renderer) updateTrailFade() {
	alpha := uint8(6)
	if rand.Float64() < 0.03 {
		alpha = 40
	}
	r.fade.Fill(color.RGBA{0, 0, 0, alpha})
	r.trail.DrawImage(r.fade, nil)
}

// drawBallsToTrail draws balls onto the trail layer
func (r *Renderer) drawBallsToTrail(balls []models.Ball) {
	for _, b := range balls {
		trailAlpha := uint8(math.Min(200, b.TTL/5.0*200))
		cc := color.RGBA{0, 0, 0, trailAlpha}
		vector.FillCircle(r.trail, float32(b.X), float32(b.Y), float32(b.R*0.6), cc, false)
	}
}

// drawTrail draws the trail layer onto the screen
func (r *Renderer) drawTrail(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(0.85)
	screen.DrawImage(r.trail, op)
}

// drawHeartGlow draws the heart glow effect
func (r *Renderer) drawHeartGlow(screen *ebiten.Image, points []models.Point) {
	glowColor := color.RGBA{255, 110, 160, 40}
	for _, p := range points {
		vector.FillCircle(screen, float32(p.X), float32(p.Y), 3.0, glowColor, false)
	}
}

// drawHeartPoints draws the main heart points
func (r *Renderer) drawHeartPoints(screen *ebiten.Image, points []models.Point) {
	mainColor := color.RGBA{255, 80, 120, 255}
	for _, p := range points {
		vector.FillRect(screen, float32(p.X), float32(p.Y), 2, 2, mainColor, false)
	}
}

// drawBalls draws the balls on top of everything
func (r *Renderer) drawBalls(screen *ebiten.Image, balls []models.Ball) {
	for _, b := range balls {
		alpha := uint8(math.Min(255, b.TTL/5.0*255))
		cc := color.RGBA{0, 0, 0, alpha}
		vector.FillCircle(screen, float32(b.X), float32(b.Y), float32(b.R), cc, false)
	}
}
