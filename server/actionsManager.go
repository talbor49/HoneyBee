package server

import (
	"fmt"
	"time"
)

func priorityQueueWorker() {
	// fmt.Println("Entered queue worker")
	// defer fmt.Println("Quit queue worker")
	for {
		if queue.Empty() {
			// fmt.Println("queue is empty :(")
			time.Sleep(1 * time.Millisecond)
		} else {
			queue.Pop()
			fmt.Println("Poped!")
		}
	}
}
