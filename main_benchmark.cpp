// main_benchmark.cpp
#include "benchmark.h"
#include <iostream>

int main() {
    std::cout << "Бенчмарк структур данных\n";
    std::cout << "=========================\n\n";
    
    Benchmark benchmark;
    
    // Запуск всех тестов
    benchmark.runAllTests();
    
    // Вывод результатов
    benchmark.printResults();
    
    // Сохранение в файл
    benchmark.saveToFile();
    
    return 0;
}