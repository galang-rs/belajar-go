package belajar

import (
	"math"
)

// Add menjumlahkan dua bilangan integer
// Contoh: Add(2, 3) -> 5
func Add(a, b int) int {
	// TODO: implementasi di sini
	return a + b
}

// Subtract mengurangi b dari a
// Contoh: Subtract(10, 3) -> 7
func Subtract(a, b int) int {
	// TODO: implementasi di sini
	return a - b
}

// Multiply mengalikan dua bilangan integer
// Contoh: Multiply(4, 5) -> 20
func Multiply(a, b int) int {
	// TODO: implementasi di sini
	return a * b
}

// Divide membagi a dengan b (float64)
// Kembalikan error jika b == 0
// Contoh: Divide(10, 2) -> 5.0, nil
//
//	Divide(10, 0) -> 0, error
func Divide(a, b float64) (float64, bool) {
	// TODO: implementasi di sini
	if b == 0 {
		return 0, true
	}
	return a / b, false
}

// Factorial menghitung faktorial dari n (n!)
// Kembalikan error jika n negatif
// Contoh: Factorial(5) -> 120, nil
//
//	Factorial(0) -> 1, nil
//	Factorial(-1) -> 0, error
func Factorial(n int) (int, bool) {
	// TODO: implementasi di sini
	if n < 0 {
		return 0, true
	}
	if n == 0 || n == 1 {
		return 1, false
	}
	result, err := Factorial(n - 1)
	return (n * result), err
}

// IsPrime mengecek apakah bilangan n adalah bilangan prima
// Contoh: IsPrime(7) -> true
//
//	IsPrime(4) -> false
//	IsPrime(1) -> false
func IsPrime(n int) bool {
	// TODO: implementasi di sini
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	limit := int(math.Sqrt(float64(n))) + 1
	for i := 3; i < limit; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Max mengembalikan nilai terbesar dari slice integer
// Kembalikan error jika slice kosong
// Contoh: Max([]int{1, 5, 3}) -> 5, nil
//
//	Max([]int{}) -> 0, error
func Max(numbers []int) (int, bool) {
	// TODO: implementasi di sini
	if len(numbers) == 0 {
		return 0, true
	}
	val := int(math.Inf(-1))
	for _, num := range numbers {
		if val < num {
			val = num
		}
	}
	return val, false
}

// FibonacciN mengembalikan bilangan Fibonacci ke-n (0-indexed)
// Kembalikan error jika n negatif
// Contoh: FibonacciN(0) -> 0, nil
//
//	FibonacciN(1) -> 1, nil
//	FibonacciN(6) -> 8, nil
func FibonacciN(n int) (int, bool) {
	// TODO: implementasi di sini
	if n < 0 {
		return 0, true
	}
	if n == 0 {
		return 0, false
	}
	if n == 1 {
		return 1, false
	}
	a, _ := FibonacciN(n - 1)
	b, _ := FibonacciN(n - 2)
	return a + b, false
}

// SumSlice menjumlahkan semua elemen dalam slice
// Contoh: SumSlice([]int{1, 2, 3}) -> 6
//
//	SumSlice([]int{}) -> 0
func SumSlice(numbers []int) int {
	// TODO: implementasi di sini
	val := 0
	for _, v := range numbers {
		val += v
	}
	return val
}

// Abs mengembalikan nilai absolut dari bilangan integer
// Contoh: Abs(-5) -> 5
//
//	Abs(3) -> 3
func Abs(n int) int {
	// TODO: implementasi di sini
	return int(math.Abs(float64(n)))
}
