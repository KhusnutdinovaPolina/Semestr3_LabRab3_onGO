package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
)

const SIZE = 10
const MAX_KICKS = 20

type CuckooHashTable struct {
	Keys1 [SIZE]string `json:"keys1"`
	Vals1 [SIZE]string `json:"vals1"`
	Keys2 [SIZE]string `json:"keys2"`
	Vals2 [SIZE]string `json:"vals2"`
}

// Конструктор
func NewCuckooHashTable() *CuckooHashTable {
	return &CuckooHashTable{}
}

// Первая хеш-функция
func (c *CuckooHashTable) hash1(key string) int {
	sum := 0
	for _, ch := range key {
		sum += int(ch)
	}
	return sum % SIZE
}

// Вторая хеш-функция
func (c *CuckooHashTable) hash2(key string) int {
	sum := 0
	for _, ch := range key {
		sum += int(ch) * 7
	}
	return sum % SIZE
}

// Вставка (с выбиванием)
func (c *CuckooHashTable) Insert(key, value string) bool {
	// Если ключ уже есть в первой таблице — обновляем
	if c.Keys1[c.hash1(key)] == key {
		c.Vals1[c.hash1(key)] = value
		return true
	}
	// Если во второй — обновляем
	if c.Keys2[c.hash2(key)] == key {
		c.Vals2[c.hash2(key)] = value
		return true
	}

	curKey, curVal := key, value
	for i := 0; i < MAX_KICKS; i++ {
		// Пробуем вставить в таблицу 1
		pos1 := c.hash1(curKey)
		if c.Keys1[pos1] == "" {
			c.Keys1[pos1], c.Vals1[pos1] = curKey, curVal
			return true
		}
		// Место занято — выбиваем
		curKey, c.Keys1[pos1] = c.Keys1[pos1], curKey
		curVal, c.Vals1[pos1] = c.Vals1[pos1], curVal

		// Пробуем вставить в таблицу 2
		pos2 := c.hash2(curKey)
		if c.Keys2[pos2] == "" {
			c.Keys2[pos2], c.Vals2[pos2] = curKey, curVal
			return true
		}
		// Место занято — выбиваем
		curKey, c.Keys2[pos2] = c.Keys2[pos2], curKey
		curVal, c.Vals2[pos2] = c.Vals2[pos2], curVal
	}
	return false // Переполнение
}

// Поиск по ключу
func (c *CuckooHashTable) Get(key string) string {
	pos1 := c.hash1(key)
	if c.Keys1[pos1] == key {
		return c.Vals1[pos1]
	}
	pos2 := c.hash2(key)
	if c.Keys2[pos2] == key {
		return c.Vals2[pos2]
	}
	return "" // Не найдено
}

// Печать обеих таблиц
func (c *CuckooHashTable) Print() {
	fmt.Println("--- Cuckoo (Table 1) ---")
	for i := 0; i < SIZE; i++ {
		if c.Keys1[i] != "" {
			fmt.Printf("[%d]: %s=%s\n", i, c.Keys1[i], c.Vals1[i])
		}
	}
	fmt.Println("--- Cuckoo (Table 2) ---")
	for i := 0; i < SIZE; i++ {
		if c.Keys2[i] != "" {
			fmt.Printf("[%d]: %s=%s\n", i, c.Keys2[i], c.Vals2[i])
		}
	}
}

// Сохранение (JSON)
func (c *CuckooHashTable) SaveToFile(filename string) error {
	data, _ := json.MarshalIndent(c, "", "    ")
	return os.WriteFile(filename, data, 0644)
}

// Загрузка (JSON)
func (c *CuckooHashTable) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, c)
}

func (c *CuckooHashTable) SaveToBinary(filename string) error {
	file, _ := os.Create(filename)
	defer file.Close()
	encoder := gob.NewEncoder(file)
	return encoder.Encode(c)
}

func (c *CuckooHashTable) LoadFromBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	return decoder.Decode(c)
}
