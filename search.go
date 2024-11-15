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
			states := state.Unwrap()
			statistic.Print("Ширину", len(states)-1)
			return states
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
			states := state.Unwrap()
			statistic.Print("Глубину", len(states))
			return states
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

type PQItem struct {
	val State
	f   int
	g   int
}

type PQItemSlice []PQItem

func AStarSearch(start, goal State, h func(start, goal State) int) []State {
	var statistic Statistic

	var openNodes PQItemSlice
	var closedNodes PQItemSlice

	openNodes = add(openNodes, PQItem{val: start, f: h(start, goal)})

	for len(openNodes) > 0 {
		curr := openNodes[0]
		openNodes = openNodes[1:]
		statistic.CollectHeuristic(openNodes, closedNodes)
		if curr.val.Equals(goal) {
			states := curr.val.Unwrap()
			statistic.Print("Эвристика на основе подзадач", len(states))
			return states
		}
		closedNodes = append(closedNodes, curr)
		for _, n := range curr.GenStates() {
			item := get(openNodes, n)
			if item != nil {
				score := h(n.val, goal) + curr.g + 1
				if score < item.f {
					item.f = score
					item.g = curr.g + 1
					item.val.prv = &curr.val
				}
				continue
			}
			item = get(closedNodes, n)
			if item != nil {
				score := h(n.val, goal) + curr.g + 1
				if score < item.f {
					closedNodes = remove(closedNodes, *item)
					item.f = score
					item.g = curr.g + 1
					item.val.prv = &curr.val
					openNodes = add(openNodes, *item)
				}
				continue
			}
			n.g = curr.g + 1
			n.f = h(n.val, goal) + n.g
			n.val.prv = &curr.val
			openNodes = add(openNodes, n)
		}
	}
	return nil
}

func (item *PQItem) GenStates() []PQItem {
	states := item.val.GenStates()
	items := make([]PQItem, 0, len(states))
	for _, s := range states {
		items = append(items, PQItem{val: s, f: item.f, g: item.g})
	}
	return items
}

func add(slice PQItemSlice, item PQItem) PQItemSlice {
	if len(slice) == 0 {
		return append(slice, item)
	}
	var idx int
	var v PQItem
	for idx, v = range slice {
		if item.f < v.f {
			break
		}
		idx += 1
	}
	res := make(PQItemSlice, len(slice))
	copy(res, slice)
	res = res[:idx]
	res = append(res, item)
	res = append(res, slice[idx:]...)
	return res
}

func get(items PQItemSlice, item PQItem) *PQItem {
	for _, i := range items {
		if item.val.Equals(i.val) {
			return &i
		}
	}
	return nil
}

func remove(items PQItemSlice, item PQItem) PQItemSlice {
	for i, v := range items {
		if item.val.Equals(v.val) {
			return append(items[:i], items[i+1:]...)
		}
	}
	return items
}

func FirstHeuristic(start, goal State) int {
	var value float32
	for row := 0; row < len(start.Content); row++ {
		for col := 0; col < len(start.Content[row]); col++ {
			if start.Content[row][col] != goal.Content[row][col] {
				value += 1
			}
		}
	}

	value /= 4
	if value < 1 {
		return 1
	}
	return int(value)
}

func SecondHeuristic(start, goal State) int {
	var value float32

	for row := 0; row < len(start.Content); row++ {
		for col := 0; col < len(start.Content[row]); col++ {
			color := start.Content[row][col]

			if start.Content[row][col] != goal.Content[row][col] {
				for targetRow := 0; targetRow < len(goal.Content); targetRow++ {
					for targetCol := 0; targetCol < len(goal.Content); targetCol++ {
						if color == goal.Content[targetRow][targetCol] {
							diff1 := row - targetRow
							diff2 := col - targetCol
							if diff1 < 0 {
								diff1 *= -1
							}
							if diff2 < 0 {
								diff2 *= -1
							}
							distance := diff1 + diff2
							if distance <= 2 {
								value += float32(distance)
							} else {
								value += 1
							}
							break
						}
					}
				}
			}
		}
	}

	value /= 4
	if value < 1 {
		return 1
	}
	return int(value)
}

func secondHeuristicForSubtask(start, goal State, colorToExclude byte) int {
	var value float32

	for row := 0; row < len(start.Content); row++ {
		for col := 0; col < len(start.Content[row]); col++ {
			color := start.Content[row][col]
			if color == colorToExclude {
				continue
			}

			if start.Content[row][col] != goal.Content[row][col] {
				for targetRow := 0; targetRow < len(goal.Content); targetRow++ {
					for targetCol := 0; targetCol < len(goal.Content); targetCol++ {
						if color == goal.Content[targetRow][targetCol] {
							diff1 := row - targetRow
							diff2 := col - targetCol
							if diff1 < 0 {
								diff1 *= -1
							}
							if diff2 < 0 {
								diff2 *= -1
							}
							distance := diff1 + diff2
							if distance <= 2 {
								value += float32(distance)
							} else {
								value += 1
							}
							break
						}
					}
				}
			}
		}
	}

	value /= 4
	if value < 1 {
		return 1
	}
	return int(value)
}

func SubtaskHeuristicWithoutSecond(start, goal State) int {
	return storage.get(start)
}

func SubtaskHeuristic(start, goal State) int {
	return secondHeuristicForSubtask(start, goal, 0) + storage.get(start)
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
