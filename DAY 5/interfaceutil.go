package belajar

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
	return 0
}

// Perimeter mengembalikan keliling Rectangle (2 * (Width + Height)).
// Contoh: Rectangle{Width: 3, Height: 4}.Perimeter() -> 14
func (r Rectangle) Perimeter() float64 {
	// TODO: implementasi di sini
	return 0
}

// Area mengembalikan luas Circle (π * r²).
// Contoh: Circle{Radius: 5}.Area() -> 78.53981633974483
// Hint: gunakan math.Pi
func (c Circle) Area() float64 {
	// TODO: implementasi di sini
	return 0
}

// Perimeter mengembalikan keliling Circle (2 * π * r).
// Contoh: Circle{Radius: 5}.Perimeter() -> 31.41592653589793
func (c Circle) Perimeter() float64 {
	// TODO: implementasi di sini
	return 0
}

// Area mengembalikan luas Triangle menggunakan rumus Heron.
// s = (A+B+C)/2, luas = sqrt(s*(s-A)*(s-B)*(s-C))
// Contoh: Triangle{A: 3, B: 4, C: 5}.Area() -> 6
func (t Triangle) Area() float64 {
	// TODO: implementasi di sini
	// Hint: gunakan math.Sqrt
	return 0
}

// Perimeter mengembalikan keliling Triangle (A + B + C).
// Contoh: Triangle{A: 3, B: 4, C: 5}.Perimeter() -> 12
func (t Triangle) Perimeter() float64 {
	// TODO: implementasi di sini
	return 0
}

// TotalArea mengembalikan total luas dari semua Shape dalam slice.
// Contoh: TotalArea([]Shape{Rectangle{3, 4}, Circle{1}}) -> 12 + π ≈ 15.14159...
//
//	TotalArea([]Shape{}) -> 0
func TotalArea(shapes []Shape) float64 {
	// TODO: implementasi di sini
	return 0
}

// TotalPerimeter mengembalikan total keliling dari semua Shape dalam slice.
// Contoh: TotalPerimeter([]Shape{Rectangle{3, 4}, Triangle{3, 4, 5}}) -> 14 + 12 = 26
//
//	TotalPerimeter([]Shape{}) -> 0
func TotalPerimeter(shapes []Shape) float64 {
	// TODO: implementasi di sini
	return 0
}

// LargestShape mengembalikan Shape dengan luas terbesar dari slice.
// Kembalikan nil dan true (error) jika slice kosong.
// Jika ada yang sama, kembalikan yang pertama ditemukan.
// Contoh: LargestShape([]Shape{Rectangle{3, 4}, Rectangle{5, 6}}) -> Rectangle{5, 6}, false
//
//	LargestShape([]Shape{}) -> nil, true
func LargestShape(shapes []Shape) (Shape, bool) {
	// TODO: implementasi di sini
	return nil, true
}

// SmallestPerimeter mengembalikan Shape dengan keliling terkecil dari slice.
// Kembalikan nil dan true (error) jika slice kosong.
// Contoh: SmallestPerimeter([]Shape{Rectangle{3, 4}, Rectangle{1, 1}}) -> Rectangle{1, 1}, false
//
//	SmallestPerimeter([]Shape{}) -> nil, true
func SmallestPerimeter(shapes []Shape) (Shape, bool) {
	// TODO: implementasi di sini
	return nil, true
}

// FilterByMinArea mengembalikan slice Shape yang luasnya >= minArea.
// Contoh: FilterByMinArea([]Shape{Rectangle{3, 4}, Rectangle{1, 1}}, 5) -> []Shape{Rectangle{3, 4}}
//
//	FilterByMinArea([]Shape{}, 10) -> []Shape{}
func FilterByMinArea(shapes []Shape, minArea float64) []Shape {
	// TODO: implementasi di sini
	return nil
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
	return ""
}
