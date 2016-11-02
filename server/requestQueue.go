// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example demonstrates a priority queue built using the heap interface.

// https://golang.org/src/container/heap/example_pq_test.go

package server

import "container/heap"

// A PriorityQueue implements heap.Interface and holds Actions.
type PriorityQueue []*Action

var Queue PriorityQueue = make(PriorityQueue, 0)

// An Action is something we manage in a priority queue.
type Action struct {
	RequestType string
	Request     interface{}
	priority    int // The priority of the action in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the action in the heap.
}

func PushRequestToActionQueue(request interface{}, requestType string, reqPriority int) {
	heap.Push(&Queue, &Action{
		Request:     request,
		priority:    reqPriority,
		RequestType: requestType,
	})
}

func InitPriorityQueue() {
	// Create a priority queue and establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
}

func PopFromRequestQueue() *Action {
	if Queue.Len() != 0 {
		return heap.Pop(&Queue).(*Action)
	} else {
		return nil
	}
}

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
	item := x.(*Action)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Action in the queue.
func (pq *PriorityQueue) Update(action *Action, requestType string, request interface{}, priority int) {
	action.RequestType = requestType
	action.Request = request
	action.priority = priority
	heap.Fix(pq, action.index)
}
