package belajar

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

// Helper: perbandingan float64 dengan toleransi
func approxEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) < tolerance
}

const epsilon = 1e-9

// =============================================================
// 1. TABLE-DRIVEN TEST - SafeDivide
// =============================================================

func TestSafeDivide(t *testing.T) {
	tests := []struct {
		name      string
		a, b      float64
		expected  float64
		expectErr error
	}{
		{"normal", 10, 2, 5, nil},
		{"desimal", 7, 2, 3.5, nil},
		{"bagi nol", 10, 0, 0, ErrDivisionByZero},
		{"nol dibagi", 0, 5, 0, nil},
		{"negatif", -10, 2, -5, nil},
		{"dua negatif", -10, -2, 5, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SafeDivide(tt.a, tt.b)

			if tt.expectErr != nil {
				if err == nil {
					t.Errorf("SafeDivide(%v, %v) expected error, got nil", tt.a, tt.b)
					return
				}
				if !errors.Is(err, tt.expectErr) {
					t.Errorf("SafeDivide(%v, %v) error = %v; want %v", tt.a, tt.b, err, tt.expectErr)
				}
				return
			}

			if err != nil {
				t.Errorf("SafeDivide(%v, %v) unexpected error: %v", tt.a, tt.b, err)
				return
			}

			if !approxEqual(result, tt.expected, epsilon) {
				t.Errorf("SafeDivide(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 2. TABLE-DRIVEN TEST - SafeSqrt
// =============================================================

func TestSafeSqrt(t *testing.T) {
	tests := []struct {
		name      string
		input     float64
		expected  float64
		expectErr error
	}{
		{"positif 25", 25, 5, nil},
		{"positif 16", 16, 4, nil},
		{"nol", 0, 0, nil},
		{"positif 2", 2, math.Sqrt(2), nil},
		{"negatif", -4, 0, ErrNegativeNumber},
		{"negatif kecil", -0.1, 0, ErrNegativeNumber},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SafeSqrt(tt.input)

			if tt.expectErr != nil {
				if err == nil {
					t.Errorf("SafeSqrt(%v) expected error, got nil", tt.input)
					return
				}
				if !errors.Is(err, tt.expectErr) {
					t.Errorf("SafeSqrt(%v) error = %v; want %v", tt.input, err, tt.expectErr)
				}
				return
			}

			if err != nil {
				t.Errorf("SafeSqrt(%v) unexpected error: %v", tt.input, err)
				return
			}

			if !approxEqual(result, tt.expected, epsilon) {
				t.Errorf("SafeSqrt(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 3. TABLE-DRIVEN TEST - SafeIndex
// =============================================================

func TestSafeIndex(t *testing.T) {
	nums := []int{10, 20, 30, 40, 50}

	tests := []struct {
		name      string
		nums      []int
		index     int
		expected  int
		expectErr error
	}{
		{"index valid 0", nums, 0, 10, nil},
		{"index valid 2", nums, 2, 30, nil},
		{"index valid terakhir", nums, 4, 50, nil},
		{"index negatif", nums, -1, 0, ErrOutOfRange},
		{"index terlalu besar", nums, 5, 0, ErrOutOfRange},
		{"index sangat besar", nums, 100, 0, ErrOutOfRange},
		{"slice kosong", []int{}, 0, 0, ErrOutOfRange},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SafeIndex(tt.nums, tt.index)

			if tt.expectErr != nil {
				if err == nil {
					t.Errorf("SafeIndex(..., %d) expected error, got nil", tt.index)
					return
				}
				if !errors.Is(err, tt.expectErr) {
					t.Errorf("SafeIndex(..., %d) error = %v; want %v", tt.index, err, tt.expectErr)
				}
				return
			}

			if err != nil {
				t.Errorf("SafeIndex(..., %d) unexpected error: %v", tt.index, err)
				return
			}

			if result != tt.expected {
				t.Errorf("SafeIndex(..., %d) = %d; want %d", tt.index, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 4. TABLE-DRIVEN TEST - ParsePositiveInt
// =============================================================

func TestParsePositiveInt(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  int
		expectErr error
	}{
		{"angka valid", "42", 42, nil},
		{"nol", "0", 0, nil},
		{"angka besar", "999999", 999999, nil},
		{"negatif", "-5", 0, ErrNegativeNumber},
		{"bukan angka", "abc", 0, ErrInvalidInput},
		{"string kosong", "", 0, ErrInvalidInput},
		{"campuran", "12abc", 0, ErrInvalidInput},
		{"spasi", " 42 ", 0, ErrInvalidInput},
		{"desimal", "3.14", 0, ErrInvalidInput},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParsePositiveInt(tt.input)

			if tt.expectErr != nil {
				if err == nil {
					t.Errorf("ParsePositiveInt(%q) expected error, got nil", tt.input)
					return
				}
				if !errors.Is(err, tt.expectErr) {
					t.Errorf("ParsePositiveInt(%q) error = %v; want %v", tt.input, err, tt.expectErr)
				}
				return
			}

			if err != nil {
				t.Errorf("ParsePositiveInt(%q) unexpected error: %v", tt.input, err)
				return
			}

			if result != tt.expected {
				t.Errorf("ParsePositiveInt(%q) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 5. TABLE-DRIVEN TEST - ParseAge
// =============================================================

func TestParseAge(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  int
		expectErr error
	}{
		{"usia valid 25", "25", 25, nil},
		{"usia valid 0", "0", 0, nil},
		{"usia valid 150", "150", 150, nil},
		{"usia terlalu besar", "200", 0, ErrOutOfBounds},
		{"usia negatif", "-5", 0, ErrOutOfBounds},
		{"bukan angka", "abc", 0, ErrInvalidInput},
		{"string kosong", "", 0, ErrInvalidInput},
		{"usia 151", "151", 0, ErrOutOfBounds},
		{"desimal", "25.5", 0, ErrInvalidInput},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseAge(tt.input)

			if tt.expectErr != nil {
				if err == nil {
					t.Errorf("ParseAge(%q) expected error, got nil", tt.input)
					return
				}
				if !errors.Is(err, tt.expectErr) {
					t.Errorf("ParseAge(%q) error = %v; want %v", tt.input, err, tt.expectErr)
				}
				return
			}

			if err != nil {
				t.Errorf("ParseAge(%q) unexpected error: %v", tt.input, err)
				return
			}

			if result != tt.expected {
				t.Errorf("ParseAge(%q) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 6. TABLE-DRIVEN TEST - ValidateEmail
// =============================================================

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expectErr bool
	}{
		{"email valid", "user@example.com", false},
		{"email valid subdomain", "user@mail.example.com", false},
		{"tanpa @", "userexample.com", true},
		{"tanpa domain", "@example.com", true},
		{"tanpa titik setelah @", "user@com", true},
		{"string kosong", "", true},
		{"titik di akhir", "user@example.", true},
		{"dua @", "user@@example.com", true},
		{"email pendek valid", "a@b.c", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEmail(tt.input)
			if tt.expectErr {
				if err == nil {
					t.Errorf("ValidateEmail(%q) expected error, got nil", tt.input)
				} else if !errors.Is(err, ErrInvalidFormat) {
					t.Errorf("ValidateEmail(%q) error = %v; want ErrInvalidFormat", tt.input, err)
				}
			} else {
				if err != nil {
					t.Errorf("ValidateEmail(%q) unexpected error: %v", tt.input, err)
				}
			}
		})
	}
}

// =============================================================
// 7. TABLE-DRIVEN TEST - SafeAverage
// =============================================================

func TestSafeAverage(t *testing.T) {
	tests := []struct {
		name      string
		input     []float64
		expected  float64
		expectErr error
	}{
		{"normal", []float64{10, 20, 30}, 20, nil},
		{"satu elemen", []float64{42}, 42, nil},
		{"slice kosong", []float64{}, 0, ErrEmptySlice},
		{"desimal", []float64{1.5, 2.5, 3.0}, 2.3333333333333335, nil},
		{"negatif", []float64{-10, 10}, 0, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SafeAverage(tt.input)

			if tt.expectErr != nil {
				if err == nil {
					t.Errorf("SafeAverage(%v) expected error, got nil", tt.input)
					return
				}
				if !errors.Is(err, tt.expectErr) {
					t.Errorf("SafeAverage(%v) error = %v; want %v", tt.input, err, tt.expectErr)
				}
				return
			}

			if err != nil {
				t.Errorf("SafeAverage(%v) unexpected error: %v", tt.input, err)
				return
			}

			if !approxEqual(result, tt.expected, epsilon) {
				t.Errorf("SafeAverage(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 8. TABLE-DRIVEN TEST - MinMax
// =============================================================

func TestMinMax(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		expectMin int
		expectMax int
		expectErr error
	}{
		{"normal", []int{3, 1, 4, 1, 5}, 1, 5, nil},
		{"satu elemen", []int{42}, 42, 42, nil},
		{"slice kosong", []int{}, 0, 0, ErrEmptySlice},
		{"semua sama", []int{7, 7, 7}, 7, 7, nil},
		{"negatif", []int{-5, -1, -10, -3}, -10, -1, nil},
		{"campuran", []int{-5, 0, 5, 10, -10}, -10, 10, nil},
		{"dua elemen", []int{100, 1}, 1, 100, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			min, max, err := MinMax(tt.input)

			if tt.expectErr != nil {
				if err == nil {
					t.Errorf("MinMax(%v) expected error, got nil", tt.input)
					return
				}
				if !errors.Is(err, tt.expectErr) {
					t.Errorf("MinMax(%v) error = %v; want %v", tt.input, err, tt.expectErr)
				}
				return
			}

			if err != nil {
				t.Errorf("MinMax(%v) unexpected error: %v", tt.input, err)
				return
			}

			if min != tt.expectMin || max != tt.expectMax {
				t.Errorf("MinMax(%v) = (%d, %d); want (%d, %d)", tt.input, min, max, tt.expectMin, tt.expectMax)
			}
		})
	}
}

// =============================================================
// 9. TABLE-DRIVEN TEST - ParseHexColor
// =============================================================

func TestParseHexColor(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		r, g, b   int
		expectErr bool
	}{
		{"putih", "#FFFFFF", 255, 255, 255, false},
		{"hitam", "#000000", 0, 0, 0, false},
		{"merah", "#FF0000", 255, 0, 0, false},
		{"hijau", "#00FF00", 0, 255, 0, false},
		{"biru", "#0000FF", 0, 0, 255, false},
		{"custom", "#FF8800", 255, 136, 0, false},
		{"lowercase", "#00ff00", 0, 255, 0, false},
		{"tanpa hash", "FF0000", 0, 0, 0, true},
		{"terlalu pendek", "#FFF", 0, 0, 0, true},
		{"terlalu panjang", "#FFFFFFF", 0, 0, 0, true},
		{"karakter invalid", "#GGGGGG", 0, 0, 0, true},
		{"string kosong", "", 0, 0, 0, true},
		{"mixed case", "#aaBBcc", 170, 187, 204, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, g, b, err := ParseHexColor(tt.input)

			if tt.expectErr {
				if err == nil {
					t.Errorf("ParseHexColor(%q) expected error, got nil", tt.input)
				}
				return
			}

			if err != nil {
				t.Errorf("ParseHexColor(%q) unexpected error: %v", tt.input, err)
				return
			}

			if r != tt.r || g != tt.g || b != tt.b {
				t.Errorf("ParseHexColor(%q) = (%d, %d, %d); want (%d, %d, %d)", tt.input, r, g, b, tt.r, tt.g, tt.b)
			}
		})
	}
}

// =============================================================
// 10. TEST - Retry
// =============================================================

func TestRetry(t *testing.T) {
	t.Run("berhasil pertama kali", func(t *testing.T) {
		fn := func() error { return nil }
		err := Retry(fn, 3)
		if err != nil {
			t.Errorf("Retry() = %v; want nil", err)
		}
	})

	t.Run("berhasil setelah 3 percobaan", func(t *testing.T) {
		attempt := 0
		fn := func() error {
			attempt++
			if attempt < 3 {
				return fmt.Errorf("gagal percobaan %d", attempt)
			}
			return nil
		}
		err := Retry(fn, 5)
		if err != nil {
			t.Errorf("Retry() = %v; want nil", err)
		}
		if attempt != 3 {
			t.Errorf("attempt = %d; want 3", attempt)
		}
	})

	t.Run("semua gagal", func(t *testing.T) {
		fn := func() error { return fmt.Errorf("selalu gagal") }
		err := Retry(fn, 3)
		if err == nil {
			t.Error("Retry() expected error, got nil")
		}
	})

	t.Run("maxAttempts nol", func(t *testing.T) {
		fn := func() error { return nil }
		err := Retry(fn, 0)
		if err == nil {
			t.Error("Retry(fn, 0) expected error, got nil")
		}
	})

	t.Run("maxAttempts negatif", func(t *testing.T) {
		fn := func() error { return nil }
		err := Retry(fn, -1)
		if err == nil {
			t.Error("Retry(fn, -1) expected error, got nil")
		}
	})

	t.Run("tepat berhasil di percobaan terakhir", func(t *testing.T) {
		attempt := 0
		fn := func() error {
			attempt++
			if attempt < 5 {
				return fmt.Errorf("gagal")
			}
			return nil
		}
		err := Retry(fn, 5)
		if err != nil {
			t.Errorf("Retry() = %v; want nil (berhasil di percobaan ke-5)", err)
		}
	})
}

// =============================================================
// 11. PARALLEL TEST - SafeDivide
// =============================================================

func TestSafeDivide_Parallel(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"case1", 100, 4, 25},
		{"case2", 9, 3, 3},
		{"case3", 1, 1, 1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result, err := SafeDivide(tt.a, tt.b)
			if err != nil {
				t.Errorf("SafeDivide(%v, %v) unexpected error: %v", tt.a, tt.b, err)
				return
			}
			if !approxEqual(result, tt.expected, epsilon) {
				t.Errorf("SafeDivide(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 12. BENCHMARK TEST
// =============================================================

func BenchmarkSafeDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SafeDivide(100, 3)
	}
}

func BenchmarkParseHexColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseHexColor("#FF8800")
	}
}

func BenchmarkRetry(b *testing.B) {
	fn := func() error { return nil }
	for i := 0; i < b.N; i++ {
		Retry(fn, 3)
	}
}
