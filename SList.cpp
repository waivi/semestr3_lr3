#include "SList.h"
#include <iostream>
#include <fstream>
#include <vector>
#include <stdexcept>

using namespace std;


void SList::printReverse(SLNode* node) const {
    if (node == nullptr) return;
    printReverse(node->next);
    cout << "\"" << node->data << "\"";
    if (node != nullptr) cout << " <- ";
}

SList::SList() {
    head = nullptr;
    size = 0;
}

SList::~SList() {
    SLCLEAR();
}

void SList::SLCLEAR() {
    SLNode* current = head;
    while (current != nullptr) {
        SLNode* temp = current;
        current = current->next;
        delete temp;
    }
    head = nullptr;
    size = 0;
    cout << "Список очищен." << endl;
}

void SList::SLPUSH_HEAD(const string& value) {
    SLNode* newNode = new SLNode;
    newNode->data = value;
    newNode->next = head;
    head = newNode;
    size++;
    //cout << "Элемент \"" << value << "\" добавлен в голову списка" << endl;
}

void SList::SLPUSH_TAIL(const string& value) {
    SLNode* newNode = new SLNode;
    newNode->data = value;
    newNode->next = nullptr;
    
    if (head == nullptr) {
        head = newNode;
    } else {
        SLNode* current = head;
        while (current->next != nullptr) {
            current = current->next;
        }
        current->next = newNode;
    }
    size++;
    //cout << "Элемент \"" << value << "\" добавлен в хвост списка" << endl;
}

void SList::SLPUSH_BEFORE(const string& target, const string& value) {
    if (head == nullptr) {
        cout << "Список пуст!" << endl;
        return;
    }
    
    if (head->data == target) {
        SLPUSH_HEAD(value);
        return;
    }
    
    SLNode* current = head;
    while (current->next != nullptr && current->next->data != target) {
        current = current->next;
    }
    
    if (current->next != nullptr) {
        SLNode* newNode = new SLNode;
        newNode->data = value;
        newNode->next = current->next;
        current->next = newNode;
        size++;
        //cout << "Элемент \"" << value << "\" добавлен перед \"" << target << "\"" << endl;
    }
    else {
        //cout << "Элемент \"" << target << "\" не найден в списке" << endl;
    }
}

void SList::SLPUSH_AFTER(const string& target, const string& value) {
    SLNode* current = head;
    while (current != nullptr && current->data != target) {
        current = current->next;
    }
    
    if (current != nullptr) {
        SLNode* newNode = new SLNode;
        newNode->data = value;
        newNode->next = current->next;
        current->next = newNode;
        size++;
        //cout << "Элемент \"" << value << "\" добавлен после \"" << target << "\"" << endl;
    }else{

            //cout << "Элемент \"" << target << "\" не найден в списке" << endl;
    }
}

void SList::SLREMOVE_HEAD() {
    if (head == nullptr) {
        cout << "Список пуст!" << endl;
        return;
    }
    
    SLNode* temp = head;
    head = head->next;
    //cout << "Элемент \"" << temp->data << "\" удалён из головы списка" << endl;
    delete temp;
    size--;
}

void SList::SLREMOVE_TAIL() {
    if (head == nullptr) {
        cout << "Список пуст!" << endl;
        return;
    }
    
    if (head->next == nullptr) {
        //cout << "Элемент \"" << head->data << "\" удалён из хвоста списка" << endl;
        delete head;
        head = nullptr;
    }
    else {
        SLNode* current = head;
        while (current->next->next != nullptr) {
            current = current->next;
        }
        //cout << "Элемент \"" << current->next->data << "\" удалён из хвоста списка" << endl;
        delete current->next;
        current->next = nullptr;
    }
    size--;
}

void SList::SLREMOVE_BEFORE(const string& target) {
    if (head == nullptr || head->next == nullptr) {
        cout << "Недостаточно элементов для удаления!" << endl;
        return;
    }
    
    if (head->data == target) {
        cout << "Невозможно удалить элемент перед головой!" << endl;
        return;
    }
    
    if (head->next->data == target) {
        SLREMOVE_HEAD();
        return;
    }
    
    SLNode* current = head;
    while (current->next->next != nullptr && current->next->next->data != target) {
        current = current->next;
    }
    
    if (current->next->next != nullptr) {
        SLNode* temp = current->next;
        current->next = current->next->next;
        //cout << "Элемент \"" << temp->data << "\" удалён перед \"" << target << "\"" << endl;
        delete temp;
        size--;
    }
    else {
        cout << "Элемент \"" << target << "\" не найден или перед ним нет элемента" << endl;
    }
}

void SList::SLREMOVE_AFTER(const string& target) {
    SLNode* current = head;
    while (current != nullptr && current->data != target) {
        current = current->next;
    }
    
    if (current != nullptr && current->next != nullptr) {
        SLNode* temp = current->next;
        current->next = current->next->next;
        //cout << "Элемент \"" << temp->data << "\" удалён после \"" << target << "\"" << endl;
        delete temp;
        size--;
    }
    else {
        cout << "Элемент \"" << target << "\" не найден или после него нет элемента" << endl;
    }
}

void SList::SLREMOVE_VALUE(const string& value) {
    if (head == nullptr) {
        cout << "Список пуст!" << endl;
        return;
    }
    
    if (head->data == value) {
        SLNode* temp = head;
        head = head->next;
        //cout << "Элемент \"" << temp->data << "\" (голова) удалён по значению" << endl;
        delete temp;
        size--;
        return;
    }
    
    SLNode* current = head;
    while (current->next != nullptr && current->next->data != value) {
        current = current->next;
    }
    
    if (current->next != nullptr) {
        SLNode* temp = current->next;
        current->next = current->next->next;
        //cout << "Элемент \"" << temp->data << "\" удалён по значению" << endl;
        delete temp;
        size--;
    }
    else {
        //cout << "Элемент \"" << value << "\" не найден в списке" << endl;
    }
}

bool SList::SLSEARCH(const string& value) const {
    SLNode* current = head;
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

string SList::SLGET(size_t index) const {
    if (index >= size) {
        throw out_of_range("Индекс " + to_string(index) + " вне диапазона [0, " + to_string(size - 1) + "]");
    }
    
    SLNode* current = head;
    for (size_t i = 0; i < index; i++) {
        current = current->next;
    }
    return current->data;
}

void SList::SLPRINT_FORWARD() const {
    if (head == nullptr) {
        cout << "Список пуст!" << endl;
        return;
    }
    
    cout << "Список [" << size << "] (прямой порядок): ";
    SLNode* current = head;
    while (current != nullptr) {
        cout << "\"" << current->data << "\"";
        if (current->next != nullptr) {
            cout << " -> ";
        }
        current = current->next;
    }
    cout << " -> NULL" << endl;
}

void SList::SLPRINT_BACKWARD() const {
    if (head == nullptr) {
        cout << "Список пуст!" << endl;
        return;
    }
    
    cout << "Список [" << size << "] (обратный порядок): NULL <- ";
    printReverse(head);
    cout << endl;
}

void SList::SLPRINT_HEAD() const {
    if (head == nullptr) {
        cout << "Список пуст!" << endl;
    }
    else {
        cout << "Голова списка: \"" << head->data << "\"" << endl;
    }
}

void SList::SLPRINT_TAIL() const {
    if (head == nullptr) {
        cout << "Список пуст!" << endl;
        return;
    }
    
    SLNode* current = head;
    while (current->next != nullptr) {
        current = current->next;
    }
    cout << "Хвост списка: \"" << current->data << "\"" << endl;
}

size_t SList::SLLENGTH() const {
    return size;
}