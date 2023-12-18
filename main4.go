package main

import "fmt"

// Структура для узла бинарного дерева
type TreeNode struct {
    Value       int
    Left, Right *TreeNode
}

// Функция для создания нового узла дерева
func newNode(value int) *TreeNode {
    return &TreeNode{Value: value, Left: nil, Right: nil}
}

// Функция для вставки узла в бинарное дерево
func insert(root *TreeNode, value int) *TreeNode {
    if root == nil {
        return newNode(value)
    }

    if value < root.Value {
        root.Left = insert(root.Left, value)
    } else if value > root.Value {
        root.Right = insert(root.Right, value)
    }

    return root
}

// Функция для вывода дерева в порядке in-order (левый - корень - правый)
func inorderTraversal(root *TreeNode) {
    if root != nil {
        inorderTraversal(root.Left)
        fmt.Printf("%d ", root.Value)
        inorderTraversal(root.Right)
    }
}

func main() {
    // Создание пустого дерева
    var root *TreeNode

    // Вставка элементов в дерево
    values := []int{5, 3, 8, 2, 4, 7, 9}
    for _, value := range values {
        root = insert(root, value)
    }

    // Вывод дерева в порядке in-order
    fmt.Println("In-order traversal:")
    inorderTraversal(root)
}