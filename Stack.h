#pragma once
#include <cstdint>
#include <string>

using namespace std;

struct SNode {
    string data;
    SNode* next;
};

class Stack {
private:
    SNode* top;
    size_t size;
public:
    Stack();
    ~Stack();
    
    void SCLEAR();
    
    void SPUSH(const string& value);
    string SPOP();
    string STOP() const;
    void SPRINT() const;
    bool isEmpty() const;
};