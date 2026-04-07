package belajar

import (
	"testing"
)

// =============================================================
// 1. BASIC TEST - Test sederhana satu per satu
// =============================================================

func TestAddBasic(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Add(2, 3) = %d; expected 5", result)
	}
}

func TestSubtractBasic(t *testing.T) {
	result := Subtract(10, 3)
	if result != 7 {
		t.Errorf("Subtract(10, 3) = %d; expected 7", result)
	}
}

func TestMultiplyBasic(t *testing.T) {
	result := Multiply(4, 5)
	if result != 20 {
		t.Errorf("Multiply(4, 5) = %d; expected 20", result)
	}
}

// =============================================================
// 2. TABLE-DRIVEN TEST - Test banyak kasus sekaligus pakai slice
// =============================================================

func TestAdd_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positif + positif", 2, 3, 5},
		{"positif + negatif", 10, -3, 7},
		{"negatif + negatif", -4, -6, -10},
		{"dengan nol", 5, 0, 5},
		{"nol + nol", 0, 0, 0},
		{"bilangan besar", 1000000, 2000000, 3000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestSubtract_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positif - positif", 10, 3, 7},
		{"hasil negatif", 3, 10, -7},
		{"dengan nol", 5, 0, 5},
		{"negatif - negatif", -4, -6, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Subtract(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Subtract(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMultiply_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positif * positif", 3, 4, 12},
		{"dengan nol", 5, 0, 0},
		{"dengan satu", 7, 1, 7},
		{"negatif * positif", -3, 4, -12},
		{"negatif * negatif", -3, -4, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 3. TEST DENGAN ERROR HANDLING - Test function yang return error
// =============================================================

func TestDivide(t *testing.T) {
	tests := []struct {
		name      string
		a, b      float64
		expected  float64
		expectErr bool
	}{
		{"normal", 10, 2, 5.0, false},
		{"desimal", 7, 2, 3.5, false},
		{"bagi nol", 10, 0, 0, true},
		{"nol dibagi", 0, 5, 0, false},
		{"negatif", -10, 2, -5.0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)

			if tt.expectErr {
				if !err {
					t.Errorf("Divide(%v, %v) expected error, got none", tt.a, tt.b)
				}
				return
			}

			if err {
				t.Errorf("Divide(%v, %v) unexpected error: %v", tt.a, tt.b, err)
				return
			}

			if result != tt.expected {
				t.Errorf("Divide(%v, %v) = %v; expected %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		name      string
		input     int
		expected  int
		expectErr bool
	}{
		{"nol", 0, 1, false},
		{"satu", 1, 1, false},
		{"lima", 5, 120, false},
		{"sepuluh", 10, 3628800, false},
		{"negatif", -1, 0, true},
		{"negatif besar", -10, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Factorial(tt.input)

			if tt.expectErr {
				if !err {
					t.Errorf("Factorial(%d) expected error, got none", tt.input)
				}
				return
			}

			if err {
				t.Errorf("Factorial(%d) unexpected error: %v", tt.input, err)
				return
			}

			if result != tt.expected {
				t.Errorf("Factorial(%d) = %d; expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 4. SUBTEST DENGAN t.Run - Test terstruktur dalam grup
// =============================================================

func TestIsPrime(t *testing.T) {
	t.Run("bilangan prima", func(t *testing.T) {
		primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 97}
		for _, p := range primes {
			if !IsPrime(p) {
				t.Errorf("IsPrime(%d) = false; expected true", p)
			}
		}
	})

	t.Run("bukan prima", func(t *testing.T) {
		notPrimes := []int{0, 1, 4, 6, 8, 9, 10, 15, 100}
		for _, n := range notPrimes {
			if IsPrime(n) {
				t.Errorf("IsPrime(%d) = true; expected false", n)
			}
		}
	})

	t.Run("bilangan negatif", func(t *testing.T) {
		negatives := []int{-1, -2, -10, -97}
		for _, n := range negatives {
			if IsPrime(n) {
				t.Errorf("IsPrime(%d) = true; expected false", n)
			}
		}
	})
}

// =============================================================
// 5. TEST SLICE OPERATIONS - Test dengan input slice
// =============================================================

func TestMax(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		expected  int
		expectErr bool
	}{
		{"normal", []int{1, 5, 3, 9, 2}, 9, false},
		{"satu elemen", []int{42}, 42, false},
		{"semua sama", []int{7, 7, 7}, 7, false},
		{"ada negatif", []int{-3, -1, -7}, -1, false},
		{"campuran", []int{-5, 0, 5, 10, -10}, 10, false},
		{"slice kosong", []int{}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Max(tt.input)

			if tt.expectErr {
				if err == false {
					t.Error("Max() expected error for empty slice, got nil")
				}
				return
			}

			if err != false {
				t.Errorf("Max() unexpected error: %v", err)
				return
			}

			if result != tt.expected {
				t.Errorf("Max(%v) = %d; expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSumSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"normal", []int{1, 2, 3, 4, 5}, 15},
		{"satu elemen", []int{10}, 10},
		{"kosong", []int{}, 0},
		{"ada negatif", []int{-1, 2, -3, 4}, 2},
		{"semua negatif", []int{-1, -2, -3}, -6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SumSlice(tt.input)
			if result != tt.expected {
				t.Errorf("SumSlice(%v) = %d; expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 6. TEST FIBONACCI - Test rekursi / iterasi
// =============================================================

func TestFibonacciN(t *testing.T) {
	tests := []struct {
		name      string
		input     int
		expected  int
		expectErr bool
	}{
		{"F(0)", 0, 0, false},
		{"F(1)", 1, 1, false},
		{"F(2)", 2, 1, false},
		{"F(3)", 3, 2, false},
		{"F(6)", 6, 8, false},
		{"F(10)", 10, 55, false},
		{"F(20)", 20, 6765, false},
		{"negatif", -1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := FibonacciN(tt.input)

			if tt.expectErr {
				if err == false {
					t.Errorf("FibonacciN(%d) expected error, got nil", tt.input)
				}
				return
			}

			if err != false {
				t.Errorf("FibonacciN(%d) unexpected error: %v", tt.input, err)
				return
			}

			if result != tt.expected {
				t.Errorf("FibonacciN(%d) = %d; expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 7. TEST ABS - Test sederhana dengan edge case
// =============================================================

func TestAbs(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"positif", 5, 5},
		{"negatif", -5, 5},
		{"nol", 0, 0},
		{"negatif satu", -1, 1},
		{"besar", -999999, 999999},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Abs(tt.input)
			if result != tt.expected {
				t.Errorf("Abs(%d) = %d; expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 8. PARALLEL TEST - Test yang jalan secara paralel
// =============================================================

func TestAdd_Parallel(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"case1", 1, 1, 2},
		{"case2", 100, 200, 300},
		{"case3", -50, 50, 0},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // test ini jalan paralel
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 9. BENCHMARK TEST - Test performa
// =============================================================

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(20)
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(97)
	}
}

func BenchmarkFibonacciN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciN(20)
	}
}
