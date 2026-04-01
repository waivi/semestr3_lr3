package main

import (
	"fmt"
)

type QNode struct {
	data string
	next *QNode
}

type Queue struct {
	front *QNode
	rear  *QNode
	size  int
}

func NewQueue() *Queue {
	return &Queue{
		front: nil,
		rear:  nil,
		size:  0,
	}
}

func (q *Queue) QCLEAR() {
	q.front = nil
	q.rear = nil
	q.size = 0
	fmt.Println("Очередь очищена.")
}

func (q *Queue) QPUSH(value string) {
	newNode := &QNode{
		data: value,
		next: nil,
	}

	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
	q.size++
}

func (q *Queue) QPOP() (string, error) {
	if q.front == nil {
		return "", fmt.Errorf("ошибка: очередь пуста")
	}

	value := q.front.data
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}
	q.size--
	return value, nil
}

func (q *Queue) QFRONT() (string, error) {
	if q.front == nil {
		return "", fmt.Errorf("ошибка: очередь пуста")
	}
	return q.front.data, nil
}

func (q *Queue) QPRINT() {
	if q.front == nil {
		fmt.Println("Очередь пуста!")
		return
	}

	fmt.Printf("Очередь [%d] (начало -> конец): ", q.size)
	current := q.front
	for current != nil {
		fmt.Printf("\"%s\"", current.data)
		if current.next != nil {
			fmt.Print(" -> ")
		}
		current = current.next
	}
	fmt.Println(" -> NULL")
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Size() int {
	return q.size
}

// Для сериализации
func (q *Queue) GetAll() []string {
	result := make([]string, 0, q.size)
	current := q.front
	for current != nil {
		result = append(result, current.data)
		current = current.next
	}
	return result
}

func (q *Queue) SetAll(items []string) {
	q.QCLEAR()
	for _, item := range items {
		q.QPUSH(item)
	}
}