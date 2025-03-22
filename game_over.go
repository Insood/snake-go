package main

import raylib "github.com/gen2brain/raylib-go/raylib"

func GameOverScreen(world *World) {
	ReadGameOverInputs(world)
	DrawGameOverScreen(world)
}

func ReadGameOverInputs(world *World) {
	if raylib.IsKeyPressed(raylib.KeySpace) {
		world.State = SPLASH_SCREEN
	}
}

func DrawGameOverScreen(world *World) {
	raylib.BeginDrawing()
	raylib.ClearBackground(Background)

	startText := "Press Space to Restart"
	textWidth := raylib.MeasureText(startText, 20)

	raylib.DrawText(startText, ScreenWidth/2-textWidth/2, 255, 20, Gray)

	raylib.EndDrawing()
}
