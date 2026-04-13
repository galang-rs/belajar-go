package belajar

// ==================== DAY 13: CLOSURE & HIGHER-ORDER FUNCTIONS ====================
// Topik: Closure (fungsi yang "mengingat" variabel dari luar) dan function factory.
// Konsep penting: fungsi di Go adalah first-class citizen, bisa jadi return value.

// MakeCounter membuat closure counter yang dimulai dari 0.
// Setiap kali fungsi yang dikembalikan dipanggil, counter bertambah 1
// dan mengembalikan nilai counter saat ini.
// Contoh:
//
//	counter := MakeCounter()
//	counter() -> 1
//	counter() -> 2
//	counter() -> 3
//
//	counter2 := MakeCounter()
//	counter2() -> 1 (independent dari counter pertama)
//
// Hint: deklarasikan variabel di dalam MakeCounter, lalu return fungsi yang mengakses variabel itu.
func MakeCounter() func() int {
	// TODO: implementasi di sini
	return nil
}

// MakeMultiplier membuat closure yang mengalikan input dengan factor tertentu.
// Contoh:
//
//	double := MakeMultiplier(2)
//	double(5) -> 10
//	double(3) -> 6
//
//	triple := MakeMultiplier(3)
//	triple(4) -> 12
//	triple(0) -> 0
func MakeMultiplier(factor int) func(int) int {
	// TODO: implementasi di sini
	return nil
}
