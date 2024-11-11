package main

// import "fmt"

type ISearch interface {
	Search()
}

type Search struct {
	states []State
	goalState State
}

func BreadthFirstSearch(start, goal State) []State {
	var statistic Statistic

	var openNodes, closedNodes []State
	openNodes = append(openNodes, start.GetCopy())

	for i := 0; len(openNodes) > 0; i++ {
		state := openNodes[0]
		openNodes = openNodes[1:]
		statistic.Collect(openNodes, closedNodes);
		if state.Equals(goal) {
			statistic.Print("Ширину");
			return state.Unwrap();
		}
		closedNodes = append(closedNodes, state)
		for _, s := range state.GenStates() {
			if !stateInStates(s, openNodes) && !stateInStates(s, closedNodes) {
				s.prv = &state;
				openNodes = append(openNodes, s)
			}
		}
	}
	return nil;
}

func DepthFirstSearch(start, goal State) []State {
	var statistic Statistic

	var openNodes, closedNodes []State
	openNodes = append(openNodes, start.GetCopy())

	for i := 0; len(openNodes) > 0; i++ {
		state := openNodes[len(openNodes)-1]
		openNodes = openNodes[:len(openNodes)-1]
		statistic.Collect(openNodes, closedNodes);
		if state.Equals(goal) {
			statistic.Print("Глубину");
			return state.Unwrap();
		}
		closedNodes = append(closedNodes, state)
		for _, s := range state.GenStates() {
			if !stateInStates(s, openNodes) && !stateInStates(s, closedNodes) {
				s.prv = &state;
				openNodes = append(openNodes, s)
			}
		}
	}
	return nil;
}

func stateInStates(s State, states []State) bool {
	for _, v := range states {
		if s.Equals(v) {
			return true
		}
	}
	return false
}
