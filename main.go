package main

import (
	"math/rand"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	GridRows              = int32(10)
	GridColumns           = int32(10)
	MaxCells              = int32(GridRows * GridColumns)
	GridThickness         = int32(2)
	GridSize              = int32(25)
	GridBroder            = int32(10)
	ScreenWidth           = GridColumns*GridSize + (GridColumns+1)*GridThickness + 2*GridBroder
	ScreenHeight          = GridRows*GridSize + (GridRows+1)*GridThickness + 2*GridBroder
	SplashScreenGrid      = 10
	SplashScreenGridSize  = int32(SplashScreenGrid * GridSize)
	SplashScreenGridCols  = int32(16)
	SplashScreenGridRows  = int32(16)
	SplashScreenGridCount = int32(2*SplashScreenGridCols + 2*SplashScreenGridRows - 4)
	StartingSpeed         = int32(22)
	MaxSpeed              = int32(5)
)

var (
	GridColor           = raylib.NewColor(218, 223, 225, 255)
	FoodColor           = raylib.NewColor(255, 0, 0, 255)
	SnakeHeadColorStart = raylib.NewColor(81, 134, 236, 255)
	SnakeHeadColorEnd   = raylib.NewColor(200, 200, 200, 255)

	Foreground = raylib.NewColor(0, 68, 130, 255)
	Background = raylib.NewColor(255, 255, 255, 255)
	Gray       = raylib.NewColor(128, 128, 128, 255)

	Up    = raylib.NewVector2(0, -1)
	Down  = raylib.NewVector2(0, 1)
	Left  = raylib.NewVector2(-1, 0)
	Right = raylib.NewVector2(1, 0)
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

func InitializeGame(game *Game) {
	game.Snake = []raylib.Vector2{{X: float32(GridColumns) / 2, Y: float32(GridRows) / 2}}
	game.Direction = raylib.NewVector2(1, 0)
	game.Speed = StartingSpeed
	game.Ticks = 0
	game.State = GAME
	PlaceFood(game)
}

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

	// raylib.DrawRectangle(
	// 	upperLeftX,
	// 	upperLeftY,
	// 	SplashScreenGridCols*SplashScreenGridSize,
	// 	SplashScreenGridRows*SplashScreenGridSize,
	// 	Foreground,
	// )

	// // raylib.DrawRectangle(
	// // 	ScreenWidth/2-SplashScreenGridCols/2*SplashScreenGridSize+SplashScreenGridSize,
	// // 	ScreenHeight/2-SplashScreenGridRows/2*SplashScreenGridSize+SplashScreenGridSize,
	// // 	(SplashScreenGridCols-2)*SplashScreenGridSize,
	// // 	(SplashScreenGridRows-2)*SplashScreenGridSize,
	// // 	Background,
	// // )

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

func GameScreen(game *Game) {
	ReadGameInputs(game)
	DrawGameScreen(game)
	UpdateGameState(game)
}

func DrawGrid() {
	// Columns
	for col := int32(0); col <= GridColumns; col++ {
		raylib.DrawRectangle(
			GridBroder+col*(GridSize+GridThickness),
			GridBroder,
			GridThickness,
			ScreenHeight-GridBroder*2,
			GridColor,
		)
	}

	// Rows
	for row := int32(0); row <= GridRows; row++ {
		raylib.DrawRectangle(
			GridBroder,
			GridBroder+row*(GridSize+GridThickness),
			ScreenWidth-GridBroder*2,
			GridThickness,
			GridColor,
		)
	}
}

func DrawBox(position raylib.Vector2, color raylib.Color) {
	raylib.DrawRectangle(
		GridBroder+GridThickness+int32(position.X)*(GridSize+GridThickness),
		GridBroder+GridThickness+int32(position.Y)*(GridSize+GridThickness),
		GridSize,
		GridSize,
		color,
	)
}

func DrawSnake(game *Game) {
	for i, segment := range game.Snake {
		segmentColor := raylib.ColorLerp(SnakeHeadColorStart, SnakeHeadColorEnd, float32(i)/float32(MaxCells))
		DrawBox(segment, segmentColor)
	}
}

func DrawGameScreen(game *Game) {
	raylib.BeginDrawing()
	raylib.ClearBackground(Background)
	DrawGrid()
	DrawSnake(game)
	DrawBox(game.Food, FoodColor)
	raylib.EndDrawing()
}

func ReadGameInputs(game *Game) {
	if raylib.IsKeyPressed(raylib.KeyUp) {
		game.LastInput = Up
	} else if raylib.IsKeyPressed(raylib.KeyDown) && game.Direction != Up {
		game.LastInput = Down
	} else if raylib.IsKeyPressed(raylib.KeyLeft) && game.Direction != Right {
		game.LastInput = Left
	} else if raylib.IsKeyPressed(raylib.KeyRight) && game.Direction != Left {
		game.LastInput = Right
	}
}

func CheckCollision(position raylib.Vector2, game *Game) bool {
	for _, segment := range game.Snake {
		if segment.X == position.X && segment.Y == position.Y {
			return true
		}
	}

	return false
}

func PlaceFood(game *Game) {
	for {
		col := float32(rand.Intn(int(GridColumns)))
		row := float32(rand.Intn(int(GridRows)))

		foodPosition := raylib.NewVector2(col, row)

		if !CheckCollision(foodPosition, game) {
			game.Food = foodPosition
			return
		}
	}
}

func UpdateGameState(game *Game) {
	game.Ticks++

	if game.Ticks < game.Speed {
		return
	}

	if game.LastInput == Left && game.Direction != Right {
		game.Direction = Left
	} else if game.LastInput == Right && game.Direction != Left {
		game.Direction = Right
	} else if game.LastInput == Up && game.Direction != Down {
		game.Direction = Up
	} else if game.LastInput == Down && game.Direction != Up {
		game.Direction = Down
	}

	game.Ticks = 0

	nextPosition := raylib.Vector2Add(game.Snake[0], game.Direction)

	if nextPosition == game.Food {
		game.Snake = append([]raylib.Vector2{nextPosition}, game.Snake...)
		PlaceFood(game)
		game.Speed = StartingSpeed - int32(float32(len(game.Snake)*1.0)/float32(MaxCells)*float32(StartingSpeed-MaxSpeed))
	} else if nextPosition.X < 0 || nextPosition.X >= float32(GridColumns) || nextPosition.Y < 0 || nextPosition.Y >= float32(GridRows) || CheckCollision(nextPosition, game) {
		game.State = GAME_OVER
	} else {
		game.Snake = append([]raylib.Vector2{nextPosition}, game.Snake...)
		game.Snake = game.Snake[:len(game.Snake)-1]
	}
}

func GameOverScreen(game *Game) {
	ReadGameOverInputs(game)
	DrawGameOverScreen(game)
}

func ReadGameOverInputs(game *Game) {
	if raylib.IsKeyPressed(raylib.KeySpace) {
		game.State = SPLASH_SCREEN
	}
}

func DrawGameOverScreen(game *Game) {
	raylib.BeginDrawing()
	raylib.ClearBackground(Background)

	startText := "Press Space to Restart"
	textWidth := raylib.MeasureText(startText, 20)

	raylib.DrawText(startText, ScreenWidth/2-textWidth/2, 255, 20, Gray)

	raylib.EndDrawing()
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
