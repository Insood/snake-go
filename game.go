package main

import (
	"math/rand"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	MaxCells = int32(GridRows * GridColumns)

	StartingSpeed = int32(22)
	MaxSpeed      = int32(5)
)

var (
	GridColor           = raylib.NewColor(218, 223, 225, 255)
	FoodColor           = raylib.NewColor(255, 0, 0, 255)
	SnakeHeadColorStart = raylib.NewColor(81, 134, 236, 255)
	SnakeHeadColorEnd   = raylib.NewColor(200, 200, 200, 255)

	Up    = raylib.NewVector2(0, -1)
	Down  = raylib.NewVector2(0, 1)
	Left  = raylib.NewVector2(-1, 0)
	Right = raylib.NewVector2(1, 0)
)

func InitializeGame(game *Game) {
	game.Snake = []raylib.Vector2{{X: float32(GridColumns) / 2, Y: float32(GridRows) / 2}}
	game.Direction = raylib.NewVector2(1, 0)
	game.Speed = StartingSpeed
	game.Ticks = 0
	game.State = GAME
	PlaceFood(game)
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
			GameBorder+col*(GridSize+GridThickness),
			GameBorder,
			GridThickness,
			ScreenHeight-GameBorder*2,
			GridColor,
		)
	}

	// Rows
	for row := int32(0); row <= GridRows; row++ {
		raylib.DrawRectangle(
			GameBorder,
			GameBorder+row*(GridSize+GridThickness),
			ScreenWidth-GameBorder*2,
			GridThickness,
			GridColor,
		)
	}
}

func DrawBox(position raylib.Vector2, color raylib.Color) {
	raylib.DrawRectangle(
		GameBorder+GridThickness+int32(position.X)*(GridSize+GridThickness),
		GameBorder+GridThickness+int32(position.Y)*(GridSize+GridThickness),
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
