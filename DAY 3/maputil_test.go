package belajar

import (
	"reflect"
	"sort"
	"testing"
)

// =============================================================
// 1. TABLE-DRIVEN TEST - WordFrequency
// =============================================================

func TestWordFrequency(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected map[string]int
	}{
		{"slice kosong", []string{}, map[string]int{}},
		{"satu kata", []string{"go"}, map[string]int{"go": 1}},
		{"kata duplikat", []string{"go", "is", "go"}, map[string]int{"go": 2, "is": 1}},
		{"semua sama", []string{"a", "a", "a"}, map[string]int{"a": 3}},
		{"semua unik", []string{"a", "b", "c"}, map[string]int{"a": 1, "b": 1, "c": 1}},
		{"case sensitive", []string{"Go", "go", "GO"}, map[string]int{"Go": 1, "go": 1, "GO": 1}},
		{"banyak kata", []string{"x", "y", "x", "z", "y", "x"}, map[string]int{"x": 3, "y": 2, "z": 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WordFrequency(tt.input)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("WordFrequency(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 2. TABLE-DRIVEN TEST - InvertMap
// =============================================================

func TestInvertMap(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]string
		expected map[string]string
	}{
		{"map kosong", map[string]string{}, map[string]string{}},
		{"satu entry", map[string]string{"a": "1"}, map[string]string{"1": "a"}},
		{"banyak entry", map[string]string{"a": "1", "b": "2", "c": "3"}, map[string]string{"1": "a", "2": "b", "3": "c"}},
		{"value jadi key", map[string]string{"nama": "galang", "kota": "jakarta"}, map[string]string{"galang": "nama", "jakarta": "kota"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := InvertMap(tt.input)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("InvertMap(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 3. TABLE-DRIVEN TEST - MergeMaps
// =============================================================

func TestMergeMaps(t *testing.T) {
	tests := []struct {
		name     string
		m1       map[string]int
		m2       map[string]int
		expected map[string]int
	}{
		{"dua map kosong", map[string]int{}, map[string]int{}, map[string]int{}},
		{"m1 kosong", map[string]int{}, map[string]int{"a": 1}, map[string]int{"a": 1}},
		{"m2 kosong", map[string]int{"a": 1}, map[string]int{}, map[string]int{"a": 1}},
		{"tanpa overlap", map[string]int{"a": 1, "b": 2}, map[string]int{"c": 3, "d": 4}, map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}},
		{"dengan overlap m2 menang", map[string]int{"a": 1, "b": 2}, map[string]int{"b": 3, "c": 4}, map[string]int{"a": 1, "b": 3, "c": 4}},
		{"semua overlap", map[string]int{"a": 1, "b": 2}, map[string]int{"a": 10, "b": 20}, map[string]int{"a": 10, "b": 20}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeMaps(tt.m1, tt.m2)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MergeMaps(%v, %v) = %v; want %v", tt.m1, tt.m2, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 4. TEST DENGAN SORT - Keys (urutan tidak ditentukan)
// =============================================================

func TestKeys(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected []string
	}{
		{"map kosong", map[string]int{}, []string{}},
		{"satu entry", map[string]int{"a": 1}, []string{"a"}},
		{"banyak entry", map[string]int{"a": 1, "b": 2, "c": 3}, []string{"a", "b", "c"}},
		{"key panjang", map[string]int{"hello": 1, "world": 2}, []string{"hello", "world"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Keys(tt.input)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			// Sort karena urutan map tidak pasti
			sort.Strings(result)
			sort.Strings(tt.expected)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Keys(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 5. TEST DENGAN SORT - Values (urutan tidak ditentukan)
// =============================================================

func TestValues(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected []int
	}{
		{"map kosong", map[string]int{}, []int{}},
		{"satu entry", map[string]int{"a": 1}, []int{1}},
		{"banyak entry", map[string]int{"a": 1, "b": 2, "c": 3}, []int{1, 2, 3}},
		{"ada duplikat value", map[string]int{"a": 1, "b": 1, "c": 2}, []int{1, 1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Values(tt.input)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			// Sort karena urutan map tidak pasti
			sort.Ints(result)
			sort.Ints(tt.expected)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Values(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 6. TABLE-DRIVEN TEST - GroupByLength
// =============================================================

func TestGroupByLength(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected map[int][]string
	}{
		{"slice kosong", []string{}, map[int][]string{}},
		{"satu kata", []string{"go"}, map[int][]string{2: {"go"}}},
		{"panjang bervariasi", []string{"go", "is", "fun", "ab"}, map[int][]string{2: {"go", "is", "ab"}, 3: {"fun"}}},
		{"semua sama panjang", []string{"aa", "bb", "cc"}, map[int][]string{2: {"aa", "bb", "cc"}}},
		{"termasuk string kosong", []string{"", "a", "bb"}, map[int][]string{0: {""}, 1: {"a"}, 2: {"bb"}}},
		{"kata panjang", []string{"golang", "python", "c", "js"}, map[int][]string{6: {"golang", "python"}, 1: {"c"}, 2: {"js"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GroupByLength(tt.input)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("GroupByLength(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 7. TABLE-DRIVEN TEST - CountCharacters
// =============================================================

func TestCountCharacters(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[rune]int
	}{
		{"string kosong", "", map[rune]int{}},
		{"satu karakter", "a", map[rune]int{'a': 1}},
		{"hello", "hello", map[rune]int{'h': 1, 'e': 1, 'l': 2, 'o': 1}},
		{"case sensitive", "aAa", map[rune]int{'a': 2, 'A': 1}},
		{"dengan spasi", "a b", map[rune]int{'a': 1, ' ': 1, 'b': 1}},
		{"unicode", "go🚀", map[rune]int{'g': 1, 'o': 1, '🚀': 1}},
		{"semua sama", "aaaa", map[rune]int{'a': 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountCharacters(tt.input)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("CountCharacters(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 8. TABLE-DRIVEN TEST - MapContainsValue
// =============================================================

func TestMapContainsValue(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		target   int
		expected bool
	}{
		{"map kosong", map[string]int{}, 1, false},
		{"ada value", map[string]int{"a": 1, "b": 2, "c": 3}, 2, true},
		{"tidak ada value", map[string]int{"a": 1, "b": 2, "c": 3}, 4, false},
		{"value nol", map[string]int{"a": 0, "b": 1}, 0, true},
		{"value negatif", map[string]int{"a": -5, "b": 10}, -5, true},
		{"value negatif tidak ada", map[string]int{"a": 1, "b": 2}, -1, false},
		{"satu entry cocok", map[string]int{"x": 42}, 42, true},
		{"satu entry tidak cocok", map[string]int{"x": 42}, 99, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapContainsValue(tt.input, tt.target)
			if result != tt.expected {
				t.Errorf("MapContainsValue(%v, %d) = %v; want %v", tt.input, tt.target, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 9. TABLE-DRIVEN TEST - DiffMaps
// =============================================================

func TestDiffMaps(t *testing.T) {
	tests := []struct {
		name     string
		m1       map[string]int
		m2       map[string]int
		expected map[string]int
	}{
		{"dua map kosong", map[string]int{}, map[string]int{}, map[string]int{}},
		{"m2 kosong", map[string]int{"a": 1, "b": 2}, map[string]int{}, map[string]int{"a": 1, "b": 2}},
		{"m1 kosong", map[string]int{}, map[string]int{"a": 1}, map[string]int{}},
		{"semua sama", map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 2}, map[string]int{}},
		{"key sama value beda", map[string]int{"a": 1, "b": 2, "c": 3}, map[string]int{"a": 1, "c": 5}, map[string]int{"b": 2, "c": 3}},
		{"key tidak ada di m2", map[string]int{"x": 10, "y": 20}, map[string]int{"z": 30}, map[string]int{"x": 10, "y": 20}},
		{"partial overlap", map[string]int{"a": 1, "b": 2, "c": 3}, map[string]int{"a": 1, "b": 99}, map[string]int{"b": 2, "c": 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DiffMaps(tt.m1, tt.m2)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("DiffMaps(%v, %v) = %v; want %v", tt.m1, tt.m2, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 10. TABLE-DRIVEN TEST - TopNFrequent
// =============================================================

func TestTopNFrequent(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		n        int
		expected []string
	}{
		{"slice kosong", []string{}, 3, []string{}},
		{"n nol", []string{"a", "b"}, 0, []string{}},
		{"satu kata", []string{"go"}, 1, []string{"go"}},
		{"top 2", []string{"go", "is", "go", "fun", "is", "go"}, 2, []string{"go", "is"}},
		{"top 1", []string{"a", "b", "a", "c", "b", "a"}, 1, []string{"a"}},
		{"n lebih besar dari unik", []string{"a", "b", "c"}, 5, []string{"a", "b", "c"}},
		{"frekuensi sama urut kemunculan", []string{"x", "y", "z", "x", "y", "z"}, 3, []string{"x", "y", "z"}},
		{"campuran frekuensi", []string{"d", "a", "b", "a", "c", "b", "a", "d", "d", "d"}, 3, []string{"d", "a", "b"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TopNFrequent(tt.words, tt.n)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("TopNFrequent(%v, %d) = %v; want %v", tt.words, tt.n, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 11. PARALLEL TEST - WordFrequency
// =============================================================

func TestWordFrequency_Parallel(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected map[string]int
	}{
		{"case1", []string{"a", "b", "a"}, map[string]int{"a": 2, "b": 1}},
		{"case2", []string{"go", "go", "go"}, map[string]int{"go": 3}},
		{"case3", []string{"x", "y", "z"}, map[string]int{"x": 1, "y": 1, "z": 1}},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := WordFrequency(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("WordFrequency(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 12. BENCHMARK TEST - Performa map operations
// =============================================================

func BenchmarkWordFrequency(b *testing.B) {
	words := []string{"go", "is", "fun", "go", "is", "great", "go", "rocks"}
	for i := 0; i < b.N; i++ {
		WordFrequency(words)
	}
}

func BenchmarkCountCharacters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountCharacters("The quick brown fox jumps over the lazy dog")
	}
}

func BenchmarkTopNFrequent(b *testing.B) {
	words := []string{"go", "is", "fun", "go", "is", "great", "go", "rocks", "fun", "is"}
	for i := 0; i < b.N; i++ {
		TopNFrequent(words, 3)
	}
}
