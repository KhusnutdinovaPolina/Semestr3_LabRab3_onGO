package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type HashTable struct {
	Table [10][]string `json:"table"`
}

// Конструктор
func NewHashTable() *HashTable {
	ht := &HashTable{}
	for i := range ht.Table {
		ht.Table[i] = []string{}
	}
	return ht
}

// Хеш-функция
func (h *HashTable) hash(key string) int {
	sum := 0
	for _, c := range key {
		sum += int(c)
	}
	return sum % 10
}

// Вставка / Обновление
func (h *HashTable) Insert(key, value string) {
	idx := h.hash(key)
	entry := key + ":" + value

	// Проверяем, есть ли уже такой ключ (обновление)
	for i, item := range h.Table[idx] {
		if strings.HasPrefix(item, key+":") {
			h.Table[idx][i] = entry
			return
		}
	}
	// Если нет — добавляем
	h.Table[idx] = append(h.Table[idx], entry)
}

// Получить значение по ключу
func (h *HashTable) Get(key string) string {
	idx := h.hash(key)
	for _, item := range h.Table[idx] {
		if strings.HasPrefix(item, key+":") {
			return strings.TrimPrefix(item, key+":")
		}
	}
	return ""
}

// Удалить по ключу
func (h *HashTable) Remove(key string) {
	idx := h.hash(key)
	for i, item := range h.Table[idx] {
		if strings.HasPrefix(item, key+":") {
			// Удаляем элемент из среза
			h.Table[idx] = append(h.Table[idx][:i], h.Table[idx][i+1:]...)
			return
		}
	}
}

// Печать таблицы
func (h *HashTable) Print() {
	fmt.Println("--- HashTable ---")
	for i, bucket := range h.Table {
		fmt.Printf("[%d]: %v\n", i, bucket)
	}
}

// Сохранение (JSON)
func (h *HashTable) SaveToFile(filename string) error {
	data, _ := json.MarshalIndent(h, "", "    ")
	return os.WriteFile(filename, data, 0644)
}

// Загрузка (JSON)
func (h *HashTable) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, h)
}

func (h *HashTable) SaveToBinary(filename string) error {
	file, _ := os.Create(filename)
	defer file.Close()
	encoder := gob.NewEncoder(file)
	return encoder.Encode(h)
}

func (h *HashTable) LoadFromBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	return decoder.Decode(h)
}
