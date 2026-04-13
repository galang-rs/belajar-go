package belajar

import (
	"reflect"
	"testing"
)

// =============================================================
// 1. TABLE-DRIVEN TEST - SwapValues
// =============================================================

func TestSwapValues(t *testing.T) {
	tests := []struct {
		name       string
		a, b       int
		expectA, expectB int
	}{
		{"positif", 10, 20, 20, 10},
		{"negatif", -5, 100, 100, -5},
		{"nol dan positif", 0, 42, 42, 0},
		{"sama", 7, 7, 7, 7},
		{"negatif dua-duanya", -3, -8, -8, -3},
		{"besar", 1000000, 999999, 999999, 1000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, b := tt.a, tt.b
			SwapValues(&a, &b)
			if a != tt.expectA || b != tt.expectB {
				t.Errorf("SwapValues(&%d, &%d) -> a=%d, b=%d; want a=%d, b=%d",
					tt.a, tt.b, a, b, tt.expectA, tt.expectB)
			}
		})
	}
}

// =============================================================
// 2. TABLE-DRIVEN TEST - DeepCopyMatrix
// =============================================================

func TestDeepCopyMatrix(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			"2x2 matrix",
			[][]int{{1, 2}, {3, 4}},
			[][]int{{1, 2}, {3, 4}},
		},
		{
			"3x3 matrix",
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
		{
			"1x1 matrix",
			[][]int{{42}},
			[][]int{{42}},
		},
		{
			"kosong",
			[][]int{},
			[][]int{},
		},
		{
			"nil",
			nil,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DeepCopyMatrix(tt.input)
			if tt.input == nil {
				if result != nil {
					t.Errorf("DeepCopyMatrix(nil) = %v; want nil", result)
				}
				return
			}
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("DeepCopyMatrix(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 3. DEEP COPY INDEPENDENCE TEST
// =============================================================

func TestDeepCopyMatrix_Independence(t *testing.T) {
	original := [][]int{{1, 2}, {3, 4}}
	copied := DeepCopyMatrix(original)

	// Ubah copy, pastikan original tidak berubah
	copied[0][0] = 99
	if original[0][0] != 1 {
		t.Errorf("Mengubah copy mempengaruhi original! original[0][0] = %d; want 1", original[0][0])
	}

	// Ubah original, pastikan copy tidak berubah
	original[1][1] = 88
	if copied[1][1] != 4 {
		t.Errorf("Mengubah original mempengaruhi copy! copied[1][1] = %d; want 4", copied[1][1])
	}
}

// =============================================================
// 4. SWAP DOUBLE TEST
// =============================================================

func TestSwapValues_Double(t *testing.T) {
	a, b := 42, 99
	SwapValues(&a, &b)
	SwapValues(&a, &b) // swap 2x harus kembali ke awal
	if a != 42 || b != 99 {
		t.Errorf("Swap 2x harus kembali ke awal: a=%d, b=%d; want a=42, b=99", a, b)
	}
}

// =============================================================
// 5. BENCHMARK
// =============================================================

func BenchmarkSwapValues(b *testing.B) {
	x, y := 10, 20
	for i := 0; i < b.N; i++ {
		SwapValues(&x, &y)
	}
}

func BenchmarkDeepCopyMatrix(b *testing.B) {
	m := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < b.N; i++ {
		DeepCopyMatrix(m)
	}
}
