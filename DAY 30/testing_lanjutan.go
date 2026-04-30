package belajar

// ==================== DAY 30: TESTING LANJUTAN — TABLE-DRIVEN, SUBTEST, BENCHMARK & MOCK ====================
//
// 🎯 FOKUS HARI INI:
//   Go memiliki ekosistem testing yang kuat dan idiomatis. Hari ini kita eksplorasi:
//     - Table-driven tests  → pola idiomatis pengujian banyak skenario sekaligus
//     - Subtests (t.Run)    → mengelompokkan & memberi nama setiap kasus uji
//     - Test helpers        → fungsi bantu yang melaporkan posisi pemanggil
//     - Benchmarks          → mengukur performa fungsi secara otomatis
//     - Interface & mock    → menguji perilaku tanpa ketergantungan eksternal
//
//   Jalankan test:
//     cd "DAY 30"
//     go test ./... -v
//
//   Jalankan benchmark:
//     go test -bench=. -benchmem
//
//   Jalankan dengan race detector:
//     go test ./... -v -race
//
// ═══════════════════════════════════════════════════════════════════════════════
// 🧠 KAPAN PAKAI APA?
// ═══════════════════════════════════════════════════════════════════════════════
//
//   ┌───────────────────────┬──────────────────────────────────────────────────┐
//   │ Teknik                │ Gunakan ketika...                                │
//   ├───────────────────────┼──────────────────────────────────────────────────┤
//   │ Table-driven test     │ banyak input/output berbeda untuk satu fungsi   │
//   │ t.Run (subtest)       │ mengelompokkan kasus & bisa di-filter            │
//   │ t.Helper()            │ fungsi bantu agar error menunjuk baris pemanggil │
//   │ Benchmark             │ mengukur dan membandingkan performa              │
//   │ Interface/mock        │ mengisolasi ketergantungan eksternal (DB, HTTP)  │
//   └───────────────────────┴──────────────────────────────────────────────────┘
//
//   Aturan emas:
//     1. Setiap test hanya menguji SATU hal.
//     2. Nama test deskriptif: Test<Fungsi>_<Skenario>
//     3. Gunakan t.Helper() di setiap fungsi bantu.
//     4. Benchmark harus idempoten — hasil tidak bergantung pada urutan.
//
// ═══════════════════════════════════════════════════════════════════════════════

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strings"
	"unicode"
)

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 1: FUNGSI MATEMATIKA DASAR (bahan untuk table-driven test)
// ═══════════════════════════════════════════════════════════════════════════════
//
// Fungsi-fungsi berikut akan diuji dengan pola table-driven test.
// Pola table-driven sangat populer di Go karena:
//   - Satu fungsi test bisa menguji puluhan skenario
//   - Mudah ditambah kasus baru tanpa mengubah logika test
//   - Output error langsung menunjuk nama kasus yang gagal
//
// Pola dasar:
//
//	tests := []struct {
//	    nama   string
//	    masukan int
//	    harap   int
//	}{
//	    {"nol", 0, 0},
//	    {"positif", 5, 25},
//	}
//	for _, tc := range tests {
//	    t.Run(tc.nama, func(t *testing.T) {
//	        hasil := Kuadrat(tc.masukan)
//	        if hasil != tc.harap { t.Errorf(...) }
//	    })
//	}
//

// Kuadrat mengembalikan n * n.
//
// Contoh:
//
//	Kuadrat(4)  // 16
//	Kuadrat(-3) // 9
func Kuadrat(n int) int {
	// TODO: implementasi di sini
	return n * n
}

// Abs mengembalikan nilai absolut dari n.
//
// Contoh:
//
//	Abs(-7) // 7
//	Abs(5)  // 5
//	Abs(0)  // 0
func Abs(n int) int {
	// TODO: implementasi di sini
	return int(math.Abs(float64(n)))
}

// Maks mengembalikan nilai terbesar dari slice.
// Mengembalikan error jika slice kosong.
//
// Contoh:
//
//	Maks([]int{3, 1, 4, 1, 5}) // 5, nil
//	Maks([]int{})               // 0, error
func Maks(nums []int) (int, error) {
	// TODO: implementasi di sini
	if len(nums) == 0 {
		return 0, fmt.Errorf("error")
	}
	val := math.Inf(-1)
	for _, v := range nums {
		if float64(v) > val {
			val = float64(v)
		}
	}
	v := int(val)
	return v, nil
}

// Min mengembalikan nilai terkecil dari slice.
// Mengembalikan error jika slice kosong.
//
// Contoh:
//
//	Min([]int{3, 1, 4, 1, 5}) // 1, nil
//	Min([]int{})               // 0, error
func Min(nums []int) (int, error) {
	// TODO: implementasi di sini
	if len(nums) == 0 {
		return 0, fmt.Errorf("error")
	}
	val := math.Inf(1)
	for _, v := range nums {
		if float64(v) < val {
			val = float64(v)
		}
	}
	v := int(val)
	return v, nil
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 2: MANIPULASI STRING (bahan untuk subtest t.Run)
// ═══════════════════════════════════════════════════════════════════════════════
//
// t.Run memungkinkan subtest: test bersarang dengan nama yang bisa di-filter.
//
// Menjalankan subtest tertentu:
//   go test -run TestNamaFungsi/nama_subtest
//
// Mengapa t.Run penting?
//   - Setiap subtest punya scope tersendiri (t.Fatal hanya memengaruhi subtest itu)
//   - Mudah dijalankan secara paralel dengan t.Parallel()
//   - Output lebih informatif: "--- FAIL: TestX/kasus_gagal"
//

// BalikString membalik urutan karakter dalam string.
//
// Contoh:
//
//	BalikString("golang") // "gnalog"
//	BalikString("")       // ""
//	BalikString("a")      // "a"
func BalikString(s string) string {
	// TODO: implementasi di sini
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}

// HitungKata mengembalikan jumlah kata dalam string.
// Kata dipisahkan oleh spasi (satu atau lebih).
//
// Contoh:
//
//	HitungKata("hello world")     // 2
//	HitungKata("  spasi  banyak") // 2
//	HitungKata("")                // 0
func HitungKata(s string) int {
	// TODO: implementasi di sini
	s = strings.ReplaceAll(s, "  ", " ")
	s = strings.TrimSpace(s)
	val := strings.Split(s, " ")
	if val[0] == "" {
		return 0
	}

	return len(val)
}

// Palindrom mengembalikan true jika s adalah palindrom (dibaca sama dari dua arah).
// Perbandingan case-insensitive. Karakter non-huruf diabaikan.
//
// Contoh:
//
//	Palindrom("racecar")       // true
//	Palindrom("A man a plan")  // false (karena "amanap lan" bukan palindrom)
//	Palindrom("kasak")         // true
//	Palindrom("Go")            // false
func Palindrom(s string) bool {
	// TODO: implementasi di sini
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)
	for k, v := range s {
		if v != rune(s[len(s)-1-k]) {
			return false
		}
	}
	return true
}

// KapitalisasiKata mengubah huruf pertama setiap kata menjadi kapital.
//
// Contoh:
//
//	KapitalisasiKata("hello world") // "Hello World"
//	KapitalisasiKata("go is fun")   // "Go Is Fun"
//	KapitalisasiKata("")            // ""
func KapitalisasiKata(s string) string {
	// TODO: implementasi di sini
	return s
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 3: KERANJANG BELANJA (bahan untuk test helper & interface)
// ═══════════════════════════════════════════════════════════════════════════════
//
// Test helper adalah fungsi bantu yang memanggil t.Helper() agar:
//   - Error melaporkan baris pemanggil (bukan baris di dalam helper)
//   - Kode test lebih bersih dan tidak duplikat logika assersi
//
// Pola test helper:
//
//	func samakanInt(t *testing.T, dapat, harap int) {
//	    t.Helper() // WAJIB — agar baris error menunjuk ke pemanggil
//	    if dapat != harap {
//	        t.Errorf("dapat %d, harusnya %d", dapat, harap)
//	    }
//	}
//

// Produk merepresentasikan item di keranjang belanja.
type Produk struct {
	Nama  string
	Harga float64
}

// Keranjang adalah kumpulan produk yang bisa dikelola.
type Keranjang struct {
	items []Produk
}

// Tambah menambahkan produk ke keranjang.
// Mengembalikan error jika harga negatif atau nama kosong.
//
// Contoh:
//
//	k := &Keranjang{}
//	err := k.Tambah(Produk{"Apel", 5000})  // nil
//	err  = k.Tambah(Produk{"", 1000})      // error: nama kosong
//	err  = k.Tambah(Produk{"Roti", -100})  // error: harga negatif
func (k *Keranjang) Tambah(p Produk) error {
	// TODO: implementasi di sini
	if p.Nama == "" {
		return fmt.Errorf("nama kosong")
	} else if p.Harga < 0 {
		return fmt.Errorf("harga negatif")
	}
	k.items = append(k.items, p)
	return nil
}

// Hapus menghapus produk pertama dengan nama yang cocok.
// Mengembalikan error jika produk tidak ditemukan.
//
// Contoh:
//
//	k.Tambah(Produk{"Apel", 5000})
//	err := k.Hapus("Apel")   // nil
//	err  = k.Hapus("Mangga") // error: tidak ditemukan
func (k *Keranjang) Hapus(nama string) error {
	// TODO: implementasi di sini
	arr := []Produk{}
	found := false
	for _, v := range k.items {
		if v.Nama != nama {
			arr = append(arr, v)
		} else {
			found = true
		}
	}
	if !found {
		return ErrTidakDitemukan
	}
	k.items = arr
	return nil
}

// TotalHarga menghitung total harga semua produk di keranjang.
//
// Contoh:
//
//	k.Tambah(Produk{"Apel", 5000})
//	k.Tambah(Produk{"Roti", 3000})
//	k.TotalHarga() // 8000
func (k *Keranjang) TotalHarga() float64 {
	// TODO: implementasi di sini
	var count float64
	for _, v := range k.items {
		count += v.Harga
	}
	return count
}

// Jumlah mengembalikan jumlah produk di keranjang.
func (k *Keranjang) Jumlah() int {
	// TODO: implementasi di sini
	return len(k.items)
}

// DaftarNama mengembalikan slice nama semua produk (terurut alfabet).
//
// Contoh:
//
//	k.Tambah(Produk{"Roti", 3000})
//	k.Tambah(Produk{"Apel", 5000})
//	k.DaftarNama() // ["Apel", "Roti"]
func (k *Keranjang) DaftarNama() []string {
	// TODO: implementasi di sini
	arr := []string{}

	for _, v := range k.items {
		arr = append(arr, v.Nama)
	}

	return arr
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 4: INTERFACE & MOCK — MENGUJI TANPA KETERGANTUNGAN EKSTERNAL
// ═══════════════════════════════════════════════════════════════════════════════
//
// Mock adalah implementasi palsu dari sebuah interface, digunakan dalam testing
// untuk menggantikan komponen eksternal (database, HTTP client, file system).
//
// Kenapa interface penting dalam testing?
//   - Kode bergantung pada interface, bukan implementasi konkret
//   - Saat testing, kita inject implementasi mock
//   - Test menjadi cepat, deterministik, dan tidak butuh infrastruktur nyata
//
// Pola:
//
//   type PenyimpanData interface {
//       Simpan(id string, data string) error
//       Ambil(id string) (string, error)
//   }
//
//   // Implementasi nyata (pakai DB)
//   type DBPenyimpan struct { ... }
//
//   // Implementasi mock (pakai map in-memory)
//   type MockPenyimpan struct { data map[string]string }
//

// PenyimpanData adalah interface untuk menyimpan dan mengambil data.
type PenyimpanData interface {
	Simpan(id string, data string) error
	Ambil(id string) (string, error)
	Hapus(id string) error
	Ada(id string) bool
}

// MockPenyimpan adalah implementasi in-memory dari PenyimpanData untuk keperluan testing.
//
// Contoh:
//
//	mock := &MockPenyimpan{}
//	mock.Simpan("u1", "galang")
//	val, _ := mock.Ambil("u1") // "galang"
type MockPenyimpan struct {
	data map[string]string
}

// Simpan menyimpan data ke map internal.
// Mengembalikan error jika id kosong.
//
// Hint: inisialisasi m.data jika nil, lalu simpan ke map.
func (m *MockPenyimpan) Simpan(id string, data string) error {
	// TODO: implementasi di sini
	if m.data == nil {
		m.data = make(map[string]string)
	}
	if m.data[id] != "" {
		return fmt.Errorf("id ada")
	}
	m.data[id] = data
	return nil
}

// Ambil mengambil data berdasarkan id.
// Mengembalikan error jika id tidak ditemukan.
//
// Hint: cek apakah key ada di map.
func (m *MockPenyimpan) Ambil(id string) (string, error) {
	// TODO: implementasi di sini
	if m.data[id] == "" {
		return "", ErrTidakDitemukan
	}

	return m.data[id], nil
}

// Hapus menghapus data berdasarkan id.
// Mengembalikan error jika id tidak ditemukan.
func (m *MockPenyimpan) Hapus(id string) error {
	// TODO: implementasi di sini
	if m.data[id] == "" {
		return ErrTidakDitemukan
	}

	delete(m.data, id)
	return nil
}

// Ada mengembalikan true jika id ada di penyimpanan.
func (m *MockPenyimpan) Ada(id string) bool {
	// TODO: implementasi di sini
	if m.data[id] == "" {
		return false
	}
	return true
}

// LayananPengguna adalah service yang bergantung pada PenyimpanData.
// Ini adalah contoh dependency injection — layanan tidak tahu apakah
// penyimpanan menggunakan DB sungguhan atau mock.
type LayananPengguna struct {
	penyimpan PenyimpanData
}

// BuatLayananPengguna membuat LayananPengguna dengan penyimpanan yang diberikan.
func BuatLayananPengguna(p PenyimpanData) *LayananPengguna {
	return &LayananPengguna{penyimpan: p}
}

// Daftar mendaftarkan pengguna baru.
// Mengembalikan error jika id sudah terdaftar atau nama kosong.
//
// Contoh:
//
//	l := BuatLayananPengguna(&MockPenyimpan{})
//	err := l.Daftar("u1", "Galang")  // nil
//	err  = l.Daftar("u1", "Budi")    // error: id sudah ada
//	err  = l.Daftar("u2", "")        // error: nama kosong
func (l *LayananPengguna) Daftar(id, nama string) error {
	// TODO: implementasi di sini
	if l.penyimpan.Ada(id) {
		return fmt.Errorf("id sudah ada")
	} else if nama == "" {
		return fmt.Errorf("nama kosong")
	}

	return nil
}

// AmbilNama mengambil nama pengguna berdasarkan id.
// Mengembalikan error jika id tidak ditemukan.
//
// Contoh:
//
//	l.Daftar("u1", "Galang")
//	nama, err := l.AmbilNama("u1")   // "Galang", nil
//	nama, err  = l.AmbilNama("x99")  // "", error
func (l *LayananPengguna) AmbilNama(id string) (string, error) {
	// TODO: implementasi di sini
	panic("belum diimplementasi")
}

// Hapus menghapus pengguna berdasarkan id.
// Mengembalikan error jika id tidak ditemukan.
func (l *LayananPengguna) Hapus(id string) error {
	// TODO: implementasi di sini
	panic("belum diimplementasi")
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 5: FUNGSI BENCHMARK (performa)
// ═══════════════════════════════════════════════════════════════════════════════
//
// Benchmark di Go berjalan dalam loop b.N kali. Go menyesuaikan b.N secara
// otomatis agar benchmark berjalan cukup lama untuk pengukuran akurat.
//
// Pola dasar benchmark:
//
//	func BenchmarkFungsi(b *testing.B) {
//	    for i := 0; i < b.N; i++ {
//	        Fungsi(masukan)
//	    }
//	}
//
// Jalankan: go test -bench=. -benchmem
// Output: BenchmarkFungsi-8   1000000   1200 ns/op   64 B/op   2 allocs/op
//
// Fungsi-fungsi berikut dirancang untuk dibandingkan performanya.
//

// GabungStringLoop menggabungkan slice string menggunakan += di dalam loop.
// Ini adalah implementasi LAMBAT karena string di Go immutable (membuat string baru tiap iterasi).
//
// Contoh:
//
//	GabungStringLoop([]string{"a", "b", "c"}) // "abc"
func GabungStringLoop(parts []string) string {
	// TODO: implementasi di sini
	panic("belum diimplementasi")
}

// GabungStringBuilder menggabungkan slice string menggunakan strings.Builder.
// Ini adalah implementasi CEPAT karena strings.Builder meminimalkan alokasi memori.
//
// Contoh:
//
//	GabungStringBuilder([]string{"a", "b", "c"}) // "abc"
//
// Hint: var sb strings.Builder → sb.WriteString → sb.String()
func GabungStringBuilder(parts []string) string {
	// TODO: implementasi di sini
	panic("belum diimplementasi")
}

// UrutkanSalinan mengurutkan slice int tanpa mengubah slice asli.
// Fungsi ini membuat salinan terlebih dahulu, lalu mengurutkannya.
//
// Contoh:
//
//	asli := []int{3, 1, 4, 1, 5}
//	terurut := UrutkanSalinan(asli)
//	// asli    = [3, 1, 4, 1, 5]  (tidak berubah)
//	// terurut = [1, 1, 3, 4, 5]
//
// Hint: make + copy + sort.Ints
func UrutkanSalinan(nums []int) []int {
	// TODO: implementasi di sini
	panic("belum diimplementasi")
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 6: PENANGANAN ERROR TERSTRUKTUR
// ═══════════════════════════════════════════════════════════════════════════════
//
// Go mendorong error sebagai nilai eksplisit. Untuk error yang kaya informasi,
// kita bisa membuat tipe error khusus yang mengimplementasi interface error.
//
// Antarmuka error standar:
//
//	type error interface {
//	    Error() string
//	}
//
// Untuk memeriksa tipe error spesifik:
//   errors.Is(err, targetErr)  → cocokkan nilai error (untuk sentinel errors)
//   errors.As(err, &target)    → cocokkan tipe error (untuk custom error types)
//

// ErrTidakDitemukan adalah sentinel error untuk kasus "data tidak ada".
var ErrTidakDitemukan = errors.New("tidak ditemukan")

// ErrMasukan adalah tipe error untuk validasi masukan yang tidak valid.
// Menyimpan nama field dan pesan detail.
//
// Contoh:
//
//	err := &ErrMasukan{Field: "nama", Pesan: "tidak boleh kosong"}
//	fmt.Println(err.Error()) // "validasi gagal pada field 'nama': tidak boleh kosong"
type ErrMasukan struct {
	Field string
	Pesan string
}

// Error mengimplementasi interface error.
// Format: "validasi gagal pada field '<Field>': <Pesan>"
func (e *ErrMasukan) Error() string {
	// TODO: implementasi di sini
	panic("belum diimplementasi")
}

// ValidasiUsia memeriksa apakah usia valid (1–150).
// Mengembalikan *ErrMasukan jika tidak valid, nil jika valid.
//
// Contoh:
//
//	ValidasiUsia(25)   // nil
//	ValidasiUsia(-1)   // &ErrMasukan{Field: "usia", ...}
//	ValidasiUsia(200)  // &ErrMasukan{Field: "usia", ...}
func ValidasiUsia(usia int) error {
	// TODO: implementasi di sini
	panic("belum diimplementasi")
}

// ValidasiEmail memeriksa apakah email mengandung karakter '@' dan '.'.
// Mengembalikan *ErrMasukan jika tidak valid, nil jika valid.
//
// Contoh:
//
//	ValidasiEmail("user@example.com") // nil
//	ValidasiEmail("tidakvalid")       // &ErrMasukan{Field: "email", ...}
//	ValidasiEmail("")                 // &ErrMasukan{Field: "email", ...}
func ValidasiEmail(email string) error {
	// TODO: implementasi di sini
	panic("belum diimplementasi")
}

// ValidasiFormulir memvalidasi semua field formulir sekaligus.
// Mengembalikan slice semua error yang ditemukan (bukan hanya pertama).
// Mengembalikan nil jika tidak ada error.
//
// Contoh:
//
//	errs := ValidasiFormulir(-1, "tidakvalid")
//	// len(errs) == 2
//
//	errs = ValidasiFormulir(25, "user@example.com")
//	// errs == nil
func ValidasiFormulir(usia int, email string) []error {
	// TODO: implementasi di sini
	panic("belum diimplementasi")
}

// ─── pastikan import tidak dilaporkan unused ───────────────────────────────────
var (
	_ = fmt.Sprintf
	_ = math.Abs
	_ = sort.Ints
	_ = strings.Builder{}
	_ = unicode.IsLetter
)
