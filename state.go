package main

import (
	"slices"
)

type State struct {
	Content Content
	prv     *State
}

func (s *State) Unwrap() []State {
	var states []State
	for it := s; it != nil; it = it.prv {
		states = append(states, *it)
	}
	slices.Reverse(states)
	return states
}

func (s *State) GenStates() []State {
	states := make([]State, 0, len(s.Content)*2)
	for j := 0; j < 2; j++ {
		for i := 0; i < len(s.Content); i++ {
			state := s.GetCopy()
			switch j {
			case 0:
				state.Content.moveCol(i)
			case 1:
				state.Content.moveRow(i)
			}
			states = append(states, state)
		}
	}
	return states
}

func (s *State) GenStatesReverse() []State {
	states := make([]State, 0, len(s.Content)*2)
	for j := 0; j < 2; j++ {
		for i := 0; i < len(s.Content); i++ {
			state := s.GetCopy()
			switch j {
			case 0:
				state.Content.moveColReverse(i)
			case 1:
				state.Content.moveRowReverse(i)
			}
			states = append(states, state)
		}
	}
	return states
}

func (s1 *State) GetCopy() State {
	var s2 State
	s2.Content = make([][]byte, len(s1.Content))
	for i, row := range s1.Content {
		s2.Content[i] = make([]byte, len(s1.Content))
		for j, v := range row {
			s2.Content[i][j] = v
		}
	}
	return s2
}

func (s1 *State) Equals(s2 State) bool {
	for i, _ := range s1.Content {
		for j, _ := range s2.Content {
			if s1.Content[i][j] != s2.Content[i][j] {
				return false
			}
		}
	}
	return true
}

func (c Content) GetState() State {
	var s State
	s.Content = c
	return s.GetCopy()
}
