package belajar

// ==================== DAY 18: DYNAMIC PROGRAMMING ====================
// Topik: Dynamic Programming (DP) - memecah masalah besar menjadi sub-masalah kecil.
// Konsep penting: menyimpan hasil perhitungan sebelumnya untuk menghindari perhitungan ulang.

// CoinChange mengembalikan jumlah minimum koin yang dibutuhkan untuk membuat sejumlah amount.
// Setiap denominasi koin bisa dipakai berkali-kali (unlimited supply).
// Jika tidak mungkin membuat amount tersebut, kembalikan -1.
// Contoh: CoinChange([]int{1, 5, 10, 25}, 30) -> 2 (25 + 5)
//
//	CoinChange([]int{1, 5, 10}, 11) -> 2 (10 + 1)
//	CoinChange([]int{2}, 3) -> -1 (tidak bisa membuat 3 dari koin 2)
//	CoinChange([]int{1}, 0) -> 0 (tidak perlu koin)
//	CoinChange([]int{1, 2, 5}, 11) -> 3 (5 + 5 + 1)
//	CoinChange([]int{}, 10) -> -1
//
// Hint: buat array dp[0..amount], dp[i] = minimum koin untuk membuat nilai i.
// dp[0] = 0, dp[i] = min(dp[i-coin] + 1) untuk setiap coin.
func CoinChange(coins []int, amount int) int {
	// TODO: implementasi di sini
	return -1
}

// LongestIncreasingSubsequence mengembalikan panjang subsequence (tidak harus berurutan)
// terpanjang yang nilainya selalu naik (strictly increasing).
// Contoh: LongestIncreasingSubsequence([]int{10, 9, 2, 5, 3, 7, 101, 18}) -> 4
//
//	// Subsequence: [2, 3, 7, 101] atau [2, 5, 7, 101] atau [2, 3, 7, 18]
//
//	LongestIncreasingSubsequence([]int{0, 1, 0, 3, 2, 3}) -> 4  // [0, 1, 2, 3]
//	LongestIncreasingSubsequence([]int{7, 7, 7, 7}) -> 1  // semua sama
//	LongestIncreasingSubsequence([]int{5}) -> 1
//	LongestIncreasingSubsequence([]int{}) -> 0
//	LongestIncreasingSubsequence([]int{1, 2, 3, 4, 5}) -> 5  // sudah terurut
//	LongestIncreasingSubsequence([]int{5, 4, 3, 2, 1}) -> 1  // terbalik
//
// Hint: buat array dp[i] = panjang LIS yang berakhir di index i.
// dp[i] = max(dp[j] + 1) untuk semua j < i dimana nums[j] < nums[i].
func LongestIncreasingSubsequence(nums []int) int {
	// TODO: implementasi di sini
	return 0
}
