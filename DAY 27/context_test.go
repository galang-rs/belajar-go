package belajar

import (
	"context"
	"sort"
	"testing"
	"time"
)

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 1: GENERATOR DENGAN CONTEXT
// ═══════════════════════════════════════════════════════════════════════════════

func TestGeneratorCtx(t *testing.T) {
	t.Log("📦 GeneratorCtx: hasilkan angka naik, berhenti saat ctx di-cancel")

	ctx, cancel := context.WithCancel(context.Background())

	gen := GeneratorCtx(ctx)
	if gen == nil {
		t.Fatal("❌ channel tidak boleh nil")
	}

	// Ambil beberapa nilai dulu
	for i := 1; i <= 5; i++ {
		v := <-gen
		if v != i {
			t.Errorf("❌ nilai ke-%d = %d, harusnya %d", i, v, i)
		}
	}

	// Batalkan context
	cancel()

	// Drain channel sampai tertutup (dengan timeout safety)
	selesai := make(chan struct{})
	go func() {
		for range gen {
		}
		close(selesai)
	}()

	select {
	case <-selesai:
		t.Log("✅ GeneratorCtx: channel tertutup setelah cancel")
	case <-time.After(500 * time.Millisecond):
		t.Error("❌ GeneratorCtx: channel tidak ditutup setelah cancel (timeout 500ms)")
	}
}

func TestGeneratorCtx_NonBlocking(t *testing.T) {
	t.Log("📦 GeneratorCtx: fungsi harus langsung return (non-blocking)")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mulai := time.Now()
	gen := GeneratorCtx(ctx)
	if time.Since(mulai) > 10*time.Millisecond {
		t.Error("❌ GeneratorCtx harus langsung return, bukan blokir")
	}
	_ = gen
	t.Log("✅ GeneratorCtx langsung return")
}

func TestGenRangeCtx(t *testing.T) {
	t.Log("📦 GenRangeCtx: hasilkan range angka, berhenti jika ctx di-cancel")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var hasil []int
	for v := range GenRangeCtx(ctx, 3, 7) {
		hasil = append(hasil, v)
	}

	harusnya := []int{3, 4, 5, 6, 7}
	if len(hasil) != len(harusnya) {
		t.Fatalf("❌ panjang %d, harusnya %d — dapat %v", len(hasil), len(harusnya), hasil)
	}
	for i, v := range harusnya {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ GenRangeCtx(ctx, 3, 7): %v", hasil)
}

func TestGenRangeCtx_CancelDiTengah(t *testing.T) {
	t.Log("📦 GenRangeCtx: cancel di tengah → hasilkan lebih sedikit dari range penuh")

	ctx, cancel := context.WithCancel(context.Background())

	// Range besar: 1..1000, tapi kita cancel setelah mendapat 3 nilai
	ch := GenRangeCtx(ctx, 1, 1000)

	var hasil []int
	for v := range ch {
		hasil = append(hasil, v)
		if len(hasil) == 3 {
			cancel() // batalkan setelah 3 nilai
			break
		}
	}
	// Drain sisa (agar goroutine tidak leak)
	go func() {
		for range ch {
		}
	}()

	if len(hasil) != 3 {
		t.Errorf("❌ sebelum cancel dapat %d nilai, harusnya 3", len(hasil))
	}
	for i, v := range hasil {
		if v != i+1 {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, v, i+1)
		}
	}
	t.Logf("✅ GenRangeCtx cancel di tengah: %v", hasil)
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 2: TRANSFORM DENGAN CONTEXT
// ═══════════════════════════════════════════════════════════════════════════════

func TestFilterCtx(t *testing.T) {
	t.Log("📦 FilterCtx: teruskan nilai yang lolos fn, hormati ctx")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	src := GenRangeCtx(ctx, 1, 10)
	genap := FilterCtx(ctx, src, func(n int) bool { return n%2 == 0 })

	if genap == nil {
		t.Fatal("❌ channel tidak boleh nil")
	}

	var hasil []int
	for v := range genap {
		hasil = append(hasil, v)
	}

	harusnya := []int{2, 4, 6, 8, 10}
	if len(hasil) != len(harusnya) {
		t.Fatalf("❌ panjang %d, harusnya %d — dapat %v", len(hasil), len(harusnya), hasil)
	}
	for i, v := range harusnya {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ FilterCtx(genap): %v", hasil)
}

func TestFilterCtx_CancelMenghentikanOutput(t *testing.T) {
	t.Log("📦 FilterCtx: cancel → channel output tertutup")

	ctx, cancel := context.WithCancel(context.Background())

	// Infinite generator, filter semua (fn selalu true)
	gen := GeneratorCtx(ctx)
	out := FilterCtx(ctx, gen, func(n int) bool { return true })

	// Ambil beberapa nilai
	for i := 0; i < 5; i++ {
		<-out
	}
	cancel()

	// Output harus tertutup
	selesai := make(chan struct{})
	go func() {
		for range out {
		}
		close(selesai)
	}()

	select {
	case <-selesai:
		t.Log("✅ FilterCtx: output tertutup setelah cancel")
	case <-time.After(500 * time.Millisecond):
		t.Error("❌ FilterCtx: output tidak tertutup setelah cancel")
	}
}

func TestTransformasiCtx(t *testing.T) {
	t.Log("📦 TransformasiCtx: terapkan fn ke tiap nilai, hormati ctx")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	src := GenRangeCtx(ctx, 1, 5)
	kuadrat := TransformasiCtx(ctx, src, func(n int) int { return n * n })

	if kuadrat == nil {
		t.Fatal("❌ channel tidak boleh nil")
	}

	var hasil []int
	for v := range kuadrat {
		hasil = append(hasil, v)
	}

	harusnya := []int{1, 4, 9, 16, 25}
	if len(hasil) != len(harusnya) {
		t.Fatalf("❌ panjang %d, harusnya %d — dapat %v", len(hasil), len(harusnya), hasil)
	}
	for i, v := range harusnya {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ TransformasiCtx(kuadrat): %v", hasil)
}

func TestTransformasiCtx_CancelMenghentikanOutput(t *testing.T) {
	t.Log("📦 TransformasiCtx: cancel → channel output tertutup")

	ctx, cancel := context.WithCancel(context.Background())

	gen := GeneratorCtx(ctx)
	out := TransformasiCtx(ctx, gen, func(n int) int { return n * 2 })

	for i := 0; i < 5; i++ {
		<-out
	}
	cancel()

	selesai := make(chan struct{})
	go func() {
		for range out {
		}
		close(selesai)
	}()

	select {
	case <-selesai:
		t.Log("✅ TransformasiCtx: output tertutup setelah cancel")
	case <-time.After(500 * time.Millisecond):
		t.Error("❌ TransformasiCtx: output tidak tertutup setelah cancel")
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 3: WORKER POOL DENGAN CONTEXT
// ═══════════════════════════════════════════════════════════════════════════════

func TestWorkerPool(t *testing.T) {
	t.Log("📦 WorkerPool: N goroutine memproses jobs secara paralel")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := GenRangeCtx(ctx, 1, 8)
	hasil_ch := WorkerPool(ctx, jobs, 3, func(n int) int { return n * n })

	if hasil_ch == nil {
		t.Fatal("❌ channel tidak boleh nil")
	}

	var hasil []int
	for v := range hasil_ch {
		hasil = append(hasil, v)
	}

	if len(hasil) != 8 {
		t.Fatalf("❌ panjang %d, harusnya 8 — dapat %v", len(hasil), hasil)
	}

	sort.Ints(hasil)
	harusnya := []int{1, 4, 9, 16, 25, 36, 49, 64}
	for i, v := range harusnya {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ WorkerPool(3 worker, kuadrat): %v", hasil)
}

func TestWorkerPool_CancelMenghentikan(t *testing.T) {
	t.Log("📦 WorkerPool: cancel → output channel tertutup")

	ctx, cancel := context.WithCancel(context.Background())

	gen := GeneratorCtx(ctx)
	hasil_ch := WorkerPool(ctx, gen, 3, func(n int) int { return n })

	// Ambil beberapa hasil
	count := 0
	for range hasil_ch {
		count++
		if count == 5 {
			cancel()
			break
		}
	}

	// Drain sisa
	selesai := make(chan struct{})
	go func() {
		for range hasil_ch {
		}
		close(selesai)
	}()

	select {
	case <-selesai:
		t.Log("✅ WorkerPool: output tertutup setelah cancel")
	case <-time.After(500 * time.Millisecond):
		t.Error("❌ WorkerPool: output tidak tertutup setelah cancel")
	}
}

func TestWorkerPool_SatuWorker(t *testing.T) {
	t.Log("📦 WorkerPool: 1 worker — harus tetap bekerja")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := GenRangeCtx(ctx, 1, 5)
	hasil_ch := WorkerPool(ctx, jobs, 1, func(n int) int { return n + 10 })

	var hasil []int
	for v := range hasil_ch {
		hasil = append(hasil, v)
	}

	sort.Ints(hasil)
	harusnya := []int{11, 12, 13, 14, 15}
	if len(hasil) != len(harusnya) {
		t.Fatalf("❌ panjang %d, harusnya 5 — dapat %v", len(hasil), hasil)
	}
	for i, v := range harusnya {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ WorkerPool(1 worker, +10): %v", hasil)
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 4: RETRY DENGAN CONTEXT
// ═══════════════════════════════════════════════════════════════════════════════

func TestCobaLagi_Berhasil(t *testing.T) {
	t.Log("📦 CobaLagi: berhasil di percobaan ke-3")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	percobaan := 0
	berhasil := CobaLagi(ctx, 5, 10*time.Millisecond, func() bool {
		percobaan++
		return percobaan >= 3
	})

	if !berhasil {
		t.Error("❌ harusnya berhasil")
	}
	if percobaan != 3 {
		t.Errorf("❌ jumlah percobaan %d, harusnya 3", percobaan)
	}
	t.Logf("✅ CobaLagi berhasil di percobaan ke-%d", percobaan)
}

func TestCobaLagi_HabisPercobaan(t *testing.T) {
	t.Log("📦 CobaLagi: maks tercapai sebelum berhasil → return false")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	percobaan := 0
	berhasil := CobaLagi(ctx, 3, 10*time.Millisecond, func() bool {
		percobaan++
		return false // selalu gagal
	})

	if berhasil {
		t.Error("❌ harusnya gagal (habis percobaan)")
	}
	if percobaan != 3 {
		t.Errorf("❌ jumlah percobaan %d, harusnya tepat 3", percobaan)
	}
	t.Logf("✅ CobaLagi habis maks(%d percobaan), return false", percobaan)
}

func TestCobaLagi_CancelMenghentikan(t *testing.T) {
	t.Log("📦 CobaLagi: ctx di-cancel → berhenti lebih awal")

	ctx, cancel := context.WithCancel(context.Background())

	percobaan := 0
	selesai := make(chan bool)
	go func() {
		result := CobaLagi(ctx, 100, 50*time.Millisecond, func() bool {
			percobaan++
			return false
		})
		selesai <- result
	}()

	// Tunggu 1 percobaan, lalu cancel
	time.Sleep(10 * time.Millisecond)
	cancel()

	select {
	case berhasil := <-selesai:
		if berhasil {
			t.Error("❌ setelah cancel harusnya return false")
		}
		if percobaan >= 100 {
			t.Error("❌ seharusnya berhenti jauh sebelum 100 percobaan")
		}
		t.Logf("✅ CobaLagi berhenti setelah cancel (%d percobaan)", percobaan)
	case <-time.After(1 * time.Second):
		t.Error("❌ CobaLagi tidak berhenti setelah cancel")
	}
}

func TestCobaLagi_LangsungBerhasil(t *testing.T) {
	t.Log("📦 CobaLagi: berhasil di percobaan pertama")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	percobaan := 0
	berhasil := CobaLagi(ctx, 5, 10*time.Millisecond, func() bool {
		percobaan++
		return true // langsung berhasil
	})

	if !berhasil {
		t.Error("❌ harusnya berhasil")
	}
	if percobaan != 1 {
		t.Errorf("❌ harusnya hanya 1 percobaan, dapat %d", percobaan)
	}
	t.Log("✅ CobaLagi berhasil di percobaan pertama")
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 5: JALANKAN PARALEL DENGAN BATAS WAKTU
// ═══════════════════════════════════════════════════════════════════════════════

func TestJalankanParalel(t *testing.T) {
	t.Log("📦 JalankanParalel: semua fn dijalankan paralel, hasilkan []int")

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	fns := []func(context.Context) int{
		func(ctx context.Context) int { return 1 },
		func(ctx context.Context) int { return 2 },
		func(ctx context.Context) int { return 3 },
	}

	hasil := JalankanParalel(ctx, fns)

	if len(hasil) != 3 {
		t.Fatalf("❌ panjang %d, harusnya 3 — dapat %v", len(hasil), hasil)
	}

	sort.Ints(hasil)
	harusnya := []int{1, 2, 3}
	for i, v := range harusnya {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ JalankanParalel: %v", hasil)
}

func TestJalankanParalel_BenarParalel(t *testing.T) {
	t.Log("📦 JalankanParalel: harus berjalan secara paralel (bukan sequential)")

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 5 fungsi masing-masing sleep 80ms
	// Sequential: ~400ms. Paralel: ~80ms.
	fns := make([]func(context.Context) int, 5)
	for i := 0; i < 5; i++ {
		i := i
		fns[i] = func(ctx context.Context) int {
			time.Sleep(80 * time.Millisecond)
			return i + 1
		}
	}

	mulai := time.Now()
	hasil := JalankanParalel(ctx, fns)
	durasi := time.Since(mulai)

	if len(hasil) != 5 {
		t.Fatalf("❌ panjang %d, harusnya 5", len(hasil))
	}
	// Paralel seharusnya selesai dalam ~100ms, bukan ~400ms
	if durasi > 200*time.Millisecond {
		t.Errorf("❌ durasi %v — harusnya paralel (~80ms), bukan sequential (~400ms)", durasi)
	}
	t.Logf("✅ JalankanParalel berjalan paralel: durasi %v", durasi)
}

func TestJalankanParalel_SatuFn(t *testing.T) {
	t.Log("📦 JalankanParalel: satu fungsi → slice berisi satu elemen")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fns := []func(context.Context) int{
		func(ctx context.Context) int { return 42 },
	}

	hasil := JalankanParalel(ctx, fns)
	if len(hasil) != 1 || hasil[0] != 42 {
		t.Errorf("❌ harusnya [42], dapat %v", hasil)
	}
	t.Log("✅ JalankanParalel(1 fn) = [42]")
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 6: PIPELINE DENGAN CONTEXT
// ═══════════════════════════════════════════════════════════════════════════════

func TestPipelineCtx(t *testing.T) {
	t.Log("📦 PipelineCtx: GenRangeCtx → FilterCtx → TransformasiCtx → []int")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Angka ganjil di [1,10] dikuadratkan: 1²,3²,5²,7²,9² = 1,9,25,49,81
	hasil := PipelineCtx(ctx, 1, 10,
		func(n int) bool { return n%2 != 0 },
		func(n int) int { return n * n },
	)

	if len(hasil) == 0 {
		t.Fatal("❌ hasil tidak boleh kosong")
	}

	harusnya := []int{1, 9, 25, 49, 81}
	if len(hasil) != len(harusnya) {
		t.Fatalf("❌ panjang %d, harusnya %d — dapat %v", len(hasil), len(harusnya), hasil)
	}
	for i, v := range harusnya {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ PipelineCtx(ganjil, kuadrat): %v", hasil)
}

func TestPipelineCtx_SemuaLolos(t *testing.T) {
	t.Log("📦 PipelineCtx: filter semua lolos (fn selalu true)")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// filter semua lolos, transform x*2
	hasil := PipelineCtx(ctx, 1, 5,
		func(n int) bool { return true },
		func(n int) int { return n * 2 },
	)

	harusnya := []int{2, 4, 6, 8, 10}
	if len(hasil) != len(harusnya) {
		t.Fatalf("❌ panjang %d, harusnya 5 — dapat %v", len(hasil), hasil)
	}
	for i, v := range harusnya {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ PipelineCtx(semua lolos, *2): %v", hasil)
}

func TestPipelineCtx_TidakAdaLolos(t *testing.T) {
	t.Log("📦 PipelineCtx: tidak ada yang lolos filter → slice kosong")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	hasil := PipelineCtx(ctx, 1, 5,
		func(n int) bool { return n > 100 }, // tidak ada yang lolos
		func(n int) int { return n },
	)

	if len(hasil) != 0 {
		t.Errorf("❌ harusnya kosong, dapat %v", hasil)
	}
	t.Log("✅ PipelineCtx(tidak ada lolos): []")
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 7: CEK ERR DARI CONTEXT
// ═══════════════════════════════════════════════════════════════════════════════

func TestAlasanBerhenti_Aktif(t *testing.T) {
	t.Log("📦 AlasanBerhenti: context aktif → \"aktif\"")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	alasan := AlasanBerhenti(ctx)
	if alasan != "aktif" {
		t.Errorf("❌ dapat %q, harusnya \"aktif\"", alasan)
	}
	t.Log("✅ AlasanBerhenti(ctx aktif) = \"aktif\"")
}

func TestAlasanBerhenti_Canceled(t *testing.T) {
	t.Log("📦 AlasanBerhenti: setelah cancel → \"dibatalkan\"")

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	alasan := AlasanBerhenti(ctx)
	if alasan != "dibatalkan" {
		t.Errorf("❌ dapat %q, harusnya \"dibatalkan\"", alasan)
	}
	t.Log("✅ AlasanBerhenti(canceled) = \"dibatalkan\"")
}

func TestAlasanBerhenti_DeadlineExceeded(t *testing.T) {
	t.Log("📦 AlasanBerhenti: setelah deadline → \"waktu habis\"")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	// Tunggu sampai deadline tercapai
	time.Sleep(10 * time.Millisecond)

	alasan := AlasanBerhenti(ctx)
	if alasan != "waktu habis" {
		t.Errorf("❌ dapat %q, harusnya \"waktu habis\"", alasan)
	}
	t.Log("✅ AlasanBerhenti(deadline exceeded) = \"waktu habis\"")
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 8: FIRST-WIN
// ═══════════════════════════════════════════════════════════════════════════════

func TestPertamaSelesai(t *testing.T) {
	t.Log("📦 PertamaSelesai: kembalikan hasil fn yang paling cepat selesai")

	fns := []func() int{
		func() int { time.Sleep(60 * time.Millisecond); return 60 },
		func() int { time.Sleep(20 * time.Millisecond); return 20 }, // paling cepat
		func() int { time.Sleep(40 * time.Millisecond); return 40 },
	}

	v := PertamaSelesai(fns)
	if v != 20 {
		t.Errorf("❌ dapat %d, harusnya 20 (yang paling cepat)", v)
	}
	t.Logf("✅ PertamaSelesai = %d (yang paling cepat)", v)
}

func TestPertamaSelesai_SatuFn(t *testing.T) {
	t.Log("📦 PertamaSelesai: satu fungsi → hasilnya langsung dikembalikan")

	fns := []func() int{
		func() int { return 99 },
	}

	v := PertamaSelesai(fns)
	if v != 99 {
		t.Errorf("❌ dapat %d, harusnya 99", v)
	}
	t.Log("✅ PertamaSelesai(1 fn) = 99")
}

func TestPertamaSelesai_TidakBlokir(t *testing.T) {
	t.Log("📦 PertamaSelesai: tidak menunggu semua fn selesai")

	// fn[0] sangat cepat, fn[1] sangat lambat
	// PertamaSelesai harus return ~segera, bukan menunggu fn[1]
	fns := []func() int{
		func() int { return 1 }, // langsung
		func() int { time.Sleep(500 * time.Millisecond); return 2 }, // lambat
	}

	mulai := time.Now()
	v := PertamaSelesai(fns)
	durasi := time.Since(mulai)

	if v != 1 {
		t.Errorf("❌ dapat %d, harusnya 1 (yang paling cepat)", v)
	}
	if durasi > 100*time.Millisecond {
		t.Errorf("❌ durasi %v — harusnya tidak menunggu fn lambat (~500ms)", durasi)
	}
	t.Logf("✅ PertamaSelesai tidak blokir: %v, nilai=%d", durasi, v)
}
