package main

import (
	"fmt"
	"os"
	"strconv"
)

// const MEASURES_DIR = "measures/"
const MEASURES_DIR = "measures_test/"

func Measure() {
	statesNum := 100

	type Case struct {
		name   string
		method func(start, goal State) ([]State, Statistic)
	}

	cases := []Case{
		Case{"bidirectional", BidirectionalSearch},
		Case{"manhatten", func(start, goal State) ([]State, Statistic) { return AStarSearch(start, goal, SecondHeuristic) }},
		Case{"subtask_1col", func(start, goal State) ([]State, Statistic) {
			return AStarSearch(start, goal, SubtaskHeuristicWithoutSecond)
		}},
		Case{"subtask", func(start, goal State) ([]State, Statistic) { return AStarSearch(start, goal, SubtaskMaxHeuristic) }},
	}

	field := NewField()
	goal := field.Content.GetState()

	for randMoves := 1; randMoves <= 5; randMoves++ {
		states := make([]State, 0, statesNum)
		for i := 0; i < statesNum; i++ {
			state := goal.GetCopy()
			state.Content.MoveRandomlyReversed(randMoves)
			states = append(states, state)
		}

		for _, c := range cases {
			writeMeasureToFile(c.name, measure(c.method, states, goal), statesNum, randMoves)
		}
	}
}

func writeMeasureToFile(filename string, measures []Statistic, statesNum int, randMoves int) {
	file, _ := os.Create(MEASURES_DIR + filename + "_" + strconv.Itoa(randMoves) + "_" + strconv.Itoa(statesNum) + ".txt")
	defer file.Close()

	fmt.Fprintf(file, "iters ")
	for _, stat := range measures {
		fmt.Fprintf(file, "%d ", stat.iters)
	}
	fmt.Fprintf(file, "\n")

	fmt.Fprintf(file, "maxOpenNodesNum ")
	for _, stat := range measures {
		fmt.Fprintf(file, "%d ", stat.maxOpenNodesNum)
	}
	fmt.Fprintf(file, "\n")

	fmt.Fprintf(file, "maxClosedNodesNum ")
	for _, stat := range measures {
		fmt.Fprintf(file, "%d ", stat.maxClosedNodesNum)
	}
	fmt.Fprintf(file, "\n")
}

func measure(
	method func(start, goal State) ([]State, Statistic),
	states []State,
	goal State,
) []Statistic {
	measure := make([]Statistic, 0, len(states))
	for _, state := range states {
		_, statistic := method(state, goal)
		measure = append(measure, statistic)
	}
	return measure
}
