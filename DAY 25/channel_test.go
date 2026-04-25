package belajar

import (
	"sort"
	"sync"
	"testing"
	"time"
)

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 1: CHANNEL DASAR
// ═══════════════════════════════════════════════════════════════════════════════

func TestKirimSatu(t *testing.T) {
	t.Log("📦 KirimSatu: kirim satu nilai ke channel")

	ch := KirimSatu(42)
	if ch == nil {
		t.Fatal("❌ channel tidak boleh nil")
	}

	v := <-ch
	if v != 42 {
		t.Errorf("❌ dapat %d, harusnya 42", v)
	}

	// Pastikan channel sudah ditutup (tidak blokir)
	_, ok := <-ch
	if ok {
		t.Error("❌ channel harusnya sudah ditutup setelah nilai diambil")
	}

	t.Log("✅ KirimSatu benar")
}

func TestKirimBanyak(t *testing.T) {
	t.Log("📦 KirimBanyak: kirim banyak nilai secara berurutan")

	ch := KirimBanyak(1, 2, 3, 4, 5)
	if ch == nil {
		t.Fatal("❌ channel tidak boleh nil")
	}

	var hasil []int
	for v := range ch {
		hasil = append(hasil, v)
	}

	harusnya := []int{1, 2, 3, 4, 5}
	if len(hasil) != len(harusnya) {
		t.Fatalf("❌ panjang berbeda: dapat %d, harusnya %d", len(hasil), len(harusnya))
	}
	for i := range harusnya {
		if hasil[i] != harusnya[i] {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], harusnya[i])
		}
	}

	t.Logf("✅ KirimBanyak urutan: %v", hasil)
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 2: GOROUTINE + CHANNEL
// ═══════════════════════════════════════════════════════════════════════════════

func TestHitungDiBackground(t *testing.T) {
	t.Log("📦 HitungDiBackground: jalankan fungsi di goroutine, ambil hasilnya")

	mulai := time.Now()
	ch := HitungDiBackground(func() int {
		time.Sleep(50 * time.Millisecond)
		return 99
	})

	if ch == nil {
		t.Fatal("❌ channel tidak boleh nil")
	}

	// Goroutine harusnya jalan di background — fungsi HitungDiBackground
	// tidak boleh blokir selama 50ms
	if time.Since(mulai) > 30*time.Millisecond {
		t.Error("❌ HitungDiBackground harusnya langsung return (non-blocking)")
	}

	hasil := <-ch
	if hasil != 99 {
		t.Errorf("❌ dapat %d, harusnya 99", hasil)
	}

	t.Log("✅ HitungDiBackground benar")
}

func TestJalankanN(t *testing.T) {
	t.Log("📦 JalankanN: jalankan fn sebanyak n kali secara paralel")

	hasil := JalankanN(func() int { return 7 }, 5)
	if hasil == nil {
		t.Fatal("❌ hasil tidak boleh nil")
	}
	if len(hasil) != 5 {
		t.Fatalf("❌ panjang %d, harusnya 5", len(hasil))
	}
	for i, v := range hasil {
		if v != 7 {
			t.Errorf("❌ hasil[%d] = %d, harusnya 7", i, v)
		}
	}

	// Test urutan TIDAK perlu sama — yang penting semua ada
	var mu sync.Mutex
	counter := 0
	hasilCounter := JalankanN(func() int {
		mu.Lock()
		defer mu.Unlock()
		counter++
		return counter
	}, 4)

	sort.Ints(hasilCounter)
	harusnya := []int{1, 2, 3, 4}
	for i, v := range harusnya {
		if hasilCounter[i] != v {
			t.Errorf("❌ hasilCounter[%d] = %d, harusnya %d", i, hasilCounter[i], v)
		}
	}

	t.Logf("✅ JalankanN(fn, 5) = %v (urutan boleh acak)", hasilCounter)
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 3: MENJAGA URUTAN
// ═══════════════════════════════════════════════════════════════════════════════

func TestTransformasiUrut(t *testing.T) {
	t.Log("📦 TransformasiUrut: hasil HARUS urut sesuai input")

	input := []int{3, 1, 4, 1, 5, 9, 2, 6}
	hasil := TransformasiUrut(input, func(n int) int { return n * n })

	if hasil == nil {
		t.Fatal("❌ hasil tidak boleh nil")
	}
	if len(hasil) != len(input) {
		t.Fatalf("❌ panjang %d, harusnya %d", len(hasil), len(input))
	}

	harusnya := []int{9, 1, 16, 1, 25, 81, 4, 36}
	for i := range harusnya {
		if hasil[i] != harusnya[i] {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d (input[%d]=%d)",
				i, hasil[i], harusnya[i], i, input[i])
		}
	}
	t.Logf("✅ TransformasiUrut: %v → %v", input, hasil)
}

func TestTransformasiUrut_DenganDelay(t *testing.T) {
	t.Log("📦 TransformasiUrut + delay: goroutine bisa selesai tidak urut, tapi hasil harus urut")

	// Input: [5, 1, 3, 2, 4]
	// Dengan delay = n ms: goroutine input[1]=1ms selesai duluan, tapi hasil[1] tetap fn(1)
	input := []int{5, 1, 3, 2, 4}
	hasil := TransformasiUrut(input, func(n int) int {
		time.Sleep(time.Duration(n) * time.Millisecond) // goroutine kecil selesai duluan!
		return n * 10
	})

	harusnya := []int{50, 10, 30, 20, 40}
	for i := range harusnya {
		if hasil[i] != harusnya[i] {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], harusnya[i])
		}
	}
	t.Logf("✅ Urutan terjaga meski goroutine selesai tidak urut: %v", hasil)
}

func TestTransformasiAcak(t *testing.T) {
	t.Log("📦 TransformasiAcak: semua hasil ada tapi urutan boleh acak")

	input := []int{1, 2, 3, 4, 5}
	hasil := TransformasiAcak(input, func(n int) int { return n * n })

	if hasil == nil {
		t.Fatal("❌ hasil tidak boleh nil")
	}
	if len(hasil) != 5 {
		t.Fatalf("❌ panjang %d, harusnya 5", len(hasil))
	}

	sort.Ints(hasil)
	harusnya := []int{1, 4, 9, 16, 25}
	for i, v := range harusnya {
		if hasil[i] != v {
			t.Errorf("❌ setelah sort: hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ TransformasiAcak (setelah sort): %v", hasil)
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 4: PIPELINE
// ═══════════════════════════════════════════════════════════════════════════════

func TestGandakan(t *testing.T) {
	t.Log("📦 Gandakan: tiap nilai dari channel × 2")

	src := KirimBanyak(1, 2, 3, 4, 5)
	doubled := Gandakan(src)

	if doubled == nil {
		t.Fatal("❌ channel tidak boleh nil")
	}

	var hasil []int
	for v := range doubled {
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
	t.Logf("✅ Gandakan: %v", hasil)
}

func TestTambah(t *testing.T) {
	t.Log("📦 Tambah: tiap nilai dari channel + n")

	src := KirimBanyak(1, 2, 3)
	added := Tambah(src, 10)

	if added == nil {
		t.Fatal("❌ channel tidak boleh nil")
	}

	var hasil []int
	for v := range added {
		hasil = append(hasil, v)
	}

	harusnya := []int{11, 12, 13}
	for i, v := range harusnya {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ Tambah: %v", hasil)
}

func TestPipelineLengkap(t *testing.T) {
	t.Log("📦 PipelineLengkap: KirimBanyak → Gandakan → Tambah")
	t.Log("   [1,2,3] → ×2 → [2,4,6] → +10 → [12,14,16]")

	out := PipelineLengkap([]int{1, 2, 3}, 10)
	if out == nil {
		t.Fatal("❌ channel tidak boleh nil")
	}

	var hasil []int
	for v := range out {
		hasil = append(hasil, v)
	}

	harusnya := []int{12, 14, 16}
	if len(hasil) != len(harusnya) {
		t.Fatalf("❌ panjang %d, harusnya %d", len(hasil), len(harusnya))
	}
	for i, v := range harusnya {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ Pipeline: %v", hasil)
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 5: FAN-IN
// ═══════════════════════════════════════════════════════════════════════════════

func TestGabungDua(t *testing.T) {
	t.Log("📦 GabungDua: gabungkan dua channel menjadi satu")

	ch1 := KirimBanyak(1, 3, 5)
	ch2 := KirimBanyak(2, 4, 6)
	gabungan := GabungDua(ch1, ch2)

	if gabungan == nil {
		t.Fatal("❌ channel tidak boleh nil")
	}

	var hasil []int
	for v := range gabungan {
		hasil = append(hasil, v)
	}

	if len(hasil) != 6 {
		t.Fatalf("❌ panjang %d, harusnya 6", len(hasil))
	}

	sort.Ints(hasil)
	harusnya := []int{1, 2, 3, 4, 5, 6}
	for i, v := range harusnya {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ GabungDua (setelah sort): %v", hasil)
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 6: SELECT
// ═══════════════════════════════════════════════════════════════════════════════

func TestAmbilTercepat(t *testing.T) {
	t.Log("📦 AmbilTercepat: return hasil goroutine yang selesai lebih dulu")

	// fn2 lebih cepat (10ms vs 50ms)
	hasil := AmbilTercepat(
		func() int { time.Sleep(50 * time.Millisecond); return 1 },
		func() int { time.Sleep(10 * time.Millisecond); return 2 },
	)

	if hasil != 2 {
		t.Errorf("❌ dapat %d, harusnya 2 (fn2 lebih cepat)", hasil)
	}
	t.Logf("✅ AmbilTercepat = %d", hasil)
}

func TestAmbilTercepat_fn1LebihCepat(t *testing.T) {
	t.Log("📦 AmbilTercepat: fn1 lebih cepat sekarang")

	hasil := AmbilTercepat(
		func() int { time.Sleep(5 * time.Millisecond); return 111 },
		func() int { time.Sleep(100 * time.Millisecond); return 222 },
	)

	if hasil != 111 {
		t.Errorf("❌ dapat %d, harusnya 111 (fn1 lebih cepat)", hasil)
	}
	t.Logf("✅ AmbilTercepat = %d", hasil)
}

func TestCekAtauDefault(t *testing.T) {
	t.Log("📦 CekAtauDefault: ambil nilai dari channel tanpa blokir")

	// Channel berisi nilai
	ch := make(chan int, 1)
	ch <- 99

	v, ada := CekAtauDefault(ch)
	if !ada {
		t.Error("❌ harusnya ada nilai (true)")
	}
	if v != 99 {
		t.Errorf("❌ dapat %d, harusnya 99", v)
	}

	// Channel kosong
	v, ada = CekAtauDefault(ch)
	if ada {
		t.Error("❌ harusnya tidak ada nilai (false)")
	}
	if v != 0 {
		t.Errorf("❌ default harus 0, dapat %d", v)
	}

	t.Log("✅ CekAtauDefault benar")
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 7: DONE CHANNEL
// ═══════════════════════════════════════════════════════════════════════════════

func TestGeneratorDenganDone(t *testing.T) {
	t.Log("📦 GeneratorDenganDone: hasilkan angka naik sampai done ditutup")

	done := make(chan struct{})
	ch := GeneratorDenganDone(done)

	if ch == nil {
		t.Fatal("❌ channel tidak boleh nil")
	}

	// Ambil 5 nilai pertama
	for i := 1; i <= 5; i++ {
		v := <-ch
		if v != i {
			t.Errorf("❌ nilai ke-%d = %d, harusnya %d", i, v, i)
		}
	}

	close(done) // stop generator

	// Pastikan channel akhirnya ditutup (tidak blokir selamanya)
	timeout := time.After(500 * time.Millisecond)
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				t.Log("✅ GeneratorDenganDone berhenti saat done ditutup")
				return
			}
		case <-timeout:
			t.Fatal("❌ channel tidak ditutup setelah done ditutup (timeout 500ms)")
		}
	}
}

func TestAmbilN(t *testing.T) {
	t.Log("📦 AmbilN: ambil tepat n nilai dari generator")

	done := make(chan struct{})
	defer close(done)

	gen := GeneratorDenganDone(done)
	hasil := AmbilN(gen, 7)

	if len(hasil) != 7 {
		t.Fatalf("❌ panjang %d, harusnya 7", len(hasil))
	}

	for i, v := range hasil {
		if v != i+1 {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, v, i+1)
		}
	}
	t.Logf("✅ AmbilN(gen, 7) = %v", hasil)
}

