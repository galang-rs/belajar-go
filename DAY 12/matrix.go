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
	if len(matrix) == 0 {
		return [][]int{}
	}

	rows := len(matrix)
	cols := len(matrix[0])

	result := make([][]int, cols)
	for i := 0; i < cols; i++ {
		result[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = matrix[i][j]
		}
	}

	return result
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
	if len(matrix) == 0 {
		return [][]int{}
	}

	rows := len(matrix)
	cols := len(matrix[0])

	result := make([][]int, cols)
	for i := 0; i < cols; i++ {
		result[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = matrix[i][j]
		}
	}

	for i := 0; i < len(result); i++ {
		for l, r := 0, len(result[i])-1; l < r; l, r = l+1, r-1 {
			result[i][l], result[i][r] = result[i][r], result[i][l]
		}
	}

	return result
}
