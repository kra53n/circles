package main

import "fmt"

type Subtask struct {
	opt []int
	v   int
}

func GenerateSubtasks() []Subtask {
	subtasks := make([]Subtask, 0, 1820)
	var goal State
	goal.Content = [][]byte{
		{0, 1, 1, 1},
		{0, 1, 1, 1},
		{0, 1, 1, 1},
		{0, 1, 1, 1},
	}
	opts := GenerateOpts()
	for _, opt := range opts {
		start := optToState(opt)
		var states []State
		if !start.Equals(goal) {
			states = BidirectionalSearch(start, goal)
		}
		subtasks = append(subtasks, Subtask{opt: opt, v: len(states) - 1})
	}
	return subtasks
}

func GenerateOpts() [][]int {
	opts := make([][]int, 0, 1820)
	for c1 := 0; c1 < 16; c1++ {
		for c2 := c1; c2 < 16; c2++ {
			for c3 := c2; c3 < 16; c3++ {
				for c4 := c3; c4 < 16; c4++ {
					opt := []int{c1, c2, c3, c4}
					if differs(opt[0], opt[1:]) {
						opts = append(opts, opt)
					}
				}
			}
		}
	}
	return opts
}

func differs(elem int, elems []int) bool {
	if len(elems) == 0 {
		return true
	}
	for _, i := range elems {
		if elem == i {
			return false
		}
	}
	return differs(elems[0], elems[1:])
}

func printOpt(opt []int) {
	for _, v := range opt {
		fmt.Print(v)
		if v <= 9 {
			fmt.Print(" ")
		}
		fmt.Print(" ")
	}
}

func optToState(opt []int) State {
	var content Content = make([][]byte, 0, 4)
	for i := 0; i < 4; i += 1 {
		content = append(content, make([]byte, 4))
		for j := 0; j < 4; j += 1 {
			content[i][j] = 1
		}
	}
	for _, v := range opt {
		content[v/4][v%4] = 0
	}
	return State{Content: content}
}
