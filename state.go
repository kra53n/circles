package main

// import "fmt"

type State struct {
	Content Content
}

func (s *State) genStates(size int) []State {
	states := make([]State, 0, size*2)
	for j := 0; j < 2; j++ {
		for i := 0; i < len(s.Content); i++ {
			state := s.getCopy()
			switch j {
			case 0:
				state.Content.moveCol(i, size)
			case 1:
				state.Content.moveRow(i, size)
			}
			states = append(states, state)
		}
	}
	return states
}

func (s1 *State) getCopy() State {
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

func (s1 *State) equals(s2 *State) bool {
	return false
}
