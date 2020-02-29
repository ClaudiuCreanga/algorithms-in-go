package algorithms_in_go

import (
	"container/heap"
)

type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

type PriorityQueue []*Item

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	pq := make(PriorityQueue, len(flights))
	i := 0
	for value, priority := range flights {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)
}
