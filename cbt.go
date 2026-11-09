package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
)

type TNode struct {
	Data  string `json:"data"`
	Left  *TNode `json:"left"`
	Right *TNode `json:"right"`
}

type CBT struct {
	Root  *TNode `json:"root"`
	Count int    `json:"count"`
}

// Конструктор
func NewCBT() *CBT {
	return &CBT{}
}

// Вспомогательный метод: поиск узла по индексу
func (t *CBT) getNodeByIndex(node *TNode, currentIndex, targetIndex int) *TNode {
	if node == nil {
		return nil
	}
	if currentIndex == targetIndex {
		return node
	}
	// Ищем в левом поддереве
	leftRes := t.getNodeByIndex(node.Left, 2*currentIndex+1, targetIndex)
	if leftRes != nil {
		return leftRes
	}
	// Если слева нет, ищем в правом
	return t.getNodeByIndex(node.Right, 2*currentIndex+2, targetIndex)
}

// Вставка нового элемента
func (t *CBT) Insert(val string) {
	newNode := &TNode{Data: val}
	if t.Count == 0 {
		t.Root = newNode
		t.Count = 1
		return
	}
	newIndex := t.Count
	parentIndex := (newIndex - 1) / 2
	isLeft := (newIndex%2 == 1)
	parent := t.getNodeByIndex(t.Root, 0, parentIndex)

	if isLeft {
		parent.Left = newNode
	} else {
		parent.Right = newNode
	}
	t.Count++
}

// Проверка на Complete
func (t *CBT) IsComplete() bool {
	if t.Root == nil {
		return true
	}

	// Используем массив как очередь
	queue := make([]*TNode, 0)
	queue = append(queue, t.Root)
	seenNull := false

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:] // Dequeue

		if curr != nil {
			if seenNull {
				return false // Нашли узел после пустоты
			}
			queue = append(queue, curr.Left)
			queue = append(queue, curr.Right)
		} else {
			seenNull = true
		}
	}
	return true
}

// Размер дерева
func (t *CBT) Size() int {
	return t.Count
}

// Вспомогательный метод для очистки
func (t *CBT) clearRecursive(node *TNode) {
	if node == nil {
		return
	}
	t.clearRecursive(node.Left)
	t.clearRecursive(node.Right)
	node = nil
}

// Очистка
func (t *CBT) Clear() {
	t.clearRecursive(t.Root)
	t.Root = nil
	t.Count = 0
}

// Вспомогательный метод для печати
func (t *CBT) printRecursive(node *TNode, prefix string, isLast bool) {
	if node == nil {
		return
	}
	fmt.Printf("%s", prefix)
	if isLast {
		fmt.Print("+-- ")
	} else {
		fmt.Print("|-- ")
	}
	fmt.Println(node.Data)

	newPrefix := prefix
	if isLast {
		newPrefix += "    "
	} else {
		newPrefix += "|   "
	}

	hasL := node.Left != nil
	hasR := node.Right != nil
	if hasL {
		t.printRecursive(node.Left, newPrefix, !hasR)
	}
	if hasR {
		t.printRecursive(node.Right, newPrefix, true)
	}
}

// Печать дерева
func (t *CBT) Print() {
	if t.Root == nil {
		fmt.Println("(дерево пусто)")
		return
	}
	t.printRecursive(t.Root, "", true)
}

// Сохранение (JSON)
func (t *CBT) SaveToFile(filename string) error {
	data, _ := json.MarshalIndent(t, "", "    ")
	return os.WriteFile(filename, data, 0644)
}

// Загрузка (JSON)
func (t *CBT) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, t)
}
func (t *CBT) SaveToBinary(filename string) error {
	file, _ := os.Create(filename)
	defer file.Close()
	encoder := gob.NewEncoder(file)
	return encoder.Encode(t)
}

func (t *CBT) LoadFromBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	return decoder.Decode(t)
}
