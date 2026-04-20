package belajar

import (
	"context"
	"errors"
	"sort"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// ═══════════════════════════════════════════════════════════════════════════
// TEST: RunInParallel
// ═══════════════════════════════════════════════════════════════════════════

func TestRunInParallel_Basic(t *testing.T) {
	// Semua task harus jalan, counter harus == jumlah task
	var count int64
	tasks := make([]func(), 10)
	for i := range tasks {
		tasks[i] = func() {
			atomic.AddInt64(&count, 1) // atomic: aman tanpa mutex
		}
	}

	RunInParallel(tasks...)

	if count != 10 {
		t.Errorf("RunInParallel: want count=10, got %d", count)
	}
}

func TestRunInParallel_Empty(t *testing.T) {
	// Tidak boleh panik kalau tasks kosong
	RunInParallel()
}

func TestRunInParallel_WaitsForAll(t *testing.T) {
	// Memastikan RunInParallel BENAR-BENAR menunggu semua goroutine selesai
	// sebelum return. Bukan hanya "dijalankan" tapi belum selesai.
	done := make([]bool, 5)
	tasks := make([]func(), 5)
	for i := range tasks {
		i := i
		tasks[i] = func() {
			time.Sleep(20 * time.Millisecond)
			done[i] = true
		}
	}

	RunInParallel(tasks...)

	for i, d := range done {
		if !d {
			t.Errorf("RunInParallel: task[%d] belum selesai saat return", i)
		}
	}
}

// ═══════════════════════════════════════════════════════════════════════════
// TEST: AsyncDouble
// ═══════════════════════════════════════════════════════════════════════════

func TestAsyncDouble_Values(t *testing.T) {
	cases := []struct {
		input int
		want  int
	}{
		{0, 0},
		{1, 2},
		{21, 42},
		{-5, -10},
		{100, 200},
	}

	for _, c := range cases {
		ch := AsyncDouble(c.input)
		if ch == nil {
			t.Fatalf("AsyncDouble(%d): channel tidak boleh nil", c.input)
		}

		select {
		case got := <-ch:
			if got != c.want {
				t.Errorf("AsyncDouble(%d): want %d, got %d", c.input, c.want, got)
			}
		case <-time.After(1 * time.Second):
			t.Errorf("AsyncDouble(%d): timeout! goroutine tidak mengirim hasil", c.input)
		}
	}
}

func TestAsyncDouble_IsActuallyAsync(t *testing.T) {
	// AsyncDouble harus return SEGERA (tidak blokir pemanggil)
	start := time.Now()
	ch := AsyncDouble(7)
	elapsed := time.Since(start)

	if elapsed > 100*time.Millisecond {
		t.Errorf("AsyncDouble tidak async: elapsed %v, harusnya < 100ms", elapsed)
	}

	// Pastikan hasilnya benar
	if got := <-ch; got != 14 {
		t.Errorf("AsyncDouble(7): want 14, got %d", got)
	}
}

// ═══════════════════════════════════════════════════════════════════════════
// TEST: GatherResults
// ═══════════════════════════════════════════════════════════════════════════

func TestGatherResults_Basic(t *testing.T) {
	hasil := GatherResults(
		func() int { return 10 },
		func() int { return 20 },
		func() int { return 30 },
	)

	if len(hasil) != 3 {
		t.Fatalf("GatherResults: want len=3, got %d", len(hasil))
	}

	// Urutkan karena goroutine tidak deterministik
	sort.Ints(hasil)
	want := []int{10, 20, 30}
	for i, v := range want {
		if hasil[i] != v {
			t.Errorf("GatherResults[%d]: want %d, got %d", i, v, hasil[i])
		}
	}
}

func TestGatherResults_Empty(t *testing.T) {
	hasil := GatherResults()
	if len(hasil) != 0 {
		t.Errorf("GatherResults kosong: want len=0, got %d", len(hasil))
	}
}

func TestGatherResults_Single(t *testing.T) {
	hasil := GatherResults(func() int { return 99 })
	if len(hasil) != 1 || hasil[0] != 99 {
		t.Errorf("GatherResults single: want [99], got %v", hasil)
	}
}

func TestGatherResults_Parallel(t *testing.T) {
	// Verifikasi bahwa tasks berjalan paralel:
	// 5 tasks, masing-masing tidur 100ms.
	// Kalau paralel: total ~100ms. Kalau sequential: ~500ms.
	start := time.Now()

	GatherResults(
		func() int { time.Sleep(100 * time.Millisecond); return 1 },
		func() int { time.Sleep(100 * time.Millisecond); return 2 },
		func() int { time.Sleep(100 * time.Millisecond); return 3 },
		func() int { time.Sleep(100 * time.Millisecond); return 4 },
		func() int { time.Sleep(100 * time.Millisecond); return 5 },
	)

	elapsed := time.Since(start)
	if elapsed > 400*time.Millisecond {
		t.Errorf("GatherResults tidak berjalan paralel: elapsed %v", elapsed)
	}
}

// ═══════════════════════════════════════════════════════════════════════════
// TEST: WithTimeout
// ═══════════════════════════════════════════════════════════════════════════

func TestWithTimeout_Success(t *testing.T) {
	// Work selesai sebelum timeout
	hasil, ok := WithTimeout(func() int {
		time.Sleep(10 * time.Millisecond)
		return 42
	}, 500*time.Millisecond)

	if !ok {
		t.Fatal("WithTimeout: harusnya berhasil (ok=true), bukan timeout")
	}
	if hasil != 42 {
		t.Errorf("WithTimeout: want 42, got %d", hasil)
	}
}

func TestWithTimeout_Timeout(t *testing.T) {
	// Work lebih lambat dari timeout
	hasil, ok := WithTimeout(func() int {
		time.Sleep(500 * time.Millisecond)
		return 99
	}, 50*time.Millisecond)

	if ok {
		t.Fatal("WithTimeout: harusnya timeout (ok=false)")
	}
	if hasil != 0 {
		t.Errorf("WithTimeout timeout: want 0, got %d", hasil)
	}
}

// ═══════════════════════════════════════════════════════════════════════════
// TEST: TryReceive
// ═══════════════════════════════════════════════════════════════════════════

func TestTryReceive_Empty(t *testing.T) {
	ch := make(chan int, 1)
	v, ok := TryReceive(ch)
	if ok || v != 0 {
		t.Errorf("TryReceive channel kosong: want (0,false), got (%d,%v)", v, ok)
	}
}

func TestTryReceive_HasValue(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 99
	v, ok := TryReceive(ch)
	if !ok || v != 99 {
		t.Errorf("TryReceive ada nilai: want (99,true), got (%d,%v)", v, ok)
	}
}

func TestTryReceive_NonBlocking(t *testing.T) {
	// TryReceive tidak boleh blokir
	ch := make(chan int) // unbuffered, tidak ada pengirim

	start := time.Now()
	_, ok := TryReceive(ch)
	elapsed := time.Since(start)

	if ok {
		t.Error("TryReceive channel kosong harusnya return false")
	}
	if elapsed > 50*time.Millisecond {
		t.Errorf("TryReceive blokir! elapsed %v, harusnya < 50ms", elapsed)
	}
}

// ═══════════════════════════════════════════════════════════════════════════
// TEST: RunWithCancel
// ═══════════════════════════════════════════════════════════════════════════

func TestRunWithCancel_Normal(t *testing.T) {
	// Work selesai normal tanpa dibatalkan
	resultCh, cancel := RunWithCancel(func(ctx context.Context) int {
		time.Sleep(10 * time.Millisecond)
		return 42
	})
	defer cancel()

	if resultCh == nil {
		t.Fatal("RunWithCancel: resultCh tidak boleh nil")
	}

	select {
	case v := <-resultCh:
		if v != 42 {
			t.Errorf("RunWithCancel: want 42, got %d", v)
		}
	case <-time.After(1 * time.Second):
		t.Error("RunWithCancel: timeout menunggu hasil")
	}
}

func TestRunWithCancel_Cancelled(t *testing.T) {
	// Work mendengarkan ctx.Done() dan berhenti saat dibatalkan
	resultCh, cancel := RunWithCancel(func(ctx context.Context) int {
		select {
		case <-ctx.Done():
			return -1 // sinyal dibatalkan
		case <-time.After(10 * time.Second):
			return 999
		}
	})

	// Batalkan segera
	cancel()

	select {
	case v := <-resultCh:
		if v != -1 {
			t.Errorf("RunWithCancel dibatalkan: want -1, got %d", v)
		}
	case <-time.After(1 * time.Second):
		t.Error("RunWithCancel: goroutine tidak merespons pembatalan")
	}
}

// ═══════════════════════════════════════════════════════════════════════════
// TEST: BoundedParallel
// ═══════════════════════════════════════════════════════════════════════════

func TestBoundedParallel_AllTasksRun(t *testing.T) {
	var count int64
	tasks := make([]func(), 50)
	for i := range tasks {
		tasks[i] = func() {
			atomic.AddInt64(&count, 1)
		}
	}

	BoundedParallel(tasks, 5)

	if count != 50 {
		t.Errorf("BoundedParallel: semua task harus jalan, want 50, got %d", count)
	}
}

func TestBoundedParallel_RespectsLimit(t *testing.T) {
	// Verifikasi bahwa tidak pernah ada lebih dari maxConcurrent goroutine aktif
	maxConcurrent := 3
	var active int64
	var maxSeen int64

	var mu sync.Mutex
	tasks := make([]func(), 30)
	for i := range tasks {
		tasks[i] = func() {
			current := atomic.AddInt64(&active, 1)

			mu.Lock()
			if current > maxSeen {
				maxSeen = current
			}
			mu.Unlock()

			time.Sleep(20 * time.Millisecond)
			atomic.AddInt64(&active, -1)
		}
	}

	BoundedParallel(tasks, maxConcurrent)

	if maxSeen > int64(maxConcurrent) {
		t.Errorf("BoundedParallel: max concurrent %d, tapi pernah ada %d goroutine aktif",
			maxConcurrent, maxSeen)
	}
}

// ═══════════════════════════════════════════════════════════════════════════
// TEST: RunAndCollectErrors
// ═══════════════════════════════════════════════════════════════════════════

func TestRunAndCollectErrors_NoErrors(t *testing.T) {
	errs := RunAndCollectErrors(
		func() error { return nil },
		func() error { return nil },
		func() error { return nil },
	)

	if len(errs) != 0 {
		t.Errorf("RunAndCollectErrors: want 0 errors, got %d: %v", len(errs), errs)
	}
}

func TestRunAndCollectErrors_SomeErrors(t *testing.T) {
	errs := RunAndCollectErrors(
		func() error { return nil },
		func() error { return errors.New("gagal A") },
		func() error { return nil },
		func() error { return errors.New("gagal B") },
	)

	if len(errs) != 2 {
		t.Errorf("RunAndCollectErrors: want 2 errors, got %d: %v", len(errs), errs)
	}
}

func TestRunAndCollectErrors_AllErrors(t *testing.T) {
	errs := RunAndCollectErrors(
		func() error { return errors.New("err1") },
		func() error { return errors.New("err2") },
	)

	if len(errs) != 2 {
		t.Errorf("RunAndCollectErrors: want 2 errors, got %d", len(errs))
	}
}

func TestRunAndCollectErrors_Empty(t *testing.T) {
	errs := RunAndCollectErrors()
	if errs != nil && len(errs) != 0 {
		t.Errorf("RunAndCollectErrors kosong: want nil/empty, got %v", errs)
	}
}

// ═══════════════════════════════════════════════════════════════════════════
// TEST: StreamNumbers
// ═══════════════════════════════════════════════════════════════════════════

func TestStreamNumbers_Sequence(t *testing.T) {
	stop := make(chan struct{})
	numCh := StreamNumbers(stop)

	if numCh == nil {
		t.Fatal("StreamNumbers: channel tidak boleh nil")
	}

	// Ambil 5 angka pertama
	for want := 1; want <= 5; want++ {
		select {
		case got := <-numCh:
			if got != want {
				t.Errorf("StreamNumbers: want %d, got %d", want, got)
			}
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("StreamNumbers: timeout menunggu angka ke-%d", want)
		}
	}

	// Hentikan stream
	close(stop)
}

func TestStreamNumbers_StopsOnClose(t *testing.T) {
	stop := make(chan struct{})
	numCh := StreamNumbers(stop)

	// Baca sedikit
	<-numCh
	<-numCh

	// Stop
	close(stop)

	// Channel harus ditutup (for-range harus berhenti)
	timer := time.After(500 * time.Millisecond)
	draining := true
	for draining {
		select {
		case _, ok := <-numCh:
			if !ok {
				draining = false // channel ditutup dengan benar
			}
		case <-timer:
			t.Error("StreamNumbers: channel tidak ditutup setelah stop")
			draining = false
		}
	}
}

// ═══════════════════════════════════════════════════════════════════════════
// TEST: SingletonInit
// ═══════════════════════════════════════════════════════════════════════════

func TestSingletonInit_OnlyOnce(t *testing.T) {
	var count int
	inc := func() { count++ }

	var once sync.Once
	var wg sync.WaitGroup

	// 100 goroutine memanggil SingletonInit bersamaan
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			SingletonInit(&once, inc)
		}()
	}

	wg.Wait()

	if count != 1 {
		t.Errorf("SingletonInit: init harus dieksekusi tepat sekali, got count=%d", count)
	}
}

func TestSingletonInit_DifferentOnce(t *testing.T) {
	// Once yang berbeda → init bisa jalan lagi (independent)
	count := 0
	inc := func() { count++ }

	var once1 sync.Once
	var once2 sync.Once

	SingletonInit(&once1, inc) // count = 1
	SingletonInit(&once1, inc) // tidak jalan (once1 sudah dipakai)
	SingletonInit(&once2, inc) // count = 2 (once2 baru)

	if count != 2 {
		t.Errorf("SingletonInit: want count=2, got %d", count)
	}
}
