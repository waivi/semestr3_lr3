package main

import (
	"fmt"
	"strings"
)

type Node struct {
	keyValue [2]int // [0] = key, [1] = value
	next     *Node
}

func NewNode(k, v int) *Node {
	return &Node{
		keyValue: [2]int{k, v},
		next:     nil,
	}
}

type ChainingHashTable struct {
	capacity int
	size     int
	table    []*Node
}

func NewChainingHashTable(capacity int) *ChainingHashTable {
	if capacity <= 0 {
		capacity = 10
	}
	return &ChainingHashTable{
		capacity: capacity,
		size:     0,
		table:    make([]*Node, capacity),
	}
}

func (ht *ChainingHashTable) hash(key int) int {
	return key % ht.capacity
}

func (ht *ChainingHashTable) Add(key, value int) {
	// Проверка на дубликаты
	if ht.Contains(key).first {
		return
	}

	h := ht.hash(key)
	newNode := NewNode(key, value)

	// Вставка в цепочку
	if ht.table[h] == nil {
		ht.table[h] = newNode
	} else {
		current := ht.table[h]
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	ht.size++
}

func (ht *ChainingHashTable) Remove(key int) {
	h := ht.hash(key)
	current := ht.table[h]
	var prev *Node

	for current != nil {
		if current.keyValue[0] == key {
			if prev == nil {
				ht.table[h] = current.next
			} else {
				prev.next = current.next
			}
			ht.size--
			return
		}
		prev = current
		current = current.next
	}
}

type ContainsResult struct {
	first  bool
	second int
}

func (ht *ChainingHashTable) Contains(key int) ContainsResult {
	h := ht.hash(key)
	current := ht.table[h]

	for current != nil {
		if current.keyValue[0] == key {
			return ContainsResult{true, current.keyValue[1]}
		}
		current = current.next
	}
	return ContainsResult{false, -1}
}

func (ht *ChainingHashTable) ToString() string {
	var sb strings.Builder

	for i := 0; i < ht.capacity; i++ {
		sb.WriteString(fmt.Sprintf("[%3d]: ", i))
		current := ht.table[i]

		if current == nil {
			sb.WriteString("empty")
		} else {
			for current != nil {
				sb.WriteString(fmt.Sprintf("(%d,%d)", current.keyValue[0], current.keyValue[1]))
				if current.next != nil {
					sb.WriteString(" -> ")
				}
				current = current.next
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func (ht *ChainingHashTable) GetAllElements() [][2]int {
	elements := make([][2]int, 0, ht.size)
	for i := 0; i < ht.capacity; i++ {
		current := ht.table[i]
		for current != nil {
			elements = append(elements, current.keyValue)
			current = current.next
		}
	}
	return elements
}

func (ht *ChainingHashTable) GetChainLengths() (minLength, maxLength int, avgLength float64) {
	minLength = int(^uint(0) >> 1) // Max int
	maxLength = 0
	totalLength := 0
	nonEmptyChains := 0

	for i := 0; i < ht.capacity; i++ {
		length := 0
		current := ht.table[i]

		for current != nil {
			length++
			current = current.next
		}

		if length > 0 {
			if length < minLength {
				minLength = length
			}
			if length > maxLength {
				maxLength = length
			}
			totalLength += length
			nonEmptyChains++
		}
	}

	// Обработка случая когда все цепочки пустые
	if nonEmptyChains == 0 {
		minLength = 0
		maxLength = 0
		avgLength = 0.0
	} else {
		avgLength = float64(totalLength) / float64(nonEmptyChains)
	}
	return
}

func (ht *ChainingHashTable) GetSize() int {
	return ht.size
}

func (ht *ChainingHashTable) GetCapacity() int {
	return ht.capacity
}

func (ht *ChainingHashTable) GetLoadFactor() float64 {
	if ht.capacity > 0 {
		return float64(ht.size) / float64(ht.capacity)
	}
	return 0.0
}

func (ht *ChainingHashTable) Clear() {
	for i := 0; i < ht.capacity; i++ {
		ht.table[i] = nil
	}
	ht.size = 0
}

func (ht *ChainingHashTable) IsEmpty() bool {
	return ht.size == 0
}
