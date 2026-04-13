package belajar

// ==================== DAY 17: ENCODING & PARSING ====================
// Topik: Parsing dan formatting data teks (CSV-like).
// Konsep penting: memecah string menjadi data terstruktur dan sebaliknya.

// ParseCSV memecah string CSV (Comma-Separated Values) menjadi slice 2D string.
// Setiap baris dipisahkan oleh newline ("\n"), setiap kolom dipisahkan oleh koma (",").
// Spasi di awal dan akhir setiap cell harus di-trim (dihapus).
// Baris kosong diabaikan.
// Contoh: ParseCSV("nama,umur,kota\nAndi,17,Jakarta\nBudi,16,Bandung") -> [][]string{
//
//	{"nama", "umur", "kota"},
//	{"Andi", "17", "Jakarta"},
//	{"Budi", "16", "Bandung"},
//
// }
//
//	ParseCSV("a, b , c") -> [][]string{{"a", "b", "c"}}
//	ParseCSV("") -> [][]string{}
//	ParseCSV("hello") -> [][]string{{"hello"}}
//
// Hint: gunakan strings.Split dan strings.TrimSpace
func ParseCSV(input string) [][]string {
	// TODO: implementasi di sini
	return nil
}

// ToCSV mengonversi slice 2D string menjadi string CSV.
// Setiap baris dipisahkan oleh newline ("\n"), setiap kolom dipisahkan oleh koma (",").
// Tidak ada trailing newline di akhir.
// Contoh: ToCSV([][]string{
//
//	{"nama", "umur", "kota"},
//	{"Andi", "17", "Jakarta"},
//
// }) -> "nama,umur,kota\nAndi,17,Jakarta"
//
//	ToCSV([][]string{{"hello"}}) -> "hello"
//	ToCSV([][]string{}) -> ""
//	ToCSV(nil) -> ""
func ToCSV(data [][]string) string {
	// TODO: implementasi di sini
	return ""
}
