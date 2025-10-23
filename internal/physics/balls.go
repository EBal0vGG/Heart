package physics

import (
	"math/rand"

	"heart/internal/models"
)

// SpawnBall creates a new ball at the specified position
func SpawnBall(x, y float64, config models.Config) models.Ball {
	vx := (rand.Float64() - 0.5) * 2
	vy := (rand.Float64() - 0.5) * 2
	return models.Ball{
		X: x, Y: y, Vx: vx, Vy: vy,
		R: config.BallRadius, TTL: config.BallTTL,
	}
}

// SpawnFlyingBall creates a ball that flies in from the screen edges
func SpawnFlyingBall(config models.Config) models.Ball {
	side := rand.Intn(4)
	var x, y, vx, vy float64
	speed := 4.0

	switch side {
	case 0: // Left
		x = -10
		y = rand.Float64() * float64(config.ScreenHeight)
		vx = speed
		vy = (rand.Float64() - 0.5) * 1.0
	case 1: // Right
		x = float64(config.ScreenWidth) + 10
		y = rand.Float64() * float64(config.ScreenHeight)
		vx = -speed
		vy = (rand.Float64() - 0.5) * 1.0
	case 2: // Top
		x = rand.Float64() * float64(config.ScreenWidth)
		y = -10
		vx = (rand.Float64() - 0.5) * 1.0
		vy = speed
	case 3: // Bottom
		x = rand.Float64() * float64(config.ScreenWidth)
		y = float64(config.ScreenHeight) + 10
		vx = (rand.Float64() - 0.5) * 1.0
		vy = -speed
	}

	return models.Ball{
		X: x, Y: y, Vx: vx, Vy: vy,
		R: config.BallRadius, TTL: config.FlyingBallTTL,
	}
}

// UpdateBall updates ball physics and handles collisions
func UpdateBall(b *models.Ball, config models.Config, dt float64) bool {
	b.TTL -= dt
	if b.TTL <= 0 {
		return false // Remove ball
	}

	// Add some randomness to movement
	b.Vx += (rand.Float64() - 0.5) * 0.08
	b.Vy += (rand.Float64() - 0.5) * 0.08
	b.X += b.Vx
	b.Y += b.Vy

	// Handle screen boundaries
	if b.X-b.R < 0 {
		b.X = b.R
		b.Vx *= -1
	}
	if b.X+b.R > float64(config.ScreenWidth) {
		b.X = float64(config.ScreenWidth) - b.R
		b.Vx *= -1
	}
	if b.Y-b.R < 0 {
		b.Y = b.R
		b.Vy *= -1
	}
	if b.Y+b.R > float64(config.ScreenHeight) {
		b.Y = float64(config.ScreenHeight) - b.R
		b.Vy *= -1
	}

	// Apply friction
	b.Vx *= 0.995
	b.Vy *= 0.995

	return true // Keep ball
}
