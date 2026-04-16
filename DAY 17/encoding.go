package belajar

import (
	"fmt"
	"strings"
)

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

	if len(input) == 0 {
		return [][]string{}
	}
	// TODO: implementasi di sini
	var data [][]string

	trim := strings.Replace(input, "\n\n", "\n", -1)
	trim = strings.Replace(input, "\n\n", "\n", -1)
	val := strings.Split(trim, "\n")

	for _, v := range val {
		va := strings.Split(v, ",")
		var dat []string
		for _, v := range va {
			vat := strings.TrimSpace(v)
			dat = append(dat, vat)
		}
		data = append(data, dat)
	}

	return data
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

	var datas string
	for k1, v1 := range data {
		for k2, v2 := range v1 {
			datas += v2
			if k2 != len(v1)-1 {
				datas += ","
			}
		}
		if k1 != len(data)-1 {
			datas += "\n"
		}
	}
	fmt.Println(datas)
	return datas
}
