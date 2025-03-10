package main

import (
	"math/rand"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	gridRows              = int32(10)
	gridColumns           = int32(10)
	maxCells              = int32(gridRows * gridColumns)
	gridThickness         = int32(2)
	gridSize              = int32(25)
	gridBorder            = int32(10)
	screenWidth           = gridColumns*gridSize + (gridColumns+1)*gridThickness + 2*gridBorder
	screenHeight          = gridRows*gridSize + (gridRows+1)*gridThickness + 2*gridBorder
	splashScreenGrid      = 10
	splashScreenGridSize  = int32(splashScreenGrid * gridSize)
	splashScreenGridCols  = int32(16)
	splashScreenGridRows  = int32(16)
	splashScreenGridCount = int32(2*splashScreenGridCols + 2*splashScreenGridRows - 4)
	startingSpeed         = int32(22)
	maxSpeed              = int32(5)
)

var (
	gridColor           = raylib.NewColor(218, 223, 225, 255)
	foodColor           = raylib.NewColor(255, 0, 0, 255)
	snakeHeadColorStart = raylib.NewColor(81, 134, 236, 255)
	snakeHeadColorEnd   = raylib.NewColor(200, 200, 200, 255)

	foreground = raylib.NewColor(0, 68, 130, 255)
	background = raylib.NewColor(255, 255, 255, 255)
	gray       = raylib.NewColor(128, 128, 128, 255)

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
	game.Snake = []raylib.Vector2{{X: float32(gridColumns) / 2, Y: float32(gridRows) / 2}}
	game.Direction = raylib.NewVector2(1, 0)
	game.Speed = startingSpeed
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
	raylib.ClearBackground(background)

	upperLeftX := screenWidth/2 - splashScreenGridCols/2*splashScreenGridSize
	upperLeftY := screenHeight/2 - splashScreenGridRows/2*splashScreenGridSize

	raylib.DrawRectangle(
		upperLeftX,
		upperLeftY,
		splashScreenGridCols*splashScreenGridSize,
		splashScreenGridRows*splashScreenGridSize,
		foreground,
	)

	// raylib.DrawRectangle(
	// 	screenWidth/2-splashScreenGridCols/2*splashScreenGridSize+splashScreenGridSize,
	// 	screenHeight/2-splashScreenGridRows/2*splashScreenGridSize+splashScreenGridSize,
	// 	(splashScreenGridCols-2)*splashScreenGridSize,
	// 	(splashScreenGridRows-2)*splashScreenGridSize,
	// 	background,
	// )

	// for i := int32(0); i < 6; i++ {
	// 	snakeOffset := (game.Ticks/5 + i) % splashScreenGridCount
	// 	gridPosition := GetOuroborosPosition(snakeOffset)
	// 	raylib.DrawRectangle(
	// 		upperLeftX+int32(gridPosition.X)*splashScreenGridSize+1,
	// 		upperLeftY+int32(gridPosition.Y)*splashScreenGridSize+1,
	// 		splashScreenGridSize-2,
	// 		splashScreenGridSize-2,
	// 		background,
	// 	)
	// }

	startText := "Press Space to Start"
	textWidth := raylib.MeasureText(startText, 20)

	raylib.DrawText("snek", 134, 182, 30, foreground)
	raylib.DrawText(startText, screenWidth/2-textWidth/2, 255, 20, gray)

	raylib.EndDrawing()
}

func GetOuroborosPosition(i int32) raylib.Vector2 {
	if i < splashScreenGridCols {
		return raylib.NewVector2(float32(i), 0.0)
	} else if i < splashScreenGridCols+splashScreenGridRows-1 {
		return raylib.NewVector2(float32(splashScreenGridCols-1), float32(i-splashScreenGridCols+1))
	} else if i < 2*splashScreenGridCols+splashScreenGridRows-2 {
		return raylib.NewVector2(float32(2*splashScreenGridCols+splashScreenGridRows-3-i), float32(splashScreenGridRows-1))
	} else {
		return raylib.NewVector2(0.0, float32(2*splashScreenGridCols+2*splashScreenGridRows-4-i))
	}
}

func GameScreen(game *Game) {
	ReadGameInputs(game)
	DrawGameScreen(game)
	UpdateGameState(game)
}

func DrawGrid() {
	// Columns
	for col := int32(0); col <= gridColumns; col++ {
		raylib.DrawRectangle(
			gridBorder+col*(gridSize+gridThickness),
			gridBorder,
			gridThickness,
			screenHeight-gridBorder*2,
			gridColor,
		)
	}

	// Rows
	for row := int32(0); row <= gridRows; row++ {
		raylib.DrawRectangle(
			gridBorder,
			gridBorder+row*(gridSize+gridThickness),
			screenWidth-gridBorder*2,
			gridThickness,
			gridColor,
		)
	}
}

func DrawBox(position raylib.Vector2, color raylib.Color) {
	raylib.DrawRectangle(
		gridBorder+gridThickness+int32(position.X)*(gridSize+gridThickness),
		gridBorder+gridThickness+int32(position.Y)*(gridSize+gridThickness),
		gridSize,
		gridSize,
		color,
	)
}

func DrawSnake(game *Game) {
	for i, segment := range game.Snake {
		segmentColor := raylib.ColorLerp(snakeHeadColorStart, snakeHeadColorEnd, float32(i)/float32(maxCells))
		DrawBox(segment, segmentColor)
	}
}

func DrawGameScreen(game *Game) {
	raylib.BeginDrawing()
	raylib.ClearBackground(background)
	DrawGrid()
	DrawSnake(game)
	DrawBox(game.Food, foodColor)
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
		col := float32(rand.Intn(int(gridColumns)))
		row := float32(rand.Intn(int(gridRows)))

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
		game.Speed = startingSpeed - int32(float32(len(game.Snake)*1.0)/float32(maxCells)*float32(startingSpeed-maxSpeed))
	} else if nextPosition.X < 0 || nextPosition.X >= float32(gridColumns) || nextPosition.Y < 0 || nextPosition.Y >= float32(gridRows) || CheckCollision(nextPosition, game) {
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
	raylib.ClearBackground(background)

	startText := "Press Space to Restart"
	textWidth := raylib.MeasureText(startText, 20)

	raylib.DrawText(startText, screenWidth/2-textWidth/2, 255, 20, gray)

	raylib.EndDrawing()
}

func main() {
	raylib.InitWindow(screenWidth, screenHeight, "Snek")
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
