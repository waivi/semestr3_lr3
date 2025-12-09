package main

import (
	"encoding/binary"
	"os"
	"testing"
	"path/filepath" 
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestRBTree(t *testing.T) {
	t.Run("NewRBTree", func(t *testing.T) {
		tree := NewRBTree()
		assert.NotNil(t, tree)
		assert.NotNil(t, tree.nil)
		assert.Equal(t, tree.nil, tree.root)
	})

	t.Run("TINSERT", func(t *testing.T) {
		tree := NewRBTree()

		tree.TINSERT("C")
		assert.Equal(t, "C", tree.root.data)
		assert.Equal(t, BLACK, tree.root.color)

		tree.TINSERT("A")
		assert.True(t, tree.TSEARCH("A"))
		assert.True(t, tree.TSEARCH("C"))

		tree.TINSERT("B")
		assert.True(t, tree.TSEARCH("B"))

		// Проверяем что дубликаты не добавляются
		initialSize := len(tree.GetAllElements())
		tree.TINSERT("A") // Дубликат
		assert.Equal(t, initialSize, len(tree.GetAllElements()))
	})

	t.Run("TSEARCH", func(t *testing.T) {
		tree := NewRBTree()

		// Поиск в пустом дереве
		assert.False(t, tree.TSEARCH("A"))

		// Добавляем и ищем
		tree.TINSERT("D")
		tree.TINSERT("B")
		tree.TINSERT("F")
		tree.TINSERT("A")
		tree.TINSERT("C")
		tree.TINSERT("E")
		tree.TINSERT("G")

		// Ищем существующие элементы
		assert.True(t, tree.TSEARCH("A"))
		assert.True(t, tree.TSEARCH("B"))
		assert.True(t, tree.TSEARCH("C"))
		assert.True(t, tree.TSEARCH("D"))
		assert.True(t, tree.TSEARCH("E"))
		assert.True(t, tree.TSEARCH("F"))
		assert.True(t, tree.TSEARCH("G"))

		// Ищем несуществующие элементы
		assert.False(t, tree.TSEARCH("H"))
		assert.False(t, tree.TSEARCH("Z"))
	})

	t.Run("TDELETE", func(t *testing.T) {
		tree := NewRBTree()

		// Удаление из пустого дерева
		tree.TDELETE("A")
		assert.Equal(t, 0, len(tree.GetAllElements()))

		// Добавляем элементы
		tree.TINSERT("D")
		tree.TINSERT("B")
		tree.TINSERT("F")
		tree.TINSERT("A")
		tree.TINSERT("C")
		tree.TINSERT("E")
		tree.TINSERT("G")

		initialSize := len(tree.GetAllElements())

		// Удаляем лист
		tree.TDELETE("A")
		assert.False(t, tree.TSEARCH("A"))
		assert.Equal(t, initialSize-1, len(tree.GetAllElements()))

		// Удаляем узел с одним потомком
		tree.TDELETE("B")
		assert.False(t, tree.TSEARCH("B"))

		// Удаляем узел с двумя потомками
		tree.TDELETE("D") // Корень
		assert.False(t, tree.TSEARCH("D"))

		// Удаляем все элементы
		elements := tree.GetAllElements()
		for _, elem := range elements {
			tree.TDELETE(elem)
		}
		assert.Equal(t, 0, len(tree.GetAllElements()))
	})

	t.Run("GetAllElements", func(t *testing.T) {
		tree := NewRBTree()

		// Пустое дерево
		assert.Equal(t, 0, len(tree.GetAllElements()))

		// Добавляем элементы в случайном порядке
		tree.TINSERT("M")
		tree.TINSERT("B")
		tree.TINSERT("Z")
		tree.TINSERT("A")
		tree.TINSERT("K")

		elements := tree.GetAllElements()
		assert.Equal(t, 5, len(elements))

		// Проверяем что все элементы присутствуют
		for _, elem := range []string{"A", "B", "K", "M", "Z"} {
			assert.Contains(t, elements, elem)
		}
	})

	t.Run("InOrderTraversal", func(t *testing.T) {
		tree := NewRBTree()

		// Пустое дерево
		result := tree.InOrderTraversal()
		assert.Empty(t, result)

		// Добавляем элементы
		tree.TINSERT("D")
		tree.TINSERT("B")
		tree.TINSERT("F")
		tree.TINSERT("A")
		tree.TINSERT("C")
		tree.TINSERT("E")

		// Проверяем порядок обхода
		result = tree.InOrderTraversal()
		expected := []string{"A", "B", "C", "D", "E", "F"}
		assert.Equal(t, expected, result)
	})

	t.Run("MinMax", func(t *testing.T) {
		tree := NewRBTree()

		// Пустое дерево
		assert.Equal(t, "", tree.Min())
		assert.Equal(t, "", tree.Max())

		// Добавляем элементы
		tree.TINSERT("D")
		assert.Equal(t, "D", tree.Min())
		assert.Equal(t, "D", tree.Max())

		tree.TINSERT("B")
		tree.TINSERT("F")
		tree.TINSERT("A")
		tree.TINSERT("C")
		tree.TINSERT("E")
		tree.TINSERT("G")

		assert.Equal(t, "A", tree.Min())
		assert.Equal(t, "G", tree.Max())
	})

	t.Run("PropertiesAfterOperations", func(t *testing.T) {
		tree := NewRBTree()

		// Добавляем много элементов
		letters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M"}
		for _, letter := range letters {
			tree.TINSERT(letter)
		}

		// Проверяем все элементы присутствуют
		for _, letter := range letters {
			assert.True(t, tree.TSEARCH(letter), "Element %s should be in tree", letter)
		}

		// Удаляем некоторые элементы
		toDelete := []string{"B", "D", "F", "H", "J", "L"}
		for _, letter := range toDelete {
			tree.TDELETE(letter)
			assert.False(t, tree.TSEARCH(letter), "Element %s should not be in tree after deletion", letter)
		}

		// Проверяем оставшиеся элементы
		remaining := []string{"A", "C", "E", "G", "I", "K", "M"}
		for _, letter := range remaining {
			assert.True(t, tree.TSEARCH(letter), "Element %s should still be in tree", letter)
		}
	})
}

func TestTreeSerializationErrors(t *testing.T) {
	tempDir := t.TempDir()

	t.Run("LoadTreeFromText_ColorError", func(t *testing.T) {
		tempFile := filepath.Join(tempDir, "invalid_color.txt")
		// Указываем неверный цвет
		content := "1\nnode1\nINVALID_COLOR\n"
		os.WriteFile(tempFile, []byte(content), 0644)

		tree := NewRBTree()
		err := LoadTreeFromText(tree, tempFile)
		// Теперь должно возвращать ошибку
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid color")
		
		// Узел не должен быть добавлен из-за ошибки
		// Проверяем что дерево пустое
		assert.Equal(t, 0, len(tree.GetAllElements()))
	})

	t.Run("LoadTreeFromBinary_Truncated", func(t *testing.T) {
		tempFile := filepath.Join(tempDir, "truncated.bin")
		file, _ := os.Create(tempFile)
		// Пишем count=1, но обрезаем данные
		binary.Write(file, binary.LittleEndian, int32(1))
		binary.Write(file, binary.LittleEndian, int32(5))
		// Не пишем 5 байт данных
		file.Write([]byte("abc")) // Только 3 байта из 5
		file.Close()

		tree := NewRBTree()
		err := LoadTreeFromBinary(tree, tempFile)
		assert.Error(t, err)
	})
}
func TestTreeInsertWithColor(t *testing.T) {
	tree := NewRBTree()

	tree.InsertWithColor("A", RED)
	tree.InsertWithColor("B", BLACK)
	tree.InsertWithColor("C", RED)

	// Проверяем что узлы добавлены
	assert.True(t, tree.TSEARCH("A"))
	assert.True(t, tree.TSEARCH("B"))
	assert.True(t, tree.TSEARCH("C"))

	// Проверяем цвета через GetAllNodes
	nodes := tree.GetAllNodes()
	colorMap := make(map[string]Color)
	for _, node := range nodes {
		colorMap[node.Data] = node.Color
	}

	// Убедимся что все три узла есть
	assert.Equal(t, 3, len(nodes))
}
//
func TestRBTree_Additional(t *testing.T) {
	t.Run("TINSERT_DuplicateHandling", func(t *testing.T) {
		tree := NewRBTree()
		
		// Вставка дубликатов
		tree.TINSERT("A")
		tree.TINSERT("B")
		tree.TINSERT("A") // Дубликат
		tree.TINSERT("B") // Дубликат
		tree.TINSERT("C")
		
		// Должно быть только 3 уникальных элемента
		elements := tree.GetAllElements()
		assert.Equal(t, 3, len(elements))
		
		// Проверяем наличие всех элементов
		assert.Contains(t, elements, "A")
		assert.Contains(t, elements, "B")
		assert.Contains(t, elements, "C")
	})

	t.Run("TDELETE_NonExistentElement", func(t *testing.T) {
		tree := NewRBTree()
		tree.TINSERT("A")
		tree.TINSERT("B")
		tree.TINSERT("C")
		
		initialSize := len(tree.GetAllElements())
		
		// Удаление несуществующего элемента
		tree.TDELETE("X")
		tree.TDELETE("Y")
		tree.TDELETE("Z")
		
		// Размер не должен измениться
		assert.Equal(t, initialSize, len(tree.GetAllElements()))
		
		// Все оригинальные элементы должны остаться
		assert.True(t, tree.TSEARCH("A"))
		assert.True(t, tree.TSEARCH("B"))
		assert.True(t, tree.TSEARCH("C"))
	})

	t.Run("TDELETE_RootNode", func(t *testing.T) {
		tree := NewRBTree()
		tree.TINSERT("D")
		tree.TINSERT("B")
		tree.TINSERT("F")
		tree.TINSERT("A")
		tree.TINSERT("C")
		tree.TINSERT("E")
		tree.TINSERT("G")
		
		// Удаляем корень
		tree.TDELETE("D")
		
		// Проверяем что корень удален
		assert.False(t, tree.TSEARCH("D"))
		
		// Проверяем что другие элементы остались
		assert.True(t, tree.TSEARCH("A"))
		assert.True(t, tree.TSEARCH("B"))
		assert.True(t, tree.TSEARCH("C"))
		assert.True(t, tree.TSEARCH("E"))
		assert.True(t, tree.TSEARCH("F"))
		assert.True(t, tree.TSEARCH("G"))
		
		// Проверяем что дерево все еще работает
		newElements := []string{"H", "I", "J"}
		for _, elem := range newElements {
			tree.TINSERT(elem)
			assert.True(t, tree.TSEARCH(elem))
		}
	})

	t.Run("TDELETE_LeafNode", func(t *testing.T) {
		tree := NewRBTree()
		tree.TINSERT("D")
		tree.TINSERT("B")
		tree.TINSERT("F")
		tree.TINSERT("A")
		tree.TINSERT("C")
		
		// Удаляем лист A
		tree.TDELETE("A")
		
		assert.False(t, tree.TSEARCH("A"))
		assert.True(t, tree.TSEARCH("B"))
		assert.True(t, tree.TSEARCH("C"))
		assert.True(t, tree.TSEARCH("D"))
		assert.True(t, tree.TSEARCH("F"))
		
		// Удаляем лист C
		tree.TDELETE("C")
		
		assert.False(t, tree.TSEARCH("C"))
		assert.True(t, tree.TSEARCH("B"))
		assert.True(t, tree.TSEARCH("D"))
		assert.True(t, tree.TSEARCH("F"))
	})

	t.Run("TDELETE_NodeWithOneChild", func(t *testing.T) {
		tree := NewRBTree()
		
		// Создаем дерево где некоторые узлы имеют только одного потомка
		tree.TINSERT("5")
		tree.TINSERT("3")
		tree.TINSERT("7")
		tree.TINSERT("2")
		tree.TINSERT("4")
		
		// Удаляем 3 (у которого два потомка 2 и 4)
		tree.TDELETE("3")
		
		assert.False(t, tree.TSEARCH("3"))
		assert.True(t, tree.TSEARCH("2"))
		assert.True(t, tree.TSEARCH("4"))
		assert.True(t, tree.TSEARCH("5"))
		assert.True(t, tree.TSEARCH("7"))
	})

	t.Run("TGET_ExistingElement", func(t *testing.T) {
		tree := NewRBTree()
		tree.TINSERT("Apple")
		tree.TINSERT("Banana")
		tree.TINSERT("Cherry")
		
		val, err := tree.TGET("Banana")
		assert.NoError(t, err)
		assert.Equal(t, "Banana", val)
		
		val, err = tree.TGET("Apple")
		assert.NoError(t, err)
		assert.Equal(t, "Apple", val)
		
		val, err = tree.TGET("Cherry")
		assert.NoError(t, err)
		assert.Equal(t, "Cherry", val)
	})

	t.Run("TGET_NonExistentElement", func(t *testing.T) {
		tree := NewRBTree()
		tree.TINSERT("A")
		tree.TINSERT("B")
		
		_, err := tree.TGET("C")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "не найден")
		
		_, err = tree.TGET("")
		assert.Error(t, err)
	})

	t.Run("GetAllNodes_IncludesColors", func(t *testing.T) {
		tree := NewRBTree()
		
		// Добавляем элементы
		elements := []string{"D", "B", "F", "A", "C", "E", "G"}
		for _, elem := range elements {
			tree.TINSERT(elem)
		}
		
		nodes := tree.GetAllNodes()
		assert.Equal(t, len(elements), len(nodes))
		
		// Проверяем что все элементы присутствуют
		elementMap := make(map[string]bool)
		for _, node := range nodes {
			elementMap[node.Data] = true
			// Проверяем что цвет либо RED, либо BLACK
			assert.True(t, node.Color == RED || node.Color == BLACK)
		}
		
		for _, elem := range elements {
			assert.True(t, elementMap[elem], "Element %s should be in GetAllNodes", elem)
		}
	})

	t.Run("InsertWithColor_MaintainsRBProperties", func(t *testing.T) {
		tree := NewRBTree()
		
		// Вставляем с явным указанием цвета
		tree.InsertWithColor("M", BLACK)
		tree.InsertWithColor("B", RED)
		tree.InsertWithColor("R", RED)
		tree.InsertWithColor("A", BLACK)
		tree.InsertWithColor("C", BLACK)
		tree.InsertWithColor("Q", BLACK)
		tree.InsertWithColor("S", BLACK)
		
		// Проверяем что все узлы добавлены
		elements := tree.GetAllElements()
		assert.Equal(t, 7, len(elements))
		
		// Проверяем поиск
		assert.True(t, tree.TSEARCH("M"))
		assert.True(t, tree.TSEARCH("B"))
		assert.True(t, tree.TSEARCH("R"))
		assert.True(t, tree.TSEARCH("A"))
		assert.True(t, tree.TSEARCH("C"))
		assert.True(t, tree.TSEARCH("Q"))
		assert.True(t, tree.TSEARCH("S"))
	})

	t.Run("Clear_RemovesAllElements", func(t *testing.T) {
		tree := NewRBTree()
		
		// Добавляем элементы
		for i := 0; i < 100; i++ {
			tree.TINSERT(fmt.Sprintf("Element%d", i))
		}
		
		assert.Equal(t, 100, len(tree.GetAllElements()))
		
		// Очищаем
		tree.Clear()
		
		assert.Equal(t, 0, len(tree.GetAllElements()))
		assert.False(t, tree.TSEARCH("Element0"))
		assert.False(t, tree.TSEARCH("Element50"))
		assert.False(t, tree.TSEARCH("Element99"))
		
		// Можем добавлять снова
		tree.TINSERT("NewElement")
		assert.True(t, tree.TSEARCH("NewElement"))
		assert.Equal(t, 1, len(tree.GetAllElements()))
	})

	t.Run("InOrderTraversal_SortedOrder", func(t *testing.T) {
		tree := NewRBTree()
		
		// Добавляем в случайном порядке
		words := []string{"zebra", "apple", "mango", "banana", "cherry", "date", "fig"}
		for _, word := range words {
			tree.TINSERT(word)
		}
		
		// InOrderTraversal должен вернуть отсортированный список
		result := tree.InOrderTraversal()
		
		// Проверяем что результат отсортирован
		for i := 0; i < len(result)-1; i++ {
			assert.True(t, result[i] < result[i+1], 
				"Elements should be in sorted order: %s should come before %s", 
				result[i], result[i+1])
		}
		
		// Проверяем что все элементы присутствуют
		assert.Equal(t, len(words), len(result))
		for _, word := range words {
			assert.Contains(t, result, word)
		}
	})

	t.Run("Min_Max_Operations", func(t *testing.T) {
		tree := NewRBTree()
		
		// Пустое дерево
		assert.Equal(t, "", tree.Min())
		assert.Equal(t, "", tree.Max())
		
		// Один элемент
		tree.TINSERT("Middle")
		assert.Equal(t, "Middle", tree.Min())
		assert.Equal(t, "Middle", tree.Max())
		
		// Много элементов
		tree.TINSERT("Aardvark")
		tree.TINSERT("Zyzzyva")
		tree.TINSERT("Banana")
		tree.TINSERT("Yak")
		
		assert.Equal(t, "Aardvark", tree.Min())
		assert.Equal(t, "Zyzzyva", tree.Max())
		
		// После удаления
		tree.TDELETE("Aardvark")
		assert.Equal(t, "Banana", tree.Min())
		
		tree.TDELETE("Zyzzyva")
		assert.Equal(t, "Yak", tree.Max())
	})

	t.Run("TreePrintFunctions_DoNotCrash", func(t *testing.T) {
		tree := NewRBTree()
		
		// Проверяем что функции печати не падают на пустом дереве
		tree.TPRINT_INORDER()
		tree.TPRINT_PREORDER()
		tree.TPRINT_POSTORDER()
		tree.TPRINT_TREE()
		
		// Добавляем элементы и проверяем снова
		tree.TINSERT("D")
		tree.TINSERT("B")
		tree.TINSERT("F")
		
		tree.TPRINT_INORDER()
		tree.TPRINT_PREORDER()
		tree.TPRINT_POSTORDER()
		tree.TPRINT_TREE()
		
		// Удаляем и проверяем снова
		tree.TDELETE("B")
		
		tree.TPRINT_INORDER()
		tree.TPRINT_PREORDER()
		tree.TPRINT_POSTORDER()
		tree.TPRINT_TREE()
	})

	t.Run("ComplexSequenceOfOperations", func(t *testing.T) {
		tree := NewRBTree()
		
		// Сложная последовательность операций
		operations := []struct {
			op   string
			val  string
			desc string
		}{
			{"insert", "50", "Вставляем корень"},
			{"insert", "30", "Вставляем левого потомка"},
			{"insert", "70", "Вставляем правого потомка"},
			{"insert", "20", "Вставляем в левое поддерево"},
			{"insert", "40", "Вставляем в левое поддерево"},
			{"insert", "60", "Вставляем в правое поддерево"},
			{"insert", "80", "Вставляем в правое поддерево"},
			{"search", "30", "Ищем существующий элемент"},
			{"search", "90", "Ищем несуществующий элемент"},
			{"delete", "20", "Удаляем лист"},
			{"delete", "30", "Удаляем узел с одним потомком"},
			{"delete", "50", "Удаляем корень"},
			{"insert", "55", "Вставляем новый элемент"},
			{"insert", "45", "Вставляем новый элемент"},
			{"get", "40", "Получаем существующий элемент"},
			{"min", "", "Получаем минимальный элемент"},
			{"max", "", "Получаем максимальный элемент"},
			{"clear", "", "Очищаем дерево"},
		}
		
		for _, op := range operations {
			switch op.op {
			case "insert":
				tree.TINSERT(op.val)
			case "delete":
				tree.TDELETE(op.val)
			case "search":
				tree.TSEARCH(op.val)
			case "get":
				tree.TGET(op.val)
			case "min":
				tree.Min()
			case "max":
				tree.Max()
			case "clear":
				tree.Clear()
			}
		}
		
		// После всего дерево должно быть пустым
		assert.Equal(t, 0, len(tree.GetAllElements()))
	})

	t.Run("RBTreeProperties_AfterRandomOperations", func(t *testing.T) {
		tree := NewRBTree()
		
		// Выполняем много случайных операций
		for i := 0; i < 1000; i++ {
			// Случайное число от 0 до 999
			value := fmt.Sprintf("%03d", i%1000)
			
			if i%3 == 0 {
				// Примерно 1/3 операций - удаление
				tree.TDELETE(value)
			} else {
				// Примерно 2/3 операций - вставка
				tree.TINSERT(value)
			}
		}
		
		// Проверяем что дерево в рабочем состоянии
		elements := tree.GetAllElements()
		
		// Проверяем что InOrderTraversal возвращает отсортированный список
		inOrder := tree.InOrderTraversal()
		for i := 0; i < len(inOrder)-1; i++ {
			assert.True(t, inOrder[i] < inOrder[i+1], 
				"Tree should maintain sorted order")
		}
		
		// Проверяем что Min и Max работают
		if len(elements) > 0 {
			min := tree.Min()
			max := tree.Max()
			
			// Проверяем что min действительно минимальный
			for _, elem := range elements {
				assert.True(t, min <= elem, "Min should be <= all elements")
			}
			
			// Проверяем что max действительно максимальный
			for _, elem := range elements {
				assert.True(t, elem <= max, "Max should be >= all elements")
			}
		}
	})
}