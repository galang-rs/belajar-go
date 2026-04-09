package belajar

import (
	"reflect"
	"testing"
)

// =============================================================
// 1. TEST - Stack Operations
// =============================================================

func TestStack_Basic(t *testing.T) {
	s := NewStack()
	if s == nil {
		t.Fatal("NewStack() returned nil")
	}

	if !s.IsEmpty() {
		t.Error("new stack should be empty")
	}

	if s.Size() != 0 {
		t.Errorf("new stack Size() = %d; want 0", s.Size())
	}
}

func TestStack_PushPop(t *testing.T) {
	s := NewStack()
	if s == nil {
		t.Fatal("NewStack() returned nil")
	}

	s.Push(10)
	s.Push(20)
	s.Push(30)

	if s.Size() != 3 {
		t.Errorf("Size() = %d; want 3", s.Size())
	}

	if s.IsEmpty() {
		t.Error("stack should not be empty after Push")
	}

	// Pop should return LIFO order
	val, err := s.Pop()
	if err || val != 30 {
		t.Errorf("Pop() = (%d, %v); want (30, false)", val, err)
	}

	val, err = s.Pop()
	if err || val != 20 {
		t.Errorf("Pop() = (%d, %v); want (20, false)", val, err)
	}

	val, err = s.Pop()
	if err || val != 10 {
		t.Errorf("Pop() = (%d, %v); want (10, false)", val, err)
	}

	// Pop from empty
	val, err = s.Pop()
	if !err {
		t.Errorf("Pop() on empty stack should return error, got (%d, %v)", val, err)
	}
}

func TestStack_Peek(t *testing.T) {
	s := NewStack()
	if s == nil {
		t.Fatal("NewStack() returned nil")
	}

	// Peek on empty
	_, err := s.Peek()
	if !err {
		t.Error("Peek() on empty stack should return error")
	}

	s.Push(42)
	val, err := s.Peek()
	if err || val != 42 {
		t.Errorf("Peek() = (%d, %v); want (42, false)", val, err)
	}

	// Peek should not remove element
	if s.Size() != 1 {
		t.Errorf("Size() after Peek = %d; want 1", s.Size())
	}

	s.Push(99)
	val, err = s.Peek()
	if err || val != 99 {
		t.Errorf("Peek() = (%d, %v); want (99, false)", val, err)
	}
}

func TestStack_ToSlice(t *testing.T) {
	s := NewStack()
	if s == nil {
		t.Fatal("NewStack() returned nil")
	}

	// Empty stack
	result := s.ToSlice()
	if len(result) != 0 {
		t.Errorf("ToSlice() on empty = %v; want []", result)
	}

	s.Push(10)
	s.Push(20)
	s.Push(30)

	result = s.ToSlice()
	expected := []int{10, 20, 30}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ToSlice() = %v; want %v", result, expected)
	}
}

// =============================================================
// 2. TEST - Queue Operations
// =============================================================

func TestQueue_Basic(t *testing.T) {
	q := NewQueue()
	if q == nil {
		t.Fatal("NewQueue() returned nil")
	}

	if !q.IsEmpty() {
		t.Error("new queue should be empty")
	}

	if q.Size() != 0 {
		t.Errorf("new queue Size() = %d; want 0", q.Size())
	}
}

func TestQueue_EnqueueDequeue(t *testing.T) {
	q := NewQueue()
	if q == nil {
		t.Fatal("NewQueue() returned nil")
	}

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	if q.Size() != 3 {
		t.Errorf("Size() = %d; want 3", q.Size())
	}

	// Dequeue should return FIFO order
	val, err := q.Dequeue()
	if err || val != 10 {
		t.Errorf("Dequeue() = (%d, %v); want (10, false)", val, err)
	}

	val, err = q.Dequeue()
	if err || val != 20 {
		t.Errorf("Dequeue() = (%d, %v); want (20, false)", val, err)
	}

	val, err = q.Dequeue()
	if err || val != 30 {
		t.Errorf("Dequeue() = (%d, %v); want (30, false)", val, err)
	}

	// Dequeue from empty
	val, err = q.Dequeue()
	if !err {
		t.Errorf("Dequeue() on empty queue should return error, got (%d, %v)", val, err)
	}
}

func TestQueue_Peek(t *testing.T) {
	q := NewQueue()
	if q == nil {
		t.Fatal("NewQueue() returned nil")
	}

	// Peek on empty
	_, err := q.Peek()
	if !err {
		t.Error("Peek() on empty queue should return error")
	}

	q.Enqueue(42)
	val, err := q.Peek()
	if err || val != 42 {
		t.Errorf("Peek() = (%d, %v); want (42, false)", val, err)
	}

	// Peek should not remove element
	if q.Size() != 1 {
		t.Errorf("Size() after Peek = %d; want 1", q.Size())
	}

	q.Enqueue(99)
	val, err = q.Peek()
	if err || val != 42 {
		t.Errorf("Peek() should still be first element: got (%d, %v); want (42, false)", val, err)
	}
}

func TestQueue_ToSlice(t *testing.T) {
	q := NewQueue()
	if q == nil {
		t.Fatal("NewQueue() returned nil")
	}

	result := q.ToSlice()
	if len(result) != 0 {
		t.Errorf("ToSlice() on empty = %v; want []", result)
	}

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	result = q.ToSlice()
	expected := []int{10, 20, 30}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ToSlice() = %v; want %v", result, expected)
	}
}

// =============================================================
// 3. TEST - Queue Mixed Operations
// =============================================================

func TestQueue_MixedOperations(t *testing.T) {
	q := NewQueue()
	if q == nil {
		t.Fatal("NewQueue() returned nil")
	}

	q.Enqueue(1)
	q.Enqueue(2)
	val, _ := q.Dequeue() // 1
	if val != 1 {
		t.Errorf("Dequeue() = %d; want 1", val)
	}

	q.Enqueue(3)
	val, _ = q.Dequeue() // 2
	if val != 2 {
		t.Errorf("Dequeue() = %d; want 2", val)
	}

	val, _ = q.Dequeue() // 3
	if val != 3 {
		t.Errorf("Dequeue() = %d; want 3", val)
	}

	if !q.IsEmpty() {
		t.Error("queue should be empty")
	}
}

// =============================================================
// 4. TEST - LinkedList Basic Operations
// =============================================================

func TestLinkedList_Basic(t *testing.T) {
	ll := NewLinkedList()
	if ll == nil {
		t.Fatal("NewLinkedList() returned nil")
	}

	if ll.Size() != 0 {
		t.Errorf("new list Size() = %d; want 0", ll.Size())
	}

	result := ll.ToSlice()
	if len(result) != 0 {
		t.Errorf("ToSlice() on empty = %v; want []", result)
	}
}

func TestLinkedList_Append(t *testing.T) {
	ll := NewLinkedList()
	if ll == nil {
		t.Fatal("NewLinkedList() returned nil")
	}

	ll.Append(10)
	ll.Append(20)
	ll.Append(30)

	if ll.Size() != 3 {
		t.Errorf("Size() = %d; want 3", ll.Size())
	}

	result := ll.ToSlice()
	expected := []int{10, 20, 30}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ToSlice() = %v; want %v", result, expected)
	}
}

func TestLinkedList_Prepend(t *testing.T) {
	ll := NewLinkedList()
	if ll == nil {
		t.Fatal("NewLinkedList() returned nil")
	}

	ll.Prepend(30)
	ll.Prepend(20)
	ll.Prepend(10)

	result := ll.ToSlice()
	expected := []int{10, 20, 30}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ToSlice() = %v; want %v", result, expected)
	}
}

func TestLinkedList_AppendAndPrepend(t *testing.T) {
	ll := NewLinkedList()
	if ll == nil {
		t.Fatal("NewLinkedList() returned nil")
	}

	ll.Append(20)
	ll.Prepend(10)
	ll.Append(30)
	ll.Prepend(5)

	result := ll.ToSlice()
	expected := []int{5, 10, 20, 30}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ToSlice() = %v; want %v", result, expected)
	}

	if ll.Size() != 4 {
		t.Errorf("Size() = %d; want 4", ll.Size())
	}
}

// =============================================================
// 5. TEST - LinkedList Contains
// =============================================================

func TestLinkedList_Contains(t *testing.T) {
	ll := NewLinkedList()
	if ll == nil {
		t.Fatal("NewLinkedList() returned nil")
	}

	ll.Append(10)
	ll.Append(20)
	ll.Append(30)

	tests := []struct {
		name     string
		val      int
		expected bool
	}{
		{"ada di awal", 10, true},
		{"ada di tengah", 20, true},
		{"ada di akhir", 30, true},
		{"tidak ada", 99, false},
		{"tidak ada negatif", -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ll.Contains(tt.val)
			if result != tt.expected {
				t.Errorf("Contains(%d) = %v; want %v", tt.val, result, tt.expected)
			}
		})
	}

	// Contains on empty list
	t.Run("list kosong", func(t *testing.T) {
		empty := NewLinkedList()
		if empty == nil {
			t.Fatal("NewLinkedList() returned nil")
		}
		if empty.Contains(1) {
			t.Error("Contains(1) on empty list should be false")
		}
	})
}

// =============================================================
// 6. TEST - LinkedList DeleteByValue
// =============================================================

func TestLinkedList_DeleteByValue(t *testing.T) {
	t.Run("hapus di awal", func(t *testing.T) {
		ll := NewLinkedList()
		if ll == nil {
			t.Fatal("NewLinkedList() returned nil")
		}
		ll.Append(10)
		ll.Append(20)
		ll.Append(30)

		ok := ll.DeleteByValue(10)
		if !ok {
			t.Error("DeleteByValue(10) should return true")
		}
		expected := []int{20, 30}
		if !reflect.DeepEqual(ll.ToSlice(), expected) {
			t.Errorf("after delete 10: %v; want %v", ll.ToSlice(), expected)
		}
	})

	t.Run("hapus di tengah", func(t *testing.T) {
		ll := NewLinkedList()
		if ll == nil {
			t.Fatal("NewLinkedList() returned nil")
		}
		ll.Append(10)
		ll.Append(20)
		ll.Append(30)

		ok := ll.DeleteByValue(20)
		if !ok {
			t.Error("DeleteByValue(20) should return true")
		}
		expected := []int{10, 30}
		if !reflect.DeepEqual(ll.ToSlice(), expected) {
			t.Errorf("after delete 20: %v; want %v", ll.ToSlice(), expected)
		}
	})

	t.Run("hapus di akhir", func(t *testing.T) {
		ll := NewLinkedList()
		if ll == nil {
			t.Fatal("NewLinkedList() returned nil")
		}
		ll.Append(10)
		ll.Append(20)
		ll.Append(30)

		ok := ll.DeleteByValue(30)
		if !ok {
			t.Error("DeleteByValue(30) should return true")
		}
		expected := []int{10, 20}
		if !reflect.DeepEqual(ll.ToSlice(), expected) {
			t.Errorf("after delete 30: %v; want %v", ll.ToSlice(), expected)
		}
	})

	t.Run("value tidak ditemukan", func(t *testing.T) {
		ll := NewLinkedList()
		if ll == nil {
			t.Fatal("NewLinkedList() returned nil")
		}
		ll.Append(10)
		ll.Append(20)

		ok := ll.DeleteByValue(99)
		if ok {
			t.Error("DeleteByValue(99) should return false")
		}
		if ll.Size() != 2 {
			t.Errorf("Size() = %d; want 2", ll.Size())
		}
	})

	t.Run("hapus satu-satunya elemen", func(t *testing.T) {
		ll := NewLinkedList()
		if ll == nil {
			t.Fatal("NewLinkedList() returned nil")
		}
		ll.Append(42)

		ok := ll.DeleteByValue(42)
		if !ok {
			t.Error("DeleteByValue(42) should return true")
		}
		if ll.Size() != 0 {
			t.Errorf("Size() = %d; want 0", ll.Size())
		}
	})

	t.Run("list kosong", func(t *testing.T) {
		ll := NewLinkedList()
		if ll == nil {
			t.Fatal("NewLinkedList() returned nil")
		}
		ok := ll.DeleteByValue(1)
		if ok {
			t.Error("DeleteByValue on empty list should return false")
		}
	})

	t.Run("hapus pertama dari duplikat", func(t *testing.T) {
		ll := NewLinkedList()
		if ll == nil {
			t.Fatal("NewLinkedList() returned nil")
		}
		ll.Append(10)
		ll.Append(20)
		ll.Append(10)
		ll.Append(30)

		ok := ll.DeleteByValue(10)
		if !ok {
			t.Error("DeleteByValue(10) should return true")
		}
		expected := []int{20, 10, 30}
		if !reflect.DeepEqual(ll.ToSlice(), expected) {
			t.Errorf("after delete first 10: %v; want %v", ll.ToSlice(), expected)
		}
	})
}

// =============================================================
// 7. TEST - LinkedList Reverse
// =============================================================

func TestLinkedList_Reverse(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		ll := NewLinkedList()
		if ll == nil {
			t.Fatal("NewLinkedList() returned nil")
		}
		ll.Append(10)
		ll.Append(20)
		ll.Append(30)
		ll.Reverse()

		expected := []int{30, 20, 10}
		if !reflect.DeepEqual(ll.ToSlice(), expected) {
			t.Errorf("after Reverse: %v; want %v", ll.ToSlice(), expected)
		}
	})

	t.Run("satu elemen", func(t *testing.T) {
		ll := NewLinkedList()
		if ll == nil {
			t.Fatal("NewLinkedList() returned nil")
		}
		ll.Append(42)
		ll.Reverse()

		expected := []int{42}
		if !reflect.DeepEqual(ll.ToSlice(), expected) {
			t.Errorf("after Reverse: %v; want %v", ll.ToSlice(), expected)
		}
	})

	t.Run("list kosong", func(t *testing.T) {
		ll := NewLinkedList()
		if ll == nil {
			t.Fatal("NewLinkedList() returned nil")
		}
		ll.Reverse() // should not panic
		if ll.Size() != 0 {
			t.Errorf("Size() = %d; want 0", ll.Size())
		}
	})

	t.Run("dua elemen", func(t *testing.T) {
		ll := NewLinkedList()
		if ll == nil {
			t.Fatal("NewLinkedList() returned nil")
		}
		ll.Append(1)
		ll.Append(2)
		ll.Reverse()

		expected := []int{2, 1}
		if !reflect.DeepEqual(ll.ToSlice(), expected) {
			t.Errorf("after Reverse: %v; want %v", ll.ToSlice(), expected)
		}
	})

	t.Run("double reverse kembali ke asli", func(t *testing.T) {
		ll := NewLinkedList()
		if ll == nil {
			t.Fatal("NewLinkedList() returned nil")
		}
		ll.Append(1)
		ll.Append(2)
		ll.Append(3)
		ll.Append(4)
		ll.Append(5)

		ll.Reverse()
		ll.Reverse()

		expected := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(ll.ToSlice(), expected) {
			t.Errorf("after double Reverse: %v; want %v", ll.ToSlice(), expected)
		}
	})
}

// =============================================================
// 8. TEST - Stack Mixed Operations
// =============================================================

func TestStack_MixedOperations(t *testing.T) {
	s := NewStack()
	if s == nil {
		t.Fatal("NewStack() returned nil")
	}

	s.Push(1)
	s.Push(2)
	val, _ := s.Pop() // 2
	if val != 2 {
		t.Errorf("Pop() = %d; want 2", val)
	}

	s.Push(3)
	s.Push(4)

	// Stack sekarang: [1, 3, 4]
	if s.Size() != 3 {
		t.Errorf("Size() = %d; want 3", s.Size())
	}

	val, _ = s.Peek() // 4
	if val != 4 {
		t.Errorf("Peek() = %d; want 4", val)
	}

	val, _ = s.Pop() // 4
	val, _ = s.Pop() // 3
	val, _ = s.Pop() // 1
	if val != 1 {
		t.Errorf("last Pop() = %d; want 1", val)
	}

	if !s.IsEmpty() {
		t.Error("stack should be empty")
	}
}

// =============================================================
// 9. BENCHMARK TEST
// =============================================================

func BenchmarkStack_PushPop(b *testing.B) {
	s := NewStack()
	for i := 0; i < b.N; i++ {
		s.Push(i)
		s.Pop()
	}
}

func BenchmarkQueue_EnqueueDequeue(b *testing.B) {
	q := NewQueue()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
		q.Dequeue()
	}
}

func BenchmarkLinkedList_Append(b *testing.B) {
	ll := NewLinkedList()
	for i := 0; i < b.N; i++ {
		ll.Append(i)
	}
}

func BenchmarkLinkedList_Prepend(b *testing.B) {
	ll := NewLinkedList()
	for i := 0; i < b.N; i++ {
		ll.Prepend(i)
	}
}
