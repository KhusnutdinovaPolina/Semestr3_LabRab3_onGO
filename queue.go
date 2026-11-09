package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
)

type Queue struct {
	Items []string `json:"items"`
}

// Конструктор
func NewQueue() *Queue {
	return &Queue{Items: []string{}}
}

// Добавить в конец
func (q *Queue) Push(val string) {
	q.Items = append(q.Items, val)
}

// Извлечь из начала
func (q *Queue) Pop() string {
	if len(q.Items) == 0 {
		return ""
	}
	val := q.Items[0]
	q.Items = q.Items[1:] // Сдвигаем срез
	return val
}

// Посмотреть начало
func (q *Queue) Front() string {
	if len(q.Items) == 0 {
		return ""
	}
	return q.Items[0]
}

// Проверить на пустоту
func (q *Queue) IsEmpty() bool {
	return len(q.Items) == 0
}

// Получить размер
func (q *Queue) GetSize() int {
	return len(q.Items)
}

// Очистить
func (q *Queue) Clear() {
	q.Items = []string{}
}

// Печать (от начала к концу)
func (q *Queue) Print() {
	fmt.Println(q.Items)
}

// Сохранение (JSON)
func (q *Queue) SaveToFile(filename string) error {
	data, _ := json.MarshalIndent(q, "", "    ")
	return os.WriteFile(filename, data, 0644)
}

// Загрузка (JSON)
func (q *Queue) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, q)
}

func (q *Queue) SaveToBinary(filename string) error {
	file, _ := os.Create(filename)
	defer file.Close()
	encoder := gob.NewEncoder(file)
	return encoder.Encode(q)
}

func (q *Queue) LoadFromBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	return decoder.Decode(q)
}
