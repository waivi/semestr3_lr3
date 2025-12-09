#include "Queue.h"
#include <iostream>
#include <fstream>
#include <vector>
#include <stdexcept>

using namespace std;

Queue::Queue() {
    front = nullptr;
    rear = nullptr;
    size = 0;
}

Queue::~Queue() {
    QCLEAR();
}

void Queue::QCLEAR() {
    while (front != nullptr) {
        QNode* temp = front;
        front = front->next;
        delete temp;
    }
    rear = nullptr;
    size = 0;
    cout << "Очередь очищена." << endl;
}

void Queue::QPUSH(const string& value) {
    QNode* newNode = new QNode{ value, nullptr };

    if (rear == nullptr) {
        // Очередь пуста - новый элемент становится началом и концом
        front = newNode;
        rear = newNode;
    }
    else {
        // Добавляем в конец очереди
        rear->next = newNode;
        rear = newNode;
    }

    size++;
    //cout << "Элемент \"" << value << "\" добавлен в очередь" << endl;
}

string Queue::QPOP() {
    if (front == nullptr) {
        throw runtime_error("Ошибка: очередь пуста!");
    }

    QNode* temp = front;
    string value = temp->data;
    front = front->next;

    // Если очередь стала пустой, обнуляем rear
    if (front == nullptr) {
        rear = nullptr;
    }

    delete temp;
    size--;

    //cout << "Элемент \"" << value << "\" извлечён из очереди" << endl;
    return value;
}

string Queue::QFRONT() const {
    if (front == nullptr) {
        throw runtime_error("Ошибка: очередь пуста!");
    }
    return front->data;
}

void Queue::QPRINT() const {
    if (front == nullptr) {
        cout << "Очередь пуста!" << endl;
        return;
    }

    cout << "Очередь [" << size << "] (начало-> конец): ";
    QNode* current = front;
    while (current != nullptr) {
        cout << "\"" << current->data << "\"";
        if (current->next != nullptr) {
            cout << " -> ";
        }
        current = current->next;
    }
    cout << " -> NULL" << endl;
}
bool Queue::isEmpty() const {
    return size == 0;
}