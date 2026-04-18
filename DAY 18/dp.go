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
	if amount == 0 {
		return 0
	}
	if len(coins) == 0 {
		return -1
	}

	max := amount + 1
	dp := make(map[int]int)

	for i := 0; i <= amount; i++ {
		dp[i] = max
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, v := range coins {
			if i-v >= 0 {
				if dp[i-v]+1 < dp[i] {
					dp[i] = dp[i-v] + 1
				}
			}
		}
	}

	if dp[amount] == max {
		return -1
	}
	return dp[amount]
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
	n := len(nums)

	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	result := 0
	for i := 0; i < n; i++ {
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}
