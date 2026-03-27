#ifndef SERIALIZATION_H
#define SERIALIZATION_H

#include <fstream>
#include <string>
#include <stdexcept>
#include <vector>
#include <utility>
#include <stack>
#include <sstream>
#include <queue>

#include "array.h" 
#include "SList.h"
#include "DList.h" 
#include "Hash_map.h"
#include "Tree.h" 
#include "Stack.h"       
#include "Queue.h"       

// MArray

// Текстовый формат
inline void saveToText(const MArray& arr, const std::string& filename) {
    std::ofstream file(filename);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    size_t size = arr.MLENGTH();
    file << size << "\n";
    for (size_t i = 0; i < size; ++i) {
        file << arr.MGETINDEX(i) << "\n";
    }
    file.close();
}

inline void loadFromText(MArray& arr, const std::string& filename) {
    std::ifstream file(filename);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    arr.MCLEAR();
    size_t size;
    file >> size;
    file.ignore();
    
    std::string line;
    for (size_t i = 0; i < size; ++i) {
        std::getline(file, line);
        arr.MADDEND(line);
    }
    file.close();
}

// Бинарный формат
inline void saveToBinary(const MArray& arr, const std::string& filename) {
    std::ofstream file(filename, std::ios::binary);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    size_t size = arr.MLENGTH();
    file.write(reinterpret_cast<const char*>(&size), sizeof(size));
    
    for (size_t i = 0; i < size; ++i) {
        std::string str = arr.MGETINDEX(i);
        size_t len = str.length();
        file.write(reinterpret_cast<const char*>(&len), sizeof(len));
        file.write(str.c_str(), len);
    }
    file.close();
}

inline void loadFromBinary(MArray& arr, const std::string& filename) {
    std::ifstream file(filename, std::ios::binary);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    arr.MCLEAR();
    size_t size;
    file.read(reinterpret_cast<char*>(&size), sizeof(size));
    
    for (size_t i = 0; i < size; ++i) {
        size_t len;
        file.read(reinterpret_cast<char*>(&len), sizeof(len));
        std::string str(len, '\0');
        file.read(&str[0], len);
        arr.MADDEND(str);
    }
    file.close();
}

// SList

// Текстовый формат
inline void saveToText(const SList& list, const std::string& filename) {
    std::ofstream file(filename);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    size_t size = list.SLLENGTH();
    file << size << "\n";
    
    for (size_t i = 0; i < size; ++i) {
        file << list.SLGET(i) << "\n";
    }
    file.close();
}

inline void loadFromText(SList& list, const std::string& filename) {
    std::ifstream file(filename);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    list.SLCLEAR();
    size_t size;
    file >> size;
    file.ignore();
    
    std::string line;
    for (size_t i = 0; i < size; ++i) {
        std::getline(file, line);
        list.SLPUSH_TAIL(line);
    }
    file.close();
}

// Бинарный формат
inline void saveToBinary(const SList& list, const std::string& filename) {
    std::ofstream file(filename, std::ios::binary);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    size_t size = list.SLLENGTH();
    file.write(reinterpret_cast<const char*>(&size), sizeof(size));
    
    for (size_t i = 0; i < size; ++i) {
        std::string str = list.SLGET(i);
        size_t len = str.length();
        file.write(reinterpret_cast<const char*>(&len), sizeof(len));
        file.write(str.c_str(), len);
    }
    file.close();
}

inline void loadFromBinary(SList& list, const std::string& filename) {
    std::ifstream file(filename, std::ios::binary);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    list.SLCLEAR();
    size_t size;
    file.read(reinterpret_cast<char*>(&size), sizeof(size));
    
    for (size_t i = 0; i < size; ++i) {
        size_t len;
        file.read(reinterpret_cast<char*>(&len), sizeof(len));
        std::string str(len, '\0');
        file.read(&str[0], len);
        list.SLPUSH_TAIL(str);
    }
    file.close();
}

//DList

// Текстовый формат
inline void saveToText(const DList& list, const std::string& filename) {
    std::ofstream file(filename);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    size_t size = list.DLLENGTH();
    file << size << "\n";
    
    for (size_t i = 0; i < size; ++i) {
        file << list.DLGET(i) << "\n";
    }
    file.close();
}

inline void loadFromText(DList& list, const std::string& filename) {
    std::ifstream file(filename);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    list.DLCLEAR();
    size_t size;
    file >> size;
    file.ignore();
    
    std::string line;
    for (size_t i = 0; i < size; ++i) {
        std::getline(file, line);
        list.DLPUSH_TAIL(line);
    }
    file.close();
}

// Бинарный формат
inline void saveToBinary(const DList& list, const std::string& filename) {
    std::ofstream file(filename, std::ios::binary);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    size_t size = list.DLLENGTH();
    file.write(reinterpret_cast<const char*>(&size), sizeof(size));
    
    for (size_t i = 0; i < size; ++i) {
        std::string str = list.DLGET(i);
        size_t len = str.length();
        file.write(reinterpret_cast<const char*>(&len), sizeof(len));
        file.write(str.c_str(), len);
    }
    file.close();
}

inline void loadFromBinary(DList& list, const std::string& filename) {
    std::ifstream file(filename, std::ios::binary);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    list.DLCLEAR();
    size_t size;
    file.read(reinterpret_cast<char*>(&size), sizeof(size));
    
    for (size_t i = 0; i < size; ++i) {
        size_t len;
        file.read(reinterpret_cast<char*>(&len), sizeof(len));
        std::string str(len, '\0');
        file.read(&str[0], len);
        list.DLPUSH_TAIL(str);
    }
    file.close();
}

// Stack

// Текстовый формат
inline void saveToText(Stack& stack, const std::string& filename) {
    std::ofstream file(filename);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    // Создаем временный стек для сохранения порядка
    Stack temp;
    std::vector<std::string> items;
    
    // Переливаем элементы во временный стек
    while (!stack.isEmpty()) {
        std::string val = stack.SPOP();
        temp.SPUSH(val);
        items.push_back(val);
    }
    
    // Восстанавливаем оригинальный стек
    while (!temp.isEmpty()) {
        stack.SPUSH(temp.SPOP());
    }
    
    // Записываем в файл (в обратном порядке, чтобы при загрузке получить правильный порядок)
    file << items.size() << "\n";
    for (auto it = items.rbegin(); it != items.rend(); ++it) {
        file << *it << "\n";
    }
    
    file.close();
}

inline void loadFromText(Stack& stack, const std::string& filename) {
    std::ifstream file(filename);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    stack.SCLEAR();
    int size;
    file >> size;
    file.ignore();
    
    std::vector<std::string> items(size);
    for (int i = 0; i < size; ++i) {
        std::getline(file, items[i]);
    }
    
    // Заполняем стек (первый элемент из файла должен быть наверху)
    for (const auto& item : items) {
        stack.SPUSH(item);
    }
    
    file.close();
}

// Бинарный формат
inline void saveToBinary(Stack& stack, const std::string& filename) {
    std::ofstream file(filename, std::ios::binary);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    // Создаем временный стек для сохранения порядка
    Stack temp;
    std::vector<std::string> items;
    
    // Переливаем элементы во временный стек
    while (!stack.isEmpty()) {
        std::string val = stack.SPOP();
        temp.SPUSH(val);
        items.push_back(val);
    }
    
    // Восстанавливаем оригинальный стек
    while (!temp.isEmpty()) {
        stack.SPUSH(temp.SPOP());
    }
    
    // Записываем в файл
    int count = items.size();
    file.write(reinterpret_cast<const char*>(&count), sizeof(count));
    
    for (auto it = items.rbegin(); it != items.rend(); ++it) {
        size_t len = it->length();
        file.write(reinterpret_cast<const char*>(&len), sizeof(len));
        file.write(it->c_str(), len);
    }
    
    file.close();
}

inline void loadFromBinary(Stack& stack, const std::string& filename) {
    std::ifstream file(filename, std::ios::binary);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    stack.SCLEAR();
    int count;
    file.read(reinterpret_cast<char*>(&count), sizeof(count));
    
    std::vector<std::string> items(count);
    for (int i = 0; i < count; ++i) {
        size_t len;
        file.read(reinterpret_cast<char*>(&len), sizeof(len));
        items[i].resize(len);
        file.read(&items[i][0], len);
    }
    
    // Заполняем стек
    for (const auto& item : items) {
        stack.SPUSH(item);
    }
    
    file.close();
}

// Queue

// Текстовый формат
inline void saveToText(Queue& queue, const std::string& filename) {
    std::ofstream file(filename);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    // Создаем временную очередь для сохранения порядка
    Queue temp;
    std::vector<std::string> items;
    
    // Переливаем элементы во временную очередь
    while (!queue.isEmpty()) {
        std::string val = queue.QPOP();
        temp.QPUSH(val);
        items.push_back(val);
    }
    
    // Восстанавливаем оригинальную очередь
    while (!temp.isEmpty()) {
        queue.QPUSH(temp.QPOP());
    }
    
    // Записываем в файл
    file << items.size() << "\n";
    for (const auto& item : items) {
        file << item << "\n";
    }
    
    file.close();
}

inline void loadFromText(Queue& queue, const std::string& filename) {
    std::ifstream file(filename);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    queue.QCLEAR();
    int size;
    file >> size;
    file.ignore();
    
    std::string line;
    for (int i = 0; i < size; ++i) {
        std::getline(file, line);
        queue.QPUSH(line);
    }
    
    file.close();
}

// Бинарный формат
inline void saveToBinary(Queue& queue, const std::string& filename) {
    std::ofstream file(filename, std::ios::binary);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    // Создаем временную очередь для сохранения порядка
    Queue temp;
    std::vector<std::string> items;
    
    // Переливаем элементы во временную очередь
    while (!queue.isEmpty()) {
        std::string val = queue.QPOP();
        temp.QPUSH(val);
        items.push_back(val);
    }
    
    // Восстанавливаем оригинальную очередь
    while (!temp.isEmpty()) {
        queue.QPUSH(temp.QPOP());
    }
    
    // Записываем в файл
    int count = items.size();
    file.write(reinterpret_cast<const char*>(&count), sizeof(count));
    
    for (const auto& item : items) {
        size_t len = item.length();
        file.write(reinterpret_cast<const char*>(&len), sizeof(len));
        file.write(item.c_str(), len);
    }
    
    file.close();
}

inline void loadFromBinary(Queue& queue, const std::string& filename) {
    std::ifstream file(filename, std::ios::binary);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    queue.QCLEAR();
    int count;
    file.read(reinterpret_cast<char*>(&count), sizeof(count));
    
    for (int i = 0; i < count; ++i) {
        size_t len;
        file.read(reinterpret_cast<char*>(&len), sizeof(len));
        std::string str(len, '\0');
        file.read(&str[0], len);
        queue.QPUSH(str);
    }
    
    file.close();
}

// ChainingHashTable

// Текстовый формат
inline void saveToText(const ChainingHashTable& ht, const std::string& filename) {
    std::ofstream file(filename);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    // Получаем все элементы и сохраняем
    auto elements = ht.getAllElements();
    file << elements.size() << "\n";
    for (const auto& elem : elements) {
        file << elem.first << " " << elem.second << "\n";
    }
    
    file.close();
}

inline void loadFromText(ChainingHashTable& ht, const std::string& filename) {
    std::ifstream file(filename);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    ht.clear();
    
    int size;
    file >> size;
    
    for (int i = 0; i < size; ++i) {
        int key, value;
        file >> key >> value;
        ht.add(std::make_pair(key, value));
    }
    
    file.close();
}

// БИНАРНЫЙ ФОРМАТ - ПОЛНОСТЬЮ РЕАЛИЗОВАН
inline void saveToBinary(const ChainingHashTable& ht, const std::string& filename) {
    std::ofstream file(filename, std::ios::binary);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    // Получаем все элементы
    auto elements = ht.getAllElements();
    int elementCount = elements.size();
    
    // Сохраняем количество элементов
    file.write(reinterpret_cast<const char*>(&elementCount), sizeof(elementCount));
    
    // Сохраняем каждый элемент (пару ключ-значение)
    for (const auto& elem : elements) {
        int key = elem.first;
        int value = elem.second;
        file.write(reinterpret_cast<const char*>(&key), sizeof(key));
        file.write(reinterpret_cast<const char*>(&value), sizeof(value));
    }
    
    file.close();
}

inline void loadFromBinary(ChainingHashTable& ht, const std::string& filename) {
    std::ifstream file(filename, std::ios::binary);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    ht.clear();
    
    // Читаем количество элементов
    int elementCount;
    file.read(reinterpret_cast<char*>(&elementCount), sizeof(elementCount));
    
    // Читаем и добавляем каждый элемент
    for (int i = 0; i < elementCount; ++i) {
        int key, value;
        file.read(reinterpret_cast<char*>(&key), sizeof(key));
        file.read(reinterpret_cast<char*>(&value), sizeof(value));
        ht.add(std::make_pair(key, value));
    }
    
    file.close();
}

// RBTree

namespace RBTreeSerializer {
    // Рекурсивный обход для сохранения (префиксный порядок)
    inline void serializeNodeBinary(std::ofstream& file, TNode* node, TNode* nil) {
        if (node == nil || node == nullptr) {
            // Сохраняем маркер NIL узла
            char marker = 'N';
            file.write(&marker, sizeof(marker));
            return;
        }
        
        // Сохраняем маркер обычного узла
        char marker = 'U';
        file.write(&marker, sizeof(marker));
        
        // Сохраняем данные
        size_t dataLen = node->data.length();
        file.write(reinterpret_cast<const char*>(&dataLen), sizeof(dataLen));
        file.write(node->data.c_str(), dataLen);
        
        // Сохраняем цвет
        char color = (node->color == RED) ? 'R' : 'B';
        file.write(&color, sizeof(color));
        
        // Рекурсивно сохраняем поддеревья
        serializeNodeBinary(file, node->left, nil);
        serializeNodeBinary(file, node->right, nil);
    }
    
    // Рекурсивное чтение и построение дерева
    inline TNode* deserializeNodeBinary(std::ifstream& file, TNode* parent, TNode* nil, RBTree& tree) {
        char marker;
        file.read(&marker, sizeof(marker));
        
        if (marker == 'N') {
            return nil; // NIL узел
        }
        
        // Читаем данные
        size_t dataLen;
        file.read(reinterpret_cast<char*>(&dataLen), sizeof(dataLen));
        std::string data(dataLen, '\0');
        file.read(&data[0], dataLen);
        
        // Читаем цвет
        char colorChar;
        file.read(&colorChar, sizeof(colorChar));
        Color color = (colorChar == 'R') ? RED : BLACK;
        
        // Создаем узел
        TNode* node = tree.createNode(data);
        node->color = color;
        node->parent = parent;
        
        // Рекурсивно создаем поддеревья
        node->left = deserializeNodeBinary(file, node, nil, tree);
        node->right = deserializeNodeBinary(file, node, nil, tree);
        
        return node;
    }
}

// Текстовый формат 
inline void saveToText(const RBTree& tree, const std::string& filename) {
    std::ofstream file(filename);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    // Получаем все узлы
    auto nodes = tree.getAllNodes();
    
    // Сохраняем количество узлов
    file << nodes.size() << "\n";
    
    // Сохраняем каждый узел
    for (const auto& node : nodes) {
        file << node.first << "\n";
        file << (node.second == RED ? "RED" : "BLACK") << "\n";
    }
    
    file.close();
}

inline void loadFromText(RBTree& tree, const std::string& filename) {
    std::ifstream file(filename);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    // Очищаем дерево
    tree.clear();
    
    int nodeCount;
    file >> nodeCount;
    file.ignore();
    
    // Читаем и вставляем узлы
    for (int i = 0; i < nodeCount; ++i) {
        std::string data;
        std::string colorStr;
        
        std::getline(file, data);
        std::getline(file, colorStr);
        
        Color color = (colorStr == "RED") ? RED : BLACK;
        
        // Вставляем узел с указанным цветом
        tree.insertWithColor(data, color);
    }
    
    file.close();
}

// бинарный фоомат
inline void saveToBinary(const RBTree& tree, const std::string& filename) {
    std::ofstream file(filename, std::ios::binary);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    // Сохраняем структуру дерева с помощью рекурсивного обхода
    RBTreeSerializer::serializeNodeBinary(file, tree.getRoot(), tree.getNil());
    
    file.close();
}

inline void loadFromBinary(RBTree& tree, const std::string& filename) {
    std::ifstream file(filename, std::ios::binary);
    if (!file.is_open()) throw std::runtime_error("Cannot open file: " + filename);
    
    // Очищаем дерево
    tree.clear();
    
    // Читаем и восстанавливаем структуру дерева
    TNode* newRoot = RBTreeSerializer::deserializeNodeBinary(file, tree.getNil(), tree.getNil(), tree);
    tree.setRoot(newRoot);
    
    file.close();
}

#endif 