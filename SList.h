#pragma once

#include <cstdint>
#include <string>

using namespace std;

struct SLNode {
    string data;
    SLNode* next;
};

class SList {
private:
    SLNode* head;
    size_t size;
    void printReverse(SLNode* node) const;

public:
    SList();
    ~SList();
    
    void SLCLEAR();
    
    // Добавление элементов (4 способа)
    void SLPUSH_HEAD(const string& value);
    void SLPUSH_TAIL(const string& value);
    void SLPUSH_BEFORE(const string& target, const string& value);
    void SLPUSH_AFTER(const string& target, const string& value);
    
    // Удаление элементов (4 способа)
    void SLREMOVE_HEAD();
    void SLREMOVE_TAIL();
    void SLREMOVE_BEFORE(const string& target);
    void SLREMOVE_AFTER(const string& target);
    
    // Удаление по значению
    void SLREMOVE_VALUE(const string& value);
    
    // Поиск по значению
    bool SLSEARCH(const string& value) const;
    
    // Получение элемента
    string SLGET(size_t index) const;
    
    // Чтение (несколько способов)
    void SLPRINT_FORWARD() const;
    void SLPRINT_BACKWARD() const;
    void SLPRINT_HEAD() const;
    void SLPRINT_TAIL() const;
    
    // Длина списка
    size_t SLLENGTH() const;
};