// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example demonstrates a priority queue built using the heap interface.
package server

import (
	"container/heap"
	"fmt"
)

// An Pair is something we manage in a priority queue.
type Pair struct {
	requestType string
	request     interface{}
	priority    int // The priority of the pair in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the pair in the heap.
}

// A PriorityQueue implements heap.Interface and holds Pairs.
type PriorityQueue []*Pair

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	pair := x.(*Pair)
	pair.index = n
	*pq = append(*pq, pair)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	pair := old[n-1]
	pair.index = -1 // for safety
	*pq = old[0 : n-1]
	return pair
}

// update modifies the priority and value of an Pair in the queue.
func (pq *PriorityQueue) update(pair *Pair, requestType string, request interface{}, priority int) {
	pair.requestType = requestType
	pair.request = request
	pair.priority = priority
	heap.Fix(pq, pair.index)
}

// This example creates a PriorityQueue with some pairs, adds and manipulates an pair,
// and then removes the pairs in priority order.
func Example_priorityQueue() {
	// Some pairs and their priorities.

	// Create a priority queue, put the pairs in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// Test priority queue
	// heap.Push(&pq, &Pair{
	// 	requestType: "GET",
	// 	priority:    6,
	// 	index:       0,
	// })
	//
	// heap.Push(&pq, &Pair{
	// 	requestType: "DELETE",
	// 	priority:    4,
	// 	index:       0,
	// })
	//
	// heap.Push(&pq, &Pair{
	// 	requestType: "SET",
	// 	priority:    8,
	// 	index:       0,
	// })

	// Way to insert pairs into queue in a more efficient way.
	// i := 0
	// for value, priority := range pairs {
	// 	pq[i] = &Pair{
	//
	// 		value:    value,
	// 		priority: priority,
	// 		index:    i,
	// 	}
	// 	i++
	// }

	// Take the pairs out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		pair := heap.Pop(&pq).(*Pair)
		fmt.Printf("%.2d:%s ", pair.priority, pair.requestType)
	}
}
