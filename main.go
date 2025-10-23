package main

import (
	"heart/internal/config"
	"heart/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// Load configuration
	cfg := config.DefaultConfig()

	// Create game instance
	gameInstance := game.NewGame(cfg)

	// Set up window
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle("Heart <3")

	// Run the game
	if err := ebiten.RunGame(gameInstance); err != nil {
		panic(err)
	}
}
