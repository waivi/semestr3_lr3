package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//  Тестовые вспомогательные функции 

func createTestFiles() (tempDir string, cleanup func()) {
	tempDir, err := os.MkdirTemp("", "serialization_test_*")
	if err != nil {
		panic(err)
	}
	return tempDir, func() {
		os.RemoveAll(tempDir)
	}
}

func createTestArray() *MArray {
	arr := NewMArray()
	arr.MADDEND("first")
	arr.MADDEND("second")
	arr.MADDEND("third")
	return arr
}

func createTestSList() *SList {
	list := NewSList()
	list.SLPUSH_TAIL("item1")
	list.SLPUSH_TAIL("item2")
	list.SLPUSH_TAIL("item3")
	return list
}

func createTestDList() *DList {
	list := NewDList()
	list.DLPUSH_TAIL("ditem1")
	list.DLPUSH_TAIL("ditem2")
	list.DLPUSH_TAIL("ditem3")
	return list
}

func createTestStack() *Stack {
	s := NewStack()
	s.SPUSH("bottom")
	s.SPUSH("middle")
	s.SPUSH("top")
	return s
}

func createTestQueue() *Queue {
	q := NewQueue()
	q.QPUSH("first")
	q.QPUSH("second")
	q.QPUSH("third")
	return q
}

func createTestTree() *RBTree {
	t := NewRBTree()
	t.TINSERT("D")
	t.TINSERT("B")
	t.TINSERT("F")
	t.TINSERT("A")
	t.TINSERT("C")
	t.TINSERT("E")
	t.TINSERT("G")
	return t
}

func createTestHashMap() *ChainingHashTable {
	ht := NewChainingHashTable(10)
	ht.Add(1, 100)
	ht.Add(2, 200)
	ht.Add(3, 300)
	ht.Add(11, 1100) // Коллизия с 1
	return ht
}

//  MArray тесты 

func TestArraySerialization(t *testing.T) {
	tempDir, cleanup := createTestFiles()
	defer cleanup()

	arr := createTestArray()

	// Текстовый формат
	t.Run("SaveArrayToText", func(t *testing.T) {
		filename := tempDir + "/array.txt"
		err := SaveArrayToText(arr, filename)
		require.NoError(t, err)

		// Проверяем что файл создан
		_, err = os.Stat(filename)
		assert.NoError(t, err)
	})

	t.Run("LoadArrayFromText", func(t *testing.T) {
		filename := tempDir + "/array.txt"
		err := SaveArrayToText(arr, filename)
		require.NoError(t, err)

		newArr := NewMArray()
		err = LoadArrayFromText(newArr, filename)
		require.NoError(t, err)

		assert.Equal(t, arr.GetAll(), newArr.GetAll())
	})

	// Бинарный формат
	t.Run("SaveArrayToBinary", func(t *testing.T) {
		filename := tempDir + "/array.bin"
		err := SaveArrayToBinary(arr, filename)
		require.NoError(t, err)

		_, err = os.Stat(filename)
		assert.NoError(t, err)
	})

	t.Run("LoadArrayFromBinary", func(t *testing.T) {
		filename := tempDir + "/array.bin"
		err := SaveArrayToBinary(arr, filename)
		require.NoError(t, err)

		newArr := NewMArray()
		err = LoadArrayFromBinary(newArr, filename)
		require.NoError(t, err)

		assert.Equal(t, arr.GetAll(), newArr.GetAll())
	})

	// Пустой массив
	t.Run("EmptyArray", func(t *testing.T) {
		emptyArr := NewMArray()
		filename := tempDir + "/empty_array.txt"

		err := SaveArrayToText(emptyArr, filename)
		require.NoError(t, err)

		loadedArr := NewMArray()
		err = LoadArrayFromText(loadedArr, filename)
		require.NoError(t, err)

		assert.Equal(t, 0, len(loadedArr.GetAll()))
	})

	// Ошибки
	t.Run("ArrayErrors", func(t *testing.T) {
		// Несуществующий файл для загрузки
		err := LoadArrayFromText(arr, "/nonexistent/file.txt")
		assert.Error(t, err)

		// Некорректный формат файла
		filename := tempDir + "/invalid.txt"
		os.WriteFile(filename, []byte("not a number\n"), 0644)
		err = LoadArrayFromText(arr, filename)
		assert.Error(t, err)
	})
}

//  SList тесты 

func TestSListSerialization(t *testing.T) {
	tempDir, cleanup := createTestFiles()
	defer cleanup()

	list := createTestSList()

	t.Run("SaveSListToText", func(t *testing.T) {
		filename := tempDir + "/slist.txt"
		err := SaveSListToText(list, filename)
		require.NoError(t, err)

		_, err = os.Stat(filename)
		assert.NoError(t, err)
	})

	t.Run("LoadSListFromText", func(t *testing.T) {
		filename := tempDir + "/slist.txt"
		err := SaveSListToText(list, filename)
		require.NoError(t, err)

		newList := NewSList()
		err = LoadSListFromText(newList, filename)
		require.NoError(t, err)

		assert.Equal(t, list.GetAll(), newList.GetAll())
	})

	t.Run("SaveSListToBinary", func(t *testing.T) {
		filename := tempDir + "/slist.bin"
		err := SaveSListToBinary(list, filename)
		require.NoError(t, err)

		_, err = os.Stat(filename)
		assert.NoError(t, err)
	})

	t.Run("LoadSListFromBinary", func(t *testing.T) {
		filename := tempDir + "/slist.bin"
		err := SaveSListToBinary(list, filename)
		require.NoError(t, err)

		newList := NewSList()
		err = LoadSListFromBinary(newList, filename)
		require.NoError(t, err)

		assert.Equal(t, list.GetAll(), newList.GetAll())
	})
}

//  DList тесты 

func TestDListSerialization(t *testing.T) {
	tempDir, cleanup := createTestFiles()
	defer cleanup()

	list := createTestDList()

	t.Run("SaveDListToText", func(t *testing.T) {
		filename := tempDir + "/dlist.txt"
		err := SaveDListToText(list, filename)
		require.NoError(t, err)

		_, err = os.Stat(filename)
		assert.NoError(t, err)
	})

	t.Run("LoadDListFromText", func(t *testing.T) {
		filename := tempDir + "/dlist.txt"
		err := SaveDListToText(list, filename)
		require.NoError(t, err)

		newList := NewDList()
		err = LoadDListFromText(newList, filename)
		require.NoError(t, err)

		assert.Equal(t, list.GetAll(), newList.GetAll())
	})

	t.Run("SaveDListToBinary", func(t *testing.T) {
		filename := tempDir + "/dlist.bin"
		err := SaveDListToBinary(list, filename)
		require.NoError(t, err)

		_, err = os.Stat(filename)
		assert.NoError(t, err)
	})

	t.Run("LoadDListFromBinary", func(t *testing.T) {
		filename := tempDir + "/dlist.bin"
		err := SaveDListToBinary(list, filename)
		require.NoError(t, err)

		newList := NewDList()
		err = LoadDListFromBinary(newList, filename)
		require.NoError(t, err)

		assert.Equal(t, list.GetAll(), newList.GetAll())
	})
}

//  Stack тесты 

func TestStackSerialization(t *testing.T) {
	tempDir, cleanup := createTestFiles()
	defer cleanup()

	stack := createTestStack()

	t.Run("SaveStackToText", func(t *testing.T) {
		filename := tempDir + "/stack.txt"
		err := SaveStackToText(stack, filename)
		require.NoError(t, err)

		_, err = os.Stat(filename)
		assert.NoError(t, err)
	})

	t.Run("LoadStackFromText", func(t *testing.T) {
		filename := tempDir + "/stack.txt"
		err := SaveStackToText(stack, filename)
		require.NoError(t, err)

		newStack := NewStack()
		err = LoadStackFromText(newStack, filename)
		require.NoError(t, err)

		// Проверяем порядок элементов (LIFO)
		assert.Equal(t, stack.GetAll(), newStack.GetAll())

		// Проверяем операцию SPOP
		originalTop, _ := stack.SPOP()
		newTop, _ := newStack.SPOP()
		assert.Equal(t, originalTop, newTop)
	})

	t.Run("SaveStackToBinary", func(t *testing.T) {
		filename := tempDir + "/stack.bin"
		err := SaveStackToBinary(stack, filename)
		require.NoError(t, err)

		_, err = os.Stat(filename)
		assert.NoError(t, err)
	})

	t.Run("LoadStackFromBinary", func(t *testing.T) {
		filename := tempDir + "/stack.bin"
		err := SaveStackToBinary(stack, filename)
		require.NoError(t, err)

		newStack := NewStack()
		err = LoadStackFromBinary(newStack, filename)
		require.NoError(t, err)

		assert.Equal(t, stack.GetAll(), newStack.GetAll())
	})
}

//  Queue тесты 

func TestQueueSerialization(t *testing.T) {
	tempDir, cleanup := createTestFiles()
	defer cleanup()

	queue := createTestQueue()

	t.Run("SaveQueueToText", func(t *testing.T) {
		filename := tempDir + "/queue.txt"
		err := SaveQueueToText(queue, filename)
		require.NoError(t, err)

		_, err = os.Stat(filename)
		assert.NoError(t, err)
	})

	t.Run("LoadQueueFromText", func(t *testing.T) {
		filename := tempDir + "/queue.txt"
		err := SaveQueueToText(queue, filename)
		require.NoError(t, err)

		newQueue := NewQueue()
		err = LoadQueueFromText(newQueue, filename)
		require.NoError(t, err)

		assert.Equal(t, queue.GetAll(), newQueue.GetAll())

		// Проверяем операцию QPOP
		originalFirst, _ := queue.QPOP()
		newFirst, _ := newQueue.QPOP()
		assert.Equal(t, originalFirst, newFirst)
	})

	t.Run("SaveQueueToBinary", func(t *testing.T) {
		filename := tempDir + "/queue.bin"
		err := SaveQueueToBinary(queue, filename)
		require.NoError(t, err)

		_, err = os.Stat(filename)
		assert.NoError(t, err)
	})

	t.Run("LoadQueueFromBinary", func(t *testing.T) {
		filename := tempDir + "/queue.bin"
		err := SaveQueueToBinary(queue, filename)
		require.NoError(t, err)

		newQueue := NewQueue()
		err = LoadQueueFromBinary(newQueue, filename)
		require.NoError(t, err)

		assert.Equal(t, queue.GetAll(), newQueue.GetAll())
	})
}

//  Tree тесты 

func TestTreeSerialization(t *testing.T) {
	tempDir, cleanup := createTestFiles()
	defer cleanup()

	tree := createTestTree()

	t.Run("SaveTreeToText", func(t *testing.T) {
		filename := tempDir + "/tree.txt"
		err := SaveTreeToText(tree, filename)
		require.NoError(t, err)

		_, err = os.Stat(filename)
		assert.NoError(t, err)
	})

	t.Run("LoadTreeFromText", func(t *testing.T) {
		filename := tempDir + "/tree.txt"
		err := SaveTreeToText(tree, filename)
		require.NoError(t, err)

		newTree := NewRBTree()
		err = LoadTreeFromText(newTree, filename)
		require.NoError(t, err)

		// Проверяем что все элементы присутствуют
		originalElements := tree.GetAllElements()
		newElements := newTree.GetAllElements()

		// Сортируем для сравнения (дерево может иметь разную структуру)
		assert.ElementsMatch(t, originalElements, newElements)

		// Проверяем поиск элементов
		for _, elem := range originalElements {
			assert.True(t, newTree.TSEARCH(elem))
		}
	})

	t.Run("SaveTreeToBinary", func(t *testing.T) {
		filename := tempDir + "/tree.bin"
		err := SaveTreeToBinary(tree, filename)
		require.NoError(t, err)

		_, err = os.Stat(filename)
		assert.NoError(t, err)
	})

	t.Run("LoadTreeFromBinary", func(t *testing.T) {
		filename := tempDir + "/tree.bin"
		err := SaveTreeToBinary(tree, filename)
		require.NoError(t, err)

		newTree := NewRBTree()
		err = LoadTreeFromBinary(newTree, filename)
		require.NoError(t, err)

		originalElements := tree.GetAllElements()
		newElements := newTree.GetAllElements()

		assert.ElementsMatch(t, originalElements, newElements)
	})

	// Тест с цветами узлов
	t.Run("TreeColors", func(t *testing.T) {
		tree := NewRBTree()
		tree.InsertWithColor("A", RED)
		tree.InsertWithColor("B", BLACK)
		tree.InsertWithColor("C", RED)

		filename := tempDir + "/tree_colors.txt"
		err := SaveTreeToText(tree, filename)
		require.NoError(t, err)

		newTree := NewRBTree()
		err = LoadTreeFromText(newTree, filename)
		require.NoError(t, err)

		// Проверяем что узлы загрузились
		assert.True(t, newTree.TSEARCH("A"))
		assert.True(t, newTree.TSEARCH("B"))
		assert.True(t, newTree.TSEARCH("C"))
	})

	// Пустое дерево
	t.Run("EmptyTree", func(t *testing.T) {
		emptyTree := NewRBTree()
		filename := tempDir + "/empty_tree.txt"

		err := SaveTreeToText(emptyTree, filename)
		require.NoError(t, err)

		loadedTree := NewRBTree()
		err = LoadTreeFromText(loadedTree, filename)
		require.NoError(t, err)

		assert.Equal(t, 0, len(loadedTree.GetAllElements()))
	})
}

//  HashMap тесты 

func TestHashMapSerialization(t *testing.T) {
	tempDir, cleanup := createTestFiles()
	defer cleanup()

	ht := createTestHashMap()

	t.Run("SaveHashMapToText", func(t *testing.T) {
		filename := tempDir + "/hashmap.txt"
		err := SaveHashMapToText(ht, filename)
		require.NoError(t, err)

		_, err = os.Stat(filename)
		assert.NoError(t, err)
	})

	t.Run("LoadHashMapFromText", func(t *testing.T) {
		filename := tempDir + "/hashmap.txt"
		err := SaveHashMapToText(ht, filename)
		require.NoError(t, err)

		newHt := NewChainingHashTable(10)
		err = LoadHashMapFromText(newHt, filename)
		require.NoError(t, err)

		// Проверяем все пары ключ-значение через Contains
		result1 := ht.Contains(1)
		result2 := newHt.Contains(1)
		assert.Equal(t, result1.first, result2.first)
		assert.Equal(t, result1.second, result2.second)

		result1 = ht.Contains(2)
		result2 = newHt.Contains(2)
		assert.Equal(t, result1.first, result2.first)
		assert.Equal(t, result1.second, result2.second)

		result1 = ht.Contains(3)
		result2 = newHt.Contains(3)
		assert.Equal(t, result1.first, result2.first)
		assert.Equal(t, result1.second, result2.second)

		result1 = ht.Contains(11)
		result2 = newHt.Contains(11)
		assert.Equal(t, result1.first, result2.first)
		assert.Equal(t, result1.second, result2.second)
	})

	t.Run("SaveHashMapToBinary", func(t *testing.T) {
		filename := tempDir + "/hashmap.bin"
		err := SaveHashMapToBinary(ht, filename)
		require.NoError(t, err)

		_, err = os.Stat(filename)
		assert.NoError(t, err)
	})

	t.Run("LoadHashMapFromBinary", func(t *testing.T) {
		filename := tempDir + "/hashmap.bin"
		err := SaveHashMapToBinary(ht, filename)
		require.NoError(t, err)

		newHt := NewChainingHashTable(10)
		err = LoadHashMapFromBinary(newHt, filename)
		require.NoError(t, err)

		result1 := ht.Contains(1)
		result2 := newHt.Contains(1)
		assert.Equal(t, result1.first, result2.first)
		assert.Equal(t, result1.second, result2.second)

		result1 = ht.Contains(2)
		result2 = newHt.Contains(2)
		assert.Equal(t, result1.first, result2.first)
		assert.Equal(t, result1.second, result2.second)

		result1 = ht.Contains(3)
		result2 = newHt.Contains(3)
		assert.Equal(t, result1.first, result2.first)
		assert.Equal(t, result1.second, result2.second)

		result1 = ht.Contains(11)
		result2 = newHt.Contains(11)
		assert.Equal(t, result1.first, result2.first)
		assert.Equal(t, result1.second, result2.second)
	})

	// Пустая хеш-таблица
	t.Run("EmptyHashMap", func(t *testing.T) {
		emptyHt := NewChainingHashTable(10)
		filename := tempDir + "/empty_hashmap.txt"

		err := SaveHashMapToText(emptyHt, filename)
		require.NoError(t, err)

		loadedHt := NewChainingHashTable(10)
		err = LoadHashMapFromText(loadedHt, filename)
		require.NoError(t, err)

		// Проверяем что таблица пуста
		assert.Equal(t, 0, len(loadedHt.GetAllElements()))
	})
}

//  Универсальные функции тесты 

func TestUniversalFunctions(t *testing.T) {
	tempDir, cleanup := createTestFiles()
	defer cleanup()

	t.Run("SaveUniversalArray", func(t *testing.T) {
		arr := createTestArray()
		filename := tempDir + "/universal_array.txt"

		err := Save(arr, ARRAY, TEXT, filename)
		require.NoError(t, err)

		// Проверяем загрузку
		newArr := NewMArray()
		err = Load(newArr, ARRAY, TEXT, filename)
		require.NoError(t, err)

		assert.Equal(t, arr.GetAll(), newArr.GetAll())
	})

	t.Run("SaveUniversalSList", func(t *testing.T) {
		list := createTestSList()
		filename := tempDir + "/universal_slist.bin"

		err := Save(list, SLIST, BINARY, filename)
		require.NoError(t, err)

		newList := NewSList()
		err = Load(newList, SLIST, BINARY, filename)
		require.NoError(t, err)

		assert.Equal(t, list.GetAll(), newList.GetAll())
	})

	t.Run("SaveUniversalTree", func(t *testing.T) {
		tree := createTestTree()
		filename := tempDir + "/universal_tree.txt"

		err := Save(tree, TREE, TEXT, filename)
		require.NoError(t, err)

		newTree := NewRBTree()
		err = Load(newTree, TREE, TEXT, filename)
		require.NoError(t, err)

		assert.ElementsMatch(t, tree.GetAllElements(), newTree.GetAllElements())
	})

	t.Run("SaveUniversalHashMap", func(t *testing.T) {
		ht := createTestHashMap()
		filename := tempDir + "/universal_hashmap.bin"

		err := Save(ht, HASHTABLE, BINARY, filename)
		require.NoError(t, err)

		newHt := NewChainingHashTable(10)
		err = Load(newHt, HASHTABLE, BINARY, filename)
		require.NoError(t, err)

		result1 := ht.Contains(1)
		result2 := newHt.Contains(1)
		assert.Equal(t, result1.first, result2.first)
		assert.Equal(t, result1.second, result2.second)
	})

	// Тесты ошибок
	t.Run("UniversalErrors", func(t *testing.T) {
		arr := createTestArray()

		// Неправильный тип структуры
		err := Save(arr, SLIST, TEXT, tempDir+"/error.txt")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "неверный тип")

		// Неподдерживаемый формат
		err = Save(arr, ARRAY, Format("UNKNOWN"), tempDir+"/error.txt")
		assert.Error(t, err)

		// Неподдерживаемый тип данных
		err = Save(arr, DataType("UNKNOWN"), TEXT, tempDir+"/error.txt")
		assert.Error(t, err)
	})

	t.Run("DirectoryCreation", func(t *testing.T) {
		arr := createTestArray()
		filename := tempDir + "/subdir/deep/nested/array.txt"

		err := Save(arr, ARRAY, TEXT, filename)
		require.NoError(t, err)

		_, err = os.Stat(filename)
		assert.NoError(t, err)
	})
}

//  Интеграционные тесты 

func TestIntegration(t *testing.T) {
	tempDir, cleanup := createTestFiles()
	defer cleanup()

	t.Run("MultipleStructures", func(t *testing.T) {
		// Сохраняем все структуры
		structures := []struct {
			name     string
			obj      interface{}
			dataType DataType
			format   Format
		}{
			{"array", createTestArray(), ARRAY, TEXT},
			{"slist", createTestSList(), SLIST, BINARY},
			{"dlist", createTestDList(), DLIST, TEXT},
			{"stack", createTestStack(), STACK, BINARY},
			{"queue", createTestQueue(), QUEUE, TEXT},
			{"tree", createTestTree(), TREE, BINARY},
			{"hashmap", createTestHashMap(), HASHTABLE, TEXT},
		}

		for _, s := range structures {
			filename := fmt.Sprintf("%s/%s.test", tempDir, s.name)
			err := Save(s.obj, s.dataType, s.format, filename)
			require.NoError(t, err, "Failed to save %s", s.name)
		}

		// Проверяем что все файлы созданы
		for _, s := range structures {
			filename := fmt.Sprintf("%s/%s.test", tempDir, s.name)
			_, err := os.Stat(filename)
			assert.NoError(t, err, "File not created for %s", s.name)
		}
	})

	t.Run("RoundTripAllFormats", func(t *testing.T) {
		// Проверяем туда-обратно для всех комбинаций
		testCases := []struct {
			name     string
			createFn func() interface{}
			dataType DataType
		}{
			{"Array", func() interface{} { return createTestArray() }, ARRAY},
			{"SList", func() interface{} { return createTestSList() }, SLIST},
			{"DList", func() interface{} { return createTestDList() }, DLIST},
			{"Stack", func() interface{} { return createTestStack() }, STACK},
			{"Queue", func() interface{} { return createTestQueue() }, QUEUE},
			{"Tree", func() interface{} { return createTestTree() }, TREE},
			{"HashMap", func() interface{} { return createTestHashMap() }, HASHTABLE},
		}

		for _, tc := range testCases {
			for _, format := range []Format{TEXT, BINARY} {
				t.Run(tc.name+"_"+string(format), func(t *testing.T) {
					original := tc.createFn()
					filename := fmt.Sprintf("%s/%s_%s.test", tempDir, tc.name, format)

					// Сохраняем
					err := Save(original, tc.dataType, format, filename)
					require.NoError(t, err)

					// Создаем новую структуру и загружаем
					var loaded interface{}
					switch tc.dataType {
					case ARRAY:
						loaded = NewMArray()
					case SLIST:
						loaded = NewSList()
					case DLIST:
						loaded = NewDList()
					case STACK:
						loaded = NewStack()
					case QUEUE:
						loaded = NewQueue()
					case TREE:
						loaded = NewRBTree()
					case HASHTABLE:
						loaded = NewChainingHashTable(10)
					}

					err = Load(loaded, tc.dataType, format, filename)
					require.NoError(t, err)

					// Проверяем эквивалентность
					switch tc.dataType {
					case ARRAY:
						assert.Equal(t, original.(*MArray).GetAll(), loaded.(*MArray).GetAll())
					case SLIST:
						assert.Equal(t, original.(*SList).GetAll(), loaded.(*SList).GetAll())
					case DLIST:
						assert.Equal(t, original.(*DList).GetAll(), loaded.(*DList).GetAll())
					case STACK:
						assert.Equal(t, original.(*Stack).GetAll(), loaded.(*Stack).GetAll())
					case QUEUE:
						assert.Equal(t, original.(*Queue).GetAll(), loaded.(*Queue).GetAll())
					case TREE:
						origElements := original.(*RBTree).GetAllElements()
						loadedElements := loaded.(*RBTree).GetAllElements()
						assert.ElementsMatch(t, origElements, loadedElements)
					case HASHTABLE:
						origElements := original.(*ChainingHashTable).GetAllElements()
						loadedElements := loaded.(*ChainingHashTable).GetAllElements()
						assert.ElementsMatch(t, origElements, loadedElements)
					}
				})
			}
		}
	})
}

//  Тесты граничных случаев 

func TestEdgeCases(t *testing.T) {
	tempDir, cleanup := createTestFiles()
	defer cleanup()

	t.Run("EmptyStringInArray", func(t *testing.T) {
		arr := NewMArray()
		arr.MADDEND("")
		arr.MADDEND("normal")
		arr.MADDEND("")

		filename := tempDir + "/empty_strings.txt"
		err := SaveArrayToText(arr, filename)
		require.NoError(t, err)

		newArr := NewMArray()
		err = LoadArrayFromText(newArr, filename)
		require.NoError(t, err)

		assert.Equal(t, arr.GetAll(), newArr.GetAll())
	})

	t.Run("LongStringInTree", func(t *testing.T) {
		tree := NewRBTree()
		longStr := "This is a very long string that should be properly serialized and deserialized without any issues."
		tree.TINSERT(longStr)
		tree.TINSERT("short")

		filename := tempDir + "/long_string.bin"
		err := SaveTreeToBinary(tree, filename)
		require.NoError(t, err)

		newTree := NewRBTree()
		err = LoadTreeFromBinary(newTree, filename)
		require.NoError(t, err)

		assert.True(t, newTree.TSEARCH(longStr))
		assert.True(t, newTree.TSEARCH("short"))
	})

	t.Run("SpecialCharacters", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_TAIL("line1\nline2")
		list.SLPUSH_TAIL("tab\ttab")
		list.SLPUSH_TAIL("unicode: © ® ™")
		list.SLPUSH_TAIL("quotes: \"double\", 'single'")

		filename := tempDir + "/special_chars.txt"
		err := SaveSListToText(list, filename)
		require.NoError(t, err)

		newList := NewSList()
		err = LoadSListFromText(newList, filename)
		require.NoError(t, err)

		assert.Equal(t, list.GetAll(), newList.GetAll())
	})

	t.Run("LargeDataStructure", func(t *testing.T) {
		ht := NewChainingHashTable(100)
		for i := 0; i < 1000; i++ {
			ht.Add(i, i*10)
		}

		filename := tempDir + "/large_hashmap.bin"
		err := SaveHashMapToBinary(ht, filename)
		require.NoError(t, err)

		newHt := NewChainingHashTable(100)
		err = LoadHashMapFromBinary(newHt, filename)
		require.NoError(t, err)

		// Проверяем несколько случайных значений через Contains
		result1 := ht.Contains(0)
		result2 := newHt.Contains(0)
		assert.Equal(t, result1.first, result2.first)
		assert.Equal(t, result1.second, result2.second)

		result1 = ht.Contains(500)
		result2 = newHt.Contains(500)
		assert.Equal(t, result1.first, result2.first)
		assert.Equal(t, result1.second, result2.second)

		result1 = ht.Contains(999)
		result2 = newHt.Contains(999)
		assert.Equal(t, result1.first, result2.first)
		assert.Equal(t, result1.second, result2.second)
	})
}

//  Тесты производительности 

func BenchmarkSerialization(b *testing.B) {
	tempDir, cleanup := createTestFiles()
	defer cleanup()

	// Подготовка данных
	tree := NewRBTree()
	for i := 0; i < 1000; i++ {
		tree.TINSERT(fmt.Sprintf("element_%d", i))
	}

	b.Run("SaveTreeToBinary", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			filename := fmt.Sprintf("%s/benchmark_%d.bin", tempDir, i)
			SaveTreeToBinary(tree, filename)
		}
	})

	b.Run("LoadTreeFromBinary", func(b *testing.B) {
		filename := tempDir + "/benchmark_load.bin"
		SaveTreeToBinary(tree, filename)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			newTree := NewRBTree()
			LoadTreeFromBinary(newTree, filename)
		}
	})
}

// В функции TestEdgeCases_Additional заменить:
func TestEdgeCases_Additional(t *testing.T) {
	tempDir := t.TempDir() // Вместо createTestFiles()

	t.Run("VeryLongString", func(t *testing.T) {
		// Создаем очень длинную строку
		longStr := ""
		for i := 0; i < 10000; i++ {
			longStr += fmt.Sprintf("%04d ", i)
		}

		list := NewSList()
		list.SLPUSH_TAIL(longStr)

		filename := filepath.Join(tempDir, "very_long.txt")
		err := SaveSListToText(list, filename)
		require.NoError(t, err)

		newList := NewSList()
		err = LoadSListFromText(newList, filename)
		require.NoError(t, err)

		assert.Equal(t, list.GetAll(), newList.GetAll())
	})

	t.Run("MixedEscapedCharacters", func(t *testing.T) {
		arr := NewMArray()
		testStrings := []string{
			"line1\nline2",    // Настоящий перевод строки
			"tab\ttab",        // Настоящая табуляция
			"backslash\\here", // Один обратный слеш
			"normal string",
			"", // Пустая строка
		}

		for _, s := range testStrings {
			arr.MADDEND(s)
		}

		filename := filepath.Join(tempDir, "mixed_escaped.txt")
		err := SaveArrayToText(arr, filename)
		require.NoError(t, err)

		newArr := NewMArray()
		err = LoadArrayFromText(newArr, filename)
		require.NoError(t, err)

		assert.Equal(t, arr.GetAll(), newArr.GetAll())
	})

	t.Run("FilePermissions", func(t *testing.T) {
		arr := NewMArray()
		arr.MADDEND("test")

		// Создаем файл без прав на чтение
		filename := filepath.Join(tempDir, "protected.txt")
		err := SaveArrayToText(arr, filename)
		require.NoError(t, err)

		os.Chmod(filename, 0222)       // Только запись
		defer os.Chmod(filename, 0644) // Восстанавливаем

		newArr := NewMArray()
		err = LoadArrayFromText(newArr, filename)
		assert.Error(t, err)
	})

	t.Run("BinaryWithZeroLengthString", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_TAIL("")
		list.SLPUSH_TAIL("normal")
		list.SLPUSH_TAIL("")

		filename := filepath.Join(tempDir, "zero_strings.bin")
		err := SaveSListToBinary(list, filename)
		require.NoError(t, err)

		newList := NewSList()
		err = LoadSListFromBinary(newList, filename)
		require.NoError(t, err)

		assert.Equal(t, list.GetAll(), newList.GetAll())
	})
}

func TestSerializationErrors_Additional(t *testing.T) {
	tempDir := t.TempDir()

	t.Run("SaveArrayToText_PermissionDenied", func(t *testing.T) {
		arr := NewMArray()
		arr.MADDEND("test")

		// Пытаемся сохранить в системную директорию
		err := SaveArrayToText(arr, "/root/test_array.txt")
		if err == nil {
			// Если тест запущен с правами root, это может не сработать
			t.Skip("Test requires non-root permissions")
		}
		assert.Error(t, err)
	})

	t.Run("LoadArrayFromText_FileTooShort", func(t *testing.T) {
		tempFile := filepath.Join(tempDir, "short_array.txt")
		// Указываем size=3, но даем только 1 элемент
		content := "3\nitem1\n"
		os.WriteFile(tempFile, []byte(content), 0644)

		arr := NewMArray()
		err := LoadArrayFromText(arr, tempFile)
		assert.Error(t, err)
	})

	t.Run("SaveSListToBinary_DirectoryInsteadOfFile", func(t *testing.T) {
		list := NewSList()
		list.SLPUSH_TAIL("test")

		// Пытаемся сохранить в директорию
		err := SaveSListToBinary(list, tempDir)
		assert.Error(t, err)
	})

	t.Run("LoadSListFromText_FileTooShort", func(t *testing.T) {
		tempFile := filepath.Join(tempDir, "short_slist.txt")
		// Указываем size=3, но даем только 2 элемента
		content := "3\nitem1\nitem2\n" // Нет третьего элемента
		os.WriteFile(tempFile, []byte(content), 0644)

		list := NewSList()
		err := LoadSListFromText(list, tempFile)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unexpected EOF")
	})

	t.Run("LoadSListFromText_InvalidSize", func(t *testing.T) {
		tempFile := filepath.Join(tempDir, "invalid_size_slist.txt")
		// Некорректный размер
		content := "abc\nitem1\n"
		os.WriteFile(tempFile, []byte(content), 0644)

		list := NewSList()
		err := LoadSListFromText(list, tempFile)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid size format")
	})

	t.Run("SaveDListToText_EmptyList", func(t *testing.T) {
		list := NewDList()

		filename := filepath.Join(tempDir, "empty_dlist.txt")
		err := SaveDListToText(list, filename)
		assert.NoError(t, err)

		// Проверяем что файл создан и содержит только размер
		content, _ := os.ReadFile(filename)
		assert.Equal(t, "0\n", string(content))
	})

	t.Run("LoadStackFromBinary_NegativeCount", func(t *testing.T) {
		tempFile := filepath.Join(tempDir, "negative_count.bin")
		file, _ := os.Create(tempFile)
		// Пишем отрицательное количество элементов
		var count int32 = -5
		binary.Write(file, binary.LittleEndian, count)
		file.Close()

		stack := NewStack()
		err := LoadStackFromBinary(stack, tempFile)
		assert.Error(t, err)
	})

	t.Run("SaveTreeToText_WithNewlinesInData", func(t *testing.T) {
		tree := NewRBTree()
		tree.TINSERT("line1\nline2")
		tree.TINSERT("normal")

		filename := filepath.Join(tempDir, "tree_with_newlines.txt")
		err := SaveTreeToText(tree, filename)
		assert.NoError(t, err)

		// Загружаем обратно
		newTree := NewRBTree()
		err = LoadTreeFromText(newTree, filename)
		assert.NoError(t, err)

		assert.True(t, newTree.TSEARCH("line1\nline2"))
		assert.True(t, newTree.TSEARCH("normal"))
	})

	t.Run("LoadHashMapFromText_NegativeNumbers", func(t *testing.T) {
		tempFile := filepath.Join(tempDir, "negative_hashmap.txt")
		content := "2\n-10 100\n20 -200\n"
		os.WriteFile(tempFile, []byte(content), 0644)

		ht := NewChainingHashTable(10)
		err := LoadHashMapFromText(ht, tempFile)
		assert.NoError(t, err)

		// Проверяем что отрицательные значения загрузились
		result1 := ht.Contains(-10)
		assert.True(t, result1.first)
		assert.Equal(t, 100, result1.second)

		result2 := ht.Contains(20)
		assert.True(t, result2.first)
		assert.Equal(t, -200, result2.second)
	})

	t.Run("SaveUniversal_InvalidDataType", func(t *testing.T) {
		arr := NewMArray()
		arr.MADDEND("test")

		// Неправильный тип данных
		err := Save(arr, DataType("INVALID"), TEXT, tempDir+"/test.txt")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "неподдерживаемый")
	})

	t.Run("LoadUniversal_FileNotFound", func(t *testing.T) {
		arr := NewMArray()

		err := Load(arr, ARRAY, TEXT, "/nonexistent/file.txt")
		assert.Error(t, err)
	})

	t.Run("BinarySerialization_VeryLongString", func(t *testing.T) {
		list := NewSList()

		// Создаем очень длинную строку
		longStr := ""
		for i := 0; i < 10000; i++ {
			longStr += "a"
		}
		list.SLPUSH_TAIL(longStr)

		filename := filepath.Join(tempDir, "long_string.bin")
		err := SaveSListToBinary(list, filename)
		assert.NoError(t, err)

		// Проверяем размер файла
		info, _ := os.Stat(filename)
		assert.True(t, info.Size() > 10000)

		// Загружаем обратно
		newList := NewSList()
		err = LoadSListFromBinary(newList, filename)
		assert.NoError(t, err)

		assert.Equal(t, list.GetAll(), newList.GetAll())
	})

	t.Run("SaveToBinary_ZeroLengthString", func(t *testing.T) {
		arr := NewMArray()
		arr.MADDEND("")
		arr.MADDEND("normal")
		arr.MADDEND("")

		filename := filepath.Join(tempDir, "zero_strings.bin")
		err := SaveArrayToBinary(arr, filename)
		assert.NoError(t, err)

		newArr := NewMArray()
		err = LoadArrayFromBinary(newArr, filename)
		assert.NoError(t, err)

		assert.Equal(t, arr.GetAll(), newArr.GetAll())
	})
}
