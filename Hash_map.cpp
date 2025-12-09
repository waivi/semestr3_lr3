#include "Hash_map.h"
#include <string>
#include <algorithm>
#include <sstream>
#include <iomanip>

Node::Node(int k, int v) : keyValue(std::make_pair(k, v)), next(nullptr) {}

// Реализация хэш-функции
int ChainingHashTable::hash(int key) {
    // Используем простое деление для хэширования
    return key % capacity;
}

ChainingHashTable::ChainingHashTable(int capacity) 
    : capacity(capacity > 0 ? capacity : 10), size(0) {
    table.resize(this->capacity, nullptr);
}

ChainingHashTable::~ChainingHashTable() {
    clear();
}

// Добавление элемента
void ChainingHashTable::add(std::pair<int, int> keyValue) {
    // Проверка на дубликаты
    auto result = contains(keyValue.first);
    if (result.first) {
        return; // элемент уже существует
    }

    int h = hash(keyValue.first);
    Node* newNode = new Node(keyValue.first, keyValue.second);

    // Вставка в цепочку
    if (table[h] == nullptr) {
        // Пустая ячейка
        table[h] = newNode;
    } else {
        // Коллизия - добавляем в конец цепочки
        Node* current = table[h];
        while (current->next != nullptr) {
            current = current->next;
        }
        current->next = newNode;
    }
    size++;
}

// Удаление элемента
void ChainingHashTable::remove(int key) {
    int h = hash(key);
    Node* current = table[h];
    Node* prev = nullptr;

    while (current != nullptr) {
        if (current->keyValue.first == key) {
            // Нашли элемент для удаления
            if (prev == nullptr) {
                // Удаляем первый элемент цепочки
                table[h] = current->next;
            } else {
                // Удаляем из середины или конца
                prev->next = current->next;
            }
            delete current;
            size--;
            return;
        }
        prev = current;
        current = current->next;
    }
    // Элемент не найден - ничего не делаем
}

// Поиск элемента
std::pair<bool, int> ChainingHashTable::contains(int key) {
    int h = hash(key);
    Node* current = table[h];

    while (current != nullptr) {
        if (current->keyValue.first == key) {
            return std::make_pair(true, current->keyValue.second);
        }
        current = current->next;
    }
    return std::make_pair(false, -1);
}

// Получение строкового представления
std::string ChainingHashTable::toString() const {
    std::stringstream ss;
    
    for (int i = 0; i < capacity; i++) {
        ss << "[" << std::setw(3) << i << "]: ";
        Node* current = table[i];
        
        if (current == nullptr) {
            ss << "empty";
        } else {
            while (current != nullptr) {
                ss << "(" << current->keyValue.first 
                   << "," << current->keyValue.second << ")";
                if (current->next != nullptr) {
                    ss << " -> ";
                }
                current = current->next;
            }
        }
        ss << "\n";
    }
    
    return ss.str();
}
std::vector<std::pair<int, int>> ChainingHashTable::getAllElements() const {
    std::vector<std::pair<int, int>> elements;
    for (int i = 0; i < capacity; ++i) {
        Node* current = table[i];
        while (current != nullptr) {
            elements.push_back(current->keyValue);
            current = current->next;
        }
    }
    return elements;
}
// Анализ длины цепочек
void ChainingHashTable::getChainLengths(int& minLength, int& maxLength, double& avgLength) const {
    minLength = INT_MAX;
    maxLength = 0;
    int totalLength = 0;
    int nonEmptyChains = 0;

    for (int i = 0; i < capacity; i++) {
        int length = 0;
        Node* current = table[i];
        
        while (current != nullptr) {
            length++;
            current = current->next;
        }

        if (length > 0) {
            minLength = std::min(minLength, length);
            maxLength = std::max(maxLength, length);
            totalLength += length;
            nonEmptyChains++;
        }
    }

    if (nonEmptyChains == 0) {
        minLength = 0;
        avgLength = 0.0;
    } else {
        avgLength = static_cast<double>(totalLength) / nonEmptyChains;
    }
}

// Геттеры
int ChainingHashTable::getSize() const {
    return size;
}

int ChainingHashTable::getCapacity() const {
    return capacity;
}

double ChainingHashTable::getLoadFactor() const {
    return (capacity > 0) ? static_cast<double>(size) / capacity : 0.0;
}

// Очистка таблицы
void ChainingHashTable::clear() {
    for (int i = 0; i < capacity; i++) {
        Node* current = table[i];
        while (current != nullptr) {
            Node* temp = current;
            current = current->next;
            delete temp;
        }
        table[i] = nullptr;
    }
    size = 0;
}

// Проверка на пустоту
bool ChainingHashTable::isEmpty() const {
    return size == 0;
}