// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example demonstrates a priority queue built using the heap interface.

// https://golang.org/src/container/heap/example_pq_test.go

package server

import "container/heap"

var queue PriorityQueue = make(PriorityQueue, 0)

func InitPriorityQueue() {
	// Create a priority queue and establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
}

func PushRequestToActionQueue(request interface{}, requestType string, priority int) {
	heap.Push(&queue, &Action{
		priority:    priority,
		request:     request,
		requestType: requestType,
	})
}

// An Action is something we manage in a priority queue.
type Action struct {
	requestType string
	request     interface{}
	priority    int // The priority of the action in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the action in the heap.
}

// A PriorityQueue implements heap.Interface and holds Actions.
type PriorityQueue []*Action

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Empty() bool { return len(pq) == 0 }

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
	action := x.(*Action)
	action.index = n
	*pq = append(*pq, action)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	action := old[n-1]
	action.index = -1 // for safety
	*pq = old[0 : n-1]
	return action
}

// update modifies the priority and value of an Action in the queue.
func (pq *PriorityQueue) update(action *Action, requestType string, request interface{}, priority int) {
	action.requestType = requestType
	action.request = request
	action.priority = priority
	heap.Fix(pq, action.index)
}
