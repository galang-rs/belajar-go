package belajar

// ==================== DAY 20: FINAL CHALLENGE ====================
// Topik: Soal tantangan akhir - menggabungkan berbagai konsep yang sudah dipelajari.
// Konsep: matrix spiral dan validasi Sudoku.

// SpiralMatrix menghasilkan matrix N x N yang diisi angka 1 sampai N² dalam pola spiral
// searah jarum jam, dimulai dari pojok kiri atas.
// Contoh: SpiralMatrix(3) -> [][]int{
//
//	{1, 2, 3},
//	{8, 9, 4},
//	{7, 6, 5},
//
// }
//
//	SpiralMatrix(4) -> [][]int{
//	    {1,  2,  3,  4},
//	    {12, 13, 14, 5},
//	    {11, 16, 15, 6},
//	    {10, 9,  8,  7},
//	}
//
//	SpiralMatrix(1) -> [][]int{{1}}
//	SpiralMatrix(0) -> [][]int{}
//
// Hint: gunakan 4 variabel batas: top, bottom, left, right.
// Isi ke kanan, lalu ke bawah, lalu ke kiri, lalu ke atas. Ulangi sambil menyempitkan batas.
func SpiralMatrix(n int) [][]int {
	// TODO: implementasi di sini
	table := make([][]int, n)
	return table
} // aku nyerah ini karena pusing saya

// ValidSudoku mengecek apakah board Sudoku 9x9 valid.
// Board direpresentasikan sebagai [][]int 9x9. Angka 0 berarti cell kosong.
// Aturan validasi:
//  1. Setiap baris harus berisi angka 1-9 tanpa duplikat (abaikan 0).
//  2. Setiap kolom harus berisi angka 1-9 tanpa duplikat (abaikan 0).
//  3. Setiap sub-box 3x3 harus berisi angka 1-9 tanpa duplikat (abaikan 0).
//
// Board tidak harus terisi penuh (boleh ada angka 0).
// Contoh:
//
//	ValidSudoku(board_valid) -> true
//	ValidSudoku(board_duplikat_baris) -> false
//	ValidSudoku(board_duplikat_kolom) -> false
//	ValidSudoku(board_duplikat_box) -> false
//
// Hint: cek duplikat di setiap baris, kolom, dan sub-box 3x3 menggunakan map atau set.
func ValidSudoku(board [][]int) bool {
	// TODO: implementasi di sini
	return false
} // aku nyerah ini karena pusing saya
