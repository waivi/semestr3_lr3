package main

import (
	"testing"
	"path/filepath" 
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"os"
)

func TestSList(t *testing.T) {
	t.Run("NewSList", func(t *testing.T) {
		list := NewSList()
		assert.NotNil(t, list)
		assert.Equal(t, 0, list.size)
		assert.Nil(t, list.head)
	})

	t.Run("SLPUSH_HEAD", func(t *testing.T) {
		list := NewSList()

		list.SLPUSH_HEAD("first")
		assert.Equal(t, 1, list.size)
		assert.Equal(t, "first", list.head.data)

		list.SLPUSH_HEAD("second")
		assert.Equal(t, 2, list.size)
		assert.Equal(t, "second", list.head.data)
		assert.Equal(t, "first", list.head.next.data)
	})

	t.Run("SLPUSH_TAIL", func(t *testing.T) {
		list := NewSList()

		list.SLPUSH_TAIL("first")
		assert.Equal(t, 1, list.size)
		assert.Equal(t, "first", list.head.data)

		list.SLPUSH_TAIL("second")
		assert.Equal(t, 2, list.size)

		// Проверяем порядок
		assert.Equal(t, "first", list.head.data)
		assert.Equal(t, "second", list.head.next.data)
	})

	t.Run("SLSEARCH", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_HEAD("A")
		list.SLPUSH_HEAD("B")
		list.SLPUSH_HEAD("C")

		assert.True(t, list.SLSEARCH("A"))
		assert.True(t, list.SLSEARCH("B"))
		assert.True(t, list.SLSEARCH("C"))
		assert.False(t, list.SLSEARCH("D"))
	})

	t.Run("SLREMOVE_HEAD", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_HEAD("C")
		list.SLPUSH_HEAD("B")
		list.SLPUSH_HEAD("A")

		list.SLREMOVE_HEAD()
		assert.Equal(t, 2, list.size)
		assert.Equal(t, "B", list.head.data)

		list.SLREMOVE_HEAD()
		assert.Equal(t, 1, list.size)
		assert.Equal(t, "C", list.head.data)

		list.SLREMOVE_HEAD()
		assert.Equal(t, 0, list.size)
		assert.Nil(t, list.head)

		// Удаление из пустого списка
		list.SLREMOVE_HEAD()
		assert.Equal(t, 0, list.size)
	})

	t.Run("SLREMOVE_TAIL", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_TAIL("A")
		list.SLPUSH_TAIL("B")
		list.SLPUSH_TAIL("C")

		list.SLREMOVE_TAIL()
		assert.Equal(t, 2, list.size)

		// Проверяем что C удален
		current := list.head
		for current.next != nil {
			current = current.next
		}
		assert.Equal(t, "B", current.data)
	})

	t.Run("SLREMOVE_VALUE", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_TAIL("A")
		list.SLPUSH_TAIL("B")
		list.SLPUSH_TAIL("C")

		// Удаляем из середины
		list.SLREMOVE_VALUE("B")
		assert.Equal(t, 2, list.size)
		assert.False(t, list.SLSEARCH("B"))

		// Удаляем голову
		list.SLREMOVE_VALUE("A")
		assert.Equal(t, 1, list.size)
		assert.Equal(t, "C", list.head.data)

		// Удаляем хвост
		list.SLREMOVE_VALUE("C")
		assert.Equal(t, 0, list.size)
		assert.Nil(t, list.head)
	})

	t.Run("SLGET", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_TAIL("A")
		list.SLPUSH_TAIL("B")
		list.SLPUSH_TAIL("C")

		val, err := list.SLGET(0)
		require.NoError(t, err)
		assert.Equal(t, "A", val)

		val, err = list.SLGET(1)
		require.NoError(t, err)
		assert.Equal(t, "B", val)

		val, err = list.SLGET(2)
		require.NoError(t, err)
		assert.Equal(t, "C", val)

		_, err = list.SLGET(3)
		assert.Error(t, err)

		_, err = list.SLGET(-1)
		assert.Error(t, err)
	})

	t.Run("SLLENGTH", func(t *testing.T) {
		list := NewSList()
		assert.Equal(t, 0, list.SLLENGTH())

		list.SLPUSH_HEAD("A")
		assert.Equal(t, 1, list.SLLENGTH())

		list.SLPUSH_HEAD("B")
		assert.Equal(t, 2, list.SLLENGTH())
	})

	t.Run("GetAll_SetAll", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_TAIL("A")
		list.SLPUSH_TAIL("B")
		list.SLPUSH_TAIL("C")

		data := list.GetAll()
		assert.Equal(t, []string{"A", "B", "C"}, data)

		newList := NewSList()
		newList.SetAll([]string{"X", "Y", "Z"})
		assert.Equal(t, 3, newList.SLLENGTH())
		assert.Equal(t, []string{"X", "Y", "Z"}, newList.GetAll())
	})
}

/// В функции TestSListSerializationErrors заменить:
func TestSListSerializationErrors(t *testing.T) {
	tempDir := t.TempDir()

	t.Run("LoadSListFromText_InvalidFile", func(t *testing.T) {
		list := NewSList()
		err := LoadSListFromText(list, "/nonexistent/file.txt")
		assert.Error(t, err)
	})

	t.Run("SaveSListToText_InvalidPath", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_TAIL("test")
		// Пытаемся сохранить в защищенную директорию
		err := SaveSListToText(list, "/proc/cpuinfo/test.txt")
		assert.Error(t, err)
	})

	t.Run("LoadSListFromBinary_InvalidFile", func(t *testing.T) {
		list := NewSList()
		err := LoadSListFromBinary(list, "/nonexistent/file.bin")
		assert.Error(t, err)
	})

	t.Run("BinaryCorruption", func(t *testing.T) {
		tempFile := filepath.Join(tempDir, "corrupted.bin")
		// Пишем неполные данные
		os.WriteFile(tempFile, []byte{0x01, 0x00, 0x00, 0x00}, 0644) // Только count=1

		list := NewSList()
		err := LoadSListFromBinary(list, tempFile)
		assert.Error(t, err)
	})
}
//
func TestSList_Additional(t *testing.T) {
	t.Run("SLPUSH_BEFORE_AtHead", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_TAIL("B")
		
		list.SLPUSH_BEFORE("B", "A")
		
		assert.Equal(t, 2, list.SLLENGTH())
		assert.Equal(t, "A", list.head.data)
		assert.Equal(t, "B", list.head.next.data)
	})

	t.Run("SLPUSH_BEFORE_InMiddle", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_TAIL("A")
		list.SLPUSH_TAIL("C")
		list.SLPUSH_TAIL("D")
		
		list.SLPUSH_BEFORE("C", "B")
		
		assert.Equal(t, 4, list.SLLENGTH())
		// Проверяем порядок
		current := list.head
		assert.Equal(t, "A", current.data)
		current = current.next
		assert.Equal(t, "B", current.data)
		current = current.next
		assert.Equal(t, "C", current.data)
		current = current.next
		assert.Equal(t, "D", current.data)
	})

	t.Run("SLPUSH_BEFORE_TargetNotFound", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_TAIL("A")
		list.SLPUSH_TAIL("B")
		
		initialSize := list.SLLENGTH()
		list.SLPUSH_BEFORE("NotFound", "C")
		
		// Размер не должен измениться
		assert.Equal(t, initialSize, list.SLLENGTH())
	})

	t.Run("SLPUSH_BEFORE_EmptyList", func(t *testing.T) {
		list := NewSList()
		
		list.SLPUSH_BEFORE("Any", "Value")
		
		assert.Equal(t, 0, list.SLLENGTH())
		assert.Nil(t, list.head)
	})

	t.Run("SLPUSH_AFTER_AtTail", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_TAIL("A")
		
		list.SLPUSH_AFTER("A", "B")
		
		assert.Equal(t, 2, list.SLLENGTH())
		assert.Equal(t, "A", list.head.data)
		assert.Equal(t, "B", list.head.next.data)
		assert.Nil(t, list.head.next.next)
	})

	t.Run("SLPUSH_AFTER_InMiddle", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_TAIL("A")
		list.SLPUSH_TAIL("B")
		list.SLPUSH_TAIL("D")
		
		list.SLPUSH_AFTER("B", "C")
		
		assert.Equal(t, 4, list.SLLENGTH())
		// Проверяем порядок через GetAll
		data := list.GetAll()
		assert.Equal(t, []string{"A", "B", "C", "D"}, data)
	})

	t.Run("SLPUSH_AFTER_TargetNotFound", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_TAIL("A")
		
		initialSize := list.SLLENGTH()
		list.SLPUSH_AFTER("NotFound", "B")
		
		assert.Equal(t, initialSize, list.SLLENGTH())
	})

	t.Run("SLREMOVE_BEFORE_FirstElement", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_TAIL("A")
		list.SLPUSH_TAIL("B")
		list.SLPUSH_TAIL("C")
		
		// Нельзя удалить перед первым элементом
		initialSize := list.SLLENGTH()
		list.SLREMOVE_BEFORE("A")
		
		assert.Equal(t, initialSize, list.SLLENGTH())
	})

	t.Run("SLREMOVE_BEFORE_SecondElement", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_TAIL("A")
		list.SLPUSH_TAIL("B")
		list.SLPUSH_TAIL("C")
		
		list.SLREMOVE_BEFORE("B") // Удаляет A
		
		assert.Equal(t, 2, list.SLLENGTH())
		assert.Equal(t, "B", list.head.data)
		assert.Equal(t, "C", list.head.next.data)
	})

	t.Run("SLREMOVE_BEFORE_TargetNotFound", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_TAIL("A")
		list.SLPUSH_TAIL("B")
		
		initialSize := list.SLLENGTH()
		list.SLREMOVE_BEFORE("NotFound")
		
		assert.Equal(t, initialSize, list.SLLENGTH())
	})

	t.Run("SLREMOVE_AFTER_LastElement", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_TAIL("A")
		list.SLPUSH_TAIL("B")
		list.SLPUSH_TAIL("C")
		
		// Нельзя удалить после последнего элемента
		initialSize := list.SLLENGTH()
		list.SLREMOVE_AFTER("C")
		
		assert.Equal(t, initialSize, list.SLLENGTH())
	})

	t.Run("SLREMOVE_AFTER_TargetNotFound", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_TAIL("A")
		
		initialSize := list.SLLENGTH()
		list.SLREMOVE_AFTER("NotFound")
		
		assert.Equal(t, initialSize, list.SLLENGTH())
	})

	t.Run("SLREMOVE_VALUE_MultipleOccurrences", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_TAIL("A")
		list.SLPUSH_TAIL("B")
		list.SLPUSH_TAIL("A") // Дубликат
		list.SLPUSH_TAIL("C")
		
		// Удаляет только первое вхождение
		list.SLREMOVE_VALUE("A")
		
		assert.Equal(t, 3, list.SLLENGTH())
		data := list.GetAll()
		assert.Equal(t, []string{"B", "A", "C"}, data)
	})

	t.Run("SLPRINT_BACKWARD_SingleElement", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_TAIL("Single")
		
		// Проверяем что не падает
		list.SLPRINT_BACKWARD()
		
		assert.Equal(t, 1, list.SLLENGTH())
	})

	t.Run("SetAll_ReverseOrder", func(t *testing.T) {
		list := NewSList()
		
		items := []string{"A", "B", "C"}
		list.SetAll(items)
		
		assert.Equal(t, 3, list.SLLENGTH())
		data := list.GetAll()
		assert.Equal(t, items, data)
		
		// Проверяем что порядок правильный
		assert.Equal(t, "A", list.head.data)
		assert.Equal(t, "B", list.head.next.data)
		assert.Equal(t, "C", list.head.next.next.data)
	})

	t.Run("ConsecutiveComplexOperations", func(t *testing.T) {
		list := NewSList()
		
		// Сложная последовательность операций
		list.SLPUSH_TAIL("Start")
		list.SLPUSH_HEAD("BeforeStart")
		list.SLPUSH_AFTER("Start", "Middle")
		list.SLPUSH_BEFORE("Middle", "BeforeMiddle")
		list.SLREMOVE_VALUE("Start")
		list.SLREMOVE_HEAD()
		list.SLREMOVE_TAIL()
		list.SLPUSH_TAIL("End")
		
		assert.Equal(t, 2, list.SLLENGTH())
		data := list.GetAll()
		assert.Equal(t, []string{"BeforeMiddle", "End"}, data)
	})

	t.Run("EmptyListOperations", func(t *testing.T) {
		list := NewSList()
		
		// Все операции должны корректно обрабатывать пустой список
		list.SLREMOVE_HEAD()
		list.SLREMOVE_TAIL()
		list.SLREMOVE_VALUE("Anything")
		list.SLREMOVE_BEFORE("X")
		list.SLREMOVE_AFTER("Y")
		
		assert.Equal(t, 0, list.SLLENGTH())
		assert.Nil(t, list.head)
	})
}
// 
func TestSList_Coverage(t *testing.T) {
	t.Run("SLPRINT_FORWARD_EmptyList", func(t *testing.T) {
		list := NewSList()
		
		// Проверяем что функция не падает
		list.SLPRINT_FORWARD()
		
		assert.Equal(t, 0, list.SLLENGTH())
		assert.Nil(t, list.head)
	})

	t.Run("SLPRINT_FORWARD_SingleElement", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_TAIL("Single")
		
		// Проверяем что функция не падает
		list.SLPRINT_FORWARD()
		
		assert.Equal(t, 1, list.SLLENGTH())
		assert.Equal(t, "Single", list.head.data)
		assert.Nil(t, list.head.next)
	})

	t.Run("SLPRINT_FORWARD_MultipleElements", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_TAIL("First")
		list.SLPUSH_TAIL("Second")
		list.SLPUSH_TAIL("Third")
		
		// Проверяем что функция не падает
		list.SLPRINT_FORWARD()
		
		assert.Equal(t, 3, list.SLLENGTH())
		
		// Проверяем порядок элементов
		current := list.head
		assert.Equal(t, "First", current.data)
		current = current.next
		assert.Equal(t, "Second", current.data)
		current = current.next
		assert.Equal(t, "Third", current.data)
		assert.Nil(t, current.next)
	})

	t.Run("SLPRINT_BACKWARD_EmptyList", func(t *testing.T) {
		list := NewSList()
		
		// Проверяем что функция не падает
		list.SLPRINT_BACKWARD()
		
		assert.Equal(t, 0, list.SLLENGTH())
	})

	t.Run("SLPRINT_BACKWARD_SingleElement", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_TAIL("Single")
		
		// Проверяем что функция не падает
		list.SLPRINT_BACKWARD()
		
		assert.Equal(t, 1, list.SLLENGTH())
	})

	t.Run("SLPRINT_BACKWARD_MultipleElements", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_TAIL("A")
		list.SLPUSH_TAIL("B")
		list.SLPUSH_TAIL("C")
		
		// Проверяем что функция не падает
		list.SLPRINT_BACKWARD()
		
		assert.Equal(t, 3, list.SLLENGTH())
	})

	t.Run("printReverse_HelperFunction", func(t *testing.T) {
		list := NewSList()
		
		// Проверяем что внутренняя функция не падает на nil
		list.printReverse(nil)
		
		// Проверяем с одним элементом
		list.SLPUSH_TAIL("Test")
		list.printReverse(list.head)
		
		assert.Equal(t, 1, list.SLLENGTH())
	})

	t.Run("EdgeCases_PrintFunctions", func(t *testing.T) {
		// Проверяем все функции печати на разных состояниях списка
		testCases := []struct {
			name   string
			setup  func() *SList
		}{
			{
				name: "EmptyList",
				setup: func() *SList {
					return NewSList()
				},
			},
			{
				name: "SingleElement",
				setup: func() *SList {
					list := NewSList()
					list.SLPUSH_TAIL("A")
					return list
				},
			},
			{
				name: "TwoElements",
				setup: func() *SList {
					list := NewSList()
					list.SLPUSH_TAIL("A")
					list.SLPUSH_TAIL("B")
					return list
				},
			},
			{
				name: "ThreeElements",
				setup: func() *SList {
					list := NewSList()
					list.SLPUSH_TAIL("A")
					list.SLPUSH_TAIL("B")
					list.SLPUSH_TAIL("C")
					return list
				},
			},
			{
				name: "AfterRemoval",
				setup: func() *SList {
					list := NewSList()
					list.SLPUSH_TAIL("A")
					list.SLPUSH_TAIL("B")
					list.SLPUSH_TAIL("C")
					list.SLREMOVE_VALUE("B")
					return list
				},
			},
			{
				name: "AfterClear",
				setup: func() *SList {
					list := NewSList()
					list.SLPUSH_TAIL("A")
					list.SLPUSH_TAIL("B")
					list.SLPUSH_TAIL("C")
					list.SLCLEAR()
					return list
				},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				list := tc.setup()
				
				// Проверяем что функции печати не падают
				list.SLPRINT_FORWARD()
				list.SLPRINT_BACKWARD()
				
				// Проверяем что список в валидном состоянии
				assert.True(t, list.SLLENGTH() >= 0)
			})
		}
	})
}