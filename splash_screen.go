package main

import raylib "github.com/gen2brain/raylib-go/raylib"

const (
	SplashScreenGrid      = 10
	SplashScreenGridSize  = int32(10)
	SplashScreenGridCols  = int32(16)
	SplashScreenGridRows  = int32(16)
	SplashScreenGridCount = int32(2*SplashScreenGridCols + 2*SplashScreenGridRows - 4)
)

var (
	Foreground = raylib.NewColor(0, 68, 130, 255)
)

func SplashScreen(game *Game) {
	ReadInputsSplashScreen(game)
	UpdateSplashScreenState(game)
	DrawSplashScreen(game)
}

func ReadInputsSplashScreen(game *Game) {
	if raylib.IsKeyPressed(raylib.KeySpace) {
		InitializeGame(game)
	}
}

func UpdateSplashScreenState(game *Game) {
	game.Ticks++
}

func DrawSplashScreen(game *Game) {
	raylib.BeginDrawing()
	raylib.ClearBackground(Background)

	upperLeftX := ScreenWidth/2 - (SplashScreenGridCols/2)*SplashScreenGridSize
	upperLeftY := ScreenHeight/2 - (SplashScreenGridRows/2)*SplashScreenGridSize

	raylib.DrawRectangle(
		upperLeftX,
		upperLeftY,
		SplashScreenGridCols*SplashScreenGridSize,
		SplashScreenGridRows*SplashScreenGridSize,
		Foreground,
	)

	raylib.DrawRectangle(
		ScreenWidth/2-SplashScreenGridCols/2*SplashScreenGridSize+SplashScreenGridSize,
		ScreenHeight/2-SplashScreenGridRows/2*SplashScreenGridSize+SplashScreenGridSize,
		(SplashScreenGridCols-2)*SplashScreenGridSize,
		(SplashScreenGridRows-2)*SplashScreenGridSize,
		Background,
	)

	for i := int32(0); i < 6; i++ {
		snakeOffset := (game.Ticks/5 + i) % SplashScreenGridCount
		gridPosition := GetOuroborosPosition(snakeOffset)
		raylib.DrawRectangle(
			upperLeftX+int32(gridPosition.X)*SplashScreenGridSize+1,
			upperLeftY+int32(gridPosition.Y)*SplashScreenGridSize+1,
			SplashScreenGridSize-2,
			SplashScreenGridSize-2,
			Background,
		)
	}

	startText := "Press Space to Start"
	textWidth := raylib.MeasureText(startText, 20)

	raylib.DrawText("snek", 134, 182, 30, Foreground)
	raylib.DrawText(startText, ScreenWidth/2-textWidth/2, 255, 20, Gray)

	raylib.EndDrawing()
}

func GetOuroborosPosition(i int32) raylib.Vector2 {
	if i < SplashScreenGridCols {
		return raylib.NewVector2(float32(i), 0.0)
	} else if i < SplashScreenGridCols+SplashScreenGridRows-1 {
		return raylib.NewVector2(float32(SplashScreenGridCols-1), float32(i-SplashScreenGridCols+1))
	} else if i < 2*SplashScreenGridCols+SplashScreenGridRows-2 {
		return raylib.NewVector2(float32(2*SplashScreenGridCols+SplashScreenGridRows-3-i), float32(SplashScreenGridRows-1))
	} else {
		return raylib.NewVector2(0.0, float32(2*SplashScreenGridCols+2*SplashScreenGridRows-4-i))
	}
}
