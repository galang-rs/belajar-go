package belajar

import (
	"reflect"
	"sort"
	"testing"
)

// =============================================================
// 1. TABLE-DRIVEN TEST - Permutations
// =============================================================

func TestPermutations(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{
			"tiga elemen",
			[]int{1, 2, 3},
			[][]int{
				{1, 2, 3}, {1, 3, 2},
				{2, 1, 3}, {2, 3, 1},
				{3, 1, 2}, {3, 2, 1},
			},
		},
		{
			"satu elemen",
			[]int{1},
			[][]int{{1}},
		},
		{
			"kosong",
			[]int{},
			[][]int{},
		},
		{
			"dua elemen",
			[]int{4, 5},
			[][]int{{4, 5}, {5, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Permutations(tt.input)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			// Sort both for comparison
			sortPerms := func(perms [][]int) {
				sort.Slice(perms, func(i, j int) bool {
					for k := 0; k < len(perms[i]) && k < len(perms[j]); k++ {
						if perms[i][k] != perms[j][k] {
							return perms[i][k] < perms[j][k]
						}
					}
					return len(perms[i]) < len(perms[j])
				})
			}
			sortPerms(result)
			sortPerms(tt.expected)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Permutations(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 2. TABLE-DRIVEN TEST - PowerSet
// =============================================================

func TestPowerSet(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{
			"tiga elemen",
			[]int{1, 2, 3},
			[][]int{
				{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3},
			},
		},
		{
			"satu elemen",
			[]int{1},
			[][]int{{}, {1}},
		},
		{
			"kosong",
			[]int{},
			[][]int{{}},
		},
		{
			"dua elemen",
			[]int{4, 5},
			[][]int{{}, {4}, {4, 5}, {5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PowerSet(tt.input)
			// Sort both for comparison
			sortSets := func(sets [][]int) {
				for i := range sets {
					sort.Ints(sets[i])
				}
				sort.Slice(sets, func(i, j int) bool {
					if len(sets[i]) != len(sets[j]) {
						return len(sets[i]) < len(sets[j])
					}
					for k := 0; k < len(sets[i]); k++ {
						if sets[i][k] != sets[j][k] {
							return sets[i][k] < sets[j][k]
						}
					}
					return false
				})
			}
			sortSets(result)
			sortSets(tt.expected)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("PowerSet(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 3. JUMLAH PERMUTASI
// =============================================================

func TestPermutations_Count(t *testing.T) {
	// 4! = 24 permutasi
	result := Permutations([]int{1, 2, 3, 4})
	if len(result) != 24 {
		t.Errorf("Permutations([1,2,3,4]) menghasilkan %d permutasi; want 24", len(result))
	}
}

// =============================================================
// 4. JUMLAH POWER SET
// =============================================================

func TestPowerSet_Count(t *testing.T) {
	// 2^4 = 16 subset
	result := PowerSet([]int{1, 2, 3, 4})
	if len(result) != 16 {
		t.Errorf("PowerSet([1,2,3,4]) menghasilkan %d subset; want 16", len(result))
	}
}

// =============================================================
// 5. BENCHMARK
// =============================================================

func BenchmarkPermutations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Permutations([]int{1, 2, 3, 4, 5})
	}
}

func BenchmarkPowerSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PowerSet([]int{1, 2, 3, 4, 5})
	}
}
