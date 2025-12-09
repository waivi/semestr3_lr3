#include "array.h"
#include <iostream>
#include <fstream>
#include <vector>
#include <stdexcept>

using namespace std;


void MArray::MRESIZE() {
    if (size < capacity) return;

    size_t newCapacity = (capacity == 0) ? 4 : capacity * 2;
    string* newData = new string[newCapacity];

    for (size_t i = 0; i < size; i++) {
        newData[i] = data[i];
    }

    delete[] data;
    data = newData;
    capacity = newCapacity;
}

MArray::MArray() {
    data = nullptr;
    size = 0;
    capacity = 0;
}

MArray::~MArray() {
    delete[] data;
}

void MArray::MCLEAR() {
    delete[] data;
    data = nullptr;
    size = 0;
    capacity = 0;
    cout << "Массив очищен." << endl;
}

void MArray::MADDINDEX(size_t index, const string& value) {
    if (index > size) {
        cout << "Ошибка: индекс " << index << " превышает размер массива (" << size << ")" << endl;
        return;
    }

    MRESIZE();

    for (size_t i = size; i > index; i--) {
        data[i] = data[i - 1];
    }

    data[index] = value;
    size++;
    //cout << "Элемент \"" << value << "\" добавлен на позицию " << index << endl;
}

void MArray::MADDEND(const string& value) {
    MRESIZE();

    data[size] = value;
    size++;
    //cout << "Элемент \"" << value << "\" добавлен в конец массива" << endl;
}

string MArray::MGETINDEX(size_t index) const {
    if (index >= size) {
        throw out_of_range("Индекс " + to_string(index) + " вне диапазона [0, " + to_string(size - 1) + "]");
    }
    return data[index];
}

void MArray::MREMOVEINDEX(size_t index) {
    if (index >= size) {
        cout << "Ошибка: индекс " << index << " вне диапазона [0, " << size - 1 << "]" << endl;
        return;
    }

    //cout << "Элемент \"" << data[index] << "\" удалён с позиции " << index << endl;

    for (size_t i = index; i < size - 1; i++) {
        data[i] = data[i + 1];
    }

    size--;
}

void MArray::MREPLACEINDEX(size_t index, const string& newValue) {
    if (index >= size) {
        cout << "Ошибка: индекс " << index << " вне диапазона [0, " << size - 1 << "]" << endl;
        return;
    } 

    string oldValue = data[index];
    data[index] = newValue;
    //cout << "Элемент \"" << oldValue << "\" заменён на \"" << newValue << "\" на позиции " << index << endl;
}

size_t MArray::MLENGTH() const {
    return size;
}

void MArray::MPRINT() const {
    if (size == 0) {
        cout << "Массив пуст!" << endl;
        return;
    }

    cout << "Динамический массив [" << size << "]: ";
    for (size_t i = 0; i < size; i++) {
        cout << "\"" << data[i] << "\"";
        if (i < size - 1) {
            cout << ", ";
        }
    }
    cout << endl;
}