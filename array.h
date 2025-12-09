#pragma once

#include <cstdint>
#include <string>

using namespace std;

class MArray {
private:
    string* data;
    size_t capacity;
    size_t size;
    void MRESIZE();

public:
    MArray();
    ~MArray();
    
    void MCLEAR();
    
    void MADDINDEX(size_t index, const string& value);
    void MADDEND(const string& value);
    
    string MGETINDEX(size_t index) const;
    void MREMOVEINDEX(size_t index);
    void MREPLACEINDEX(size_t index, const string& newValue);
    
    size_t MLENGTH() const;
    void MPRINT() const;
};