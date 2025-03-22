package main

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

type InputSystem struct {
}

type SegmentUpdateSystem struct {
}

type GameDrawingSystem struct {
}

type SnakeCollisionSystem struct {
}

func (system InputSystem) Update(world *World) {
	Map4(world, func(headID ID, head *Head, pos *Position, movement *Movement, headSegment *Segment) {
		disallowedX := headSegment.LastX
		disallowedY := headSegment.LastY

		if raylib.IsKeyPressed(raylib.KeyUp) && pos.Y-1 != disallowedY {
			movement.Y = -1
			movement.X = 0
		} else if raylib.IsKeyPressed(raylib.KeyDown) && pos.Y+1 != disallowedY {
			movement.Y = 1
			movement.X = 0
		} else if raylib.IsKeyPressed(raylib.KeyLeft) && pos.X-1 != disallowedX {
			movement.Y = 0
			movement.X = -1
		} else if raylib.IsKeyPressed(raylib.KeyRight) && pos.X+1 != disallowedX {
			movement.Y = 0
			movement.X = 1
		}
	})
}

func (system GameDrawingSystem) Update(world *World) {
	raylib.BeginDrawing()
	raylib.ClearBackground(Background)
	DrawGrid()

	Map2(world, func(id ID, position *Position, color *Color) {
		DrawBox(raylib.Vector2{X: position.X, Y: position.Y}, color.Color)
	})

	raylib.EndDrawing()
}

func (system SegmentUpdateSystem) Update(world *World) {
	Map2(world, func(headID ID, head *Head, movement *Movement) {
		if world.Ticks%movement.Speed != 0 {
			return
		}

		segments := *ComponentsOfType[Segment](world)
		positions := *ComponentsOfType[Position](world)

		headPosition := positions[headID].(*Position)

		nextX := headPosition.X + movement.X
		nextY := headPosition.Y + movement.Y
		segmentID := headID

		for {
			pos := positions[segmentID].(*Position)
			nextX, pos.X = pos.X, nextX
			nextY, pos.Y = pos.Y, nextY

			segment := segments[segmentID].(*Segment)
			segment.LastX = nextX
			segment.LastY = nextY

			if segment.NextSegmentID >= 0 {
				segmentID = segment.NextSegmentID
			} else {
				break
			}
		}
	})
}

func (system SnakeCollisionSystem) Update(world *World) {
	Map2(world, func(headID ID, headPosition *Position, head *Head) {
		// Game grid bounds checking
		if headPosition.X < 0 || int32(headPosition.X) >= GridColumns || headPosition.Y < 0 || int32(headPosition.Y) >= GridRows {
			world.State = GAME_OVER
			return
		}

		// Collision with self
		if CheckCollision(world, headID) {
			world.State = GAME_OVER
			return
		}

		// Collision with food
		Map2(world, func(foodID ID, foodPosition *Position, food *Food) {
			if headPosition.X == foodPosition.X && headPosition.Y == foodPosition.Y {
				newPos := FindEmptySpot(world)
				foodPosition.X = newPos.X
				foodPosition.Y = newPos.Y

				ExtendSnake(world, headID)
			}
		})
	})
}
