package belajar

import (
	"reflect"
	"testing"
)

// =============================================================
// 1. TEST - BST Insert & InOrder
// =============================================================

func TestBST_InsertAndInOrder(t *testing.T) {
	tests := []struct {
		name     string
		inserts  []int
		expected []int
	}{
		{
			"standar",
			[]int{5, 3, 7, 1, 4},
			[]int{1, 3, 4, 5, 7},
		},
		{
			"satu elemen",
			[]int{42},
			[]int{42},
		},
		{
			"sudah terurut",
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"terbalik",
			[]int{5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"dengan duplikat",
			[]int{5, 3, 7, 3, 5, 7},
			[]int{3, 5, 7},
		},
		{
			"kosong",
			[]int{},
			[]int{},
		},
		{
			"negatif dan positif",
			[]int{0, -5, 5, -3, 3},
			[]int{-5, -3, 0, 3, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewBST()
			for _, v := range tt.inserts {
				tree.Insert(v)
			}
			result := tree.InOrder()
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Insert%v lalu InOrder() = %v; want %v", tt.inserts, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 2. TABLE-DRIVEN TEST - BST Search
// =============================================================

func TestBST_Search(t *testing.T) {
	tree := NewBST()
	values := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	for _, v := range values {
		tree.Insert(v)
	}

	tests := []struct {
		name     string
		target   int
		expected bool
	}{
		{"cari root", 8, true},
		{"cari kiri", 3, true},
		{"cari kanan", 10, true},
		{"cari leaf kiri", 1, true},
		{"cari leaf kanan", 7, true},
		{"cari 14", 14, true},
		{"tidak ada 2", 2, false},
		{"tidak ada 5", 5, false},
		{"tidak ada 100", 100, false},
		{"tidak ada negatif", -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tree.Search(tt.target)
			if result != tt.expected {
				t.Errorf("Search(%d) = %v; want %v", tt.target, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 3. TEST - BST Kosong
// =============================================================

func TestBST_Empty(t *testing.T) {
	tree := NewBST()

	if tree.Search(1) {
		t.Errorf("Search pada BST kosong harus false")
	}

	result := tree.InOrder()
	if len(result) != 0 {
		t.Errorf("InOrder pada BST kosong = %v; want []", result)
	}
}

// =============================================================
// 4. TEST - InOrder Selalu Terurut
// =============================================================

func TestBST_InOrderAlwaysSorted(t *testing.T) {
	tree := NewBST()
	inserts := []int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45}
	for _, v := range inserts {
		tree.Insert(v)
	}
	result := tree.InOrder()
	for i := 1; i < len(result); i++ {
		if result[i] <= result[i-1] {
			t.Errorf("InOrder tidak terurut di posisi %d: %v", i, result)
			break
		}
	}
}

// =============================================================
// 5. BENCHMARK
// =============================================================

func BenchmarkBST_Insert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := NewBST()
		for j := 0; j < 100; j++ {
			tree.Insert(j)
		}
	}
}

func BenchmarkBST_Search(b *testing.B) {
	tree := NewBST()
	for j := 0; j < 100; j++ {
		tree.Insert(j)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Search(50)
	}
}
