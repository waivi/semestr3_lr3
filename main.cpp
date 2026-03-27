#include <iostream>
#include <string>
#include <vector>
#include <sstream>
#include <map>
#include <cstring>
#include "array.h"
#include "SList.h"
#include "DList.h"
#include "Stack.h"
#include "Queue.h"
#include "Tree.h"
#include "Hash_map.h"
#include "serialisation.h"
using namespace std;

void printHelp();
void handleSerializationCommand(const vector<string>& tokens);
// Глобальные объекты для хранения структур данных
map<string, MArray> arrays;
map<string, SList> singleLists;
map<string, DList> doubleLists;
map<string, Stack> stacks;
map<string, Queue> queues;
map<string, RBTree> trees;
map<string, ChainingHashTable> hashTables;

string currentFilename = "data.txt";

// Функция для разбора запроса
vector<string> parseQuery(const string& query) {
    vector<string> tokens;
    stringstream ss(query);
    string token;
    
    while (ss >> token) {
        // Обработка строк в кавычках (двойных или одинарных)
        if (!token.empty() && (token.front() == '"' || token.front() == '\'')) {
            char quoteChar = token.front();
            string quotedStr = token;
            
            // Если закрывающей кавычки нет в первом токене
            if (token.length() == 1 || token.back() != quoteChar) {
                // Собираем оставшуюся часть строки
                string nextToken;
                while (ss >> nextToken) {
                    quotedStr += " " + nextToken;
                    if (!nextToken.empty() && nextToken.back() == quoteChar) break;
                }
            }
            tokens.push_back(quotedStr);
        } else {
            tokens.push_back(token);
        }
    }
    
    
    return tokens;
}

// Функция для удаления кавычек из строки
string removeQuotes(const string& str) {
    if (str.length() >= 2) {
        // Проверяем двойные кавычки
        if (str.front() == '"' && str.back() == '"') {
            return str.substr(1, str.length() - 2);
        }
        // Проверяем одинарные кавычки
        if (str.front() == '\'' && str.back() == '\'') {
            return str.substr(1, str.length() - 2);
        }
    }
    return str;

    //if (str.length() >= 2 && str.front() == '"' && str.back() == '"') {
      //  return str.substr(1, str.length() - 2);
    //}
    //return str;
}


// Функция для обработки команд массивов
void handleArrayCommand(const vector<string>& tokens) {
    if (tokens.size() < 3) {
        cout << "Ошибка: недостаточно параметров для массива" << endl;
        return;
    }
    
    string arrayName = tokens[1];
    string command = tokens[2];
    
    // Создаем массив если его нет
    if (arrays.find(arrayName) == arrays.end()) {
        arrays[arrayName] = MArray();
    }
    
    MArray& arr = arrays[arrayName];
    
    if (command == "MADDEND" && tokens.size() >= 4) {
        arr.MADDEND(removeQuotes(tokens[3]));
    }
    else if (command == "MADDINDEX" && tokens.size() >= 5) {
        size_t index = stoul(tokens[3]);
        arr.MADDINDEX(index, removeQuotes(tokens[4]));
    }
    else if (command == "MREMOVEINDEX" && tokens.size() >= 4) {
        size_t index = stoul(tokens[3]);
        arr.MREMOVEINDEX(index);
    }
    else if (command == "MREPLACEINDEX" && tokens.size() >= 5) {
        size_t index = stoul(tokens[3]);
        arr.MREPLACEINDEX(index, removeQuotes(tokens[4]));
    }
    else if (command == "MGETINDEX" && tokens.size() >= 4) {
        try {
            size_t index = stoul(tokens[3]);
            string value = arr.MGETINDEX(index);
            cout << value << endl;
        } catch (const exception& e) {
            cout << "Ошибка: " << e.what() << endl;
        }
    }
    else if (command == "MLENGTH") {
        cout << arr.MLENGTH() << endl;
    }
    else if (command == "MPRINT") {
        arr.MPRINT();
    }
    else if (command == "MCLEAR") {
        arr.MCLEAR();
    }
    else {
        cout << "Ошибка: неизвестная команда массива или недостаточно параметров" << endl;
    }
}

// Функция для обработки команд односвязных списков
void handleSListCommand(const vector<string>& tokens) {
    if (tokens.size() < 3) {
        cout << "Ошибка: недостаточно параметров для списка" << endl;
        return;
    }
    
    string listName = tokens[1];
    string command = tokens[2];
    
    // Создаем список если его нет
    if (singleLists.find(listName) == singleLists.end()) {
        singleLists[listName] = SList();
    }
    
    SList& list = singleLists[listName];
    
    if (command == "SLPUSH_HEAD" && tokens.size() >= 4) {
        list.SLPUSH_HEAD(removeQuotes(tokens[3]));
    }
    else if (command == "SLPUSH_TAIL" && tokens.size() >= 4) {
        list.SLPUSH_TAIL(removeQuotes(tokens[3]));
    }
    else if (command == "SLPUSH_BEFORE" && tokens.size() >= 5) {
        list.SLPUSH_BEFORE(removeQuotes(tokens[3]), removeQuotes(tokens[4]));
    }
    else if (command == "SLPUSH_AFTER" && tokens.size() >= 5) {
        list.SLPUSH_AFTER(removeQuotes(tokens[3]), removeQuotes(tokens[4]));
    }
    else if (command == "SLREMOVE_HEAD") {
        list.SLREMOVE_HEAD();
    }
    else if (command == "SLREMOVE_TAIL") {
        list.SLREMOVE_TAIL();
    }
    else if (command == "SLREMOVE_BEFORE" && tokens.size() >= 4) {
        list.SLREMOVE_BEFORE(removeQuotes(tokens[3]));
    }
    else if (command == "SLREMOVE_AFTER" && tokens.size() >= 4) {
        list.SLREMOVE_AFTER(removeQuotes(tokens[3]));
    }
    else if (command == "SLREMOVE_VALUE" && tokens.size() >= 4) {
        list.SLREMOVE_VALUE(removeQuotes(tokens[3]));
    }
    else if (command == "SLSEARCH" && tokens.size() >= 4) {
        bool found = list.SLSEARCH(removeQuotes(tokens[3]));
        cout << (found ? "Найдено" : "Не найдено") << endl;
    }
    else if (command == "SLGET" && tokens.size() >= 4) {
        try {
            size_t index = stoul(tokens[3]);
            string value = list.SLGET(index);
            cout << value << endl;
        } catch (const exception& e) {
            cout << "Ошибка: " << e.what() << endl;
        }
    }
    else if (command == "SLPRINT_FORWARD") {
        list.SLPRINT_FORWARD();
    }
    else if (command == "SLPRINT_BACKWARD") {
        list.SLPRINT_BACKWARD();
    }
    else if (command == "SLLENGTH") {
        cout << list.SLLENGTH() << endl;
    }
    else {
        cout << "Ошибка: неизвестная команда списка или недостаточно параметров" << endl;
    }
}

// Функция для обработки команд двусвязных списков
void handleDListCommand(const vector<string>& tokens) {
    if (tokens.size() < 3) {
        cout << "Ошибка: недостаточно параметров для двусвязного списка" << endl;
        return;
    }
    
    string listName = tokens[1];
    string command = tokens[2];
    
    // Создаем список если его нет
    if (doubleLists.find(listName) == doubleLists.end()) {
        doubleLists[listName] = DList();
    }
    
    DList& list = doubleLists[listName];
    
    if (command == "DLPUSH_HEAD" && tokens.size() >= 4) {
        list.DLPUSH_HEAD(removeQuotes(tokens[3]));
    }
    else if (command == "DLPUSH_TAIL" && tokens.size() >= 4) {
        list.DLPUSH_TAIL(removeQuotes(tokens[3]));
    }
    else if (command == "DLPUSH_BEFORE" && tokens.size() >= 5) {
        list.DLPUSH_BEFORE(removeQuotes(tokens[3]), removeQuotes(tokens[4]));
    }
    else if (command == "DLPUSH_AFTER" && tokens.size() >= 5) {
        list.DLPUSH_AFTER(removeQuotes(tokens[3]), removeQuotes(tokens[4]));
    }
    else if (command == "DLREMOVE_HEAD") {
        list.DLREMOVE_HEAD();
    }
    else if (command == "DLREMOVE_TAIL") {
        list.DLREMOVE_TAIL();
    }
    else if (command == "DLREMOVE_BEFORE" && tokens.size() >= 4) {
        list.DLREMOVE_BEFORE(removeQuotes(tokens[3]));
    }
    else if (command == "DLREMOVE_AFTER" && tokens.size() >= 4) {
        list.DLREMOVE_AFTER(removeQuotes(tokens[3]));
    }
    else if (command == "DLREMOVE_VALUE" && tokens.size() >= 4) {
        list.DLREMOVE_VALUE(removeQuotes(tokens[3]));
    }
    else if (command == "DLSEARCH" && tokens.size() >= 4) {
        bool found = list.DLSEARCH(removeQuotes(tokens[3]));
        cout << (found ? "Найдено" : "Не найдено") << endl;
    }
    else if (command == "DLGET" && tokens.size() >= 4) {
        try {
            size_t index = stoul(tokens[3]);
            string value = list.DLGET(index);
            cout << value << endl;
        } catch (const exception& e) {
            cout << "Ошибка: " << e.what() << endl;
        }
    }
    else if (command == "DLPRINT_FORWARD") {
        list.DLPRINT_FORWARD();
    }
    else if (command == "DLPRINT_BACKWARD") {
        list.DLPRINT_BACKWARD();
    }
    else if (command == "DLLENGTH") {
        cout << list.DLLENGTH() << endl;
    }
    else {
        cout << "Ошибка: неизвестная команда двусвязного списка или недостаточно параметров" << endl;
    }
}

// Функция для обработки команд стеков
void handleStackCommand(const vector<string>& tokens) {
    if (tokens.size() < 3) {
        cout << "Ошибка: недостаточно параметров для стека" << endl;
        return;
    }
    
    string stackName = tokens[1];
    string command = tokens[2];
    
    // Создаем стек если его нет
    if (stacks.find(stackName) == stacks.end()) {
        stacks[stackName] = Stack();
    }
    
    Stack& stack = stacks[stackName];
    
    if (command == "SPUSH" && tokens.size() >= 4) {
        stack.SPUSH(removeQuotes(tokens[3]));
    }
    else if (command == "SPOP") {
        try {
            string value = stack.SPOP();
            cout << value << endl;
        } catch (const exception& e) {
            cout << "Ошибка: " << e.what() << endl;
        }
    }
    else if (command == "STOP") {
        try {
            string value = stack.STOP();
            cout << value << endl;
        } catch (const exception& e) {
            cout << "Ошибка: " << e.what() << endl;
        }
    }
    else if (command == "SPRINT") {
        stack.SPRINT();
    }
    else {
        cout << "Ошибка: неизвестная команда стека или недостаточно параметров" << endl;
    }
}

// Функция для обработки команд очередей
void handleQueueCommand(const vector<string>& tokens) {
    if (tokens.size() < 3) {
        cout << "Ошибка: недостаточно параметров для очереди" << endl;
        return;
    }
    
    string queueName = tokens[1];
    string command = tokens[2];
    
    // Создаем очередь если ее нет
    if (queues.find(queueName) == queues.end()) {
        queues[queueName] = Queue();
    }
    
    Queue& queue = queues[queueName];
    
    if (command == "QPUSH" && tokens.size() >= 4) {
        queue.QPUSH(removeQuotes(tokens[3]));
    }
    else if (command == "QPOP") {
        try {
            string value = queue.QPOP();
            cout << value << endl;
        } catch (const exception& e) {
            cout << "Ошибка: " << e.what() << endl;
        }
    }
    else if (command == "QFRONT") {
        try {
            string value = queue.QFRONT();
            cout << value << endl;
        } catch (const exception& e) {
            cout << "Ошибка: " << e.what() << endl;
        }
    }
    else if (command == "QPRINT") {
        queue.QPRINT();
    }
    else {
        cout << "Ошибка: неизвестная команда очереди или недостаточно параметров" << endl;
    }
}

// Функция для обработки команд деревьев
void handleTreeCommand(const vector<string>& tokens) {
    cout << "DEBUG handleTreeCommand: всего токенов = " << tokens.size() << endl;
    for (size_t i = 0; i < tokens.size(); i++) {
        cout << "  tokens[" << i << "] = '" << tokens[i] << "'" << endl;
    }
    if (tokens.size() < 3) {
        cout << "Ошибка: недостаточно параметров для дерева" << endl;
        return;
    }
    
    string treeName = tokens[1];
    string command = tokens[2];
    
    // Создаем дерево если его нет
    if (trees.find(treeName) == trees.end()) {
        trees[treeName] = RBTree();
    }
    
    RBTree& tree = trees[treeName];
    
    if (command == "TINSERT" && tokens.size() >= 4) {
        tree.TINSERT(removeQuotes(tokens[3]));
    }
    else if (command == "TDELETE" && tokens.size() >= 4) {
        tree.TDELETE(removeQuotes(tokens[3]));
    }
    else if (command == "TSEARCH" && tokens.size() >= 4) {
        bool found = tree.TSEARCH(removeQuotes(tokens[3]));
        cout << (found ? "Найдено" : "Не найдено") << endl;
    }
    else if (command == "TGET" && tokens.size() >= 4) {
        string value = tree.TGET(removeQuotes(tokens[3]));
        if (!value.empty()) {
            cout << value << endl;
        } else {
            cout << "Элемент не найден" << endl;
        }
    }
    else if (command == "TPRINT_INORDER") {
        tree.TPRINT_INORDER();
    }
    else if (command == "TPRINT_PREORDER") {
        tree.TPRINT_PREORDER();
    }
    else if (command == "TPRINT_POSTORDER") {
        tree.TPRINT_POSTORDER();
    }
    else if (command == "TPRINT_TREE") {
        tree.TPRINT_TREE();
    }
    else {
        cout << "Ошибка: неизвестная команда дерева или недостаточно параметров" << endl;
    }
}

// Функция для обработки команд хэш-таблиц
void handleHashTableCommand(const vector<string>& tokens) {
    if (tokens.size() < 4) {
        cout << "Ошибка: недостаточно параметров для хэш-таблицы" << endl;
        cout << "Использование: HSET <имя> HADD <ключ> <значение>" << endl;
        cout << "            или HSET <имя> HGET <ключ>" << endl;
        return;
    }
    
    string hashName = tokens[1];
    string command = tokens[2];
    
    // Создаем хэш-таблицу если ее нет (емкость по умолчанию 10)
    if (hashTables.find(hashName) == hashTables.end()) {
        hashTables.emplace(hashName, 10);
    }
    
    ChainingHashTable& hashTable = hashTables[hashName];
    
    if (command == "HADD" && tokens.size() >= 5) {
        try {
            int key = stoi(tokens[3]);
            int value = stoi(tokens[4]);
            hashTable.add(make_pair(key, value));
            cout << "Добавлено: ключ=" << key << ", значение=" << value << endl;
        } catch (const exception& e) {
            cout << "Ошибка: ключ и значение должны быть целыми числами" << endl;
        }
    }
    else if (command == "HREMOVE" && tokens.size() >= 4) {
        try {
            int key = stoi(tokens[3]);
            hashTable.remove(key);
            cout << "Удален ключ: " << key << endl;
        } catch (const exception& e) {
            cout << "Ошибка: ключ должен быть целым числом" << endl;
        }
    }
    else if (command == "HGET" && tokens.size() >= 4) {
        try {
            int key = stoi(tokens[3]);
            auto result = hashTable.contains(key);
            if (result.first) {
                cout << "Найдено: ключ=" << key << ", значение=" << result.second << endl;
            } else {
                cout << "Ключ " << key << " не найден" << endl;
            }
        } catch (const exception& e) {
            cout << "Ошибка: ключ должен быть целым числом" << endl;
        }
    }
    else if (command == "HPRINT") {
        cout << hashTable.toString();
    }
    else if (command == "HSIZE") {
        cout << "Размер: " << hashTable.getSize() << endl;
        cout << "Емкость: " << hashTable.getCapacity() << endl;
        cout << "Коэффициент загрузки: " << hashTable.getLoadFactor() << endl;
    }
    else if (command == "HSTATS") {
        int minLen, maxLen;
        double avgLen;
        hashTable.getChainLengths(minLen, maxLen, avgLen);
        cout << "Статистика цепочек:" << endl;
        cout << "  Минимальная длина: " << minLen << endl;
        cout << "  Максимальная длина: " << maxLen << endl;
        cout << "  Средняя длина: " << avgLen << endl;
        cout << "  Коэффициент загрузки: " << hashTable.getLoadFactor() << endl;
    }
    else if (command == "HCLEAR") {
        hashTable.clear();
        cout << "Хэш-таблица очищена" << endl;
    }
    else if (command == "HEMPTY") {
        cout << (hashTable.isEmpty() ? "Пустая" : "Не пустая") << endl;
    }
    else {
        cout << "Ошибка: неизвестная команда хэш-таблицы" << endl;
        cout << "Доступные команды: HADD, HREMOVE, HGET, HPRINT, HSIZE, HSTATS, HCLEAR, HEMPTY" << endl;
    }
}

// Функция для обработки команд сериализации
void handleSerializationCommand(const vector<string>& tokens) {
    if (tokens.size() < 4) {
        cout << "Ошибка: недостаточно параметров для сериализации" << endl;
        cout << "Использование: SAVE <тип> <имя> <формат> <файл>" << endl;
        cout << "            или LOAD <тип> <имя> <формат> <файл>" << endl;
        cout << "Типы: ARRAY, SLIST, DLIST, STACK, QUEUE, TREE, HASHTABLE" << endl;
        cout << "Форматы: TEXT, BINARY" << endl;
        return;
    }
    
    string command = tokens[0];
    string type = tokens[1];
    string name = tokens[2];
    string format = tokens[3];
    string filename;
    
    if (tokens.size() >= 5) {
        filename = tokens[4];
    } else {
        // Генерируем имя файла по умолчанию
        filename = name + "_" + (command == "SAVE" ? "saved" : "loaded") + 
                   (format == "BINARY" ? ".bin" : ".txt");
    }
    
    if (command == "SAVE") {
        if (type == "ARRAY") {
            if (arrays.find(name) == arrays.end()) {
                cout << "Ошибка: массив '" << name << "' не найден" << endl;
                return;
            }
            if (format == "TEXT") {
                saveToText(arrays[name], filename);
                cout << "Массив '" << name << "' сохранен в текстовый файл: " << filename << endl;
            } else if (format == "BINARY") {
                saveToBinary(arrays[name], filename);
                cout << "Массив '" << name << "' сохранен в бинарный файл: " << filename << endl;
            }
        }
        else if (type == "SLIST") {
            if (singleLists.find(name) == singleLists.end()) {
                cout << "Ошибка: список '" << name << "' не найден" << endl;
                return;
            }
            if (format == "TEXT") {
                saveToText(singleLists[name], filename);
                cout << "Список '" << name << "' сохранен в текстовый файл: " << filename << endl;
            } else if (format == "BINARY") {
                saveToBinary(singleLists[name], filename);
                cout << "Список '" << name << "' сохранен в бинарный файл: " << filename << endl;
            }
        }
        else if (type == "DLIST") {
            if (doubleLists.find(name) == doubleLists.end()) {
                cout << "Ошибка: двусвязный список '" << name << "' не найден" << endl;
                return;
            }
            if (format == "TEXT") {
                saveToText(doubleLists[name], filename);
                cout << "Двусвязный список '" << name << "' сохранен в текстовый файл: " << filename << endl;
            } else if (format == "BINARY") {
                saveToBinary(doubleLists[name], filename);
                cout << "Двусвязный список '" << name << "' сохранен в бинарный файл: " << filename << endl;
            }
        }
        else if (type == "STACK") {
            if (stacks.find(name) == stacks.end()) {
                cout << "Ошибка: стек '" << name << "' не найден" << endl;
                return;
            }
            if (format == "TEXT") {
                saveToText(stacks[name], filename);
                cout << "Стек '" << name << "' сохранен в текстовый файл: " << filename << endl;
            } else if (format == "BINARY") {
                saveToBinary(stacks[name], filename);
                cout << "Стек '" << name << "' сохранен в бинарный файл: " << filename << endl;
            }
        }
        else if (type == "QUEUE") {
            if (queues.find(name) == queues.end()) {
                cout << "Ошибка: очередь '" << name << "' не найден" << endl;
                return;
            }
            if (format == "TEXT") {
                saveToText(queues[name], filename);
                cout << "Очередь '" << name << "' сохранена в текстовый файл: " << filename << endl;
            } else if (format == "BINARY") {
                saveToBinary(queues[name], filename);
                cout << "Очередь '" << name << "' сохранена в бинарный файл: " << filename << endl;
            }
        }
        else if (type == "TREE") {
            if (trees.find(name) == trees.end()) {
                cout << "Ошибка: дерево '" << name << "' не найдено" << endl;
                return;
            }
            if (format == "TEXT") {
                saveToText(trees[name], filename);
                cout << "Дерево '" << name << "' сохранено в текстовый файл: " << filename << endl;
            } else if (format == "BINARY") {
                saveToBinary(trees[name], filename);
                cout << "Дерево '" << name << "' сохранено в бинарный файл: " << filename << endl;
            }
        }
        else if (type == "HASHTABLE") {
            if (hashTables.find(name) == hashTables.end()) {
                cout << "Ошибка: хэш-таблица '" << name << "' не найдена" << endl;
                return;
            }
            if (format == "TEXT") {
                saveToText(hashTables[name], filename);
                cout << "Хэш-таблица '" << name << "' сохранена в текстовый файл: " << filename << endl;
            } else if (format == "BINARY") {
                saveToBinary(hashTables[name], filename);
                cout << "Хэш-таблица '" << name << "' сохранена в бинарный файл: " << filename << endl;
            }
        }
        else {
            cout << "Ошибка: неизвестный тип для сериализации: " << type << endl;
        }
    }
    else if (command == "LOAD") {
        if (type == "ARRAY") {
            if (format == "TEXT") {
                loadFromText(arrays[name], filename);
                cout << "Массив загружен из текстового файла: " << filename << endl;
            } else if (format == "BINARY") {
                loadFromBinary(arrays[name], filename);
                cout << "Массив загружен из бинарного файла: " << filename << endl;
            }
        }
        else if (type == "SLIST") {
            if (format == "TEXT") {
                loadFromText(singleLists[name], filename);
                cout << "Список загружен из текстового файла: " << filename << endl;
            } else if (format == "BINARY") {
                loadFromBinary(singleLists[name], filename);
                cout << "Список загружен из бинарного файла: " << filename << endl;
            }
        }
        else if (type == "DLIST") {
            if (format == "TEXT") {
                loadFromText(doubleLists[name], filename);
                cout << "Двусвязный список загружен из текстового файла: " << filename << endl;
            } else if (format == "BINARY") {
                loadFromBinary(doubleLists[name], filename);
                cout << "Двусвязный список загружен из бинарного файла: " << filename << endl;
            }
        }
        else if (type == "STACK") {
            if (format == "TEXT") {
                loadFromText(stacks[name], filename);
                cout << "Стек загружен из текстового файла: " << filename << endl;
            } else if (format == "BINARY") {
                loadFromBinary(stacks[name], filename);
                cout << "Стек загружен из бинарного файла: " << filename << endl;
            }
        }
        else if (type == "QUEUE") {
            if (format == "TEXT") {
                loadFromText(queues[name], filename);
                cout << "Очередь загружена из текстового файла: " << filename << endl;
            } else if (format == "BINARY") {
                loadFromBinary(queues[name], filename);
                cout << "Очередь загружена из бинарного файла: " << filename << endl;
            }
        }
        else if (type == "TREE") {
            if (format == "TEXT") {
                loadFromText(trees[name], filename);
                cout << "Дерево загружено из текстового файла: " << filename << endl;
            } else if (format == "BINARY") {
                loadFromBinary(trees[name], filename);
                cout << "Дерево загружено из бинарного файла: " << filename << endl;
            }
        }
        else if (type == "HASHTABLE") {
            if (format == "TEXT") {
                loadFromText(hashTables[name], filename);
                cout << "Хэш-таблица загружена из текстового файла: " << filename << endl;
            } else if (format == "BINARY") {
                loadFromBinary(hashTables[name], filename);
                cout << "Хэш-таблица загружена из бинарного файла: " << filename << endl;
            }
        }
        else {
            cout << "Ошибка: неизвестный тип для загрузки: " << type << endl;
        }
    }
    else {
        cout << "Ошибка: неизвестная команда сериализации: " << command << endl;
    }
}

// Функция для обработки запроса
void processQuery(const string& query) {
    cout << "DEBUG processQuery: входная строка = '" << query << "'" << endl;
    vector<string> tokens = parseQuery(query);
    cout << "DEBUG: Получено токенов: " << tokens.size() << endl;
    for (size_t i = 0; i < tokens.size(); i++) {
        cout << "  tokens[" << i << "] = '" << tokens[i] << "'" << endl;
    }
    if (tokens.empty()) {
        cout << "Ошибка: пустой запрос" << endl;
        return;
    }
    
    string dataType = tokens[0];
    
    if (dataType == "MARRAY") {
        handleArrayCommand(tokens);
    }
    else if (dataType == "SLIST") {
        handleSListCommand(tokens);
    }
    else if (dataType == "DLIST") {
        handleDListCommand(tokens);
    }
    else if (dataType == "STACK") {
        handleStackCommand(tokens);
    }
    else if (dataType == "QUEUE") {
        handleQueueCommand(tokens);
    }
    else if (dataType == "TREE") {
        handleTreeCommand(tokens);
    }
    else if (dataType == "HSET") {
        handleHashTableCommand(tokens);
    }
    else if (dataType == "HGET") {
        handleHashTableCommand(tokens);
    }
    else if (dataType == "SAVE" || dataType == "LOAD") {
        handleSerializationCommand(tokens);
    }
    else if (dataType == "HELP") {
        printHelp();
    }
    else if (dataType == "CLEAR") {
        arrays.clear();
        singleLists.clear();
        doubleLists.clear();
        stacks.clear();
        queues.clear();
        trees.clear();
        hashTables.clear();
        cout << "Все структуры данных очищены" << endl;
    }
    else if (dataType == "LIST") {
        cout << "Доступные структуры:" << endl;
        cout << "Массивы: ";
        for (const auto& pair : arrays) cout << pair.first << " ";
        cout << endl << "Односвязные списки: ";
        for (const auto& pair : singleLists) cout << pair.first << " ";
        cout << endl << "Двусвязные списки: ";
        for (const auto& pair : doubleLists) cout << pair.first << " ";
        cout << endl << "Стеки: ";
        for (const auto& pair : stacks) cout << pair.first << " ";
        cout << endl << "Очереди: ";
        for (const auto& pair : queues) cout << pair.first << " ";
        cout << endl << "Деревья: ";
        for (const auto& pair : trees) cout << pair.first << " ";
        cout << endl << "Хэш-таблицы: ";
        for (const auto& pair : hashTables) cout << pair.first << " ";
        cout << endl;
    }
    else {
        cout << "Ошибка: неизвестный тип данных: " << dataType << endl;
        cout << "Доступные команды: MARRAY, SLIST, DLIST, STACK, QUEUE, TREE, HSET, HGET, SAVE, LOAD, HELP, CLEAR, LIST" << endl;
    }
}

// Функция для вывода справки
void printHelp() {
    cout << "==================================================" << endl;
    cout << "           СИСТЕМА УПРАВЛЕНИЯ ДАННЫМИ" << endl;
    cout << "==================================================" << endl;
    cout << endl;
    cout << "ОСНОВНЫЕ КОМАНДЫ:" << endl;
    cout << "  HELP           - Показать эту справку" << endl;
    cout << "  CLEAR          - Очистить все структуры данных" << endl;
    cout << "  LIST           - Показать все доступные структуры" << endl;
    cout << endl;
    cout << "МАССИВЫ (MARRAY):" << endl;
    cout << "  MARRAY <имя> MADDEND <значение>" << endl;
    cout << "  MARRAY <имя> MADDINDEX <индекс> <значение>" << endl;
    cout << "  MARRAY <имя> MGETINDEX <индекс>" << endl;
    cout << "  MARRAY <имя> MPRINT" << endl;
    cout << "  MARRAY <имя> MLENGTH" << endl;
    cout << endl;
    cout << "ОДНОСВЯЗНЫЕ СПИСКИ (SLIST):" << endl;
    cout << "  SLIST <имя> SLPUSH_HEAD <значение>" << endl;
    cout << "  SLIST <имя> SLPUSH_TAIL <значение>" << endl;
    cout << "  SLIST <имя> SLSEARCH <значение>" << endl;
    cout << "  SLIST <имя> SLPRINT_FORWARD" << endl;
    cout << "  SLIST <имя> SLLENGTH" << endl;
    cout << endl;
    cout << "ДВУСВЯЗНЫЕ СПИСКИ (DLIST):" << endl;
    cout << "  DLIST <имя> DLPUSH_HEAD <значение>" << endl;
    cout << "  DLIST <имя> DLPUSH_TAIL <значение>" << endl;
    cout << "  DLIST <имя> DLPRINT_FORWARD" << endl;
    cout << "  DLIST <имя> DLPRINT_BACKWARD" << endl;
    cout << endl;
    cout << "СТЕКИ (STACK):" << endl;
    cout << "  STACK <имя> SPUSH <значение>" << endl;
    cout << "  STACK <имя> SPOP" << endl;
    cout << "  STACK <имя> SPRINT" << endl;
    cout << endl;
    cout << "ОЧЕРЕДИ (QUEUE):" << endl;
    cout << "  QUEUE <имя> QPUSH <значение>" << endl;
    cout << "  QUEUE <имя> QPOP" << endl;
    cout << "  QUEUE <имя> QPRINT" << endl;
    cout << endl;
    cout << "ДЕРЕВЬЯ (TREE):" << endl;
    cout << "  TREE <имя> TINSERT <значение>" << endl;
    cout << "  TREE <имя> TSEARCH <значение>" << endl;
    cout << "  TREE <имя> TPRINT_INORDER" << endl;
    cout << "  TREE <имя> TPRINT_TREE" << endl;
    cout << endl;
    cout << "ХЭШ-ТАБЛИЦЫ (HSET/HGET):" << endl;
    cout << "  HSET <имя> HADD <ключ> <значение>" << endl;
    cout << "  HSET <имя> HGET <ключ>" << endl;
    cout << "  HSET <имя> HPRINT" << endl;
    cout << "  HSET <имя> HSIZE" << endl;
    cout << endl;
    cout << "СЕРИАЛИЗАЦИЯ И ДЕСЕРИАЛИЗАЦИЯ:" << endl;
    cout << "  SAVE <тип> <имя> <формат> [файл]" << endl;
    cout << "  LOAD <тип> <имя> <формат> [файл]" << endl;
    cout << "  Типы: ARRAY, SLIST, DLIST, STACK, QUEUE, TREE, HASHTABLE, ALL" << endl;
    cout << "  Форматы: TEXT, BINARY" << endl;
    cout << endl;
    cout << "ПРИМЕРЫ СЕРИАЛИЗАЦИИ:" << endl;
    cout << "  SAVE ARRAY myarray TEXT myarray.txt" << endl;
    cout << "  LOAD ARRAY newarray BINARY myarray.bin" << endl;
    cout << "  SAVE TREE mytree BINARY tree_data.bin" << endl;
    cout << "  SAVE ALL project TEXT project_data" << endl;
    cout << "  LOAD ALL project BINARY project_data" << endl;
    cout << endl;
    cout << "==================================================" << endl;
}

int main(int argc, char* argv[]) {
    string query;
    string filename = "data.txt";
    
    // Разбор аргументов командной строки
    for (int i = 1; i < argc; i++) {
        string arg = argv[i];
        
        if (arg == "--file" && i + 1 < argc) {
            filename = argv[++i];
            currentFilename = filename;
        }
        else if (arg == "--query" && i + 1 < argc) {
            query = argv[++i];
        }
        else if (arg == "--help" || arg == "-h") {
            printHelp();
            return 0;
        }
        else if (arg == "--interactive" || arg == "-i") {
            // Интерактивный режим
            cout << "Режим интерактивной сериализации" << endl;
            cout << "Введите команды (или 'exit' для выхода):" << endl;
            
            string line;
            while (true) {
                cout << "> ";
                getline(cin, line);
                
                if (line == "exit" || line == "quit") {
                    break;
                }
                if (line == "help") {
                    printHelp();
                    continue;
                }
                
                try {
                    processQuery(line);
                } catch (const exception& e) {
                    cout << "Ошибка: " << e.what() << endl;
                }
            }
            return 0;
        }
        else if (arg == "--test") {
            // Автоматический тест сериализации
            cout << "Запуск автоматического теста сериализации..." << endl;
            
            // Создаем и тестируем все структуры
            cout << "\n=== ТЕСТ МАССИВА ===" << endl;
            MArray testArray;
            testArray.MADDEND("Элемент1");
            testArray.MADDEND("Элемент2");
            testArray.MADDEND("Элемент3");
            cout << "Исходный массив: ";
            testArray.MPRINT();
            
            saveToText(testArray, "test_array.txt");
            cout << "Сохранен в test_array.txt" << endl;
            
            MArray loadedArray;
            loadFromText(loadedArray, "test_array.txt");
            cout << "Загруженный массив: ";
            loadedArray.MPRINT();
            
            // Тест бинарной сериализации
            saveToBinary(testArray, "test_array.bin");
            MArray binArray;
            loadFromBinary(binArray, "test_array.bin");
            cout << "Бинарно загруженный массив: ";
            binArray.MPRINT();
            
            cout << "\n=== ТЕСТ ЗАВЕРШЕН ===" << endl;
            return 0;
        }
    }
    
    if (query.empty()) {
        cout << "Использование: ./dbms [опции]" << endl;
        cout << "Опции:" << endl;
        cout << "  --query 'команда'     - Выполнить одну команду" << endl;
        cout << "  --interactive         - Интерактивный режим" << endl;
        cout << "  --test                - Запустить автоматический тест" << endl;
        cout << "  --help                - Показать справку" << endl;
        return 1;
    }
    
    cout << "Файл: " << filename << endl;
    cout << "Запрос: " << query << endl;
    cout << "Результат: ";
    
    try {
        processQuery(query);
    } catch (const exception& e) {
        cout << "Ошибка выполнения: " << e.what() << endl;
        return 1;
    }
    
    return 0;
}