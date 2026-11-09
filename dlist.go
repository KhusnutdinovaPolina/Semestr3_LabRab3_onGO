package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
)

type DNode struct {
	Data string `json:"data"`
	Prev *DNode `json:"prev"`
	Next *DNode `json:"next"`
}

type DList struct {
	Head *DNode `json:"head"`
	Tail *DNode `json:"tail"`
	Size int    `json:"size"`
}

func NewDList() *DList { return &DList{} }

func (l *DList) PushHead(val string) {
	newNode := &DNode{Data: val, Next: l.Head}
	if l.Head != nil {
		l.Head.Prev = newNode
	}
	l.Head = newNode
	if l.Tail == nil {
		l.Tail = newNode
	}
	l.Size++
}

func (l *DList) PushTail(val string) {
	newNode := &DNode{Data: val, Prev: l.Tail}
	if l.Tail != nil {
		l.Tail.Next = newNode
	}
	l.Tail = newNode
	if l.Head == nil {
		l.Head = newNode
	}
	l.Size++
}

// Вставить после
func (l *DList) InsertAfter(target, val string) {
	curr := l.Head
	for curr != nil && curr.Data != target {
		curr = curr.Next
	}
	if curr == nil {
		return
	}
	newNode := &DNode{Data: val, Prev: curr, Next: curr.Next}
	if curr.Next != nil {
		curr.Next.Prev = newNode
	} else {
		l.Tail = newNode
	}
	curr.Next = newNode
	l.Size++
}

// Вставить перед
func (l *DList) InsertBefore(target, val string) {
	curr := l.Head
	for curr != nil && curr.Data != target {
		curr = curr.Next
	}
	if curr == nil {
		return
	}
	newNode := &DNode{Data: val, Prev: curr.Prev, Next: curr}
	if curr.Prev != nil {
		curr.Prev.Next = newNode
	} else {
		l.Head = newNode
	}
	curr.Prev = newNode
	l.Size++
}

// Удалить после
func (l *DList) PopAfter(target string) {
	curr := l.Head
	for curr != nil && curr.Data != target {
		curr = curr.Next
	}
	if curr == nil || curr.Next == nil {
		return
	}
	toDelete := curr.Next
	curr.Next = toDelete.Next
	if toDelete.Next != nil {
		toDelete.Next.Prev = curr
	} else {
		l.Tail = curr
	}
	l.Size--
}

// Удалить перед
func (l *DList) PopBefore(target string) {
	curr := l.Head
	for curr != nil && curr.Data != target {
		curr = curr.Next
	}
	if curr == nil || curr.Prev == nil {
		return
	}
	toDelete := curr.Prev
	curr.Prev = toDelete.Prev
	if toDelete.Prev != nil {
		toDelete.Prev.Next = curr
	} else {
		l.Head = curr
	}
	l.Size--
}

func (l *DList) PopHead() {
	if l.Head == nil {
		return
	}
	l.Head = l.Head.Next
	if l.Head != nil {
		l.Head.Prev = nil
	} else {
		l.Tail = nil
	}
	l.Size--
}

func (l *DList) PopTail() {
	if l.Tail == nil {
		return
	}
	l.Tail = l.Tail.Prev
	if l.Tail != nil {
		l.Tail.Next = nil
	} else {
		l.Head = nil
	}
	l.Size--
}

func (l *DList) Find(val string) bool {
	curr := l.Head
	for curr != nil {
		if curr.Data == val {
			return true
		}
		curr = curr.Next
	}
	return false
}

func (l *DList) Get(index int) string {
	if index < 0 || index >= l.Size {
		return ""
	}
	curr := l.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	return curr.Data
}

func (l *DList) Length() int { return l.Size }
func (l *DList) Clear()      { l.Head = nil; l.Tail = nil; l.Size = 0 }

func (l *DList) PrintForward() {
	curr := l.Head
	fmt.Print("[")
	for curr != nil {
		fmt.Print(curr.Data)
		if curr.Next != nil {
			fmt.Print(" <-> ")
		}
		curr = curr.Next
	}
	fmt.Println("]")
}

func (l *DList) PrintBackward() {
	curr := l.Tail
	fmt.Print("[")
	for curr != nil {
		fmt.Print(curr.Data)
		if curr.Prev != nil {
			fmt.Print(" <-> ")
		}
		curr = curr.Prev
	}
	fmt.Println("]")
}

func (l *DList) SaveToFile(filename string) error {
	data, _ := json.MarshalIndent(l, "", "    ")
	return os.WriteFile(filename, data, 0644)
}

func (l *DList) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, l)
}

// Бинарное сохранение (Gob)
func (l *DList) SaveToBinary(filename string) error {
	file, _ := os.Create(filename)
	defer file.Close()
	encoder := gob.NewEncoder(file)
	return encoder.Encode(l)
}

// Бинарная загрузка (Gob)
func (l *DList) LoadFromBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	return decoder.Decode(l)
}
