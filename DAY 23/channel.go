package belajar

// ==================== DAY 23: CHANNEL & ASYNC — LATIHAN MANDIRI ====================
//
// 🎯 TUJUAN HARI INI:
//   Kamu akan mengimplementasikan sendiri fungsi-fungsi yang berhubungan dengan
//   channel dan goroutine. Setiap fungsi sudah dilengkapi penjelasan dan hint.
//   Jalankan test untuk mengecek hasil kerjamu:
//
//     cd "DAY 23"
//     go test ./... -v
//
// ═══════════════════════════════════════════════════════════════════════════════
// 📖 KILAT REVIEW: CHANNEL DI GO
// ═══════════════════════════════════════════════════════════════════════════════
//
//  MEMBUAT CHANNEL:
//    ch := make(chan int)      → unbuffered (blokir sampai ada penerima)
//    ch := make(chan int, 5)   → buffered kapasitas 5 (tidak blokir sampai penuh)
//
//  KIRIM & TERIMA:
//    ch <- 42       → kirim 42 ke channel
//    v := <-ch      → terima dari channel, simpan ke v
//    v, ok := <-ch  → ok=false jika channel sudah ditutup dan kosong
//
//  TUTUP CHANNEL:
//    close(ch)      → sinyal "tidak ada data lagi"
//                     ⚠️ hanya PENGIRIM yang boleh menutup!
//
//  RANGE CHANNEL:
//    for v := range ch { }  → baca terus sampai channel ditutup
//
//  SELECT:
//    select {
//    case v := <-ch1:  // ch1 punya data
//    case ch2 <- x:    // ch2 siap menerima
//    default:          // tidak ada yang siap → langsung lanjut (non-blocking)
//    }
//
//  GOROUTINE:
//    go fungsi()       → jalankan fungsi di goroutine terpisah (async)
//
// ═══════════════════════════════════════════════════════════════════════════════

import (
	"fmt"
	"sync"
	"time"
)

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 1: CHANNEL DASAR
// ═══════════════════════════════════════════════════════════════════════════════

// KirimSemua mengirimkan semua nilai dalam `data` ke channel `ch` secara berurutan.
// Setelah semua terkirim, channel DITUTUP.
//
// ⚠️ Pastikan ch punya kapasitas cukup (buffered) agar tidak deadlock.
//
// Contoh:
//
//	ch := make(chan int, 3)
//	KirimSemua(ch, 10, 20, 30)
//	// ch sekarang berisi 10, 20, 30 lalu ditutup
//	fmt.Println(<-ch) // → 10
//	fmt.Println(<-ch) // → 20
//	fmt.Println(<-ch) // → 30
//
// Hint:
//
//	for _, v := range data { ch <- v }
//	close(ch)
func KirimSemua(ch chan<- int, data ...int) {
	// TODO: implementasi di sini
	go func() {
		for v := range data { // ← baca dari channel input
			fmt.Println(data[v])
			ch <- data[v]
		}
		close(ch)
	}()
}

// TerimaSemuaSlice membaca SEMUA nilai dari channel `ch` sampai channel ditutup,
// dan mengembalikannya sebagai slice.
//
// Fungsi ini BLOKIR sampai channel ditutup.
//
// Contoh:
//
//	ch := make(chan int, 3)
//	ch <- 1
//	ch <- 2
//	ch <- 3
//	close(ch)
//	hasil := TerimaSemuaSlice(ch)
//	// hasil → [1, 2, 3]
//
// Hint: gunakan `for v := range ch`
func TerimaSemuaSlice(ch <-chan int) []int {
	// TODO: implementasi di sini
	result := []int{}
	done := make(chan struct{})

	go func() {
		for v := range ch {
			result = append(result, v)
		}
		close(done)
	}()

	<-done
	return result

}

// CekChannelTertutup membaca satu nilai dari channel.
// Mengembalikan (nilai, true) jika berhasil.
// Mengembalikan (0, false) jika channel sudah DITUTUP dan kosong.
//
// Contoh:
//
//	ch := make(chan int, 1)
//	ch <- 99
//	v, ok := CekChannelTertutup(ch)  // → (99, true)
//	close(ch)
//	v, ok = CekChannelTertutup(ch)   // → (0, false)
//
// Hint: gunakan `v, ok := <-ch`
func CekChannelTertutup(ch <-chan int) (int, bool) {
	// TODO: implementasi di sini
	v, ok := <-ch
	return v, ok
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 2: GOROUTINE + CHANNEL (ASYNC)
// ═══════════════════════════════════════════════════════════════════════════════

// HitungAsync menerima dua angka a dan b, menghitung a+b di goroutine terpisah,
// dan mengembalikan hasilnya lewat channel (read-only).
//
// Pemanggil tidak perlu menunggu — hasil datang saat goroutine selesai.
//
// Contoh:
//
//	ch := HitungAsync(3, 7)
//	result := <-ch  // → 10
//
// Hint:
//
//	out := make(chan int, 1)   ← buffered agar goroutine tidak blokir
//	go func() { out <- a + b }()
//	return out
func HitungAsync(a, b int) <-chan int {
	// TODO: implementasi di sini
	out := make(chan int, 1)
	go func() {
		out <- a + b
		close(out)
	}()
	return out
}

// KuadratAsync menerima slice angka, menghitung kuadrat tiap angka
// di goroutine terpisah, dan mengembalikan hasilnya lewat satu channel.
// Urutan hasil tidak harus sama dengan urutan input.
//
// Contoh:
//
//	ch := KuadratAsync(2, 3, 4)
//	// ch akan mengalirkan: 4, 9, 16 (urutan bisa beda)
//
// Hint:
//   - out := make(chan int, len(nums))
//   - for _, n := range nums { go func(n int) { out <- n*n }(n) }
//   - return out (JANGAN close di sini — biarkan goroutine yang isi)
func KuadratAsync(nums ...int) <-chan int {
	// TODO: implementasi di sini
	out := make(chan int, len(nums))
	for _, n := range nums {
		go func(n int) {
			out <- n * n
		}(n)
	}
	return out
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 3: SELECT
// ═══════════════════════════════════════════════════════════════════════════════

// AmbilTercepat menjalankan dua "pekerjaan" secara async dan mengembalikan
// hasil dari yang PERTAMA selesai. Yang satunya diabaikan.
//
// Contoh:
//
//	hasil := AmbilTercepat(
//	    func() int { time.Sleep(300*time.Millisecond); return 1 },
//	    func() int { time.Sleep(100*time.Millisecond); return 2 },
//	)
//	// hasil → 2  (pekerjaan kedua lebih cepat)
//
// Hint:
//
//	ch1 := make(chan int, 1)
//	ch2 := make(chan int, 1)
//	go func() { ch1 <- kerjaan1() }()
//	go func() { ch2 <- kerjaan2() }()
//	select {
//	case v := <-ch1: return v
//	case v := <-ch2: return v
//	}
func AmbilTercepat(kerjaan1, kerjaan2 func() int) int {
	// TODO: implementasi di sini
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	go func() { ch1 <- kerjaan1() }()
	go func() { ch2 <- kerjaan2() }()
	select {
	case v := <-ch1:
		return v
	case v := <-ch2:
		return v
	}
}

// CobaAmbil mencoba membaca dari channel TANPA MENUNGGU (non-blocking).
// Jika ada data → return (data, true).
// Jika kosong   → return (0, false) SEGERA.
//
// Contoh:
//
//	ch := make(chan int, 1)
//	v, ok := CobaAmbil(ch)   // kosong → (0, false)
//	ch <- 55
//	v, ok = CobaAmbil(ch)    // ada isi → (55, true)
//
// Hint: gunakan `select { case v := <-ch: ... default: ... }`
func CobaAmbil(ch <-chan int) (int, bool) {
	// TODO: implementasi di sini
	select {
	case v := <-ch:
		return v, true
	default:
		return 0, false
	}
}

// TungguDenganBatas menjalankan `kerja` secara async.
// Jika selesai sebelum `batas` waktu → return (hasil, true).
// Jika waktu habis duluan           → return (0, false).
//
// Contoh:
//
//	// cepat → berhasil
//	hasil, ok := TungguDenganBatas(func() int { return 42 }, 1*time.Second)
//	// hasil=42, ok=true
//
//	// lambat → timeout
//	hasil, ok = TungguDenganBatas(func() int {
//	    time.Sleep(5*time.Second); return 99
//	}, 100*time.Millisecond)
//	// hasil=0, ok=false
//
// Hint:
//
//	ch := make(chan int, 1)
//	go func() { ch <- kerja() }()
//	select {
//	case v := <-ch:           return v, true
//	case <-time.After(batas): return 0, false
//	}
func TungguDenganBatas(kerja func() int, batas time.Duration) (int, bool) {
	// TODO: implementasi di sini
	ch := make(chan int, 1)
	go func() { ch <- kerja() }()
	select {
	case v := <-ch:
		return v, true
	case <-time.After(batas):
		return 0, false
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 4: GOROUTINE + WAITGROUP
// ═══════════════════════════════════════════════════════════════════════════════

// JalankanParalel menjalankan semua fungsi dalam `tugas` secara paralel
// (setiap fungsi di goroutine terpisah) dan MENUNGGU sampai semua selesai.
//
// Contoh:
//
//	var counter int64
//	JalankanParalel(
//	    func() { atomic.AddInt64(&counter, 1) },
//	    func() { atomic.AddInt64(&counter, 1) },
//	    func() { atomic.AddInt64(&counter, 1) },
//	)
//	// counter == 3
//
// Hint:
//
//	var wg sync.WaitGroup
//	for _, t := range tugas {
//	    wg.Add(1)
//	    go func(fn func()) { defer wg.Done(); fn() }(t)
//	}
//	wg.Wait()
func JalankanParalel(tugas ...func()) {
	// TODO: implementasi di sini
	var wg sync.WaitGroup
	for _, t := range tugas {
		wg.Add(1)
		go func(fn func()) { defer wg.Done(); fn() }(t)
	}
	wg.Wait()
}

// KumpulkanHasil menjalankan semua `penghasil` secara paralel dan
// mengumpulkan semua hasilnya ke dalam slice.
// Urutan hasil tidak harus sama dengan urutan penghasil.
//
// Contoh:
//
//	hasil := KumpulkanHasil(
//	    func() int { return 10 },
//	    func() int { return 20 },
//	    func() int { return 30 },
//	)
//	sort.Ints(hasil) // → [10, 20, 30]
//
// Hint:
//   - ch := make(chan int, len(penghasil))
//   - goroutine kirim hasil ke ch
//   - kumpulkan len(penghasil) nilai
func KumpulkanHasil(penghasil ...func() int) []int {
	// TODO: implementasi di sini
	ch := make(chan int, len(penghasil))
	result := []int{}
	for k := range penghasil {
		result = append(result, penghasil[k]())
	}
	close(ch)
	return result
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 5: DONE CHANNEL — Hentikan goroutine
// ═══════════════════════════════════════════════════════════════════════════════

// GeneratorAngka menghasilkan angka 1, 2, 3, ... secara terus-menerus
// ke channel output, sampai channel `selesai` ditutup.
// Setelah `selesai` ditutup, channel output juga DITUTUP.
//
// Contoh:
//
//	selesai := make(chan struct{})
//	angkaCh := GeneratorAngka(selesai)
//
//	fmt.Println(<-angkaCh)  // 1
//	fmt.Println(<-angkaCh)  // 2
//	fmt.Println(<-angkaCh)  // 3
//	close(selesai)          // hentikan generator
//
// Hint:
//
//	out := make(chan int)
//	go func() {
//	    n := 1
//	    for {
//	        select {
//	        case <-selesai:
//	            close(out)
//	            return
//	        case out <- n:
//	            n++
//	        }
//	    }
//	}()
//	return out
func GeneratorAngka(selesai <-chan struct{}) <-chan int {
	// TODO: implementasi di sini
	out := make(chan int)
	go func() {
		n := 1
		for {
			select {
			case <-selesai:
				close(out)
				return
			case out <- n:
				n++
			}
		}
	}()
	return out
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 6: FAN-OUT & FAN-IN
// ═══════════════════════════════════════════════════════════════════════════════

// SebarKeSemua (Fan-Out) membaca tiap nilai dari channel `masuk` dan
// mengirimkan nilai yang SAMA ke semua channel dalam `keluar`.
// Berhenti saat channel `masuk` ditutup, lalu menutup semua channel `keluar`.
//
// Analogi: satu mikrofon → siaran ke banyak speaker.
//
// Contoh:
//
//	masuk := make(chan int, 3)
//	out1  := make(chan int, 3)
//	out2  := make(chan int, 3)
//	masuk <- 1; masuk <- 2; masuk <- 3
//	close(masuk)
//
//	SebarKeSemua(masuk, out1, out2)
//
//	// out1 berisi: 1, 2, 3
//	// out2 berisi: 1, 2, 3
//
// Hint:
//
//	for v := range masuk {
//	    for _, out := range keluar { out <- v }
//	}
//	for _, out := range keluar { close(out) }
func SebarKeSemua(masuk <-chan int, keluar ...chan int) {
	// TODO: implementasi di sini
	for v := range masuk {
		for _, out := range keluar {
			out <- v
		}
	}
	for _, out := range keluar {
		close(out)
	}
}

// GabungkanChannel (Fan-In) menggabungkan BANYAK channel menjadi SATU.
// Nilai dari semua channel input akan dialirkan ke channel output.
// Output ditutup setelah SEMUA input ditutup.
//
// Analogi: banyak speaker → satu rekaman.
//
// Contoh:
//
//	ch1 := Generate(1, 3, 5)
//	ch2 := Generate(2, 4, 6)
//	gabung := GabungkanChannel(ch1, ch2)
//	var hasil []int
//	for v := range gabung { hasil = append(hasil, v) }
//	sort.Ints(hasil) // → [1, 2, 3, 4, 5, 6]
//
// Hint:
//   - var wg sync.WaitGroup
//   - untuk tiap channel: go func() { defer wg.Done(); for v := range ch { out <- v } }()
//   - go func() { wg.Wait(); close(out) }()
func GabungkanChannel(channels ...<-chan int) <-chan int {
	// TODO: implementasi di sini
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(len(channels))
	go func() {
		for k, v := range channels {
			go func() {
				defer wg.Done()
				for value := range v {
					ch <- value
					fmt.Println(value, k, v)
				}
			}()
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 7: WORKER POOL
// ═══════════════════════════════════════════════════════════════════════════════

// WorkerPool menjalankan `proses` terhadap setiap angka dari `pekerjaan`
// menggunakan sejumlah `jumlahWorker` goroutine secara paralel.
// Mengembalikan channel yang mengalirkan semua hasil.
// Channel hasil ditutup setelah semua pekerjaan selesai diproses.
//
// Analogi: kasir supermarket — N kasir, banyak pelanggan antri.
//
// Contoh:
//
//	jobs := make(chan int, 5)
//	jobs <- 2; jobs <- 3; jobs <- 4
//	close(jobs)
//
//	hasilCh := WorkerPool(jobs, 2, func(n int) int { return n * n })
//	for v := range hasilCh {
//	    fmt.Println(v)  // 4, 9, 16 (urutan bisa beda)
//	}
//
// Hint:
//
//	out := make(chan int, cap(pekerjaan))
//	var wg sync.WaitGroup
//	for i := 0; i < jumlahWorker; i++ {
//	    wg.Add(1)
//	    go func() {
//	        defer wg.Done()
//	        for job := range pekerjaan {
//	            out <- proses(job)
//	        }
//	    }()
//	}
//	go func() { wg.Wait(); close(out) }()
//	return out
func WorkerPool(pekerjaan <-chan int, jumlahWorker int, proses func(int) int) <-chan int {
	// TODO: implementasi di sini
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(jumlahWorker)

	go func() {
		for i := 0; i < jumlahWorker; i++ {
			go func() {
				defer wg.Done()
				for job := range pekerjaan {
					ch <- proses(job)
					fmt.Println(proses(job), job, i)
				}
			}()
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 8: PIPELINE
// ═══════════════════════════════════════════════════════════════════════════════

// Sumber membuat channel yang mengalirkan angka-angka dari `data`,
// lalu menutup channel setelah semua terkirim.
// Ini adalah TAHAP PERTAMA pipeline (sumber data).
//
// Contoh:
//
//	ch := Sumber(1, 2, 3)
//	fmt.Println(<-ch) // 1
//	fmt.Println(<-ch) // 2
//	fmt.Println(<-ch) // 3
//	// channel ditutup otomatis
func Sumber(data ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, v := range data {
			ch <- v
		}
	}()
	return ch
}

// Kalikan membaca tiap nilai dari `masuk`, mengalikannya dengan `pengali`,
// dan mengalirkan hasilnya ke channel baru.
// Channel output ditutup saat channel input ditutup.
// Ini adalah TAHAP TENGAH pipeline (transformasi).
//
// Contoh:
//
//	src := Sumber(1, 2, 3)
//	hasil := Kalikan(src, 3)
//	for v := range hasil { fmt.Println(v) } // → 3, 6, 9
func Kalikan(masuk <-chan int, pengali int) <-chan int {
	// TODO: implementasi di sini
	ch := make(chan int)
	go func() {
		defer close(ch)
		for v := range masuk {
			ch <- v * pengali
		}
	}()
	return ch
}

// Tambahkan membaca tiap nilai dari `masuk`, menambahkannya dengan `n`,
// dan mengalirkan hasilnya ke channel baru.
// Channel output ditutup saat channel input ditutup.
//
// Contoh pipeline lengkap:
//
//	src    := Sumber(1, 2, 3)       // → 1, 2, 3
//	kali   := Kalikan(src, 2)       // → 2, 4, 6
//	tambah := Tambahkan(kali, 10)   // → 12, 14, 16
//	for v := range tambah { fmt.Println(v) }

func Tambahkan(masuk <-chan int, n int) <-chan int {
	// TODO: implementasi di sini
	ch := make(chan int)
	go func() {
		defer close(ch)
		for v := range masuk {
			ch <- v + n
		}
	}()
	return ch
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 9: SEMAPHORE — Batasi goroutine dengan buffered channel
// ═══════════════════════════════════════════════════════════════════════════════

// ParalelTerbatas menjalankan semua `tugas`, tapi maksimum `maks` goroutine
// yang berjalan BERSAMAAN. Menunggu semua tugas selesai sebelum return.
//
// Analogi: antrian loket — 1000 orang, hanya 3 loket buka.
//
// Contoh:
//
//	// Verifikasi: tidak pernah lebih dari 3 goroutine aktif bersamaan
//	var aktif int64
//	tasks := make([]func(), 20)
//	for i := range tasks {
//	    tasks[i] = func() {
//	        atomic.AddInt64(&aktif, 1)
//	        time.Sleep(10*time.Millisecond)
//	        atomic.AddInt64(&aktif, -1)
//	    }
//	}
//	ParalelTerbatas(3, tasks...)
//
// Hint (semaphore dengan buffered channel):
//
//	sem := make(chan struct{}, maks)
//	var wg sync.WaitGroup
//	for _, t := range tugas {
//	    wg.Add(1)
//	    sem <- struct{}{}              // acquire: tunggu slot
//	    go func(fn func()) {
//	        defer wg.Done()
//	        defer func() { <-sem }()  // release: kembalikan slot
//	        fn()
//	    }(t)
//	}
//	wg.Wait()
func ParalelTerbatas(maks int, tugas ...func()) {
	// TODO: implementasi di sini
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🔧 HELPER (sudah diimplementasikan, boleh dipakai di kode kamu)
// ═══════════════════════════════════════════════════════════════════════════════

// buatChannel adalah helper untuk membuat channel yang sudah berisi data.
// Dipakai di beberapa test. Sudah diimplementasikan — tidak perlu diubah.
func buatChannel(data ...int) <-chan int {
	ch := make(chan int, len(data))
	for _, v := range data {
		ch <- v
	}
	close(ch)
	return ch
}

// _ dipakai agar import sync dan time tidak error sebelum kamu implementasikan.
var _ = sync.WaitGroup{}
var _ = time.Second
