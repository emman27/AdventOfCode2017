package utils

import (
	"bufio"
	"errors"
	"os"
	"time"
)

func ReadFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// LinkedList defintes a LinkedList
type LinkedList interface{}

// Queue implements an asynchronous queue structure
type Queue struct {
	LinkedList
	Items []interface{}
}

// Pop removes the first item from the queue if available
func (q *Queue) Pop(timeout time.Duration) (interface{}, error) {
	ch := make(chan interface{})
	go func() {
		for true {
			if len(q.Items) != 0 {
				popped := q.Items[0]
				q.Items = q.Items[1:]
				ch <- popped
			}
		}
	}()
	select {
	case res := <-ch:
		return res, nil
	case <-time.After(timeout * time.Second):
		return nil, errors.New("couldn't pop Queue")
	}
}

// Push adds an item into a Queue
func (q *Queue) Push(item interface{}) {
	q.Items = append(q.Items, item)
}
