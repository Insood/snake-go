package main

import raylib "github.com/gen2brain/raylib-go/raylib"

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
