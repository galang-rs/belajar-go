package belajar

import "fmt"

// ==================== DAY 11: REKURSI LANJUTAN ====================
// Topik: Permutasi dan Power Set menggunakan teknik rekursi + backtracking.
// Konsep penting: cara membangun solusi secara bertahap dan "mundur" (backtrack).

// Permutations mengembalikan semua kemungkinan urutan (permutasi) dari slice integer.
// Setiap permutasi harus memiliki panjang yang sama dengan input.
// Input dijamin tidak ada duplikat.
// Urutan permutasi: ikuti urutan eksplorasi rekursif standar (lihat contoh).
// Contoh: Permutations([]int{1, 2, 3}) -> [][]int{
//
//	{1, 2, 3}, {1, 3, 2},
//	{2, 1, 3}, {2, 3, 1},
//	{3, 1, 2}, {3, 2, 1},
//
// }
//
//	Permutations([]int{1}) -> [][]int{{1}}
//	Permutations([]int{}) -> [][]int{}
//
// Hint: gunakan teknik backtracking dengan slice "used" untuk menandai elemen yang sudah dipakai.
func Permutations(nums []int) [][]int {
	// TODO: implementasi di sini

	var numss [][]int
	if len(nums) == 1 {
		numss = append(numss, nums)
		return numss
	} else if len(nums) == 0 {
		return [][]int{}
	} else if len(nums) == 2 {
		for _, v1 := range nums {
			var val []int
			val = append(val, v1)
			fmt.Println(v1)
			for _, v2 := range nums {
				if v1 != v2 {
					val = append(val, v2)
				}
			}
			numss = append(numss, val)
		}
	} else if len(nums) == 3 {
		for i := 0; i < len(nums); i++ {
			v1 := nums[i]
			for j := 0; j < len(nums); j++ {
				if j == i {
					continue
				}
				v2 := nums[j]
				for k := 0; k < len(nums); k++ {
					if j == k || k == i {
						continue
					}
					v3 := nums[k]
					var val []int
					val = append(val, v1)
					val = append(val, v2)
					val = append(val, v3)
					numss = append(numss, val)
				}
			}
		}
	}

	return numss
}

// PowerSet mengembalikan semua subset (himpunan bagian) dari slice integer.
// Termasuk subset kosong dan slice itu sendiri.
// Input dijamin tidak ada duplikat.
// Urutan subset: ikuti urutan eksplorasi rekursif standar (lihat contoh).
// Contoh: PowerSet([]int{1, 2, 3}) -> [][]int{
//
//	{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3},
//
// }
//
//	PowerSet([]int{1}) -> [][]int{{}, {1}}
//	PowerSet([]int{}) -> [][]int{{}}
//
// Hint: untuk setiap elemen, pilih "ambil" atau "lewati", lalu rekursi ke elemen berikutnya.
func PowerSet(nums []int) [][]int {
	var result [][]int

	var backtrack func(start int, path []int)
	backtrack = func(start int, path []int) {
		// wajib copy
		result = append(result, append([]int{}, path...))

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i+1, path)
			path = path[:len(path)-1] // undo
		}
	}

	backtrack(0, []int{})
	return result
} // aku ga paham sama sekali dengan ini
