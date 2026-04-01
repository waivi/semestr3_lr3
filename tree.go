package main

import (
	"fmt"
)

type Color int

const (
	RED Color = iota
	BLACK
)

type TNode struct {
	data   string
	color  Color
	left   *TNode
	right  *TNode
	parent *TNode
}

type RBTree struct {
	root *TNode
	nil  *TNode // NIL-лист
}

func NewRBTree() *RBTree {
	nilNode := &TNode{
		data:   "",
		color:  BLACK,
		left:   nil,
		right:  nil,
		parent: nil,
	}
	return &RBTree{
		root: nilNode,
		nil:  nilNode,
	}
}

func (t *RBTree) createNode(value string) *TNode {
	return &TNode{
		data:   value,
		color:  RED,
		left:   t.nil,
		right:  t.nil,
		parent: t.nil,
	}
}

func (t *RBTree) TINSERT(value string) {
	z := t.createNode(value)
	y := t.nil
	x := t.root

	// Поиск места для вставки
	for x != t.nil {
		y = x
		if z.data < x.data {
			x = x.left
		} else if z.data > x.data {
			x = x.right
		} else {
			// Элемент уже существует
			fmt.Printf("Элемент \"%s\" уже существует в дереве\n", value)
			return
		}
	}

	// Вставляем новый узел
	z.parent = y
	if y == t.nil {
		t.root = z
	} else if z.data < y.data {
		y.left = z
	} else {
		y.right = z
	}

	t.insertFixup(z)
}

func (t *RBTree) leftRotate(x *TNode) {
	y := x.right
	x.right = y.left

	if y.left != t.nil {
		y.left.parent = x
	}

	y.parent = x.parent

	if x.parent == t.nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x
	x.parent = y
}

func (t *RBTree) rightRotate(y *TNode) {
	x := y.left
	y.left = x.right

	if x.right != t.nil {
		x.right.parent = y
	}

	x.parent = y.parent

	if y.parent == t.nil {
		t.root = x
	} else if y == y.parent.right {
		y.parent.right = x
	} else {
		y.parent.left = x
	}

	x.right = y
	y.parent = x
}

func (t *RBTree) insertFixup(z *TNode) {
	for z.parent.color == RED {
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right // дядя

			if y.color == RED {
				// Случай 1
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					// Случай 2
					z = z.parent
					t.leftRotate(z)
				}
				// Случай 3
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.rightRotate(z.parent.parent)
			}
		} else {
			// Симметричный случай
			y := z.parent.parent.left

			if y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					t.rightRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.leftRotate(z.parent.parent)
			}
		}
	}
	t.root.color = BLACK
}

func (t *RBTree) minimum(node *TNode) *TNode {
	for node.left != t.nil {
		node = node.left
	}
	return node
}

func (t *RBTree) transplant(u, v *TNode) {
	if u.parent == t.nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

func (t *RBTree) deleteFixup(x *TNode) {
	for x != t.root && x.color == BLACK {
		if x == x.parent.left {
			w := x.parent.right

			if w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				t.leftRotate(x.parent)
				w = x.parent.right
			}

			if w.left.color == BLACK && w.right.color == BLACK {
				w.color = RED
				x = x.parent
			} else {
				if w.right.color == BLACK {
					w.left.color = BLACK
					w.color = RED
					t.rightRotate(w)
					w = x.parent.right
				}

				w.color = x.parent.color
				x.parent.color = BLACK
				w.right.color = BLACK
				t.leftRotate(x.parent)
				x = t.root
			}
		} else {
			// Симметричный случай
			w := x.parent.left

			if w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				t.rightRotate(x.parent)
				w = x.parent.left
			}

			if w.right.color == BLACK && w.left.color == BLACK {
				w.color = RED
				x = x.parent
			} else {
				if w.left.color == BLACK {
					w.right.color = BLACK
					w.color = RED
					t.leftRotate(w)
					w = x.parent.left
				}

				w.color = x.parent.color
				x.parent.color = BLACK
				w.left.color = BLACK
				t.rightRotate(x.parent)
				x = t.root
			}
		}
	}
	x.color = BLACK
}

func (t *RBTree) TDELETE(value string) {
	z := t.root
	// Поиск удаляемого элемента
	for z != t.nil {
		if value == z.data {
			break
		} else if value < z.data {
			z = z.left
		} else {
			z = z.right
		}
	}

	if z == t.nil {
		return
	}

	y := z
	yOriginalColor := y.color
	var x *TNode

	if z.left == t.nil {
		x = z.right
		t.transplant(z, z.right)
	} else if z.right == t.nil {
		x = z.left
		t.transplant(z, z.left)
	} else {
		y = t.minimum(z.right)
		yOriginalColor = y.color
		x = y.right

		if y.parent == z {
			x.parent = y
		} else {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}

		t.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}

	if yOriginalColor == BLACK {
		t.deleteFixup(x)
	}
}

func (t *RBTree) TSEARCH(value string) bool {
	current := t.root
	for current != t.nil {
		if value == current.data {
			return true
		} else if value < current.data {
			current = current.left
		} else {
			current = current.right
		}
	}
	return false
}

func (t *RBTree) TGET(value string) (string, error) {
	current := t.root
	for current != t.nil {
		if value == current.data {
			return current.data, nil
		} else if value < current.data {
			current = current.left
		} else {
			current = current.right
		}
	}
	return "", fmt.Errorf("элемент не найден")
}

func (t *RBTree) inorder(node *TNode, result *[]string) {
	if node != t.nil {
		t.inorder(node.left, result)
		*result = append(*result, node.data)
		t.inorder(node.right, result)
	}
}

func (t *RBTree) TPRINT_INORDER() {
	if t.root == t.nil {
		fmt.Println("Дерево пусто!")
		return
	}

	fmt.Print("Красно-чёрное дерево (инфиксный обход): ")
	result := make([]string, 0)
	t.inorder(t.root, &result)
	for i, val := range result {
		fmt.Printf("\"%s\"", val)
		if i < len(result)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func (t *RBTree) preorder(node *TNode) {
	if node != t.nil {
		fmt.Printf("\"%s\"(%s) ", node.data, colorToString(node.color))
		t.preorder(node.left)
		t.preorder(node.right)
	}
}

func (t *RBTree) TPRINT_PREORDER() {
	if t.root == t.nil {
		fmt.Println("Дерево пусто!")
		return
	}

	fmt.Print("Красно-чёрное дерево (префиксный обход): ")
	t.preorder(t.root)
	fmt.Println()
}

func (t *RBTree) postorder(node *TNode) {
	if node != t.nil {
		t.postorder(node.left)
		t.postorder(node.right)
		fmt.Printf("\"%s\"(%s) ", node.data, colorToString(node.color))
	}
}

func (t *RBTree) TPRINT_POSTORDER() {
	if t.root == t.nil {
		fmt.Println("Дерево пусто!")
		return
	}

	fmt.Print("Красно-чёрное дерево (постфиксный обход): ")
	t.postorder(t.root)
	fmt.Println()
}

func (t *RBTree) printTree(node *TNode, prefix string, isLeft bool) {
	if node != t.nil {
		fmt.Print(prefix)
		if isLeft {
			fmt.Print("|--")
		} else {
			fmt.Print("|__")
		}
		fmt.Printf("%s (%s)\n", node.data, colorToString(node.color))

		t.printTree(node.left, prefix+func() string {
			if isLeft {
				return "|   "
			}
			return "    "
		}(), true)
		t.printTree(node.right, prefix+func() string {
			if isLeft {
				return "|   "
			}
			return "    "
		}(), false)
	}
}

func (t *RBTree) TPRINT_TREE() {
	if t.root == t.nil {
		fmt.Println("Дерево пусто!")
		return
	}

	fmt.Println("\nКрасно-чёрное дерево (структура):")
	t.printTree(t.root, "", false)
}

func colorToString(color Color) string {
	if color == RED {
		return "R"
	}
	return "B"
}

func (t *RBTree) clearTree(node *TNode) {
	if node != t.nil {
		t.clearTree(node.left)
		t.clearTree(node.right)
	}
}

func (t *RBTree) Clear() {
	t.clearTree(t.root)
	t.root = t.nil
}

// Для сериализации
func (t *RBTree) GetAllElements() []string {
	result := make([]string, 0)
	if t.root == t.nil {
		return result
	}

	stack := []*TNode{t.root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node != t.nil {
			result = append(result, node.data)
			stack = append(stack, node.right, node.left)
		}
	}
	return result
}

func (t *RBTree) GetAllNodes() []struct {
	Data  string
	Color Color
} {
	result := make([]struct {
		Data  string
		Color Color
	}, 0)
	if t.root == t.nil {
		return result
	}

	stack := []*TNode{t.root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node != t.nil {
			result = append(result, struct {
				Data  string
				Color Color
			}{
				Data:  node.data,
				Color: node.color,
			})
			stack = append(stack, node.right, node.left)
		}
	}
	return result
}

func (t *RBTree) SetRoot(newRoot *TNode) {
	t.root = newRoot
}

func (t *RBTree) InsertWithColor(value string, color Color) {
	node := t.createNode(value)
	node.color = color

	y := t.nil
	x := t.root

	// Поиск места для вставки
	for x != t.nil {
		y = x
		if node.data < x.data {
			x = x.left
		} else {
			x = x.right
		}
	}

	// Вставляем новый узел
	node.parent = y
	if y == t.nil {
		t.root = node
	} else if node.data < y.data {
		y.left = node
	} else {
		y.right = node
	}
}

// InOrderTraversal возвращает элементы дерева в порядке возрастания
func (t *RBTree) InOrderTraversal() []string {
	result := make([]string, 0)
	t.inorderTraversalHelper(t.root, &result)
	return result
}

func (t *RBTree) inorderTraversalHelper(node *TNode, result *[]string) {
	if node != t.nil {
		t.inorderTraversalHelper(node.left, result)
		*result = append(*result, node.data)
		t.inorderTraversalHelper(node.right, result)
	}
}

// Min возвращает минимальный элемент дерева
func (t *RBTree) Min() string {
	if t.root == t.nil {
		return ""
	}
	node := t.minimum(t.root)
	return node.data
}

// Max возвращает максимальный элемент дерева
func (t *RBTree) Max() string {
	if t.root == t.nil {
		return ""
	}
	node := t.root
	for node.right != t.nil {
		node = node.right
	}
	return node.data
}
