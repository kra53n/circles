package main

// import "fmt"

type ISearch interface {
	Search()
}

type Search struct {
	states    []State
	goalState State
}

func BreadthFirstSearch(start, goal State) []State {
	var statistic Statistic

	var openNodes, closedNodes []State
	openNodes = append(openNodes, start.GetCopy())

	for i := 0; len(openNodes) > 0; i++ {
		state := openNodes[0]
		openNodes = openNodes[1:]
		statistic.Collect(openNodes, closedNodes)
		if state.Equals(goal) {
			statistic.Print("Ширину")
			return state.Unwrap()
		}
		closedNodes = append(closedNodes, state)
		for _, s := range state.GenStates() {
			if !stateInStates(s, openNodes) && !stateInStates(s, closedNodes) {
				s.prv = &state
				openNodes = append(openNodes, s)
			}
		}
	}
	return nil
}

func DepthFirstSearch(start, goal State) []State {
	var statistic Statistic

	var openNodes, closedNodes []State
	openNodes = append(openNodes, start.GetCopy())

	for i := 0; len(openNodes) > 0; i++ {
		state := openNodes[len(openNodes)-1]
		openNodes = openNodes[:len(openNodes)-1]
		statistic.Collect(openNodes, closedNodes)
		if state.Equals(goal) {
			statistic.Print("Глубину")
			return state.Unwrap()
		}
		closedNodes = append(closedNodes, state)
		for _, s := range state.GenStates() {
			if !stateInStates(s, openNodes) && !stateInStates(s, closedNodes) {
				s.prv = &state
				openNodes = append(openNodes, s)
			}
		}
	}
	return nil
}

func BidirectionalSearch(start, goal State) []State {
	var openNodes, closedNodes, openNodesR, closedNodesR, newO []State
	openNodes = append(openNodes, start.GetCopy())
	openNodesR = append(openNodes, goal.GetCopy())

	for {
		newO = nil
		for _, node := range openNodes {
			// collection here statistic
			closedNodes = append(closedNodes, node)
			for _, n := range node.GenStates() {
				nodeReversePtr := getStateInStates(n, openNodesR)
				if nodeReversePtr != nil {
					return UnwrapBidirectionalStates(n, *nodeReversePtr)
				}
				if !stateInStates(n, openNodes) && !stateInStates(n, closedNodes) {
					n.prv = &node
					newO = append(openNodes, n)
				}
			}
		}
		openNodes = newO
		
		newO = nil
		for _, node := range openNodesR {
			// collection here statistic
			closedNodesR = append(closedNodesR, node)
			for _, n := range node.GenReversedStates() {
				nodePtr := getStateInStates(n, openNodes)
				if nodePtr != nil {
					return UnwrapBidirectionalStates(*nodePtr, n)
				}
				if !stateInStates(n, openNodesR) && !stateInStates(n, closedNodesR) {
					n.prv = &node
					newO = append(openNodes, n)
				}
			}
		}
		openNodesR = newO
	}
}


func stateInStates(s State, states []State) bool {
	for _, v := range states {
		if s.Equals(v) {
			return true
		}
	}
	return false
}

func getStateInStates(s State, states []State) *State {
	for _, v := range states {
		if s.Equals(v) {
			return &v
		}
	}
	return nil
}
