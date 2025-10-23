package config

import "heart/internal/models"

// DefaultConfig returns the default game configuration
func DefaultConfig() models.Config {
	return models.Config{
		ScreenWidth:     800,
		ScreenHeight:    600,
		HeartPointCount: 5000,
		AutoSpawnDelay:  2.0,
		BallSpawnCount:  3,
		BallRadius:      4.0,
		BallTTL:         5.0,
		FlyingBallTTL:   8.0,
	}
}
