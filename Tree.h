#pragma once
#include <cstdint>
#include <string>

using namespace std;

enum Color { RED, BLACK };

struct TNode {
    string data;    // Значение узла
    Color color;    // Цвет узла
    TNode* left;    // Левый потомок
    TNode* right;   // Правый потомок
    TNode* parent;  // Родительский узел
};

class RBTree {
private:
    TNode* root;    // Корень дерева
    TNode* nil;     // Специальный NIL-узел (черный лист)
    
    // Вспомогательные методы
    
    void clearTree(TNode* node);
    void leftRotate(TNode* x);
    void rightRotate(TNode* y);
    void insertFixup(TNode* z);
    TNode* minimum(TNode* node);
    void transplant(TNode* u, TNode* v);
    void deleteFixup(TNode* x);
    
    // Рекурсивные методы обхода
    void inorder(TNode* node) const;
    void preorder(TNode* node) const;
    void postorder(TNode* node) const;
    void printTree(TNode* node, const string& prefix, bool isLeft) const;

public:
    RBTree();
    ~RBTree();
    TNode* createNode(const string& value);
    TNode* getRoot() const;
    TNode* getNil() const;

    // Основные операции
    void TINSERT(const string& value);
    void TDELETE(const string& value);
    bool TSEARCH(const string& value) const;
    string TGET(const string& value) const;
    
    // Вывод
    void TPRINT_INORDER() const;
    void TPRINT_PREORDER() const;
    void TPRINT_POSTORDER() const;
    void TPRINT_TREE() const;

    void clear();
    std::vector<std::string> getAllElements() const ;
    std::vector<std::pair<std::string, Color>> getAllNodes() const ;
    void setRoot(TNode* newRoot);
    void insertWithColor(const std::string& value, Color color);

};