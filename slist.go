package main

import (
	"fmt"
)

type SLNode struct {
	data string
	next *SLNode
}

type SList struct {
	head *SLNode
	size int
}

func NewSList() *SList {
	return &SList{
		head: nil,
		size: 0,
	}
}

func (l *SList) SLCLEAR() {
	l.head = nil
	l.size = 0
	fmt.Println("Список очищен.")
}

func (l *SList) SLPUSH_HEAD(value string) {
	newNode := &SLNode{
		data: value,
		next: l.head,
	}
	l.head = newNode
	l.size++
}

func (l *SList) SLPUSH_TAIL(value string) {
	newNode := &SLNode{
		data: value,
		next: nil,
	}

	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	l.size++
}

func (l *SList) SLPUSH_BEFORE(target, value string) {
	if l.head == nil {
		fmt.Println("Список пуст!")
		return
	}

	if l.head.data == target {
		l.SLPUSH_HEAD(value)
		return
	}

	current := l.head
	for current.next != nil && current.next.data != target {
		current = current.next
	}

	if current.next != nil {
		newNode := &SLNode{
			data: value,
			next: current.next,
		}
		current.next = newNode
		l.size++
	}
}

func (l *SList) SLPUSH_AFTER(target, value string) {
	current := l.head
	for current != nil && current.data != target {
		current = current.next
	}

	if current != nil {
		newNode := &SLNode{
			data: value,
			next: current.next,
		}
		current.next = newNode
		l.size++
	}
}

func (l *SList) SLREMOVE_HEAD() {
	if l.head == nil {
		fmt.Println("Список пуст!")
		return
	}

	l.head = l.head.next
	l.size--
}

func (l *SList) SLREMOVE_TAIL() {
	if l.head == nil {
		fmt.Println("Список пуст!")
		return
	}

	if l.head.next == nil {
		l.head = nil
	} else {
		current := l.head
		for current.next.next != nil {
			current = current.next
		}
		current.next = nil
	}
	l.size--
}

func (l *SList) SLREMOVE_BEFORE(target string) {
	if l.head == nil || l.head.next == nil {
		fmt.Println("Недостаточно элементов для удаления!")
		return
	}

	if l.head.data == target {
		fmt.Println("Невозможно удалить элемент перед головой!")
		return
	}

	if l.head.next.data == target {
		l.SLREMOVE_HEAD()
		return
	}

	current := l.head
	for current.next.next != nil && current.next.next.data != target {
		current = current.next
	}

	if current.next.next != nil {
		current.next = current.next.next
		l.size--
	}
}

func (l *SList) SLREMOVE_AFTER(target string) {
	current := l.head
	for current != nil && current.data != target {
		current = current.next
	}

	if current != nil && current.next != nil {
		current.next = current.next.next
		l.size--
	}
}

func (l *SList) SLREMOVE_VALUE(value string) {
	if l.head == nil {
		fmt.Println("Список пуст!")
		return
	}

	if l.head.data == value {
		l.SLREMOVE_HEAD()
		return
	}

	current := l.head
	for current.next != nil && current.next.data != value {
		current = current.next
	}

	if current.next != nil {
		current.next = current.next.next
		l.size--
	}
}

func (l *SList) SLSEARCH(value string) bool {
	current := l.head
	index := 0
	for current != nil {
		if current.data == value {
			return true
		}
		current = current.next
		index++
	}
	return false
}

func (l *SList) SLGET(index int) (string, error) {
	if index >= l.size || index < 0 {
		return "", fmt.Errorf("индекс %d вне диапазона [0, %d]", index, l.size-1)
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.data, nil
}

func (l *SList) SLPRINT_FORWARD() {
	if l.head == nil {
		fmt.Println("Список пуст!")
		return
	}

	fmt.Printf("Список [%d] (прямой порядок): ", l.size)
	current := l.head
	for current != nil {
		fmt.Printf("\"%s\"", current.data)
		if current.next != nil {
			fmt.Print(" -> ")
		}
		current = current.next
	}
	fmt.Println(" -> NULL")
}

func (l *SList) printReverse(node *SLNode) {
	if node == nil {
		return
	}
	l.printReverse(node.next)
	fmt.Printf("\"%s\"", node.data)
	if node != l.head {
		fmt.Print(" <- ")
	}
}

func (l *SList) SLPRINT_BACKWARD() {
	if l.head == nil {
		fmt.Println("Список пуст!")
		return
	}

	fmt.Printf("Список [%d] (обратный порядок): NULL <- ", l.size)
	l.printReverse(l.head)
	fmt.Println()
}

func (l *SList) SLLENGTH() int {
	return l.size
}

// Для сериализации
func (l *SList) GetAll() []string {
	result := make([]string, 0, l.size)
	current := l.head
	for current != nil {
		result = append(result, current.data)
		current = current.next
	}
	return result
}

func (l *SList) SetAll(items []string) {
	l.SLCLEAR()
	for i := len(items) - 1; i >= 0; i-- {
		l.SLPUSH_HEAD(items[i])
	}
}