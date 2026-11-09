package main

import (
	"os"
	"testing"
)

// Вспомогательная функция для отключения вывода в консоль
func silenceOutput() func() {
	old := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	return func() { os.Stdout = old }
}

func TestArray(t *testing.T) {
	restore := silenceOutput()
	defer restore()

	arr := NewArray()
	arr.PushBack("A")
	arr.PushBack("B")
	arr.InsertAt(1, "X") // Вставка по индексу
	_ = arr.Get(0)
	_ = arr.Get(99)
	arr.Set(1, "Y")
	arr.RemoveAt(0) // Удаление по индексу
	_ = arr.Length()
	arr.Print()
	arr.Clear()

	// JSON
	arr.SaveToFile("test_a.json")
	arr2 := NewArray()
	arr2.LoadFromFile("test_a.json")
	arr.SaveToFile("bad_dir/test.json")
	arr2.LoadFromFile("bad.json")

	// Gob
	arr.SaveToBinary("test_a.gob")
	arr3 := NewArray()
	arr3.LoadFromBinary("test_a.gob")
	arr.SaveToBinary("bad_dir/test.gob")
	arr3.LoadFromBinary("bad.gob")

	os.Remove("test_a.json")
	os.Remove("test_a.gob")
}

func TestSList(t *testing.T) {
	restore := silenceOutput()
	defer restore()

	// 1. Работа с пустым списком (покрывает ветки if l.Head == nil)
	l := NewSList()
	l.PopHead()              // Удаление из пустого
	l.PopTail()              // Удаление из пустого
	l.PopAfter("A")          // Удаление после несуществующего
	l.PopBefore("A")         // Удаление перед несуществующим
	l.InsertAfter("A", "X")  // Вставка после несуществующего
	l.InsertBefore("A", "X") // Вставка перед несуществующим

	// 2. Обычные операции
	l.PushHead("A")
	l.PushTail("B")
	l.PushTail("C") // A -> B -> C

	_ = l.Find("B")
	_ = l.Find("Z")
	_ = l.Get(1)
	_ = l.Get(99) // Выход за границы
	_ = l.Length()
	l.Print()

	// 3. Вставки / Удаления (покрывают середину)
	l.InsertAfter("A", "X")  // A -> X -> B -> C
	l.InsertBefore("B", "Y") // A -> X -> Y -> B -> C
	l.PopAfter("X")          // A -> X -> B -> C
	l.PopBefore("C")         // A -> X -> B -> C

	// 4. Вставка в голову / хвост (покрывают границы)
	l.InsertBefore("A", "Z") // Z -> A -> X -> B -> C
	l.InsertAfter("C", "W")  // Z -> A -> X -> B -> C -> W

	// 5. Удаление (покрывает PopHead / PopTail)
	l.PopHead()
	l.PopTail()

	// JSON
	l.SaveToFile("test_s.json")
	l2 := NewSList()
	l2.LoadFromFile("test_s.json")
	l.SaveToFile("bad_dir/test.json")
	l2.LoadFromFile("bad.json")

	// Gob
	l.SaveToBinary("test_s.gob")
	l3 := NewSList()
	l3.LoadFromBinary("test_s.gob")
	l.SaveToBinary("bad_dir/test.gob")
	l3.LoadFromBinary("bad.gob")

	// 7. Очистка
	l.Clear()
	os.Remove("test_s.json")
}

func TestDList(t *testing.T) {
	restore := silenceOutput()
	defer restore()

	d := NewDList()
	d.PushHead("A")
	d.PushTail("B")
	d.PushTail("C") // A <-> B <-> C

	_ = d.Find("B")
	_ = d.Find("Z")
	_ = d.Get(1)
	_ = d.Length()
	d.PrintForward()
	d.PrintBackward()

	d.InsertAfter("A", "X")  // A <-> X <-> B <-> C
	d.InsertBefore("B", "Y") // A <-> X <-> Y <-> B <-> C
	d.PopAfter("X")          // A <-> X <-> B <-> C
	d.PopBefore("C")         // A <-> X <-> B <-> C

	d.PopHead() // X <-> B <-> C
	d.PopTail() // X <-> B

	// JSON
	d.SaveToFile("test_d.json")
	d2 := NewDList()
	d2.LoadFromFile("test_d.json")
	d.SaveToFile("bad_dir/test.json")
	d2.LoadFromFile("bad.json")

	// Gob
	d.SaveToBinary("test_d.gob")
	d3 := NewDList()
	d3.LoadFromBinary("test_d.gob")
	d.SaveToBinary("bad_dir/test.gob")
	d3.LoadFromBinary("bad.gob")

	d.Clear()
	os.Remove("test_d.json")
	os.Remove("test_d.gob")

}

func TestStack(t *testing.T) {
	restore := silenceOutput()
	defer restore()

	s := NewStack()
	s.Push("A")
	s.Push("B")

	_ = s.GetSize()
	_ = s.Top()
	_ = s.IsEmpty()
	s.Print()

	s.Pop()
	s.Pop()
	_ = s.Pop() // Пустой
	_ = s.Top() // Пустой
	_ = s.IsEmpty()

	// JSON
	s.SaveToFile("test_st.json")
	s2 := NewStack()
	s2.LoadFromFile("test_st.json")
	s.SaveToFile("bad_dir/test.json")
	s2.LoadFromFile("bad.json")

	// Gob
	s.SaveToBinary("test_st.gob")
	s3 := NewStack()
	s3.LoadFromBinary("test_st.gob")
	s.SaveToBinary("bad_dir/test.gob")
	s3.LoadFromBinary("bad.gob")

	s.Clear()
	os.Remove("test_st.json")
	os.Remove("test_st.gob")
}

func TestQueue(t *testing.T) {
	restore := silenceOutput()
	defer restore()

	q := NewQueue()
	q.Push("A")
	q.Push("B")

	_ = q.GetSize()
	_ = q.Front()
	_ = q.IsEmpty()
	q.Print()

	q.Pop()
	q.Pop()
	_ = q.Pop()   // Пустая
	_ = q.Front() // Пустая
	_ = q.IsEmpty()

	// JSON
	q.SaveToFile("test_q.json")
	q2 := NewQueue()
	q2.LoadFromFile("test_q.json")
	q.SaveToFile("bad_dir/test.json")
	q2.LoadFromFile("bad.json")

	// Gob
	q.SaveToBinary("test_q.gob")
	q3 := NewQueue()
	q3.LoadFromBinary("test_q.gob")
	q.SaveToBinary("bad_dir/test.gob")
	q3.LoadFromBinary("bad.gob")

	q.Clear()
	os.Remove("test_q.json")
	os.Remove("test_q.gob")
}

func TestHashTable(t *testing.T) {
	restore := silenceOutput()
	defer restore()

	h := NewHashTable()
	h.Insert("apple", "fruit")
	h.Insert("apple", "fruit2")
	h.Insert("banana", "fruit3")

	_ = h.Get("apple")
	_ = h.Get("none")
	h.Print()

	h.Remove("banana")
	_ = h.Get("banana")

	// JSON
	h.SaveToFile("test_h.json")
	h2 := NewHashTable()
	h2.LoadFromFile("test_h.json")
	h.SaveToFile("bad_dir/test.json")
	h2.LoadFromFile("bad.json")

	// Gob
	h.SaveToBinary("test_h.gob")
	h3 := NewHashTable()
	h3.LoadFromBinary("test_h.gob")
	h.SaveToBinary("bad_dir/test.gob")
	h3.LoadFromBinary("bad.gob")

	os.Remove("test_h.json")
	os.Remove("test_h.gob")
}

func TestCuckoo(t *testing.T) {
	restore := silenceOutput()
	defer restore()

	c := NewCuckooHashTable()
	c.Insert("a", "1")
	c.Insert("k", "2")
	c.Insert("a", "3")

	_ = c.Get("a")
	_ = c.Get("none")
	c.Print()

	// JSON
	c.SaveToFile("test_c.json")
	c2 := NewCuckooHashTable()
	c2.LoadFromFile("test_c.json")
	c.SaveToFile("bad_dir/test.json")
	c2.LoadFromFile("bad.json")

	// Gob
	c.SaveToBinary("test_c.gob")
	c3 := NewCuckooHashTable()
	c3.LoadFromBinary("test_c.gob")
	c.SaveToBinary("bad_dir/test.gob")
	c3.LoadFromBinary("bad.gob")

	os.Remove("test_c.json")
	os.Remove("test_c.gob")
}

func TestCBT(t *testing.T) {
	restore := silenceOutput()
	defer restore()

	tree := NewCBT()
	tree.Print()
	tree.Insert("Root")
	tree.Insert("Left")
	tree.Insert("Right")
	tree.Insert("LeftLeft")
	tree.Insert("RightRight")
	tree.Print()

	_ = tree.Size()
	_ = tree.IsComplete()

	// JSON
	tree.SaveToFile("test_cbt.json")
	tree2 := NewCBT()
	tree2.LoadFromFile("test_cbt.json")
	tree.SaveToFile("bad_dir/test.json")
	tree2.LoadFromFile("bad.json")

	// Gob
	tree.SaveToBinary("test_cbt.gob")
	tree3 := NewCBT()
	tree3.LoadFromBinary("test_cbt.gob")
	tree.SaveToBinary("bad_dir/test.gob")
	tree3.LoadFromBinary("bad.gob")

	tree.Clear()
	os.Remove("test_cbt.json")
	os.Remove("test_cbt.gob")
}

func TestCLI(t *testing.T) {
	restore := silenceOutput()
	defer restore()

	RunCLI([]string{}) // Без аргументов

	// Массив
	RunCLI([]string{"--file", "test_cli.json", "--query", "MPUSH", "Apple"})
	RunCLI([]string{"--file", "test_cli.json", "--query", "MINSERT", "Banana"})
	RunCLI([]string{"--file", "test_cli.json", "--query", "MREMOVE", "1"})
	RunCLI([]string{"--file", "test_cli.json", "--query", "MPRINT"})

	// Списки
	RunCLI([]string{"--file", "test_cli.json", "--query", "FPUSH", "A"})
	RunCLI([]string{"--file", "test_cli.json", "--query", "FPOP"})
	RunCLI([]string{"--file", "test_cli.json", "--query", "FPRINT"})

	// Стек
	RunCLI([]string{"--file", "test_cli.json", "--query", "SPUSH", "A"})
	RunCLI([]string{"--file", "test_cli.json", "--query", "SPOP"})
	RunCLI([]string{"--file", "test_cli.json", "--query", "STOP"})
	RunCLI([]string{"--file", "test_cli.json", "--query", "SPRINT"})

	// Очередь
	RunCLI([]string{"--file", "test_cli.json", "--query", "QPUSH", "B"})
	RunCLI([]string{"--file", "test_cli.json", "--query", "QPOP"})
	RunCLI([]string{"--file", "test_cli.json", "--query", "QFRONT"})
	RunCLI([]string{"--file", "test_cli.json", "--query", "QPRINT"})

	// Дерево
	RunCLI([]string{"--file", "test_cli.json", "--query", "TINSERT", "Root"})
	RunCLI([]string{"--file", "test_cli.json", "--query", "TCHECK"})
	RunCLI([]string{"--file", "test_cli.json", "--query", "TPRINT"})

	// Хеш-таблица
	RunCLI([]string{"--file", "test_cli.json", "--query", "HSET", "name", "Polina"})
	RunCLI([]string{"--file", "test_cli.json", "--query", "HGET", "name"})
	RunCLI([]string{"--file", "test_cli.json", "--query", "HDEL", "name"})
	RunCLI([]string{"--file", "test_cli.json", "--query", "HPRINT"})

	// Кукушка
	RunCLI([]string{"--file", "test_cli.json", "--query", "CSET", "city", "Moscow"})
	RunCLI([]string{"--file", "test_cli.json", "--query", "CGET", "city"})
	RunCLI([]string{"--file", "test_cli.json", "--query", "CPRINT"})

	// Ошибки (недостаточно аргументов)
	RunCLI([]string{"--file", "test_cli.json", "--query", "MPUSH"})
	RunCLI([]string{"--file", "test_cli.json", "--query", "UNKNOWN"})

	os.Remove("test_cli.json")
}
