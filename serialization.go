package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// MArray

// Текстовый формат
func SaveArrayToText(arr *MArray, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data := arr.GetAll()
	file.WriteString(strconv.Itoa(len(data)) + "\n")
	for _, item := range data {
		// Экранируем переводы строк в строке
		escaped := strings.ReplaceAll(item, "\n", "\\n")
		escaped = strings.ReplaceAll(escaped, "\t", "\\t")
		escaped = strings.ReplaceAll(escaped, "\\", "\\\\")
		file.WriteString(escaped + "\n")
	}

	return nil
}

func LoadArrayFromText(arr *MArray, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Читаем размер
	if !scanner.Scan() {
		return fmt.Errorf("failed to read size: empty file")
	}

	sizeStr := scanner.Text()
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return fmt.Errorf("invalid size format '%s': %w", sizeStr, err)
	}

	if size < 0 {
		return fmt.Errorf("negative size: %d", size)
	}

	items := make([]string, 0, size)
	for i := 0; i < size; i++ {
		if !scanner.Scan() {
			return fmt.Errorf("unexpected EOF: expected %d items, got %d", size, i)
		}

		line := scanner.Text()
		// Восстанавливаем специальные символы
		line = strings.ReplaceAll(line, "\\\\", "\\")
		line = strings.ReplaceAll(line, "\\n", "\n")
		line = strings.ReplaceAll(line, "\\t", "\t")
		items = append(items, line)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	arr.SetAll(items)
	return nil
}

// Бинарный формат
func SaveArrayToBinary(arr *MArray, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data := arr.GetAll()
	count := int32(len(data))

	// Записываем количество элементов
	err = binary.Write(file, binary.LittleEndian, count)
	if err != nil {
		return err
	}

	// Записываем каждый элемент
	for _, str := range data {
		length := int32(len(str))
		err = binary.Write(file, binary.LittleEndian, length)
		if err != nil {
			return err
		}

		_, err = file.Write([]byte(str))
		if err != nil {
			return err
		}
	}

	return nil
}

func LoadArrayFromBinary(arr *MArray, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var count int32
	err = binary.Read(file, binary.LittleEndian, &count)
	if err != nil {
		return err
	}

	items := make([]string, count)
	for i := int32(0); i < count; i++ {
		var length int32
		err = binary.Read(file, binary.LittleEndian, &length)
		if err != nil {
			return err
		}

		buffer := make([]byte, length)
		_, err = file.Read(buffer)
		if err != nil {
			return err
		}

		items[i] = string(buffer)
	}

	arr.SetAll(items)
	return nil
}

// SList

// Текстовый формат
func SaveSListToText(list *SList, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data := list.GetAll()
	file.WriteString(strconv.Itoa(len(data)) + "\n")
	for _, item := range data {
		// Экранируем специальные символы
		escaped := strings.ReplaceAll(item, "\n", "\\n")
		escaped = strings.ReplaceAll(escaped, "\t", "\\t")
		escaped = strings.ReplaceAll(escaped, "\\", "\\\\")
		file.WriteString(escaped + "\n")
	}

	return nil
}

func LoadSListFromText(list *SList, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Читаем размер
	if !scanner.Scan() {
		return fmt.Errorf("failed to read size: empty file")
	}

	sizeStr := scanner.Text()
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return fmt.Errorf("invalid size format '%s': %w", sizeStr, err)
	}

	if size < 0 {
		return fmt.Errorf("negative size: %d", size)
	}

	items := make([]string, 0, size)
	for i := 0; i < size; i++ {
		if !scanner.Scan() {
			return fmt.Errorf("unexpected EOF: expected %d items, got %d", size, i)
		}

		line := scanner.Text()
		// Восстанавливаем специальные символы
		line = strings.ReplaceAll(line, "\\\\", "\\")
		line = strings.ReplaceAll(line, "\\n", "\n")
		line = strings.ReplaceAll(line, "\\t", "\t")
		items = append(items, line)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	list.SetAll(items)
	return nil
}

// Бинарный формат
func SaveSListToBinary(list *SList, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data := list.GetAll()
	count := int32(len(data))

	// Записываем количество элементов
	err = binary.Write(file, binary.LittleEndian, count)
	if err != nil {
		return err
	}

	// Записываем каждый элемент
	for _, str := range data {
		length := int32(len(str))
		err = binary.Write(file, binary.LittleEndian, length)
		if err != nil {
			return err
		}

		_, err = file.Write([]byte(str))
		if err != nil {
			return err
		}
	}

	return nil
}

func LoadSListFromBinary(list *SList, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var count int32
	err = binary.Read(file, binary.LittleEndian, &count)
	if err != nil {
		return fmt.Errorf("failed to read count: %w", err)
	}

	if count < 0 {
		return fmt.Errorf("negative count: %d", count)
	}

	items := make([]string, 0, count)
	for i := int32(0); i < count; i++ {
		var length int32
		err = binary.Read(file, binary.LittleEndian, &length)
		if err != nil {
			return fmt.Errorf("failed to read length for item %d: %w", i, err)
		}

		if length < 0 {
			return fmt.Errorf("negative length for item %d: %d", i, length)
		}

		buffer := make([]byte, length)
		n, err := file.Read(buffer)
		if err != nil {
			return fmt.Errorf("failed to read data for item %d: %w", i, err)
		}
		if int32(n) != length {
			return fmt.Errorf("short read for item %d: expected %d bytes, got %d", i, length, n)
		}

		items = append(items, string(buffer))
	}

	list.SetAll(items)
	return nil
}

// DList

// Текстовый формат
func SaveDListToText(list *DList, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data := list.GetAll()
	file.WriteString(strconv.Itoa(len(data)) + "\n")
	for _, item := range data {
		// Экранируем специальные символы
		escaped := strings.ReplaceAll(item, "\n", "\\n")
		escaped = strings.ReplaceAll(escaped, "\t", "\\t")
		escaped = strings.ReplaceAll(escaped, "\\", "\\\\")
		file.WriteString(escaped + "\n")
	}

	return nil
}

func LoadDListFromText(list *DList, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Читаем размер
	if !scanner.Scan() {
		return fmt.Errorf("failed to read size: empty file")
	}

	sizeStr := scanner.Text()
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return fmt.Errorf("invalid size format '%s': %w", sizeStr, err)
	}

	if size < 0 {
		return fmt.Errorf("negative size: %d", size)
	}

	items := make([]string, 0, size)
	for i := 0; i < size; i++ {
		if !scanner.Scan() {
			return fmt.Errorf("unexpected EOF: expected %d items, got %d", size, i)
		}

		line := scanner.Text()
		// Восстанавливаем специальные символы
		line = strings.ReplaceAll(line, "\\\\", "\\")
		line = strings.ReplaceAll(line, "\\n", "\n")
		line = strings.ReplaceAll(line, "\\t", "\t")
		items = append(items, line)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	list.SetAll(items)
	return nil
}

// Бинарный формат
func SaveDListToBinary(list *DList, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data := list.GetAll()
	count := int32(len(data))

	// Записываем количество элементов
	err = binary.Write(file, binary.LittleEndian, count)
	if err != nil {
		return err
	}

	// Записываем каждый элемент
	for _, str := range data {
		length := int32(len(str))
		err = binary.Write(file, binary.LittleEndian, length)
		if err != nil {
			return err
		}

		_, err = file.Write([]byte(str))
		if err != nil {
			return err
		}
	}

	return nil
}

func LoadDListFromBinary(list *DList, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var count int32
	err = binary.Read(file, binary.LittleEndian, &count)
	if err != nil {
		return err
	}

	items := make([]string, count)
	for i := int32(0); i < count; i++ {
		var length int32
		err = binary.Read(file, binary.LittleEndian, &length)
		if err != nil {
			return err
		}

		buffer := make([]byte, length)
		_, err = file.Read(buffer)
		if err != nil {
			return err
		}

		items[i] = string(buffer)
	}

	list.SetAll(items)
	return nil
}

// Stack

// Текстовый формат
func SaveStackToText(s *Stack, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	items := s.GetAll()
	file.WriteString(strconv.Itoa(len(items)) + "\n")
	for _, item := range items {
		// Экранируем специальные символы
		escaped := strings.ReplaceAll(item, "\n", "\\n")
		escaped = strings.ReplaceAll(escaped, "\t", "\\t")
		escaped = strings.ReplaceAll(escaped, "\\", "\\\\")
		file.WriteString(escaped + "\n")
	}

	return nil
}

func LoadStackFromText(s *Stack, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Читаем размер
	if !scanner.Scan() {
		return fmt.Errorf("failed to read size: empty file")
	}

	sizeStr := scanner.Text()
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return fmt.Errorf("invalid size format '%s': %w", sizeStr, err)
	}

	if size < 0 {
		return fmt.Errorf("negative size: %d", size)
	}

	items := make([]string, 0, size)
	for i := 0; i < size; i++ {
		if !scanner.Scan() {
			return fmt.Errorf("unexpected EOF: expected %d items, got %d", size, i)
		}

		line := scanner.Text()
		// Восстанавливаем специальные символы
		line = strings.ReplaceAll(line, "\\\\", "\\")
		line = strings.ReplaceAll(line, "\\n", "\n")
		line = strings.ReplaceAll(line, "\\t", "\t")
		items = append(items, line)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	s.SetAll(items)
	return nil
}

// Бинарный формат
func SaveStackToBinary(s *Stack, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	items := s.GetAll()
	count := int32(len(items))

	// Записываем количество элементов
	err = binary.Write(file, binary.LittleEndian, count)
	if err != nil {
		return err
	}

	// Записываем каждый элемент в исходном порядке
	for _, str := range items {
		length := int32(len(str))
		err = binary.Write(file, binary.LittleEndian, length)
		if err != nil {
			return err
		}

		_, err = file.Write([]byte(str))
		if err != nil {
			return err
		}
	}

	return nil
}

func LoadStackFromBinary(s *Stack, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var count int32
	err = binary.Read(file, binary.LittleEndian, &count)
	if err != nil {
		return fmt.Errorf("failed to read count: %w", err)
	}

	if count < 0 {
		return fmt.Errorf("negative count: %d", count)
	}

	items := make([]string, 0, count)
	for i := int32(0); i < count; i++ {
		var length int32
		err = binary.Read(file, binary.LittleEndian, &length)
		if err != nil {
			return fmt.Errorf("failed to read length for item %d: %w", i, err)
		}

		if length < 0 {
			return fmt.Errorf("negative length for item %d: %d", i, length)
		}

		buffer := make([]byte, length)
		n, err := file.Read(buffer)
		if err != nil {
			return fmt.Errorf("failed to read data for item %d: %w", i, err)
		}
		if int32(n) != length {
			return fmt.Errorf("short read for item %d: expected %d bytes, got %d", i, length, n)
		}

		items = append(items, string(buffer))
	}

	s.SetAll(items)
	return nil
}

// Queue

// Текстовый формат
func SaveQueueToText(q *Queue, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data := q.GetAll()
	file.WriteString(strconv.Itoa(len(data)) + "\n")
	for _, item := range data {
		// Экранируем специальные символы
		escaped := strings.ReplaceAll(item, "\n", "\\n")
		escaped = strings.ReplaceAll(escaped, "\t", "\\t")
		escaped = strings.ReplaceAll(escaped, "\\", "\\\\")
		file.WriteString(escaped + "\n")
	}

	return nil
}

func LoadQueueFromText(q *Queue, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Читаем размер
	if !scanner.Scan() {
		return fmt.Errorf("failed to read size: empty file")
	}

	sizeStr := scanner.Text()
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return fmt.Errorf("invalid size format '%s': %w", sizeStr, err)
	}

	if size < 0 {
		return fmt.Errorf("negative size: %d", size)
	}

	items := make([]string, 0, size)
	for i := 0; i < size; i++ {
		if !scanner.Scan() {
			return fmt.Errorf("unexpected EOF: expected %d items, got %d", size, i)
		}

		line := scanner.Text()
		// Восстанавливаем специальные символы
		line = strings.ReplaceAll(line, "\\\\", "\\")
		line = strings.ReplaceAll(line, "\\n", "\n")
		line = strings.ReplaceAll(line, "\\t", "\t")
		items = append(items, line)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	q.SetAll(items)
	return nil
}

// Бинарный формат
func SaveQueueToBinary(q *Queue, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data := q.GetAll()
	count := int32(len(data))

	// Записываем количество элементов
	err = binary.Write(file, binary.LittleEndian, count)
	if err != nil {
		return err
	}

	// Записываем каждый элемент
	for _, str := range data {
		length := int32(len(str))
		err = binary.Write(file, binary.LittleEndian, length)
		if err != nil {
			return err
		}

		_, err = file.Write([]byte(str))
		if err != nil {
			return err
		}
	}

	return nil
}

func LoadQueueFromBinary(q *Queue, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var count int32
	err = binary.Read(file, binary.LittleEndian, &count)
	if err != nil {
		return fmt.Errorf("failed to read count: %w", err)
	}

	if count < 0 {
		return fmt.Errorf("negative count: %d", count)
	}

	items := make([]string, 0, count)
	for i := int32(0); i < count; i++ {
		var length int32
		err = binary.Read(file, binary.LittleEndian, &length)
		if err != nil {
			return fmt.Errorf("failed to read length for item %d: %w", i, err)
		}

		if length < 0 {
			return fmt.Errorf("negative length for item %d: %d", i, length)
		}

		buffer := make([]byte, length)
		n, err := file.Read(buffer)
		if err != nil {
			return fmt.Errorf("failed to read data for item %d: %w", i, err)
		}
		if int32(n) != length {
			return fmt.Errorf("short read for item %d: expected %d bytes, got %d", i, length, n)
		}

		items = append(items, string(buffer))
	}

	q.SetAll(items)
	return nil
}

// Tree

// Текстовый формат
func SaveTreeToText(t *RBTree, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	nodes := t.GetAllNodes()
	file.WriteString(strconv.Itoa(len(nodes)) + "\n")
	for _, node := range nodes {
		colorStr := "RED"
		if node.Color == BLACK {
			colorStr = "BLACK"
		}
		// Экранируем данные узла
		data := strings.ReplaceAll(node.Data, "\n", "\\n")
		data = strings.ReplaceAll(data, "\t", "\\t")
		data = strings.ReplaceAll(data, "\\", "\\\\")
		file.WriteString(data + "\n")
		file.WriteString(colorStr + "\n")
	}

	return nil
}

func LoadTreeFromText(t *RBTree, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Читаем размер
	if !scanner.Scan() {
		return fmt.Errorf("failed to read size: empty file")
	}

	sizeStr := scanner.Text()
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return fmt.Errorf("invalid size format '%s': %w", sizeStr, err)
	}

	if size < 0 {
		return fmt.Errorf("negative size: %d", size)
	}

	t.Clear()
	for i := 0; i < size; i++ {
		// Читаем данные
		if !scanner.Scan() {
			return fmt.Errorf("unexpected EOF: expected data for node %d", i)
		}
		data := scanner.Text()
		// Восстанавливаем специальные символы
		data = strings.ReplaceAll(data, "\\\\", "\\")
		data = strings.ReplaceAll(data, "\\n", "\n")
		data = strings.ReplaceAll(data, "\\t", "\t")

		// Читаем цвет
		if !scanner.Scan() {
			return fmt.Errorf("unexpected EOF: expected color for node %d", i)
		}
		colorStr := scanner.Text()

		var color Color
		if strings.ToUpper(colorStr) == "RED" {
			color = RED
		} else if strings.ToUpper(colorStr) == "BLACK" {
			color = BLACK
		} else {
			return fmt.Errorf("invalid color '%s' for node %d", colorStr, i)
		}

		t.InsertWithColor(data, color)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	return nil
}

// Бинарный формат
func SaveTreeToBinary(t *RBTree, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	nodes := t.GetAllNodes()
	count := int32(len(nodes))

	// Записываем количество узлов
	err = binary.Write(file, binary.LittleEndian, count)
	if err != nil {
		return err
	}

	// Записываем каждый узел
	for _, node := range nodes {
		// Записываем данные
		dataLen := int32(len(node.Data))
		err = binary.Write(file, binary.LittleEndian, dataLen)
		if err != nil {
			return err
		}
		_, err = file.Write([]byte(node.Data))
		if err != nil {
			return err
		}

		// Записываем цвет (1 байт: 0 = RED, 1 = BLACK)
		var colorByte byte
		if node.Color == RED {
			colorByte = 0
		} else {
			colorByte = 1
		}
		err = binary.Write(file, binary.LittleEndian, colorByte)
		if err != nil {
			return err
		}
	}

	return nil
}

func LoadTreeFromBinary(t *RBTree, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	t.Clear()

	// Читаем количество узлов
	var count int32
	err = binary.Read(file, binary.LittleEndian, &count)
	if err != nil {
		return fmt.Errorf("failed to read count: %w", err)
	}

	if count < 0 {
		return fmt.Errorf("negative count: %d", count)
	}

	// Читаем каждый узел
	for i := int32(0); i < count; i++ {
		// Читаем данные
		var dataLen int32
		err = binary.Read(file, binary.LittleEndian, &dataLen)
		if err != nil {
			return fmt.Errorf("failed to read data length for node %d: %w", i, err)
		}

		if dataLen < 0 {
			return fmt.Errorf("negative data length for node %d: %d", i, dataLen)
		}

		dataBytes := make([]byte, dataLen)
		n, err := file.Read(dataBytes)
		if err != nil {
			return fmt.Errorf("failed to read data for node %d: %w", i, err)
		}
		if int32(n) != dataLen {
			return fmt.Errorf("short read for node %d: expected %d bytes, got %d", i, dataLen, n)
		}
		data := string(dataBytes)

		// Читаем цвет
		var colorByte byte
		err = binary.Read(file, binary.LittleEndian, &colorByte)
		if err != nil {
			return fmt.Errorf("failed to read color for node %d: %w", i, err)
		}

		var color Color
		if colorByte == 0 {
			color = RED
		} else if colorByte == 1 {
			color = BLACK
		} else {
			return fmt.Errorf("invalid color byte %d for node %d", colorByte, i)
		}

		// Вставляем узел с указанным цветом
		t.InsertWithColor(data, color)
	}

	return nil
}

// Hashmap

// Текстовый формат
func SaveHashMapToText(ht *ChainingHashTable, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	elements := ht.GetAllElements()
	file.WriteString(strconv.Itoa(len(elements)) + "\n")
	for _, elem := range elements {
		file.WriteString(fmt.Sprintf("%d %d\n", elem[0], elem[1]))
	}

	return nil
}

func LoadHashMapFromText(ht *ChainingHashTable, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// Читаем размер
	sizeLine, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	size, err := strconv.Atoi(strings.TrimSpace(sizeLine))
	if err != nil {
		return err
	}

	ht.Clear()
	for i := 0; i < size; i++ {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}

		// Парсим ключ и значение
		var key, value int
		_, err = fmt.Sscanf(strings.TrimSpace(line), "%d %d", &key, &value)
		if err != nil {
			return err
		}
		ht.Add(key, value)
	}

	return nil
}

// Бинарный формат

func SaveHashMapToBinary(ht *ChainingHashTable, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	elements := ht.GetAllElements()
	count := int32(len(elements))

	// Записываем количество элементов
	err = binary.Write(file, binary.LittleEndian, count)
	if err != nil {
		return err
	}

	// Записываем каждую пару ключ-значение
	for _, elem := range elements {
		key := int32(elem[0])
		value := int32(elem[1])

		err = binary.Write(file, binary.LittleEndian, key)
		if err != nil {
			return err
		}

		err = binary.Write(file, binary.LittleEndian, value)
		if err != nil {
			return err
		}
	}

	return file.Sync() // Принудительная запись на диск
}

func LoadHashMapFromBinary(ht *ChainingHashTable, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Читаем количество элементов
	var count int32
	err = binary.Read(file, binary.LittleEndian, &count)
	if err != nil {
		return fmt.Errorf("ошибка чтения количества элементов: %v", err)
	}

	if count < 0 {
		return fmt.Errorf("некорректное количество элементов: %d", count)
	}

	// Проверяем файл на EOF перед чтением элементов
	ht.Clear()
	for i := int32(0); i < count; i++ {
		var key, value int32

		err = binary.Read(file, binary.LittleEndian, &key)
		if err != nil {
			if err == io.EOF {
				return fmt.Errorf("неожиданный конец файла: ожидалось %d элементов, но найдено только %d", count, i)
			}
			return fmt.Errorf("ошибка чтения ключа #%d: %v", i+1, err)
		}

		err = binary.Read(file, binary.LittleEndian, &value)
		if err != nil {
			if err == io.EOF {
				return fmt.Errorf("неожиданный конец файла: неполная пара ключ-значение #%d", i+1)
			}
			return fmt.Errorf("ошибка чтения значения #%d: %v", i+1, err)
		}

		ht.Add(int(key), int(value))
	}

	return nil
}

// Универсальные функции

type DataType string

const (
	ARRAY     DataType = "ARRAY"
	SLIST     DataType = "SLIST"
	DLIST     DataType = "DLIST"
	STACK     DataType = "STACK"
	QUEUE     DataType = "QUEUE"
	TREE      DataType = "TREE"
	HASHTABLE DataType = "HASHTABLE"
)

type Format string

const (
	TEXT   Format = "TEXT"
	BINARY Format = "BINARY"
)

// Save универсальная функция сохранения
func Save(structure interface{}, dataType DataType, format Format, filename string) error {
	// Создаем директорию если не существует
	dir := filepath.Dir(filename)
	if dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	switch dataType {
	case ARRAY:
		arr, ok := structure.(*MArray)
		if !ok {
			return fmt.Errorf("неверный тип для ARRAY")
		}
		switch format {
		case TEXT:
			return SaveArrayToText(arr, filename)
		case BINARY:
			return SaveArrayToBinary(arr, filename)
		}

	case SLIST:
		list, ok := structure.(*SList)
		if !ok {
			return fmt.Errorf("неверный тип для SLIST")
		}
		switch format {
		case TEXT:
			return SaveSListToText(list, filename)
		case BINARY:
			return SaveSListToBinary(list, filename)
		}

	case DLIST:
		list, ok := structure.(*DList)
		if !ok {
			return fmt.Errorf("неверный тип для DLIST")
		}
		switch format {
		case TEXT:
			return SaveDListToText(list, filename)
		case BINARY:
			return SaveDListToBinary(list, filename)
		}

	case STACK:
		s, ok := structure.(*Stack)
		if !ok {
			return fmt.Errorf("неверный тип для STACK")
		}
		switch format {
		case TEXT:
			return SaveStackToText(s, filename)
		case BINARY:
			return SaveStackToBinary(s, filename)
		}

	case QUEUE:
		q, ok := structure.(*Queue)
		if !ok {
			return fmt.Errorf("неверный тип для QUEUE")
		}
		switch format {
		case TEXT:
			return SaveQueueToText(q, filename)
		case BINARY:
			return SaveQueueToBinary(q, filename)
		}

	case TREE:
		t, ok := structure.(*RBTree)
		if !ok {
			return fmt.Errorf("неверный тип для TREE")
		}
		switch format {
		case TEXT:
			return SaveTreeToText(t, filename)
		case BINARY:
			return SaveTreeToBinary(t, filename)
		}

	case HASHTABLE:
		ht, ok := structure.(*ChainingHashTable)
		if !ok {
			return fmt.Errorf("неверный тип для HASHTABLE")
		}
		switch format {
		case TEXT:
			return SaveHashMapToText(ht, filename)
		case BINARY:
			return SaveHashMapToBinary(ht, filename)
		}
	}

	return fmt.Errorf("неподдерживаемый формат или тип данных")
}

// Load универсальная функция загрузки
func Load(structure interface{}, dataType DataType, format Format, filename string) error {
	switch dataType {
	case ARRAY:
		arr, ok := structure.(*MArray)
		if !ok {
			return fmt.Errorf("неверный тип для ARRAY")
		}
		switch format {
		case TEXT:
			return LoadArrayFromText(arr, filename)
		case BINARY:
			return LoadArrayFromBinary(arr, filename)
		}

	case SLIST:
		list, ok := structure.(*SList)
		if !ok {
			return fmt.Errorf("неверный тип для SLIST")
		}
		switch format {
		case TEXT:
			return LoadSListFromText(list, filename)
		case BINARY:
			return LoadSListFromBinary(list, filename)
		}

	case DLIST:
		list, ok := structure.(*DList)
		if !ok {
			return fmt.Errorf("неверный тип для DLIST")
		}
		switch format {
		case TEXT:
			return LoadDListFromText(list, filename)
		case BINARY:
			return LoadDListFromBinary(list, filename)
		}

	case STACK:
		s, ok := structure.(*Stack)
		if !ok {
			return fmt.Errorf("неверный тип для STACK")
		}
		switch format {
		case TEXT:
			return LoadStackFromText(s, filename)
		case BINARY:
			return LoadStackFromBinary(s, filename)
		}

	case QUEUE:
		q, ok := structure.(*Queue)
		if !ok {
			return fmt.Errorf("неверный тип для QUEUE")
		}
		switch format {
		case TEXT:
			return LoadQueueFromText(q, filename)
		case BINARY:
			return LoadQueueFromBinary(q, filename)
		}

	case TREE:
		t, ok := structure.(*RBTree)
		if !ok {
			return fmt.Errorf("неверный тип для TREE")
		}
		switch format {
		case TEXT:
			return LoadTreeFromText(t, filename)
		case BINARY:
			return LoadTreeFromBinary(t, filename)
		}

	case HASHTABLE:
		ht, ok := structure.(*ChainingHashTable)
		if !ok {
			return fmt.Errorf("неверный тип для HASHTABLE")
		}
		switch format {
		case TEXT:
			return LoadHashMapFromText(ht, filename)
		case BINARY:
			return LoadHashMapFromBinary(ht, filename)
		}
	}

	return fmt.Errorf("неподдерживаемый формат или тип данных")
}
