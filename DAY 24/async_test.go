package belajar

import (
	"context"
	"errors"
	"sort"
	"testing"
	"time"
)

// ═══════════════════════════════════════════════════════════════════════════════
// TEST BAGIAN 1: ERRGROUP PATTERN
// ═══════════════════════════════════════════════════════════════════════════════

func TestAsyncJalankanKumpulkanError_SemuaBerhasil(t *testing.T) {
	errs := JalankanKumpulkanError(
		func() error { return nil },
		func() error { return nil },
		func() error { return nil },
	)
	if len(errs) != 0 {
		t.Errorf("JalankanKumpulkanError: want 0 error, got %d: %v", len(errs), errs)
	}
}

func TestAsyncJalankanKumpulkanError_AdaYangGagal(t *testing.T) {
	errs := JalankanKumpulkanError(
		func() error { return nil },
		func() error { return errors.New("gagal A") },
		func() error { return nil },
		func() error { return errors.New("gagal B") },
	)
	if len(errs) != 2 {
		t.Errorf("JalankanKumpulkanError: want 2 error, got %d: %v", len(errs), errs)
	}
}

func TestAsyncJalankanKumpulkanError_SemuaGagal(t *testing.T) {
	errs := JalankanKumpulkanError(
		func() error { return errors.New("a") },
		func() error { return errors.New("b") },
		func() error { return errors.New("c") },
	)
	if len(errs) != 3 {
		t.Errorf("JalankanKumpulkanError: want 3 error, got %d", len(errs))
	}
}

func TestAsyncJalankanKumpulkanError_Paralel(t *testing.T) {
	// 5 tugas × 100ms → kalau paralel ~100ms, sequential ~500ms
	start := time.Now()
	JalankanKumpulkanError(
		func() error { time.Sleep(100 * time.Millisecond); return nil },
		func() error { time.Sleep(100 * time.Millisecond); return nil },
		func() error { time.Sleep(100 * time.Millisecond); return nil },
		func() error { time.Sleep(100 * time.Millisecond); return nil },
		func() error { time.Sleep(100 * time.Millisecond); return nil },
	)
	elapsed := time.Since(start)
	if elapsed > 400*time.Millisecond {
		t.Errorf("JalankanKumpulkanError tidak paralel: elapsed %v", elapsed)
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// TEST BAGIAN 2: FIRST-SUCCESS PATTERN
// ═══════════════════════════════════════════════════════════════════════════════

func TestAsyncAmbilPertamaBerhasil_AdaYangBerhasil(t *testing.T) {
	hasil, err := AmbilPertamaBerhasil(
		func() (int, error) { return 0, errors.New("timeout") },
		func() (int, error) { return 0, errors.New("not found") },
		func() (int, error) { return 42, nil },
	)
	if err != nil {
		t.Fatalf("AmbilPertamaBerhasil: unexpected error: %v", err)
	}
	if hasil != 42 {
		t.Errorf("AmbilPertamaBerhasil: want 42, got %d", hasil)
	}
}

func TestAsyncAmbilPertamaBerhasil_SemuaGagal(t *testing.T) {
	_, err := AmbilPertamaBerhasil(
		func() (int, error) { return 0, errors.New("gagal 1") },
		func() (int, error) { return 0, errors.New("gagal 2") },
		func() (int, error) { return 0, errors.New("gagal 3") },
	)
	if err == nil {
		t.Error("AmbilPertamaBerhasil: harus return error jika semua gagal")
	}
}

func TestAsyncAmbilPertamaBerhasil_PertamaBerhasil(t *testing.T) {
	hasil, err := AmbilPertamaBerhasil(
		func() (int, error) { return 99, nil },
		func() (int, error) { return 0, errors.New("gagal") },
	)
	if err != nil || hasil != 99 {
		t.Errorf("AmbilPertamaBerhasil: want (99, nil), got (%d, %v)", hasil, err)
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// TEST BAGIAN 3: CONTEXT TIMEOUT
// ═══════════════════════════════════════════════════════════════════════════════

func TestAsyncJalankanDenganContext_Berhasil(t *testing.T) {
	ctx := context.Background()
	hasil, err := JalankanDenganContext(ctx, func() int {
		return 42
	})
	if err != nil {
		t.Fatalf("JalankanDenganContext: unexpected error: %v", err)
	}
	if hasil != 42 {
		t.Errorf("JalankanDenganContext: want 42, got %d", hasil)
	}
}

func TestAsyncJalankanDenganContext_Timeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	hasil, err := JalankanDenganContext(ctx, func() int {
		time.Sleep(1 * time.Second)
		return 99
	})
	if err == nil {
		t.Error("JalankanDenganContext: harus return error saat timeout")
	}
	if hasil != 0 {
		t.Errorf("JalankanDenganContext timeout: want 0, got %d", hasil)
	}
}

func TestAsyncJalankanDenganContext_DibatalkanManual(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(30 * time.Millisecond)
		cancel()
	}()

	_, err := JalankanDenganContext(ctx, func() int {
		time.Sleep(1 * time.Second)
		return 99
	})
	if err == nil {
		t.Error("JalankanDenganContext: harus return error saat ctx dibatalkan")
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// TEST BAGIAN 4: PARALLEL MAP
// ═══════════════════════════════════════════════════════════════════════════════

func TestAsyncParallelMap_Kuadrat(t *testing.T) {
	hasil := ParallelMap([]int{1, 2, 3, 4, 5}, func(n int) int { return n * n })

	if len(hasil) != 5 {
		t.Fatalf("ParallelMap: want len=5, got %d", len(hasil))
	}
	want := []int{1, 4, 9, 16, 25}
	for i, w := range want {
		if hasil[i] != w {
			t.Errorf("ParallelMap[%d]: want %d, got %d", i, w, hasil[i])
		}
	}
}

func TestAsyncParallelMap_UrutanTerjaga(t *testing.T) {
	// Elemen pertama sengaja lambat — pastikan urutan tetap terjaga
	hasil := ParallelMap([]int{1, 2, 3, 4, 5}, func(n int) int {
		if n == 1 {
			time.Sleep(50 * time.Millisecond)
		}
		return n * 100
	})

	want := []int{100, 200, 300, 400, 500}
	for i, w := range want {
		if hasil[i] != w {
			t.Errorf("ParallelMap urutan[%d]: want %d, got %d", i, w, hasil[i])
		}
	}
}

func TestAsyncParallelMap_Kosong(t *testing.T) {
	hasil := ParallelMap([]int{}, func(n int) int { return n })
	if len(hasil) != 0 {
		t.Errorf("ParallelMap kosong: want len=0, got %d", len(hasil))
	}
}

func TestAsyncParallelMap_Paralel(t *testing.T) {
	// 5 elemen × 100ms → kalau paralel ~100ms, sequential ~500ms
	start := time.Now()
	ParallelMap([]int{1, 2, 3, 4, 5}, func(n int) int {
		time.Sleep(100 * time.Millisecond)
		return n
	})
	elapsed := time.Since(start)
	if elapsed > 400*time.Millisecond {
		t.Errorf("ParallelMap tidak paralel: elapsed %v", elapsed)
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// TEST BAGIAN 5: SCATTER-GATHER
// ═══════════════════════════════════════════════════════════════════════════════

func TestAsyncScatterGather_SemuaHasil(t *testing.T) {
	hasil := ScatterGather(
		func() int { return 100 },
		func() int { return 50 },
		func() int { return 75 },
	)

	if len(hasil) != 3 {
		t.Fatalf("ScatterGather: want 3 hasil, got %d", len(hasil))
	}
	sort.Ints(hasil)
	want := []int{50, 75, 100}
	for i, w := range want {
		if hasil[i] != w {
			t.Errorf("ScatterGather[%d]: want %d, got %d", i, w, hasil[i])
		}
	}
}

func TestAsyncScatterGather_Paralel(t *testing.T) {
	// 4 sumber × 100ms → kalau paralel ~100ms, sequential ~400ms
	start := time.Now()
	ScatterGather(
		func() int { time.Sleep(100 * time.Millisecond); return 1 },
		func() int { time.Sleep(100 * time.Millisecond); return 2 },
		func() int { time.Sleep(100 * time.Millisecond); return 3 },
		func() int { time.Sleep(100 * time.Millisecond); return 4 },
	)
	elapsed := time.Since(start)
	if elapsed > 350*time.Millisecond {
		t.Errorf("ScatterGather tidak paralel: elapsed %v", elapsed)
	}
}

func TestAsyncScatterGather_Kosong(t *testing.T) {
	hasil := ScatterGather()
	if len(hasil) != 0 {
		t.Errorf("ScatterGather kosong: want len=0, got %d", len(hasil))
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// TEST BAGIAN 6: ASYNC PIPELINE
// ═══════════════════════════════════════════════════════════════════════════════

func TestAsyncSumberAsync_Nilai(t *testing.T) {
	ctx := context.Background()
	ch := SumberAsync(ctx, 1, 2, 3, 4, 5)
	if ch == nil {
		t.Fatal("SumberAsync: channel tidak boleh nil")
	}

	var hasil []int
	for v := range ch {
		hasil = append(hasil, v)
	}

	want := []int{1, 2, 3, 4, 5}
	if len(hasil) != len(want) {
		t.Fatalf("SumberAsync: want %v, got %v", want, hasil)
	}
	for i, w := range want {
		if hasil[i] != w {
			t.Errorf("SumberAsync[%d]: want %d, got %d", i, w, hasil[i])
		}
	}
}

func TestAsyncSumberAsync_DibatalkanDiTengah(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	// Batalkan setelah sedikit delay
	go func() {
		time.Sleep(20 * time.Millisecond)
		cancel()
	}()

	// Channel dengan data tak terbatas (angka besar, pasti belum selesai)
	ch := SumberAsync(ctx, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// Drain channel — harus berhenti (tidak deadlock) setelah ctx dibatalkan
	timer := time.After(500 * time.Millisecond)
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				return // channel ditutup — sukses
			}
		case <-timer:
			t.Error("SumberAsync: tidak berhenti setelah ctx dibatalkan (kemungkinan deadlock)")
			return
		}
	}
}

func TestAsyncProsesAsync_Nilai(t *testing.T) {
	ctx := context.Background()
	src := SumberAsync(ctx, 1, 2, 3)
	doubled := ProsesAsync(ctx, src, func(n int) int { return n * 2 })

	if doubled == nil {
		t.Fatal("ProsesAsync: channel tidak boleh nil")
	}

	var hasil []int
	for v := range doubled {
		hasil = append(hasil, v)
	}

	want := []int{2, 4, 6}
	if len(hasil) != len(want) {
		t.Fatalf("ProsesAsync: want %v, got %v", want, hasil)
	}
	for i, w := range want {
		if hasil[i] != w {
			t.Errorf("ProsesAsync[%d]: want %d, got %d", i, w, hasil[i])
		}
	}
}

func TestAsyncPipelineLengkap(t *testing.T) {
	// SumberAsync(1,2,3,4,5) → ProsesAsync(×2) → ProsesAsync(+10)
	// Ekspektasi: 12, 14, 16, 18, 20
	ctx := context.Background()
	src := SumberAsync(ctx, 1, 2, 3, 4, 5)
	doubled := ProsesAsync(ctx, src, func(n int) int { return n * 2 })
	added := ProsesAsync(ctx, doubled, func(n int) int { return n + 10 })

	var hasil []int
	for v := range added {
		hasil = append(hasil, v)
	}

	want := []int{12, 14, 16, 18, 20}
	if len(hasil) != len(want) {
		t.Fatalf("PipelineAsync: want %v, got %v", want, hasil)
	}
	for i, w := range want {
		if hasil[i] != w {
			t.Errorf("PipelineAsync[%d]: want %d, got %d", i, w, hasil[i])
		}
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// TEST BAGIAN 7: WORKER POOL DENGAN ERROR
// ═══════════════════════════════════════════════════════════════════════════════

func TestAsyncWorkerPoolDenganError_SemuaBerhasil(t *testing.T) {
	jobs := make(chan int, 3)
	jobs <- 2
	jobs <- 3
	jobs <- 4
	close(jobs)

	hasilCh := WorkerPoolDenganError(jobs, 2, func(n int) (int, error) {
		return n * n, nil
	})
	if hasilCh == nil {
		t.Fatal("WorkerPoolDenganError: channel tidak boleh nil")
	}

	var sukses []int
	var gagal []error
	for h := range hasilCh {
		if h.Err != nil {
			gagal = append(gagal, h.Err)
		} else {
			sukses = append(sukses, h.Nilai)
		}
	}

	if len(gagal) != 0 {
		t.Errorf("WorkerPoolDenganError: want 0 error, got %d", len(gagal))
	}
	if len(sukses) != 3 {
		t.Errorf("WorkerPoolDenganError: want 3 sukses, got %d", len(sukses))
	}
	sort.Ints(sukses)
	want := []int{4, 9, 16}
	for i, w := range want {
		if sukses[i] != w {
			t.Errorf("WorkerPoolDenganError[%d]: want %d, got %d", i, w, sukses[i])
		}
	}
}

func TestAsyncWorkerPoolDenganError_AdaYangGagal(t *testing.T) {
	jobs := make(chan int, 4)
	for _, v := range []int{1, 2, 3, 4} {
		jobs <- v
	}
	close(jobs)

	hasilCh := WorkerPoolDenganError(jobs, 2, func(n int) (int, error) {
		if n%2 == 0 {
			return 0, errors.New("angka genap ditolak")
		}
		return n * 10, nil
	})

	var sukses, gagal int
	for h := range hasilCh {
		if h.Err != nil {
			gagal++
		} else {
			sukses++
		}
	}

	if sukses != 2 {
		t.Errorf("WorkerPoolDenganError: want 2 sukses, got %d", sukses)
	}
	if gagal != 2 {
		t.Errorf("WorkerPoolDenganError: want 2 gagal, got %d", gagal)
	}
}

func TestAsyncWorkerPoolDenganError_ChannelDitutup(t *testing.T) {
	jobs := make(chan int, 2)
	jobs <- 1
	jobs <- 2
	close(jobs)

	hasilCh := WorkerPoolDenganError(jobs, 2, func(n int) (int, error) {
		return n, nil
	})

	// Drain dan pastikan channel tertutup (tidak deadlock)
	timer := time.After(1 * time.Second)
	for {
		select {
		case _, ok := <-hasilCh:
			if !ok {
				return // sukses: channel ditutup
			}
		case <-timer:
			t.Error("WorkerPoolDenganError: channel tidak ditutup setelah semua pekerjaan selesai")
			return
		}
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// TEST BAGIAN 8: THROTTLED ASYNC
// ═══════════════════════════════════════════════════════════════════════════════

func TestAsyncJalankanTerbatas_Nilai(t *testing.T) {
	hasil := JalankanTerbatas([]int{1, 2, 3, 4, 5}, func(n int) int {
		return n * 10
	}, 10) // rate tinggi agar test tidak terlalu lama

	if len(hasil) != 5 {
		t.Fatalf("JalankanTerbatas: want 5 hasil, got %d", len(hasil))
	}
	sort.Ints(hasil)
	want := []int{10, 20, 30, 40, 50}
	for i, w := range want {
		if hasil[i] != w {
			t.Errorf("JalankanTerbatas[%d]: want %d, got %d", i, w, hasil[i])
		}
	}
}

func TestAsyncJalankanTerbatas_TerThrottle(t *testing.T) {
	// 4 tugas, rate 2/detik → minimal ~1.5 detik, tidak bisa <1 detik
	start := time.Now()
	JalankanTerbatas([]int{1, 2, 3, 4}, func(n int) int { return n }, 2)
	elapsed := time.Since(start)

	// Dengan rate 2/detik dan 4 tugas: tugas 1 & 2 di detik ke-0,
	// tugas 3 & 4 di detik ke-1 → minimal ~1 detik
	if elapsed < 900*time.Millisecond {
		t.Errorf("JalankanTerbatas: terlalu cepat (%v) — rate limiter tidak bekerja", elapsed)
	}
}

func TestAsyncJalankanTerbatas_Kosong(t *testing.T) {
	hasil := JalankanTerbatas([]int{}, func(n int) int { return n }, 5)
	if len(hasil) != 0 {
		t.Errorf("JalankanTerbatas kosong: want len=0, got %d", len(hasil))
	}
}
