package belajar

// ==================== STACK ====================

// Stack merepresentasikan struktur data stack (LIFO - Last In First Out).
type Stack struct {
	items []int
}

// NewStack membuat Stack baru yang kosong.
// Contoh: s := NewStack() -> stack kosong
func NewStack() *Stack {
	// TODO: implementasi di sini
	return nil
}

// Push menambahkan elemen ke puncak stack.
// Contoh: s.Push(10) -> stack: [10]
//
//	s.Push(20) -> stack: [10, 20]
func (s *Stack) Push(val int) {
	// TODO: implementasi di sini
}

// Pop menghapus dan mengembalikan elemen dari puncak stack.
// Kembalikan 0 dan true (error) jika stack kosong.
// Contoh: stack [10, 20] -> s.Pop() -> 20, false; stack menjadi [10]
//
//	stack [] -> s.Pop() -> 0, true
func (s *Stack) Pop() (int, bool) {
	// TODO: implementasi di sini
	return 0, true
}

// Peek mengembalikan elemen di puncak stack TANPA menghapusnya.
// Kembalikan 0 dan true (error) jika stack kosong.
// Contoh: stack [10, 20] -> s.Peek() -> 20, false; stack tetap [10, 20]
//
//	stack [] -> s.Peek() -> 0, true
func (s *Stack) Peek() (int, bool) {
	// TODO: implementasi di sini
	return 0, true
}

// IsEmpty mengecek apakah stack kosong.
// Contoh: NewStack().IsEmpty() -> true
func (s *Stack) IsEmpty() bool {
	// TODO: implementasi di sini
	return true
}

// Size mengembalikan jumlah elemen dalam stack.
// Contoh: stack [10, 20, 30] -> s.Size() -> 3
func (s *Stack) Size() int {
	// TODO: implementasi di sini
	return 0
}

// ToSlice mengembalikan isi stack sebagai slice (dari bawah ke atas).
// Contoh: stack [10, 20, 30] -> s.ToSlice() -> []int{10, 20, 30}
func (s *Stack) ToSlice() []int {
	// TODO: implementasi di sini
	return nil
}

// ==================== QUEUE ====================

// Queue merepresentasikan struktur data queue (FIFO - First In First Out).
type Queue struct {
	items []int
}

// NewQueue membuat Queue baru yang kosong.
// Contoh: q := NewQueue() -> queue kosong
func NewQueue() *Queue {
	// TODO: implementasi di sini
	return nil
}

// Enqueue menambahkan elemen ke belakang queue.
// Contoh: q.Enqueue(10) -> queue: [10]
//
//	q.Enqueue(20) -> queue: [10, 20]
func (q *Queue) Enqueue(val int) {
	// TODO: implementasi di sini
}

// Dequeue menghapus dan mengembalikan elemen dari depan queue.
// Kembalikan 0 dan true (error) jika queue kosong.
// Contoh: queue [10, 20] -> q.Dequeue() -> 10, false; queue menjadi [20]
//
//	queue [] -> q.Dequeue() -> 0, true
func (q *Queue) Dequeue() (int, bool) {
	// TODO: implementasi di sini
	return 0, true
}

// Peek mengembalikan elemen di depan queue TANPA menghapusnya.
// Kembalikan 0 dan true (error) jika queue kosong.
// Contoh: queue [10, 20] -> q.Peek() -> 10, false
func (q *Queue) Peek() (int, bool) {
	// TODO: implementasi di sini
	return 0, true
}

// IsEmpty mengecek apakah queue kosong.
func (q *Queue) IsEmpty() bool {
	// TODO: implementasi di sini
	return true
}

// Size mengembalikan jumlah elemen dalam queue.
func (q *Queue) Size() int {
	// TODO: implementasi di sini
	return 0
}

// ToSlice mengembalikan isi queue sebagai slice (dari depan ke belakang).
func (q *Queue) ToSlice() []int {
	// TODO: implementasi di sini
	return nil
}

// ==================== LINKED LIST ====================

// Node merepresentasikan satu node dalam singly linked list.
type Node struct {
	Value int
	Next  *Node
}

// LinkedList merepresentasikan singly linked list.
type LinkedList struct {
	Head *Node
	size int
}

// NewLinkedList membuat LinkedList baru yang kosong.
func NewLinkedList() *LinkedList {
	// TODO: implementasi di sini
	return nil
}

// Append menambahkan elemen di akhir linked list.
// Contoh: list [] -> Append(10) -> list [10]
//
//	list [10] -> Append(20) -> list [10, 20]
func (ll *LinkedList) Append(val int) {
	// TODO: implementasi di sini
}

// Prepend menambahkan elemen di awal linked list.
// Contoh: list [20, 30] -> Prepend(10) -> list [10, 20, 30]
func (ll *LinkedList) Prepend(val int) {
	// TODO: implementasi di sini
}

// DeleteByValue menghapus node pertama dengan value tertentu.
// Kembalikan true jika berhasil dihapus, false jika value tidak ditemukan.
// Contoh: list [10, 20, 30] -> DeleteByValue(20) -> true, list menjadi [10, 30]
//
//	list [10, 20, 30] -> DeleteByValue(99) -> false, list tetap [10, 20, 30]
func (ll *LinkedList) DeleteByValue(val int) bool {
	// TODO: implementasi di sini
	return false
}

// Contains mengecek apakah value ada di linked list.
// Contoh: list [10, 20, 30] -> Contains(20) -> true
//
//	list [10, 20, 30] -> Contains(99) -> false
func (ll *LinkedList) Contains(val int) bool {
	// TODO: implementasi di sini
	return false
}

// Size mengembalikan jumlah elemen dalam linked list.
func (ll *LinkedList) Size() int {
	// TODO: implementasi di sini
	return 0
}

// ToSlice mengonversi linked list ke slice int.
// Contoh: list [10, 20, 30] -> ToSlice() -> []int{10, 20, 30}
//
//	list [] -> ToSlice() -> []int{}
func (ll *LinkedList) ToSlice() []int {
	// TODO: implementasi di sini
	return nil
}

// Reverse membalik urutan elemen dalam linked list (in-place).
// Contoh: list [10, 20, 30] -> Reverse() -> list [30, 20, 10]
func (ll *LinkedList) Reverse() {
	// TODO: implementasi di sini
}
