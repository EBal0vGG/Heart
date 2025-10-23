package physics

import (
	"math"

	"heart/internal/models"
)

// UpdateHeartAnimation updates the heart shape animation
func UpdateHeartAnimation(points []models.Point, time float64, balls []models.Ball) {
	scale := 1.0 + math.Sin(time*2.0)*0.04

	for i := range points {
		p := &points[i]

		// Calculate target position with scaling
		targetX := 400 + (p.HomeX-400)*scale
		targetY := 300 + (p.HomeY-300)*scale

		// Add noise for organic movement
		pxNoise := math.Sin(float64(i%233)/233*math.Pi*2+time*2.3) * 0.35
		pyNoise := math.Cos(float64(i%317)/317*math.Pi*2+time*1.7) * 0.35

		// Apply return force to target position
		p.ApplyReturn(targetX+pxNoise, targetY+pyNoise)

		// Apply repulsion from balls
		for _, b := range balls {
			p.ApplyRepulsion(b)
		}

		// Update point position
		p.Update()
	}
}
