package belajar

// ==================== DAY 27: CONTEXT — CANCELLATION & TIMEOUT ====================
//
// 🎯 FOKUS HARI INI:
//   Menguasai context.Context sebagai cara idiomatis Go untuk:
//   - Membatalkan (cancel) goroutine dari luar
//   - Memberi deadline / timeout pada operasi
//   - Meneruskan sinyal berhenti lewat seluruh pipeline
//
//   Jalankan test:
//     cd "DAY 27"
//     go test ./... -v
//
// ═══════════════════════════════════════════════════════════════════════════════
// 🧠 KENAPA CONTEXT?
// ═══════════════════════════════════════════════════════════════════════════════
//
// Day 26 kamu pakai `done <-chan struct{}` untuk sinyal berhenti.
// Itu sudah bagus — tapi hanya bisa "cancel", tidak bisa membawa:
//   - deadline (kapan harus selesai)
//   - informasi penyebab pembatalan
//   - nilai yang diteruskan ke seluruh call stack (request-scoped value)
//
// context.Context menggabungkan semua itu dalam satu interface:
//
//   type Context interface {
//       Done() <-chan struct{}    ← sama seperti done channel Day 26!
//       Err()  error             ← kenapa selesai: Canceled / DeadlineExceeded
//       Deadline() (time.Time, bool)
//       Value(key any) any
//   }
//
// POLA UMUM dalam goroutine:
//
//   select {
//   case <-ctx.Done():
//       return ctx.Err()   // atau return saja
//   case ch <- v:
//   }
//
// ═══════════════════════════════════════════════════════════════════════════════
// 📦 TIGA CONSTRUCTOR UTAMA
// ═══════════════════════════════════════════════════════════════════════════════
//
//   ctx, cancel := context.WithCancel(parent)
//       → batalkan secara manual dengan memanggil cancel()
//
//   ctx, cancel := context.WithTimeout(parent, dur)
//       → otomatis cancel setelah dur; atau manual dengan cancel()
//
//   ctx, cancel := context.WithDeadline(parent, t)
//       → otomatis cancel pada waktu t; atau manual dengan cancel()
//
//   ATURAN WAJIB: selalu defer cancel() setelah membuat context!
//   (mencegah goroutine / timer leak)
//
// ═══════════════════════════════════════════════════════════════════════════════

import (
	"context"
	"sync"
	"time"
)

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 1: GENERATOR DENGAN CONTEXT
// ═══════════════════════════════════════════════════════════════════════════════
//
// Generator "naik" yang bisa dihentikan lewat context — mirip GeneratorDenganDone
// di Day 26, tapi kini menggunakan ctx.Done() bukan done chan struct{}.
//

// GeneratorCtx menghasilkan angka naik (1, 2, 3, ...) tanpa batas.
// Berhenti dan menutup channel output saat ctx dibatalkan.
//
// Contoh:
//
//	ctx, cancel := context.WithCancel(context.Background())
//	gen := GeneratorCtx(ctx)
//	fmt.Println(<-gen) // 1
//	fmt.Println(<-gen) // 2
//	cancel()           // sinyal berhenti
//	// gen tertutup setelah ini
//
// Hint:
//   - Buat channel output (unbuffered).
//   - Goroutine: for-loop i=1; select { case <-ctx.Done(): return; case ch <- i: i++ }.
//   - defer close(ch) di awal goroutine.
func GeneratorCtx(ctx context.Context) <-chan int {
	// TODO: implementasi di sini
	ch := make(chan int)

	go func() {
		defer close(ch)
		i := 1
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
				i++
			}
		}

	}()
	return ch
}

// GenRangeCtx menghasilkan angka dari `dari` sampai `sampai` (inklusif).
// Berhenti lebih awal jika ctx dibatalkan sebelum selesai.
// Channel output selalu ditutup (baik selesai normal maupun karena cancel).
//
// Contoh:
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//	for v := range GenRangeCtx(ctx, 1, 5) {
//	    fmt.Println(v) // → 1, 2, 3, 4, 5  (atau lebih sedikit jika di-cancel)
//	}
//
// Hint:
//   - Goroutine: for i := dari; i <= sampai; i++ { select { case <-ctx.Done(): return; case ch <- i: } }
func GenRangeCtx(ctx context.Context, dari, sampai int) <-chan int {
	// TODO: implementasi di sini
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := dari; i <= sampai; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:

			}
		}
	}()

	return ch
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 2: TRANSFORM DENGAN CONTEXT
// ═══════════════════════════════════════════════════════════════════════════════
//
// Versi context-aware dari Filter dan TransformasiChan (Day 26).
// Jika ctx dibatalkan di tengah jalan, goroutine berhenti dan menutup output.
//

// FilterCtx meneruskan nilai dari src ke output HANYA jika fn(nilai) == true.
// Berhenti lebih awal jika ctx dibatalkan.
// Channel output ditutup saat src ditutup ATAU ctx dibatalkan.
//
// Contoh:
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//	src := GenRangeCtx(ctx, 1, 10)
//	genap := FilterCtx(ctx, src, func(n int) bool { return n%2 == 0 })
//	for v := range genap { fmt.Println(v) } // → 2, 4, 6, 8, 10
//
// Hint:
//   - Goroutine dengan for-select:
//     case <-ctx.Done(): return
//     case v, ok := <-src: if !ok { return }; if fn(v) { out <- v }
func FilterCtx(ctx context.Context, src <-chan int, fn func(int) bool) <-chan int {
	// TODO: implementasi di sini
	ch := make(chan int)

	go func() {
		defer close(ch)
		for v := range src {
			select {
			case <-ctx.Done():
				return
			default:
				if fn(v) {
					ch <- v
				}
			}
		}
	}()

	return ch
}

// TransformasiCtx menerapkan fn ke setiap nilai dari src.
// Berhenti lebih awal jika ctx dibatalkan.
//
// Contoh:
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//	src := GenRangeCtx(ctx, 1, 5)
//	kuadrat := TransformasiCtx(ctx, src, func(n int) int { return n * n })
//	for v := range kuadrat { fmt.Println(v) } // → 1, 4, 9, 16, 25
func TransformasiCtx(ctx context.Context, src <-chan int, fn func(int) int) <-chan int {
	// TODO: implementasi di sini
	ch := make(chan int)

	go func() {
		defer close(ch)
		for v := range src {
			select {
			case <-ctx.Done():
				return
			case ch <- fn(v):
			}
		}
	}()

	return ch
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 3: WORKER POOL DENGAN CONTEXT
// ═══════════════════════════════════════════════════════════════════════════════
//
// Worker pool = N goroutine yang sama-sama membaca dari satu input channel
// dan menulis hasil ke satu output channel (fan-out + fan-in sekaligus).
// Dengan context, seluruh pool bisa dibatalkan dari luar.
//
//   jobs ──┬── worker0 ──┐
//          ├── worker1 ──┼── hasil
//          └── worker2 ──┘
//

// WorkerPool menjalankan `jumlah` goroutine worker.
// Setiap worker membaca nilai dari `jobs`, menerapkan `fn`, dan mengirim ke output.
// Berhenti saat ctx dibatalkan ATAU jobs ditutup.
// Output channel ditutup setelah semua worker selesai.
//
// Contoh:
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//	jobs := GenRangeCtx(ctx, 1, 8)
//	hasil := WorkerPool(ctx, jobs, 3, func(n int) int { return n * n })
//	var out []int
//	for v := range hasil { out = append(out, v) }
//	sort.Ints(out) // order tidak dijamin (paralel)
//	// out = [1, 4, 9, 16, 25, 36, 49, 64]
//
// Hint:
//   - Buat output channel.
//   - Gunakan WaitGroup: Add(jumlah).
//   - Untuk tiap worker: goroutine dengan for-select { case <-ctx.Done(); case v,ok := <-jobs }.
//   - Goroutine terpisah: wg.Wait() → close(out).
func WorkerPool(ctx context.Context, jobs <-chan int, jumlah int, fn func(int) int) <-chan int {
	// TODO: implementasi di sini
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for v := range jobs {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- fn(v)
			}
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 4: RETRY DENGAN CONTEXT
// ═══════════════════════════════════════════════════════════════════════════════
//
// Operasi yang mungkin gagal (misalnya HTTP request) perlu di-retry.
// Dengan context, retry langsung berhenti jika ctx dibatalkan.
//

// CobaLagi menjalankan `fn` berulang kali sampai berhasil (return true)
// atau ctx dibatalkan atau sudah dicoba `maks` kali.
// Mengembalikan true jika berhasil, false jika tidak.
// Setiap percobaan gagal tunggu `jeda` sebelum mencoba lagi.
//
// Contoh:
//
//	percobaan := 0
//	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
//	defer cancel()
//	berhasil := CobaLagi(ctx, 5, 10*time.Millisecond, func() bool {
//	    percobaan++
//	    return percobaan >= 3 // berhasil di percobaan ke-3
//	})
//	// berhasil == true, percobaan == 3
//
// Hint:
//   - Loop sampai maks kali.
//   - Panggil fn(). Jika true, return true.
//   - Jika gagal: select { case <-ctx.Done(): return false; case <-time.After(jeda): }
func CobaLagi(ctx context.Context, maks int, jeda time.Duration, fn func() bool) bool {
	// TODO: implementasi di sini

	d := time.NewTicker(jeda)

	for i := 0; i < maks; i++ {
		select {
		case <-ctx.Done():
			return false
		case <-d.C:
			return false
		default:
			if fn() {
				return true
			}
		}
	}

	return false
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 5: JALANKAN PARALEL DENGAN BATAS WAKTU
// ═══════════════════════════════════════════════════════════════════════════════
//
// Menjalankan banyak fungsi secara paralel dengan satu deadline bersama.
// Jika salah satu selesai lebih awal, fungsi lain TETAP berjalan sampai selesai
// atau deadline tercapai.
//

// JalankanParalel menjalankan semua fungsi dalam `fns` secara bersamaan.
// Setiap fn menerima ctx yang sama. Hasil dikumpulkan ke []int.
// Menunggu semua selesai atau ctx dibatalkan (mana yang lebih dulu).
// Urutan hasil TIDAK dijamin (paralel).
//
// Contoh:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
//	defer cancel()
//	fns := []func(context.Context) int{
//	    func(ctx context.Context) int { return 1 },
//	    func(ctx context.Context) int { return 2 },
//	    func(ctx context.Context) int { return 3 },
//	}
//	hasil := JalankanParalel(ctx, fns)
//	sort.Ints(hasil)
//	// hasil = [1, 2, 3]
//
// Hint:
//   - Buat channel berkapasitas len(fns).
//   - Untuk tiap fn: goroutine → panggil fn(ctx) → kirim hasilnya ke channel.
//   - Setelah semua goroutine mulai, loop len(fns) kali untuk menerima hasil.
//   - Gunakan select: case v := <-ch: append; case <-ctx.Done(): break loop.
func JalankanParalel(ctx context.Context, fns []func(context.Context) int) []int {
	// TODO: implementasi di sini

	ch := make(chan int)
	var wg sync.WaitGroup

	for _, v := range fns {
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case <-ctx.Done():
			default:
				ch <- v(ctx)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	result := []int{}

	for v := range ch {
		result = append(result, v)
	}

	return result
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 6: PIPELINE DENGAN CONTEXT (COMPOSE)
// ═══════════════════════════════════════════════════════════════════════════════
//
// Gabungkan semua yang sudah dipelajari hari ini menjadi satu pipeline
// yang bisa dibatalkan dari luar lewat satu context.
//

// PipelineCtx membangun pipeline:
//
//	GenRangeCtx(ctx, dari, sampai)
//	→ FilterCtx(ctx, _, fn)
//	→ TransformasiCtx(ctx, _, transform)
//	→ kumpulkan ke []int
//
// Mengembalikan slice hasil. Jika ctx dibatalkan di tengah jalan,
// hasilnya adalah nilai yang sudah sempat dikumpulkan (bisa lebih sedikit).
//
// Contoh:
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//	// angka ganjil di [1,10], dikuadratkan: 1,9,25,49,81
//	hasil := PipelineCtx(ctx, 1, 10,
//	    func(n int) bool { return n%2 != 0 },
//	    func(n int) int  { return n * n },
//	)
//	// hasil = [1, 9, 25, 49, 81]
func PipelineCtx(ctx context.Context, dari, sampai int, filter func(int) bool, transform func(int) int) []int {
	// TODO: implementasi di sini
	ch := make(chan int, sampai-dari+1)
	result := []int{}

	for i := dari; i <= sampai; i++ {
		select {
		case <-ctx.Done():
			close(ch)
			for v := range ch {
				result = append(result, v)
			}
			return result
		default:
			if filter(i) {
				ch <- transform(i)
			}
		}
	}
	close(ch)

	for v := range ch {
		result = append(result, v)
	}

	return result
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 7: CEK ERR DARI CONTEXT
// ═══════════════════════════════════════════════════════════════════════════════
//
// ctx.Err() memberikan informasi KENAPA context selesai:
//   - context.Canceled       → di-cancel secara manual
//   - context.DeadlineExceeded → deadline/timeout tercapai
//   - nil                    → context masih aktif
//

// AlasanBerhenti mengembalikan string deskripsi alasan context berhenti.
// Kembalikan:
//   - "aktif"              jika ctx belum selesai (ctx.Err() == nil)
//   - "dibatalkan"         jika ctx.Err() == context.Canceled
//   - "waktu habis"        jika ctx.Err() == context.DeadlineExceeded
//   - "tidak diketahui"    untuk error lain
//
// Contoh:
//
//	ctx, cancel := context.WithCancel(context.Background())
//	fmt.Println(AlasanBerhenti(ctx)) // → "aktif"
//	cancel()
//	fmt.Println(AlasanBerhenti(ctx)) // → "dibatalkan"
func AlasanBerhenti(ctx context.Context) string {
	// TODO: implementasi di sini
	if ctx.Err() == nil {
		return "aktif"
	}

	switch ctx.Err() {
	case context.Canceled:
		return "dibatalkan"
	case context.DeadlineExceeded:
		return "waktu habis"
	default:
		return "tidak diketahui"
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 8: FIRST-WIN — PAKAI YANG PALING CEPAT SELESAI
// ═══════════════════════════════════════════════════════════════════════════════
//
// Terkadang kamu ingin menjalankan N komputasi paralel dan menggunakan
// HASIL PERTAMA yang selesai, lalu membatalkan sisanya.
//

// PertamaSelesai menjalankan semua fungsi dalam `fns` secara paralel.
// Mengembalikan hasil dari fn yang PERTAMA kali selesai.
// Setelah ada satu yang selesai, ctx yang dibuat internal di-cancel
// agar goroutine lain bisa berhenti (best-effort).
//
// Contoh:
//
//	fns := []func() int{
//	    func() int { time.Sleep(30*time.Millisecond); return 30 },
//	    func() int { time.Sleep(10*time.Millisecond); return 10 }, // ini menang
//	    func() int { time.Sleep(20*time.Millisecond); return 20 },
//	}
//	v := PertamaSelesai(fns)
//	// v == 10
//
// Hint:
//   - Buat buffered channel kapasitas len(fns).
//   - Jalankan semua fn dalam goroutine, kirim hasilnya ke channel.
//   - return <-ch  (ambil yang pertama masuk).
func PertamaSelesai(fns []func() int) int {
	// TODO: implementasi di sini
	ch := make(chan int)
	var wg sync.WaitGroup

	for _, v := range fns {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- v()
		}()
	}

	return <-ch
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🔧 HELPER — sudah diimplementasikan, jangan diubah
// ═══════════════════════════════════════════════════════════════════════════════

// KumpulkanSemua membaca semua nilai dari ch hingga channel ditutup,
// lalu mengembalikannya sebagai []int. Digunakan oleh beberapa test.
func KumpulkanSemua(ch <-chan int) []int {
	var hasil []int
	for v := range ch {
		hasil = append(hasil, v)
	}
	return hasil
}

// _ mencegah error "imported and not used"
var (
	_ *sync.WaitGroup
	_ = time.Second
	_ = context.Background
)
