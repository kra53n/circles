package main

import (
	_ "fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const WDT = 1280
const HGT = 720

func main() {
	rl.InitWindow(WDT, HGT, "circle")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	field := NewField()
	var baseState State
	baseState.Content = field.Content
	baseState.genStates(field.Size)

	for !rl.WindowShouldClose() {

		field.Update()
		rl.BeginDrawing()
		{
			rl.ClearBackground(rl.Blue)
			field.Draw()
		}
		rl.EndDrawing()
	}
}
