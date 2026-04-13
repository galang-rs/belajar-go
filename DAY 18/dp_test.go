package belajar

import "testing"

// =============================================================
// 1. TABLE-DRIVEN TEST - CoinChange
// =============================================================

func TestCoinChange(t *testing.T) {
	tests := []struct {
		name     string
		coins    []int
		amount   int
		expected int
	}{
		{"standar 25+5", []int{1, 5, 10, 25}, 30, 2},
		{"10+1", []int{1, 5, 10}, 11, 2},
		{"tidak mungkin", []int{2}, 3, -1},
		{"amount nol", []int{1}, 0, 0},
		{"5+5+1", []int{1, 2, 5}, 11, 3},
		{"koin kosong", []int{}, 10, -1},
		{"satu koin pas", []int{5}, 5, 1},
		{"satu koin banyak", []int{1}, 7, 7},
		{"koin besar", []int{3, 7}, 14, 2},
		{"koin besar tidak bisa", []int{3, 7}, 5, -1},
		{"kombinasi optimal", []int{1, 3, 4}, 6, 2}, // 3+3, bukan 4+1+1
		{"amount 1", []int{1, 2, 5}, 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CoinChange(tt.coins, tt.amount)
			if result != tt.expected {
				t.Errorf("CoinChange(%v, %d) = %d; want %d", tt.coins, tt.amount, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 2. TABLE-DRIVEN TEST - LongestIncreasingSubsequence
// =============================================================

func TestLongestIncreasingSubsequence(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"standar", []int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{"dengan pengulangan", []int{0, 1, 0, 3, 2, 3}, 4},
		{"semua sama", []int{7, 7, 7, 7}, 1},
		{"satu elemen", []int{5}, 1},
		{"kosong", []int{}, 0},
		{"sudah terurut", []int{1, 2, 3, 4, 5}, 5},
		{"terbalik", []int{5, 4, 3, 2, 1}, 1},
		{"zigzag", []int{1, 3, 2, 4, 3, 5}, 4},
		{"dua elemen naik", []int{1, 2}, 2},
		{"dua elemen turun", []int{2, 1}, 1},
		{"panjang", []int{3, 5, 6, 2, 5, 4, 19, 5, 6, 7, 12}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LongestIncreasingSubsequence(tt.input)
			if result != tt.expected {
				t.Errorf("LongestIncreasingSubsequence(%v) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 3. EDGE CASE TEST
// =============================================================

func TestCoinChange_EdgeCases(t *testing.T) {
	// Koin [1] selalu bisa membuat amount berapa pun
	for amount := 0; amount <= 20; amount++ {
		result := CoinChange([]int{1}, amount)
		if result != amount {
			t.Errorf("CoinChange([1], %d) = %d; want %d", amount, result, amount)
		}
	}
}

// =============================================================
// 4. BENCHMARK
// =============================================================

func BenchmarkCoinChange(b *testing.B) {
	coins := []int{1, 5, 10, 25}
	for i := 0; i < b.N; i++ {
		CoinChange(coins, 100)
	}
}

func BenchmarkLongestIncreasingSubsequence(b *testing.B) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18, 1, 4, 6, 8}
	for i := 0; i < b.N; i++ {
		LongestIncreasingSubsequence(nums)
	}
}
