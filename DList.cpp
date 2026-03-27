#include "DList.h"
#include <iostream>
#include <fstream>
#include <vector>
#include <stdexcept>

using namespace std;

DList::DList() {
    head = nullptr;
    tail = nullptr;
    size = 0;
}

DList::~DList() {
    DLCLEAR();
}

void DList::DLCLEAR() {
    DNode* current = head;
    while (current != nullptr) {
        DNode* temp = current;
        current = current->next;
        delete temp;
    }
    head = nullptr;
    tail = nullptr;
    size = 0;
    cout << "Двусвязный список очищен." << endl;
}

void DList::DLPUSH_HEAD(const string& value) {
    DNode* newNode = new DNode{ value, head, nullptr };

    if (head != nullptr) {
        head->prev = newNode;
    }
    else {
        tail = newNode;
    }

    head = newNode;
    size++;
    //cout << "Элемент \"" << value << "\" добавлен в голову списка" << endl;
}

void DList::DLPUSH_TAIL(const string& value) {
    DNode* newNode = new DNode{ value, nullptr, tail };

    if (tail != nullptr) {
        tail->next = newNode;
    }
    else {
        head = newNode;
    }

    tail = newNode;
    size++;
    //cout << "Элемент \"" << value << "\" добавлен в хвост списка" << endl;
}

void DList::DLPUSH_BEFORE(const string& target, const string& value) {
    if (head == nullptr) {
        cout << "Список пуст!" << endl;
        return;
    }

    DNode* current = head;
    while (current != nullptr && current->data != target) {
        current = current->next;
    }

    if (current == nullptr) {
        //cout << "Элемент \"" << target << "\" не найден в списке" << endl;
        return;
    }

    if (current == head) {
        DLPUSH_HEAD(value);
        return;
    }

    DNode* newNode = new DNode{ value, current, current->prev };
    current->prev->next = newNode;
    current->prev = newNode;
    size++;
    //cout << "Элемент \"" << value << "\" добавлен перед \"" << target << "\"" << endl;
}

void DList::DLPUSH_AFTER(const string& target, const string& value) {
    if (head == nullptr) {
        cout << "Список пуст!" << endl;
        return;
    }

    DNode* current = head;
    while (current != nullptr && current->data != target) {
        current = current->next;
    }

    if (current == nullptr) {
        //cout << "Элемент \"" << target << "\" не найден в списке" << endl;
        return;
    }

    if (current == tail) {
        DLPUSH_TAIL(value);
        return;
    }

    DNode* newNode = new DNode{ value, current->next, current };
    current->next->prev = newNode;
    current->next = newNode;
    size++;
    //cout << "Элемент \"" << value << "\" добавлен после \"" << target << "\"" << endl;
}

void DList::DLREMOVE_HEAD() {
    if (head == nullptr) {
        cout << "Список пуст!" << endl;
        return;
    }

    DNode* temp = head;
    head = head->next;

    if (head != nullptr) {
        head->prev = nullptr;
    }
    else {
        tail = nullptr;
    }

    //cout << "Элемент \"" << temp->data << "\" удалён из головы списка" << endl;
    delete temp;
    size--;
}

void DList::DLREMOVE_TAIL() {
    if (tail == nullptr) {
        cout << "Список пуст!" << endl;
        return;
    }

    DNode* temp = tail;
    tail = tail->prev;

    if (tail != nullptr) {
        tail->next = nullptr;
    }
    else {
        head = nullptr;
    }

    //cout << "Элемент \"" << temp->data << "\" удалён из хвоста списка" << endl;
    delete temp;
    size--;
}

void DList::DLREMOVE_BEFORE(const string& target) {
    if (head == nullptr || head->next == nullptr) {
        cout << "Недостаточно элементов для удаления!" << endl;
        return;
    }

    DNode* current = head;
    while (current != nullptr && current->data != target) {
        current = current->next;
    }

    if (current == nullptr) {
        //cout << "Элемент \"" << target << "\" не найден в списке" << endl;
        return;
    }

    if (current->prev == nullptr) {
        cout << "Невозможно удалить элемент перед головой!" << endl;
        return;
    }

    if (current->prev == head) {
        DLREMOVE_HEAD();
        return;
    }

    DNode* temp = current->prev;
    temp->prev->next = current;
    current->prev = temp->prev;

    //cout << "Элемент \"" << temp->data << "\" удалён перед \"" << target << "\"" << endl;
    delete temp;
    size--;
}

void DList::DLREMOVE_AFTER(const string& target) {
    if (head == nullptr || head->next == nullptr) {
        cout << "Недостаточно элементов для удаления!" << endl;
        return;
    }

    DNode* current = head;
    while (current != nullptr && current->data != target) {
        current = current->next;
    }

    if (current == nullptr) {
        //cout << "Элемент \"" << target << "\" не найден в списке" << endl;
        return;
    }

    if (current->next == nullptr) {
        cout << "Невозможно удалить элемент после хвоста!" << endl;
        return;
    }

    if (current->next == tail) {
        DLREMOVE_TAIL();
        return;
    }

    DNode* temp = current->next;
    current->next = temp->next;
    temp->next->prev = current;

    //cout << "Элемент \"" << temp->data << "\" удалён после \"" << target << "\"" << endl;
    delete temp;
    size--;
}

void DList::DLREMOVE_VALUE(const string& value) {
    if (head == nullptr) {
        cout << "Список пуст!" << endl;
        return;
    }
    DNode* current = head;

    while (current != nullptr && current->data != value) {
        current = current->next;
    }

    if (current == nullptr) {
        //cout << "Элемент \"" << value << "\" не найден в списке" << endl;
        return;
    }

    if (current == head) {
        head = current->next;
        if (head != nullptr) {
            head->prev = nullptr;
        }
        else {
            tail = nullptr;
        }

        //cout << "Элемент \"" << current->data << "\" (голова) удалён по значению" << endl;
        delete current;
        size--;
        return;
    }

    if (current == tail) {
        tail = current->prev;
        if (tail != nullptr) {
            tail->next = nullptr;
        }
        else {
            head = nullptr;
        }

        //cout << "Элемент \"" << current->data << "\" (хвост) удалён по значению" << endl;
        delete current;
        size--;
        return;
    }

    current->prev->next = current->next;
    current->next->prev = current->prev;

    //cout << "Элемент \"" << current->data << "\" удалён по значению" << endl;
    delete current;
    size--;
}

bool DList::DLSEARCH(const string& value) const {
    DNode* current = head;
    size_t index = 0;
    while (current != nullptr) {
        if (current->data == value) {
            //cout << "Элемент \"" << value << "\" найден на позиции " << index << endl;
            return true;
        }
        current = current->next;
        index++;
    }
    //cout << "Элемент \"" << value << "\" не найден в списке" << endl;
    return false;
}

string DList::DLGET(size_t index) const {
    if (index >= size) {
        throw out_of_range("Индекс " + to_string(index) + " вне диапазона [0, " + to_string(size - 1) + "]");
    }

    DNode* current = head;
    for (size_t i = 0; i < index; i++) {
        current = current->next;
    }
    return current->data;
}

void DList::DLPRINT_FORWARD() const {
    if (head == nullptr) {
        cout << "Список пуст!" << endl;
        return;
    }

    cout << "Двусвязный список [" << size << "] (прямой порядок): NULL <- ";
    DNode* current = head;
    while (current != nullptr) {
        cout << "\"" << current->data << "\"";
        if (current->next != nullptr) {
            cout << " <-> ";
        }
        current = current->next;
    }
    cout << " -> NULL" << endl;
}

void DList::DLPRINT_BACKWARD() const {
    if (tail == nullptr) {
        cout << "Список пуст!" << endl;
        return;
    }

    cout << "Двусвязный список [" << size << "] (обратный порядок): NULL <- ";
    DNode* current = tail;
    while (current != nullptr) {
        cout << "\"" << current->data << "\"";
        if (current->prev != nullptr) {
            cout << " <-> ";
        }
        current = current->prev;
    }
    cout << " -> NULL" << endl;
}

void DList::DLPRINT_HEAD() const {
    if (head == nullptr) {
        cout << "Список пуст!" << endl;
    }
    else {
        cout << "Голова списка: \"" << head->data << "\"" << endl;
    }
}

void DList::DLPRINT_TAIL() const {
    if (tail == nullptr) {
        cout << "Список пуст!" << endl;
    }
    else {
        cout << "Хвост списка: \"" << tail->data << "\"" << endl;
    }
}

size_t DList::DLLENGTH() const {
    return size;
}
