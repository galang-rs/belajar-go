package belajar

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Empty slice", []int{}, []int{}},
		{"Single element", []int{1}, []int{1}},
		{"Even number of elements", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{"Odd number of elements", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)
			
			Reverse(inputCopy)
			if !reflect.DeepEqual(inputCopy, tt.expected) {
				t.Errorf("Reverse(%v) = %v; want %v", tt.input, inputCopy, tt.expected)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Empty slice", []int{}, []int{}},
		{"No duplicates", []int{1, 2, 3}, []int{1, 2, 3}},
		{"With duplicates", []int{1, 2, 2, 3, 1, 4}, []int{1, 2, 3, 4}},
		{"All duplicates", []int{2, 2, 2, 2}, []int{2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Unique(tt.input)
			if len(result) == 0 && len(tt.expected) == 0 {
				return // both empty
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Unique(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected []int
	}{
		{"Empty slices", []int{}, []int{}, []int{}},
		{"One empty", []int{1, 2}, []int{}, []int{}},
		{"No intersection", []int{1, 2, 3}, []int{4, 5, 6}, []int{}},
		{"Normal intersection", []int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
		{"Multiple intersection", []int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}}, 
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Intersect(tt.nums1, tt.nums2)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			set := make(map[int]bool)
			for _, v := range result {
				set[v] = true
			}
			setExpected := make(map[int]bool)
			for _, v := range tt.expected {
				setExpected[v] = true
			}
			if len(set) != len(setExpected) {
				t.Errorf("Intersect(%v, %v) = %v; want elements %v", tt.nums1, tt.nums2, result, tt.expected)
				return
			}
			for k := range setExpected {
				if !set[k] {
					t.Errorf("Intersect(%v, %v) = %v; want elements %v (missing %d)", tt.nums1, tt.nums2, result, tt.expected, k)
					return
				}
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		target   int
		expected bool
	}{
		{"Empty slice", []int{}, 1, false},
		{"Contains item", []int{1, 2, 3}, 2, true},
		{"Does not contain", []int{1, 2, 3}, 4, false},
		{"Multiple occurrences", []int{2, 2, 2}, 2, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Contains(tt.input, tt.target)
			if result != tt.expected {
				t.Errorf("Contains(%v, %d) = %v; want %v", tt.input, tt.target, result, tt.expected)
			}
		})
	}
}

func TestGroupByParity(t *testing.T) {
	tests := []struct {
		name          string
		input         []int
		expectedEven  []int
		expectedOdd   []int
	}{
		{"Empty slice", []int{}, []int{}, []int{}},
		{"All evens", []int{2, 4, 6}, []int{2, 4, 6}, []int{}},
		{"All odds", []int{1, 3, 5}, []int{}, []int{1, 3, 5}},
		{"Mixed", []int{1, 2, 3, 4, 5, 6}, []int{2, 4, 6}, []int{1, 3, 5}},
		{"Negative mixed", []int{-1, 0, 2}, []int{0, 2}, []int{-1}}, // -1 % 2 != 0
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			even, odd := GroupByParity(tt.input)
			if len(even) == 0 && len(tt.expectedEven) == 0 {
				even = []int{}
				tt.expectedEven = []int{}
			}
			if len(odd) == 0 && len(tt.expectedOdd) == 0 {
				odd = []int{}
				tt.expectedOdd = []int{}
			}
			if !reflect.DeepEqual(even, tt.expectedEven) || !reflect.DeepEqual(odd, tt.expectedOdd) {
				t.Errorf("GroupByParity(%v) = %v, %v; want %v, %v", tt.input, even, odd, tt.expectedEven, tt.expectedOdd)
			}
		})
	}
}

func TestRemoveAt(t *testing.T) {
	tests := []struct {
		name        string
		input       []int
		index       int
		expected    []int
		expectedErr bool
	}{
		{"Valid index middle", []int{10, 20, 30}, 1, []int{10, 30}, false},
		{"Valid index end", []int{10, 20, 30}, 2, []int{10, 20}, false},
		{"Valid index start", []int{10, 20, 30}, 0, []int{20, 30}, false},
		{"Index out of bounds positive", []int{10, 20, 30}, 3, nil, true},
		{"Index out of bounds negative", []int{10, 20, 30}, -1, nil, true},
		{"Empty slice", []int{}, 0, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := RemoveAt(tt.input, tt.index)
			if err != tt.expectedErr {
				t.Errorf("RemoveAt(%v, %d) err = %v; want err = %v", tt.input, tt.index, err, tt.expectedErr)
			}
			if !err {
				if len(result) == 0 && len(tt.expected) == 0 {
					return
				}
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("RemoveAt(%v, %d) = %v; want %v", tt.input, tt.index, result, tt.expected)
				}
			}
		})
	}
}
