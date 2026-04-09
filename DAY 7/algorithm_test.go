package belajar

import (
	"reflect"
	"testing"
)

// =============================================================
// 1. TABLE-DRIVEN TEST - BinarySearch
// =============================================================

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		target    int
		expectIdx int
		expectOk  bool
	}{
		{"ditemukan di tengah", []int{1, 3, 5, 7, 9}, 5, 2, true},
		{"ditemukan di awal", []int{1, 3, 5, 7, 9}, 1, 0, true},
		{"ditemukan di akhir", []int{1, 3, 5, 7, 9}, 9, 4, true},
		{"tidak ditemukan", []int{1, 3, 5, 7, 9}, 4, -1, false},
		{"slice kosong", []int{}, 1, -1, false},
		{"satu elemen ditemukan", []int{5}, 5, 0, true},
		{"satu elemen tidak ditemukan", []int{5}, 3, -1, false},
		{"target di luar range kiri", []int{10, 20, 30}, 5, -1, false},
		{"target di luar range kanan", []int{10, 20, 30}, 35, -1, false},
		{"slice besar", []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, 14, 6, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx, ok := BinarySearch(tt.nums, tt.target)
			if idx != tt.expectIdx || ok != tt.expectOk {
				t.Errorf("BinarySearch(%v, %d) = (%d, %v); want (%d, %v)", tt.nums, tt.target, idx, ok, tt.expectIdx, tt.expectOk)
			}
		})
	}
}

// =============================================================
// 2. TABLE-DRIVEN TEST - BubbleSort
// =============================================================

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"normal", []int{5, 3, 1, 4, 2}, []int{1, 2, 3, 4, 5}},
		{"sudah terurut", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"terbalik", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"satu elemen", []int{42}, []int{42}},
		{"slice kosong", []int{}, []int{}},
		{"ada duplikat", []int{3, 1, 4, 1, 5, 9, 2, 6}, []int{1, 1, 2, 3, 4, 5, 6, 9}},
		{"semua sama", []int{7, 7, 7}, []int{7, 7, 7}},
		{"negatif", []int{-3, -1, -4, -1, -5}, []int{-5, -4, -3, -1, -1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BubbleSort(tt.input)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BubbleSort(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}

	// Test: tidak mengubah slice asli
	t.Run("tidak mengubah slice asli", func(t *testing.T) {
		original := []int{5, 3, 1}
		originalCopy := []int{5, 3, 1}
		BubbleSort(original)
		if !reflect.DeepEqual(original, originalCopy) {
			t.Errorf("BubbleSort() mengubah slice asli")
		}
	})
}

// =============================================================
// 3. TABLE-DRIVEN TEST - SelectionSort
// =============================================================

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"normal", []int{64, 25, 12, 22, 11}, []int{11, 12, 22, 25, 64}},
		{"sudah terurut", []int{1, 2, 3}, []int{1, 2, 3}},
		{"terbalik", []int{3, 2, 1}, []int{1, 2, 3}},
		{"satu elemen", []int{1}, []int{1}},
		{"slice kosong", []int{}, []int{}},
		{"ada duplikat", []int{5, 3, 5, 1}, []int{1, 3, 5, 5}},
		{"negatif", []int{0, -5, 3, -2, 8}, []int{-5, -2, 0, 3, 8}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SelectionSort(tt.input)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SelectionSort(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 4. TABLE-DRIVEN TEST - InsertionSort
// =============================================================

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"normal", []int{5, 2, 4, 6, 1, 3}, []int{1, 2, 3, 4, 5, 6}},
		{"sudah terurut", []int{1, 2, 3}, []int{1, 2, 3}},
		{"terbalik", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"satu elemen", []int{42}, []int{42}},
		{"slice kosong", []int{}, []int{}},
		{"ada duplikat", []int{4, 2, 4, 1}, []int{1, 2, 4, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := InsertionSort(tt.input)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("InsertionSort(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 5. TABLE-DRIVEN TEST - MergeSort
// =============================================================

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"normal", []int{38, 27, 43, 3, 9, 82, 10}, []int{3, 9, 10, 27, 38, 43, 82}},
		{"sudah terurut", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"terbalik", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"satu elemen", []int{1}, []int{1}},
		{"slice kosong", []int{}, []int{}},
		{"dua elemen", []int{2, 1}, []int{1, 2}},
		{"ada duplikat", []int{5, 1, 5, 3, 1}, []int{1, 1, 3, 5, 5}},
		{"negatif dan positif", []int{-5, 3, -1, 0, 7}, []int{-5, -1, 0, 3, 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeSort(tt.input)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MergeSort(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 6. TABLE-DRIVEN TEST - GCD
// =============================================================

func TestGCD(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"12 dan 8", 12, 8, 4},
		{"17 dan 5", 17, 5, 1},
		{"100 dan 75", 100, 75, 25},
		{"0 dan 5", 0, 5, 5},
		{"5 dan 0", 5, 0, 5},
		{"sama", 7, 7, 7},
		{"satu dan banyak", 1, 100, 1},
		{"bilangan besar", 48, 18, 6},
		{"prima", 13, 17, 1},
		{"kelipatan", 15, 45, 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GCD(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("GCD(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 7. TABLE-DRIVEN TEST - LCM
// =============================================================

func TestLCM(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"4 dan 6", 4, 6, 12},
		{"3 dan 5", 3, 5, 15},
		{"0 dan 5", 0, 5, 0},
		{"5 dan 0", 5, 0, 0},
		{"sama", 7, 7, 7},
		{"1 dan angka", 1, 10, 10},
		{"12 dan 18", 12, 18, 36},
		{"prima", 3, 7, 21},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LCM(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("LCM(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 8. TABLE-DRIVEN TEST - Power
// =============================================================

func TestPower(t *testing.T) {
	tests := []struct {
		name     string
		base     int
		exp      int
		expected int
	}{
		{"2^10", 2, 10, 1024},
		{"3^3", 3, 3, 27},
		{"5^0", 5, 0, 1},
		{"0^5", 0, 5, 0},
		{"1^100", 1, 100, 1},
		{"2^0", 2, 0, 1},
		{"10^3", 10, 3, 1000},
		{"2^1", 2, 1, 2},
		{"7^2", 7, 2, 49},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Power(tt.base, tt.exp)
			if result != tt.expected {
				t.Errorf("Power(%d, %d) = %d; want %d", tt.base, tt.exp, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 9. TABLE-DRIVEN TEST - SumDigits
// =============================================================

func TestSumDigits(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"123", 123, 6},
		{"negartif -456", -456, 15},
		{"nol", 0, 0},
		{"satu digit", 9, 9},
		{"100", 100, 1},
		{"999", 999, 27},
		{"12345", 12345, 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SumDigits(tt.input)
			if result != tt.expected {
				t.Errorf("SumDigits(%d) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 10. TABLE-DRIVEN TEST - IsSorted
// =============================================================

func TestIsSorted(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{"terurut", []int{1, 2, 3, 4, 5}, true},
		{"tidak terurut", []int{1, 3, 2}, false},
		{"terurut dengan duplikat", []int{1, 1, 2, 2, 3}, true},
		{"slice kosong", []int{}, true},
		{"satu elemen", []int{42}, true},
		{"terbalik", []int{5, 4, 3, 2, 1}, false},
		{"semua sama", []int{3, 3, 3}, true},
		{"dua elemen terurut", []int{1, 2}, true},
		{"dua elemen tidak terurut", []int{2, 1}, false},
		{"hampir terurut", []int{1, 2, 3, 5, 4}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsSorted(tt.input)
			if result != tt.expected {
				t.Errorf("IsSorted(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 11. PARALLEL TEST - Sorting Algorithms
// =============================================================

func TestSortingAlgorithms_Parallel(t *testing.T) {
	input := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	t.Run("BubbleSort", func(t *testing.T) {
		t.Parallel()
		result := BubbleSort(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("BubbleSort() = %v; want %v", result, expected)
		}
	})

	t.Run("SelectionSort", func(t *testing.T) {
		t.Parallel()
		result := SelectionSort(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SelectionSort() = %v; want %v", result, expected)
		}
	})

	t.Run("InsertionSort", func(t *testing.T) {
		t.Parallel()
		result := InsertionSort(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("InsertionSort() = %v; want %v", result, expected)
		}
	})

	t.Run("MergeSort", func(t *testing.T) {
		t.Parallel()
		result := MergeSort(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("MergeSort() = %v; want %v", result, expected)
		}
	})
}

// =============================================================
// 12. BENCHMARK TEST
// =============================================================

func BenchmarkBubbleSort(b *testing.B) {
	input := []int{38, 27, 43, 3, 9, 82, 10, 45, 12, 67}
	for i := 0; i < b.N; i++ {
		BubbleSort(input)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	input := []int{38, 27, 43, 3, 9, 82, 10, 45, 12, 67}
	for i := 0; i < b.N; i++ {
		MergeSort(input)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	input := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	for i := 0; i < b.N; i++ {
		BinarySearch(input, 13)
	}
}

func BenchmarkGCD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GCD(48, 18)
	}
}

func BenchmarkPower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Power(2, 20)
	}
}
