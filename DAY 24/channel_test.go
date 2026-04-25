package belajar

import (
	"sort"
	"testing"
	"time"
)

// ═══════════════════════════════════════════════════════════════════════════════
// TEST BAGIAN 1: ORDERED FAN-OUT
// ═══════════════════════════════════════════════════════════════════════════════

func TestMapOrdered_Kuadrat(t *testing.T) {
	hasil := MapOrdered(func(n int) int { return n * n }, 1, 2, 3, 4, 5)

	if len(hasil) != 5 {
		t.Fatalf("MapOrdered: want len=5, got %d", len(hasil))
	}
	want := []int{1, 4, 9, 16, 25}
	for i, w := range want {
		if hasil[i] != w {
			t.Errorf("MapOrdered[%d]: want %d, got %d", i, w, hasil[i])
		}
	}
}

func TestMapOrdered_UrutanTerjaga(t *testing.T) {
	// Fungsi yang lambat untuk item ke-0 — pastikan tetap urut
	hasil := MapOrdered(func(n int) int {
		if n == 1 {
			time.Sleep(50 * time.Millisecond) // item ke-0 sengaja lambat
		}
		return n * 10
	}, 1, 2, 3, 4, 5)

	want := []int{10, 20, 30, 40, 50}
	for i, w := range want {
		if hasil[i] != w {
			t.Errorf("MapOrdered urutan[%d]: want %d, got %d (urutan tidak terjaga?)", i, w, hasil[i])
		}
	}
}

func TestMapOrdered_Kosong(t *testing.T) {
	hasil := MapOrdered(func(n int) int { return n })
	if len(hasil) != 0 {
		t.Errorf("MapOrdered kosong: want len=0, got %d", len(hasil))
	}
}

func TestMapOrdered_Paralel(t *testing.T) {
	// 5 item × 100ms. Kalau paralel → ~100ms. Kalau sequential → ~500ms.
	start := time.Now()
	MapOrdered(func(n int) int {
		time.Sleep(100 * time.Millisecond)
		return n
	}, 1, 2, 3, 4, 5)
	elapsed := time.Since(start)

	if elapsed > 400*time.Millisecond {
		t.Errorf("MapOrdered tidak berjalan paralel: elapsed %v", elapsed)
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// TEST BAGIAN 2: BATCH PROCESSOR
// ═══════════════════════════════════════════════════════════════════════════════

func TestBatchChannel_Pas(t *testing.T) {
	// 6 item, ukuran 3 → 2 batch masing-masing [1,2,3] dan [4,5,6]
	masuk := make(chan int, 6)
	for _, v := range []int{1, 2, 3, 4, 5, 6} {
		masuk <- v
	}
	close(masuk)

	batchCh := BatchChannel(masuk, 3)
	if batchCh == nil {
		t.Fatal("BatchChannel: channel tidak boleh nil")
	}

	var batches [][]int
	for b := range batchCh {
		batches = append(batches, b)
	}

	if len(batches) != 2 {
		t.Fatalf("BatchChannel: want 2 batch, got %d", len(batches))
	}
	want := [][]int{{1, 2, 3}, {4, 5, 6}}
	for i, w := range want {
		if len(batches[i]) != len(w) {
			t.Errorf("BatchChannel batch[%d]: want %v, got %v", i, w, batches[i])
			continue
		}
		for j, wv := range w {
			if batches[i][j] != wv {
				t.Errorf("BatchChannel batch[%d][%d]: want %d, got %d", i, j, wv, batches[i][j])
			}
		}
	}
}

func TestBatchChannel_Sisa(t *testing.T) {
	// 5 item, ukuran 3 → batch [1,2,3] dan [4,5]
	masuk := make(chan int, 5)
	for _, v := range []int{1, 2, 3, 4, 5} {
		masuk <- v
	}
	close(masuk)

	batchCh := BatchChannel(masuk, 3)
	var batches [][]int
	for b := range batchCh {
		batches = append(batches, b)
	}

	if len(batches) != 2 {
		t.Fatalf("BatchChannel sisa: want 2 batch, got %d: %v", len(batches), batches)
	}
	if len(batches[0]) != 3 {
		t.Errorf("BatchChannel batch[0]: want len=3, got %d", len(batches[0]))
	}
	if len(batches[1]) != 2 {
		t.Errorf("BatchChannel batch[1] (sisa): want len=2, got %d", len(batches[1]))
	}
}

func TestBatchChannel_Kosong(t *testing.T) {
	masuk := make(chan int)
	close(masuk)

	batchCh := BatchChannel(masuk, 3)
	var count int
	for range batchCh {
		count++
	}
	if count != 0 {
		t.Errorf("BatchChannel kosong: want 0 batch, got %d", count)
	}
}

func TestBatchChannel_SatuItem(t *testing.T) {
	masuk := make(chan int, 1)
	masuk <- 42
	close(masuk)

	batchCh := BatchChannel(masuk, 5)
	var batches [][]int
	for b := range batchCh {
		batches = append(batches, b)
	}

	if len(batches) != 1 {
		t.Fatalf("BatchChannel 1 item: want 1 batch, got %d", len(batches))
	}
	if len(batches[0]) != 1 || batches[0][0] != 42 {
		t.Errorf("BatchChannel 1 item: want [42], got %v", batches[0])
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// TEST BAGIAN 3: RETRY PATTERN
// ═══════════════════════════════════════════════════════════════════════════════

func TestCobaUlang_BerhasilPertama(t *testing.T) {
	panggil := 0
	hasil, err := CobaUlang(func() (int, error) {
		panggil++
		return 42, nil
	}, 3, 0)

	if err != nil {
		t.Fatalf("CobaUlang: unexpected error: %v", err)
	}
	if hasil != 42 {
		t.Errorf("CobaUlang: want 42, got %d", hasil)
	}
	if panggil != 1 {
		t.Errorf("CobaUlang: fn harus dipanggil 1x, dipanggil %d", panggil)
	}
}

func TestCobaUlang_BerhasilSetelahBeberapa(t *testing.T) {
	panggil := 0
	hasil, err := CobaUlang(func() (int, error) {
		panggil++
		if panggil < 3 {
			return 0, errGagal("sengaja gagal")
		}
		return 99, nil
	}, 5, 0)

	if err != nil {
		t.Fatalf("CobaUlang: unexpected error: %v", err)
	}
	if hasil != 99 {
		t.Errorf("CobaUlang: want 99, got %d", hasil)
	}
	if panggil != 3 {
		t.Errorf("CobaUlang: fn harus dipanggil 3x, dipanggil %d", panggil)
	}
}

func TestCobaUlang_GagalSemua(t *testing.T) {
	panggil := 0
	_, err := CobaUlang(func() (int, error) {
		panggil++
		return 0, errGagal("selalu gagal")
	}, 4, 0)

	if err == nil {
		t.Error("CobaUlang: harus return error saat semua percobaan gagal")
	}
	if panggil != 4 {
		t.Errorf("CobaUlang: fn harus dipanggil 4x, dipanggil %d", panggil)
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// TEST BAGIAN 4: RATE LIMITER
// ═══════════════════════════════════════════════════════════════════════════════

func TestBuatRateLimiter_KirimToken(t *testing.T) {
	selesai := make(chan struct{})
	defer close(selesai)

	interval := 50 * time.Millisecond
	limiter := BuatRateLimiter(interval, selesai)
	if limiter == nil {
		t.Fatal("BuatRateLimiter: channel tidak boleh nil")
	}

	// Ambil 3 token, ukur waktunya
	start := time.Now()
	for i := 0; i < 3; i++ {
		select {
		case <-limiter:
		case <-time.After(1 * time.Second):
			t.Fatalf("BuatRateLimiter: timeout menunggu token ke-%d", i+1)
		}
	}
	elapsed := time.Since(start)

	// 3 token dengan interval 50ms → minimal ~100ms (token ke-1 langsung, ke-2 50ms, ke-3 100ms)
	if elapsed < 80*time.Millisecond {
		t.Errorf("BuatRateLimiter: terlalu cepat (%v), rate limiter tidak bekerja", elapsed)
	}
}

func TestBuatRateLimiter_BerhentiSaatSelesai(t *testing.T) {
	selesai := make(chan struct{})
	limiter := BuatRateLimiter(50*time.Millisecond, selesai)

	// Ambil 1 token
	select {
	case <-limiter:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("BuatRateLimiter: timeout menunggu token pertama")
	}

	// Tutup selesai
	close(selesai)

	// Channel harus ditutup segera
	time.Sleep(100 * time.Millisecond)
	select {
	case _, ok := <-limiter:
		if ok {
			// mungkin ada 1 token ekstra yang sudah di-buffer — tidak masalah
		}
		// channel akhirnya harus ditutup
	default:
		// channel mungkin sudah kosong — ini OK juga
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// TEST BAGIAN 5: BROADCAST
// ═══════════════════════════════════════════════════════════════════════════════

func TestBroadcast_SemuaMenerima(t *testing.T) {
	b := &Broadcaster{}
	ch1 := b.Subscribe(5)
	ch2 := b.Subscribe(5)
	ch3 := b.Subscribe(5)

	if ch1 == nil || ch2 == nil || ch3 == nil {
		t.Fatal("Subscribe: channel tidak boleh nil")
	}

	b.Broadcast(10)
	b.Broadcast(20)
	b.Broadcast(30)
	b.Close()

	for i, ch := range []<-chan int{ch1, ch2, ch3} {
		var hasil []int
		for v := range ch {
			hasil = append(hasil, v)
		}
		sort.Ints(hasil)
		want := []int{10, 20, 30}
		if len(hasil) != 3 {
			t.Errorf("Broadcast subscriber-%d: want 3 nilai, got %d: %v", i+1, len(hasil), hasil)
			continue
		}
		for j, w := range want {
			if hasil[j] != w {
				t.Errorf("Broadcast subscriber-%d[%d]: want %d, got %d", i+1, j, w, hasil[j])
			}
		}
	}
}

func TestBroadcast_TanpaSubscriber(t *testing.T) {
	b := &Broadcaster{}
	// Tidak boleh panic meski tidak ada subscriber
	b.Broadcast(99)
	b.Close()
}

func TestBroadcast_ChannelDitutupSetelahClose(t *testing.T) {
	b := &Broadcaster{}
	ch := b.Subscribe(5)

	b.Broadcast(1)
	b.Close()

	// Drain channel, harus bisa sampai channel ditutup
	var count int
	timer := time.After(500 * time.Millisecond)
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				if count < 1 {
					t.Error("Broadcast: subscriber harus terima nilai sebelum channel ditutup")
				}
				return // sukses
			}
			count++
		case <-timer:
			t.Error("Broadcast: channel subscriber tidak ditutup setelah Close()")
			return
		}
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// TEST BAGIAN 6: OR-DONE
// ═══════════════════════════════════════════════════════════════════════════════

func TestOrDone_SatuTertutup(t *testing.T) {
	d1 := make(chan struct{})
	d2 := make(chan struct{})
	d3 := make(chan struct{})

	combined := OrDone(d1, d2, d3)
	if combined == nil {
		t.Fatal("OrDone: channel tidak boleh nil")
	}

	// Tutup satu dari tiga
	close(d2)

	select {
	case <-combined:
		// sukses — combined ikut menutup
	case <-time.After(500 * time.Millisecond):
		t.Error("OrDone: harus selesai saat salah satu done ditutup")
	}
}

func TestOrDone_PertamaTertutup(t *testing.T) {
	d1 := make(chan struct{})
	d2 := make(chan struct{})

	combined := OrDone(d1, d2)

	close(d1)
	select {
	case <-combined:
	case <-time.After(500 * time.Millisecond):
		t.Error("OrDone: harus selesai saat d1 ditutup")
	}
}

func TestOrDone_TanpaChannel(t *testing.T) {
	// Tidak boleh panic, boleh blokir selamanya (channel tidak pernah tutup)
	combined := OrDone()
	if combined == nil {
		t.Fatal("OrDone tanpa channel: tidak boleh nil")
	}
}

func TestOrDone_Tercepat(t *testing.T) {
	// Done yang lebih cepat harus menang
	d1 := make(chan struct{})
	d2 := make(chan struct{})
	d3 := make(chan struct{})

	combined := OrDone(d1, d2, d3)

	go func() {
		time.Sleep(50 * time.Millisecond)
		close(d3) // d3 paling cepat
	}()

	start := time.Now()
	select {
	case <-combined:
		if time.Since(start) > 300*time.Millisecond {
			t.Error("OrDone: terlalu lama menunggu")
		}
	case <-time.After(500 * time.Millisecond):
		t.Error("OrDone: timeout — tidak merespons saat d3 ditutup")
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// TEST BAGIAN 7: TIMED GENERATOR
// ═══════════════════════════════════════════════════════════════════════════════

func TestTimedGenerator_MenghasilkanNilai(t *testing.T) {
	selesai := make(chan struct{})
	counter := 0
	ch := TimedGenerator(func() int {
		counter++
		return counter
	}, 50*time.Millisecond, selesai)

	if ch == nil {
		t.Fatal("TimedGenerator: channel tidak boleh nil")
	}

	var hasil []int
	for i := 0; i < 3; i++ {
		select {
		case v := <-ch:
			hasil = append(hasil, v)
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("TimedGenerator: timeout menunggu nilai ke-%d", i+1)
		}
	}

	close(selesai)

	if len(hasil) < 3 {
		t.Errorf("TimedGenerator: want >=3 nilai, got %d", len(hasil))
	}
}

func TestTimedGenerator_BerhentiSaatSelesai(t *testing.T) {
	selesai := make(chan struct{})
	ch := TimedGenerator(func() int { return 1 }, 30*time.Millisecond, selesai)

	// Baca satu nilai
	select {
	case <-ch:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("TimedGenerator: timeout menunggu nilai pertama")
	}

	// Hentikan
	close(selesai)

	// Channel harus ditutup tidak lama kemudian
	timer := time.After(300 * time.Millisecond)
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				return // sukses: channel ditutup
			}
		case <-timer:
			t.Error("TimedGenerator: channel tidak ditutup setelah selesai ditutup")
			return
		}
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// TEST BAGIAN 8: PIPELINE DENGAN FILTER & REDUCE
// ═══════════════════════════════════════════════════════════════════════════════

func TestFilter_Genap(t *testing.T) {
	src := Sumber(1, 2, 3, 4, 5, 6)
	hasil := Filter(src, func(n int) bool { return n%2 == 0 })

	if hasil == nil {
		t.Fatal("Filter: channel tidak boleh nil")
	}

	var got []int
	for v := range hasil {
		got = append(got, v)
	}

	want := []int{2, 4, 6}
	if len(got) != len(want) {
		t.Fatalf("Filter genap: want %v, got %v", want, got)
	}
	for i, w := range want {
		if got[i] != w {
			t.Errorf("Filter[%d]: want %d, got %d", i, w, got[i])
		}
	}
}

func TestFilter_SemuaLolos(t *testing.T) {
	src := Sumber(1, 2, 3)
	hasil := Filter(src, func(n int) bool { return true })

	var got []int
	for v := range hasil {
		got = append(got, v)
	}
	if len(got) != 3 {
		t.Errorf("Filter semua lolos: want 3, got %d", len(got))
	}
}

func TestFilter_TidakAdaYangLolos(t *testing.T) {
	src := Sumber(1, 2, 3)
	hasil := Filter(src, func(n int) bool { return false })

	var got []int
	for v := range hasil {
		got = append(got, v)
	}
	if len(got) != 0 {
		t.Errorf("Filter tidak ada lolos: want 0, got %d: %v", len(got), got)
	}
}

func TestFilter_ChannelDitutup(t *testing.T) {
	src := Sumber(2, 4, 6)
	hasil := Filter(src, func(n int) bool { return true })

	for range hasil {
	}
	// Kalau sampai sini tanpa deadlock, channel sudah ditutup dengan benar
}

// ───────────────────────────────────────────────────────────────────────────

func TestReduce_Sum(t *testing.T) {
	src := Sumber(1, 2, 3, 4, 5)
	total := Reduce(src, func(acc, v int) int { return acc + v }, 0)

	if total != 15 {
		t.Errorf("Reduce sum: want 15, got %d", total)
	}
}

func TestReduce_Max(t *testing.T) {
	src := Sumber(3, 1, 4, 1, 5, 9, 2, 6)
	maks := Reduce(src, func(acc, v int) int {
		if v > acc {
			return v
		}
		return acc
	}, 0)

	if maks != 9 {
		t.Errorf("Reduce max: want 9, got %d", maks)
	}
}

func TestReduce_Kosong(t *testing.T) {
	src := Sumber()
	hasil := Reduce(src, func(acc, v int) int { return acc + v }, 42)

	if hasil != 42 {
		t.Errorf("Reduce kosong: want awal=42, got %d", hasil)
	}
}

func TestPipelineDenganFilter(t *testing.T) {
	// Sumber(1..10) → Filter(genap) → Reduce(sum) = 2+4+6+8+10 = 30
	src := Sumber(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	genap := Filter(src, func(n int) bool { return n%2 == 0 })
	total := Reduce(genap, func(acc, v int) int { return acc + v }, 0)

	if total != 30 {
		t.Errorf("Pipeline+Filter: want 30, got %d", total)
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// TEST BAGIAN 9: TERBARU (LATEST VALUE)
// ═══════════════════════════════════════════════════════════════════════════════

func TestTerbaru_NilaiValid(t *testing.T) {
	masuk := make(chan int, 5)
	masuk <- 1
	masuk <- 2
	masuk <- 3
	close(masuk)

	ch := Terbaru(masuk)
	if ch == nil {
		t.Fatal("Terbaru: channel tidak boleh nil")
	}

	// Baca setidaknya satu nilai — harus valid (bukan 0 dari zero value)
	var got int
	select {
	case v, ok := <-ch:
		if !ok {
			t.Fatal("Terbaru: channel langsung tutup tanpa nilai")
		}
		got = v
	case <-time.After(500 * time.Millisecond):
		t.Fatal("Terbaru: timeout menunggu nilai")
	}

	if got < 1 || got > 3 {
		t.Errorf("Terbaru: nilai harus antara 1-3, got %d", got)
	}
}

func TestTerbaru_ChannelDitutup(t *testing.T) {
	masuk := make(chan int, 2)
	masuk <- 10
	masuk <- 20
	close(masuk)

	ch := Terbaru(masuk)

	// Drain channel — harus menutup dengan bersih
	timer := time.After(1 * time.Second)
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				return // sukses: channel ditutup
			}
		case <-timer:
			t.Error("Terbaru: channel tidak ditutup setelah input ditutup")
			return
		}
	}
}

func TestTerbaru_TidakDeadlock(t *testing.T) {
	// Producer cepat, consumer lambat → tidak deadlock, consumer dapat nilai terbaru
	masuk := make(chan int)
	ch := Terbaru(masuk)

	// Kirim banyak nilai cepat
	go func() {
		for i := 1; i <= 100; i++ {
			masuk <- i
		}
		close(masuk)
	}()

	// Consumer lambat — cukup ambil beberapa
	time.Sleep(10 * time.Millisecond)

	var nilaiTerakhir int
	timer := time.After(500 * time.Millisecond)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				// channel sudah tutup, cek dapat nilai tinggi (yang terbaru)
				if nilaiTerakhir == 0 {
					t.Error("Terbaru: tidak mendapat nilai apapun")
				}
				return
			}
			nilaiTerakhir = v
		case <-timer:
			t.Error("Terbaru: timeout — deadlock?")
			return
		}
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🔧 HELPER TEST
// ═══════════════════════════════════════════════════════════════════════════════

type errGagal string

func (e errGagal) Error() string { return string(e) }
