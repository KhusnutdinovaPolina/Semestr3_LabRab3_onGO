package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
)

type Array struct {
	Data []string `json:"data"`
}

// Конструктор
func NewArray() *Array {
	return &Array{Data: []string{}}
}

// Добавить в конец
func (a *Array) PushBack(val string) {
	a.Data = append(a.Data, val)
}

// Вставить по индексу (аналог insertAt)
func (a *Array) InsertAt(index int, val string) {
	if index < 0 || index > len(a.Data) {
		return
	}
	// Вставляем элемент в середину через срезы
	a.Data = append(a.Data[:index], append([]string{val}, a.Data[index:]...)...)
}

// Получить по индексу
func (a *Array) Get(index int) string {
	if index < 0 || index >= len(a.Data) {
		return ""
	}
	return a.Data[index]
}

// Заменить по индексу
func (a *Array) Set(index int, val string) {
	if index >= 0 && index < len(a.Data) {
		a.Data[index] = val
	}
}

// Удалить по индексу
func (a *Array) RemoveAt(index int) {
	if index >= 0 && index < len(a.Data) {
		a.Data = append(a.Data[:index], a.Data[index+1:]...)
	}
}

// Длина
func (a *Array) Length() int {
	return len(a.Data)
}

// Очистить
func (a *Array) Clear() {
	a.Data = []string{}
}

// Печать
func (a *Array) Print() {
	fmt.Println(a.Data)
}

// Сохранение (JSON)
func (a *Array) SaveToFile(filename string) error {
	data, _ := json.MarshalIndent(a, "", "    ")
	return os.WriteFile(filename, data, 0644)
}

// Загрузка (JSON)
func (a *Array) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, a)
}

// Бинарное сохранение (Gob)
func (a *Array) SaveToBinary(filename string) error {
	file, _ := os.Create(filename)
	defer file.Close()
	encoder := gob.NewEncoder(file)
	return encoder.Encode(a)
}

// Бинарная загрузка (Gob)
func (a *Array) LoadFromBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	return decoder.Decode(a)
}
