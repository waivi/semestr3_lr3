package main

import (
	"fmt"
)

type DNode struct {
	data string
	next *DNode
	prev *DNode
}

type DList struct {
	head *DNode
	tail *DNode
	size int
}

func NewDList() *DList {
	return &DList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (l *DList) DLCLEAR() {
	l.head = nil
	l.tail = nil
	l.size = 0
	fmt.Println("Двусвязный список очищен.")
}

func (l *DList) DLPUSH_HEAD(value string) {
	newNode := &DNode{
		data: value,
		next: l.head,
		prev: nil,
	}

	if l.head != nil {
		l.head.prev = newNode
	} else {
		l.tail = newNode
	}

	l.head = newNode
	l.size++
}

func (l *DList) DLPUSH_TAIL(value string) {
	newNode := &DNode{
		data: value,
		next: nil,
		prev: l.tail,
	}

	if l.tail != nil {
		l.tail.next = newNode
	} else {
		l.head = newNode
	}

	l.tail = newNode
	l.size++
}

func (l *DList) DLPUSH_BEFORE(target, value string) {
	if l.head == nil {
		fmt.Println("Список пуст!")
		return
	}

	current := l.head
	for current != nil && current.data != target {
		current = current.next
	}

	if current == nil {
		return
	}

	if current == l.head {
		l.DLPUSH_HEAD(value)
		return
	}

	newNode := &DNode{
		data: value,
		next: current,
		prev: current.prev,
	}
	current.prev.next = newNode
	current.prev = newNode
	l.size++
}

func (l *DList) DLPUSH_AFTER(target, value string) {
	if l.head == nil {
		fmt.Println("Список пуст!")
		return
	}

	current := l.head
	for current != nil && current.data != target {
		current = current.next
	}

	if current == nil {
		return
	}

	if current == l.tail {
		l.DLPUSH_TAIL(value)
		return
	}

	newNode := &DNode{
		data: value,
		next: current.next,
		prev: current,
	}
	current.next.prev = newNode
	current.next = newNode
	l.size++
}

func (l *DList) DLREMOVE_HEAD() {
	if l.head == nil {
		fmt.Println("Список пуст!")
		return
	}

	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
		l.head.prev = nil
	}
	l.size--
}

func (l *DList) DLREMOVE_TAIL() {
	if l.tail == nil {
		fmt.Println("Список пуст!")
		return
	}

	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = l.tail.prev
		l.tail.next = nil
	}
	l.size--
}

func (l *DList) DLREMOVE_BEFORE(target string) {
	if l.head == nil || l.head.next == nil {
		fmt.Println("Недостаточно элементов для удаления!")
		return
	}

	current := l.head
	for current != nil && current.data != target {
		current = current.next
	}

	if current == nil || current.prev == nil {
		return
	}

	if current.prev == l.head {
		l.DLREMOVE_HEAD()
		return
	}

	nodeToRemove := current.prev
	nodeToRemove.prev.next = current
	current.prev = nodeToRemove.prev
	l.size--
}

func (l *DList) DLREMOVE_AFTER(target string) {
	if l.head == nil || l.head.next == nil {
		fmt.Println("Недостаточно элементов для удаления!")
		return
	}

	current := l.head
	for current != nil && current.data != target {
		current = current.next
	}

	if current == nil || current.next == nil {
		return
	}

	if current.next == l.tail {
		l.DLREMOVE_TAIL()
		return
	}

	nodeToRemove := current.next
	current.next = nodeToRemove.next
	nodeToRemove.next.prev = current
	l.size--
}

func (l *DList) DLREMOVE_VALUE(value string) {
	if l.head == nil {
		fmt.Println("Список пуст!")
		return
	}

	current := l.head
	for current != nil && current.data != value {
		current = current.next
	}

	if current == nil {
		return
	}

	if current == l.head {
		l.DLREMOVE_HEAD()
	} else if current == l.tail {
		l.DLREMOVE_TAIL()
	} else {
		current.prev.next = current.next
		current.next.prev = current.prev
		l.size--
	}
}

func (l *DList) DLSEARCH(value string) bool {
	current := l.head
	for current != nil {
		if current.data == value {
			return true
		}
		current = current.next
	}
	return false
}

func (l *DList) DLGET(index int) (string, error) {
	if index >= l.size || index < 0 {
		return "", fmt.Errorf("индекс %d вне диапазона [0, %d]", index, l.size-1)
	}

	var current *DNode
	// Оптимизация: идем с начала или с конца в зависимости от индекса
	if index < l.size/2 {
		current = l.head
		for i := 0; i < index; i++ {
			current = current.next
		}
	} else {
		current = l.tail
		for i := l.size - 1; i > index; i-- {
			current = current.prev
		}
	}
	return current.data, nil
}

func (l *DList) DLPRINT_FORWARD() {
	if l.head == nil {
		fmt.Println("Список пуст!")
		return
	}

	fmt.Printf("Двусвязный список [%d] (прямой порядок): NULL <- ", l.size)
	current := l.head
	for current != nil {
		fmt.Printf("\"%s\"", current.data)
		if current.next != nil {
			fmt.Print(" <-> ")
		}
		current = current.next
	}
	fmt.Println(" -> NULL")
}

func (l *DList) DLPRINT_BACKWARD() {
	if l.tail == nil {
		fmt.Println("Список пуст!")
		return
	}

	fmt.Printf("Двусвязный список [%d] (обратный порядок): NULL <- ", l.size)
	current := l.tail
	for current != nil {
		fmt.Printf("\"%s\"", current.data)
		if current.prev != nil {
			fmt.Print(" <-> ")
		}
		current = current.prev
	}
	fmt.Println(" -> NULL")
}

func (l *DList) DLLENGTH() int {
	return l.size
}

// Для сериализации
func (l *DList) GetAll() []string {
	result := make([]string, 0, l.size)
	current := l.head
	for current != nil {
		result = append(result, current.data)
		current = current.next
	}
	return result
}

func (l *DList) SetAll(items []string) {
	l.DLCLEAR()
	for _, item := range items {
		l.DLPUSH_TAIL(item)
	}
}