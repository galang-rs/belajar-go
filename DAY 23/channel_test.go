package belajar

import (
	"sort"
	"sync/atomic"
	"testing"
	"time"
)

// ═══════════════════════════════════════════════════════════════════════════════
// TEST BAGIAN 1: CHANNEL DASAR
// ═══════════════════════════════════════════════════════════════════════════════

func TestKirimSemua_Basic(t *testing.T) {
	ch := make(chan int, 3)
	KirimSemua(ch, 10, 20, 30)

	want := []int{10, 20, 30}
	for i, w := range want {
		// Coba terima dengan timeout
		select {
		case got := <-ch:
			if got != w {
				t.Errorf("KirimSemua: nilai ke-%d want %d, got %d", i, w, got)
			}
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("KirimSemua: timeout membaca nilai ke-%d", i)
		}
	}
}

func TestKirimSemua_ChannelDitutup(t *testing.T) {
	// Setelah KirimSemua, channel harus ditutup
	ch := make(chan int, 5)
	KirimSemua(ch, 1, 2, 3)

	// Baca semua nilai dulu
	<-ch
	<-ch
	<-ch

	// Channel harus sudah ditutup → ok harus false
	_, ok := <-ch
	if ok {
		t.Error("KirimSemua: channel harus ditutup setelah semua data terkirim")
	}
}

func TestKirimSemua_Kosong(t *testing.T) {
	ch := make(chan int, 1)
	KirimSemua(ch) // tanpa data

	_, ok := <-ch
	if ok {
		t.Error("KirimSemua tanpa data: channel harus langsung ditutup")
	}
}

// ───────────────────────────────────────────────────────────────────────────

func TestTerimaSemuaSlice_Basic(t *testing.T) {
	ch := buatChannel(1, 2, 3, 4, 5)
	hasil := TerimaSemuaSlice(ch)

	if len(hasil) != 5 {
		t.Fatalf("TerimaSemuaSlice: want len=5, got %d", len(hasil))
	}
	want := []int{1, 2, 3, 4, 5}
	for i, w := range want {
		if hasil[i] != w {
			t.Errorf("TerimaSemuaSlice[%d]: want %d, got %d", i, w, hasil[i])
		}
	}
}

func TestTerimaSemuaSlice_Kosong(t *testing.T) {
	ch := buatChannel() // channel kosong langsung ditutup
	hasil := TerimaSemuaSlice(ch)
	if len(hasil) != 0 {
		t.Errorf("TerimaSemuaSlice channel kosong: want len=0, got %d", len(hasil))
	}
}

// ───────────────────────────────────────────────────────────────────────────

func TestCekChannelTertutup_AdaData(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 99

	v, ok := CekChannelTertutup(ch)
	if !ok || v != 99 {
		t.Errorf("CekChannelTertutup ada data: want (99, true), got (%d, %v)", v, ok)
	}
}

func TestCekChannelTertutup_Ditutup(t *testing.T) {
	ch := make(chan int)
	close(ch)

	v, ok := CekChannelTertutup(ch)
	if ok || v != 0 {
		t.Errorf("CekChannelTertutup ditutup: want (0, false), got (%d, %v)", v, ok)
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// TEST BAGIAN 2: GOROUTINE + CHANNEL (ASYNC)
// ═══════════════════════════════════════════════════════════════════════════════

func TestHitungAsync_Nilai(t *testing.T) {
	cases := []struct{ a, b, want int }{
		{3, 7, 10},
		{0, 0, 0},
		{-5, 5, 0},
		{100, 200, 300},
		{1, -1, 0},
	}
	for _, c := range cases {
		ch := HitungAsync(c.a, c.b)
		if ch == nil {
			t.Fatalf("HitungAsync(%d, %d): channel tidak boleh nil", c.a, c.b)
		}
		select {
		case got := <-ch:
			if got != c.want {
				t.Errorf("HitungAsync(%d, %d): want %d, got %d", c.a, c.b, c.want, got)
			}
		case <-time.After(1 * time.Second):
			t.Errorf("HitungAsync(%d, %d): timeout", c.a, c.b)
		}
	}
}

func TestHitungAsync_LangsungReturn(t *testing.T) {
	// HitungAsync harus return channel SEGERA tanpa blocking
	start := time.Now()
	ch := HitungAsync(10, 20)
	elapsed := time.Since(start)

	if ch == nil {
		t.Fatal("HitungAsync: channel nil")
	}
	if elapsed > 100*time.Millisecond {
		t.Errorf("HitungAsync tidak async: elapsed %v, harus < 100ms", elapsed)
	}
	<-ch // pastikan goroutine selesai
}

// ───────────────────────────────────────────────────────────────────────────

func TestKuadratAsync_Nilai(t *testing.T) {
	ch := KuadratAsync(2, 3, 4, 5)
	if ch == nil {
		t.Fatal("KuadratAsync: channel tidak boleh nil")
	}

	// Kumpulkan 4 hasil
	var hasil []int
	for i := 0; i < 4; i++ {
		select {
		case v := <-ch:
			hasil = append(hasil, v)
		case <-time.After(1 * time.Second):
			t.Fatalf("KuadratAsync: timeout setelah %d hasil", i)
		}
	}

	sort.Ints(hasil)
	want := []int{4, 9, 16, 25}
	for i, w := range want {
		if hasil[i] != w {
			t.Errorf("KuadratAsync[%d]: want %d, got %d", i, w, hasil[i])
		}
	}
}

func TestKuadratAsync_Kosong(t *testing.T) {
	ch := KuadratAsync()
	if ch == nil {
		t.Fatal("KuadratAsync kosong: channel tidak boleh nil")
	}
	// Tidak ada yang dikirim, jangan deadlock
	// (channel mungkin kosong, cukup pastikan tidak panic)
}

// ═══════════════════════════════════════════════════════════════════════════════
// TEST BAGIAN 3: SELECT
// ═══════════════════════════════════════════════════════════════════════════════

func TestAmbilTercepat_PertamaLebihCepat(t *testing.T) {
	hasil := AmbilTercepat(
		func() int { return 1 }, // sangat cepat
		func() int { time.Sleep(500 * time.Millisecond); return 2 },
	)
	if hasil != 1 {
		t.Errorf("AmbilTercepat: want 1 (lebih cepat), got %d", hasil)
	}
}

func TestAmbilTercepat_KeduaLebihCepat(t *testing.T) {
	hasil := AmbilTercepat(
		func() int { time.Sleep(500 * time.Millisecond); return 1 },
		func() int { return 2 }, // sangat cepat
	)
	if hasil != 2 {
		t.Errorf("AmbilTercepat: want 2 (lebih cepat), got %d", hasil)
	}
}

func TestAmbilTercepat_LebihCepatDariKeduanya(t *testing.T) {
	// Pastikan lebih cepat dari jika dijalankan sequential
	start := time.Now()
	AmbilTercepat(
		func() int { time.Sleep(200 * time.Millisecond); return 1 },
		func() int { time.Sleep(200 * time.Millisecond); return 2 },
	)
	elapsed := time.Since(start)
	// Harus selesai ~ 200ms, bukan 400ms
	if elapsed > 350*time.Millisecond {
		t.Errorf("AmbilTercepat tidak paralel: elapsed %v", elapsed)
	}
}

// ───────────────────────────────────────────────────────────────────────────

func TestCobaAmbil_ChannelKosong(t *testing.T) {
	ch := make(chan int, 1)
	v, ok := CobaAmbil(ch)
	if ok || v != 0 {
		t.Errorf("CobaAmbil kosong: want (0, false), got (%d, %v)", v, ok)
	}
}

func TestCobaAmbil_AdaData(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 55
	v, ok := CobaAmbil(ch)
	if !ok || v != 55 {
		t.Errorf("CobaAmbil ada data: want (55, true), got (%d, %v)", v, ok)
	}
}

func TestCobaAmbil_TidakBlokir(t *testing.T) {
	ch := make(chan int) // unbuffered, tidak ada pengirim
	start := time.Now()
	_, ok := CobaAmbil(ch)
	elapsed := time.Since(start)

	if ok {
		t.Error("CobaAmbil: harusnya false pada channel kosong")
	}
	if elapsed > 50*time.Millisecond {
		t.Errorf("CobaAmbil blokir! elapsed %v, harus < 50ms", elapsed)
	}
}

// ───────────────────────────────────────────────────────────────────────────

func TestTungguDenganBatas_Berhasil(t *testing.T) {
	hasil, ok := TungguDenganBatas(func() int {
		time.Sleep(10 * time.Millisecond)
		return 42
	}, 500*time.Millisecond)

	if !ok {
		t.Fatal("TungguDenganBatas: harusnya berhasil (ok=true)")
	}
	if hasil != 42 {
		t.Errorf("TungguDenganBatas: want 42, got %d", hasil)
	}
}

func TestTungguDenganBatas_Timeout(t *testing.T) {
	hasil, ok := TungguDenganBatas(func() int {
		time.Sleep(1 * time.Second)
		return 99
	}, 50*time.Millisecond)

	if ok {
		t.Fatal("TungguDenganBatas: harusnya timeout (ok=false)")
	}
	if hasil != 0 {
		t.Errorf("TungguDenganBatas timeout: want 0, got %d", hasil)
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// TEST BAGIAN 4: GOROUTINE + WAITGROUP
// ═══════════════════════════════════════════════════════════════════════════════

func TestJalankanParalel_SemuaJalan(t *testing.T) {
	var counter int64
	tugas := make([]func(), 10)
	for i := range tugas {
		tugas[i] = func() { atomic.AddInt64(&counter, 1) }
	}

	JalankanParalel(tugas...)

	if counter != 10 {
		t.Errorf("JalankanParalel: want counter=10, got %d", counter)
	}
}

func TestJalankanParalel_MenungguSemua(t *testing.T) {
	selesai := make([]bool, 5)
	tugas := make([]func(), 5)
	for i := range tugas {
		i := i
		tugas[i] = func() {
			time.Sleep(30 * time.Millisecond)
			selesai[i] = true
		}
	}

	JalankanParalel(tugas...)

	for i, s := range selesai {
		if !s {
			t.Errorf("JalankanParalel: tugas[%d] belum selesai saat return", i)
		}
	}
}

func TestJalankanParalel_Kosong(t *testing.T) {
	// Tidak boleh panik
	JalankanParalel()
}

func TestJalankanParalel_Paralel(t *testing.T) {
	// 5 tugas × 100ms. Kalau paralel → ~100ms. Sequential → ~500ms.
	start := time.Now()
	JalankanParalel(
		func() { time.Sleep(100 * time.Millisecond) },
		func() { time.Sleep(100 * time.Millisecond) },
		func() { time.Sleep(100 * time.Millisecond) },
		func() { time.Sleep(100 * time.Millisecond) },
		func() { time.Sleep(100 * time.Millisecond) },
	)
	elapsed := time.Since(start)
	if elapsed > 400*time.Millisecond {
		t.Errorf("JalankanParalel tidak berjalan paralel: elapsed %v", elapsed)
	}
}

// ───────────────────────────────────────────────────────────────────────────

func TestKumpulkanHasil_Nilai(t *testing.T) {
	hasil := KumpulkanHasil(
		func() int { return 10 },
		func() int { return 20 },
		func() int { return 30 },
	)

	if len(hasil) != 3 {
		t.Fatalf("KumpulkanHasil: want len=3, got %d", len(hasil))
	}
	sort.Ints(hasil)
	want := []int{10, 20, 30}
	for i, w := range want {
		if hasil[i] != w {
			t.Errorf("KumpulkanHasil[%d]: want %d, got %d", i, w, hasil[i])
		}
	}
}

func TestKumpulkanHasil_Kosong(t *testing.T) {
	hasil := KumpulkanHasil()
	if len(hasil) != 0 {
		t.Errorf("KumpulkanHasil kosong: want len=0, got %d", len(hasil))
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// TEST BAGIAN 5: DONE CHANNEL
// ═══════════════════════════════════════════════════════════════════════════════

func TestGeneratorAngka_Urutan(t *testing.T) {
	selesai := make(chan struct{})
	ch := GeneratorAngka(selesai)

	if ch == nil {
		t.Fatal("GeneratorAngka: channel tidak boleh nil")
	}

	// Ambil 5 angka pertama, verifikasi urutannya
	for want := 1; want <= 5; want++ {
		select {
		case got := <-ch:
			if got != want {
				t.Errorf("GeneratorAngka: want %d, got %d", want, got)
			}
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("GeneratorAngka: timeout menunggu angka ke-%d", want)
		}
	}

	close(selesai)
}

func TestGeneratorAngka_BerhentiSaatDitutup(t *testing.T) {
	selesai := make(chan struct{})
	ch := GeneratorAngka(selesai)

	// Baca beberapa
	<-ch
	<-ch

	// Hentikan
	close(selesai)

	// Channel harus ditutup setelah stop
	timer := time.After(500 * time.Millisecond)
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				return // sukses: channel ditutup
			}
		case <-timer:
			t.Error("GeneratorAngka: channel tidak ditutup setelah selesai ditutup")
			return
		}
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// TEST BAGIAN 6: FAN-OUT & FAN-IN
// ═══════════════════════════════════════════════════════════════════════════════

func TestSebarKeSemua_DuaOutput(t *testing.T) {
	masuk := make(chan int, 3)
	out1 := make(chan int, 3)
	out2 := make(chan int, 3)

	masuk <- 10
	masuk <- 20
	masuk <- 30
	close(masuk)

	SebarKeSemua(masuk, out1, out2)

	// Kedua channel harus berisi nilai yang sama
	for _, ch := range []chan int{out1, out2} {
		var hasil []int
		for {
			select {
			case v, ok := <-ch:
				if !ok {
					goto cekHasil
				}
				hasil = append(hasil, v)
			default:
				goto cekHasil
			}
		}
	cekHasil:
		sort.Ints(hasil)
		want := []int{10, 20, 30}
		if len(hasil) != 3 {
			t.Errorf("SebarKeSemua: want 3 nilai, got %d: %v", len(hasil), hasil)
			continue
		}
		for i, w := range want {
			if hasil[i] != w {
				t.Errorf("SebarKeSemua[%d]: want %d, got %d", i, w, hasil[i])
			}
		}
	}
}

func TestSebarKeSemua_OutputDitutup(t *testing.T) {
	masuk := make(chan int, 1)
	out1 := make(chan int, 1)
	out2 := make(chan int, 1)

	close(masuk) // langsung tutup
	SebarKeSemua(masuk, out1, out2)

	// Semua output harus ditutup
	for i, ch := range []chan int{out1, out2} {
		_, ok := <-ch
		if ok {
			t.Errorf("SebarKeSemua: out%d harus ditutup setelah masuk ditutup", i+1)
		}
	}
}

// ───────────────────────────────────────────────────────────────────────────

func TestGabungkanChannel_DuaChannel(t *testing.T) {
	ch1 := buatChannel(1, 3, 5)
	ch2 := buatChannel(2, 4, 6)
	gabung := GabungkanChannel(ch1, ch2)

	if gabung == nil {
		t.Fatal("GabungkanChannel: channel tidak boleh nil")
	}

	var hasil []int
	for v := range gabung {
		hasil = append(hasil, v)
	}

	if len(hasil) != 6 {
		t.Fatalf("GabungkanChannel: want 6 nilai, got %d", len(hasil))
	}
	sort.Ints(hasil)
	want := []int{1, 2, 3, 4, 5, 6}
	for i, w := range want {
		if hasil[i] != w {
			t.Errorf("GabungkanChannel[%d]: want %d, got %d", i, w, hasil[i])
		}
	}
}

func TestGabungkanChannel_BanyakChannel(t *testing.T) {
	channels := make([]<-chan int, 5)
	for i := range channels {
		i := i
		ch := make(chan int, 1)
		ch <- (i + 1) * 10 // 10, 20, 30, 40, 50
		close(ch)
		channels[i] = ch
	}

	gabung := GabungkanChannel(channels...)
	var hasil []int
	for v := range gabung {
		hasil = append(hasil, v)
	}

	sort.Ints(hasil)
	want := []int{10, 20, 30, 40, 50}
	if len(hasil) != len(want) {
		t.Fatalf("GabungkanChannel: want %v, got %v", want, hasil)
	}
	for i, w := range want {
		if hasil[i] != w {
			t.Errorf("GabungkanChannel[%d]: want %d, got %d", i, w, hasil[i])
		}
	}
}

func TestGabungkanChannel_Ditutup(t *testing.T) {
	// Output harus ditutup setelah semua input selesai
	ch1 := buatChannel(1)
	ch2 := buatChannel(2)
	gabung := GabungkanChannel(ch1, ch2)

	var hasil []int
	timer := time.After(1 * time.Second)
	draining := true
	for draining {
		select {
		case v, ok := <-gabung:
			if !ok {
				draining = false
			} else {
				hasil = append(hasil, v)
			}
		case <-timer:
			t.Error("GabungkanChannel: channel tidak ditutup — timeout")
			draining = false
		}
	}

	if len(hasil) != 2 {
		t.Errorf("GabungkanChannel: want 2 nilai, got %d", len(hasil))
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// TEST BAGIAN 7: WORKER POOL
// ═══════════════════════════════════════════════════════════════════════════════

func TestWorkerPool_Kuadrat(t *testing.T) {
	jobs := make(chan int, 5)
	jobs <- 2
	jobs <- 3
	jobs <- 4
	jobs <- 5
	jobs <- 6
	close(jobs)

	hasilCh := WorkerPool(jobs, 3, func(n int) int { return n * n })
	if hasilCh == nil {
		t.Fatal("WorkerPool: channel tidak boleh nil")
	}

	var hasil []int
	for v := range hasilCh {
		hasil = append(hasil, v)
	}

	if len(hasil) != 5 {
		t.Fatalf("WorkerPool: want 5 hasil, got %d", len(hasil))
	}
	sort.Ints(hasil)
	want := []int{4, 9, 16, 25, 36}
	for i, w := range want {
		if hasil[i] != w {
			t.Errorf("WorkerPool[%d]: want %d, got %d", i, w, hasil[i])
		}
	}
}

func TestWorkerPool_Paralel(t *testing.T) {
	// 10 jobs × 50ms. Dengan 5 worker → ~100ms. Sequential → ~500ms.
	jobs := make(chan int, 10)
	for i := 0; i < 10; i++ {
		jobs <- i
	}
	close(jobs)

	start := time.Now()
	hasilCh := WorkerPool(jobs, 5, func(n int) int {
		time.Sleep(50 * time.Millisecond)
		return n
	})
	for range hasilCh {
	}
	elapsed := time.Since(start)

	if elapsed > 400*time.Millisecond {
		t.Errorf("WorkerPool tidak paralel: elapsed %v", elapsed)
	}
}

func TestWorkerPool_JobsKosong(t *testing.T) {
	jobs := make(chan int)
	close(jobs) // langsung tutup

	hasilCh := WorkerPool(jobs, 3, func(n int) int { return n })
	if hasilCh == nil {
		t.Fatal("WorkerPool: channel tidak boleh nil")
	}

	var hasil []int
	for v := range hasilCh {
		hasil = append(hasil, v)
	}

	if len(hasil) != 0 {
		t.Errorf("WorkerPool jobs kosong: want 0 hasil, got %d", len(hasil))
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// TEST BAGIAN 8: PIPELINE
// ═══════════════════════════════════════════════════════════════════════════════

func TestSumber_Nilai(t *testing.T) {
	ch := Sumber(1, 2, 3, 4, 5)
	if ch == nil {
		t.Fatal("Sumber: channel tidak boleh nil")
	}

	var hasil []int
	for v := range ch {
		hasil = append(hasil, v)
	}

	want := []int{1, 2, 3, 4, 5}
	if len(hasil) != len(want) {
		t.Fatalf("Sumber: want %v, got %v", want, hasil)
	}
	for i, w := range want {
		if hasil[i] != w {
			t.Errorf("Sumber[%d]: want %d, got %d", i, w, hasil[i])
		}
	}
}

func TestSumber_Ditutup(t *testing.T) {
	ch := Sumber(1, 2)
	<-ch
	<-ch
	// Channel harus ditutup
	_, ok := <-ch
	if ok {
		t.Error("Sumber: channel harus ditutup setelah semua data terkirim")
	}
}

// ───────────────────────────────────────────────────────────────────────────

func TestKalikan_Nilai(t *testing.T) {
	src := Sumber(1, 2, 3, 4, 5)
	hasil := Kalikan(src, 3)

	if hasil == nil {
		t.Fatal("Kalikan: channel tidak boleh nil")
	}

	var got []int
	for v := range hasil {
		got = append(got, v)
	}

	want := []int{3, 6, 9, 12, 15}
	if len(got) != len(want) {
		t.Fatalf("Kalikan: want %v, got %v", want, got)
	}
	for i, w := range want {
		if got[i] != w {
			t.Errorf("Kalikan[%d]: want %d, got %d", i, w, got[i])
		}
	}
}

// ───────────────────────────────────────────────────────────────────────────

func TestTambahkan_Nilai(t *testing.T) {
	src := Sumber(1, 2, 3)
	hasil := Tambahkan(src, 10)

	if hasil == nil {
		t.Fatal("Tambahkan: channel tidak boleh nil")
	}

	var got []int
	for v := range hasil {
		got = append(got, v)
	}

	want := []int{11, 12, 13}
	if len(got) != len(want) {
		t.Fatalf("Tambahkan: want %v, got %v", want, got)
	}
	for i, w := range want {
		if got[i] != w {
			t.Errorf("Tambahkan[%d]: want %d, got %d", i, w, got[i])
		}
	}
}

func TestPipelineLengkap(t *testing.T) {
	// Sumber(1,2,3,4,5) → Kalikan(×2) → Tambahkan(+10)
	// Ekspektasi: 12, 14, 16, 18, 20
	src := Sumber(1, 2, 3, 4, 5)
	kali := Kalikan(src, 2)
	tambah := Tambahkan(kali, 10)

	var hasil []int
	for v := range tambah {
		hasil = append(hasil, v)
	}

	want := []int{12, 14, 16, 18, 20}
	if len(hasil) != len(want) {
		t.Fatalf("PipelineLengkap: want %v, got %v", want, hasil)
	}
	for i, w := range want {
		if hasil[i] != w {
			t.Errorf("PipelineLengkap[%d]: want %d, got %d", i, w, hasil[i])
		}
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// TEST BAGIAN 9: SEMAPHORE
// ═══════════════════════════════════════════════════════════════════════════════

func TestParalelTerbatas_SemuaJalan(t *testing.T) {
	var counter int64
	tugas := make([]func(), 30)
	for i := range tugas {
		tugas[i] = func() { atomic.AddInt64(&counter, 1) }
	}

	ParalelTerbatas(5, tugas...)

	if counter != 30 {
		t.Errorf("ParalelTerbatas: semua harus jalan, want 30, got %d", counter)
	}
}

func TestParalelTerbatas_BatasDipatuhi(t *testing.T) {
	maks := 3
	var aktif int64
	var puncak int64

	tugas := make([]func(), 30)
	for i := range tugas {
		tugas[i] = func() {
			current := atomic.AddInt64(&aktif, 1)
			// update puncak secara atomic
			for {
				old := atomic.LoadInt64(&puncak)
				if current <= old || atomic.CompareAndSwapInt64(&puncak, old, current) {
					break
				}
			}
			time.Sleep(20 * time.Millisecond)
			atomic.AddInt64(&aktif, -1)
		}
	}

	ParalelTerbatas(maks, tugas...)

	if puncak > int64(maks) {
		t.Errorf("ParalelTerbatas: maks %d goroutine, tapi puncak %d goroutine aktif bersamaan",
			maks, puncak)
	}
}

func TestParalelTerbatas_MenungguSemua(t *testing.T) {
	selesai := make([]bool, 10)
	tugas := make([]func(), 10)
	for i := range tugas {
		i := i
		tugas[i] = func() {
			time.Sleep(20 * time.Millisecond)
			selesai[i] = true
		}
	}

	ParalelTerbatas(3, tugas...)

	for i, s := range selesai {
		if !s {
			t.Errorf("ParalelTerbatas: tugas[%d] belum selesai saat return", i)
		}
	}
}
