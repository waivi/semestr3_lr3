package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Глобальные хранилища структур данных
var (
	arrays      = make(map[string]*MArray)
	singleLists = make(map[string]*SList)
	doubleLists = make(map[string]*DList)
	stacks      = make(map[string]*Stack)
	queues      = make(map[string]*Queue)
	trees       = make(map[string]*RBTree)
	hashTables  = make(map[string]*ChainingHashTable)
)

func main() {
	fmt.Println("=== Система управления структурами данных на Go ===")
	fmt.Println("Введите команды (help для справки, exit для выхода)")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		if line == "exit" || line == "quit" {
			break
		}

		if line == "" {
			continue
		}

		processCommand(line)
	}

	fmt.Println("Программа завершена.")
}

func processCommand(command string) {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return
	}

	switch parts[0] {
	case "help":
		printHelp()
	case "demo":
		runDemo()
	case "list":
		listStructures()
	case "clear":
		clearAll()
	case "MARRAY":
		handleArrayCommand(parts)
	case "SLIST":
		handleSListCommand(parts)
	case "DLIST":
		handleDListCommand(parts)
	case "STACK":
		handleStackCommand(parts)
	case "QUEUE":
		handleQueueCommand(parts)
	case "TREE":
		handleTreeCommand(parts)
	case "HSET", "HGET":
		handleHashTableCommand(parts)
	case "SAVE":
		handleSaveCommand(parts)
	case "LOAD":
		handleLoadCommand(parts)
	default:
		fmt.Printf("Неизвестная команда: %s\n", parts[0])
		fmt.Println("Введите help для списка команд")
	}
}

func printHelp() {
	fmt.Print(`
Доступные команды:

ОБЩИЕ:
  help                    - Показать эту справку
  demo                    - Запустить демонстрацию
  list                    - Показать все структуры
  clear                   - Очистить все структуры

МАССИВЫ (MARRAY):
  MARRAY <имя> MADDEND <значение>
  MARRAY <имя> MADDINDEX <индекс> <значение>
  MARRAY <имя> MGETINDEX <индекс>
  MARRAY <имя> MPRINT
  MARRAY <имя> MLENGTH

ОДНОСВЯЗНЫЕ СПИСКИ (SLIST):
  SLIST <имя> SLPUSH_HEAD <значение>
  SLIST <имя> SLPUSH_TAIL <значение>
  SLIST <имя> SLSEARCH <значение>
  SLIST <имя> SLPRINT_FORWARD
  SLIST <имя> SLLENGTH

ДВУСВЯЗНЫЕ СПИСКИ (DLIST):
  DLIST <имя> DLPUSH_HEAD <значение>
  DLIST <имя> DLPUSH_TAIL <значение>
  DLIST <имя> DLPRINT_FORWARD
  DLIST <имя> DLPRINT_BACKWARD
  DLIST <имя> DLLENGTH

СТЕКИ (STACK):
  STACK <имя> SPUSH <значение>
  STACK <имя> SPOP
  STACK <имя> SPRINT

ОЧЕРЕДИ (QUEUE):
  QUEUE <имя> QPUSH <значение>
  QUEUE <имя> QPOP
  QUEUE <имя> QPRINT

ДЕРЕВЬЯ (TREE):
  TREE <имя> TINSERT <значение>
  TREE <имя> TSEARCH <значение>
  TREE <имя> TPRINT_INORDER
  TREE <имя> TPRINT_TREE

ХЭШ-ТАБЛИЦЫ (HSET/HGET):
  HSET <имя> HADD <ключ> <значение>
  HSET <имя> HGET <ключ>
  HSET <имя> HPRINT
  HSET <имя> HSIZE

СЕРИАЛИЗАЦИЯ:
  SAVE <тип> <имя> <формат> [файл]
  LOAD <тип> <имя> <формат> <файл>
  
  Типы: ARRAY, SLIST, DLIST, STACK, QUEUE, TREE, HASHTABLE
  Форматы: TEXT, JSON, BINARY

ПРИМЕРЫ:
  MARRAY myarr MADDEND "Hello"
  SLIST mylist SLPUSH_HEAD "World"
  SAVE ARRAY myarr TEXT myarray.txt
  LOAD ARRAY newarr TEXT myarray.txt
`)
}

func runDemo() {
	fmt.Println("\n=== ДЕМОНСТРАЦИЯ РАБОТЫ СТРУКТУР ДАННЫХ ===")

	// 1. Демонстрация массива
	fmt.Println("\n1. Работа с массивом:")
	arr := NewMArray()
	arr.MADDEND("Первый")
	arr.MADDEND("Второй")
	arr.MADDEND("Третий")
	arr.MPRINT()

	// 2. Демонстрация списка
	fmt.Println("\n2. Работа с односвязным списком:")
	list := NewSList()
	list.SLPUSH_HEAD("Голова")
	list.SLPUSH_TAIL("Хвост")
	list.SLPUSH_HEAD("Новая голова")
	list.SLPRINT_FORWARD()

	// 3. Демонстрация стека
	fmt.Println("\n3. Работа со стеком:")
	s := NewStack()
	s.SPUSH("Первый в стек")
	s.SPUSH("Второй в стек")
	s.SPUSH("Третий в стек")
	s.SPRINT()
	if val, err := s.SPOP(); err == nil {
		fmt.Printf("Извлечен из стека: %s\n", val)
	}
	s.SPRINT()

	// 4. Демонстрация очереди
	fmt.Println("\n4. Работа с очередью:")
	q := NewQueue()
	q.QPUSH("Первый в очереди")
	q.QPUSH("Второй в очереди")
	q.QPUSH("Третий в очереди")
	q.QPRINT()
	if val, err := q.QPOP(); err == nil {
		fmt.Printf("Извлечен из очереди: %s\n", val)
	}
	q.QPRINT()

	// 5. Демонстрация дерева
	fmt.Println("\n5. Работа с красно-черным деревом:")
	tr := NewRBTree()
	tr.TINSERT("Apple")
	tr.TINSERT("Banana")
	tr.TINSERT("Cherry")
	tr.TINSERT("Date")
	tr.TPRINT_INORDER()
	tr.TPRINT_TREE()

	// 6. Демонстрация хэш-таблицы
	fmt.Println("\n6. Работа с хэш-таблицей:")
	ht := NewChainingHashTable(10)
	ht.Add(1, 100)
	ht.Add(2, 200)
	ht.Add(11, 1100) // Коллизия с 1
	fmt.Print(ht.ToString())

	// 7. Демонстрация сериализации
	fmt.Println("\n7. Демонстрация сериализации:")
	
	// Сохраняем массив
	testArray := NewMArray()
	testArray.MADDEND("Элемент 1")
	testArray.MADDEND("Элемент 2")
	testArray.MADDEND("Элемент 3")
	
	fmt.Println("Исходный массив:")
	testArray.MPRINT()
	
	// Сохраняем в файл
	err := SaveArrayToText(testArray, "demo_array.txt")
	if err != nil {
		fmt.Printf("Ошибка сохранения: %v\n", err)
	} else {
		fmt.Println("Массив сохранен в demo_array.txt")
	}
	
	// Загружаем обратно
	loadedArray := NewMArray()
	err = LoadArrayFromText(loadedArray, "demo_array.txt")
	if err != nil {
		fmt.Printf("Ошибка загрузки: %v\n", err)
	} else {
		fmt.Println("Загруженный массив:")
		loadedArray.MPRINT()
	}
}

func listStructures() {
	fmt.Println("Доступные структуры данных:")
	
	fmt.Print("  Массивы: ")
	if len(arrays) == 0 {
		fmt.Println("нет")
	} else {
		for name := range arrays {
			fmt.Printf("%s ", name)
		}
		fmt.Println()
	}
	
	fmt.Print("  Односвязные списки: ")
	if len(singleLists) == 0 {
		fmt.Println("нет")
	} else {
		for name := range singleLists {
			fmt.Printf("%s ", name)
		}
		fmt.Println()
	}
	
	fmt.Print("  Двусвязные списки: ")
	if len(doubleLists) == 0 {
		fmt.Println("нет")
	} else {
		for name := range doubleLists {
			fmt.Printf("%s ", name)
		}
		fmt.Println()
	}
	
	fmt.Print("  Стеки: ")
	if len(stacks) == 0 {
		fmt.Println("нет")
	} else {
		for name := range stacks {
			fmt.Printf("%s ", name)
		}
		fmt.Println()
	}
	
	fmt.Print("  Очереди: ")
	if len(queues) == 0 {
		fmt.Println("нет")
	} else {
		for name := range queues {
			fmt.Printf("%s ", name)
		}
		fmt.Println()
	}
	
	fmt.Print("  Деревья: ")
	if len(trees) == 0 {
		fmt.Println("нет")
	} else {
		for name := range trees {
			fmt.Printf("%s ", name)
		}
		fmt.Println()
	}
	
	fmt.Print("  Хэш-таблицы: ")
	if len(hashTables) == 0 {
		fmt.Println("нет")
	} else {
		for name := range hashTables {
			fmt.Printf("%s ", name)
		}
		fmt.Println()
	}
}

func clearAll() {
	arrays = make(map[string]*MArray)
	singleLists = make(map[string]*SList)
	doubleLists = make(map[string]*DList)
	stacks = make(map[string]*Stack)
	queues = make(map[string]*Queue)
	trees = make(map[string]*RBTree)
	hashTables = make(map[string]*ChainingHashTable)
	fmt.Println("Все структуры данных очищены.")
}

// Обработчики команд для каждой структуры данных
func handleArrayCommand(parts []string) {
	if len(parts) < 3 {
		fmt.Println("Использование: MARRAY <имя> <команда> [параметры]")
		return
	}

	name := parts[1]
	command := parts[2]

	// Создаем массив если его нет
	if _, exists := arrays[name]; !exists {
		arrays[name] = NewMArray()
	}

	arr := arrays[name]

	switch command {
	case "MADDEND":
		if len(parts) < 4 {
			fmt.Println("Использование: MARRAY <имя> MADDEND <значение>")
			return
		}
		arr.MADDEND(parts[3])
	
	case "MADDINDEX":
		if len(parts) < 5 {
			fmt.Println("Использование: MARRAY <имя> MADDINDEX <индекс> <значение>")
			return
		}
		index, err := strconv.Atoi(parts[3])
		if err != nil {
			fmt.Printf("Ошибка: индекс должен быть числом: %v\n", err)
			return
		}
		arr.MADDINDEX(index, parts[4])
	
	case "MGETINDEX":
		if len(parts) < 4 {
			fmt.Println("Использование: MARRAY <имя> MGETINDEX <индекс>")
			return
		}
		index, err := strconv.Atoi(parts[3])
		if err != nil {
			fmt.Printf("Ошибка: индекс должен быть числом: %v\n", err)
			return
		}
		value, err := arr.MGETINDEX(index)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		} else {
			fmt.Println(value)
		}
	
	case "MPRINT":
		arr.MPRINT()
	
	case "MLENGTH":
		fmt.Println(arr.MLENGTH())
	
	case "MCLEAR":
		arr.MCLEAR()
	
	default:
		fmt.Printf("Неизвестная команда массива: %s\n", command)
	}
}

func handleSListCommand(parts []string) {
	if len(parts) < 3 {
		fmt.Println("Использование: SLIST <имя> <команда> [параметры]")
		return
	}

	name := parts[1]
	command := parts[2]

	if _, exists := singleLists[name]; !exists {
		singleLists[name] = NewSList()
	}

	list := singleLists[name]

	switch command {
	case "SLPUSH_HEAD":
		if len(parts) < 4 {
			fmt.Println("Использование: SLIST <имя> SLPUSH_HEAD <значение>")
			return
		}
		list.SLPUSH_HEAD(parts[3])
	
	case "SLPUSH_TAIL":
		if len(parts) < 4 {
			fmt.Println("Использование: SLIST <имя> SLPUSH_TAIL <значение>")
			return
		}
		list.SLPUSH_TAIL(parts[3])
	
	case "SLSEARCH":
		if len(parts) < 4 {
			fmt.Println("Использование: SLIST <имя> SLSEARCH <значение>")
			return
		}
		if list.SLSEARCH(parts[3]) {
			fmt.Println("Найдено")
		} else {
			fmt.Println("Не найдено")
		}
	
	case "SLPRINT_FORWARD":
		list.SLPRINT_FORWARD()
	
	case "SLLENGTH":
		fmt.Println(list.SLLENGTH())
	
	default:
		fmt.Printf("Неизвестная команда списка: %s\n", command)
	}
}

func handleDListCommand(parts []string) {
	if len(parts) < 3 {
		fmt.Println("Использование: DLIST <имя> <команда> [параметры]")
		return
	}

	name := parts[1]
	command := parts[2]

	if _, exists := doubleLists[name]; !exists {
		doubleLists[name] = NewDList()
	}

	list := doubleLists[name]

	switch command {
	case "DLPUSH_HEAD":
		if len(parts) < 4 {
			fmt.Println("Использование: DLIST <имя> DLPUSH_HEAD <значение>")
			return
		}
		list.DLPUSH_HEAD(parts[3])
	
	case "DLPUSH_TAIL":
		if len(parts) < 4 {
			fmt.Println("Использование: DLIST <имя> DLPUSH_TAIL <значение>")
			return
		}
		list.DLPUSH_TAIL(parts[3])
	
	case "DLPRINT_FORWARD":
		list.DLPRINT_FORWARD()
	
	case "DLPRINT_BACKWARD":
		list.DLPRINT_BACKWARD()
	
	case "DLLENGTH":
		fmt.Println(list.DLLENGTH())
	
	default:
		fmt.Printf("Неизвестная команда двусвязного списка: %s\n", command)
	}
}

func handleStackCommand(parts []string) {
	if len(parts) < 3 {
		fmt.Println("Использование: STACK <имя> <команда> [параметры]")
		return
	}

	name := parts[1]
	command := parts[2]

	if _, exists := stacks[name]; !exists {
		stacks[name] = NewStack()
	}

	s := stacks[name]

	switch command {
	case "SPUSH":
		if len(parts) < 4 {
			fmt.Println("Использование: STACK <имя> SPUSH <значение>")
			return
		}
		s.SPUSH(parts[3])
	
	case "SPOP":
		val, err := s.SPOP()
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		} else {
			fmt.Println(val)
		}
	
	case "SPRINT":
		s.SPRINT()
	
	default:
		fmt.Printf("Неизвестная команда стека: %s\n", command)
	}
}

func handleQueueCommand(parts []string) {
	if len(parts) < 3 {
		fmt.Println("Использование: QUEUE <имя> <команда> [параметры]")
		return
	}

	name := parts[1]
	command := parts[2]

	if _, exists := queues[name]; !exists {
		queues[name] = NewQueue()
	}

	q := queues[name]

	switch command {
	case "QPUSH":
		if len(parts) < 4 {
			fmt.Println("Использование: QUEUE <имя> QPUSH <значение>")
			return
		}
		q.QPUSH(parts[3])
	
	case "QPOP":
		val, err := q.QPOP()
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		} else {
			fmt.Println(val)
		}
	
	case "QPRINT":
		q.QPRINT()
	
	default:
		fmt.Printf("Неизвестная команда очереди: %s\n", command)
	}
}

func handleTreeCommand(parts []string) {
	if len(parts) < 3 {
		fmt.Println("Использование: TREE <имя> <команда> [параметры]")
		return
	}

	name := parts[1]
	command := parts[2]

	if _, exists := trees[name]; !exists {
		trees[name] = NewRBTree()
	}

	t := trees[name]

	switch command {
	case "TINSERT":
		if len(parts) < 4 {
			fmt.Println("Использование: TREE <имя> TINSERT <значение>")
			return
		}
		t.TINSERT(parts[3])
	
	case "TSEARCH":
		if len(parts) < 4 {
			fmt.Println("Использование: TREE <имя> TSEARCH <значение>")
			return
		}
		if t.TSEARCH(parts[3]) {
			fmt.Println("Найдено")
		} else {
			fmt.Println("Не найдено")
		}
	
	case "TPRINT_INORDER":
		t.TPRINT_INORDER()
	
	case "TPRINT_TREE":
		t.TPRINT_TREE()
	
	default:
		fmt.Printf("Неизвестная команда дерева: %s\n", command)
	}
}

func handleHashTableCommand(parts []string) {
	if len(parts) < 4 {
		fmt.Println("Использование: HSET <имя> <команда> [параметры]")
		fmt.Println("Использование: HGET <имя> <команда> [параметры]")
		return
	}

	name := parts[1]
	command := parts[2]

	if _, exists := hashTables[name]; !exists {
		hashTables[name] = NewChainingHashTable(10)
	}

	ht := hashTables[name]

	switch command {
	case "HADD":
		if len(parts) < 5 {
			fmt.Println("Использование: HSET <имя> HADD <ключ> <значение>")
			return
		}
		key, err := strconv.Atoi(parts[3])
		if err != nil {
			fmt.Println("Ошибка: ключ должен быть целым числом")
			return
		}
		value, err := strconv.Atoi(parts[4])
		if err != nil {
			fmt.Println("Ошибка: значение должно быть целым числом")
			return
		}
		ht.Add(key, value)
		fmt.Printf("Добавлено: ключ=%d, значение=%d\n", key, value)
	
	case "HGET":
		if len(parts) < 4 {
			fmt.Println("Использование: HSET <имя> HGET <ключ>")
			return
		}
		key, err := strconv.Atoi(parts[3])
		if err != nil {
			fmt.Println("Ошибка: ключ должен быть целым числом")
			return
		}
		result := ht.Contains(key)
		if result.first {
			fmt.Printf("Найдено: ключ=%d, значение=%d\n", key, result.second)
		} else {
			fmt.Printf("Ключ %d не найден\n", key)
		}
	
	case "HPRINT":
		fmt.Print(ht.ToString())
	
	case "HSIZE":
		fmt.Printf("Размер: %d\n", ht.GetSize())
		fmt.Printf("Емкость: %d\n", ht.GetCapacity())
		fmt.Printf("Коэффициент загрузки: %.2f\n", ht.GetLoadFactor())
	
	default:
		fmt.Printf("Неизвестная команда хэш-таблицы: %s\n", command)
	}
}

func handleSaveCommand(parts []string) {
	if len(parts) < 5 {
		fmt.Println("Использование: SAVE <тип> <имя> <формат> [файл]")
		fmt.Println("Типы: ARRAY, SLIST, DLIST, STACK, QUEUE, TREE, HASHTABLE")
		fmt.Println("Форматы: TEXT, BINARY")
		return
	}

	dataType := DataType(strings.ToUpper(parts[1]))
	name := parts[2]
	format := Format(strings.ToUpper(parts[3]))
	filename := parts[4]

	// Если имя файла не указано, генерируем по умолчанию
	if filename == "" && len(parts) > 4 {
		filename = parts[4]
	} else if filename == "" {
		filename = name + "_" + strings.ToLower(string(format))
		switch format {
		case "TEXT":
			filename += ".txt"
		case "BINARY":
			filename += ".bin"
		}
	}

	var structure interface{}
	var found bool

	switch dataType {
	case "ARRAY":
		structure, found = arrays[name]
	case "SLIST":
		structure, found = singleLists[name]
	case "DLIST":
		structure, found = doubleLists[name]
	case "STACK":
		structure, found = stacks[name]
	case "QUEUE":
		structure, found = queues[name]
	case "TREE":
		structure, found = trees[name]
	case "HASHTABLE":
		structure, found = hashTables[name]
	default:
		fmt.Printf("Неизвестный тип данных: %s\n", dataType)
		return
	}

	if !found {
		fmt.Printf("Структура '%s' типа %s не найдена\n", name, dataType)
		return
	}

	err := Save(structure, dataType, format, filename)
	if err != nil {
		fmt.Printf("Ошибка сохранения: %v\n", err)
	} else {
		fmt.Printf("Структура '%s' сохранена в файл: %s (формат: %s)\n", name, filename, format)
	}
}

func handleLoadCommand(parts []string) {
	if len(parts) < 6 {
		fmt.Println("Использование: LOAD <тип> <имя> <формат> <файл> <новое_имя>")
		fmt.Println("Типы: ARRAY, SLIST, DLIST, STACK, QUEUE, TREE, HASHTABLE")
		fmt.Println("Форматы: TEXT, JSON, BINARY")
		return
	}

	dataType := DataType(strings.ToUpper(parts[1]))
	//name := parts[2] // Имя для сохранения в программе
	format := Format(strings.ToUpper(parts[3]))
	filename := parts[4]
	newName := parts[5] // Новое имя структуры в программе

	var structure interface{}

	// Создаем новую структуру нужного типа
	switch dataType {
	case "ARRAY":
		structure = NewMArray()
	case "SLIST":
		structure = NewSList()
	case "DLIST":
		structure = NewDList()
	case "STACK":
		structure = NewStack()
	case "QUEUE":
		structure = NewQueue()
	case "TREE":
		structure = NewRBTree()
	case "HASHTABLE":
		structure = NewChainingHashTable(10)
	default:
		fmt.Printf("Неизвестный тип данных: %s\n", dataType)
		return
	}

	err := Load(structure, dataType, format, filename)
	if err != nil {
		fmt.Printf("Ошибка загрузки: %v\n", err)
		return
	}

	// Сохраняем загруженную структуру
	switch dataType {
	case "ARRAY":
		arrays[newName] = structure.(*MArray)
	case "SLIST":
		singleLists[newName] = structure.(*SList)
	case "DLIST":
		doubleLists[newName] = structure.(*DList)
	case "STACK":
		stacks[newName] = structure.(*Stack)
	case "QUEUE":
		queues[newName] = structure.(*Queue)
	case "TREE":
		trees[newName] = structure.(*RBTree)
	case "HASHTABLE":
		hashTables[newName] = structure.(*ChainingHashTable)
	}

	fmt.Printf("Структура загружена из файла %s и сохранена как '%s'\n", filename, newName)
}