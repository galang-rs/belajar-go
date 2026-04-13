package belajar

import (
	"reflect"
	"testing"
)

// =============================================================
// 1. TABLE-DRIVEN TEST - ParseCSV
// =============================================================

func TestParseCSV(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected [][]string
	}{
		{
			"standar 3 baris",
			"nama,umur,kota\nAndi,17,Jakarta\nBudi,16,Bandung",
			[][]string{
				{"nama", "umur", "kota"},
				{"Andi", "17", "Jakarta"},
				{"Budi", "16", "Bandung"},
			},
		},
		{
			"dengan spasi",
			"a, b , c",
			[][]string{{"a", "b", "c"}},
		},
		{
			"kosong",
			"",
			[][]string{},
		},
		{
			"satu cell",
			"hello",
			[][]string{{"hello"}},
		},
		{
			"baris kosong diabaikan",
			"a,b\n\nc,d",
			[][]string{{"a", "b"}, {"c", "d"}},
		},
		{
			"satu kolom banyak baris",
			"alpha\nbeta\ngamma",
			[][]string{{"alpha"}, {"beta"}, {"gamma"}},
		},
		{
			"spasi banyak",
			"  x  ,  y  ,  z  ",
			[][]string{{"x", "y", "z"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseCSV(tt.input)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ParseCSV(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 2. TABLE-DRIVEN TEST - ToCSV
// =============================================================

func TestToCSV(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]string
		expected string
	}{
		{
			"standar 2 baris",
			[][]string{{"nama", "umur", "kota"}, {"Andi", "17", "Jakarta"}},
			"nama,umur,kota\nAndi,17,Jakarta",
		},
		{
			"satu cell",
			[][]string{{"hello"}},
			"hello",
		},
		{
			"kosong",
			[][]string{},
			"",
		},
		{
			"nil",
			nil,
			"",
		},
		{
			"satu kolom",
			[][]string{{"a"}, {"b"}, {"c"}},
			"a\nb\nc",
		},
		{
			"3x3",
			[][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
			"1,2,3\n4,5,6\n7,8,9",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToCSV(tt.input)
			if result != tt.expected {
				t.Errorf("ToCSV(%v) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 3. INTEGRATION TEST - Round Trip
// =============================================================

func TestCSV_RoundTrip(t *testing.T) {
	original := [][]string{
		{"nama", "umur", "kota"},
		{"Andi", "17", "Jakarta"},
		{"Budi", "16", "Bandung"},
	}
	csv := ToCSV(original)
	parsed := ParseCSV(csv)
	if !reflect.DeepEqual(parsed, original) {
		t.Errorf("Round trip gagal: ToCSV -> ParseCSV = %v; want %v", parsed, original)
	}
}

// =============================================================
// 4. BENCHMARK
// =============================================================

func BenchmarkParseCSV(b *testing.B) {
	input := "nama,umur,kota\nAndi,17,Jakarta\nBudi,16,Bandung\nCici,18,Surabaya"
	for i := 0; i < b.N; i++ {
		ParseCSV(input)
	}
}

func BenchmarkToCSV(b *testing.B) {
	data := [][]string{{"nama", "umur"}, {"Andi", "17"}, {"Budi", "16"}}
	for i := 0; i < b.N; i++ {
		ToCSV(data)
	}
}
