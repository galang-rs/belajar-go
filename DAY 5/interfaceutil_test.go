package belajar

import (
	"math"
	"testing"
)

// Helper: perbandingan float64 dengan toleransi
func approxEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) < tolerance
}

const epsilon = 1e-6

// =============================================================
// 1. TABLE-DRIVEN TEST - Rectangle Area & Perimeter
// =============================================================

func TestRectangleArea(t *testing.T) {
	tests := []struct {
		name     string
		rect     Rectangle
		expected float64
	}{
		{"normal", Rectangle{3, 4}, 12},
		{"persegi", Rectangle{5, 5}, 25},
		{"sisi nol", Rectangle{0, 10}, 0},
		{"desimal", Rectangle{2.5, 4}, 10},
		{"besar", Rectangle{100, 200}, 20000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.rect.Area()
			if !approxEqual(result, tt.expected, epsilon) {
				t.Errorf("Rectangle{%v, %v}.Area() = %v; want %v", tt.rect.Width, tt.rect.Height, result, tt.expected)
			}
		})
	}
}

func TestRectanglePerimeter(t *testing.T) {
	tests := []struct {
		name     string
		rect     Rectangle
		expected float64
	}{
		{"normal", Rectangle{3, 4}, 14},
		{"persegi", Rectangle{5, 5}, 20},
		{"sisi nol", Rectangle{0, 10}, 20},
		{"desimal", Rectangle{2.5, 4}, 13},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.rect.Perimeter()
			if !approxEqual(result, tt.expected, epsilon) {
				t.Errorf("Rectangle{%v, %v}.Perimeter() = %v; want %v", tt.rect.Width, tt.rect.Height, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 2. TABLE-DRIVEN TEST - Circle Area & Perimeter
// =============================================================

func TestCircleArea(t *testing.T) {
	tests := []struct {
		name     string
		circle   Circle
		expected float64
	}{
		{"radius 1", Circle{1}, math.Pi},
		{"radius 5", Circle{5}, math.Pi * 25},
		{"radius 0", Circle{0}, 0},
		{"radius 10", Circle{10}, math.Pi * 100},
		{"radius 0.5", Circle{0.5}, math.Pi * 0.25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.circle.Area()
			if !approxEqual(result, tt.expected, epsilon) {
				t.Errorf("Circle{%v}.Area() = %v; want %v", tt.circle.Radius, result, tt.expected)
			}
		})
	}
}

func TestCirclePerimeter(t *testing.T) {
	tests := []struct {
		name     string
		circle   Circle
		expected float64
	}{
		{"radius 1", Circle{1}, 2 * math.Pi},
		{"radius 5", Circle{5}, 2 * math.Pi * 5},
		{"radius 0", Circle{0}, 0},
		{"radius 10", Circle{10}, 2 * math.Pi * 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.circle.Perimeter()
			if !approxEqual(result, tt.expected, epsilon) {
				t.Errorf("Circle{%v}.Perimeter() = %v; want %v", tt.circle.Radius, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 3. TABLE-DRIVEN TEST - Triangle Area & Perimeter
// =============================================================

func TestTriangleArea(t *testing.T) {
	tests := []struct {
		name     string
		tri      Triangle
		expected float64
	}{
		{"segitiga siku-siku 3-4-5", Triangle{3, 4, 5}, 6},
		{"segitiga sama sisi 3", Triangle{3, 3, 3}, math.Sqrt(3) / 4 * 9},
		{"segitiga besar", Triangle{13, 14, 15}, 84},
		{"segitiga 5-12-13", Triangle{5, 12, 13}, 30},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.tri.Area()
			if !approxEqual(result, tt.expected, epsilon) {
				t.Errorf("Triangle{%v, %v, %v}.Area() = %v; want %v", tt.tri.A, tt.tri.B, tt.tri.C, result, tt.expected)
			}
		})
	}
}

func TestTrianglePerimeter(t *testing.T) {
	tests := []struct {
		name     string
		tri      Triangle
		expected float64
	}{
		{"3-4-5", Triangle{3, 4, 5}, 12},
		{"sama sisi", Triangle{3, 3, 3}, 9},
		{"besar", Triangle{10, 20, 25}, 55},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.tri.Perimeter()
			if !approxEqual(result, tt.expected, epsilon) {
				t.Errorf("Triangle{%v, %v, %v}.Perimeter() = %v; want %v", tt.tri.A, tt.tri.B, tt.tri.C, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 4. TABLE-DRIVEN TEST - TotalArea
// =============================================================

func TestTotalArea(t *testing.T) {
	tests := []struct {
		name     string
		shapes   []Shape
		expected float64
	}{
		{"kosong", []Shape{}, 0},
		{"satu rectangle", []Shape{Rectangle{3, 4}}, 12},
		{"rectangle + circle", []Shape{Rectangle{3, 4}, Circle{1}}, 12 + math.Pi},
		{"campuran", []Shape{Rectangle{3, 4}, Circle{5}, Triangle{3, 4, 5}}, 12 + math.Pi*25 + 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TotalArea(tt.shapes)
			if !approxEqual(result, tt.expected, epsilon) {
				t.Errorf("TotalArea(...) = %v; want %v", result, tt.expected)
			}
		})
	}
}

// =============================================================
// 5. TABLE-DRIVEN TEST - TotalPerimeter
// =============================================================

func TestTotalPerimeter(t *testing.T) {
	tests := []struct {
		name     string
		shapes   []Shape
		expected float64
	}{
		{"kosong", []Shape{}, 0},
		{"satu rectangle", []Shape{Rectangle{3, 4}}, 14},
		{"rectangle + triangle", []Shape{Rectangle{3, 4}, Triangle{3, 4, 5}}, 14 + 12},
		{"campuran", []Shape{Rectangle{3, 4}, Circle{1}, Triangle{3, 4, 5}}, 14 + 2*math.Pi + 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TotalPerimeter(tt.shapes)
			if !approxEqual(result, tt.expected, epsilon) {
				t.Errorf("TotalPerimeter(...) = %v; want %v", result, tt.expected)
			}
		})
	}
}

// =============================================================
// 6. TEST - LargestShape
// =============================================================

func TestLargestShape(t *testing.T) {
	t.Run("slice kosong", func(t *testing.T) {
		_, err := LargestShape([]Shape{})
		if !err {
			t.Error("LargestShape([]) expected error, got none")
		}
	})

	t.Run("satu shape", func(t *testing.T) {
		r := Rectangle{3, 4}
		result, err := LargestShape([]Shape{r})
		if err {
			t.Error("LargestShape() unexpected error")
			return
		}
		if result != r {
			t.Errorf("LargestShape([Rectangle{3,4}]) = %v; want %v", result, r)
		}
	})

	t.Run("rectangle vs circle", func(t *testing.T) {
		r := Rectangle{3, 4}     // area = 12
		c := Circle{5}           // area = ~78.5
		result, err := LargestShape([]Shape{r, c})
		if err {
			t.Error("unexpected error")
			return
		}
		if result != c {
			t.Errorf("LargestShape() = %v; want Circle{5}", result)
		}
	})

	t.Run("banyak shapes", func(t *testing.T) {
		shapes := []Shape{
			Rectangle{2, 3},  // area = 6
			Triangle{3, 4, 5}, // area = 6
			Circle{10},        // area = ~314
			Rectangle{20, 20}, // area = 400
		}
		result, err := LargestShape(shapes)
		if err {
			t.Error("unexpected error")
			return
		}
		if result != shapes[3] {
			t.Errorf("LargestShape() = %v; want Rectangle{20, 20}", result)
		}
	})
}

// =============================================================
// 7. TEST - SmallestPerimeter
// =============================================================

func TestSmallestPerimeter(t *testing.T) {
	t.Run("slice kosong", func(t *testing.T) {
		_, err := SmallestPerimeter([]Shape{})
		if !err {
			t.Error("SmallestPerimeter([]) expected error, got none")
		}
	})

	t.Run("satu shape", func(t *testing.T) {
		r := Rectangle{1, 1}
		result, err := SmallestPerimeter([]Shape{r})
		if err {
			t.Error("unexpected error")
			return
		}
		if result != r {
			t.Errorf("SmallestPerimeter() = %v; want %v", result, r)
		}
	})

	t.Run("banyak shapes", func(t *testing.T) {
		shapes := []Shape{
			Rectangle{10, 20}, // perimeter = 60
			Circle{1},         // perimeter = ~6.28
			Triangle{3, 4, 5}, // perimeter = 12
		}
		result, err := SmallestPerimeter(shapes)
		if err {
			t.Error("unexpected error")
			return
		}
		if result != shapes[1] {
			t.Errorf("SmallestPerimeter() = %v; want Circle{1}", result)
		}
	})
}

// =============================================================
// 8. TABLE-DRIVEN TEST - FilterByMinArea
// =============================================================

func TestFilterByMinArea(t *testing.T) {
	shapes := []Shape{
		Rectangle{2, 3},   // area = 6
		Circle{1},          // area = ~3.14
		Triangle{3, 4, 5},  // area = 6
		Rectangle{10, 10},  // area = 100
	}

	t.Run("filter area >= 5", func(t *testing.T) {
		result := FilterByMinArea(shapes, 5)
		if len(result) != 3 {
			t.Errorf("FilterByMinArea(..., 5) returned %d shapes; want 3", len(result))
		}
	})

	t.Run("filter area >= 50", func(t *testing.T) {
		result := FilterByMinArea(shapes, 50)
		if len(result) != 1 {
			t.Errorf("FilterByMinArea(..., 50) returned %d shapes; want 1", len(result))
		}
	})

	t.Run("filter area >= 0", func(t *testing.T) {
		result := FilterByMinArea(shapes, 0)
		if len(result) != 4 {
			t.Errorf("FilterByMinArea(..., 0) returned %d shapes; want 4", len(result))
		}
	})

	t.Run("tidak ada yang lolos", func(t *testing.T) {
		result := FilterByMinArea(shapes, 500)
		if len(result) != 0 {
			t.Errorf("FilterByMinArea(..., 500) returned %d shapes; want 0", len(result))
		}
	})

	t.Run("slice kosong", func(t *testing.T) {
		result := FilterByMinArea([]Shape{}, 10)
		if len(result) != 0 {
			t.Errorf("FilterByMinArea([], 10) returned %d shapes; want 0", len(result))
		}
	})
}

// =============================================================
// 9. TABLE-DRIVEN TEST - Describe
// =============================================================

func TestDescribe(t *testing.T) {
	tests := []struct {
		name     string
		shape    Shape
		expected string
	}{
		{"rectangle", Rectangle{3, 4}, "Persegi Panjang: 3.00 x 4.00"},
		{"rectangle desimal", Rectangle{2.5, 3.75}, "Persegi Panjang: 2.50 x 3.75"},
		{"circle", Circle{5}, "Lingkaran: r=5.00"},
		{"circle kecil", Circle{0.5}, "Lingkaran: r=0.50"},
		{"triangle", Triangle{3, 4, 5}, "Segitiga: 3.00, 4.00, 5.00"},
		{"triangle desimal", Triangle{2.5, 3.5, 4.5}, "Segitiga: 2.50, 3.50, 4.50"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Describe(tt.shape)
			if result != tt.expected {
				t.Errorf("Describe(%v) = %q; want %q", tt.shape, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 10. INTERFACE COMPLIANCE TEST
// =============================================================

func TestShapeInterfaceCompliance(t *testing.T) {
	// Memastikan semua tipe mengimplementasikan Shape interface
	var _ Shape = Rectangle{}
	var _ Shape = Circle{}
	var _ Shape = Triangle{}
}

// =============================================================
// 11. PARALLEL TEST
// =============================================================

func TestTotalArea_Parallel(t *testing.T) {
	tests := []struct {
		name     string
		shapes   []Shape
		expected float64
	}{
		{"case1", []Shape{Rectangle{1, 1}}, 1},
		{"case2", []Shape{Circle{1}}, math.Pi},
		{"case3", []Shape{Triangle{3, 4, 5}}, 6},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := TotalArea(tt.shapes)
			if !approxEqual(result, tt.expected, epsilon) {
				t.Errorf("TotalArea(...) = %v; want %v", result, tt.expected)
			}
		})
	}
}

// =============================================================
// 12. BENCHMARK TEST
// =============================================================

func BenchmarkTotalArea(b *testing.B) {
	shapes := []Shape{
		Rectangle{3, 4},
		Circle{5},
		Triangle{3, 4, 5},
		Rectangle{10, 20},
		Circle{7},
	}
	for i := 0; i < b.N; i++ {
		TotalArea(shapes)
	}
}

func BenchmarkDescribe(b *testing.B) {
	shapes := []Shape{Rectangle{3, 4}, Circle{5}, Triangle{3, 4, 5}}
	for i := 0; i < b.N; i++ {
		for _, s := range shapes {
			Describe(s)
		}
	}
}
