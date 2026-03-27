// benchmark.h
#pragma once

#include <chrono>
#include <iostream>
#include <iomanip>
#include <string>
#include <vector>
#include <random>

class Benchmark {
private:
    struct Result {
        std::string structure;
        long long time_microseconds;
        int operations;
    };
    
    std::vector<Result> results;
    std::mt19937 rng;
    std::uniform_int_distribution<int> dist;
    
public:
    Benchmark();
    
    void runAllTests();
    void printResults() const;
    void saveToFile(const std::string& filename = "benchmark_results.txt") const;
    
private:
    void testArray();
    void testSinglyLinkedList();
    void testDoublyLinkedList();
    void testStack();
    void testQueue();
    void testTree();
    void testHashTable();
    
    template<typename Func>
    long long measureTime(Func func);
    
    std::string generateRandomString(int length = 10);
    int generateRandomNumber();
};