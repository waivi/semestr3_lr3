package main

import (
	"fmt"
)

type SNode struct {
	data string
	next *SNode
}

type Stack struct {
	top  *SNode
	size int
}

func NewStack() *Stack {
	return &Stack{
		top:  nil,
		size: 0,
	}
}

func (s *Stack) SCLEAR() {
	s.top = nil
	s.size = 0
	fmt.Println("Стек очищен.")
}

func (s *Stack) SPUSH(value string) {
	newNode := &SNode{
		data: value,
		next: s.top,
	}
	s.top = newNode
	s.size++
}

func (s *Stack) SPOP() (string, error) {
	if s.top == nil {
		return "", fmt.Errorf("ошибка: стек пуст")
	}

	value := s.top.data
	s.top = s.top.next
	s.size--
	return value, nil
}

func (s *Stack) STOP() (string, error) {
	if s.top == nil {
		return "", fmt.Errorf("ошибка: стек пуст")
	}
	return s.top.data, nil
}

func (s *Stack) SPRINT() {
	if s.top == nil {
		fmt.Println("Стек пуст!")
		return
	}

	fmt.Printf("Стек [%d] (вершина -> основание): ", s.size)
	current := s.top
	for current != nil {
		fmt.Printf("\"%s\"", current.data)
		if current.next != nil {
			fmt.Print(" -> ")
		}
		current = current.next
	}
	fmt.Println(" -> NULL")
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Size() int {
	return s.size
}

// Для сериализации
func (s *Stack) GetAll() []string {
	result := make([]string, 0, s.size)
	current := s.top
	for current != nil {
		result = append(result, current.data)
		current = current.next
	}
	return result
}

func (s *Stack) SetAll(items []string) {
	s.SCLEAR()
	for i := len(items) - 1; i >= 0; i-- {
		s.SPUSH(items[i])
	}
}