package belajar

import "testing"

// =============================================================
// 1. TEST - MakeCounter
// =============================================================

func TestMakeCounter(t *testing.T) {
	counter := MakeCounter()

	tests := []struct {
		name     string
		expected int
	}{
		{"panggilan pertama", 1},
		{"panggilan kedua", 2},
		{"panggilan ketiga", 3},
		{"panggilan keempat", 4},
		{"panggilan kelima", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := counter()
			if result != tt.expected {
				t.Errorf("counter() = %d; want %d", result, tt.expected)
			}
		})
	}
}

func TestMakeCounter_Independent(t *testing.T) {
	counter1 := MakeCounter()
	counter2 := MakeCounter()

	// counter1 dipanggil 3x
	counter1()
	counter1()
	val1 := counter1()

	// counter2 dipanggil 1x (harus independen)
	val2 := counter2()

	if val1 != 3 {
		t.Errorf("counter1 panggilan ke-3 = %d; want 3", val1)
	}
	if val2 != 1 {
		t.Errorf("counter2 panggilan pertama = %d; want 1 (harus independen)", val2)
	}
}

// =============================================================
// 2. TABLE-DRIVEN TEST - MakeMultiplier
// =============================================================

func TestMakeMultiplier(t *testing.T) {
	tests := []struct {
		name     string
		factor   int
		input    int
		expected int
	}{
		{"double 5", 2, 5, 10},
		{"double 3", 2, 3, 6},
		{"double 0", 2, 0, 0},
		{"triple 4", 3, 4, 12},
		{"triple 10", 3, 10, 30},
		{"kali 10", 10, 7, 70},
		{"kali 0", 0, 100, 0},
		{"kali 1", 1, 42, 42},
		{"negatif", -2, 5, -10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := MakeMultiplier(tt.factor)
			result := fn(tt.input)
			if result != tt.expected {
				t.Errorf("MakeMultiplier(%d)(%d) = %d; want %d", tt.factor, tt.input, result, tt.expected)
			}
		})
	}
}

func TestMakeMultiplier_Independent(t *testing.T) {
	double := MakeMultiplier(2)
	triple := MakeMultiplier(3)

	// Pastikan tidak saling pengaruh
	r1 := double(10)
	r2 := triple(10)

	if r1 != 20 {
		t.Errorf("double(10) = %d; want 20", r1)
	}
	if r2 != 30 {
		t.Errorf("triple(10) = %d; want 30", r2)
	}
}

// =============================================================
// 3. BENCHMARK
// =============================================================

func BenchmarkMakeCounter(b *testing.B) {
	counter := MakeCounter()
	for i := 0; i < b.N; i++ {
		counter()
	}
}

func BenchmarkMakeMultiplier(b *testing.B) {
	double := MakeMultiplier(2)
	for i := 0; i < b.N; i++ {
		double(42)
	}
}
