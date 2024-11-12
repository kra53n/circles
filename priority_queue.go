// priority queue implementation for AStartSearch based on "container/heap"
package main

type PQItem struct {
	val    State
	f int
	g int
	idx int
}

func (item *PQItem) GenStates() []PQItem {
	states := item.val.GenStates()
	items := make([]PQItem, 0, len(states))
	for _, s := range states {
		items = append(items, PQItem{val: s, f: item.f, g: item.g, idx: item.idx})
	}
	return items
}

type PriorityQueue []*PQItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].f > pq[j].f
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].idx = i
	pq[j].idx = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*PQItem)
	item.idx = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.idx = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) GetItem(item PQItem) *PQItem {
	for _, i := range *pq {
		if item.val.Equals(i.val) {
			return i
		}
	}
	return nil
}
