package belajar

import (
	"sort"
	"testing"
	"time"
)

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 1: ASYNC RETURN CHANNEL
// ═══════════════════════════════════════════════════════════════════════════════

func TestKuadratAsync(t *testing.T) {
	t.Log("📦 KuadratAsync: hitung n*n di goroutine, return channel")

	ch := KuadratAsync(7)
	if ch == nil {
		t.Fatal("❌ channel tidak boleh nil")
	}

	// Fungsi harus langsung return (non-blocking)
	mulai := time.Now()
	_ = ch
	if time.Since(mulai) > 10*time.Millisecond {
		t.Error("❌ KuadratAsync harus langsung return, bukan blokir")
	}

	hasil := <-ch
	if hasil != 49 {
		t.Errorf("❌ dapat %d, harusnya 49 (7*7)", hasil)
	}

	t.Logf("✅ KuadratAsync(7) = %d", hasil)
}

func TestKuadratAsync_BerbagaiNilai(t *testing.T) {
	kasus := []struct{ n, harusnya int }{
		{0, 0},
		{1, 1},
		{5, 25},
		{10, 100},
		{12, 144},
	}

	for _, k := range kasus {
		ch := KuadratAsync(k.n)
		hasil := <-ch
		if hasil != k.harusnya {
			t.Errorf("❌ KuadratAsync(%d) = %d, harusnya %d", k.n, hasil, k.harusnya)
		}
	}
	t.Log("✅ KuadratAsync semua kasus benar")
}

func TestJumlahAsync(t *testing.T) {
	t.Log("📦 JumlahAsync: hitung a+b di goroutine, return channel")

	kasus := []struct{ a, b, harusnya int }{
		{3, 4, 7},
		{0, 0, 0},
		{10, -3, 7},
		{100, 200, 300},
	}

	for _, k := range kasus {
		ch := JumlahAsync(k.a, k.b)
		if ch == nil {
			t.Fatalf("❌ JumlahAsync(%d, %d): channel tidak boleh nil", k.a, k.b)
		}
		hasil := <-ch
		if hasil != k.harusnya {
			t.Errorf("❌ JumlahAsync(%d, %d) = %d, harusnya %d", k.a, k.b, hasil, k.harusnya)
		}
	}
	t.Log("✅ JumlahAsync semua kasus benar")
}

func TestFactorialAsync(t *testing.T) {
	t.Log("📦 FactorialAsync: hitung n! di goroutine, return channel")

	kasus := []struct{ n, harusnya int }{
		{0, 1},
		{1, 1},
		{2, 2},
		{5, 120},
		{6, 720},
		{10, 3628800},
	}

	for _, k := range kasus {
		ch := FactorialAsync(k.n)
		if ch == nil {
			t.Fatalf("❌ FactorialAsync(%d): channel tidak boleh nil", k.n)
		}
		hasil := <-ch
		if hasil != k.harusnya {
			t.Errorf("❌ FactorialAsync(%d) = %d, harusnya %d", k.n, hasil, k.harusnya)
		}
	}
	t.Log("✅ FactorialAsync semua kasus benar")
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 2: TRANSFORM CHANNEL → CHANNEL
// ═══════════════════════════════════════════════════════════════════════════════

func TestFilter(t *testing.T) {
	t.Log("📦 Filter: teruskan nilai yang lolos fn saja")

	// Filter genap
	src := GenRange(1, 10)
	genap := Filter(src, func(n int) bool { return n%2 == 0 })

	if genap == nil {
		t.Fatal("❌ channel tidak boleh nil")
	}

	var hasil []int
	for v := range genap {
		hasil = append(hasil, v)
	}

	harusnya := []int{2, 4, 6, 8, 10}
	if len(hasil) != len(harusnya) {
		t.Fatalf("❌ panjang %d, harusnya %d", len(hasil), len(harusnya))
	}
	for i, v := range harusnya {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ Filter(genap): %v", hasil)
}

func TestFilter_TidakAdaYangLolos(t *testing.T) {
	src := GenRange(1, 5)
	// Filter nilai > 100 — tidak ada yang lolos
	out := Filter(src, func(n int) bool { return n > 100 })

	var hasil []int
	for v := range out {
		hasil = append(hasil, v)
	}

	if len(hasil) != 0 {
		t.Errorf("❌ harusnya tidak ada hasil, dapat %v", hasil)
	}
	t.Log("✅ Filter (tidak ada yang lolos): channel kosong dan tertutup")
}

func TestKali(t *testing.T) {
	t.Log("📦 Kali: tiap nilai × faktor")

	src := GenRange(1, 5)
	out := Kali(src, 3)

	if out == nil {
		t.Fatal("❌ channel tidak boleh nil")
	}

	var hasil []int
	for v := range out {
		hasil = append(hasil, v)
	}

	harusnya := []int{3, 6, 9, 12, 15}
	if len(hasil) != len(harusnya) {
		t.Fatalf("❌ panjang %d, harusnya %d", len(hasil), len(harusnya))
	}
	for i, v := range harusnya {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ Kali(src, 3): %v", hasil)
}

func TestKurangi(t *testing.T) {
	t.Log("📦 Kurangi: tiap nilai - n")

	src := GenRange(5, 9)
	out := Kurangi(src, 3)

	if out == nil {
		t.Fatal("❌ channel tidak boleh nil")
	}

	var hasil []int
	for v := range out {
		hasil = append(hasil, v)
	}

	harusnya := []int{2, 3, 4, 5, 6}
	if len(hasil) != len(harusnya) {
		t.Fatalf("❌ panjang %d, harusnya %d", len(hasil), len(harusnya))
	}
	for i, v := range harusnya {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ Kurangi(src, 3): %v", hasil)
}

func TestAmbil(t *testing.T) {
	t.Log("📦 Ambil: ambil n pertama dari channel, return <-chan int")

	src := GenRange(1, 100) // 1..100
	lima := Ambil(src, 5)

	if lima == nil {
		t.Fatal("❌ channel tidak boleh nil")
	}

	var hasil []int
	for v := range lima {
		hasil = append(hasil, v)
	}

	if len(hasil) != 5 {
		t.Fatalf("❌ panjang %d, harusnya 5", len(hasil))
	}
	harusnya := []int{1, 2, 3, 4, 5}
	for i, v := range harusnya {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ Ambil(src, 5): %v", hasil)
}

func TestAmbil_BisaDiPipe(t *testing.T) {
	t.Log("📦 Ambil: hasilnya bisa langsung di-pipe ke Kali")

	// Tanpa konversi manual: ambil 3 pertama → kali 10 → langsung for range
	src := GenRange(1, 100)
	tiga := Ambil(src, 3)
	hasil_pipe := Kali(tiga, 10)

	var hasil []int
	for v := range hasil_pipe {
		hasil = append(hasil, v)
	}

	harusnya := []int{10, 20, 30}
	for i, v := range harusnya {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ Ambil → Kali langsung pipe: %v", hasil)
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 3: GENERATOR
// ═══════════════════════════════════════════════════════════════════════════════

func TestGenRange(t *testing.T) {
	t.Log("📦 GenRange: hasilkan angka dari..sampai (inklusif)")

	var hasil []int
	for v := range GenRange(3, 7) {
		hasil = append(hasil, v)
	}

	harusnya := []int{3, 4, 5, 6, 7}
	if len(hasil) != len(harusnya) {
		t.Fatalf("❌ panjang %d, harusnya %d", len(hasil), len(harusnya))
	}
	for i, v := range harusnya {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ GenRange(3, 7): %v", hasil)
}

func TestGenRange_SatuElemen(t *testing.T) {
	var hasil []int
	for v := range GenRange(5, 5) {
		hasil = append(hasil, v)
	}
	if len(hasil) != 1 || hasil[0] != 5 {
		t.Errorf("❌ GenRange(5,5) harusnya [5], dapat %v", hasil)
	}
	t.Log("✅ GenRange(5, 5) = [5]")
}

func TestGenUlangi(t *testing.T) {
	t.Log("📦 GenUlangi: hasilkan nilai v sebanyak n kali")

	var hasil []int
	for v := range GenUlangi(7, 4) {
		hasil = append(hasil, v)
	}

	if len(hasil) != 4 {
		t.Fatalf("❌ panjang %d, harusnya 4", len(hasil))
	}
	for i, v := range hasil {
		if v != 7 {
			t.Errorf("❌ hasil[%d] = %d, harusnya 7", i, v)
		}
	}
	t.Logf("✅ GenUlangi(7, 4): %v", hasil)
}

func TestGenUlangi_Nol(t *testing.T) {
	var hasil []int
	for v := range GenUlangi(99, 0) {
		hasil = append(hasil, v)
	}
	if len(hasil) != 0 {
		t.Errorf("❌ GenUlangi(99, 0) harusnya kosong, dapat %v", hasil)
	}
	t.Log("✅ GenUlangi(99, 0) = []")
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 4: FAN-OUT
// ═══════════════════════════════════════════════════════════════════════════════

func TestSebarKe(t *testing.T) {
	t.Log("📦 SebarKe: bagi nilai dari src ke n channel (round-robin)")

	src := GenRange(1, 6) // 1,2,3,4,5,6
	chs := SebarKe(src, 3)

	if len(chs) != 3 {
		t.Fatalf("❌ jumlah channel %d, harusnya 3", len(chs))
	}

	// channels[0] → 1, 4
	// channels[1] → 2, 5
	// channels[2] → 3, 6
	harusnya := [][]int{{1, 4}, {2, 5}, {3, 6}}

	for i, ch := range chs {
		var hasil []int
		for v := range ch {
			hasil = append(hasil, v)
		}
		if len(hasil) != len(harusnya[i]) {
			t.Errorf("❌ channels[%d] panjang %d, harusnya %d", i, len(hasil), len(harusnya[i]))
			continue
		}
		for j, v := range harusnya[i] {
			if hasil[j] != v {
				t.Errorf("❌ channels[%d][%d] = %d, harusnya %d", i, j, hasil[j], v)
			}
		}
	}
	t.Log("✅ SebarKe round-robin benar")
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 5: FAN-IN
// ═══════════════════════════════════════════════════════════════════════════════

func TestGabungSemua(t *testing.T) {
	t.Log("📦 GabungSemua: gabungkan N channel menjadi satu")

	ch1 := GenRange(1, 3)
	ch2 := GenRange(4, 6)
	ch3 := GenRange(7, 9)
	gabung := GabungSemua(ch1, ch2, ch3)

	if gabung == nil {
		t.Fatal("❌ channel tidak boleh nil")
	}

	var hasil []int
	for v := range gabung {
		hasil = append(hasil, v)
	}

	if len(hasil) != 9 {
		t.Fatalf("❌ panjang %d, harusnya 9", len(hasil))
	}

	sort.Ints(hasil)
	for i, v := range hasil {
		if v != i+1 {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, v, i+1)
		}
	}
	t.Logf("✅ GabungSemua (setelah sort): %v", hasil)
}

func TestGabungSemua_SatuChannel(t *testing.T) {
	ch := GenRange(1, 5)
	gabung := GabungSemua(ch)

	var hasil []int
	for v := range gabung {
		hasil = append(hasil, v)
	}

	if len(hasil) != 5 {
		t.Fatalf("❌ panjang %d, harusnya 5", len(hasil))
	}
	t.Logf("✅ GabungSemua (1 channel): %v", hasil)
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 6: TIMEOUT
// ═══════════════════════════════════════════════════════════════════════════════

func TestDenganTimeout(t *testing.T) {
	t.Log("📦 DenganTimeout: teruskan nilai sampai timeout lalu tutup channel")

	// Buat infinite generator (dari Day 25)
	done := make(chan struct{})
	defer close(done)

	gen := GeneratorDenganDone(done)
	terbatas := DenganTimeout(gen, 50*time.Millisecond)

	if terbatas == nil {
		t.Fatal("❌ channel tidak boleh nil")
	}

	// Paling tidak dapat beberapa nilai sebelum timeout
	var hasil []int
	for v := range terbatas {
		hasil = append(hasil, v)
	}

	if len(hasil) == 0 {
		t.Error("❌ harusnya dapat minimal 1 nilai sebelum timeout")
	}
	// Nilai harus mulai dari 1 dan berurutan
	for i, v := range hasil {
		if v != i+1 {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, v, i+1)
		}
	}
	t.Logf("✅ DenganTimeout(50ms): dapat %d nilai: %v", len(hasil), hasil)
}

func TestDenganTimeout_ChannelSelesaiSebelumTimeout(t *testing.T) {
	t.Log("📦 DenganTimeout: channel selesai sebelum timeout → output juga selesai")

	src := GenRange(1, 3) // hanya 3 nilai, pasti selesai sebelum 500ms
	out := DenganTimeout(src, 500*time.Millisecond)

	var hasil []int
	for v := range out {
		hasil = append(hasil, v)
	}

	harusnya := []int{1, 2, 3}
	if len(hasil) != len(harusnya) {
		t.Fatalf("❌ panjang %d, harusnya %d", len(hasil), len(harusnya))
	}
	for i, v := range harusnya {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ DenganTimeout (selesai sebelum timeout): %v", hasil)
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 7: BATCH
// ═══════════════════════════════════════════════════════════════════════════════

func TestBatch(t *testing.T) {
	t.Log("📦 Batch: kumpulkan n nilai jadi []int per batch")

	src := GenRange(1, 7) // 1,2,3,4,5,6,7
	batched := Batch(src, 3)

	if batched == nil {
		t.Fatal("❌ channel tidak boleh nil")
	}

	var semua [][]int
	for grup := range batched {
		semua = append(semua, grup)
	}

	// Harusnya: [1,2,3], [4,5,6], [7]
	if len(semua) != 3 {
		t.Fatalf("❌ jumlah batch %d, harusnya 3", len(semua))
	}

	harusnya := [][]int{{1, 2, 3}, {4, 5, 6}, {7}}
	for i, grup := range harusnya {
		if len(semua[i]) != len(grup) {
			t.Errorf("❌ batch[%d] panjang %d, harusnya %d", i, len(semua[i]), len(grup))
			continue
		}
		for j, v := range grup {
			if semua[i][j] != v {
				t.Errorf("❌ batch[%d][%d] = %d, harusnya %d", i, j, semua[i][j], v)
			}
		}
	}
	t.Logf("✅ Batch(src, 3): %v", semua)
}

func TestBatch_PasGenap(t *testing.T) {
	src := GenRange(1, 6) // 6 nilai, size=2 → 3 batch penuh
	batched := Batch(src, 2)

	var semua [][]int
	for grup := range batched {
		semua = append(semua, grup)
	}

	if len(semua) != 3 {
		t.Fatalf("❌ jumlah batch %d, harusnya 3", len(semua))
	}
	harusnya := [][]int{{1, 2}, {3, 4}, {5, 6}}
	for i, g := range harusnya {
		for j, v := range g {
			if semua[i][j] != v {
				t.Errorf("❌ batch[%d][%d] = %d, harusnya %d", i, j, semua[i][j], v)
			}
		}
	}
	t.Logf("✅ Batch (pas genap): %v", semua)
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 8: ZIP
// ═══════════════════════════════════════════════════════════════════════════════

func TestZip(t *testing.T) {
	t.Log("📦 Zip: pasangkan nilai dari dua channel")

	a := GenRange(1, 3)
	b := GenRange(10, 12)
	zipped := Zip(a, b)

	if zipped == nil {
		t.Fatal("❌ channel tidak boleh nil")
	}

	var hasil [][2]int
	for pair := range zipped {
		hasil = append(hasil, pair)
	}

	harusnya := [][2]int{{1, 10}, {2, 11}, {3, 12}}
	if len(hasil) != len(harusnya) {
		t.Fatalf("❌ panjang %d, harusnya %d", len(hasil), len(harusnya))
	}
	for i, v := range harusnya {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %v, harusnya %v", i, hasil[i], v)
		}
	}
	t.Logf("✅ Zip: %v", hasil)
}

func TestZip_BedaPanjang(t *testing.T) {
	t.Log("📦 Zip: berhenti saat salah satu channel habis")

	a := GenRange(1, 5)  // 5 nilai
	b := GenRange(1, 3)  // 3 nilai (lebih pendek)
	zipped := Zip(a, b)

	var hasil [][2]int
	for pair := range zipped {
		hasil = append(hasil, pair)
	}

	// Harus berhenti setelah b habis → 3 pasang
	if len(hasil) != 3 {
		t.Fatalf("❌ panjang %d, harusnya 3", len(hasil))
	}
	t.Logf("✅ Zip (beda panjang, berhenti di yg terpendek): %v", hasil)
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 9: COMPOSE PIPELINE LENGKAP
// ═══════════════════════════════════════════════════════════════════════════════

func TestTransformasiChan(t *testing.T) {
	t.Log("📦 TransformasiChan: terapkan fn ke tiap nilai, return <-chan int")

	src := GenRange(1, 5)
	kuadrat := TransformasiChan(src, func(n int) int { return n * n })

	if kuadrat == nil {
		t.Fatal("❌ channel tidak boleh nil")
	}

	var hasil []int
	for v := range kuadrat {
		hasil = append(hasil, v)
	}

	harusnya := []int{1, 4, 9, 16, 25}
	if len(hasil) != len(harusnya) {
		t.Fatalf("❌ panjang %d, harusnya %d", len(hasil), len(harusnya))
	}
	for i, v := range harusnya {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ TransformasiChan (kuadrat): %v", hasil)
}

func TestHitungGenapKuadrat(t *testing.T) {
	t.Log("📦 HitungGenapKuadrat: pipeline lengkap GenRange→Filter→TransformasiChan")
	t.Log("   [1..6] → filter genap → [2,4,6] → kuadrat → [4,16,36]")

	out := HitungGenapKuadrat(1, 6)
	if out == nil {
		t.Fatal("❌ channel tidak boleh nil")
	}

	var hasil []int
	for v := range out {
		hasil = append(hasil, v)
	}

	harusnya := []int{4, 16, 36}
	if len(hasil) != len(harusnya) {
		t.Fatalf("❌ panjang %d, harusnya %d — dapat %v", len(hasil), len(harusnya), hasil)
	}
	for i, v := range harusnya {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ HitungGenapKuadrat(1,6): %v", hasil)
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST KOMPOSISI — SEMUA DIGABUNG TANPA SATU PUN KONVERSI MANUAL
// ═══════════════════════════════════════════════════════════════════════════════

func TestPipelineKomposisi(t *testing.T) {
	t.Log("📦 Pipeline komposisi: GenRange → Filter → TransformasiChan → Kali → Ambil")
	t.Log("   Tanpa satu pun konversi int manual!")

	// [1..20] → ganjil → *3 → ambil 4
	// ganjil di 1..20: 1,3,5,7,9,11,13,15,17,19
	// *3: 3,9,15,21,27,...
	// ambil 4: [3,9,15,21]
	out := Ambil(
		Kali(
			Filter(GenRange(1, 20), func(n int) bool { return n%2 != 0 }),
			3,
		),
		4,
	)

	var hasil []int
	for v := range out {
		hasil = append(hasil, v)
	}

	harusnya := []int{3, 9, 15, 21}
	if len(hasil) != len(harusnya) {
		t.Fatalf("❌ panjang %d, harusnya %d — dapat %v", len(hasil), len(harusnya), hasil)
	}
	for i, v := range harusnya {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ Pipeline komposisi: %v (zero konversi manual!)", hasil)
}
