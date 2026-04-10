package belajar

import (
	"fmt"
	"math"
)

// Shape adalah interface untuk bentuk geometri.
// Setiap bentuk harus bisa menghitung luas dan keliling.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle merepresentasikan persegi panjang.
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle merepresentasikan lingkaran.
type Circle struct {
	Radius float64
}

// Triangle merepresentasikan segitiga dengan tiga sisi.
type Triangle struct {
	A, B, C float64
}

// Area mengembalikan luas Rectangle (Width * Height).
// Contoh: Rectangle{Width: 3, Height: 4}.Area() -> 12
func (r Rectangle) Area() float64 {
	// TODO: implementasi di sini
	return r.Width * r.Height
}

// Perimeter mengembalikan keliling Rectangle (2 * (Width + Height)).
// Contoh: Rectangle{Width: 3, Height: 4}.Perimeter() -> 14
func (r Rectangle) Perimeter() float64 {
	// TODO: implementasi di sini
	return (2 * (r.Width + r.Height))
}

// Area mengembalikan luas Circle (π * r²).
// Contoh: Circle{Radius: 5}.Area() -> 78.53981633974483
// Hint: gunakan math.Pi
func (c Circle) Area() float64 {
	// TODO: implementasi di sini
	return math.Pi * (c.Radius * c.Radius)
}

// Perimeter mengembalikan keliling Circle (2 * π * r).
// Contoh: Circle{Radius: 5}.Perimeter() -> 31.41592653589793
func (c Circle) Perimeter() float64 {
	// TODO: implementasi di sini
	return 2 * math.Pi * c.Radius
}

// Area mengembalikan luas Triangle menggunakan rumus Heron.
// s = (A+B+C)/2, luas = sqrt(s*(s-A)*(s-B)*(s-C))
// Contoh: Triangle{A: 3, B: 4, C: 5}.Area() -> 6
func (t Triangle) Area() float64 {
	// TODO: implementasi di sini
	// Hint: gunakan math.Sqrt
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

// Perimeter mengembalikan keliling Triangle (A + B + C).
// Contoh: Triangle{A: 3, B: 4, C: 5}.Perimeter() -> 12
func (t Triangle) Perimeter() float64 {
	// TODO: implementasi di sini
	return t.A + t.B + t.C
}

// TotalArea mengembalikan total luas dari semua Shape dalam slice.
// Contoh: TotalArea([]Shape{Rectangle{3, 4}, Circle{1}}) -> 12 + π ≈ 15.14159...
//
//	TotalArea([]Shape{}) -> 0
func TotalArea(shapes []Shape) float64 {
	// TODO: implementasi di sini
	var totalVal float64
	for k := range shapes {
		totalVal += shapes[k].Area()
	}
	return totalVal
}

// TotalPerimeter mengembalikan total keliling dari semua Shape dalam slice.
// Contoh: TotalPerimeter([]Shape{Rectangle{3, 4}, Triangle{3, 4, 5}}) -> 14 + 12 = 26
//
//	TotalPerimeter([]Shape{}) -> 0
func TotalPerimeter(shapes []Shape) float64 {
	// TODO: implementasi di sini
	var totalVal float64
	for k := range shapes {
		totalVal += shapes[k].Perimeter()
	}
	return totalVal
}

// LargestShape mengembalikan Shape dengan luas terbesar dari slice.
// Kembalikan nil dan true (error) jika slice kosong.
// Jika ada yang sama, kembalikan yang pertama ditemukan.
// Contoh: LargestShape([]Shape{Rectangle{3, 4}, Rectangle{5, 6}}) -> Rectangle{5, 6}, false
//
//	LargestShape([]Shape{}) -> nil, true
func LargestShape(shapes []Shape) (Shape, bool) {
	// TODO: implementasi di sini
	if len(shapes) == 0 {
		return nil, true
	}

	var shape Shape
	var high float64
	for k := range shapes {
		if shapes[k].Area() > high {
			high = shapes[k].Area()
			shape = shapes[k]
		}
	}
	return shape, false
}

// SmallestPerimeter mengembalikan Shape dengan keliling terkecil dari slice.
// Kembalikan nil dan true (error) jika slice kosong.
// Contoh: SmallestPerimeter([]Shape{Rectangle{3, 4}, Rectangle{1, 1}}) -> Rectangle{1, 1}, false
//
//	SmallestPerimeter([]Shape{}) -> nil, true
func SmallestPerimeter(shapes []Shape) (Shape, bool) {
	// TODO: implementasi di sini
	if len(shapes) == 0 {
		return nil, true
	}

	var shape Shape
	var lower float64 = math.Inf(1)
	for k := range shapes {
		if shapes[k].Perimeter() < lower {
			lower = shapes[k].Perimeter()
			shape = shapes[k]
		}
	}
	return shape, false
}

// FilterByMinArea mengembalikan slice Shape yang luasnya >= minArea.
// Contoh: FilterByMinArea([]Shape{Rectangle{3, 4}, Rectangle{1, 1}}, 5) -> []Shape{Rectangle{3, 4}}
//
//	FilterByMinArea([]Shape{}, 10) -> []Shape{}
func FilterByMinArea(shapes []Shape, minArea float64) []Shape {
	// TODO: implementasi di sini
	var shapesVal []Shape

	for k := range shapes {
		if shapes[k].Area() >= minArea {
			shapesVal = append(shapesVal, shapes[k])
		}
	}

	return shapesVal
}

// Describe mengembalikan deskripsi string dari sebuah Shape berdasarkan tipenya.
// Format:
//   - Rectangle -> "Persegi Panjang: 3.00 x 4.00"
//   - Circle    -> "Lingkaran: r=5.00"
//   - Triangle  -> "Segitiga: 3.00, 4.00, 5.00"
//   - Lainnya   -> "Bentuk tidak dikenal"
//
// Hint: gunakan type switch dan fmt.Sprintf
func Describe(s Shape) string {
	// TODO: implementasi di sini
	switch v := s.(type) {
	case Rectangle:
		return fmt.Sprintf("Persegi Panjang: %.2f x %.2f", v.Width, v.Height)
	case Circle:
		return fmt.Sprintf("Lingkaran: r=%.2f", v.Radius)
	case Triangle:
		return fmt.Sprintf("Segitiga: %.2f, %.2f, %.2f", v.A, v.B, v.C)
	default:
		return "Bentuk tidak dikenal"
	}
}
