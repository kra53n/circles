package main

import (
	"fmt"
	"os"
	"strconv"
)

const MEASURES_DIR = "measures/"

func Measure() {
	statesNum := 100

	type Case struct {
		name   string
		method func(start, goal State) ([]State, Iters)
	}

	cases := []Case{
		Case{"bidirectional", BidirectionalSearch},
		Case{"manhatten", func(start, goal State) ([]State, Iters) { return AStarSearch(start, goal, SecondHeuristic) }},
		Case{"subtask_1col", func(start, goal State) ([]State, Iters) {
			return AStarSearch(start, goal, SubtaskHeuristicWithoutSecond)
		}},
		Case{"subtask", func(start, goal State) ([]State, Iters) { return AStarSearch(start, goal, SubtaskMaxHeuristic) }},
	}

	field := NewField()
	goal := field.Content.GetState()

	for randMoves := 1; randMoves <= 5; randMoves++ {
		states := make([]State, 0, statesNum)
		for i := 0; i < statesNum; i++ {
			state := goal.GetCopy()
			state.Content.MoveRandomly(randMoves)
			states = append(states, state)
		}

		for _, c := range cases {
			writeMeasureToFile(c.name, measure(c.method, states, goal), statesNum, randMoves)
		}
	}
}

func writeMeasureToFile(filename string, measures []Iters, statesNum int, randMoves int) {
	file, _ := os.Create(MEASURES_DIR + filename + "_" + strconv.Itoa(randMoves) + "_" + strconv.Itoa(statesNum) + ".txt")
	defer file.Close()
	for _, iters := range measures {
		fmt.Fprintf(file, "%d ", iters)
	}
	fmt.Fprintf(file, "\n")
}

func measure(
	method func(start, goal State) ([]State, Iters),
	states []State,
	goal State,
) []Iters {
	measure := make([]Iters, 0, len(states))
	for _, state := range states {
		_, iters := method(state, goal)
		measure = append(measure, iters)
	}
	return measure
}
