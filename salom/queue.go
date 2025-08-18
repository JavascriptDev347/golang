package main

import "fmt"

// new Queue create
type Queue struct {
	data []int
}

// newQueue yangi Queue obektini yaratib beradi
func newQueue() *Queue {
	return &Queue{
		data: []int{},
	}
}

// isEmpty queueni bo'sh yoki bo'sh emasligini qaytaradi
func (q *Queue) isEmpty() bool {
	return len(q.data) == 0
}

// Peek queuedagi keyingi elementni aniqlab berad
func (q *Queue) Peek() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Queue is empty")
		
	}
	return q.data[0], nil
}

// Enqueue elementni Queuega qo'shib beradi
func (q *Queue) Enqueue(value int) *Queue {
	q.data = append(q.data, value)
	return q
}

// Dequeue elementni Queuedan olib tashlab beradi
func (q *Queue) Dequeue() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Queue is empty")
	}
	value := q.data[0]
	q.data = q.data[1:]
	return value, nil
}
