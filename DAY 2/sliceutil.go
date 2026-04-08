package belajar

// Reverse membalikkan urutan elemen dalam slice.
// Contoh: nums := []int{1, 2, 3}; Reverse(nums) -> nums menjadi []int{3, 2, 1}
func Reverse(nums []int) []int {
	// TODO: implementasi di sini
	numbers := nums
	for a, b := 0, len(numbers)-1; a < b; a, b = a+1, b-1 {
		numbers[a], numbers[b] = numbers[b], numbers[a]
	}
	return numbers
}

// Unique membuang elemen duplikat dan mengembalikan slice baru dengan elemen unik.
// Urutan elemen unik harus sesuai dengan kemunculan pertama di slice asli.
// Contoh: Unique([]int{1, 2, 2, 3, 1}) -> []int{1, 2, 3}
func Unique(nums []int) []int {
	// TODO: implementasi di sini
	seen := make(map[int]bool)
	result := []int{}
	for _, v := range nums {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// Intersect mengembalikan irisan dari dua slice (elemen yang ada di kedua slice).
// Tiap angka pada irisan hanya boleh muncul sekali (unik).
// Contoh: Intersect([]int{1, 2, 2, 1}, []int{2, 2, 3}) -> []int{2}
func Intersect(nums1 []int, nums2 []int) []int {
	// TODO: implementasi di sini
	seen := make(map[int]bool)
	result := []int{}
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			if v1 == v2 && !seen[v1] && !seen[v2] {
				seen[v1] = true
				result = append(result, v1)
			}
		}
	}
	return result
}

// Contains mengecek apakah nilai target terdapat di dalam slice.
// Contoh: Contains([]int{1, 2, 3}, 2) -> true
//
//	Contains([]int{1, 2, 3}, 4) -> false
func Contains(nums []int, target int) bool {
	// TODO: implementasi di sini
	for _, v := range nums {
		if v == target {
			return true
		}
	}
	return false
}

// GroupByParity mengelompokkan elemen menjadi bilangan genap dan ganjil.
// Mengembalikan slice genap pada return pertama dan slice ganjil pada return kedua.
// Contoh: GroupByParity([]int{1, 2, 3, 4, 5}) -> []int{2, 4}, []int{1, 3, 5}
func GroupByParity(nums []int) ([]int, []int) {
	// TODO: implementasi di sini
	seenGenap := make(map[int]bool)
	genap := []int{}
	ganjil := []int{}
	for _, v := range nums {
		if v%2 == 0 {
			seenGenap[v] = true
			genap = append(genap, v)
		} else {
			ganjil = append(ganjil, v)
		}
	}
	return genap, ganjil
}

// RemoveAt menghapus elemen pada indeks tertentu dan mengembalikan slice baru.
// Kembalikan bool true jika terjadi error / index tidak valid.
// Contoh: RemoveAt([]int{10, 20, 30}, 1) -> []int{10, 30}, false
//
//	RemoveAt([]int{10, 20, 30}, 5) -> nil, true
func RemoveAt(nums []int, index int) ([]int, bool) {
	// TODO: implementasi di sini
	if len(nums) <= 0 || index < 0 || len(nums) <= index {
		return nil, true
	}
	result := []int{}
	for k, v := range nums {
		if !(k == index) {
			result = append(result, v)
		}
	}
	return result, false
}

// Chunk membagi slice menjadi beberapa sub-slice dengan ukuran tertentu.
// Sub-slice terakhir bisa lebih kecil jika elemen tidak habis dibagi.
// Kembalikan nil jika size <= 0.
// Contoh: Chunk([]int{1, 2, 3, 4, 5}, 2) -> [][]int{{1, 2}, {3, 4}, {5}}
//
//	Chunk([]int{1, 2, 3}, 5) -> [][]int{{1, 2, 3}}
//	Chunk([]int{}, 3) -> [][]int{}
func Chunk(nums []int, size int) [][]int {
	// TODO: implementasi di sini
	if size <= 0 {
		return nil
	}

	var result [][]int
	var chunk []int

	for _, v := range nums {
		chunk = append(chunk, v)

		if len(chunk) == size {
			result = append(result, chunk)
			chunk = []int{}
		}
	}

	// sisa (kalau ada)
	if len(chunk) > 0 {
		result = append(result, chunk)
	}

	return result
}

// Flatten menggabungkan slice 2D menjadi slice 1D.
// Contoh: Flatten([][]int{{1, 2}, {3, 4}, {5}}) -> []int{1, 2, 3, 4, 5}
//
//	Flatten([][]int{{}, {1}, {}}) -> []int{1}
//	Flatten([][]int{}) -> []int{}
func Flatten(nums [][]int) []int {
	// TODO: implementasi di sini
	var result []int

	for _, v := range nums {
		for _, v1 := range v {
			result = append(result, v1)
		}
	}

	return result
}

// Map menerapkan fungsi transformasi ke setiap elemen slice dan mengembalikan slice baru.
// Contoh: Map([]int{1, 2, 3}, func(n int) int { return n * 2 }) -> []int{2, 4, 6}
//
//	Map([]int{}, func(n int) int { return n + 1 }) -> []int{}
func Map(nums []int, fn func(int) int) []int {
	// TODO: implementasi di sini
	var data []int
	for _, v := range nums {
		data = append(data, fn(v))
	}
	return data
}

// Filter mengembalikan slice baru berisi elemen yang memenuhi kondisi (fungsi predicate).
// Contoh: Filter([]int{1, 2, 3, 4, 5}, func(n int) bool { return n > 3 }) -> []int{4, 5}
//
//	Filter([]int{1, 2, 3}, func(n int) bool { return n > 10 }) -> []int{}
func Filter(nums []int, predicate func(int) bool) []int {
	// TODO: implementasi di sini
	var data []int
	for _, v := range nums {
		if predicate(v) {
			data = append(data, v)
		}
	}
	return data
}

// Reduce mengakumulasi nilai dari slice menggunakan fungsi akumulator dan nilai awal.
// Contoh: Reduce([]int{1, 2, 3, 4}, func(acc, n int) int { return acc + n }, 0) -> 10
//
//	Reduce([]int{1, 2, 3}, func(acc, n int) int { return acc * n }, 1) -> 6
//	Reduce([]int{}, func(acc, n int) int { return acc + n }, 5) -> 5
func Reduce(nums []int, fn func(int, int) int, initial int) int {
	// TODO: implementasi di sini
	number := initial
	for _, v := range nums {
		number = fn(number, v)
	}
	return number
}

// Zip menggabungkan dua slice menjadi slice pasangan [2]int.
// Panjang hasil ditentukan oleh slice yang lebih pendek.
// Contoh: Zip([]int{1, 2, 3}, []int{4, 5, 6}) -> [][2]int{{1, 4}, {2, 5}, {3, 6}}
//
//	Zip([]int{1, 2}, []int{3, 4, 5}) -> [][2]int{{1, 3}, {2, 4}}
//	Zip([]int{}, []int{1, 2}) -> [][2]int{}
func Zip(a, b []int) [][2]int {
	// TODO: implementasi di sini
	length := len(a)
	if len(a) > len(b) {
		length = len(b)
	}
	var result [][2]int
	for i := 0; i < length; i++ {
		result = append(result, [2]int{a[i], b[i]})
	}
	return result
}
