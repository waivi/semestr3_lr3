package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	t.Run("NewQueue", func(t *testing.T) {
		q := NewQueue()
		assert.NotNil(t, q)
		assert.Equal(t, 0, q.size)
		assert.Nil(t, q.front)
		assert.Nil(t, q.rear)
	})

	t.Run("QPUSH", func(t *testing.T) {
		q := NewQueue()

		q.QPUSH("first")
		assert.Equal(t, 1, q.size)
		assert.Equal(t, "first", q.front.data)
		assert.Equal(t, "first", q.rear.data)
		assert.Equal(t, q.front, q.rear)

		q.QPUSH("second")
		assert.Equal(t, 2, q.size)
		assert.Equal(t, "first", q.front.data)
		assert.Equal(t, "second", q.rear.data)
		assert.Equal(t, q.front.next, q.rear)
	})

	t.Run("QPOP", func(t *testing.T) {
		q := NewQueue()
		q.QPUSH("A")
		q.QPUSH("B")
		q.QPUSH("C")

		val, err := q.QPOP()
		require.NoError(t, err)
		assert.Equal(t, "A", val)
		assert.Equal(t, 2, q.size)
		assert.Equal(t, "B", q.front.data)

		val, err = q.QPOP()
		require.NoError(t, err)
		assert.Equal(t, "B", val)
		assert.Equal(t, 1, q.size)
		assert.Equal(t, "C", q.front.data)
		assert.Equal(t, q.front, q.rear)

		val, err = q.QPOP()
		require.NoError(t, err)
		assert.Equal(t, "C", val)
		assert.Equal(t, 0, q.size)
		assert.Nil(t, q.front)
		assert.Nil(t, q.rear)

		// Попытка извлечь из пустой очереди
		_, err = q.QPOP()
		assert.Error(t, err)
	})

	t.Run("QFRONT", func(t *testing.T) {
		q := NewQueue()
		q.QPUSH("A")
		q.QPUSH("B")

		val, err := q.QFRONT()
		require.NoError(t, err)
		assert.Equal(t, "A", val)
		assert.Equal(t, 2, q.size) // Размер не должен измениться

		q.QPOP()
		val, err = q.QFRONT()
		require.NoError(t, err)
		assert.Equal(t, "B", val)

		q.QPOP()
		_, err = q.QFRONT()
		assert.Error(t, err)
	})

	t.Run("IsEmpty", func(t *testing.T) {
		q := NewQueue()
		assert.True(t, q.IsEmpty())

		q.QPUSH("A")
		assert.False(t, q.IsEmpty())

		q.QPOP()
		assert.True(t, q.IsEmpty())
	})

	t.Run("Size", func(t *testing.T) {
		q := NewQueue()
		assert.Equal(t, 0, q.Size())

		q.QPUSH("A")
		assert.Equal(t, 1, q.Size())

		q.QPUSH("B")
		assert.Equal(t, 2, q.Size())

		q.QPOP()
		assert.Equal(t, 1, q.Size())
	})
}

func TestQueueSerializationErrors(t *testing.T) {
	tempDir := t.TempDir()

	t.Run("LoadQueueFromText_MissingData", func(t *testing.T) {
		tempFile := filepath.Join(tempDir, "missing_data.txt")
		// Указываем size=3, но даем только 2 элемента
		content := "3\nitem1\nitem2\n" // Добавили новую строку в конце
		os.WriteFile(tempFile, []byte(content), 0644)

		q := NewQueue()
		err := LoadQueueFromText(q, tempFile)
		// Теперь должен вернуть ошибку из-за недостающих данных
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unexpected EOF")
	})

	t.Run("LoadQueueFromText_InvalidSize", func(t *testing.T) {
		tempFile := filepath.Join(tempDir, "invalid_size.txt")
		os.WriteFile(tempFile, []byte("abc\nitem1\n"), 0644)

		q := NewQueue()
		err := LoadQueueFromText(q, tempFile)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid size format")
	})

	t.Run("BinaryWriteError", func(t *testing.T) {
		q := NewQueue()
		q.QPUSH("test")
		// Пытаемся записать в защищенную директорию
		err := SaveQueueToBinary(q, "/proc/cpuinfo/test.bin")
		assert.Error(t, err)
	})
}

func TestQueue_Additional(t *testing.T) {
	t.Run("QPUSH_ManyElements", func(t *testing.T) {
		q := NewQueue()

		// Добавляем много элементов
		for i := 0; i < 1000; i++ {
			q.QPUSH(fmt.Sprintf("Element%d", i))
		}

		assert.Equal(t, 1000, q.Size())

		// Проверяем что front и rear правильные
		front, err := q.QFRONT()
		assert.NoError(t, err)
		assert.Equal(t, "Element0", front)

		// Проверяем rear через извлечение всех элементов
		for i := 0; i < 999; i++ {
			_, err := q.QPOP()
			assert.NoError(t, err)
		}

		last, err := q.QPOP()
		assert.NoError(t, err)
		assert.Equal(t, "Element999", last)
		assert.True(t, q.IsEmpty())
	})

	t.Run("QPOP_AllElements", func(t *testing.T) {
		q := NewQueue()
		q.QPUSH("A")
		q.QPUSH("B")
		q.QPUSH("C")

		val1, err := q.QPOP()
		assert.NoError(t, err)
		assert.Equal(t, "A", val1)

		val2, err := q.QPOP()
		assert.NoError(t, err)
		assert.Equal(t, "B", val2)

		val3, err := q.QPOP()
		assert.NoError(t, err)
		assert.Equal(t, "C", val3)

		assert.True(t, q.IsEmpty())
		assert.Nil(t, q.front)
		assert.Nil(t, q.rear)

		// Еще одна попытка извлечь из пустой очереди
		_, err = q.QPOP()
		assert.Error(t, err)
	})

	t.Run("QFRONT_WithoutModification", func(t *testing.T) {
		q := NewQueue()
		q.QPUSH("A")
		q.QPUSH("B")

		// Многократный вызов QFRONT не должен менять очередь
		front1, err := q.QFRONT()
		assert.NoError(t, err)
		assert.Equal(t, "A", front1)
		assert.Equal(t, 2, q.Size())

		front2, err := q.QFRONT()
		assert.NoError(t, err)
		assert.Equal(t, "A", front2)
		assert.Equal(t, 2, q.Size())

		// Проверяем что QPOP работает после QFRONT
		pop, err := q.QPOP()
		assert.NoError(t, err)
		assert.Equal(t, "A", pop)
		assert.Equal(t, 1, q.Size())
	})

	t.Run("QCLEAR_AfterOperations", func(t *testing.T) {
		q := NewQueue()

		// Добавляем элементы
		for i := 0; i < 10; i++ {
			q.QPUSH(fmt.Sprintf("Data%d", i))
		}

		assert.Equal(t, 10, q.Size())
		assert.False(t, q.IsEmpty())

		// Очищаем
		q.QCLEAR()

		assert.Equal(t, 0, q.Size())
		assert.True(t, q.IsEmpty())
		assert.Nil(t, q.front)
		assert.Nil(t, q.rear)

		// Можем использовать снова
		q.QPUSH("New")
		assert.Equal(t, 1, q.Size())
		front, err := q.QFRONT()
		assert.NoError(t, err)
		assert.Equal(t, "New", front)
		assert.Equal(t, q.front, q.rear)
	})

	t.Run("GetAll_PreservesFIFOOrder", func(t *testing.T) {
		q := NewQueue()

		// Добавляем в порядке A, B, C
		q.QPUSH("A")
		q.QPUSH("B")
		q.QPUSH("C")

		// GetAll должен вернуть [A, B, C] (порядок FIFO)
		data := q.GetAll()
		assert.Equal(t, []string{"A", "B", "C"}, data)
	})

	t.Run("SetAll_ReconstructsQueue", func(t *testing.T) {
		q := NewQueue()

		items := []string{"A", "B", "C"} // Порядок FIFO
		q.SetAll(items)

		assert.Equal(t, 3, q.Size())

		// Проверяем что очередь работает правильно
		front, err := q.QFRONT()
		assert.NoError(t, err)
		assert.Equal(t, "A", front)

		pop1, err := q.QPOP()
		assert.NoError(t, err)
		assert.Equal(t, "A", pop1)

		pop2, err := q.QPOP()
		assert.NoError(t, err)
		assert.Equal(t, "B", pop2)

		pop3, err := q.QPOP()
		assert.NoError(t, err)
		assert.Equal(t, "C", pop3)

		assert.True(t, q.IsEmpty())
	})

	t.Run("SetAll_EmptySlice", func(t *testing.T) {
		q := NewQueue()
		q.QPUSH("SomeData")

		q.SetAll([]string{}) // Устанавливаем пустую очередь

		assert.Equal(t, 0, q.Size())
		assert.True(t, q.IsEmpty())
		assert.Nil(t, q.front)
		assert.Nil(t, q.rear)
	})

	t.Run("SetAll_NilSlice", func(t *testing.T) {
		q := NewQueue()
		q.QPUSH("SomeData")

		q.SetAll(nil) // Устанавливаем nil

		assert.Equal(t, 0, q.Size())
		assert.True(t, q.IsEmpty())
		assert.Nil(t, q.front)
		assert.Nil(t, q.rear)
	})

	t.Run("ConsecutiveOperations_Complex", func(t *testing.T) {
		q := NewQueue()

		// Сложная последовательность операций
		q.QPUSH("A")
		q.QPUSH("B")
		q.QPOP() // Удаляет A
		q.QPUSH("C")
		q.QPUSH("D")
		q.QFRONT() // Просматривает B
		q.QCLEAR()
		q.QPUSH("E")
		q.QPUSH("F")
		q.QPOP() // Удаляет E

		assert.Equal(t, 1, q.Size())
		assert.False(t, q.IsEmpty())

		front, err := q.QFRONT()
		assert.NoError(t, err)
		assert.Equal(t, "F", front)

		data := q.GetAll()
		assert.Equal(t, []string{"F"}, data)

		// Проверяем что front и rear указывают на один узел
		assert.Equal(t, q.front, q.rear)
		assert.Equal(t, "F", q.front.data)
		assert.Nil(t, q.front.next)
	})

	t.Run("QPRINT_EmptyQueue", func(t *testing.T) {
		q := NewQueue()

		// Проверяем что не падает
		q.QPRINT()

		assert.True(t, q.IsEmpty())
	})

	t.Run("QPRINT_SingleElement", func(t *testing.T) {
		q := NewQueue()
		q.QPUSH("Single")

		// Проверяем что не падает
		q.QPRINT()

		assert.Equal(t, 1, q.Size())
	})

	t.Run("EdgeCases_LargeQueue", func(t *testing.T) {
		q := NewQueue()

		// Создаем большую очередь
		for i := 0; i < 10000; i++ {
			q.QPUSH(fmt.Sprintf("Item%d", i))
		}

		assert.Equal(t, 10000, q.Size())

		// Извлекаем половину
		for i := 0; i < 5000; i++ {
			_, err := q.QPOP()
			assert.NoError(t, err)
		}

		assert.Equal(t, 5000, q.Size())

		// Проверяем front элемент
		front, err := q.QFRONT()
		assert.NoError(t, err)
		assert.Equal(t, "Item5000", front)

		// Проверяем что rear правильный
		// Извлекаем все кроме последнего
		for i := 0; i < 4999; i++ {
			_, err := q.QPOP()
			assert.NoError(t, err)
		}

		last, err := q.QPOP()
		assert.NoError(t, err)
		assert.Equal(t, "Item9999", last)
		assert.True(t, q.IsEmpty())
	})

	t.Run("MultipleClearOperations", func(t *testing.T) {
		q := NewQueue()

		// Многократное очищение
		q.QCLEAR()
		q.QCLEAR() // Очистка уже пустой очереди
		q.QCLEAR()

		assert.True(t, q.IsEmpty())
		assert.Equal(t, 0, q.Size())

		// Добавляем после многократного очищения
		q.QPUSH("Test")
		assert.Equal(t, 1, q.Size())
	})

	t.Run("SingleElementQueueOperations", func(t *testing.T) {
		q := NewQueue()
		q.QPUSH("Single")

		assert.Equal(t, 1, q.Size())
		assert.Equal(t, q.front, q.rear)
		assert.Nil(t, q.front.next)

		// Проверяем QFRONT и QPOP
		front, err := q.QFRONT()
		assert.NoError(t, err)
		assert.Equal(t, "Single", front)

		pop, err := q.QPOP()
		assert.NoError(t, err)
		assert.Equal(t, "Single", pop)

		assert.True(t, q.IsEmpty())
		assert.Nil(t, q.front)
		assert.Nil(t, q.rear)
	})
}
