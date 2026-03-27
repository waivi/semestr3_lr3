#pragma once

#include <cstdint>
#include <string>

using namespace std;

struct DNode {
    string data;
    DNode* next;
    DNode* prev;
};

class DList {
private:
    DNode* head;
    DNode* tail;
    size_t size;

public:
    DList();
    ~DList();
    
    void DLCLEAR();
    
    // Добавление элементов (4 способа)
    void DLPUSH_HEAD(const string& value);
    void DLPUSH_TAIL(const string& value);
    void DLPUSH_BEFORE(const string& target, const string& value);
    void DLPUSH_AFTER(const string& target, const string& value);
    
    // Удаление элементов (4 способа)
    void DLREMOVE_HEAD();
    void DLREMOVE_TAIL();
    void DLREMOVE_BEFORE(const string& target);
    void DLREMOVE_AFTER(const string& target);
    
    // Удаление по значению
    void DLREMOVE_VALUE(const string& value);
    
    // Поиск по значению
    bool DLSEARCH(const string& value) const;
    
    // Получение элемента
    string DLGET(size_t index) const;
    
    // Чтение (несколько способов)
    void DLPRINT_FORWARD() const;
    void DLPRINT_BACKWARD() const;
    void DLPRINT_HEAD() const;
    void DLPRINT_TAIL() const;
    
    // Длина списка
    size_t DLLENGTH() const;
};