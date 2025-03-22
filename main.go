package main

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	GridRows      = int32(10)
	GridColumns   = int32(10)
	GridThickness = int32(2)
	GridSize      = int32(25)
	GameBorder    = int32(10)
	ScreenWidth   = GridColumns*GridSize + (GridColumns+1)*GridThickness + 2*GameBorder
	ScreenHeight  = GridRows*GridSize + (GridRows+1)*GridThickness + 2*GameBorder
)

var (
	Background = raylib.NewColor(255, 255, 255, 255)
	Gray       = raylib.NewColor(128, 128, 128, 255)
)

func main() {
	raylib.InitWindow(ScreenWidth, ScreenHeight, "Snek")
	defer raylib.CloseWindow()

	raylib.SetTargetFPS(60)

	world := NewWorld()

	for !raylib.WindowShouldClose() {
		if world.State == SPLASH_SCREEN {
			SplashScreen(world)
		} else if world.State == GAME {
			GameScreen(world)
		} else if world.State == GAME_OVER {
			GameOverScreen(world)
		}

		world.Ticks += 1
	}
}
