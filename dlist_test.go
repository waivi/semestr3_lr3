package main

import (
	"encoding/binary"
	"testing"
	"path/filepath" 
	"os"
	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDList(t *testing.T) {
	t.Run("NewDList", func(t *testing.T) {
		list := NewDList()
		assert.NotNil(t, list)
		assert.Equal(t, 0, list.size)
		assert.Nil(t, list.head)
		assert.Nil(t, list.tail)
	})

	t.Run("DLPUSH_HEAD", func(t *testing.T) {
		list := NewDList()

		list.DLPUSH_HEAD("first")
		assert.Equal(t, 1, list.size)
		assert.Equal(t, "first", list.head.data)
		assert.Equal(t, "first", list.tail.data)
		assert.Equal(t, list.head, list.tail)

		list.DLPUSH_HEAD("second")
		assert.Equal(t, 2, list.size)
		assert.Equal(t, "second", list.head.data)
		assert.Equal(t, "first", list.tail.data)
		assert.Equal(t, list.head.next, list.tail)
		assert.Equal(t, list.tail.prev, list.head)
	})

	t.Run("DLPUSH_TAIL", func(t *testing.T) {
		list := NewDList()

		list.DLPUSH_TAIL("first")
		assert.Equal(t, 1, list.size)
		assert.Equal(t, "first", list.head.data)
		assert.Equal(t, "first", list.tail.data)

		list.DLPUSH_TAIL("second")
		assert.Equal(t, 2, list.size)
		assert.Equal(t, "first", list.head.data)
		assert.Equal(t, "second", list.tail.data)
		assert.Equal(t, list.head.next, list.tail)
		assert.Equal(t, list.tail.prev, list.head)
	})

	t.Run("DLSEARCH", func(t *testing.T) {
		list := NewDList()
		list.DLPUSH_TAIL("A")
		list.DLPUSH_TAIL("B")
		list.DLPUSH_TAIL("C")

		assert.True(t, list.DLSEARCH("A"))
		assert.True(t, list.DLSEARCH("B"))
		assert.True(t, list.DLSEARCH("C"))
		assert.False(t, list.DLSEARCH("D"))
	})

	t.Run("DLGET", func(t *testing.T) {
		list := NewDList()
		list.DLPUSH_TAIL("A")
		list.DLPUSH_TAIL("B")
		list.DLPUSH_TAIL("C")
		list.DLPUSH_TAIL("D")

		// Проверяем доступ с начала (индексы 0-1)
		val, err := list.DLGET(0)
		require.NoError(t, err)
		assert.Equal(t, "A", val)

		val, err = list.DLGET(1)
		require.NoError(t, err)
		assert.Equal(t, "B", val)

		// Проверяем доступ с конца (индексы 2-3)
		val, err = list.DLGET(2)
		require.NoError(t, err)
		assert.Equal(t, "C", val)

		val, err = list.DLGET(3)
		require.NoError(t, err)
		assert.Equal(t, "D", val)

		// Проверяем ошибки
		_, err = list.DLGET(4)
		assert.Error(t, err)

		_, err = list.DLGET(-1)
		assert.Error(t, err)
	})

	t.Run("DLREMOVE_HEAD", func(t *testing.T) {
		list := NewDList()
		list.DLPUSH_TAIL("A")
		list.DLPUSH_TAIL("B")
		list.DLPUSH_TAIL("C")

		list.DLREMOVE_HEAD()
		assert.Equal(t, 2, list.size)
		assert.Equal(t, "B", list.head.data)
		assert.Nil(t, list.head.prev)

		list.DLREMOVE_HEAD()
		assert.Equal(t, 1, list.size)
		assert.Equal(t, "C", list.head.data)
		assert.Equal(t, list.head, list.tail)

		list.DLREMOVE_HEAD()
		assert.Equal(t, 0, list.size)
		assert.Nil(t, list.head)
		assert.Nil(t, list.tail)
	})

	t.Run("DLREMOVE_TAIL", func(t *testing.T) {
		list := NewDList()
		list.DLPUSH_TAIL("A")
		list.DLPUSH_TAIL("B")
		list.DLPUSH_TAIL("C")

		list.DLREMOVE_TAIL()
		assert.Equal(t, 2, list.size)
		assert.Equal(t, "B", list.tail.data)
		assert.Nil(t, list.tail.next)

		list.DLREMOVE_TAIL()
		assert.Equal(t, 1, list.size)
		assert.Equal(t, "A", list.tail.data)
		assert.Equal(t, list.head, list.tail)

		list.DLREMOVE_TAIL()
		assert.Equal(t, 0, list.size)
		assert.Nil(t, list.head)
		assert.Nil(t, list.tail)
	})

	t.Run("DLREMOVE_VALUE", func(t *testing.T) {
		list := NewDList()
		list.DLPUSH_TAIL("A")
		list.DLPUSH_TAIL("B")
		list.DLPUSH_TAIL("C")

		// Удаляем из середины
		list.DLREMOVE_VALUE("B")
		assert.Equal(t, 2, list.size)
		assert.Equal(t, "A", list.head.data)
		assert.Equal(t, "C", list.tail.data)
		assert.Equal(t, list.head.next, list.tail)
		assert.Equal(t, list.tail.prev, list.head)

		// Удаляем голову
		list.DLREMOVE_VALUE("A")
		assert.Equal(t, 1, list.size)
		assert.Equal(t, "C", list.head.data)
		assert.Equal(t, list.head, list.tail)

		// Удаляем хвост
		list.DLREMOVE_VALUE("C")
		assert.Equal(t, 0, list.size)
		assert.Nil(t, list.head)
		assert.Nil(t, list.tail)
	})

	t.Run("DLLENGTH", func(t *testing.T) {
		list := NewDList()
		assert.Equal(t, 0, list.DLLENGTH())

		list.DLPUSH_HEAD("A")
		assert.Equal(t, 1, list.DLLENGTH())

		list.DLPUSH_HEAD("B")
		assert.Equal(t, 2, list.DLLENGTH())
	})
}
func TestDListSerializationErrors(t *testing.T) {
	tempDir := t.TempDir()

	t.Run("LoadDListFromText_PermissionDenied", func(t *testing.T) {
		list := NewDList()
		err := LoadDListFromText(list, "/proc/cpuinfo")
		assert.Error(t, err)
	})

	t.Run("SaveDListToText_CreateDirectory", func(t *testing.T) {
		list := NewDList()
		list.DLPUSH_TAIL("test")
		// Пытаемся сохранить в директорию (не файл)
		err := SaveDListToText(list, tempDir)
		assert.Error(t, err)
	})

	t.Run("BinaryReadError", func(t *testing.T) {
		tempFile := filepath.Join(tempDir, "test_dlist_error.bin")
		file, _ := os.Create(tempFile)
		// Пишем count=1, но не пишем данные
		binary.Write(file, binary.LittleEndian, int32(1))
		file.Close()

		list := NewDList()
		err := LoadDListFromBinary(list, tempFile)
		assert.Error(t, err)
	})
}
//
func TestDList_Additional(t *testing.T) {
	t.Run("DLPUSH_BEFORE_AtHead", func(t *testing.T) {
		list := NewDList()
		list.DLPUSH_TAIL("B")
		
		list.DLPUSH_BEFORE("B", "A")
		
		assert.Equal(t, 2, list.DLLENGTH())
		assert.Equal(t, "A", list.head.data)
		assert.Equal(t, "B", list.tail.data)
		assert.Equal(t, list.head.next, list.tail)
		assert.Equal(t, list.tail.prev, list.head)
	})

	t.Run("DLPUSH_BEFORE_InMiddle", func(t *testing.T) {
		list := NewDList()
		list.DLPUSH_TAIL("A")
		list.DLPUSH_TAIL("C")
		list.DLPUSH_TAIL("D")
		
		list.DLPUSH_BEFORE("C", "B")
		
		assert.Equal(t, 4, list.DLLENGTH())
		
		// Проверяем связи
		node1 := list.head // A
		node2 := node1.next // B
		node3 := node2.next // C
		node4 := node3.next // D
		
		assert.Equal(t, "A", node1.data)
		assert.Equal(t, "B", node2.data)
		assert.Equal(t, "C", node3.data)
		assert.Equal(t, "D", node4.data)
		
		// Проверяем next связи
		assert.Equal(t, node2, node1.next)
		assert.Equal(t, node3, node2.next)
		assert.Equal(t, node4, node3.next)
		assert.Nil(t, node4.next)
		
		// Проверяем prev связи
		assert.Nil(t, node1.prev)
		assert.Equal(t, node1, node2.prev)
		assert.Equal(t, node2, node3.prev)
		assert.Equal(t, node3, node4.prev)
	})

	t.Run("DLPUSH_AFTER_AtTail", func(t *testing.T) {
		list := NewDList()
		list.DLPUSH_TAIL("A")
		
		list.DLPUSH_AFTER("A", "B")
		
		assert.Equal(t, 2, list.DLLENGTH())
		assert.Equal(t, "A", list.head.data)
		assert.Equal(t, "B", list.tail.data)
		assert.Equal(t, list.tail, list.head.next)
		assert.Equal(t, list.head, list.tail.prev)
	})

	t.Run("DLPUSH_AFTER_InMiddle", func(t *testing.T) {
		list := NewDList()
		list.DLPUSH_TAIL("A")
		list.DLPUSH_TAIL("B")
		list.DLPUSH_TAIL("D")
		
		list.DLPUSH_AFTER("B", "C")
		
		assert.Equal(t, 4, list.DLLENGTH())
		
		// Проверяем через GetAll
		data := list.GetAll()
		assert.Equal(t, []string{"A", "B", "C", "D"}, data)
		
		// Проверяем что tail правильный
		assert.Equal(t, "D", list.tail.data)
		assert.Nil(t, list.tail.next)
	})

	t.Run("DLREMOVE_BEFORE_MiddleNode", func(t *testing.T) {
		list := NewDList()
		list.DLPUSH_TAIL("A")
		list.DLPUSH_TAIL("B")
		list.DLPUSH_TAIL("C")
		list.DLPUSH_TAIL("D")
		
		list.DLREMOVE_BEFORE("C") // Удаляет B
		
		assert.Equal(t, 3, list.DLLENGTH())
		data := list.GetAll()
		assert.Equal(t, []string{"A", "C", "D"}, data)
		
		// Проверяем связи
		assert.Equal(t, "A", list.head.data)
		assert.Equal(t, "C", list.head.next.data)
		assert.Equal(t, "D", list.tail.data)
		assert.Equal(t, list.head, list.tail.prev.prev)
	})

	t.Run("DLREMOVE_AFTER_MiddleNode", func(t *testing.T) {
		list := NewDList()
		list.DLPUSH_TAIL("A")
		list.DLPUSH_TAIL("B")
		list.DLPUSH_TAIL("C")
		list.DLPUSH_TAIL("D")
		
		list.DLREMOVE_AFTER("B") // Удаляет C
		
		assert.Equal(t, 3, list.DLLENGTH())
		data := list.GetAll()
		assert.Equal(t, []string{"A", "B", "D"}, data)
		
		// Проверяем связи
		assert.Equal(t, "A", list.head.data)
		assert.Equal(t, "B", list.head.next.data)
		assert.Equal(t, "D", list.tail.data)
		assert.Equal(t, list.tail, list.head.next.next)
		assert.Equal(t, list.head.next, list.tail.prev)
	})

	t.Run("DLREMOVE_VALUE_MiddleNode", func(t *testing.T) {
		list := NewDList()
		list.DLPUSH_TAIL("A")
		list.DLPUSH_TAIL("B")
		list.DLPUSH_TAIL("C")
		
		list.DLREMOVE_VALUE("B")
		
		assert.Equal(t, 2, list.DLLENGTH())
		assert.Equal(t, "A", list.head.data)
		assert.Equal(t, "C", list.tail.data)
		assert.Equal(t, list.tail, list.head.next)
		assert.Equal(t, list.head, list.tail.prev)
	})

	t.Run("DLGET_OptimizationCheck", func(t *testing.T) {
		list := NewDList()
		
		// Добавляем много элементов
		for i := 0; i < 100; i++ {
			list.DLPUSH_TAIL(fmt.Sprintf("Element%d", i))
		}
		
		// Проверяем доступ к элементам в начале (должен идти с головы)
		val, err := list.DLGET(0)
		assert.NoError(t, err)
		assert.Equal(t, "Element0", val)
		
		val, err = list.DLGET(1)
		assert.NoError(t, err)
		assert.Equal(t, "Element1", val)
		
		// Проверяем доступ к элементам в конце (должен идти с хвоста)
		val, err = list.DLGET(98)
		assert.NoError(t, err)
		assert.Equal(t, "Element98", val)
		
		val, err = list.DLGET(99)
		assert.NoError(t, err)
		assert.Equal(t, "Element99", val)
		
		// Проверяем доступ к среднему элементу
		val, err = list.DLGET(50)
		assert.NoError(t, err)
		assert.Equal(t, "Element50", val)
	})

	t.Run("DLPRINT_FORWARD_AND_BACKWARD", func(t *testing.T) {
		list := NewDList()
		list.DLPUSH_TAIL("A")
		list.DLPUSH_TAIL("B")
		list.DLPUSH_TAIL("C")
		
		// Просто проверяем что функции не падают
		list.DLPRINT_FORWARD()
		list.DLPRINT_BACKWARD()
		
		assert.Equal(t, 3, list.DLLENGTH())
	})

	t.Run("DLCLEAR_ThenReuse", func(t *testing.T) {
		list := NewDList()
		list.DLPUSH_TAIL("A")
		list.DLPUSH_TAIL("B")
		list.DLPUSH_TAIL("C")
		
		assert.Equal(t, 3, list.DLLENGTH())
		
		list.DLCLEAR()
		
		assert.Equal(t, 0, list.DLLENGTH())
		assert.Nil(t, list.head)
		assert.Nil(t, list.tail)
		
		// Используем снова
		list.DLPUSH_HEAD("New")
		
		assert.Equal(t, 1, list.DLLENGTH())
		assert.Equal(t, "New", list.head.data)
		assert.Equal(t, list.head, list.tail)
	})

	t.Run("SetAll_PreservesOrder", func(t *testing.T) {
		list := NewDList()
		
		items := []string{"First", "Second", "Third", "Fourth"}
		list.SetAll(items)
		
		assert.Equal(t, 4, list.DLLENGTH())
		
		// Проверяем прямые связи
		current := list.head
		for i := 0; i < len(items); i++ {
			assert.Equal(t, items[i], current.data)
			if i < len(items)-1 {
				assert.NotNil(t, current.next)
				assert.Equal(t, current, current.next.prev)
			}
			current = current.next
		}
		
		// Проверяем обратные связи
		current = list.tail
		for i := len(items) - 1; i >= 0; i-- {
			assert.Equal(t, items[i], current.data)
			if i > 0 {
				assert.NotNil(t, current.prev)
				assert.Equal(t, current, current.prev.next)
			}
			current = current.prev
		}
	})

	t.Run("EdgeCases_SingleElementList", func(t *testing.T) {
		list := NewDList()
		list.DLPUSH_TAIL("Single")
		
		// Проверяем все операции на списке с одним элементом
		assert.Equal(t, 1, list.DLLENGTH())
		assert.Equal(t, list.head, list.tail)
		assert.Nil(t, list.head.next)
		assert.Nil(t, list.head.prev)
		
		// Удаление перед/после не должно ничего делать
		list.DLREMOVE_BEFORE("Single")
		list.DLREMOVE_AFTER("Single")
		assert.Equal(t, 1, list.DLLENGTH())
		
		// Удаление по значению
		list.DLREMOVE_VALUE("Single")
		assert.Equal(t, 0, list.DLLENGTH())
		assert.Nil(t, list.head)
		assert.Nil(t, list.tail)
	})

	t.Run("EdgeCases_EmptyListOperations", func(t *testing.T) {
		list := NewDList()
		
		// Все операции должны корректно обрабатывать пустой список
		list.DLREMOVE_HEAD()
		list.DLREMOVE_TAIL()
		list.DLREMOVE_VALUE("Anything")
		list.DLREMOVE_BEFORE("X")
		list.DLREMOVE_AFTER("Y")
		list.DLPUSH_BEFORE("A", "B")
		list.DLPUSH_AFTER("A", "B")
		
		assert.Equal(t, 0, list.DLLENGTH())
		assert.Nil(t, list.head)
		assert.Nil(t, list.tail)
	})
}