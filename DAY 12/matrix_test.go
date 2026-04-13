package belajar

import (
	"reflect"
	"testing"
)

// =============================================================
// 1. TABLE-DRIVEN TEST - TransposeMatrix
// =============================================================

func TestTransposeMatrix(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			"2x3 matrix",
			[][]int{{1, 2, 3}, {4, 5, 6}},
			[][]int{{1, 4}, {2, 5}, {3, 6}},
		},
		{
			"3x3 matrix",
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}},
		},
		{
			"1x1 matrix",
			[][]int{{1}},
			[][]int{{1}},
		},
		{
			"kosong",
			[][]int{},
			[][]int{},
		},
		{
			"1x3 matrix",
			[][]int{{1, 2, 3}},
			[][]int{{1}, {2}, {3}},
		},
		{
			"3x1 matrix",
			[][]int{{1}, {2}, {3}},
			[][]int{{1, 2, 3}},
		},
		{
			"2x2 matrix",
			[][]int{{10, 20}, {30, 40}},
			[][]int{{10, 30}, {20, 40}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TransposeMatrix(tt.input)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("TransposeMatrix(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 2. TABLE-DRIVEN TEST - RotateMatrix90
// =============================================================

func TestRotateMatrix90(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			"3x3 matrix",
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			"2x2 matrix",
			[][]int{{1, 2}, {3, 4}},
			[][]int{{3, 1}, {4, 2}},
		},
		{
			"1x1 matrix",
			[][]int{{5}},
			[][]int{{5}},
		},
		{
			"kosong",
			[][]int{},
			[][]int{},
		},
		{
			"4x4 matrix",
			[][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			[][]int{
				{13, 9, 5, 1},
				{14, 10, 6, 2},
				{15, 11, 7, 3},
				{16, 12, 8, 4},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RotateMatrix90(tt.input)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("RotateMatrix90(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 3. INTEGRATION TEST - Rotate 4x = Identity
// =============================================================

func TestRotateMatrix90_FourTimes(t *testing.T) {
	original := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	result := RotateMatrix90(original)
	result = RotateMatrix90(result)
	result = RotateMatrix90(result)
	result = RotateMatrix90(result)
	if !reflect.DeepEqual(result, original) {
		t.Errorf("Rotasi 4x seharusnya kembali ke matrix asli, got %v", result)
	}
}

// =============================================================
// 4. INTEGRATION TEST - Transpose 2x = Identity
// =============================================================

func TestTransposeMatrix_Twice(t *testing.T) {
	original := [][]int{{1, 2, 3}, {4, 5, 6}}
	result := TransposeMatrix(TransposeMatrix(original))
	if !reflect.DeepEqual(result, original) {
		t.Errorf("Transpose 2x seharusnya kembali ke matrix asli, got %v", result)
	}
}

// =============================================================
// 5. BENCHMARK
// =============================================================

func BenchmarkTransposeMatrix(b *testing.B) {
	m := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < b.N; i++ {
		TransposeMatrix(m)
	}
}

func BenchmarkRotateMatrix90(b *testing.B) {
	m := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < b.N; i++ {
		RotateMatrix90(m)
	}
}
