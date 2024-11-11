package main

import (
	// "fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const WDT = 1280
const HGT = 720

func main() {
	rl.InitWindow(WDT, HGT, "circle")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	field := NewField()
	baseState := field.Content.GetState()
	var animation Animation

	for !rl.WindowShouldClose() {

		field.Update()

		if rl.IsKeyPressed(rl.KeyZero + 1) {
			rl.SetWindowTitle("Запущен поиск в ширину")
			states := BreadthFirstSearch(field.Content.GetState(), baseState)
			animation.Load(states)
			animation.Play()
			rl.SetWindowTitle("Завершён поиск в ширину")
		}

		if animation.Animate {
			animation.Play()
			s := animation.GetCurrState()
			field.Content = s.Content
			if s.Equals(baseState) {
				animation.Stop()
			}
		}

		if rl.IsKeyPressed(rl.KeySpace) {
			if animation.Animate {
				animation.Stop()
			} else {
				animation.Play()
			}
		}

		if rl.IsKeyPressed(rl.KeyR) && animation.States[0].Content != nil {
			field.Content = animation.States[0].Content
			animation.Play()
			animation.Stop()
		}

		rl.BeginDrawing()
		{
			rl.ClearBackground(rl.Blue)
			field.Draw()
		}
		rl.EndDrawing()
	}
}
