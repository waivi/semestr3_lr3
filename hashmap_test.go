package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChainingHashTable(t *testing.T) {
	t.Run("NewChainingHashTable", func(t *testing.T) {
		hm := NewChainingHashTable(10)
		assert.NotNil(t, hm)
		assert.Equal(t, 10, hm.capacity)
		assert.NotNil(t, hm.table)
		assert.Equal(t, 10, len(hm.table))
		assert.Equal(t, 0, hm.size)
	})

	t.Run("Add", func(t *testing.T) {
		hm := NewChainingHashTable(5)

		// Добавление элементов
		hm.Add(1, 100)
		hm.Add(2, 200)
		hm.Add(3, 300)

		// Проверка добавления
		result1 := hm.Contains(1)
		assert.True(t, result1.first)
		assert.Equal(t, 100, result1.second)

		result2 := hm.Contains(2)
		assert.True(t, result2.first)
		assert.Equal(t, 200, result2.second)

		// Обновление значения (дубликаты не должны добавляться)
		hm.Add(1, 150) // Не должно обновить, так как дубликаты не добавляются
		result1 = hm.Contains(1)
		assert.True(t, result1.first)
		assert.Equal(t, 100, result1.second) // Должно остаться старое значение

		// Коллизии
		hmSmall := NewChainingHashTable(2)
		hmSmall.Add(1, 100)
		hmSmall.Add(2, 200)
		hmSmall.Add(3, 300) // Должна быть коллизия

		resultA := hmSmall.Contains(1)
		assert.True(t, resultA.first)
		assert.Equal(t, 100, resultA.second)

		resultB := hmSmall.Contains(2)
		assert.True(t, resultB.first)
		assert.Equal(t, 200, resultB.second)

		resultC := hmSmall.Contains(3)
		assert.True(t, resultC.first)
		assert.Equal(t, 300, resultC.second)
	})

	t.Run("Contains", func(t *testing.T) {
		hm := NewChainingHashTable(10)

		// Поиск в пустой таблице
		result := hm.Contains(999)
		assert.False(t, result.first)
		assert.Equal(t, -1, result.second)

		// Добавляем элементы
		hm.Add(10, 100)
		hm.Add(20, 200)
		hm.Add(30, 300)

		// Поиск существующих элементов
		result1 := hm.Contains(10)
		assert.True(t, result1.first)
		assert.Equal(t, 100, result1.second)

		result2 := hm.Contains(20)
		assert.True(t, result2.first)
		assert.Equal(t, 200, result2.second)

		// Поиск несуществующего элемента
		result = hm.Contains(40)
		assert.False(t, result.first)
		assert.Equal(t, -1, result.second)
	})

	t.Run("Remove", func(t *testing.T) {
		hm := NewChainingHashTable(10)

		// Удаление из пустой таблицы
		initialSize := hm.size
		hm.Remove(999)
		assert.Equal(t, initialSize, hm.size)

		// Добавляем элементы
		hm.Add(10, 100)
		hm.Add(20, 200)
		hm.Add(30, 300)

		// Удаление существующего элемента
		hm.Remove(20)

		// Проверяем что элемент удален
		result := hm.Contains(20)
		assert.False(t, result.first)

		// Проверяем что другие элементы остались
		result1 := hm.Contains(10)
		assert.True(t, result1.first)
		assert.Equal(t, 100, result1.second)

		result3 := hm.Contains(30)
		assert.True(t, result3.first)
		assert.Equal(t, 300, result3.second)

		// Удаление несуществующего элемента
		initialSize = hm.size
		hm.Remove(999)
		assert.Equal(t, initialSize, hm.size)
	})

	t.Run("GetAllElements", func(t *testing.T) {
		hm := NewChainingHashTable(5)

		// Пустая таблица
		all := hm.GetAllElements()
		assert.Equal(t, 0, len(all))

		// Добавляем элементы
		hm.Add(1, 100)
		hm.Add(2, 200)
		hm.Add(3, 300)

		// Получаем все элементы
		all = hm.GetAllElements()
		assert.Equal(t, 3, len(all))

		// Проверяем наличие всех пар
		expected := map[int]int{
			1: 100,
			2: 200,
			3: 300,
		}

		for _, elem := range all {
			expectedValue, exists := expected[elem[0]]
			assert.True(t, exists)
			assert.Equal(t, expectedValue, elem[1])
		}

		// Проверяем после удаления
		hm.Remove(2)
		all = hm.GetAllElements()
		assert.Equal(t, 2, len(all))

		found := false
		for _, elem := range all {
			if elem[0] == 2 {
				found = true
			}
		}
		assert.False(t, found)
	})

	t.Run("CollisionHandling", func(t *testing.T) {
		// Создаем маленькую таблицу для создания коллизий
		hm := NewChainingHashTable(2)

		// Эти ключи должны давать одинаковый хэш
		hm.Add(1, 100) // 1 % 2 = 1
		hm.Add(3, 300) // 3 % 2 = 1 - коллизия
		hm.Add(5, 500) // 5 % 2 = 1 - коллизия
		hm.Add(7, 700) // 7 % 2 = 1 - коллизия

		// Проверяем что все значения доступны
		result1 := hm.Contains(1)
		assert.True(t, result1.first)
		assert.Equal(t, 100, result1.second)

		result3 := hm.Contains(3)
		assert.True(t, result3.first)
		assert.Equal(t, 300, result3.second)

		result5 := hm.Contains(5)
		assert.True(t, result5.first)
		assert.Equal(t, 500, result5.second)

		result7 := hm.Contains(7)
		assert.True(t, result7.first)
		assert.Equal(t, 700, result7.second)

		// Удаляем элемент в цепочке
		hm.Remove(3)

		result3 = hm.Contains(3)
		assert.False(t, result3.first)

		// Проверяем что другие элементы цепочки остались
		result1 = hm.Contains(1)
		assert.True(t, result1.first)
		assert.Equal(t, 100, result1.second)

		result5 = hm.Contains(5)
		assert.True(t, result5.first)
		assert.Equal(t, 500, result5.second)
	})

	t.Run("GetChainLengths", func(t *testing.T) {
		hm := NewChainingHashTable(10)

		// Пустая таблица
		minLen, maxLen, avgLen := hm.GetChainLengths()
		assert.Equal(t, 0, minLen)
		assert.Equal(t, 0, maxLen)
		assert.Equal(t, 0.0, avgLen)

		// Добавляем элементы
		for i := 0; i < 50; i++ {
			hm.Add(i, i*10)
		}

		minLen, maxLen, avgLen = hm.GetChainLengths()
		assert.True(t, minLen >= 0)
		assert.True(t, maxLen > 0)
		assert.True(t, avgLen > 0.0)
	})

	t.Run("SizeAndCapacity", func(t *testing.T) {
		hm := NewChainingHashTable(20)
		assert.Equal(t, 20, hm.GetCapacity())
		assert.Equal(t, 0, hm.GetSize())

		// Добавляем элементы
		for i := 0; i < 15; i++ {
			hm.Add(i, i*100)
		}

		assert.Equal(t, 15, hm.GetSize())
		assert.Equal(t, 20, hm.GetCapacity())

		// Удаляем часть элементов
		for i := 0; i < 5; i++ {
			hm.Remove(i)
		}

		assert.Equal(t, 10, hm.GetSize())
	})

	t.Run("LoadFactor", func(t *testing.T) {
		hm := NewChainingHashTable(10)

		// Пустая таблица
		assert.Equal(t, 0.0, hm.GetLoadFactor())

		// Добавляем 5 элементов
		for i := 0; i < 5; i++ {
			hm.Add(i, i*10)
		}

		// Фактор загрузки = 5/10 = 0.5
		assert.Equal(t, 0.5, hm.GetLoadFactor())

		// Добавляем еще 5 элементов
		for i := 5; i < 10; i++ {
			hm.Add(i, i*10)
		}

		// Фактор загрузки = 10/10 = 1.0
		assert.Equal(t, 1.0, hm.GetLoadFactor())
	})

	t.Run("Clear", func(t *testing.T) {
		hm := NewChainingHashTable(10)

		// Добавляем элементы
		for i := 0; i < 10; i++ {
			hm.Add(i, i*10)
		}

		assert.Equal(t, 10, hm.GetSize())
		assert.False(t, hm.IsEmpty())

		// Очищаем
		hm.Clear()

		assert.Equal(t, 0, hm.GetSize())
		assert.True(t, hm.IsEmpty())

		// Проверяем что элементы действительно удалены
		for i := 0; i < 10; i++ {
			result := hm.Contains(i)
			assert.False(t, result.first)
		}
	})

	t.Run("ToString", func(t *testing.T) {
		hm := NewChainingHashTable(3)
		hm.Add(1, 100)
		hm.Add(2, 200)
		hm.Add(3, 300)

		str := hm.ToString()
		assert.Contains(t, str, "(1,100)")
		assert.Contains(t, str, "(2,200)")
		assert.Contains(t, str, "(3,300)")
	})

	t.Run("PerformanceTest", func(t *testing.T) {
		hm := NewChainingHashTable(100)

		// Добавляем много элементов
		for i := 0; i < 1000; i++ {
			hm.Add(i, i*2)
		}

		// Проверяем поиск
		for i := 0; i < 100; i++ {
			result := hm.Contains(i)
			assert.True(t, result.first)
			assert.Equal(t, i*2, result.second)
		}

		// Удаляем часть элементов
		for i := 0; i < 500; i++ {
			hm.Remove(i)
		}

		// Проверяем что остальные элементы доступны
		for i := 500; i < 1000; i++ {
			result := hm.Contains(i)
			assert.True(t, result.first)
			assert.Equal(t, i*2, result.second)
		}
	})
}

func TestHashMapSerializationErrors(t *testing.T) {
	tempDir := t.TempDir()

	t.Run("LoadHashMapFromText_FormatError", func(t *testing.T) {
		tempFile := filepath.Join(tempDir, "invalid_hashmap.txt")
		// Неправильный формат строки
		content := "2\n1 100\ninvalid_line\n"
		os.WriteFile(tempFile, []byte(content), 0644)

		ht := NewChainingHashTable(10)
		err := LoadHashMapFromText(ht, tempFile)
		assert.Error(t, err)
	})

	t.Run("SaveHashMapToText_PermissionError", func(t *testing.T) {
		ht := NewChainingHashTable(10)
		ht.Add(1, 100)
		// Пытаемся записать в защищенную директорию
		err := SaveHashMapToText(ht, "/proc/cpuinfo/hashmap.txt")
		assert.Error(t, err)
	})

	t.Run("LoadHashMapFromBinary_Corrupted", func(t *testing.T) {
		tempFile := filepath.Join(tempDir, "corrupted_hashmap.bin")
		// Коррумпированные данные
		os.WriteFile(tempFile, []byte{0xFF, 0xFF, 0xFF, 0xFF}, 0644)

		ht := NewChainingHashTable(10)
		err := LoadHashMapFromBinary(ht, tempFile)
		assert.Error(t, err)
	})

	t.Run("LoadHashMapFromBinary_InvalidData", func(t *testing.T) {
		tempFile := filepath.Join(tempDir, "invalid_data.bin")
		file, _ := os.Create(tempFile)
		// Пишем count=1, но не пишем пару ключ-значение полностью
		binary.Write(file, binary.LittleEndian, int32(1))
		binary.Write(file, binary.LittleEndian, int32(10)) // key
		// Не пишем value
		file.Close()

		ht := NewChainingHashTable(10)
		err := LoadHashMapFromBinary(ht, tempFile)
		assert.Error(t, err)
	})

	t.Run("LoadHashMapFromBinary_Corrupted", func(t *testing.T) {
		tempFile := filepath.Join(tempDir, "corrupted_hashmap.bin")
		// Коррумпированные данные - записываем меньше байт, чем ожидается
		os.WriteFile(tempFile, []byte{0x01, 0x00, 0x00, 0x00}, 0644) // count = 1

		ht := NewChainingHashTable(10)
		err := LoadHashMapFromBinary(ht, tempFile)
		// Ожидаем ошибку, потому что count=1, но нет данных
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "неожиданный конец файла")
	})
}
func TestChainingHashTable_Additional(t *testing.T) {
	t.Run("Add_DuplicateKeyShouldNotUpdate", func(t *testing.T) {
		ht := NewChainingHashTable(10)
		ht.Add(1, 100)
		ht.Add(1, 200) // Дубликат

		result := ht.Contains(1)
		assert.True(t, result.first)
		assert.Equal(t, 100, result.second) // Должно остаться старое значение
		assert.Equal(t, 1, ht.GetSize())
	})

	t.Run("Add_MultipleCollisions", func(t *testing.T) {
		ht := NewChainingHashTable(2) // Маленькая таблица для гарантии коллизий

		// Добавляем элементы, которые будут иметь одинаковый хэш
		for i := 0; i < 10; i++ {
			ht.Add(i, i*10)
		}

		assert.Equal(t, 10, ht.GetSize())

		// Проверяем все элементы
		for i := 0; i < 10; i++ {
			result := ht.Contains(i)
			assert.True(t, result.first)
			assert.Equal(t, i*10, result.second)
		}
	})

	t.Run("Remove_MiddleOfChain", func(t *testing.T) {
		ht := NewChainingHashTable(2) // Все элементы в одной цепочке

		ht.Add(1, 100) // Все имеют хэш 1 % 2 = 1
		ht.Add(3, 300)
		ht.Add(5, 500)
		ht.Add(7, 700)

		assert.Equal(t, 4, ht.GetSize())

		// Удаляем элемент из середины цепочки
		ht.Remove(3)

		assert.Equal(t, 3, ht.GetSize())
		assert.False(t, ht.Contains(3).first)
		assert.True(t, ht.Contains(1).first)
		assert.True(t, ht.Contains(5).first)
		assert.True(t, ht.Contains(7).first)
	})

	t.Run("Remove_HeadOfChain", func(t *testing.T) {
		ht := NewChainingHashTable(2)

		ht.Add(1, 100)
		ht.Add(3, 300)
		ht.Add(5, 500)

		// Удаляем голову цепочки
		ht.Remove(1)

		assert.Equal(t, 2, ht.GetSize())
		assert.False(t, ht.Contains(1).first)
		assert.True(t, ht.Contains(3).first)
		assert.True(t, ht.Contains(5).first)
	})

	t.Run("ToString_EmptyTable", func(t *testing.T) {
		ht := NewChainingHashTable(5)

		str := ht.ToString()
		assert.Contains(t, str, "[  0]: empty")
		assert.Contains(t, str, "[  4]: empty")
	})

	t.Run("GetAllElements_OrderConsistency", func(t *testing.T) {
		ht := NewChainingHashTable(5)

		// Добавляем элементы в случайном порядке
		ht.Add(15, 150)
		ht.Add(3, 30)
		ht.Add(8, 80)
		ht.Add(20, 200)
		ht.Add(11, 110)

		elements := ht.GetAllElements()
		assert.Equal(t, 5, len(elements))

		// Проверяем что все элементы присутствуют
		expectedMap := map[int]int{
			15: 150, 3: 30, 8: 80, 20: 200, 11: 110,
		}

		for _, elem := range elements {
			expectedValue, exists := expectedMap[elem[0]]
			assert.True(t, exists)
			assert.Equal(t, expectedValue, elem[1])
		}
	})

	t.Run("Clear_ThenAdd", func(t *testing.T) {
		ht := NewChainingHashTable(10)

		// Добавляем элементы
		for i := 0; i < 5; i++ {
			ht.Add(i, i*10)
		}

		assert.Equal(t, 5, ht.GetSize())
		assert.Equal(t, 0.5, ht.GetLoadFactor())

		// Очищаем
		ht.Clear()

		assert.Equal(t, 0, ht.GetSize())
		assert.Equal(t, 0.0, ht.GetLoadFactor())
		assert.True(t, ht.IsEmpty())

		// Добавляем снова
		ht.Add(100, 1000)
		assert.Equal(t, 1, ht.GetSize())
		assert.True(t, ht.Contains(100).first)
	})

	t.Run("GetChainLengths_VariousScenarios", func(t *testing.T) {
		t.Run("AllEmptyChains", func(t *testing.T) {
			ht := NewChainingHashTable(10)
			min, max, avg := ht.GetChainLengths()
			assert.Equal(t, 0, min)
			assert.Equal(t, 0, max)
			assert.Equal(t, 0.0, avg)
		})

	t.Run("SingleElementPerChain", func(t *testing.T) {
		ht := NewChainingHashTable(5)
		
		// Добавляем элементы с разными хэшами
		for i := 0; i < 5; i++ {
			ht.Add(i, i*100) // Разные хэши: i % 5 = i
		}

		min, max, avg := ht.GetChainLengths()
		assert.Equal(t, 1, min)
		assert.Equal(t, 1, max)
		assert.Equal(t, 1.0, avg)
	})

		t.Run("MultipleElementsInSingleChain", func(t *testing.T) {
			ht := NewChainingHashTable(1) // Все в одной цепочке
			for i := 0; i < 5; i++ {
				ht.Add(i, i*10)
			}

			min, max, avg := ht.GetChainLengths()
			assert.Equal(t, 5, min)
			assert.Equal(t, 5, max)
			assert.Equal(t, 5.0, avg)
		})
	})

	t.Run("LoadFactor_BoundaryCases", func(t *testing.T) {
		ht := NewChainingHashTable(0) // capacity = 0
		assert.Equal(t, 0.0, ht.GetLoadFactor())

		ht2 := NewChainingHashTable(10)
		assert.Equal(t, 0.0, ht2.GetLoadFactor()) // Пустая таблица

		// Добавляем до capacity
		for i := 0; i < 10; i++ {
			ht2.Add(i, i*10)
		}
		assert.Equal(t, 1.0, ht2.GetLoadFactor())

		// Превышаем capacity
		ht2.Add(10, 100)
		assert.Equal(t, 1.1, ht2.GetLoadFactor())
	})

	t.Run("Contains_AfterMultipleOperations", func(t *testing.T) {
		ht := NewChainingHashTable(10)

		// Серия операций
		ht.Add(1, 100)
		assert.True(t, ht.Contains(1).first)

		ht.Add(2, 200)
		assert.True(t, ht.Contains(2).first)

		ht.Remove(1)
		assert.False(t, ht.Contains(1).first)
		assert.True(t, ht.Contains(2).first)

		ht.Clear()
		assert.False(t, ht.Contains(2).first)

		ht.Add(3, 300)
		assert.True(t, ht.Contains(3).first)
	})

	t.Run("NewChainingHashTable_InvalidCapacity", func(t *testing.T) {
		// Проверяем что конструктор обрабатывает некорректную емкость
		ht := NewChainingHashTable(0)
		assert.Equal(t, 10, ht.GetCapacity()) // Должна быть установлена емкость по умолчанию

		ht2 := NewChainingHashTable(-5)
		assert.Equal(t, 10, ht2.GetCapacity()) // Должна быть установлена емкость по умолчанию
	})
}

func TestChainingHashTable_Coverage(t *testing.T) {
	t.Run("Remove_NonExistentKey", func(t *testing.T) {
		ht := NewChainingHashTable(10)

		ht.Add(1, 100)
		ht.Add(2, 200)

		initialSize := ht.GetSize()

		// Удаляем несуществующий ключ
		ht.Remove(999)

		assert.Equal(t, initialSize, ht.GetSize())

		// Проверяем что существующие элементы остались
		assert.True(t, ht.Contains(1).first)
		assert.True(t, ht.Contains(2).first)
	})

	t.Run("Remove_FromEmptyTable", func(t *testing.T) {
		ht := NewChainingHashTable(10)

		initialSize := ht.GetSize()

		// Удаляем из пустой таблицы
		ht.Remove(1)
		ht.Remove(2)
		ht.Remove(3)

		assert.Equal(t, initialSize, ht.GetSize())
		assert.Equal(t, 0, ht.GetSize())
	})

	t.Run("Remove_FromChain", func(t *testing.T) {
		ht := NewChainingHashTable(2) // Маленькая емкость для коллизий

		// Добавляем элементы которые будут в одной цепочке
		ht.Add(1, 100) // 1 % 2 = 1
		ht.Add(3, 300) // 3 % 2 = 1 (коллизия)
		ht.Add(5, 500) // 5 % 2 = 1 (коллизия)

		assert.Equal(t, 3, ht.GetSize())

		// Удаляем из середины цепочки
		ht.Remove(3)

		assert.Equal(t, 2, ht.GetSize())
		assert.False(t, ht.Contains(3).first)
		assert.True(t, ht.Contains(1).first)
		assert.True(t, ht.Contains(5).first)

		// Удаляем голову цепочки
		ht.Remove(1)

		assert.Equal(t, 1, ht.GetSize())
		assert.False(t, ht.Contains(1).first)
		assert.True(t, ht.Contains(5).first)

		// Удаляем последний элемент цепочки
		ht.Remove(5)

		assert.Equal(t, 0, ht.GetSize())
		assert.False(t, ht.Contains(5).first)
	})

	t.Run("ToString_EmptyTable", func(t *testing.T) {
		ht := NewChainingHashTable(5)

		str := ht.ToString()

		assert.NotEmpty(t, str)
		assert.Contains(t, str, "[  0]:")
		assert.Contains(t, str, "[  4]:")
		assert.Contains(t, str, "empty")

		// Проверяем что все строки имеют правильный формат
		lines := strings.Split(strings.TrimSpace(str), "\n")
		assert.Equal(t, 5, len(lines)) // 5 строк для емкости 5

		for i, line := range lines {
			assert.Contains(t, line, fmt.Sprintf("[%3d]:", i))
		}
	})

	t.Run("ToString_WithElements", func(t *testing.T) {
		ht := NewChainingHashTable(3)

		ht.Add(1, 100)
		ht.Add(2, 200)
		ht.Add(4, 400) // 4 % 3 = 1 (коллизия с 1)

		str := ht.ToString()

		assert.NotEmpty(t, str)

		// Проверяем наличие элементов
		assert.Contains(t, str, "(1,100)")
		assert.Contains(t, str, "(2,200)")
		assert.Contains(t, str, "(4,400)")

		// Проверяем формат цепочек
		lines := strings.Split(strings.TrimSpace(str), "\n")

		// Находим строку с цепочкой
		for _, line := range lines {
			if strings.Contains(line, "(1,100)") {
				// Проверяем что есть стрелка для цепочки
				assert.Contains(t, line, "->")
			}
		}
	})

	t.Run("ToString_SingleElementPerBucket", func(t *testing.T) {
		ht := NewChainingHashTable(3)

		// Добавляем элементы в разные бакеты
		ht.Add(1, 100) // 1 % 3 = 1
		ht.Add(2, 200) // 2 % 3 = 2
		ht.Add(3, 300) // 3 % 3 = 0

		str := ht.ToString()

		// Проверяем что нет стрелок (одиночные элементы)
		assert.Contains(t, str, "(1,100)")
		assert.Contains(t, str, "(2,200)")
		assert.Contains(t, str, "(3,300)")

		// Проверяем что нет стрелок для одиночных элементов
		lines := strings.Split(strings.TrimSpace(str), "\n")
		for _, line := range lines {
			if strings.Contains(line, "(1,100)") && !strings.Contains(line, "(4,400)") {
				assert.NotContains(t, line, "->")
			}
			if strings.Contains(line, "(2,200)") {
				assert.NotContains(t, line, "->")
			}
			if strings.Contains(line, "(3,300)") {
				assert.NotContains(t, line, "->")
			}
		}
	})

	t.Run("GetChainLengths_AllScenarios", func(t *testing.T) {
		t.Run("EmptyTable", func(t *testing.T) {
			ht := NewChainingHashTable(10)

			min, max, avg := ht.GetChainLengths()

			assert.Equal(t, 0, min)
			assert.Equal(t, 0, max)
			assert.Equal(t, 0.0, avg)
		})

		t.Run("SomeEmptyBuckets", func(t *testing.T) {
			ht := NewChainingHashTable(5)

			// Добавляем элементы только в некоторые бакеты
			ht.Add(1, 100)   // Бакет 1
			ht.Add(2, 200)   // Бакет 2
			ht.Add(6, 600)   // Бакет 1 (коллизия)
			ht.Add(7, 700)   // Бакет 2 (коллизия)
			ht.Add(11, 1100) // Бакет 1 (еще коллизия)

			min, max, avg := ht.GetChainLengths()

			// Должно быть 2 непустых бакета
			// Бакет 1: 3 элемента, Бакет 2: 2 элемента
			assert.Equal(t, 2, min)   // Минимальная цепочка
			assert.Equal(t, 3, max)   // Максимальная цепочка
			assert.Equal(t, 2.5, avg) // Среднее: (3+2)/2 = 2.5
		})

		t.Run("AllBucketsEmptyExceptOne", func(t *testing.T) {
			ht := NewChainingHashTable(10)

			// Все элементы в одном бакете
			for i := 0; i < 5; i++ {
				ht.Add(i*10, i*100) // Все имеют хэш 0
			}

			min, max, avg := ht.GetChainLengths()

			assert.Equal(t, 5, min) // Все в одной цепочке
			assert.Equal(t, 5, max)
			assert.Equal(t, 5.0, avg)
		})

		t.Run("SingleElementPerBucket", func(t *testing.T) {
			ht := NewChainingHashTable(5)

			// Каждый элемент в своем бакете
			for i := 0; i < 5; i++ {
				ht.Add(i, i*10)
			}

			min, max, avg := ht.GetChainLengths()

			assert.Equal(t, 1, min)
			assert.Equal(t, 1, max)
			assert.Equal(t, 1.0, avg)
		})
	})

	t.Run("GetChainLengths_LargeValues", func(t *testing.T) {
		ht := NewChainingHashTable(100)

		// Добавляем много элементов
		for i := 0; i < 1000; i++ {
			ht.Add(i, i*10)
		}

		min, max, avg := ht.GetChainLengths()

		// Проверяем что значения в разумных пределах
		assert.True(t, min >= 0)
		assert.True(t, max > 0)
		assert.True(t, avg > 0.0)

		// Проверяем что min <= avg <= max
		assert.LessOrEqual(t, float64(min), avg)
		assert.LessOrEqual(t, avg, float64(max))
	})
}

// ----
func TestChainingHashTable_EdgeCases(t *testing.T) {
	t.Run("Remove_NonExistentKey", func(t *testing.T) {
		ht := NewChainingHashTable(10)

		// Удаление из пустой таблицы
		initialSize := ht.GetSize()
		ht.Remove(999)
		assert.Equal(t, initialSize, ht.GetSize())

		// Добавляем элемент
		ht.Add(1, 100)

		// Удаление несуществующего ключа
		ht.Remove(999)
		assert.Equal(t, 1, ht.GetSize())
		assert.True(t, ht.Contains(1).first)
	})

	t.Run("Remove_SingleElementChain", func(t *testing.T) {
		ht := NewChainingHashTable(5)
		ht.Add(1, 100)

		assert.Equal(t, 1, ht.GetSize())
		assert.True(t, ht.Contains(1).first)

		// Удаляем единственный элемент
		ht.Remove(1)

		assert.Equal(t, 0, ht.GetSize())
		assert.False(t, ht.Contains(1).first)
	})

	t.Run("Remove_MiddleOfChain_Detailed", func(t *testing.T) {
		// Создаем таблицу с гарантированной коллизией
		ht := NewChainingHashTable(3)

		// Все элементы попадут в бакет 0
		ht.Add(0, 100) // 0 % 3 = 0
		ht.Add(3, 300) // 3 % 3 = 0
		ht.Add(6, 600) // 6 % 3 = 0
		ht.Add(9, 900) // 9 % 3 = 0

		assert.Equal(t, 4, ht.GetSize())

		// Удаляем элемент из середины цепочки
		ht.Remove(3)

		assert.Equal(t, 3, ht.GetSize())

		// Проверяем оставшиеся элементы
		assert.True(t, ht.Contains(0).first)
		assert.False(t, ht.Contains(3).first)
		assert.True(t, ht.Contains(6).first)
		assert.True(t, ht.Contains(9).first)

		// Удаляем голову цепочки
		ht.Remove(0)

		assert.Equal(t, 2, ht.GetSize())
		assert.False(t, ht.Contains(0).first)
		assert.True(t, ht.Contains(6).first)
		assert.True(t, ht.Contains(9).first)
	})
}

func TestChainingHashTable_ToString_Coverage(t *testing.T) {
	t.Run("ToString_EmptyTable", func(t *testing.T) {
		ht := NewChainingHashTable(3)

		str := ht.ToString()

		// Проверяем формат
		lines := strings.Split(strings.TrimSpace(str), "\n")
		assert.Equal(t, 3, len(lines))

		for i, line := range lines {
			assert.Contains(t, line, fmt.Sprintf("[%3d]: empty", i))
		}
	})

	t.Run("ToString_SingleElement", func(t *testing.T) {
		ht := NewChainingHashTable(5)
		ht.Add(42, 100)

		str := ht.ToString()

		assert.Contains(t, str, "(42,100)")
		assert.NotContains(t, str, "->") // Не должно быть стрелки для одного элемента
	})

	t.Run("ToString_MultipleElementsInChain", func(t *testing.T) {
		ht := NewChainingHashTable(2) // Гарантируем коллизии

		ht.Add(1, 100) // 1 % 2 = 1
		ht.Add(3, 300) // 3 % 2 = 1 (коллизия)
		ht.Add(5, 500) // 5 % 2 = 1 (коллизия)

		str := ht.ToString()

		// Проверяем наличие всех элементов
		assert.Contains(t, str, "(1,100)")
		assert.Contains(t, str, "(3,300)")
		assert.Contains(t, str, "(5,500)")

		// Проверяем наличие стрелок для цепочки
		lines := strings.Split(str, "\n")

		for _, line := range lines {
			if strings.Contains(line, "(1,100)") {
				assert.Contains(t, line, "->")
			}
		}
	})

	t.Run("ToString_MixedBuckets", func(t *testing.T) {
		ht := NewChainingHashTable(4)

		// Заполняем разные бакеты
		ht.Add(1, 100) // Бакет 1
		ht.Add(2, 200) // Бакет 2
		ht.Add(5, 500) // Бакет 1 (коллизия)
		ht.Add(6, 600) // Бакет 2 (коллизия)
		ht.Add(9, 900) // Бакет 1 (коллизия)

		str := ht.ToString()

		// Проверяем наличие всех элементов
		for _, expected := range []string{"(1,100)", "(2,200)", "(5,500)", "(6,600)", "(9,900)"} {
			assert.Contains(t, str, expected)
		}
	})
}

func TestChainingHashTable_GetChainLengths_Detailed(t *testing.T) {
	t.Run("GetChainLengths_EmptyTable", func(t *testing.T) {
		ht := NewChainingHashTable(10)

		min, max, avg := ht.GetChainLengths()

		assert.Equal(t, 0, min)
		assert.Equal(t, 0, max)
		assert.Equal(t, 0.0, avg)
	})

	t.Run("GetChainLengths_SingleElementPerBucket", func(t *testing.T) {
		ht := NewChainingHashTable(5)

		// Добавляем по одному элементу в разные бакеты
		ht.Add(0, 100) // Бакет 0
		ht.Add(1, 200) // Бакет 1
		ht.Add(2, 300) // Бакет 2

		min, max, avg := ht.GetChainLengths()

		assert.Equal(t, 1, min)
		assert.Equal(t, 1, max)
		assert.Equal(t, 1.0, avg)
	})

	t.Run("GetChainLengths_UnevenDistribution", func(t *testing.T) {
		ht := NewChainingHashTable(4)

		// Создаем неравномерное распределение
		ht.Add(0, 100) // Бакет 0: 1 элемент
		ht.Add(4, 400) // Бакет 0: 2 элемент (коллизия)
		ht.Add(8, 800) // Бакет 0: 3 элемент (коллизия)

		ht.Add(1, 200) // Бакет 1: 1 элемент
		ht.Add(5, 500) // Бакет 1: 2 элемент (коллизия)

		ht.Add(2, 300) // Бакет 2: 1 элемент

		// Бакет 3: пустой

		min, max, avg := ht.GetChainLengths()

		// 3 непустых бакета: длины 3, 2, 1
		assert.Equal(t, 1, min)   // Минимальная цепочка
		assert.Equal(t, 3, max)   // Максимальная цепочка
		assert.Equal(t, 2.0, avg) // Среднее: (3+2+1)/3 = 2.0
	})

	t.Run("GetChainLengths_AllInOneBucket", func(t *testing.T) {
		ht := NewChainingHashTable(10)

		// Все элементы в одном бакете
		for i := 0; i < 7; i++ {
			ht.Add(i*10, i*100) // Все имеют хэш 0
		}

		min, max, avg := ht.GetChainLengths()

		assert.Equal(t, 7, min)
		assert.Equal(t, 7, max)
		assert.Equal(t, 7.0, avg)
	})

	t.Run("GetChainLengths_LargeTable", func(t *testing.T) {
		ht := NewChainingHashTable(100)

		// Добавляем много элементов
		for i := 0; i < 1000; i++ {
			ht.Add(i, i*10)
		}

		min, max, avg := ht.GetChainLengths()

		// Проверяем разумность значений
		assert.True(t, min >= 0)
		assert.True(t, max > 0)
		assert.True(t, avg > 0.0)

		// Проверяем, что min <= avg <= max
		assert.LessOrEqual(t, float64(min), avg)
		assert.LessOrEqual(t, avg, float64(max))

		// Проверяем вычисления
		totalElements := 0
		nonEmptyBuckets := 0

		for i := 0; i < ht.capacity; i++ {
			length := 0
			current := ht.table[i]

			for current != nil {
				length++
				current = current.next
			}

			if length > 0 {
				totalElements += length
				nonEmptyBuckets++
			}
		}

		// Проверяем согласованность
		if nonEmptyBuckets > 0 {
			expectedAvg := float64(totalElements) / float64(nonEmptyBuckets)
			assert.Equal(t, expectedAvg, avg)
		}
	})
}

func TestChainingHashTable_GetMethods(t *testing.T) {
	t.Run("GetSize_VariousScenarios", func(t *testing.T) {
		ht := NewChainingHashTable(10)

		assert.Equal(t, 0, ht.GetSize())

		ht.Add(1, 100)
		assert.Equal(t, 1, ht.GetSize())

		ht.Add(2, 200)
		ht.Add(3, 300)
		assert.Equal(t, 3, ht.GetSize())

		ht.Remove(2)
		assert.Equal(t, 2, ht.GetSize())

		ht.Clear()
		assert.Equal(t, 0, ht.GetSize())
	})

	t.Run("GetCapacity_EdgeCases", func(t *testing.T) {
		// Различные значения емкости
		testCases := []struct {
			input    int
			expected int
		}{
			{5, 5},
			{10, 10},
			{100, 100},
			{0, 10},  // По умолчанию
			{-5, 10}, // По умолчанию
		}

		for _, tc := range testCases {
			ht := NewChainingHashTable(tc.input)
			assert.Equal(t, tc.expected, ht.GetCapacity())
		}
	})

	t.Run("GetLoadFactor_Detailed", func(t *testing.T) {
		t.Run("EmptyTable", func(t *testing.T) {
			ht := NewChainingHashTable(10)
			assert.Equal(t, 0.0, ht.GetLoadFactor())
		})

		t.Run("HalfFull", func(t *testing.T) {
			ht := NewChainingHashTable(10)

			for i := 0; i < 5; i++ {
				ht.Add(i, i*10)
			}

			assert.Equal(t, 0.5, ht.GetLoadFactor())
		})

		t.Run("Full", func(t *testing.T) {
			ht := NewChainingHashTable(5)

			for i := 0; i < 5; i++ {
				ht.Add(i, i*10)
			}

			assert.Equal(t, 1.0, ht.GetLoadFactor())
		})

		t.Run("OverCapacity", func(t *testing.T) {
			ht := NewChainingHashTable(3)

			for i := 0; i < 10; i++ {
				ht.Add(i, i*10)
			}

			// Коэффициент загрузки может быть больше 1.0
			assert.Greater(t, ht.GetLoadFactor(), 1.0)
		})

		t.Run("ZeroCapacity", func(t *testing.T) {
			ht := NewChainingHashTable(0) // Будет установлена 10
			assert.Equal(t, 0.0, ht.GetLoadFactor())

			ht.Add(1, 100)
			assert.Equal(t, 0.1, ht.GetLoadFactor())
		})
	})

	t.Run("GetLoadFactor_AfterOperations", func(t *testing.T) {
		ht := NewChainingHashTable(10)

		// Начальное состояние
		assert.Equal(t, 0.0, ht.GetLoadFactor())

		// Добавляем элементы
		ht.Add(1, 100)
		ht.Add(2, 200)
		ht.Add(3, 300)
		assert.Equal(t, 0.3, ht.GetLoadFactor())

		// Удаляем элемент
		ht.Remove(2)
		assert.Equal(t, 0.2, ht.GetLoadFactor())

		// Очищаем
		ht.Clear()
		assert.Equal(t, 0.0, ht.GetLoadFactor())

		// Добавляем снова
		for i := 0; i < 15; i++ {
			ht.Add(i, i*10)
		}
		assert.Equal(t, 1.5, ht.GetLoadFactor())
	})
}
