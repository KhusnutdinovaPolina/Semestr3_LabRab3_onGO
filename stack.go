package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
)

type Stack struct {
	Items []string `json:"items"`
}

// Конструктор
func NewStack() *Stack {
	return &Stack{Items: []string{}}
}

// Добавить на вершину
func (s *Stack) Push(val string) {
	s.Items = append(s.Items, val)
}

// Снять с вершины (и вернуть)
func (s *Stack) Pop() string {
	if len(s.Items) == 0 {
		return ""
	}
	val := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return val
}

// Посмотреть вершину
func (s *Stack) Top() string {
	if len(s.Items) == 0 {
		return ""
	}
	return s.Items[len(s.Items)-1]
}

// Проверить на пустоту
func (s *Stack) IsEmpty() bool {
	return len(s.Items) == 0
}

// Получить размер
func (s *Stack) GetSize() int {
	return len(s.Items)
}

// Очистить
func (s *Stack) Clear() {
	s.Items = []string{}
}

// Печать (сверху вниз)
func (s *Stack) Print() {
	// Печатаем в обратном порядке, так как вершина в конце среза
	fmt.Print("[")
	for i := len(s.Items) - 1; i >= 0; i-- {
		fmt.Print(s.Items[i])
		if i > 0 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

// Сохранение (JSON)
func (s *Stack) SaveToFile(filename string) error {
	data, _ := json.MarshalIndent(s, "", "    ")
	return os.WriteFile(filename, data, 0644)
}

// Загрузка (JSON)
func (s *Stack) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, s)
}

func (s *Stack) SaveToBinary(filename string) error {
	file, _ := os.Create(filename)
	defer file.Close()
	encoder := gob.NewEncoder(file)
	return encoder.Encode(s)
}

func (s *Stack) LoadFromBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	return decoder.Decode(s)
}
