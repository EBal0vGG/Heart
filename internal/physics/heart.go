package physics

import (
	"math"
	"math/rand"

	"heart/internal/models"
)

// GenerateHeartPoints creates points that form a heart shape
func GenerateHeartPoints(count int) []models.Point {
	points := make([]models.Point, 0, count)
	for i := 0; i < count; i++ {
		t := rand.Float64() * 2 * math.Pi
		r := rand.Float64()
		x := 16 * math.Pow(math.Sin(t), 3) * r
		y := (13*math.Cos(t) - 5*math.Cos(2*t) - 2*math.Cos(3*t) - math.Cos(4*t)) * r

		px := 400 + x*10
		py := 300 - y*10

		points = append(points, models.Point{
			X: px, Y: py, HomeX: px, HomeY: py,
		})
	}
	return points
}
