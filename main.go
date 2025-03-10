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

type GameState int

const (
	SPLASH_SCREEN GameState = iota
	GAME
	GAME_OVER
)

type Game struct {
	Snake     []raylib.Vector2
	Direction raylib.Vector2
	LastInput raylib.Vector2
	Food      raylib.Vector2
	Speed     int32
	Ticks     int32
	State     GameState
}

func main() {
	raylib.InitWindow(ScreenWidth, ScreenHeight, "Snek")
	defer raylib.CloseWindow()

	raylib.SetTargetFPS(60)

	game := Game{}
	game.State = SPLASH_SCREEN

	for !raylib.WindowShouldClose() {
		if game.State == SPLASH_SCREEN {
			SplashScreen(&game)
		} else if game.State == GAME {
			GameScreen(&game)
		} else if game.State == GAME_OVER {
			GameOverScreen(&game)
		}
	}
}
