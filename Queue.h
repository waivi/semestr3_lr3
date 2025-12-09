#pragma once

#include <cstdint>
#include <string>

using namespace std;

struct QNode {
    string data;
    QNode* next;
};

class Queue {
private:
    QNode* front;   // Указатель на начало очереди
    QNode* rear;    // Указатель на конец очереди  
    size_t size;

public:
    Queue();
    ~Queue();
    
    void QCLEAR();
    
    // Основные операции
    void QPUSH(const string& value);
    string QPOP();
    string QFRONT() const;
    
    // Чтение и вывод
    void QPRINT() const;
    bool isEmpty() const;
};