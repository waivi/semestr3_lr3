package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	t.Run("NewStack", func(t *testing.T) {
		s := NewStack()
		assert.NotNil(t, s)
		assert.Equal(t, 0, s.size)
		assert.Nil(t, s.top)
	})

	t.Run("SPUSH", func(t *testing.T) {
		s := NewStack()

		s.SPUSH("first")
		assert.Equal(t, 1, s.size)
		assert.Equal(t, "first", s.top.data)

		s.SPUSH("second")
		assert.Equal(t, 2, s.size)
		assert.Equal(t, "second", s.top.data)
		assert.Equal(t, "first", s.top.next.data)
	})

	t.Run("SPOP", func(t *testing.T) {
		s := NewStack()
		s.SPUSH("A")
		s.SPUSH("B")
		s.SPUSH("C")

		val, err := s.SPOP()
		require.NoError(t, err)
		assert.Equal(t, "C", val)
		assert.Equal(t, 2, s.size)

		val, err = s.SPOP()
		require.NoError(t, err)
		assert.Equal(t, "B", val)
		assert.Equal(t, 1, s.size)

		val, err = s.SPOP()
		require.NoError(t, err)
		assert.Equal(t, "A", val)
		assert.Equal(t, 0, s.size)
		assert.Nil(t, s.top)

		// Попытка извлечь из пустого стека
		_, err = s.SPOP()
		assert.Error(t, err)
	})

	t.Run("STOP", func(t *testing.T) {
		s := NewStack()
		s.SPUSH("A")
		s.SPUSH("B")

		val, err := s.STOP()
		require.NoError(t, err)
		assert.Equal(t, "B", val)
		assert.Equal(t, 2, s.size) // Размер не должен измениться

		s.SPOP()
		val, err = s.STOP()
		require.NoError(t, err)
		assert.Equal(t, "A", val)

		s.SPOP()
		_, err = s.STOP()
		assert.Error(t, err)
	})

	t.Run("IsEmpty", func(t *testing.T) {
		s := NewStack()
		assert.True(t, s.IsEmpty())

		s.SPUSH("A")
		assert.False(t, s.IsEmpty())

		s.SPOP()
		assert.True(t, s.IsEmpty())
	})

	t.Run("Size", func(t *testing.T) {
		s := NewStack()
		assert.Equal(t, 0, s.Size())

		s.SPUSH("A")
		assert.Equal(t, 1, s.Size())

		s.SPUSH("B")
		assert.Equal(t, 2, s.Size())

		s.SPOP()
		assert.Equal(t, 1, s.Size())
	})
}

func TestStackSerializationErrors(t *testing.T) {
	tempDir := t.TempDir()

	t.Run("LoadStackFromText_EmptyFile", func(t *testing.T) {
		tempFile := filepath.Join(tempDir, "empty.txt")
		os.WriteFile(tempFile, []byte(""), 0644)

		s := NewStack()
		err := LoadStackFromText(s, tempFile)
		assert.Error(t, err)
	})

	t.Run("LoadStackFromText_InvalidSize", func(t *testing.T) {
		tempFile := filepath.Join(tempDir, "invalid_size.txt")
		os.WriteFile(tempFile, []byte("abc\n"), 0644)

		s := NewStack()
		err := LoadStackFromText(s, tempFile)
		assert.Error(t, err)
	})

	t.Run("SaveStackToText_PermissionDenied", func(t *testing.T) {
		s := NewStack()
		s.SPUSH("test")
		// Пытаемся сохранить в системную директорию
		err := SaveStackToText(s, "/proc/cpuinfo/test.txt")
		assert.Error(t, err)
	})

	t.Run("SaveStackToBinary_InvalidPath", func(t *testing.T) {
		s := NewStack()
		s.SPUSH("test")
		// Пытаемся сохранить в несуществующую директорию без прав
		err := SaveStackToBinary(s, "/root/nonexistent/test.bin")
		assert.Error(t, err)
	})
}

func TestStack_Additional(t *testing.T) {
	t.Run("SPUSH_ManyElements", func(t *testing.T) {
		s := NewStack()

		// Добавляем много элементов
		for i := 0; i < 1000; i++ {
			s.SPUSH(fmt.Sprintf("Element%d", i))
		}

		assert.Equal(t, 1000, s.Size())

		// Проверяем что верхний элемент правильный
		top, err := s.STOP()
		assert.NoError(t, err)
		assert.Equal(t, "Element999", top)
	})

	t.Run("SPOP_AllElements", func(t *testing.T) {
		s := NewStack()
		s.SPUSH("A")
		s.SPUSH("B")
		s.SPUSH("C")

		val1, err := s.SPOP()
		assert.NoError(t, err)
		assert.Equal(t, "C", val1)

		val2, err := s.SPOP()
		assert.NoError(t, err)
		assert.Equal(t, "B", val2)

		val3, err := s.SPOP()
		assert.NoError(t, err)
		assert.Equal(t, "A", val3)

		assert.True(t, s.IsEmpty())

		// Еще одна попытка извлечь из пустого стека
		_, err = s.SPOP()
		assert.Error(t, err)
	})

	t.Run("STOP_WithoutModification", func(t *testing.T) {
		s := NewStack()
		s.SPUSH("A")
		s.SPUSH("B")

		// Многократный вызов STOP не должен менять стек
		top1, err := s.STOP()
		assert.NoError(t, err)
		assert.Equal(t, "B", top1)
		assert.Equal(t, 2, s.Size())

		top2, err := s.STOP()
		assert.NoError(t, err)
		assert.Equal(t, "B", top2)
		assert.Equal(t, 2, s.Size())

		// Проверяем что SPOP работает после STOP
		pop, err := s.SPOP()
		assert.NoError(t, err)
		assert.Equal(t, "B", pop)
		assert.Equal(t, 1, s.Size())
	})

	t.Run("SCLEAR_AfterOperations", func(t *testing.T) {
		s := NewStack()

		// Добавляем элементы
		for i := 0; i < 10; i++ {
			s.SPUSH(fmt.Sprintf("Data%d", i))
		}

		assert.Equal(t, 10, s.Size())
		assert.False(t, s.IsEmpty())

		// Очищаем
		s.SCLEAR()

		assert.Equal(t, 0, s.Size())
		assert.True(t, s.IsEmpty())
		assert.Nil(t, s.top)

		// Можем использовать снова
		s.SPUSH("New")
		assert.Equal(t, 1, s.Size())
		top, err := s.STOP()
		assert.NoError(t, err)
		assert.Equal(t, "New", top)
	})

	t.Run("GetAll_PreservesOrder", func(t *testing.T) {
		s := NewStack()

		// Добавляем в порядке A, B, C
		s.SPUSH("A")
		s.SPUSH("B")
		s.SPUSH("C")

		// GetAll должен вернуть [C, B, A] (от вершины к основанию)
		data := s.GetAll()
		assert.Equal(t, []string{"C", "B", "A"}, data)
	})

	t.Run("SetAll_ReconstructsStack", func(t *testing.T) {
		s := NewStack()

		items := []string{"C", "B", "A"} // Порядок от вершины к основанию
		s.SetAll(items)

		assert.Equal(t, 3, s.Size())

		// Проверяем что стек работает правильно
		top, err := s.STOP()
		assert.NoError(t, err)
		assert.Equal(t, "C", top)

		pop1, err := s.SPOP()
		assert.NoError(t, err)
		assert.Equal(t, "C", pop1)

		pop2, err := s.SPOP()
		assert.NoError(t, err)
		assert.Equal(t, "B", pop2)

		pop3, err := s.SPOP()
		assert.NoError(t, err)
		assert.Equal(t, "A", pop3)

		assert.True(t, s.IsEmpty())
	})

	t.Run("SetAll_EmptySlice", func(t *testing.T) {
		s := NewStack()
		s.SPUSH("SomeData")

		s.SetAll([]string{}) // Устанавливаем пустой стек

		assert.Equal(t, 0, s.Size())
		assert.True(t, s.IsEmpty())
		assert.Nil(t, s.top)
	})

	t.Run("SetAll_NilSlice", func(t *testing.T) {
		s := NewStack()
		s.SPUSH("SomeData")

		s.SetAll(nil) // Устанавливаем nil

		assert.Equal(t, 0, s.Size())
		assert.True(t, s.IsEmpty())
		assert.Nil(t, s.top)
	})

	t.Run("ConsecutiveOperations_Complex", func(t *testing.T) {
		s := NewStack()

		// Сложная последовательность операций
		s.SPUSH("A")
		s.SPUSH("B")
		s.SPOP() // Удаляет B
		s.SPUSH("C")
		s.SPUSH("D")
		s.STOP() // Просматривает D
		s.SCLEAR()
		s.SPUSH("E")
		s.SPUSH("F")
		s.SPOP() // Удаляет F

		assert.Equal(t, 1, s.Size())
		assert.False(t, s.IsEmpty())

		top, err := s.STOP()
		assert.NoError(t, err)
		assert.Equal(t, "E", top)

		data := s.GetAll()
		assert.Equal(t, []string{"E"}, data)
	})

	t.Run("SPRINT_EmptyStack", func(t *testing.T) {
		s := NewStack()

		// Проверяем что не падает
		s.SPRINT()

		assert.True(t, s.IsEmpty())
	})

	t.Run("SPRINT_SingleElement", func(t *testing.T) {
		s := NewStack()
		s.SPUSH("Single")

		// Проверяем что не падает
		s.SPRINT()

		assert.Equal(t, 1, s.Size())
	})

	t.Run("EdgeCases_LargeStack", func(t *testing.T) {
		s := NewStack()

		// Создаем большой стек
		for i := 0; i < 10000; i++ {
			s.SPUSH(fmt.Sprintf("Item%d", i))
		}

		assert.Equal(t, 10000, s.Size())

		// Извлекаем половину
		for i := 0; i < 5000; i++ {
			_, err := s.SPOP()
			assert.NoError(t, err)
		}

		assert.Equal(t, 5000, s.Size())

		// Проверяем верхний элемент
		top, err := s.STOP()
		assert.NoError(t, err)
		assert.Equal(t, "Item4999", top)
	})

	t.Run("MultipleClearOperations", func(t *testing.T) {
		s := NewStack()

		// Многократное очищение
		s.SCLEAR()
		s.SCLEAR() // Очистка уже пустого стека
		s.SCLEAR()

		assert.True(t, s.IsEmpty())
		assert.Equal(t, 0, s.Size())

		// Добавляем после многократного очищения
		s.SPUSH("Test")
		assert.Equal(t, 1, s.Size())
	})
}
