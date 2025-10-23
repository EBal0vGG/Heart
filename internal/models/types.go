package models

import "math"

// Point represents a particle in the heart shape
type Point struct {
	X, Y         float64
	HomeX, HomeY float64
	Vx, Vy       float64
}

// Ball represents a bouncing ball particle
type Ball struct {
	X, Y, R float64
	Vx, Vy  float64
	TTL     float64
}

// GameState holds all the game data
type GameState struct {
	Points           []Point
	Balls            []Ball
	PrevMousePressed bool
	Time             float64
	Trail            interface{} // Will be *ebiten.Image
	Fade             interface{} // Will be *ebiten.Image
	AutoSpawnTimer   float64
}

// Config holds game configuration
type Config struct {
	ScreenWidth     int
	ScreenHeight    int
	HeartPointCount int
	AutoSpawnDelay  float64
	BallSpawnCount  int
	BallRadius      float64
	BallTTL         float64
	FlyingBallTTL   float64
}

// ApplyRepulsion applies repulsion force from a ball to a point
func (p *Point) ApplyRepulsion(b Ball) {
	dx := p.X - b.X
	dy := p.Y - b.Y
	dist := math.Hypot(dx, dy)

	repulsionRadius := b.R * 12
	if dist < repulsionRadius && dist > 0 {
		force := (repulsionRadius - dist) * 0.04
		p.Vx += (dx / dist) * force
		p.Vy += (dy / dist) * force
	}
}

// ApplyReturn applies return force to bring point back to target position
func (p *Point) ApplyReturn(targetX, targetY float64) {
	dx := targetX - p.X
	dy := targetY - p.Y
	p.Vx += dx * 0.005
	p.Vy += dy * 0.005
}

// Update updates point position based on velocity
func (p *Point) Update() {
	p.Vx *= 0.88
	p.Vy *= 0.88
	p.X += p.Vx
	p.Y += p.Vy
}
