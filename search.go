package main

// import "fmt"
import "container/heap"

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
				openNodes = append(openNodes, s)
			}
		}
	}
	return nil
}

func BidirectionalSearch(start, goal State) []State {
	var statistic BidirectionalStatistic

	var openNodes, closedNodes, openNodesR, closedNodesR, newO []State
	openNodes = append(openNodes, start.GetCopy())
	openNodesR = append(openNodesR, goal.GetCopy())

	for {
		newO = nil
		for _, node := range openNodes {
			statistic.Collect(openNodes, closedNodes)
			closedNodes = append(closedNodes, node)
			for _, n := range node.GenStates() {
				nodeReversePtr := getStateInStates(n, openNodesR)
				if nodeReversePtr != nil {
					states := UnwrapBidirectionalStates(n, *nodeReversePtr)
					statistic.Print(len(states))
					return states
				}
				if !stateInStates(n, openNodes) && !stateInStates(n, closedNodes) {
					
					newO = append(newO, n)
				}
			}
		}
		openNodes = newO
		
		newO = nil
		for _, node := range openNodesR {
			statistic.Collect(openNodesR, closedNodesR)
			closedNodesR = append(closedNodesR, node)
			for _, n := range node.GenReversedStates() {
				nodePtr := getStateInStates(n, openNodes)
				if nodePtr != nil {
					states := UnwrapBidirectionalStates(*nodePtr, n)
					statistic.Print(len(states))
					return states
				}
				if !stateInStates(n, openNodesR) && !stateInStates(n, closedNodesR) {
					newO = append(newO, n)
				}
			}
		}
		openNodesR = newO
	}
}

type PQItemSlice []*PQItem

func AStarSearch(start, goal State, h func(s State) int) []State {
	var openNodes PriorityQueue
	var closedNodes PQItemSlice
	// add start to openNodes
	
	for openNodes.Len() > 0 {
		curr := heap.Pop(&openNodes).(*PQItem)
		if curr.val.Equals(goal) {
			// statistic.pathLenght = curr.g;
			// statistic.Print();
			return curr.val.Unwrap();
		}
		// statistic.Collect(openNodes, closedNodes);
		closedNodes = append(closedNodes, curr)
		for _, n := range curr.GenStates() {
			item := openNodes.GetItem(n)
			if item != nil {
				score := h(n.val) + curr.g + 1
				if score < item.f {
					item.f = score
					item.g = curr.g + 1
					item.val.prv = &curr.val
				}
				continue
			}
			item = closedNodes.GetItem(n)
			if item != nil {
				score := h(n.val) + curr.g + 1
				if score < item.f {
					closedNodes.Remove(*item)
					item.f = score
					item.g = curr.g + 1
					item.val.prv = &curr.val
					heap.Push(&openNodes, item)
				}
				continue
			}
			n.g = curr.g + 1
			n.f = h(n.val) + n.g
			n.val.prv = &curr.val
			heap.Push(&openNodes, n)
		}
	}
	return nil
}

func (items PQItemSlice) GetItem(item PQItem) *PQItem {
	for _, i := range items {
		if item.val.Equals(i.val) {
			return i
		}
	}
	return nil
}

func (items *PQItemSlice) Remove(item PQItem) {
	for i, v := range *items {
		if item.val.Equals(v.val) {
			*items = append((*items)[:i], (*items)[i+1:]...)
			return
		}
	}
	return
}

func FirstHeuristic(s State) int {
	return 1
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
