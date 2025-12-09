package main

import (
	"fmt"
)

type MArray struct {
	data []string
	size int
}

func NewMArray() *MArray {
	return &MArray{
		data: make([]string, 0),
		size: 0,
	}
}

func (a *MArray) MCLEAR() {
	a.data = make([]string, 0)
	a.size = 0
	fmt.Println("Массив очищен.")
}

func (a *MArray) mresize() {
	// В Go слайсы автоматически расширяются
	// Эта функция для совместимости с C++ API
	if a.size >= len(a.data) {
		newCapacity := len(a.data) * 2
		if newCapacity == 0 {
			newCapacity = 4
		}
		newData := make([]string, newCapacity)
		// Копируем только существующие элементы (первые a.size)
		copy(newData, a.data[:a.size])
		a.data = newData
	}
}

func (a *MArray) MADDINDEX(index int, value string) {
	if index > a.size {
		fmt.Printf("Ошибка: индекс %d превышает размер массива (%d)\n", index, a.size)
		return
	}

	a.mresize()

	// Если добавляем в конец
	if index == a.size {
		a.data = append(a.data[:a.size], value)
	} else {
		// Сдвигаем элементы вправо
		a.data = append(a.data[:index+1], a.data[index:]...)
		a.data[index] = value
	}
	a.size++
}

func (a *MArray) MADDEND(value string) {
	a.mresize()
	// Добавляем элемент в конец
	if a.size < len(a.data) {
		a.data[a.size] = value
	} else {
		a.data = append(a.data, value)
	}
	a.size++
}

func (a *MArray) MGETINDEX(index int) (string, error) {
	if index >= a.size || index < 0 {
		return "", fmt.Errorf("индекс %d вне диапазона [0, %d]", index, a.size-1)
	}
	return a.data[index], nil
}

func (a *MArray) MREMOVEINDEX(index int) {
	if index >= a.size || index < 0 {
		fmt.Printf("Ошибка: индекс %d вне диапазона [0, %d]\n", index, a.size-1)
		return
	}

	// Сдвигаем элементы влево
	copy(a.data[index:], a.data[index+1:a.size])
	a.size--
}

func (a *MArray) MREPLACEINDEX(index int, newValue string) {
	if index >= a.size || index < 0 {
		fmt.Printf("Ошибка: индекс %d вне диапазона [0, %d]\n", index, a.size-1)
		return
	}

	oldValue := a.data[index]
	a.data[index] = newValue
	fmt.Printf("Элемент \"%s\" заменён на \"%s\" на позиции %d\n", oldValue, newValue, index)
}

func (a *MArray) MLENGTH() int {
	return a.size
}

func (a *MArray) MPRINT() {
	if a.size == 0 {
		fmt.Println("Массив пуст!")
		return
	}

	fmt.Printf("Динамический массив [%d]: ", a.size)
	for i := 0; i < a.size; i++ {
		fmt.Printf("\"%s\"", a.data[i])
		if i < a.size-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}

// Для сериализации
func (a *MArray) GetAll() []string {
	// Возвращаем копию данных
	result := make([]string, a.size)
	copy(result, a.data[:a.size])
	return result
}

func (a *MArray) SetAll(items []string) {
	a.data = make([]string, len(items))
	copy(a.data, items)
	a.size = len(items)
}
