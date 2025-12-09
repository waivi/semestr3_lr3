#define BOOST_TEST_MODULE DataStructuresExtendedTest
#include <boost/test/included/unit_test.hpp>
#include <boost/test/data/test_case.hpp>
#include <boost/test/data/monomorphic.hpp>
#include <boost/test/tools/output_test_stream.hpp>

#include "array.h"
#include "SList.h"
#include "DList.h"
#include "Stack.h"
#include "Queue.h"
#include "Tree.h"
#include "Hash_map.h"
#include "serialisation.h"

#include <fstream>
#include <random>
#include <chrono>
#include <sstream>
#include <algorithm>

BOOST_AUTO_TEST_SUITE(FullCoverageTests)
// MArray 

BOOST_AUTO_TEST_CASE(MArray_EdgeCases) {
    MArray arr;
    
    // Тест пустого массива
    BOOST_TEST(arr.MLENGTH() == 0);
    
    // Тест добавления большого количества элементов
    for(int i = 0; i < 1000; i++) {
        arr.MADDEND("element_" + std::to_string(i));
    }
    BOOST_TEST(arr.MLENGTH() == 1000);
    
    // Тест получения всех элементов
    for(int i = 0; i < 1000; i++) {
        BOOST_TEST(arr.MGETINDEX(i) == "element_" + std::to_string(i));
    }
    
    // Тест добавления по индексу в начало
    arr.MADDINDEX(0, "new_first");
    BOOST_TEST(arr.MGETINDEX(0) == "new_first");
    BOOST_TEST(arr.MLENGTH() == 1001);
    
    // Тест добавления по индексу в середину
    arr.MADDINDEX(500, "middle_element");
    BOOST_TEST(arr.MGETINDEX(500) == "middle_element");
    
    // Тест добавления по индексу в конец
    arr.MADDINDEX(arr.MLENGTH(), "new_last");
    BOOST_TEST(arr.MGETINDEX(arr.MLENGTH() - 1) == "new_last");
    
    // Тест замены элементов
    arr.MREPLACEINDEX(0, "replaced_first");
    BOOST_TEST(arr.MGETINDEX(0) == "replaced_first");
    
    arr.MREPLACEINDEX(500, "replaced_middle");
    BOOST_TEST(arr.MGETINDEX(500) == "replaced_middle");
    
    arr.MREPLACEINDEX(arr.MLENGTH() - 1, "replaced_last");
    BOOST_TEST(arr.MGETINDEX(arr.MLENGTH() - 1) == "replaced_last");
    
    // Тест удаления из разных позиций
    size_t initialLength = arr.MLENGTH();
    
    // Удаление из начала
    arr.MREMOVEINDEX(0);
    BOOST_TEST(arr.MLENGTH() == initialLength - 1);
    
    // Удаление из середины
    arr.MREMOVEINDEX(250);
    BOOST_TEST(arr.MLENGTH() == initialLength - 2);
    
    // Удаление из конца
    arr.MREMOVEINDEX(arr.MLENGTH() - 1);
    BOOST_TEST(arr.MLENGTH() == initialLength - 3);
    
    // Тест метода MPRINT (не проверяем вывод, только что не падает)
    arr.MPRINT();
    
    // Тест очистки
    arr.MCLEAR();
    BOOST_TEST(arr.MLENGTH() == 0);
    
    // Тест повторного добавления после очистки
    arr.MADDEND("after_clear");
    BOOST_TEST(arr.MLENGTH() == 1);
    BOOST_TEST(arr.MGETINDEX(0) == "after_clear");
}

BOOST_AUTO_TEST_CASE(MArray_ExceptionHandling) {
    MArray arr;
    
    // Исключения при пустом массиве
    BOOST_CHECK_THROW(arr.MGETINDEX(0), std::out_of_range);
    BOOST_CHECK_THROW(arr.MGETINDEX(100), std::out_of_range);
    
    // Добавляем элементы
    arr.MADDEND("one");
    arr.MADDEND("two");
    arr.MADDEND("three");
    
    // Валидные индексы не должны кидать исключения
    BOOST_CHECK_NO_THROW(arr.MGETINDEX(0));
    BOOST_CHECK_NO_THROW(arr.MGETINDEX(1));
    BOOST_CHECK_NO_THROW(arr.MGETINDEX(2));
    
    // Невалидные индексы должны кидать исключения
    BOOST_CHECK_THROW(arr.MGETINDEX(3), std::out_of_range);
    BOOST_CHECK_THROW(arr.MGETINDEX(100), std::out_of_range);
    
    // Удаление с невалидным индексом (не должно кидать исключение, но и не должно падать)
    BOOST_CHECK_NO_THROW(arr.MREMOVEINDEX(10)); // Должно вывести сообщение об ошибке
    
    // Замена с невалидным индексом
    BOOST_CHECK_NO_THROW(arr.MREPLACEINDEX(10, "test"));
    
    // Добавление с индексом больше размера
    BOOST_CHECK_NO_THROW(arr.MADDINDEX(10, "test"));
}

// SList 

BOOST_AUTO_TEST_CASE(SList_FullOperations) {
    SList list;
    
    // Тест пустого списка
    BOOST_TEST(list.SLLENGTH() == 0);
    
    // Тест всех методов добавления
    list.SLPUSH_HEAD("head");
    BOOST_TEST(list.SLLENGTH() == 1);
    BOOST_TEST(list.SLGET(0) == "head");
    
    list.SLPUSH_TAIL("tail");
    BOOST_TEST(list.SLLENGTH() == 2);
    BOOST_TEST(list.SLGET(1) == "tail");
    
    list.SLPUSH_BEFORE("tail", "before_tail");
    BOOST_TEST(list.SLLENGTH() == 3);
    BOOST_TEST(list.SLSEARCH("before_tail") == true);
    
    list.SLPUSH_AFTER("head", "after_head");
    BOOST_TEST(list.SLLENGTH() == 4);
    
    // Проверяем порядок
    BOOST_TEST(list.SLGET(0) == "head");
    BOOST_TEST(list.SLGET(1) == "after_head");
    
    // Тест поиска несуществующих элементов
    BOOST_TEST(list.SLSEARCH("non_existent") == false);
    
    // Тест всех методов удаления
    list.SLREMOVE_HEAD();
    BOOST_TEST(list.SLLENGTH() == 3);
    BOOST_TEST(list.SLGET(0) != "head");
    
    list.SLREMOVE_TAIL();
    BOOST_TEST(list.SLLENGTH() == 2);
    
    list.SLPUSH_HEAD("new_head");
    list.SLPUSH_TAIL("new_tail");
    list.SLPUSH_TAIL("target");
    list.SLPUSH_TAIL("after_target");
    
    list.SLREMOVE_BEFORE("target");
    BOOST_TEST(list.SLSEARCH("new_tail") == false);
    
    list.SLREMOVE_AFTER("target");
    BOOST_TEST(list.SLLENGTH() == 4);
    
    list.SLREMOVE_VALUE("target");
    BOOST_TEST(list.SLSEARCH("target") == false);
    BOOST_TEST(list.SLLENGTH() == 3);
    
    // Тест получения по индексу
    BOOST_CHECK_NO_THROW(list.SLGET(0));
    BOOST_CHECK_NO_THROW(list.SLGET(list.SLLENGTH() - 1));
    BOOST_CHECK_THROW(list.SLGET(list.SLLENGTH()), std::out_of_range);
    
    // Тест методов вывода
    list.SLPRINT_FORWARD();
    list.SLPRINT_BACKWARD();
    list.SLPRINT_HEAD();
    list.SLPRINT_TAIL();
    
    // Тест очистки
    list.SLCLEAR();
    BOOST_TEST(list.SLLENGTH() == 0);
    
    // Тест работы с пустым списком после очистки
    BOOST_CHECK_NO_THROW(list.SLREMOVE_HEAD()); // Должно вывести сообщение
    BOOST_CHECK_NO_THROW(list.SLREMOVE_TAIL());
    BOOST_TEST(list.SLSEARCH("anything") == false);
}
BOOST_AUTO_TEST_CASE(SList_EdgeCase_BranchCoverage) {
    SList list;
    
    // 1. Тест всех ветвей в SLPUSH_BEFORE
    // Пустой список + элемент не найден
    list.SLPUSH_BEFORE("non_existent", "test1");
    BOOST_TEST(list.SLLENGTH() == 0);
    
    // Добавляем элемент и тестируем перед ним
    list.SLPUSH_HEAD("target");
    list.SLPUSH_BEFORE("target", "before_target");
    BOOST_TEST(list.SLLENGTH() == 2);
    BOOST_TEST(list.SLGET(0) == "before_target");
    
    // Попытка добавить перед несуществующим элементом в непустом списке
    list.SLPUSH_BEFORE("non_existent2", "test2");
    BOOST_TEST(list.SLLENGTH() == 2);
    
    // 2. Тест всех ветвей в SLPUSH_AFTER
    // Пустой список (не обрабатывается отдельно, но проверим)
    SList list2;
    list2.SLPUSH_AFTER("non_existent", "test");
    BOOST_TEST(list2.SLLENGTH() == 0);
    
    // После существующего элемента
    list2.SLPUSH_HEAD("base");
    list2.SLPUSH_AFTER("base", "after_base");
    BOOST_TEST(list2.SLLENGTH() == 2);
    BOOST_TEST(list2.SLGET(1) == "after_base");
    
    // 3. Тест всех ветвей в SLREMOVE_BEFORE
    // Недостаточно элементов
    SList list3;
    list3.SLREMOVE_BEFORE("anything");
    BOOST_TEST(list3.SLLENGTH() == 0);
    
    // Только один элемент
    list3.SLPUSH_HEAD("single");
    list3.SLREMOVE_BEFORE("single");
    BOOST_TEST(list3.SLLENGTH() == 1);
    
    // Два элемента, удаление перед головой (невозможно)
    list3.SLPUSH_HEAD("new_head");
    list3.SLREMOVE_BEFORE("new_head");
    BOOST_TEST(list3.SLLENGTH() == 2); // Исправлено: осталось 2 элемента
    
    // Элемент не найден или перед ним нет элемента
    list3.SLPUSH_HEAD("A");
    list3.SLREMOVE_BEFORE("non_existent");
    BOOST_TEST(list3.SLLENGTH() == 3); // Исправлено: осталось 3 элемента
    
    // 4. Тест всех ветвей в SLREMOVE_AFTER
    // Недостаточно элементов
    SList list4;
    list4.SLREMOVE_AFTER("anything");
    BOOST_TEST(list4.SLLENGTH() == 0);
    
    // Элемент найден, но после него нет элемента
    list4.SLPUSH_HEAD("last");
    list4.SLREMOVE_AFTER("last");
    BOOST_TEST(list4.SLLENGTH() == 1);
    
    // 5. Тест всех ветвей в SLREMOVE_VALUE
    // Пустой список
    SList list5;
    list5.SLREMOVE_VALUE("anything");
    BOOST_TEST(list5.SLLENGTH() == 0);
    
    // Удаление единственного элемента (головы)
    list5.SLPUSH_HEAD("only");
    list5.SLREMOVE_VALUE("only");
    BOOST_TEST(list5.SLLENGTH() == 0);
    BOOST_TEST(list5.SLSEARCH("only") == false);
    
    // Удаление несуществующего элемента в непустом списке
    list5.SLPUSH_HEAD("exists");
    list5.SLREMOVE_VALUE("non_existent");
    BOOST_TEST(list5.SLLENGTH() == 1);
    
    // 6. Тест рекурсивного вывода (метод printReverse)
    SList list6;
    list6.SLPUSH_HEAD("C");
    list6.SLPUSH_HEAD("B");
    list6.SLPUSH_HEAD("A");
    
    // Просто вызываем, чтобы покрыть ветви в printReverse
    list6.SLPRINT_BACKWARD();
    
    // 7. Тест SLGET с разными индексами (граничные случаи)
    SList list7;
    for(int i = 0; i < 5; i++) {
        list7.SLPUSH_TAIL("item_" + std::to_string(i));
    }
    
    // Валидные индексы
    BOOST_TEST(list7.SLGET(0) == "item_0");
    BOOST_TEST(list7.SLGET(2) == "item_2");
    BOOST_TEST(list7.SLGET(4) == "item_4");
    
    // Невалидный индекс должен бросать исключение
    BOOST_CHECK_THROW(list7.SLGET(5), std::out_of_range);
    BOOST_CHECK_THROW(list7.SLGET(10), std::out_of_range);
}
// DList 

BOOST_AUTO_TEST_CASE(DList_FullOperations) {
    DList list;
    
    // Тест пустого списка
    BOOST_TEST(list.DLLENGTH() == 0);
    
    // Тест всех методов добавления
    list.DLPUSH_HEAD("head");
    BOOST_TEST(list.DLLENGTH() == 1);
    BOOST_TEST(list.DLGET(0) == "head");
    
    list.DLPUSH_TAIL("tail");
    BOOST_TEST(list.DLLENGTH() == 2);
    BOOST_TEST(list.DLGET(1) == "tail");
    
    list.DLPUSH_BEFORE("tail", "before_tail");
    BOOST_TEST(list.DLLENGTH() == 3);
    
    list.DLPUSH_AFTER("head", "after_head");
    BOOST_TEST(list.DLLENGTH() == 4);
    
    // Проверяем порядок
    BOOST_TEST(list.DLGET(0) == "head");
    BOOST_TEST(list.DLGET(1) == "after_head");
    
    // Тест поиска
    BOOST_TEST(list.DLSEARCH("after_head") == true);
    BOOST_TEST(list.DLSEARCH("non_existent") == false);
    
    // Тест всех методов удаления
    list.DLREMOVE_HEAD();
    BOOST_TEST(list.DLLENGTH() == 3);
    
    list.DLREMOVE_TAIL();
    BOOST_TEST(list.DLLENGTH() == 2);
    
    // Добавляем элементы для тестирования остальных методов удаления
    list.DLPUSH_HEAD("A");
    list.DLPUSH_TAIL("B");
    list.DLPUSH_TAIL("C");
    list.DLPUSH_TAIL("D");
    
    list.DLREMOVE_BEFORE("C");
    BOOST_TEST(list.DLSEARCH("B") == false);
    
    list.DLREMOVE_AFTER("C");
    BOOST_TEST(list.DLSEARCH("D") == false);
    
    list.DLREMOVE_VALUE("C");
    BOOST_TEST(list.DLSEARCH("C") == false);
    
    // Тест методов вывода
    list.DLPRINT_FORWARD();
    list.DLPRINT_BACKWARD();
    list.DLPRINT_HEAD();
    list.DLPRINT_TAIL();
    
    // Тест очистки
    list.DLCLEAR();
    BOOST_TEST(list.DLLENGTH() == 0);
}

BOOST_AUTO_TEST_CASE(DList_EdgeCase_BranchCoverage) {
    DList list;
    
    // 1. Тест DLPUSH_BEFORE
    // Пустой список
    list.DLPUSH_BEFORE("non_existent", "test");
    BOOST_TEST(list.DLLENGTH() == 0);
    
    // Добавляем в начало (target - голова)
    list.DLPUSH_HEAD("target");
    list.DLPUSH_BEFORE("target", "before_head");
    BOOST_TEST(list.DLLENGTH() == 2);
    BOOST_TEST(list.DLGET(0) == "before_head");
    
    // Элемент не найден
    list.DLPUSH_BEFORE("non_existent2", "test2");
    BOOST_TEST(list.DLLENGTH() == 2);
    
    // 2. Тест DLPUSH_AFTER
    // Пустой список
    DList list2;
    list2.DLPUSH_AFTER("non_existent", "test");
    BOOST_TEST(list2.DLLENGTH() == 0);
    
    // Добавляем после хвоста (target - tail)
    list2.DLPUSH_HEAD("base");
    list2.DLPUSH_AFTER("base", "after_tail");
    BOOST_TEST(list2.DLLENGTH() == 2);
    BOOST_TEST(list2.DLGET(1) == "after_tail");
    
    // 3. Тест DLREMOVE_BEFORE
    // Недостаточно элементов
    DList list3;
    list3.DLREMOVE_BEFORE("anything");
    BOOST_TEST(list3.DLLENGTH() == 0);
    
    // Только один элемент
    list3.DLPUSH_HEAD("single");
    list3.DLREMOVE_BEFORE("single");
    BOOST_TEST(list3.DLLENGTH() == 1);
    
    // Два элемента, удаление перед головой (невозможно)
    list3.DLPUSH_HEAD("new_head");
    list3.DLREMOVE_BEFORE("new_head");
    BOOST_TEST(list3.DLLENGTH() == 2); // Исправлено: осталось 2 элемента
    
    // Элемент не найден
    list3.DLPUSH_HEAD("A");
    list3.DLREMOVE_BEFORE("non_existent");
    BOOST_TEST(list3.DLLENGTH() == 3); // Исправлено: осталось 3 элемента
    
    // 4. Тест DLREMOVE_AFTER
    // Недостаточно элементов
    DList list4;
    list4.DLREMOVE_AFTER("anything");
    BOOST_TEST(list4.DLLENGTH() == 0);
    
    // Невозможно удалить после хвоста (один элемент)
    list4.DLPUSH_HEAD("tail");
    list4.DLREMOVE_AFTER("tail");
    BOOST_TEST(list4.DLLENGTH() == 1);
    
    // Невозможно удалить после хвоста (несколько элементов)
    list4.DLPUSH_HEAD("A");
    list4.DLPUSH_HEAD("B");
    list4.DLREMOVE_AFTER("tail"); // tail - последний элемент
    BOOST_TEST(list4.DLLENGTH() == 3);
    
    // 5. Тест DLREMOVE_VALUE
    // Пустой список
    DList list5;
    list5.DLREMOVE_VALUE("anything");
    BOOST_TEST(list5.DLLENGTH() == 0);
    
    // Удаление единственного элемента (головы и хвоста одновременно)
    list5.DLPUSH_HEAD("only");
    list5.DLREMOVE_VALUE("only");
    BOOST_TEST(list5.DLLENGTH() == 0);
    BOOST_TEST(list5.DLSEARCH("only") == false);
    
    // Удаление головы (не единственной)
    list5.DLPUSH_TAIL("first");
    list5.DLPUSH_TAIL("second");
    list5.DLPUSH_TAIL("third");
    list5.DLREMOVE_VALUE("first");
    BOOST_TEST(list5.DLLENGTH() == 2);
    BOOST_TEST(list5.DLGET(0) == "second");
    
    // Удаление хвоста (не единственного)
    list5.DLREMOVE_VALUE("third");
    BOOST_TEST(list5.DLLENGTH() == 1);
    BOOST_TEST(list5.DLGET(0) == "second");
    
    // Удаление несуществующего элемента
    list5.DLREMOVE_VALUE("non_existent");
    BOOST_TEST(list5.DLLENGTH() == 1);
    
    // 6. Тест методов вывода для разных состояний списка
    DList list6;
    
    // Пустой список - все методы вывода
    list6.DLPRINT_FORWARD();
    list6.DLPRINT_BACKWARD();
    list6.DLPRINT_HEAD();
    list6.DLPRINT_TAIL();
    
    // Список с одним элементом
    list6.DLPUSH_HEAD("single");
    list6.DLPRINT_FORWARD();
    list6.DLPRINT_BACKWARD();
    list6.DLPRINT_HEAD();
    list6.DLPRINT_TAIL();
    
    // Список с несколькими элементами
    list6.DLPUSH_TAIL("second");
    list6.DLPUSH_TAIL("third");
    list6.DLPRINT_FORWARD();
    list6.DLPRINT_BACKWARD();
    
    // 7. Тест DLGET с граничными индексами
    DList list7;
    for(int i = 0; i < 5; i++) {
        list7.DLPUSH_TAIL("ditem_" + std::to_string(i));
    }
    
    // Первый элемент
    BOOST_TEST(list7.DLGET(0) == "ditem_0");
    // Средний элемент
    BOOST_TEST(list7.DLGET(2) == "ditem_2");
    // Последний элемент
    BOOST_TEST(list7.DLGET(4) == "ditem_4");
    
    // Исключения
    BOOST_CHECK_THROW(list7.DLGET(5), std::out_of_range);
    BOOST_CHECK_THROW(list7.DLGET(100), std::out_of_range);
    
    // 8. Тест связей prev/next при различных операциях
    DList list8;
    list8.DLPUSH_HEAD("A");
    list8.DLPUSH_TAIL("B");
    list8.DLPUSH_TAIL("C");
    
    // Проверяем связи
    BOOST_TEST(list8.DLGET(0) == "A");
    BOOST_TEST(list8.DLGET(1) == "B");
    BOOST_TEST(list8.DLGET(2) == "C");
    
    // Удаляем средний элемент
    list8.DLREMOVE_VALUE("B");
    BOOST_TEST(list8.DLLENGTH() == 2);
    BOOST_TEST(list8.DLGET(0) == "A");
    BOOST_TEST(list8.DLGET(1) == "C");
    
    // Добавляем между оставшимися
    list8.DLPUSH_BEFORE("C", "B_new");
    BOOST_TEST(list8.DLLENGTH() == 3);
    BOOST_TEST(list8.DLGET(1) == "B_new");
    
    // Проверяем что хвост корректно обновляется
    list8.DLPUSH_AFTER("C", "D");
    BOOST_TEST(list8.DLLENGTH() == 4);
    BOOST_TEST(list8.DLGET(3) == "D");
}

// Queue 

BOOST_AUTO_TEST_CASE(Queue_FullOperations) {
    Queue queue;
    
    // Тест пустой очереди
    BOOST_TEST(queue.isEmpty() == true);
    
    // Тест добавления множества элементов
    for(int i = 0; i < 100; i++) {
        queue.QPUSH("element_" + std::to_string(i));
        BOOST_TEST(queue.isEmpty() == false);
    }
    
    // Тест QFRONT без удаления
    BOOST_TEST(queue.QFRONT() == "element_0");
    
    // Тест извлечения всех элементов
    for(int i = 0; i < 100; i++) {
        BOOST_TEST(queue.QPOP() == "element_" + std::to_string(i));
    }
    
    BOOST_TEST(queue.isEmpty() == true);
    
    // Тест исключений при пустой очереди
    BOOST_CHECK_THROW(queue.QPOP(), std::runtime_error);
    BOOST_CHECK_THROW(queue.QFRONT(), std::runtime_error);
    
    // Тест QPRINT
    queue.QPUSH("test_print");
    queue.QPRINT();
    
    // Тест очистки
    queue.QCLEAR();
    BOOST_TEST(queue.isEmpty() == true);
    
    // Тест повторного использования после очистки
    queue.QPUSH("after_clear");
    BOOST_TEST(queue.QFRONT() == "after_clear");
}

// Tree 

BOOST_AUTO_TEST_CASE(RBTree_FullOperations) {
    RBTree tree;
    
    // Тест пустого дерева
    BOOST_TEST(tree.TSEARCH("anything") == false);
    
    // Тест вставки множества элементов
    std::vector<std::string> elements = {
        "apple", "banana", "cherry", "date", "fig", "grape", "kiwi", "lemon", "mango"
    };
    
    for(const auto& elem : elements) {
        tree.TINSERT(elem);
        BOOST_TEST(tree.TSEARCH(elem) == true);
    }
    
    // Тест поиска несуществующих элементов
    BOOST_TEST(tree.TSEARCH("watermelon") == false);
    BOOST_TEST(tree.TSEARCH("") == false);
    
    // Тест метода TGET
    for(const auto& elem : elements) {
        BOOST_TEST(tree.TGET(elem) == elem);
    }
    BOOST_TEST(tree.TGET("non_existent").empty());
    
    // Тест удаления элементов
    tree.TDELETE("banana");
    BOOST_TEST(tree.TSEARCH("banana") == false);
    
    tree.TDELETE("fig");
    BOOST_TEST(tree.TSEARCH("fig") == false);
    
    // Удаляем остальные элементы
    for(const auto& elem : elements) {
        if(elem != "banana" && elem != "fig") {
            tree.TDELETE(elem);
            BOOST_TEST(tree.TSEARCH(elem) == false);
        }
    }
    
    // Тест методов вывода
    tree.TINSERT("test");
    tree.TPRINT_INORDER();
    tree.TPRINT_PREORDER();
    tree.TPRINT_POSTORDER();
    tree.TPRINT_TREE();
    
    // Тест clear и setRoot
    tree.clear();
    BOOST_TEST(tree.TSEARCH("test") == false);
    
    // Тест вспомогательных методов
    BOOST_CHECK_NO_THROW(tree.getRoot());
    BOOST_CHECK_NO_THROW(tree.getNil());
}

BOOST_AUTO_TEST_CASE(RBTree_AdvancedOperations) {
    RBTree tree;
    
    // Вставляем элементы в случайном порядке
    std::vector<std::string> elements;
    for(int i = 0; i < 100; i++) {
        elements.push_back("node_" + std::to_string(i));
    }
    
    std::random_device rd;
    std::mt19937 g(rd());
    std::shuffle(elements.begin(), elements.end(), g);
    
    for(const auto& elem : elements) {
        tree.TINSERT(elem);
    }
    
    // Проверяем что все элементы существуют
    for(const auto& elem : elements) {
        BOOST_TEST(tree.TSEARCH(elem) == true);
    }
    
    // Удаляем половину элементов
    for(int i = 0; i < 50; i++) {
        tree.TDELETE(elements[i]);
        BOOST_TEST(tree.TSEARCH(elements[i]) == false);
    }
    
    // Проверяем что оставшиеся элементы все еще существуют
    for(int i = 50; i < 100; i++) {
        BOOST_TEST(tree.TSEARCH(elements[i]) == true);
    }
    
    // Тест метода getAllNodes
    auto nodes = tree.getAllNodes();
    BOOST_TEST(nodes.size() == 50);
    
    // Тест очистки
    tree.clear();
    BOOST_TEST(tree.getAllNodes().size() == 0);
}

// Hash Table 

BOOST_AUTO_TEST_CASE(ChainingHashTable_FullOperations) {
    // Тест с разной емкостью
    ChainingHashTable ht1(5);
    ChainingHashTable ht2(50);
    ChainingHashTable ht3(100);
    
    // Тест добавления
    ht1.add(std::make_pair(1, 100));
    ht1.add(std::make_pair(2, 200));
    ht1.add(std::make_pair(3, 300));
    
    // Проверяем добавление
    auto result = ht1.contains(1);
    BOOST_TEST(result.first == true);
    BOOST_TEST(result.second == 100);
    
    result = ht1.contains(2);
    BOOST_TEST(result.first == true);
    BOOST_TEST(result.second == 200);
    
    result = ht1.contains(3);
    BOOST_TEST(result.first == true);
    BOOST_TEST(result.second == 300);
    
    // Тест дублирования ключей
    ht1.add(std::make_pair(1, 999)); // Не должно изменить значение
    result = ht1.contains(1);
    BOOST_TEST(result.second == 100); // Должно остаться прежним
    
    // Тест удаления
    ht1.remove(2);
    result = ht1.contains(2);
    BOOST_TEST(result.first == false);
    
    // Удаление несуществующего ключа
    BOOST_CHECK_NO_THROW(ht1.remove(999));
    
    // Тест getAllElements
    auto elements = ht1.getAllElements();
    BOOST_TEST(elements.size() == 2); // Осталось 2 элемента
    
    // Тест toString
    std::string str = ht1.toString();
    BOOST_TEST(!str.empty());
    
    // Тест статистики
    int minLen, maxLen;
    double avgLen;
    ht1.getChainLengths(minLen, maxLen, avgLen);
    
    // Тест геттеров
    BOOST_TEST(ht1.getSize() == 2);
    BOOST_TEST(ht1.getCapacity() == 5);
    BOOST_TEST(ht1.getLoadFactor() == 2.0 / 5.0);
    
    // Тест очистки
    ht1.clear();
    BOOST_TEST(ht1.isEmpty() == true);
    BOOST_TEST(ht1.getSize() == 0);
    
    // Тест повторного использования после очистки
    ht1.add(std::make_pair(10, 1000));
    BOOST_TEST(ht1.contains(10).first == true);
}

BOOST_AUTO_TEST_CASE(ChainingHashTable_PerformanceTest) {
    ChainingHashTable ht(100);
    
    // Добавляем много элементов
    for(int i = 0; i < 1000; i++) {
        ht.add(std::make_pair(i, i * 10));
    }
    
    BOOST_TEST(ht.getSize() == 1000);
    
    // Проверяем все элементы
    for(int i = 0; i < 1000; i++) {
        auto result = ht.contains(i);
        BOOST_TEST(result.first == true);
        BOOST_TEST(result.second == i * 10);
    }
    
    // Удаляем половину
    for(int i = 0; i < 500; i++) {
        ht.remove(i);
    }
    
    BOOST_TEST(ht.getSize() == 500);
    
    // Проверяем статистику
    int minLen, maxLen;
    double avgLen;
    ht.getChainLengths(minLen, maxLen, avgLen);
    
    // Тест isEmpty
    BOOST_TEST(ht.isEmpty() == false);
    
    // Полная очистка
    ht.clear();
    BOOST_TEST(ht.isEmpty() == true);
    BOOST_TEST(ht.getSize() == 0);
}

// Serialization 

BOOST_AUTO_TEST_CASE(Serialization_AllFormats) {
    // Тест MArray сериализации
    {
        MArray arr;
        arr.MADDEND("serial_test_1");
        arr.MADDEND("serial_test_2");
        arr.MADDEND("serial_test_3");
        
        // Текстовая сериализация
        saveToText(arr, "test_serial_array.txt");
        
        MArray arr_loaded;
        loadFromText(arr_loaded, "test_serial_array.txt");
        
        BOOST_TEST(arr_loaded.MLENGTH() == 3);
        BOOST_TEST(arr_loaded.MGETINDEX(0) == "serial_test_1");
        BOOST_TEST(arr_loaded.MGETINDEX(2) == "serial_test_3");
        
        // Бинарная сериализация
        saveToBinary(arr, "test_serial_array.bin");
        
        MArray arr_bin_loaded;
        loadFromBinary(arr_bin_loaded, "test_serial_array.bin");
        
        BOOST_TEST(arr_bin_loaded.MLENGTH() == 3);
        BOOST_TEST(arr_bin_loaded.MGETINDEX(1) == "serial_test_2");
        
        // Очистка
        std::remove("test_serial_array.txt");
        std::remove("test_serial_array.bin");
    }
    
    // Тест SList сериализации
    {
        SList list;
        list.SLPUSH_TAIL("list_item_1");
        list.SLPUSH_TAIL("list_item_2");
        list.SLPUSH_TAIL("list_item_3");
        
        saveToText(list, "test_serial_list.txt");
        
        SList list_loaded;
        loadFromText(list_loaded, "test_serial_list.txt");
        
        BOOST_TEST(list_loaded.SLLENGTH() == 3);
        BOOST_TEST(list_loaded.SLGET(0) == "list_item_1");
        
        std::remove("test_serial_list.txt");
    }
    
    // Тест Stack сериализации
    {
        Stack stack;
        stack.SPUSH("stack_bottom");
        stack.SPUSH("stack_middle");
        stack.SPUSH("stack_top");
        
        saveToText(stack, "test_serial_stack.txt");
        
        Stack stack_loaded;
        loadFromText(stack_loaded, "test_serial_stack.txt");
        
        BOOST_TEST(stack_loaded.STOP() == "stack_top");
        stack_loaded.SPOP();
        BOOST_TEST(stack_loaded.STOP() == "stack_middle");
        
        std::remove("test_serial_stack.txt");
    }
    
    // Тест Queue сериализации
    {
        Queue queue;
        queue.QPUSH("queue_first");
        queue.QPUSH("queue_second");
        queue.QPUSH("queue_third");
        
        saveToText(queue, "test_serial_queue.txt");
        
        Queue queue_loaded;
        loadFromText(queue_loaded, "test_serial_queue.txt");
        
        BOOST_TEST(queue_loaded.QFRONT() == "queue_first");
        queue_loaded.QPOP();
        BOOST_TEST(queue_loaded.QFRONT() == "queue_second");
        
        std::remove("test_serial_queue.txt");
    }
    
    // Тест HashTable сериализации
    {
        ChainingHashTable ht(10);
        ht.add(std::make_pair(1, 100));
        ht.add(std::make_pair(2, 200));
        ht.add(std::make_pair(3, 300));
        
        saveToText(ht, "test_serial_ht.txt");
        
        ChainingHashTable ht_loaded(10);
        loadFromText(ht_loaded, "test_serial_ht.txt");
        
        BOOST_TEST(ht_loaded.contains(1).first == true);
        BOOST_TEST(ht_loaded.contains(1).second == 100);
        BOOST_TEST(ht_loaded.contains(2).first == true);
        BOOST_TEST(ht_loaded.contains(3).first == true);
        
        // Бинарная сериализация
        saveToBinary(ht, "test_serial_ht.bin");
        
        ChainingHashTable ht_bin_loaded(10);
        loadFromBinary(ht_bin_loaded, "test_serial_ht.bin");
        
        BOOST_TEST(ht_bin_loaded.contains(1).first == true);
        
        std::remove("test_serial_ht.txt");
        std::remove("test_serial_ht.bin");
    }
    
    // Тест RBTree сериализации
    {
        RBTree tree;
        tree.TINSERT("tree_apple");
        tree.TINSERT("tree_banana");
        tree.TINSERT("tree_cherry");
        
        // Текстовая сериализация
        saveToText(tree, "test_serial_tree.txt");
        
        RBTree tree_loaded;
        loadFromText(tree_loaded, "test_serial_tree.txt");
        
        BOOST_TEST(tree_loaded.TSEARCH("tree_apple") == true);
        BOOST_TEST(tree_loaded.TSEARCH("tree_banana") == true);
        BOOST_TEST(tree_loaded.TSEARCH("tree_cherry") == true);
        
        
        RBTree tree_bin_loaded;
        
        BOOST_TEST(tree_bin_loaded.TSEARCH("tree_apple") == true);
        
        std::remove("test_serial_tree.txt");
        std::remove("test_serial_tree.bin");
    }
}

BOOST_AUTO_TEST_CASE(Serialization_ErrorHandling) {
    MArray arr;
    
    // Тест исключений при открытии файлов
    BOOST_CHECK_THROW(saveToText(arr, "/nonexistent/path/file.txt"), std::runtime_error);
    BOOST_CHECK_THROW(loadFromText(arr, "/nonexistent/path/file.txt"), std::runtime_error);
    
    BOOST_CHECK_THROW(saveToBinary(arr, "/nonexistent/path/file.bin"), std::runtime_error);
    BOOST_CHECK_THROW(loadFromBinary(arr, "/nonexistent/path/file.bin"), std::runtime_error);
    
    // Создаем реальный файл для тестирования бинарного чтения с ошибками
    {
        std::ofstream bad_file("bad_binary.bin", std::ios::binary);
        bad_file << "not binary data";
        bad_file.close();
        
        // Проверяем что функция не падает
        try {
            loadFromBinary(arr, "bad_binary.bin");
            // Если дошло сюда - значит исключения не было
            arr.MCLEAR();
        } catch (const std::exception& e) {
            // Исключение - это нормально для поврежденного бинарного файла
        } catch (...) {
        }
        
        std::remove("bad_binary.bin");
    }
}

BOOST_AUTO_TEST_CASE(Integration_AllStructures) {
    // Создаем и тестируем все структуры вместе
    MArray arr;
    SList slist;
    DList dlist;
    Stack stack;
    Queue queue;
    RBTree tree;
    ChainingHashTable ht(20);
    
    // Заполняем все структуры
    for(int i = 0; i < 10; i++) {
        std::string val = "value_" + std::to_string(i);
        int num = i * 100;
        
        arr.MADDEND(val);
        slist.SLPUSH_TAIL(val);
        dlist.DLPUSH_TAIL(val);
        stack.SPUSH(val);
        queue.QPUSH(val);
        tree.TINSERT(val);
        ht.add(std::make_pair(i, num));
    }
    
    // Проверяем все структуры
    BOOST_TEST(arr.MLENGTH() == 10);
    BOOST_TEST(slist.SLLENGTH() == 10);
    BOOST_TEST(dlist.DLLENGTH() == 10);
    BOOST_TEST(tree.TSEARCH("value_5") == true);
    BOOST_TEST(ht.contains(5).first == true);
    BOOST_TEST(ht.contains(5).second == 500);
    
    // Последовательные операции
    for(int i = 0; i < 5; i++) {
        arr.MREMOVEINDEX(0);
        slist.SLREMOVE_HEAD();
        dlist.DLREMOVE_HEAD();
        stack.SPOP();
        queue.QPOP();
        tree.TDELETE("value_" + std::to_string(i));
        ht.remove(i);
    }
    
    BOOST_TEST(arr.MLENGTH() == 5);
    BOOST_TEST(slist.SLLENGTH() == 5);
    BOOST_TEST(dlist.DLLENGTH() == 5);
    BOOST_TEST(tree.TSEARCH("value_0") == false);
    BOOST_TEST(ht.contains(0).first == false);
    
    // Сериализация всех структур
    saveToText(arr, "integration_arr.txt");
    saveToText(slist, "integration_slist.txt");
    saveToText(dlist, "integration_dlist.txt");
    saveToText(stack, "integration_stack.txt");
    saveToText(queue, "integration_queue.txt");
    saveToText(tree, "integration_tree.txt");
    saveToText(ht, "integration_ht.txt");
    
    // Очистка тестовых файлов
    std::remove("integration_arr.txt");
    std::remove("integration_slist.txt");
    std::remove("integration_dlist.txt");
    std::remove("integration_stack.txt");
    std::remove("integration_queue.txt");
    std::remove("integration_tree.txt");
    std::remove("integration_ht.txt");
}

BOOST_AUTO_TEST_SUITE_END()
BOOST_AUTO_TEST_SUITE(ExtraCoverageTests)

BOOST_AUTO_TEST_CASE(PrintMethods_DoNotCrash) {
    // Тестируем что методы вывода не падают
    MArray arr;
    arr.MPRINT(); // Пустой массив
    
    arr.MADDEND("test");
    arr.MPRINT(); // Непустой массив
    
    SList slist;
    slist.SLPRINT_FORWARD();
    slist.SLPRINT_BACKWARD();
    slist.SLPRINT_HEAD();
    slist.SLPRINT_TAIL();
    
    slist.SLPUSH_TAIL("test");
    slist.SLPRINT_FORWARD();
    slist.SLPRINT_BACKWARD();
    slist.SLPRINT_HEAD();
    slist.SLPRINT_TAIL();
    
    DList dlist;
    dlist.DLPRINT_FORWARD();
    dlist.DLPRINT_BACKWARD();
    dlist.DLPRINT_HEAD();
    dlist.DLPRINT_TAIL();
    
    Stack stack;
    stack.SPRINT();
    
    Queue queue;
    queue.QPRINT();
    
    RBTree tree;
    tree.TPRINT_INORDER();
    tree.TPRINT_PREORDER();
    tree.TPRINT_POSTORDER();
    tree.TPRINT_TREE();
    
    tree.TINSERT("test");
    tree.TPRINT_INORDER();
    tree.TPRINT_PREORDER();
    tree.TPRINT_POSTORDER();
    tree.TPRINT_TREE();
}

BOOST_AUTO_TEST_CASE(Destructors_DoNotCrash) {
    // Тест что деструкторы работают корректно
    {
        MArray arr;
        for(int i = 0; i < 100; i++) {
            arr.MADDEND("test_" + std::to_string(i));
        }
    } // Деструктор вызывается здесь
    
    {
        SList list;
        for(int i = 0; i < 100; i++) {
            list.SLPUSH_TAIL("test_" + std::to_string(i));
        }
    }
    
    {
        DList list;
        for(int i = 0; i < 100; i++) {
            list.DLPUSH_TAIL("test_" + std::to_string(i));
        }
    }
    
    {
        Stack stack;
        for(int i = 0; i < 100; i++) {
            stack.SPUSH("test_" + std::to_string(i));
        }
    }
    
    {
        Queue queue;
        for(int i = 0; i < 100; i++) {
            queue.QPUSH("test_" + std::to_string(i));
        }
    }
    
    {
        RBTree tree;
        for(int i = 0; i < 100; i++) {
            tree.TINSERT("test_" + std::to_string(i));
        }
    }
    
    {
        ChainingHashTable ht(10);
        for(int i = 0; i < 100; i++) {
            ht.add(std::make_pair(i, i * 10));
        }
    }
}

BOOST_AUTO_TEST_CASE(EdgeCase_Operations) {
    // MArray - добавление в начало пустого массива
    {
        MArray arr;
        arr.MADDINDEX(0, "first");
        BOOST_TEST(arr.MLENGTH() == 1);
        BOOST_TEST(arr.MGETINDEX(0) == "first");
    }
    
    // SList - добавление перед/после в пустом списке
    {
        SList list;
        list.SLPUSH_BEFORE("nonexistent", "test"); // Должно вывести сообщение
        list.SLPUSH_AFTER("nonexistent", "test"); // Должно вывести сообщение
        BOOST_TEST(list.SLLENGTH() == 0);
    }
    
    // DList - аналогично
    {
        DList list;
        list.DLPUSH_BEFORE("nonexistent", "test");
        list.DLPUSH_AFTER("nonexistent", "test");
        BOOST_TEST(list.DLLENGTH() == 0);
    }
    
    // Stack - многократный push/pop
    {
        Stack stack;
        for(int i = 0; i < 1000; i++) {
            stack.SPUSH("elem_" + std::to_string(i));
        }
        BOOST_TEST(stack.isEmpty() == false);
        
        for(int i = 0; i < 1000; i++) {
            stack.SPOP();
        }
        BOOST_TEST(stack.isEmpty() == true);
    }
    
    // Queue - аналогично
    {
        Queue queue;
        for(int i = 0; i < 1000; i++) {
            queue.QPUSH("elem_" + std::to_string(i));
        }
        
        for(int i = 0; i < 1000; i++) {
            queue.QPOP();
        }
        BOOST_TEST(queue.isEmpty() == true);
    }
    
    // RBTree - вставка дубликатов
    {
        RBTree tree;
        tree.TINSERT("duplicate");
        tree.TINSERT("duplicate"); // Должно вывести сообщение
        tree.TINSERT("duplicate"); // Должно вывести сообщение
        BOOST_TEST(tree.TSEARCH("duplicate") == true);
    }
    
    // HashTable - коллизии
    {
        ChainingHashTable ht(1); // Очень маленькая емкость для принудительных коллизий
        for(int i = 0; i < 10; i++) {
            ht.add(std::make_pair(i, i * 100));
        }
        BOOST_TEST(ht.getSize() == 10);
        BOOST_TEST(ht.getLoadFactor() == 10.0);
    }
}
BOOST_AUTO_TEST_CASE(SList_Complex_Scenario) {
    SList list;
    
    // Сценарий: добавление, удаление, повторное добавление
    list.SLPUSH_HEAD("start");
    for(int i = 1; i < 10; i++) {
        list.SLPUSH_TAIL("node_" + std::to_string(i));
    }
    BOOST_TEST(list.SLLENGTH() == 10);
    // Удаляем элементы с четными номерами в имени
    list.SLREMOVE_VALUE("node_2");
    list.SLREMOVE_VALUE("node_4");
    list.SLREMOVE_VALUE("node_6");
    list.SLREMOVE_VALUE("node_8");
    
    BOOST_TEST(list.SLLENGTH() == 6); // start + node_1,3,5,7,9
    
    // Добавляем перед и после оставшихся
    list.SLPUSH_BEFORE("start", "new_start");
    // Ищем элемент, который точно существует
    if(list.SLSEARCH("node_8")) {
        list.SLPUSH_AFTER("node_8", "after_8");
    } else {
        // node_8 был удален, добавляем после node_7
        list.SLPUSH_AFTER("node_7", "after_7");
    }
    
    // Проверяем поиск
    BOOST_TEST(list.SLSEARCH("new_start") == true);
    BOOST_TEST(list.SLSEARCH("node_1") == true);
    BOOST_TEST(list.SLSEARCH("node_2") == false); // Был удален
    
    // Очищаем и проверяем
    list.SLCLEAR();
    BOOST_TEST(list.SLLENGTH() == 0);
    BOOST_TEST(list.SLSEARCH("anything") == false);
}

BOOST_AUTO_TEST_CASE(DList_Complex_Scenario) {
    DList list;
    
    // Построение сложной структуры
    list.DLPUSH_HEAD("center");
    
    // Добавляем симметрично в обе стороны
    list.DLPUSH_BEFORE("center", "left_1");
    list.DLPUSH_AFTER("center", "right_1");
    list.DLPUSH_BEFORE("left_1", "left_2");
    list.DLPUSH_AFTER("right_1", "right_2");
    
    BOOST_TEST(list.DLLENGTH() == 5);
    
    // Проверяем порядок
    BOOST_TEST(list.DLGET(0) == "left_2");
    BOOST_TEST(list.DLGET(1) == "left_1");
    BOOST_TEST(list.DLGET(2) == "center");
    BOOST_TEST(list.DLGET(3) == "right_1");
    BOOST_TEST(list.DLGET(4) == "right_2");
    
    // Удаляем симметрично
    list.DLREMOVE_BEFORE("center");
    list.DLREMOVE_AFTER("center");
    
    BOOST_TEST(list.DLLENGTH() == 3);
    BOOST_TEST(list.DLGET(0) == "left_2");
    BOOST_TEST(list.DLGET(1) == "center");
    BOOST_TEST(list.DLGET(2) == "right_2");
    
    // Проверяем вывод в обоих направлениях
    list.DLPRINT_FORWARD();
    list.DLPRINT_BACKWARD();
    
    // Имитация использования как стека/очереди
    for(int i = 0; i < 5; i++) {
        if(i % 2 == 0) {
            list.DLPUSH_HEAD("stack_" + std::to_string(i));
        } else {
            list.DLPUSH_TAIL("queue_" + std::to_string(i));
        }
    }
    
    BOOST_TEST(list.DLLENGTH() == 8); // 3 существующих + 5 новых
    
    // Удаляем с обоих концов
    list.DLREMOVE_HEAD();
    list.DLREMOVE_TAIL();
    
    BOOST_TEST(list.DLLENGTH() == 6);
}

BOOST_AUTO_TEST_SUITE_END()