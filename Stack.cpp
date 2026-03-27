#include "Stack.h"
#include <iostream>
#include <fstream>
#include <vector>
#include <stdexcept>

using namespace std;

Stack::Stack() {
    top = nullptr;
    size = 0;
}

Stack::~Stack() {
    SCLEAR();
}

void Stack::SCLEAR() {
    while (top != nullptr) {
        SNode* temp = top;
        top = top->next;
        delete temp;
    }
    size = 0;
    cout << "Стек очищен." << endl;

}

void Stack::SPUSH(const string& value) {
    SNode* newNode = new SNode;
    newNode->data = value;
    newNode->next = top;
    top = newNode;
    size++;
    //cout << "Элемент \"" << value << "\" добавлен в стек" << endl;

}

string Stack::SPOP() {
    if (top == nullptr) {
        throw runtime_error("Ошибка: стек пуст!");
    }

    SNode* temp = top;
    string value = temp->data;
    top = top->next;
    delete temp;
    size--;

    //cout << "Элемент \"" << value << "\" извлечён из стека" << endl;

    return value;
}

string Stack::STOP() const {
    if (top == nullptr) {
        throw runtime_error("Ошибка: стек пуст!");
    }
    return top->data;
}

void Stack::SPRINT() const {
    if (top == nullptr) {
        cout << "Стек пуст!" << endl;
        return;
    }

    cout << "Стек [" << size << "] (вершина -> основание): ";
    SNode* current = top;
    while (current != nullptr) {
        cout << "\"" << current->data << "\"";
        if (current->next != nullptr) {
            cout << " -> ";
        }
        current = current->next;
    }
    cout << " -> NULL" << endl;
}

bool Stack::isEmpty() const {
        return size == 0;
    }