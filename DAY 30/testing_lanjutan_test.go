package belajar

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

// ═══════════════════════════════════════════════════════════════════════════════
// 🛠 TEST HELPER — Fungsi bantu asertif (selalu panggil t.Helper()!)
// ═══════════════════════════════════════════════════════════════════════════════

// samakanInt memeriksa bahwa `dapat` == `harap`.
// Error menunjuk ke baris pemanggil, bukan ke dalam fungsi ini.
func samakanInt(t *testing.T, label string, dapat, harap int) {
	t.Helper()
	if dapat != harap {
		t.Errorf("❌ %s: dapat %d, harusnya %d", label, dapat, harap)
	}
}

// samakanFloat64 memeriksa bahwa `dapat` == `harap`.
func samakanFloat64(t *testing.T, label string, dapat, harap float64) {
	t.Helper()
	if dapat != harap {
		t.Errorf("❌ %s: dapat %.2f, harusnya %.2f", label, dapat, harap)
	}
}

// samakanString memeriksa bahwa `dapat` == `harap`.
func samakanString(t *testing.T, label string, dapat, harap string) {
	t.Helper()
	if dapat != harap {
		t.Errorf("❌ %s: dapat %q, harusnya %q", label, dapat, harap)
	}
}

// samakanBool memeriksa bahwa `dapat` == `harap`.
func samakanBool(t *testing.T, label string, dapat, harap bool) {
	t.Helper()
	if dapat != harap {
		t.Errorf("❌ %s: dapat %v, harusnya %v", label, dapat, harap)
	}
}

// tidakError memeriksa bahwa err == nil.
func tidakError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("❌ error tidak terduga: %v", err)
	}
}

// adaError memeriksa bahwa err != nil.
func adaError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("❌ harusnya ada error, tapi nil")
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 1: TABLE-DRIVEN TEST — FUNGSI MATEMATIKA
// ═══════════════════════════════════════════════════════════════════════════════

func TestKuadrat(t *testing.T) {
	tests := []struct {
		nama   string
		masukan int
		harap   int
	}{
		{"nol", 0, 0},
		{"satu", 1, 1},
		{"positif", 5, 25},
		{"negatif", -3, 9},
		{"besar", 12, 144},
	}
	for _, tc := range tests {
		t.Run(tc.nama, func(t *testing.T) {
			hasil := Kuadrat(tc.masukan)
			samakanInt(t, fmt.Sprintf("Kuadrat(%d)", tc.masukan), hasil, tc.harap)
		})
	}
	t.Log("✅ Kuadrat: semua kasus benar")
}

func TestAbs(t *testing.T) {
	tests := []struct {
		nama   string
		masukan int
		harap   int
	}{
		{"nol", 0, 0},
		{"positif", 7, 7},
		{"negatif", -7, 7},
		{"negatif-besar", -100, 100},
	}
	for _, tc := range tests {
		t.Run(tc.nama, func(t *testing.T) {
			hasil := Abs(tc.masukan)
			samakanInt(t, fmt.Sprintf("Abs(%d)", tc.masukan), hasil, tc.harap)
		})
	}
	t.Log("✅ Abs: semua kasus benar")
}

func TestMaks(t *testing.T) {
	t.Run("slice normal", func(t *testing.T) {
		hasil, err := Maks([]int{3, 1, 4, 1, 5, 9, 2, 6})
		tidakError(t, err)
		samakanInt(t, "Maks", hasil, 9)
	})
	t.Run("satu elemen", func(t *testing.T) {
		hasil, err := Maks([]int{42})
		tidakError(t, err)
		samakanInt(t, "Maks satu elemen", hasil, 42)
	})
	t.Run("semua sama", func(t *testing.T) {
		hasil, err := Maks([]int{7, 7, 7})
		tidakError(t, err)
		samakanInt(t, "Maks semua sama", hasil, 7)
	})
	t.Run("dengan negatif", func(t *testing.T) {
		hasil, err := Maks([]int{-5, -2, -8})
		tidakError(t, err)
		samakanInt(t, "Maks negatif", hasil, -2)
	})
	t.Run("slice kosong", func(t *testing.T) {
		_, err := Maks([]int{})
		adaError(t, err)
		t.Log("✅ Maks slice kosong menghasilkan error")
	})
	t.Log("✅ Maks: semua kasus benar")
}

func TestMin(t *testing.T) {
	t.Run("slice normal", func(t *testing.T) {
		hasil, err := Min([]int{3, 1, 4, 1, 5, 9, 2, 6})
		tidakError(t, err)
		samakanInt(t, "Min", hasil, 1)
	})
	t.Run("satu elemen", func(t *testing.T) {
		hasil, err := Min([]int{-10})
		tidakError(t, err)
		samakanInt(t, "Min satu elemen", hasil, -10)
	})
	t.Run("slice kosong", func(t *testing.T) {
		_, err := Min([]int{})
		adaError(t, err)
		t.Log("✅ Min slice kosong menghasilkan error")
	})
	t.Log("✅ Min: semua kasus benar")
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 2: SUBTEST t.Run — MANIPULASI STRING
// ═══════════════════════════════════════════════════════════════════════════════

func TestBalikString(t *testing.T) {
	tests := []struct {
		masukan string
		harap   string
	}{
		{"golang", "gnalog"},
		{"a", "a"},
		{"", ""},
		{"ab", "ba"},
		{"racecar", "racecar"},
		{"12345", "54321"},
	}
	for _, tc := range tests {
		nama := fmt.Sprintf("masukan=%q", tc.masukan)
		t.Run(nama, func(t *testing.T) {
			hasil := BalikString(tc.masukan)
			samakanString(t, "BalikString", hasil, tc.harap)
		})
	}
	t.Log("✅ BalikString: semua kasus benar")
}

func TestHitungKata(t *testing.T) {
	tests := []struct {
		nama   string
		masukan string
		harap   int
	}{
		{"kosong", "", 0},
		{"satu kata", "halo", 1},
		{"dua kata", "hello world", 2},
		{"spasi banyak", "  spasi  banyak  ", 2},
		{"tiga kata", "go is fun", 3},
		{"hanya spasi", "   ", 0},
	}
	for _, tc := range tests {
		t.Run(tc.nama, func(t *testing.T) {
			hasil := HitungKata(tc.masukan)
			samakanInt(t, fmt.Sprintf("HitungKata(%q)", tc.masukan), hasil, tc.harap)
		})
	}
	t.Log("✅ HitungKata: semua kasus benar")
}

func TestPalindrom(t *testing.T) {
	tests := []struct {
		nama   string
		masukan string
		harap   bool
	}{
		{"racecar", "racecar", true},
		{"kasak", "kasak", true},
		{"satu huruf", "a", true},
		{"kosong", "", true},
		{"bukan palindrom", "hello", false},
		{"Go kapital", "Go", false},
		{"case insensitive", "Racecar", true},
		{"angka", "12321", true},
		{"angka bukan palindrom", "12345", false},
	}
	for _, tc := range tests {
		t.Run(tc.nama, func(t *testing.T) {
			hasil := Palindrom(tc.masukan)
			samakanBool(t, fmt.Sprintf("Palindrom(%q)", tc.masukan), hasil, tc.harap)
		})
	}
	t.Log("✅ Palindrom: semua kasus benar")
}

func TestKapitalisasiKata(t *testing.T) {
	tests := []struct {
		masukan string
		harap   string
	}{
		{"hello world", "Hello World"},
		{"go is fun", "Go Is Fun"},
		{"", ""},
		{"a", "A"},
		{"satu", "Satu"},
	}
	for _, tc := range tests {
		nama := fmt.Sprintf("masukan=%q", tc.masukan)
		t.Run(nama, func(t *testing.T) {
			hasil := KapitalisasiKata(tc.masukan)
			samakanString(t, "KapitalisasiKata", hasil, tc.harap)
		})
	}
	t.Log("✅ KapitalisasiKata: semua kasus benar")
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 3: TEST HELPER — KERANJANG BELANJA
// ═══════════════════════════════════════════════════════════════════════════════

func buatKeranjangIsi(t *testing.T) *Keranjang {
	t.Helper()
	k := &Keranjang{}
	tidakError(t, k.Tambah(Produk{"Apel", 5000}))
	tidakError(t, k.Tambah(Produk{"Roti", 3000}))
	tidakError(t, k.Tambah(Produk{"Susu", 8000}))
	return k
}

func TestKeranjang_Tambah(t *testing.T) {
	t.Run("produk valid", func(t *testing.T) {
		k := &Keranjang{}
		err := k.Tambah(Produk{"Apel", 5000})
		tidakError(t, err)
		samakanInt(t, "jumlah setelah tambah", k.Jumlah(), 1)
	})
	t.Run("nama kosong", func(t *testing.T) {
		k := &Keranjang{}
		err := k.Tambah(Produk{"", 5000})
		adaError(t, err)
		t.Logf("✅ error nama kosong: %v", err)
	})
	t.Run("harga negatif", func(t *testing.T) {
		k := &Keranjang{}
		err := k.Tambah(Produk{"Apel", -100})
		adaError(t, err)
		t.Logf("✅ error harga negatif: %v", err)
	})
	t.Run("harga nol boleh", func(t *testing.T) {
		k := &Keranjang{}
		err := k.Tambah(Produk{"Gratis", 0})
		tidakError(t, err)
	})
	t.Log("✅ Keranjang.Tambah: semua kasus benar")
}

func TestKeranjang_Hapus(t *testing.T) {
	t.Run("hapus ada", func(t *testing.T) {
		k := buatKeranjangIsi(t)
		err := k.Hapus("Roti")
		tidakError(t, err)
		samakanInt(t, "jumlah setelah hapus", k.Jumlah(), 2)
	})
	t.Run("hapus tidak ada", func(t *testing.T) {
		k := buatKeranjangIsi(t)
		err := k.Hapus("Mangga")
		adaError(t, err)
		if !errors.Is(err, ErrTidakDitemukan) {
			t.Errorf("❌ harusnya ErrTidakDitemukan, dapat: %v", err)
		}
	})
	t.Log("✅ Keranjang.Hapus: semua kasus benar")
}

func TestKeranjang_TotalHarga(t *testing.T) {
	k := buatKeranjangIsi(t) // Apel 5000 + Roti 3000 + Susu 8000 = 16000
	samakanFloat64(t, "TotalHarga", k.TotalHarga(), 16000)
	t.Log("✅ Keranjang.TotalHarga benar")
}

func TestKeranjang_DaftarNama(t *testing.T) {
	k := buatKeranjangIsi(t)
	daftar := k.DaftarNama()

	harap := []string{"Apel", "Roti", "Susu"}
	if len(daftar) != len(harap) {
		t.Fatalf("❌ panjang %d, harusnya %d", len(daftar), len(harap))
	}
	for i, v := range harap {
		samakanString(t, fmt.Sprintf("DaftarNama[%d]", i), daftar[i], v)
	}
	t.Logf("✅ DaftarNama: %v", daftar)
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 4: INTERFACE & MOCK
// ═══════════════════════════════════════════════════════════════════════════════

func buatLayanan() (*LayananPengguna, *MockPenyimpan) {
	mock := &MockPenyimpan{}
	layanan := BuatLayananPengguna(mock)
	return layanan, mock
}

func TestMockPenyimpan(t *testing.T) {
	mock := &MockPenyimpan{}

	t.Run("simpan dan ambil", func(t *testing.T) {
		err := mock.Simpan("k1", "nilai satu")
		tidakError(t, err)

		val, err := mock.Ambil("k1")
		tidakError(t, err)
		samakanString(t, "Ambil k1", val, "nilai satu")
	})

	t.Run("ambil tidak ada", func(t *testing.T) {
		_, err := mock.Ambil("tidak-ada")
		adaError(t, err)
		if !errors.Is(err, ErrTidakDitemukan) {
			t.Errorf("❌ harusnya ErrTidakDitemukan, dapat: %v", err)
		}
	})

	t.Run("ada dan tidak ada", func(t *testing.T) {
		mock.Simpan("ada", "ya")
		samakanBool(t, "Ada(ada)", mock.Ada("ada"), true)
		samakanBool(t, "Ada(kosong)", mock.Ada("tidak-ada-sama-sekali"), false)
	})

	t.Run("hapus", func(t *testing.T) {
		mock.Simpan("hapus-ini", "data")
		err := mock.Hapus("hapus-ini")
		tidakError(t, err)
		samakanBool(t, "setelah hapus", mock.Ada("hapus-ini"), false)
	})

	t.Run("hapus tidak ada", func(t *testing.T) {
		err := mock.Hapus("tidak-pernah-ada")
		adaError(t, err)
	})

	t.Log("✅ MockPenyimpan: semua operasi CRUD benar")
}

func TestLayananPengguna_Daftar(t *testing.T) {
	t.Run("daftar baru", func(t *testing.T) {
		l, _ := buatLayanan()
		err := l.Daftar("u1", "Galang")
		tidakError(t, err)
	})
	t.Run("id duplikat", func(t *testing.T) {
		l, _ := buatLayanan()
		tidakError(t, l.Daftar("u1", "Galang"))
		err := l.Daftar("u1", "Budi")
		adaError(t, err)
		t.Logf("✅ error duplikat: %v", err)
	})
	t.Run("nama kosong", func(t *testing.T) {
		l, _ := buatLayanan()
		err := l.Daftar("u2", "")
		adaError(t, err)
		t.Logf("✅ error nama kosong: %v", err)
	})
	t.Log("✅ LayananPengguna.Daftar: semua kasus benar")
}

func TestLayananPengguna_AmbilNama(t *testing.T) {
	l, _ := buatLayanan()
	tidakError(t, l.Daftar("u1", "Galang"))

	t.Run("pengguna ada", func(t *testing.T) {
		nama, err := l.AmbilNama("u1")
		tidakError(t, err)
		samakanString(t, "AmbilNama", nama, "Galang")
	})
	t.Run("pengguna tidak ada", func(t *testing.T) {
		_, err := l.AmbilNama("x99")
		adaError(t, err)
	})
	t.Log("✅ LayananPengguna.AmbilNama: semua kasus benar")
}

func TestLayananPengguna_Hapus(t *testing.T) {
	t.Run("hapus ada", func(t *testing.T) {
		l, _ := buatLayanan()
		tidakError(t, l.Daftar("u1", "Galang"))
		tidakError(t, l.Hapus("u1"))
		_, err := l.AmbilNama("u1")
		adaError(t, err)
	})
	t.Run("hapus tidak ada", func(t *testing.T) {
		l, _ := buatLayanan()
		err := l.Hapus("tidak-ada")
		adaError(t, err)
	})
	t.Log("✅ LayananPengguna.Hapus: semua kasus benar")
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 5: BENCHMARK
// ═══════════════════════════════════════════════════════════════════════════════

// Pengujian kebenaran GabungString (bukan benchmark, tapi correctness test)
func TestGabungString(t *testing.T) {
	parts := []string{"Go", " ", "is", " ", "awesome"}
	harap := "Go is awesome"

	t.Run("loop", func(t *testing.T) {
		hasil := GabungStringLoop(parts)
		samakanString(t, "GabungStringLoop", hasil, harap)
	})
	t.Run("builder", func(t *testing.T) {
		hasil := GabungStringBuilder(parts)
		samakanString(t, "GabungStringBuilder", hasil, harap)
	})
	t.Log("✅ GabungString: kedua implementasi menghasilkan output sama")
}

func TestUrutkanSalinan(t *testing.T) {
	asli := []int{3, 1, 4, 1, 5, 9, 2, 6}
	asliSalinan := make([]int, len(asli))
	copy(asliSalinan, asli)

	terurut := UrutkanSalinan(asli)

	// asli tidak boleh berubah
	for i, v := range asliSalinan {
		if asli[i] != v {
			t.Errorf("❌ slice asli berubah pada indeks %d: %d → %d", i, v, asli[i])
		}
	}

	harap := []int{1, 1, 2, 3, 4, 5, 6, 9}
	if len(terurut) != len(harap) {
		t.Fatalf("❌ panjang %d, harusnya %d", len(terurut), len(harap))
	}
	for i, v := range harap {
		samakanInt(t, fmt.Sprintf("terurut[%d]", i), terurut[i], v)
	}
	t.Log("✅ UrutkanSalinan: slice asli tidak berubah, hasil terurut benar")
}

// Benchmark: bandingkan GabungStringLoop vs GabungStringBuilder
func BenchmarkGabungStringLoop(b *testing.B) {
	parts := make([]string, 100)
	for i := range parts {
		parts[i] = "golang"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GabungStringLoop(parts)
	}
}

func BenchmarkGabungStringBuilder(b *testing.B) {
	parts := make([]string, 100)
	for i := range parts {
		parts[i] = "golang"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GabungStringBuilder(parts)
	}
}

func BenchmarkUrutkanSalinan(b *testing.B) {
	nums := []int{9, 3, 7, 1, 5, 8, 2, 6, 4, 0}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UrutkanSalinan(nums)
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 6: PENANGANAN ERROR TERSTRUKTUR
// ═══════════════════════════════════════════════════════════════════════════════

func TestErrMasukan_Error(t *testing.T) {
	err := &ErrMasukan{Field: "nama", Pesan: "tidak boleh kosong"}
	harap := "validasi gagal pada field 'nama': tidak boleh kosong"
	samakanString(t, "ErrMasukan.Error()", err.Error(), harap)
	t.Log("✅ ErrMasukan.Error() format benar")
}

func TestValidasiUsia(t *testing.T) {
	tests := []struct {
		nama    string
		usia    int
		adaErr  bool
	}{
		{"valid-muda", 1, false},
		{"valid-normal", 25, false},
		{"valid-tua", 150, false},
		{"terlalu-muda", 0, true},
		{"negatif", -1, true},
		{"terlalu-tua", 151, true},
	}
	for _, tc := range tests {
		t.Run(tc.nama, func(t *testing.T) {
			err := ValidasiUsia(tc.usia)
			if tc.adaErr {
				adaError(t, err)
				var errMasukan *ErrMasukan
				if !errors.As(err, &errMasukan) {
					t.Errorf("❌ tipe error harusnya *ErrMasukan, dapat: %T", err)
				} else {
					samakanString(t, "Field", errMasukan.Field, "usia")
				}
			} else {
				tidakError(t, err)
			}
		})
	}
	t.Log("✅ ValidasiUsia: semua kasus benar")
}

func TestValidasiEmail(t *testing.T) {
	tests := []struct {
		nama   string
		email  string
		adaErr bool
	}{
		{"valid", "user@example.com", false},
		{"valid-subdomain", "a@b.co", false},
		{"kosong", "", true},
		{"tanpa-at", "tidakvalid.com", true},
		{"tanpa-titik", "user@domain", true},
		{"hanya-at", "@", true},
	}
	for _, tc := range tests {
		t.Run(tc.nama, func(t *testing.T) {
			err := ValidasiEmail(tc.email)
			if tc.adaErr {
				adaError(t, err)
				var errMasukan *ErrMasukan
				if !errors.As(err, &errMasukan) {
					t.Errorf("❌ tipe error harusnya *ErrMasukan, dapat: %T", err)
				} else {
					samakanString(t, "Field", errMasukan.Field, "email")
				}
			} else {
				tidakError(t, err)
			}
		})
	}
	t.Log("✅ ValidasiEmail: semua kasus benar")
}

func TestValidasiFormulir(t *testing.T) {
	t.Run("semua valid", func(t *testing.T) {
		errs := ValidasiFormulir(25, "user@example.com")
		if errs != nil {
			t.Errorf("❌ harusnya nil, dapat: %v", errs)
		}
	})
	t.Run("semua invalid", func(t *testing.T) {
		errs := ValidasiFormulir(-1, "tidakvalid")
		if len(errs) != 2 {
			t.Errorf("❌ harusnya 2 error, dapat %d: %v", len(errs), errs)
		}
	})
	t.Run("hanya usia invalid", func(t *testing.T) {
		errs := ValidasiFormulir(200, "ok@ok.com")
		if len(errs) != 1 {
			t.Errorf("❌ harusnya 1 error, dapat %d", len(errs))
		}
	})
	t.Run("hanya email invalid", func(t *testing.T) {
		errs := ValidasiFormulir(25, "tidakvalid")
		if len(errs) != 1 {
			t.Errorf("❌ harusnya 1 error, dapat %d", len(errs))
		}
	})
	t.Log("✅ ValidasiFormulir: semua kasus benar")
}

// ─── pastikan import tidak dilaporkan unused ───────────────────────────────────
var (
	_ = strings.Builder{}
	_ = fmt.Sprintf
)
