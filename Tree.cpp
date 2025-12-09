#include <iostream>
#include <fstream>
#include <vector>
#include <string>
#include <queue>
#include <limits>
#include "Tree.h"

using namespace std;

// Безопасный ввод строки с обработкой ошибок
string safeGetline() {
    string input;
    while (true) {
        getline(cin, input);
        if (cin.fail()) {
            cin.clear(); // Сбрасываем флаги ошибок
            cin.ignore(numeric_limits<streamsize>::max(), '\n'); // очищаем буфер
            cout << "Ошибка ввода. Попробуйте снова: ";
        }
        else {
            break;
        }
    }
    return input;
}

//Безопасный ввод числа с обработкой ошибок
int safeGetInt() {
    int value;
    while (true) {
        cin >> value;
        if (cin.fail()) {
            cin.clear();
            cin.ignore(numeric_limits<streamsize>::max(), '\n');
            cout << "Ошибка: введите целое число: ";
        }
        else {
            cin.ignore(numeric_limits<streamsize>::max(), '\n'); // очищаем буфер
            break;
        }
    }
    return value;
}

TNode* RBTree::createNode(const string& value) {
    TNode* newNode = new TNode;
    newNode->data = value;
    newNode->color = RED;  // новые узлы всегда красные
    newNode->left = nil;
    newNode->right = nil;
    newNode->parent = nil;
    return newNode;
}

RBTree::RBTree() {
    // создаём NIL-узел (чёрный лист)
    nil = new TNode;
    nil->color = BLACK;
    nil->left = nullptr;
    nil->right = nullptr;
    nil->parent = nullptr;

    // Корень изначально указывает на NIL
    root = nil;
}

void RBTree::clearTree(TNode* node) {
    if (node != nil) {
        clearTree(node->left);
        clearTree(node->right);
        delete node;
    }
}

RBTree::~RBTree() {
    // Рекурсично удаляем все узлы, кроме NIL
    clearTree(root);

    // Удаляем NIL-узлы
    delete nil;
    nil = nullptr;
    root = nullptr;
}

void RBTree::leftRotate(TNode* x) {
    TNode* y = x->right; //y - правый потомок x
    x->right = y->left; 

    if (y->left != nil) {
        y->left->parent = x; //x - родитель левого поддерева y
    }

    y->parent = x->parent;

    if (x->parent == nil) { //если х был корнем
        root = y;
    }
    else if (x == x->parent->left) { //если х был левым потомком родителя
        x->parent->left = y;
    }
    else {
        x->parent->right = y;
    }

    y->left = x;
    x->parent = y;
}

void RBTree::rightRotate(TNode* y) {
    TNode* x = y->left;
    y->left = x->right;

    if (x->right != nil) {
        x->right->parent = y;
    }

    x->parent = y->parent;

    if (y->parent == nil) {
        root = x;
    }
    else if (y == y->parent->right) {
        y->parent->right = x;
    }
    else {
        y->parent->left = x;
    }

    x->right = y;
    y->parent = x;
}

void RBTree::insertFixup(TNode* z) {
    // Пока родитель красный (нарушение свойства)
    while (z != root && z->parent->color == RED) {
        // Проверяем, что дедушка существует
        if (z->parent->parent == nil) {
            break; // Не должно происходить, но на всякий случай
        }
        
        if (z->parent == z->parent->parent->left) {
            // Родитель - левый потомок дедушки
            TNode* y = z->parent->parent->right; // дядя z
            
            // Случай 1: дядя красный
            if (y != nil && y->color == RED) {
                z->parent->color = BLACK;
                y->color = BLACK;
                z->parent->parent->color = RED;
                z = z->parent->parent;
            } else {
                // Случай 2: дядя чёрный, z - правый потомок
                if (z == z->parent->right) {
                    z = z->parent;
                    leftRotate(z);
                }
                // Случай 3: дядя чёрный, z - левый потомок
                z->parent->color = BLACK;
                z->parent->parent->color = RED;
                rightRotate(z->parent->parent);
            }
        } else {
            // Симметричный случай: родитель - правый потомок дедушки
            TNode* y = z->parent->parent->left; // дядя z (левый потомок деда)
            
            // Случай 1: дядя красный
            if (y != nil && y->color == RED) {
                z->parent->color = BLACK;
                y->color = BLACK;
                z->parent->parent->color = RED;
                z = z->parent->parent;
            } else {
                // Случай 2: дядя чёрный, z - левый потомок
                if (z == z->parent->left) {
                    z = z->parent;
                    rightRotate(z);
                }
                // Случай 3: дядя чёрный, z - правый потомок
                z->parent->color = BLACK;
                z->parent->parent->color = RED;
                leftRotate(z->parent->parent);
            }
        }
    }
    
    // Корень всегда должен быть чёрным
    root->color = BLACK;
}

void RBTree::TINSERT(const string& value) {
    cout<<"start ins"<<endl;
    TNode* z = createNode(value);
    TNode* y = nil;
    TNode* x = root;

    // Поиск места для вставки пока не дойдём до nil
    while (x != nil) {
        y = x; //сохраняем текущий узел как родителя
        if (z->data < x->data) {
            x = x->left;
        }
        else if (z->data > x->data) {
            x = x->right;
        }
        else {
            // Элемент уже существует
            delete z;
            cout << "Элемент \"" << value << "\" уже существует в дереве" << endl;
            return;
        }
    }
    cout<<"---"<<endl;
    // Вставляем новый узел
    z->parent = y;
    if (y == nil) {
        root = z; //дерево было пусто
    }
    else if (z->data < y->data) {
        y->left = z;
    }
    else {
        y->right = z;
    }
    //устанавливаем потомков нового узла в nil и цвет в красный

    z->left = nil;
    z->right = nil;
    z->color = RED;
    cout<<"+"<<endl;
    // Исправляем свойства красно-чёрного дерева
    insertFixup(z);

    cout << "Элемент \"" << value << "\" добавлен в дерево" << endl;
}

TNode* RBTree::minimum(TNode* node) {
    while (node->left != nil) { // идём до самого левого узла
        node = node->left;
    }
    return node;
}

void RBTree::transplant(TNode* u, TNode* v) {
    if (u->parent == nil) {
        root = v;//u - корень
    }
    else if (u == u->parent->left) {
        u->parent->left = v; //u- левый потомок
    }
    else {
        u->parent->right = v; //u-правый потомок
    }
    v->parent = u->parent; //v - наследует родителя u
}

void RBTree::deleteFixup(TNode* x) {
    while (x != root && x->color == BLACK) {
        if (x == x->parent->left) {
            TNode* w = x->parent->right; // w - брат х

            if (w->color == RED) { //Случай 1: w красный
                w->color = BLACK;
                x->parent->color = RED;
                leftRotate(x->parent);
                w = x->parent->right;
            }
            //случай 2: оба ребёнка чёрные
            if (w->left->color == BLACK && w->right->color == BLACK) {
                w->color = RED;
                x = x->parent;
            }
            else {
            //случай 3: левый ребёнок w красный, правый чёрный
                if (w->right->color == BLACK) {
                    w->left->color = BLACK;
                    w->color = RED;
                    rightRotate(w);
                    w = x->parent->right;
                }
                //случай 4: правый ребёнок w красный
                w->color = x->parent->color;
                x->parent->color = BLACK;
                w->right->color = BLACK;
                leftRotate(x->parent);
                x = root;
            }
        }
        else {
            // Симметричный случай: x - правый потомок 
            TNode* w = x->parent->left;

            if (w->color == RED) {
                w->color = BLACK;
                x->parent->color = RED;
                rightRotate(x->parent);
                w = x->parent->left;
            }

            if (w->right->color == BLACK && w->left->color == BLACK) {
                w->color = RED;
                x = x->parent;
            }
            else {
                if (w->left->color == BLACK) {
                    w->right->color = BLACK;
                    w->color = RED;
                    leftRotate(w);
                    w = x->parent->left;
                }

                w->color = x->parent->color;
                x->parent->color = BLACK;
                w->left->color = BLACK;
                rightRotate(x->parent);
                x = root;
            }
        }
    }
    x->color = BLACK; //чтобы сохранить свойство 1
}

void RBTree::TDELETE(const string& value) {
    TNode* z = root;
    //поиск удаляемого элемента
    while (z != nil) {
        if (value == z->data) {
            break;
        }
        else if (value < z->data) {
            z = z->left;
        }
        else {
            z = z->right;
        }
    }

    if (z == nil) {
        //cout << "Элемент \"" << value << "\" не найден в дереве" << endl;
        return;
    }
    //у – узел который фактически удаляем
    TNode* y = z;
    TNode* x; //займёт место у или nil если детей нет
    Color y_original_color = y->color; //сохраняем исходный цвет у

    if (z->left == nil) { //случай 1: z не имеет левого потомка
        x = z->right;
        transplant(z, z->right);
    }
    else if (z->right == nil) { //случай 2: z не имеет правого потомка
        x = z->left;
        transplant(z, z->left);
    }
    else { //у z есть оба потомка
        y = minimum(z->right); //ищем минимальный узел в правом поддереве z(преемник)
        y_original_color = y->color; //сохраняем цвет преемника
        x = y->right; //x - правый потомок у

        if (y->parent == z) {//у правый потомок z
            x->parent = y;
        }
        else {
            transplant(y, y->right); //y заменяем на правого потомка
            y->right = z->right;//подвешиваем правое поддерево z к y
            y->right->parent = y;
        }

        transplant(z, y);
        y->left = z->left;
        y->left->parent = y;
        y->color = z->color;
    }

    if (y_original_color == BLACK) {
        deleteFixup(x);
    }

    delete z;
    //cout << "Элемент \"" << value << "\" удалён из дереве" << endl;
}

bool RBTree::TSEARCH(const string& value) const {
    TNode* current = root;
    while (current != nil) {
        if (value == current->data) {
            //cout << "Элемент \"" << value << "\" найден в дереве" << endl;
            return true;
        }
        else if (value < current->data) {
            current = current->left;
        }
        else {
            current = current->right;
        }
    }
    //cout << "Элемент \"" << value << "\" не найден в дереве" << endl;
    return false;
}

string RBTree::TGET(const string& value) const {
    TNode* current = root;
    while (current != nil) {
        if (value == current->data) {
            return current->data;
        }
        else if (value < current->data) {
            current = current->left;
        }
        else {
            current = current->right;
        }
    }
    return "";
}

void RBTree::inorder(TNode* node) const {
    if (node != nil) {
        inorder(node->left);
        cout << "\"" << node->data << "\"(" << (node->color == RED ? "R" : "B") << ") ";
        inorder(node->right);
    }
}

void RBTree::preorder(TNode* node) const {
    if (node != nil) {
        cout << "\"" << node->data << "\"(" << (node->color == RED ? "R" : "B") << ") ";
        preorder(node->left);
        preorder(node->right);
    }
}

void RBTree::postorder(TNode* node) const {
    if (node != nil) {
        postorder(node->left);
        postorder(node->right);
        cout << "\"" << node->data << "\"(" << (node->color == RED ? "R" : "B") << ") ";
    }
}

void RBTree::printTree(TNode* node, const string& prefix, bool isLeft) const {
    if (node != nil) {
        cout << prefix;
        cout << (isLeft ? "|--" : "|__");
        cout << node->data << (node->color == RED ? " (R)" : " (B)") << endl;

        // Рекурсивно выводим левое и правое поддеревья
        printTree(node->left, prefix + (isLeft ? "|   " : "    "), true);
        printTree(node->right, prefix + (isLeft ? "|   " : "    "), false);
    }
}

void RBTree::TPRINT_INORDER() const {
    cout << "Красно-чёрное дерево (инфиксный обход): ";
    inorder(root);
    cout << endl;
}

void RBTree::TPRINT_PREORDER() const {
    cout << "Красно-чёрное дерево (префиксный обход): ";
    preorder(root);
    cout << endl;
}

void RBTree::TPRINT_POSTORDER() const {
    cout << "Красно-чёрное дерево (постфиксный обход): ";
    postorder(root);
    cout << endl;
}

void RBTree::TPRINT_TREE() const {
    if (root == nil) {
        cout << "Дерево пусто!" << endl;
        return;
    }

    cout << "\nКрасно-чёрное дерево (структура):" << endl;
    printTree(root, "", false);
}
// Реализация методов сериализации
TNode* RBTree::getRoot() const { 
    return root; 
}

TNode* RBTree::getNil() const { 
    return nil; 
}

std::vector<std::pair<std::string, Color>> RBTree::getAllNodes() const {
    std::vector<std::pair<std::string, Color>> result;
    if (root == nil) return result;
    
    std::queue<TNode*> nodeQueue;
    nodeQueue.push(root);
    
    while (!nodeQueue.empty()) {
        TNode* current = nodeQueue.front();
        nodeQueue.pop();
        
        result.push_back({current->data, current->color});
        
        if (current->left != nil) nodeQueue.push(current->left);
        if (current->right != nil) nodeQueue.push(current->right);
    }
    return result;
}

void RBTree::clear() {
    clearTree(root);
    root = nil;
}

void RBTree::setRoot(TNode* newRoot) {
    if (root != nil && root != newRoot) {
        clearTree(root);
    }
    root = newRoot;
    if (root != nil) {
        root->parent = nil;
    }
}

/*void RBTree::insertWithColor(const std::string& value, Color color) {
    TNode* z = createNode(value);
    z->color = color;
    
    TNode* y = nil;
    TNode* x = root;
    
    while (x != nil) {
        y = x;
        if (z->data < x->data) {
            x = x->left;
        } else {
            x = x->right;
        }
    }
    
    z->parent = y;
    if (y == nil) {
        root = z;
    } else if (z->data < y->data) {
        y->left = z;
    } else {
        y->right = z;
    }
    
    z->left = nil;
    z->right = nil;
    
    // Для сохранения свойств RB-дерева вызываем insertFixup
    // Но с заданным цветом это может нарушить свойства
    insertFixup(z);
}*/
void RBTree::insertWithColor(const std::string& value, Color color) {
    cout<<"IWC"<<endl;
    TNode* z = createNode(value);
    z->color = color;
    
    TNode* y = nil;
    TNode* x = root;
    
    while (x != nil) {
        y = x;
        if (z->data < x->data) {
            x = x->left;
        } else if (z->data > x->data) {
            x = x->right;
        } else {
            // Элемент уже существует
            delete z;
            cout << "Элемент \"" << value << "\" уже существует в дереве" << endl;
            return;
        }
    }
    
    z->parent = y;
    if (y == nil) {
        root = z;
    } else if (z->data < y->data) {
        y->left = z;
    } else {
        y->right = z;
    }
    
    z->left = nil;
    z->right = nil;
    
    // ВАЖНО: При вставке с заданным цветом
    // нам не нужно вызывать insertFixup, так как
    // мы предполагаем, что дерево уже сбалансировано
    // Но если вы вставляете отдельные узлы, это может сломать дерево
}