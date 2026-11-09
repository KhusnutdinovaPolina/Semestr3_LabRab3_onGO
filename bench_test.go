package main

import "testing"

// 1. Массив: добавление в конец
func BenchmarkArrayPushBack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := NewArray()
		for j := 0; j < 1000; j++ {
			arr.PushBack("test")
		}
	}
}

// 2. Односвязный список: добавление в начало
func BenchmarkSListPushHead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := NewSList()
		for j := 0; j < 1000; j++ {
			l.PushHead("test")
		}
	}
}

// 3. Двусвязный список: добавление в конец
func BenchmarkDListPushTail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := NewDList()
		for j := 0; j < 1000; j++ {
			l.PushTail("test")
		}
	}
}

// 4. Стек: push
func BenchmarkStackPush(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := NewStack()
		for j := 0; j < 1000; j++ {
			s.Push("test")
		}
	}
}

// 5. Очередь: push
func BenchmarkQueuePush(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := NewQueue()
		for j := 0; j < 1000; j++ {
			q.Push("test")
		}
	}
}

// 6. Хеш-таблица (цепочки): вставка
func BenchmarkHashTableInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := NewHashTable()
		for j := 0; j < 1000; j++ {
			h.Insert("key"+string(rune(j)), "val")
		}
	}
}

// 7. Кукушка: вставка (ограничим 10 элементами, чтобы не зациклилась)
func BenchmarkCuckooInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := NewCuckooHashTable()
		for j := 0; j < 10; j++ {
			c.Insert("key"+string(rune(j)), "val")
		}
	}
}

// 8. Дерево CBT: вставка
func BenchmarkCBTInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := NewCBT()
		for j := 0; j < 1000; j++ {
			t.Insert("test")
		}
	}
}
