import (
	"container/heap"
	"sort"
)

// https://golang.org/pkg/container/heap/
type MinHeap []int

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *MinHeap) Pop() interface{} {
	n := len(*m)
	x := (*m)[n-1]
	*m = (*m)[0 : n-1]
	return x
}

func maxEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	events = append(events, []int{100001, 100001})
	pq := &MinHeap{}
	heap.Init(pq)
	max_attends, day := 0, 1
	for _, e := range events {
		for pq.Len() > 0 && day < e[0] {
			top := (*pq)[0]
			heap.Pop(pq)
			if top >= day {
				max_attends++
				day++
			}
		}
		if pq.Len() == 0 {
			day = e[0]
		}
		heap.Push(pq, e[1])
	}
	return max_attends
}