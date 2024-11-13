package main

import (
	"os"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const WDT = 1000
const HGT = 720

func main() {
	if len(os.Args) > 1 && os.Args[1] == "subtask" {
		if len(os.Args) > 2 && os.Args[2] == "write" {
			filename := "subtask.txt"
			err := WriteSubtask(filename, GenerateSubtask())
			if err != nil {
				fmt.Printf("Could not write to file due %s\n", err)
			}
			fmt.Printf("Subtask was written to file %s\n", filename)
			return
		}
		GenerateSubtask()
		return
	}

	rl.SetTraceLogLevel(rl.LogNone)
	rl.SetTargetFPS(60)
	rl.InitWindow(WDT, HGT, "circle")
	defer rl.CloseWindow()

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
		if rl.IsKeyPressed(rl.KeyZero + 3) {
			processSearch(BidirectionalSearch, "двунаправленный поиск")
		}
		if rl.IsKeyPressed(rl.KeyZero + 4) {
			processSearch(func(start, goal State) []State { return AStarSearch(start, goal, FirstHeuristic) }, "1 эвристика")
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
