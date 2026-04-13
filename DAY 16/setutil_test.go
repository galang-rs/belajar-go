package belajar

import (
	"reflect"
	"testing"
)

// =============================================================
// 1. TABLE-DRIVEN TEST - Union
// =============================================================

func TestUnion(t *testing.T) {
	tests := []struct {
		name     string
		a, b     []int
		expected []int
	}{
		{
			"standar",
			[]int{1, 2, 3}, []int{3, 4, 5},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"dengan duplikat",
			[]int{1, 1, 2}, []int{2, 3, 3},
			[]int{1, 2, 3},
		},
		{
			"a kosong",
			[]int{}, []int{1, 2},
			[]int{1, 2},
		},
		{
			"b kosong",
			[]int{5, 6}, []int{},
			[]int{5, 6},
		},
		{
			"dua-duanya kosong",
			[]int{}, []int{},
			[]int{},
		},
		{
			"sama persis",
			[]int{5}, []int{5},
			[]int{5},
		},
		{
			"tidak ada irisan",
			[]int{1, 2}, []int{3, 4},
			[]int{1, 2, 3, 4},
		},
		{
			"banyak duplikat",
			[]int{1, 1, 1, 2, 2}, []int{2, 2, 3, 3, 3},
			[]int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Union(tt.a, tt.b)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Union(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 2. TABLE-DRIVEN TEST - SymmetricDifference
// =============================================================

func TestSymmetricDifference(t *testing.T) {
	tests := []struct {
		name     string
		a, b     []int
		expected []int
	}{
		{
			"standar",
			[]int{1, 2, 3}, []int{3, 4, 5},
			[]int{1, 2, 4, 5},
		},
		{
			"sama persis",
			[]int{1, 2}, []int{1, 2},
			[]int{},
		},
		{
			"tidak ada irisan",
			[]int{1, 2, 3}, []int{4, 5, 6},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"a kosong",
			[]int{}, []int{1, 2},
			[]int{1, 2},
		},
		{
			"b kosong",
			[]int{3, 4}, []int{},
			[]int{3, 4},
		},
		{
			"dua-duanya kosong",
			[]int{}, []int{},
			[]int{},
		},
		{
			"dengan duplikat",
			[]int{1, 1, 2}, []int{2, 3, 3},
			[]int{1, 3},
		},
		{
			"satu elemen sama",
			[]int{5}, []int{5},
			[]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SymmetricDifference(tt.a, tt.b)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SymmetricDifference(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 3. INTEGRATION TEST - SymDiff(a,a) = kosong
// =============================================================

func TestSymmetricDifference_SameSlice(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	result := SymmetricDifference(a, a)
	if len(result) != 0 {
		t.Errorf("SymmetricDifference(a, a) harus kosong, got %v", result)
	}
}

// =============================================================
// 4. BENCHMARK
// =============================================================

func BenchmarkUnion(b *testing.B) {
	a := []int{1, 2, 3, 4, 5}
	c := []int{4, 5, 6, 7, 8}
	for i := 0; i < b.N; i++ {
		Union(a, c)
	}
}

func BenchmarkSymmetricDifference(b *testing.B) {
	a := []int{1, 2, 3, 4, 5}
	c := []int{4, 5, 6, 7, 8}
	for i := 0; i < b.N; i++ {
		SymmetricDifference(a, c)
	}
}
