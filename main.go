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

	processSearch := func(searchMethod func(start, goal State) []State, name string) {
		rl.SetWindowTitle("Запущен " + name)
		states := searchMethod(field.Content.GetState(), baseState)
		animation.Load(states)
		animation.Play()
		rl.SetWindowTitle("Завершён " + name)
	}

	for !rl.WindowShouldClose() {

		field.Update()

		if rl.IsKeyPressed(rl.KeyZero + 1) {
			processSearch(BreadthFirstSearch, "поиск в ширину")
		}
		if rl.IsKeyPressed(rl.KeyZero + 2) {
			processSearch(DepthFirstSearch, "поиск в глубину")
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
