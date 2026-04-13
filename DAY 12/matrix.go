package belajar

// ==================== DAY 12: OPERASI MATRIX ====================
// Topik: Manipulasi matrix (slice 2D) - transpose dan rotasi.
// Konsep penting: cara mengakses dan memanipulasi elemen di array 2 dimensi.

// TransposeMatrix menukar baris menjadi kolom dan sebaliknya dari matrix integer.
// Matrix tidak harus persegi (bisa M x N).
// Contoh: TransposeMatrix([][]int{
//
//	{1, 2, 3},
//	{4, 5, 6},
//
// }) -> [][]int{
//
//	{1, 4},
//	{2, 5},
//	{3, 6},
//
// }
//
//	TransposeMatrix([][]int{{1}}) -> [][]int{{1}}
//	TransposeMatrix([][]int{}) -> [][]int{}
//
// Hint: baris ke-i kolom ke-j di matrix asli menjadi baris ke-j kolom ke-i di matrix hasil.
func TransposeMatrix(matrix [][]int) [][]int {
	// TODO: implementasi di sini
	return nil
}

// RotateMatrix90 memutar matrix persegi 90 derajat searah jarum jam (clockwise).
// Matrix dijamin persegi (N x N).
// Kembalikan matrix baru (TIDAK mengubah matrix asli).
// Contoh: RotateMatrix90([][]int{
//
//	{1, 2, 3},
//	{4, 5, 6},
//	{7, 8, 9},
//
// }) -> [][]int{
//
//	{7, 4, 1},
//	{8, 5, 2},
//	{9, 6, 3},
//
// }
//
//	RotateMatrix90([][]int{{1}}) -> [][]int{{1}}
//	RotateMatrix90([][]int{}) -> [][]int{}
//
// Hint: result[j][n-1-i] = matrix[i][j], atau transpose lalu reverse setiap baris.
func RotateMatrix90(matrix [][]int) [][]int {
	// TODO: implementasi di sini
	return nil
}
