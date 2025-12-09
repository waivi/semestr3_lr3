#pragma once

#include <vector>
#include <string>
#include <climits>
#include <utility>

// Узел для метода цепочек
struct Node {
    std::pair<int, int> keyValue; // пара ключ-значение
    Node* next;
    Node(int k, int v);
};

// Класс хэш-таблицы методом цепочек
class ChainingHashTable {
private:
    int capacity; // размер таблицы
    int size;     // количество элементов
    std::vector<Node*> table; // вектор указателей на узлы

    // Хэш-функция
    int hash(int key);

public:
    ChainingHashTable(int capacity);
    ChainingHashTable() : ChainingHashTable(10) {}
    ~ChainingHashTable();
    ChainingHashTable(const ChainingHashTable&) = delete;
    ChainingHashTable& operator=(const ChainingHashTable&) = delete;

    // Основные операции
    void add(std::pair<int, int> keyValue);
    void remove(int key);
    std::pair<bool, int> contains(int key);
    
    // Получение строкового представления
    std::string toString() const;
    std::vector<std::pair<int, int>> getAllElements() const;

    // Анализ длины цепочек
    void getChainLengths(int& minLength, int& maxLength, double& avgLength) const;
    
    // Геттеры
    int getSize() const;
    int getCapacity() const;
    double getLoadFactor() const;
    
    // Вспомогательные методы для отладки
    void clear();
    bool isEmpty() const;
};
