package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
)

type SNode struct {
	Data string `json:"data"`
	Next *SNode `json:"next"`
}

type SList struct {
	Head *SNode `json:"head"`
	Size int    `json:"size"`
}

func NewSList() *SList { return &SList{} }

func (l *SList) PushHead(val string) {
	l.Head = &SNode{Data: val, Next: l.Head}
	l.Size++
}

func (l *SList) PushTail(val string) {
	newNode := &SNode{Data: val}
	if l.Head == nil {
		l.Head = newNode
	} else {
		curr := l.Head
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = newNode
	}
	l.Size++
}

func (l *SList) PopHead() {
	if l.Head == nil {
		return
	}
	l.Head = l.Head.Next
	l.Size--
}

func (l *SList) PopTail() {
	if l.Head == nil {
		return
	}
	if l.Head.Next == nil {
		l.Head = nil
		l.Size--
		return
	}
	curr := l.Head
	for curr.Next.Next != nil {
		curr = curr.Next
	}
	curr.Next = nil
	l.Size--
}

// Вставить после узла target
func (l *SList) InsertAfter(target, val string) {
	curr := l.Head
	for curr != nil && curr.Data != target {
		curr = curr.Next
	}
	if curr == nil {
		return
	}
	newNode := &SNode{Data: val, Next: curr.Next}
	curr.Next = newNode
	l.Size++
}

// Вставить перед узлом target
func (l *SList) InsertBefore(target, val string) {
	if l.Head == nil {
		return
	}
	if l.Head.Data == target {
		l.PushHead(val)
		return
	}
	curr := l.Head
	for curr.Next != nil && curr.Next.Data != target {
		curr = curr.Next
	}
	if curr.Next == nil {
		return
	}
	newNode := &SNode{Data: val, Next: curr.Next}
	curr.Next = newNode
	l.Size++
}

// Удалить ПОСЛЕ узла target
func (l *SList) PopAfter(target string) {
	curr := l.Head
	for curr != nil && curr.Data != target {
		curr = curr.Next
	}
	if curr == nil || curr.Next == nil {
		return
	}
	curr.Next = curr.Next.Next
	l.Size--
}

// Удалить ПЕРЕД узлом target
func (l *SList) PopBefore(target string) {
	if l.Head == nil || l.Head.Data == target {
		return
	}
	if l.Head.Next != nil && l.Head.Next.Data == target {
		l.PopHead()
		return
	}
	curr := l.Head
	for curr.Next != nil && curr.Next.Next != nil && curr.Next.Next.Data != target {
		curr = curr.Next
	}
	if curr.Next == nil || curr.Next.Next == nil {
		return
	}
	curr.Next = curr.Next.Next
	l.Size--
}

func (l *SList) Find(val string) bool {
	curr := l.Head
	for curr != nil {
		if curr.Data == val {
			return true
		}
		curr = curr.Next
	}
	return false
}

func (l *SList) Get(index int) string {
	if index < 0 || index >= l.Size {
		return ""
	}
	curr := l.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	return curr.Data
}

func (l *SList) Length() int { return l.Size }
func (l *SList) Clear()      { l.Head = nil; l.Size = 0 }

func (l *SList) Print() {
	curr := l.Head
	fmt.Print("[")
	for curr != nil {
		fmt.Print(curr.Data)
		if curr.Next != nil {
			fmt.Print(" -> ")
		}
		curr = curr.Next
	}
	fmt.Println("]")
}

func (l *SList) SaveToFile(filename string) error {
	data, _ := json.MarshalIndent(l, "", "    ")
	return os.WriteFile(filename, data, 0644)
}

func (l *SList) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, l)

}

func (l *SList) SaveToBinary(filename string) error {
	file, _ := os.Create(filename)
	defer file.Close()
	encoder := gob.NewEncoder(file)
	return encoder.Encode(l)
}

func (l *SList) LoadFromBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	return decoder.Decode(l)
}
