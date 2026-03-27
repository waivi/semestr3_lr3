// benchmark.cpp
#include "benchmark.h"
#include "array.h"
#include "SList.h"
#include "DList.h"
#include "Stack.h"
#include "Queue.h"
#include "Tree.h"
#include "Hash_map.h"
#include <fstream>
#include <algorithm>

Benchmark::Benchmark() 
    : rng(std::random_device{}()), 
      dist(1, 10000) {}

std::string Benchmark::generateRandomString(int length) {
    static const char alphanum[] =
        "0123456789"
        "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
        "abcdefghijklmnopqrstuvwxyz";
    
    std::string result;
    result.reserve(length);
    
    for (int i = 0; i < length; ++i) {
        result += alphanum[dist(rng) % (sizeof(alphanum) - 1)];
    }
    
    return result;
}

int Benchmark::generateRandomNumber() {
    return dist(rng);
}

template<typename Func>
long long Benchmark::measureTime(Func func) {
    auto start = std::chrono::high_resolution_clock::now();
    func();
    auto end = std::chrono::high_resolution_clock::now();
    return std::chrono::duration_cast<std::chrono::microseconds>(end - start).count();
}

void Benchmark::testArray() {
    std::cout << "Тестирование массива... ";
    std::cout.flush();
    
    // Убрали параметр filename из конструктора
    MArray arr;
    int operations = 2500;  // Как на картинке
    
    auto testFunc = [&]() {
        // Вставка элементов
        for (int i = 0; i < operations; i++) {
            arr.MADDEND(generateRandomString(5));
        }
        
        // Поиск элементов (по индексу)
        for (int i = 0; i < operations; i++) {
            try {
                if (i % 100 == 0 && arr.MLENGTH() > 0) {
                    arr.MGETINDEX(i % arr.MLENGTH());
                }
            } catch (...) {}
        }
        
        // Удаление элементов
        for (int i = 0; i < operations / 3; i++) {
            if (arr.MLENGTH() > 0) {
                arr.MREMOVEINDEX(arr.MLENGTH() - 1);
            }
        }
    };
    
    long long time = measureTime(testFunc);
    results.push_back({"ARRAY", time, operations});
    std::cout << time << " мкс\n";
}

void Benchmark::testSinglyLinkedList() {
    std::cout << "Тестирование односвязного списка... ";
    std::cout.flush();
    
    // Убрали параметр filename из конструктора
    SList list;
    int operations = 1500;
    
    auto testFunc = [&]() {
        // Вставка в начало
        for (int i = 0; i < operations / 2; i++) {
            list.SLPUSH_HEAD(generateRandomString(5));
        }
        
        // Вставка в конец
        for (int i = 0; i < operations / 2; i++) {
            list.SLPUSH_TAIL(generateRandomString(5));
        }
        
        // Поиск элементов
        for (int i = 0; i < operations / 10; i++) {
            list.SLSEARCH(generateRandomString(5));
        }
    };
    
    long long time = measureTime(testFunc);
    results.push_back({"LNODE", time, operations});
    std::cout << time << " мкс\n";
}

void Benchmark::testDoublyLinkedList() {
    std::cout << "Тестирование двусвязного списка... ";
    std::cout.flush();
    
    // Убрали параметр filename из конструктора
    DList list;
    int operations = 1500;
    
    auto testFunc = [&]() {
        // Вставка в начало
        for (int i = 0; i < operations / 2; i++) {
            list.DLPUSH_HEAD(generateRandomString(5));
        }
        
        // Вставка в конец
        for (int i = 0; i < operations / 2; i++) {
            list.DLPUSH_TAIL(generateRandomString(5));
        }
        
        // Поиск элементов
        for (int i = 0; i < operations / 10; i++) {
            list.DLSEARCH(generateRandomString(5));
        }
    };
    
    long long time = measureTime(testFunc);
    results.push_back({"FNODE", time, operations});
    std::cout << time << " мкс\n";
}

void Benchmark::testStack() {
    std::cout << "Тестирование стека... ";
    std::cout.flush();
    
    // Убрали параметр filename из конструктора
    Stack stack;
    int operations = 7500;
    
    auto testFunc = [&]() {
        // Push операций
        for (int i = 0; i < operations; i++) {
            stack.SPUSH(generateRandomString(5));
        }
        
        // Pop операций
        for (int i = 0; i < operations / 2; i++) {
            try {
                stack.SPOP();
            } catch (...) {
                break;
            }
        }
    };
    
    long long time = measureTime(testFunc);
    results.push_back({"STACK", time, operations});
    std::cout << time << " мкс\n";
}

void Benchmark::testQueue() {
    std::cout << "Тестирование очереди... ";
    std::cout.flush();
    
    // Убрали параметр filename из конструктора
    Queue queue;
    int operations = 7500;
    
    auto testFunc = [&]() {
        // Enqueue операций
        for (int i = 0; i < operations; i++) {
            queue.QPUSH(generateRandomString(5));
        }
        
        // Dequeue операций
        for (int i = 0; i < operations / 2; i++) {
            try {
                queue.QPOP();
            } catch (...) {
                break;
            }
        }
    };
    
    long long time = measureTime(testFunc);
    results.push_back({"QUEUE", time, operations});
    std::cout << time << " мкс\n";
}

void Benchmark::testTree() {
    std::cout << "Тестирование дерева... ";
    std::cout.flush();
    
    RBTree tree;
    int operations = 1500;
    
    auto testFunc = [&]() {
        // Вставка элементов
        for (int i = 0; i < operations; i++) {
            tree.TINSERT(generateRandomString(5));
        }
        
        // Поиск элементов
        for (int i = 0; i < operations / 2; i++) {
            tree.TSEARCH(generateRandomString(5));
        }
        
        // Удаление некоторых элементов
        for (int i = 0; i < operations / 4; i++) {
            tree.TDELETE(generateRandomString(5));
        }
    };
    
    long long time = measureTime(testFunc);
    results.push_back({"TREE", time, operations});
    std::cout << time << " мкс\n";
}

void Benchmark::testHashTable() {
    std::cout << "Тестирование хэш-таблицы... ";
    std::cout.flush();
    
    ChainingHashTable hashTable(100);
    int operations = 2500;
    
    auto testFunc = [&]() {
        // Добавление элементов
        for (int i = 0; i < operations; i++) {
            hashTable.add(std::make_pair(generateRandomNumber(), generateRandomNumber()));
        }
        
        // Поиск элементов
        for (int i = 0; i < operations / 2; i++) {
            hashTable.contains(generateRandomNumber());
        }
        
        // Удаление элементов
        for (int i = 0; i < operations / 4; i++) {
            hashTable.remove(generateRandomNumber());
        }
    };
    
    long long time = measureTime(testFunc);
    results.push_back({"HASH", time, operations});
    std::cout << time << " мкс\n";
}

void Benchmark::runAllTests() {
    std::cout << "Запуск бенчмарков...\n\n";
    
    results.clear();
    
    // Запуск всех тестов в том же порядке, что и на картинке
    testArray();
    testSinglyLinkedList();
    testDoublyLinkedList();
    testStack();
    testQueue();
    testTree();
    testHashTable();
    
    std::cout << "\nВсе тесты завершены\n";
}

void Benchmark::printResults() const {
    std::cout << "\n" << std::string(50, '=') << "\n";
    std::cout << "РЕЗУЛЬТАТЫ БЕНЧМАРКОВ\n";
    std::cout << std::string(50, '=') << "\n\n";
    
    // Заголовок таблицы
    std::cout << std::left << std::setw(15) << "Структура "
              << std::setw(15) << " Время (µs) "
              << std::setw(15) << " Операции"
              << std::setw(15) << " Время/оп (µs)" << "\n";
    
    std::cout << std::string(60, '-') << "\n";
    
    // Данные
    for (const auto& result : results) {
        double time_per_op = static_cast<double>(result.time_microseconds) / result.operations;
        
        std::cout << std::left << std::setw(15) << result.structure
                  << std::setw(15) << result.time_microseconds
                  << std::setw(15) << result.operations
                  << std::setw(15) << std::fixed << std::setprecision(2) << time_per_op << "\n";
    }
    
    // Итоговая таблица в формате с картинки
    std::cout << "\n" << std::string(50, '-') << "\n";
    std::cout << "| Структура    | Время (µs) | Операции |\n";
    std::cout << "|--------------|------------|----------|\n";
    
    // Исправленная карта имен - нужно согласовать с вашими именами
    std::vector<std::pair<std::string, std::string>> nameMap = {
        {"ARRAY", "ARRAY"},      // В бенчмарке: ARRAY, на картинке: ARRAY
        {"LNODE", "LNODE"},      // В бенчмарке: LNODE (односвязный список), на картинке: LNODE
        {"FNODE", "FNODE"},      // В бенчмарке: FNODE (двусвязный список), на картинке: FNODE
        {"STACK", "STACK"},      // В бенчмарке: STACK, на картинке: STACK
        {"QUEUE", "QUEUE"},      // В бенчмарке: QUEUE, на картинке: QUEUE
        {"TREE", "TREE"},        // В бенчмарке: TREE, на картинке: TREE
        {"HASH", "HASH"}         // В бенчмарке: HASH, на картинке: HASH
    };
    
    for (size_t i = 0; i < results.size(); i++) {
        std::cout << "| " << std::left << std::setw(12) << nameMap[i].second
                  << "| " << std::setw(11) << results[i].time_microseconds
                  << "| " << std::setw(9) << results[i].operations << "|\n";
    }
    
    std::cout << std::string(50, '-') << "\n";
}

void Benchmark::saveToFile(const std::string& filename) const {
    std::ofstream file(filename);
    if (!file.is_open()) {
        std::cerr << "Ошибка открытия файла для записи результатов!\n";
        return;
    }
    
    file << "БЕНЧМАРК СТРУКТУР ДАННЫХ\n";
    file << "========================\n\n";
    
    file << std::left << std::setw(15) << "Структура"
         << std::setw(15) << "Время (µs)"
         << std::setw(15) << "Операций"
         << std::setw(15) << "Время/оп (µs)" << "\n";
    
    file << std::string(60, '-') << "\n";
    
    for (const auto& result : results) {
        double time_per_op = static_cast<double>(result.time_microseconds) / result.operations;
        
        file << std::left << std::setw(15) << result.structure
             << std::setw(15) << result.time_microseconds
             << std::setw(15) << result.operations
             << std::setw(15) << std::fixed << std::setprecision(2) << time_per_op << "\n";
    }
    
    file.close();
    std::cout << "Результаты сохранены в файл: " << filename << "\n";
}