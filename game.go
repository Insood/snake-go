package main

import (
	"math/rand"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	MaxCells = int32(GridRows * GridColumns)

	StartingSpeed = int64(22)
	MaxSpeed      = int64(5)
)

var (
	GridColor           = raylib.NewColor(218, 223, 225, 255)
	FoodColor           = raylib.NewColor(255, 0, 0, 255)
	SnakeHeadColorStart = raylib.NewColor(81, 134, 236, 255)
	SnakeHeadColorEnd   = raylib.NewColor(200, 200, 200, 255)
)

func NewSnake(world *World) {
	id := world.AllocateID()
	pos := NewPosition(float32(GridColumns)/2, float32(GridRows)/2)
	dir := NewMovement(StartingSpeed, 1, 0)
	color := NewColor(SnakeHeadColorStart)
	head := NewHead()
	segment := NewSegment()

	world.AddComponent(id, pos)
	world.AddComponent(id, dir)
	world.AddComponent(id, color)
	world.AddComponent(id, head)
	world.AddComponent(id, segment)
}

func CreateFood(world *World) {
	id := world.AllocateID()
	foodColor := NewColor(FoodColor)
	pos := FindEmptySpot(world)
	food := NewFood()

	world.AddComponent(id, food)
	world.AddComponent(id, pos)
	world.AddComponent(id, foodColor)
}

func CreateSegment(world *World, X float32, Y float32, segmentCount int) ID {
	id := world.AllocateID()
	pos := NewPosition(X, Y)

	segmentColor := raylib.ColorLerp(SnakeHeadColorStart, SnakeHeadColorEnd, float32(segmentCount)*1.0/float32(MaxCells))

	color := NewColor(segmentColor)
	segment := NewSegment()
	world.AddComponent(id, pos)
	world.AddComponent(id, color)
	world.AddComponent(id, segment)
	return id
}

func InitializeGame(world *World) {
	world.Nuke()
	world.State = GAME

	NewSnake(world)
	CreateFood(world)

	world.State = GAME
	world.AddSystem(InputSystem{})
	world.AddSystem(SnakeCollisionSystem{})
	world.AddSystem(GameDrawingSystem{})
	world.AddSystem(SegmentUpdateSystem{})
}

func GameScreen(world *World) {
	world.UpdateSystems()
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

func CheckCollision(world *World, headID ID) bool {
	positions := *ComponentsOfType[Position](world)
	headX := positions[headID].(*Position).X
	headY := positions[headID].(*Position).Y
	collision := false

	Map2(world, func(componentID ID, segment *Segment, position *Position) {
		if componentID != headID && position.X == headX && position.Y == headY {
			collision = true
		}
	})
	return collision
}

func FindEmptySpot(world *World) *Position {
	positions := *ComponentsOfType[Position](world)

	for {
		X := float32(rand.Intn(int(GridColumns)))
		Y := float32(rand.Intn(int(GridRows)))

		overlap := false

		for _, occupied_position := range positions {
			occupied_position := occupied_position.(*Position)

			if occupied_position.X == X && occupied_position.Y == Y {
				overlap = true
				break
			}
		}

		if !overlap {
			return NewPosition(X, Y)
		}
	}
}

// Find the last segment of the snake and then get that segment's last position
// This is the position where we want to insert an additional segment
func ExtendSnake(world *World, headID ID) {
	segments := *ComponentsOfType[Segment](world)

	segmentID := headID

	count := 0
	for {
		count += 1
		thisSegment := segments[segmentID].(*Segment)

		if thisSegment.NextSegmentID < 0 {
			X := thisSegment.LastX
			Y := thisSegment.LastY

			id := CreateSegment(world, X, Y, count)
			thisSegment.NextSegmentID = id

			break
		}

		segmentID = thisSegment.NextSegmentID
	}
}
