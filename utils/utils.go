package utils

import (
	"bufio"
	"errors"
	"math"
	"os"
	"strings"
	"sync"
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

// BoolToInt converts a boolean to an integer for easy hacks
func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// Queue implements an asynchronous queue structure
type Queue struct {
	Items []interface{}
	mux   sync.Mutex
}

// Pop removes the first item from the queue if available
func (q *Queue) Pop(timeout time.Duration) (interface{}, error) {
	ch := make(chan interface{}, 1)
	quit := make(chan bool, 1)
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				q.mux.Lock()
				if len(q.Items) != 0 {
					popped := q.Items[0]
					q.Items = q.Items[1:]
					ch <- popped
					q.mux.Unlock()
					return
				}
				q.mux.Unlock()
			}
		}
	}()
	select {
	case res := <-ch:
		return res, nil
	case <-time.After(timeout):
		quit <- true
		return nil, errors.New("couldn't pop Queue")
	}
}

// Push adds an item into a Queue
func (q *Queue) Push(item interface{}) {
	q.mux.Lock()
	defer q.mux.Unlock()
	q.Items = append(q.Items, item)
}

// AbsInt gives the absolute value of an integer
func AbsInt(i int) int {
	return int(math.Abs(float64(i)))
}

// SplitByComma returns an array of strings split by commas
func SplitByComma(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool { return r == ',' })
}
