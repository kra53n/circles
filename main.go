package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"os"
)

const WDT = 1000
const HGT = 720

var COL_BACKGROUND = rl.Color{24, 24, 37, 255}
var COL_BOX_BACKGROUND = rl.Color{17, 17, 27, 255}
var COL_ARROW = rl.Color{88, 91, 112, 255}
var COL_CIRC1 = rl.Color{230, 15, 57, 255}
var COL_CIRC2 = rl.Color{223, 142, 29, 255}
var COL_CIRC3 = rl.Color{64, 160, 43, 255}
var COL_CIRC4 = rl.Color{30, 102, 245, 255}

var storage Storage
var randMoves int = 3

func main() {
	filename := "subtask.txt"
	storage = ReadSubtask(filename)

	if len(os.Args) > 1 && os.Args[1] == "subtask" {
		if len(os.Args) > 2 {
			if os.Args[2] == "read" {
				fmt.Println(storage)
				return
			}

			if os.Args[2] == "write" {
				err := WriteSubtask(filename, GenerateSubtask())
				if err != nil {
					fmt.Printf("Could not write to file due %s\n", err)
				}
				fmt.Printf("Subtask was written to file %s\n", filename)
				return
			}
		}
		fmt.Println("use `read` or `write` as subcommands")
		return
	}

	printUsage()

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
		if rl.IsKeyPressed(rl.KeyZero + 5) {
			processSearch(func(start, goal State) []State { return AStarSearch(start, goal, SecondHeuristic) }, "2 эвристика")
		}
		if rl.IsKeyPressed(rl.KeyZero + 6) {
			processSearch(func(start, goal State) []State { return AStarSearch(start, goal, SubtaskHeuristicWithoutSecond) }, "эвристика на основе подзадач без 2 эвристики")
		}
		if rl.IsKeyPressed(rl.KeyZero + 7) {
			processSearch(func(start, goal State) []State { return AStarSearch(start, goal, SubtaskHeuristic) }, "эвристика на основе подзадач")
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
			if len(animation.States) > 0 {
				field.Content = animation.States[0].Content
				animation.Play()
			}
			animation.Stop()
		}
		if rl.IsKeyPressed(rl.KeyC) {
			if len(animation.States) > 0 {
				field.Content = animation.States[0].Content
				animation.Play()
			}
			animation.Stop()
			field.MoveRandomly(randMoves)
		}

		printRandMoves := false
		if rl.IsKeyDown(rl.KeyLeftControl) || rl.IsKeyDown(rl.KeyRightControl) {
			if rl.IsKeyPressed(rl.KeyMinus) {
				randMoves -= 1
				printRandMoves = true
			}
			if rl.IsKeyPressed(rl.KeyEqual) {
				randMoves += 1
				printRandMoves = true
			}
		}
		if rl.IsKeyPressed(rl.KeyP) {
			printRandMoves = true
		}
		if printRandMoves {
			fmt.Println("Rand moves:", randMoves)
		}

		rl.BeginDrawing()
		{
			rl.ClearBackground(COL_BACKGROUND)
			field.Draw()
		}
		rl.EndDrawing()
	}
}

func printUsage() {
	fmt.Println("1) поиск в ширину")
	fmt.Println("2) поиск в глубину")
	fmt.Println("3) двунаправленный поиск")
	fmt.Println("4) 1 эвристика")
	fmt.Println("5) 2 эвристика")
	fmt.Println("6) эвристика на основе подзадач без 2 эвристики")
	fmt.Println("7) эвристика на основе подзадач")
	fmt.Println()
}
