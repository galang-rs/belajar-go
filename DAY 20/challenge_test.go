package belajar

import (
	"reflect"
	"testing"
)

// =============================================================
// 1. TABLE-DRIVEN TEST - SpiralMatrix
// =============================================================

func TestSpiralMatrix(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected [][]int
	}{
		{
			"3x3",
			3,
			[][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			"4x4",
			4,
			[][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
		},
		{
			"1x1",
			1,
			[][]int{{1}},
		},
		{
			"0",
			0,
			[][]int{},
		},
		{
			"2x2",
			2,
			[][]int{
				{1, 2},
				{4, 3},
			},
		},
		{
			"5x5",
			5,
			[][]int{
				{1, 2, 3, 4, 5},
				{16, 17, 18, 19, 6},
				{15, 24, 25, 20, 7},
				{14, 23, 22, 21, 8},
				{13, 12, 11, 10, 9},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SpiralMatrix(tt.n)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SpiralMatrix(%d) = %v; want %v", tt.n, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 2. TEST - SpiralMatrix berisi semua angka 1..N²
// =============================================================

func TestSpiralMatrix_AllNumbers(t *testing.T) {
	n := 4
	result := SpiralMatrix(n)
	seen := make(map[int]bool)
	for _, row := range result {
		for _, val := range row {
			seen[val] = true
		}
	}
	for i := 1; i <= n*n; i++ {
		if !seen[i] {
			t.Errorf("SpiralMatrix(%d) tidak mengandung angka %d", n, i)
		}
	}
}

// =============================================================
// 3. TABLE-DRIVEN TEST - ValidSudoku
// =============================================================

func TestValidSudoku(t *testing.T) {
	validBoard := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	dupBaris := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 5, 0}, // baris 0: angka 5 muncul 2x
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	dupKolom := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{5, 0, 0, 0, 2, 0, 0, 0, 6}, // kolom 0: angka 5 muncul di baris 0 dan 5
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	dupBox := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 5, 0, 1, 9, 5, 0, 0, 0}, // box kiri-atas: angka 5 di (0,0) dan (1,1)
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	boardKosong := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	tests := []struct {
		name     string
		board    [][]int
		expected bool
	}{
		{"board valid", validBoard, true},
		{"duplikat baris", dupBaris, false},
		{"duplikat kolom", dupKolom, false},
		{"duplikat box 3x3", dupBox, false},
		{"board kosong (semua 0)", boardKosong, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidSudoku(tt.board)
			if result != tt.expected {
				t.Errorf("ValidSudoku() = %v; want %v", result, tt.expected)
			}
		})
	}
}

// =============================================================
// 4. BENCHMARK
// =============================================================

func BenchmarkSpiralMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SpiralMatrix(10)
	}
}

func BenchmarkValidSudoku(b *testing.B) {
	board := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	for i := 0; i < b.N; i++ {
		ValidSudoku(board)
	}
}
