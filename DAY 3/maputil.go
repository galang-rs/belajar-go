package belajar

import "sort"

// WordFrequency menghitung frekuensi kemunculan setiap kata dalam slice string.
// Penghitungan bersifat case-sensitive.
// Contoh: WordFrequency([]string{"go", "is", "go"}) -> map[string]int{"go": 2, "is": 1}
//
//	WordFrequency([]string{}) -> map[string]int{}
func WordFrequency(words []string) map[string]int {
	// TODO: implementasi di sini
	val := make(map[string]int)
	for _, v := range words {
		if val[v] == 0 {
			val[v] = 1
		} else {
			val[v]++
		}
	}
	return val
}

// InvertMap membalikkan key dan value dari map[string]string.
// Jika ada value duplikat di map asli, key mana pun yang diambil boleh (tidak ditentukan).
// Contoh: InvertMap(map[string]string{"a": "1", "b": "2"}) -> map[string]string{"1": "a", "2": "b"}
//
//	InvertMap(map[string]string{}) -> map[string]string{}
func InvertMap(m map[string]string) map[string]string {
	// TODO: implementasi di sini
	val := make(map[string]string)
	for k, v := range m {
		val[v] = k
	}

	return val
}

// MergeMaps menggabungkan dua map menjadi satu map baru.
// Jika ada key yang sama, value dari map kedua (m2) yang digunakan.
// Contoh: MergeMaps(map[string]int{"a": 1, "b": 2}, map[string]int{"b": 3, "c": 4})
//
//	-> map[string]int{"a": 1, "b": 3, "c": 4}
func MergeMaps(m1, m2 map[string]int) map[string]int {
	// TODO: implementasi di sini

	val := make(map[string]int)
	for k, v := range m1 {
		val[k] = v
	}
	for k, v := range m2 {
		val[k] = v
	}

	return val
}

// Keys mengembalikan semua key dari map dalam bentuk slice string.
// Urutan key tidak ditentukan (karena map di Go tidak terurut).
// Contoh: Keys(map[string]int{"a": 1, "b": 2, "c": 3}) -> []string{"a", "b", "c"} (urutan bisa beda)
func Keys(m map[string]int) []string {
	// TODO: implementasi di sini
	var val []string
	for k, _ := range m {
		val = append(val, k)
	}
	return val
}

// Values mengembalikan semua value dari map dalam bentuk slice int.
// Urutan value tidak ditentukan.
// Contoh: Values(map[string]int{"a": 1, "b": 2}) -> []int{1, 2} (urutan bisa beda)
func Values(m map[string]int) []int {
	// TODO: implementasi di sini
	var val []int
	for _, v := range m {
		val = append(val, v)
	}

	return val
}

// GroupByLength mengelompokkan slice string berdasarkan panjang string.
// Contoh: GroupByLength([]string{"go", "is", "fun", "ab"})
//
//	-> map[int][]string{2: {"go", "is", "ab"}, 3: {"fun"}}
//	GroupByLength([]string{}) -> map[int][]string{}
func GroupByLength(words []string) map[int][]string {
	// TODO: implementasi di sini
	val := make(map[int][]string)
	for _, v := range words {
		val[len(v)] = append(val[len(v)], v)
	}
	return val
}

// CountCharacters menghitung frekuensi setiap karakter (rune) dalam string.
// Case-sensitive.
// Contoh: CountCharacters("hello") -> map[rune]int{'h': 1, 'e': 1, 'l': 2, 'o': 1}
//
//	CountCharacters("") -> map[rune]int{}
func CountCharacters(s string) map[rune]int {
	// TODO: implementasi di sini
	val := make(map[rune]int)
	for _, v := range s {
		if val[rune(v)] == 0 {
			val[rune(v)] = 1
		} else {
			val[rune(v)]++
		}
	}
	return val
}

// MapContainsValue mengecek apakah sebuah value tertentu ada di dalam map.
// Contoh: MapContainsValue(map[string]int{"a": 1, "b": 2}, 2) -> true
//
//	MapContainsValue(map[string]int{"a": 1, "b": 2}, 3) -> false
//	MapContainsValue(map[string]int{}, 1) -> false
func MapContainsValue(m map[string]int, target int) bool {
	// TODO: implementasi di sini
	for _, v := range m {
		if v == target {
			return true
		}
	}
	return false
}

// DiffMaps mengembalikan map berisi key-value yang hanya ada di m1 tapi tidak ada di m2.
// Key dianggap "ada" di m2 hanya jika key tersebut ada DAN value-nya sama.
// Contoh: DiffMaps(map[string]int{"a": 1, "b": 2, "c": 3}, map[string]int{"a": 1, "c": 5})
//
//	-> map[string]int{"b": 2, "c": 3}
//	DiffMaps(map[string]int{"a": 1}, map[string]int{"a": 1}) -> map[string]int{}
func DiffMaps(m1, m2 map[string]int) map[string]int {
	// TODO: implementasi di sini
	for k1, v1 := range m1 {
		boolean := m2[k1]
		if boolean == v1 {
			delete(m1, k1)
		}
	}
	return m1
}

// TopNFrequent mengembalikan N kata yang paling sering muncul dari slice string,
// diurutkan dari frekuensi tertinggi ke terendah.
// Jika ada frekuensi yang sama, urutan berdasarkan kemunculan pertama di slice asli.
// Jika n lebih besar dari jumlah kata unik, kembalikan semua kata unik.
// Contoh: TopNFrequent([]string{"go", "is", "go", "fun", "is", "go"}, 2) -> []string{"go", "is"}
//
//	TopNFrequent([]string{"a", "b", "c"}, 5) -> []string{"a", "b", "c"}
//	TopNFrequent([]string{}, 3) -> []string{}
func TopNFrequent(words []string, n int) []string {
	// TODO: implementasi di sini
	val := make(map[string]int)
	for _, v := range words {
		val[v]++
	}

	key := make([]string, 0, len(val))
	for k := range val {
		key = append(key, k)
	}

	sort.Slice(key, func(i, j int) bool {
		if val[key[i]] == val[key[j]] {
			return key[i] < key[j]
		}
		return val[key[i]] > val[key[j]]
	})

	if n > len(key) {
		n = len(key)
	}
	return key[:n] // susah karena sortingnya urutannya karena terlalu ribet
}
