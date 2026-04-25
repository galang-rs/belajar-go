package belajar

// ==================== DAY 24: ASYNC LANJUTAN — LATIHAN MANDIRI ====================
//
// 🎯 TUJUAN HARI INI:
//   Mendalami pola async yang lebih kompleks dan realistis:
//   error handling antar goroutine, context cancellation, pipeline async,
//   dan pola-pola yang sering muncul di kode produksi Go.
//
//     cd "DAY 24"
//     go test -run "Async" ./... -v
//
// ═══════════════════════════════════════════════════════════════════════════════
// 📖 REVIEW KILAT: APA YANG SUDAH KITA TAHU
// ═══════════════════════════════════════════════════════════════════════════════
//
//  DAY 22: goroutine dasar, WaitGroup, channel Future, timeout, semaphore
//  DAY 23: fan-out, fan-in, pipeline, worker pool, done channel
//
// ═══════════════════════════════════════════════════════════════════════════════
// 📖 KONSEP BARU HARI INI
// ═══════════════════════════════════════════════════════════════════════════════
//
//  1. ERRGROUP PATTERN   — Jalankan goroutine, tangkap error pertama
//  2. ASYNC PIPELINE     — Pipeline multi-tahap yang berjalan concurrent
//  3. CONTEXT TIMEOUT    — Beri deadline pada goroutine
//  4. FIRST-SUCCESS      — Jalankan N goroutine, ambil hasil pertama yang SUKSES
//  5. PARALLEL MAP       — Map setiap elemen secara paralel, kumpulkan hasilnya
//  6. SCATTER-GATHER     — Sebar request ke banyak sumber, kumpulkan semua hasil
//
// ═══════════════════════════════════════════════════════════════════════════════

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 1: ERRGROUP PATTERN
// ═══════════════════════════════════════════════════════════════════════════════
//
// Masalah: WaitGroup hanya menunggu goroutine selesai, tapi tidak bisa
//          mengumpulkan error dari goroutine.
//
// Solusi: channel error + WaitGroup.
//   - Tiap goroutine kirim error-nya ke channel.
//   - Setelah semua selesai, kumpulkan error yang ada.
//
// ⚠️ Penting: channel error harus BUFFERED (kapasitas = jumlah goroutine)
//    agar goroutine tidak blokir saat kirim error ke channel.

// JalankanKumpulkanError menjalankan semua `tugas` secara paralel.
// Tiap tugas bisa return nil (berhasil) atau error (gagal).
// Mengembalikan slice semua error non-nil yang terjadi.
// Jika semua berhasil, mengembalikan slice kosong.
//
// Contoh:
//
//	errs := JalankanKumpulkanError(
//	    func() error { return nil },
//	    func() error { return errors.New("gagal A") },
//	    func() error { return errors.New("gagal B") },
//	)
//	// len(errs) == 2
//
// Hint:
//   - Mirip dengan RunAndCollectErrors di DAY 22 (lihat async.go DAY 22 untuk inspirasi).
//   - Gunakan buffered channel untuk menampung error dari semua goroutine.
//   - Gunakan WaitGroup untuk menunggu semua goroutine selesai sebelum drain channel.
func JalankanKumpulkanError(tugas ...func() error) []error {
	// TODO: implementasi di sini
	var err []error
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, v := range tugas {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result := v()
			if result != nil {
				mu.Lock()
				err = append(err, result)
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	return err
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 2: FIRST-SUCCESS PATTERN
// ═══════════════════════════════════════════════════════════════════════════════
//
// Berbeda dengan AmbilTercepat (DAY 23) yang ambil hasil pertama apapun —
// FirstSuccess hanya ambil hasil pertama yang BERHASIL (error == nil).
//
// Contoh nyata: query ke 3 database replica. Pakai hasil dari yang pertama
// menjawab dengan benar. Sisanya diabaikan.
//
// Skema:
//   goroutine A → gagal (error)
//   goroutine B → gagal (error)   ← diabaikan
//   goroutine C → BERHASIL (99) ← ini yang dipakai!

// AmbilPertamaBerhasil menjalankan semua `kandidat` secara paralel.
// Mengembalikan (hasil, nil) dari kandidat PERTAMA yang berhasil (error == nil).
// Jika SEMUA kandidat gagal, mengembalikan (0, error terakhir).
//
// Contoh:
//
//	hasil, err := AmbilPertamaBerhasil(
//	    func() (int, error) { return 0, errors.New("timeout") },
//	    func() (int, error) { return 0, errors.New("not found") },
//	    func() (int, error) { return 42, nil },
//	)
//	// hasil=42, err=nil
//
// Hint:
//   - Gunakan channel struct{ val int; err error } untuk menampung hasil tiap goroutine.
//   - Kumpulkan semua hasil (len(kandidat) kali), cari yang pertama err == nil.
//   - Alternatif: gunakan select + context untuk batalkan goroutine lain saat ada yang sukses.
func AmbilPertamaBerhasil(kandidat ...func() (int, error)) (int, error) {
	// TODO: implementasi di sini
	var wg sync.WaitGroup
	var mu sync.Mutex
	result, err := 0, fmt.Errorf("err")
	for _, v := range kandidat {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer mu.Unlock()
			mu.Lock()
			result, err = v()
			if err == nil {
				return
			}
		}()
	}
	wg.Wait()
	return result, err
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 3: CONTEXT TIMEOUT
// ═══════════════════════════════════════════════════════════════════════════════
//
// Context membawa "sinyal pembatalan" yang bisa diteruskan ke goroutine.
// context.WithTimeout → otomatis dibatalkan setelah durasi tertentu.
//
//   ctx, cancel := context.WithTimeout(parent, 1*time.Second)
//   defer cancel()
//
//   select {
//   case <-ctx.Done():  // waktu habis atau dibatalkan
//   case v := <-ch:    // hasil datang tepat waktu
//   }

// JalankanDenganContext menjalankan `kerja` di goroutine terpisah
// dengan context yang sudah dibuat pemanggil.
// Jika ctx dibatalkan atau timeout sebelum kerja selesai → return (0, ctx.Err()).
// Jika kerja selesai tepat waktu → return (hasil, nil).
//
// Contoh:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
//	defer cancel()
//
//	// Kerja cepat → berhasil
//	hasil, err := JalankanDenganContext(ctx, func() int {
//	    return 42
//	})
//	// hasil=42, err=nil
//
//	// Kerja lambat → timeout
//	hasil, err = JalankanDenganContext(ctx, func() int {
//	    time.Sleep(1 * time.Second)
//	    return 99
//	})
//	// hasil=0, err=context.DeadlineExceeded
//
// Hint:
//   - Buat buffered channel, jalankan kerja() di goroutine.
//   - Gunakan select: tunggu hasil dari channel ATAU ctx.Done().
//   - Jika ctx.Done() yang menang → return 0, ctx.Err().
func JalankanDenganContext(ctx context.Context, kerja func() int) (int, error) {
	// TODO: implementasi di sini
	resultCh := make(chan int, 1)

	go func() {
		resultCh <- kerja()
	}()

	select {
	case res := <-resultCh:
		return res, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 4: PARALLEL MAP
// ═══════════════════════════════════════════════════════════════════════════════
//
// Pola umum: punya daftar item, ingin transformasi tiap item secara paralel.
// Mirip MapOrdered di channel.go, tapi input/output berupa slice langsung
// (bukan streaming via channel).
//
// Kegunaan nyata:
//   - Resize banyak gambar secara paralel
//   - Fetch URL dari daftar secara bersamaan
//   - Enkripsi banyak file sekaligus

// ParallelMap menerapkan `fn` ke setiap elemen `input` secara PARALEL
// dan mengembalikan slice hasil dengan URUTAN YANG SAMA seperti input.
//
// Contoh:
//
//	hasil := ParallelMap([]int{1, 2, 3, 4, 5}, func(n int) int {
//	    time.Sleep(10 * time.Millisecond) // simulasi kerja berat
//	    return n * n
//	})
//	// hasil → [1, 4, 9, 16, 25]  ← urutan terjaga, tapi dikerjakan paralel
//
// Hint:
//   - Buat slice hasil berukuran sama dengan input.
//   - Gunakan WaitGroup dan goroutine untuk setiap elemen.
//   - Tiap goroutine boleh langsung tulis ke hasil[i] — aman karena tiap
//     goroutine menulis ke indeks yang berbeda.
//   - Tunggu semua goroutine selesai sebelum return.
func ParallelMap(input []int, fn func(int) int) []int {
	// TODO: implementasi di sini
	var wg sync.WaitGroup
	result := make([]int, len(input))

	for k, v := range input {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result[k] = fn(v)
		}()
	}
	wg.Wait()

	return result
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 5: SCATTER-GATHER
// ═══════════════════════════════════════════════════════════════════════════════
//
// Scatter-Gather = sebar ke banyak sumber, kumpulkan semua jawaban.
//
// Berbeda dengan AmbilTercepat (ambil satu saja) —
// Scatter-Gather mengambil SEMUA hasil dari semua sumber.
//
// Contoh nyata: tanya ke 5 server harga barang, kumpulkan semua, pilih termurah.
//
//   Scatter: kirim request ke server A, B, C bersamaan
//   Gather:  tunggu semua selesai, kumpulkan hasilnya
//   Proses:  pilih yang terbaik dari semua hasil

// ScatterGather menjalankan semua `sumber` secara paralel dan
// mengumpulkan SEMUA hasilnya ke dalam slice.
// Menunggu sampai SEMUA sumber selesai (berbeda dengan AmbilTercepat).
// Urutan hasil tidak harus sama dengan urutan sumber.
//
// Contoh:
//
//	hasil := ScatterGather(
//	    func() int { time.Sleep(30*time.Millisecond); return 100 },
//	    func() int { time.Sleep(10*time.Millisecond); return 50  },
//	    func() int { time.Sleep(20*time.Millisecond); return 75  },
//	)
//	sort.Ints(hasil) // → [50, 75, 100]
//
// Hint:
//   - Mirip GatherResults dari DAY 22 — coba ingat kembali polanya.
//   - Gunakan buffered channel dengan kapasitas = len(sumber).
//   - Jalankan semua sumber di goroutine, kumpulkan len(sumber) hasil.
func ScatterGather(sumber ...func() int) []int {
	// TODO: implementasi di sini
	var wg sync.WaitGroup
	wg.Add(len(sumber))
	result := make([]int, len(sumber))

	for k, v := range sumber {
		go func() {
			defer wg.Done()
			result[k] = v()
		}()
	}

	wg.Wait()

	return result
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 6: ASYNC PIPELINE
// ═══════════════════════════════════════════════════════════════════════════════
//
// Menggabungkan semua konsep pipeline + async + context menjadi satu.
// Pipeline ini bisa dibatalkan di tengah jalan via context.
//
// Alur:
//   input []int → [SumberAsync] → channel → [Proses] → channel → hasil
//
// Setiap tahap berjalan di goroutine sendiri, dihubungkan dengan channel.
// Jika ctx dibatalkan, semua tahap berhenti dengan bersih.

// SumberAsync menghasilkan nilai dari `data` ke channel, dengan menghormati
// konteks. Jika ctx dibatalkan di tengah pengiriman, goroutine berhenti
// dan channel output ditutup.
//
// Contoh:
//
//	ctx, cancel := context.WithCancel(context.Background())
//	ch := SumberAsync(ctx, 1, 2, 3, 4, 5)
//	cancel() // batalkan di tengah jalan
//	// ch tidak lagi menghasilkan nilai baru setelah cancel
//
// Hint:
//   - Mirip Sumber() di channel.go, tapi setiap pengiriman harus lewat select.
//   - select: coba kirim nilai ke out, ATAU deteksi ctx.Done() dan return.
func SumberAsync(ctx context.Context, data ...int) <-chan int {
	// TODO: implementasi di sini
	ch := make(chan int, len(data))
	var wg sync.WaitGroup

	for _, v := range data {
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			case ch <- v:
			}
		}()
	}
	wg.Wait()
	close(ch)
	return ch
}

// ProsesAsync membaca nilai dari `masuk`, menerapkan `fn` ke setiap nilai,
// dan mengalirkan hasilnya ke channel output. Berhenti jika ctx dibatalkan
// atau channel input ditutup.
//
// Contoh:
//
//	ctx := context.Background()
//	src := SumberAsync(ctx, 1, 2, 3)
//	doubled := ProsesAsync(ctx, src, func(n int) int { return n * 2 })
//	for v := range doubled { fmt.Println(v) } // → 2, 4, 6
//
// Hint:
//   - Goroutine dengan loop: select antara `masuk` dan `ctx.Done()`.
//   - Jika masuk ditutup (ok == false), tutup output dan return.
//   - Jika ctx.Done() → tutup output dan return.
func ProsesAsync(ctx context.Context, masuk <-chan int, fn func(int) int) <-chan int {
	// TODO: implementasi di sini
	out := make(chan int)

	go func() {
		defer close(out)

		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-masuk:
				if !ok {
					return
				}

				out <- fn(v) // urut 100%
			}
		}
	}()

	return out

}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 7: WORKER POOL DENGAN ERROR HANDLING
// ═══════════════════════════════════════════════════════════════════════════════
//
// Melanjutkan WorkerPool dari DAY 23 — kali ini worker bisa return error.
// Jika ada job yang gagal, error-nya dikumpulkan, tapi job lain tetap jalan.
//
// Kegunaan nyata: batch processing file. File yang gagal diproses dicatat,
// sisanya tetap diproses sampai selesai.

// HasilKerja menyimpan hasil dari satu job: nilai hasil dan error-nya.
type HasilKerja struct {
	Nilai int
	Err   error
}

// WorkerPoolDenganError menjalankan `proses` terhadap setiap angka dari `pekerjaan`
// menggunakan `jumlahWorker` goroutine secara paralel.
// Proses bisa berhasil (return nilai, nil) atau gagal (return 0, error).
// Mengembalikan channel HasilKerja yang mengalirkan SEMUA hasil (sukses maupun gagal).
// Channel hasil ditutup setelah semua pekerjaan selesai diproses.
//
// Contoh:
//
//	jobs := make(chan int, 3)
//	jobs <- 1; jobs <- 2; jobs <- 3
//	close(jobs)
//
//	hasilCh := WorkerPoolDenganError(jobs, 2, func(n int) (int, error) {
//	    if n == 2 { return 0, errors.New("angka 2 ditolak") }
//	    return n * n, nil
//	})
//
//	for h := range hasilCh {
//	    if h.Err != nil { fmt.Println("gagal:", h.Err) }
//	    if h.Err == nil { fmt.Println("berhasil:", h.Nilai) }
//	}
//
// Hint:
//   - Mirip WorkerPool dari DAY 23, tapi channel output bertipe chan HasilKerja.
//   - Tiap worker kirim HasilKerja{Nilai: v, Err: nil} atau HasilKerja{Err: err}.
//   - Gunakan WaitGroup + goroutine penutup untuk menutup channel setelah semua selesai.
func WorkerPoolDenganError(pekerjaan <-chan int, jumlahWorker int, proses func(int) (int, error)) <-chan HasilKerja {
	// TODO: implementasi di sini
	return nil
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 8: THROTTLED ASYNC
// ═══════════════════════════════════════════════════════════════════════════════
//
// Menggabungkan Rate Limiter (channel.go Bagian 4) dengan eksekusi async.
// Berguna saat harus panggil API eksternal yang punya rate limit.
//
// Skema:
//   [jobs] → tunggu token dari rate limiter → jalankan job → [hasil]

// JalankanTerbatas menjalankan setiap elemen `tugas` menggunakan `fn`,
// dengan pembatasan: maksimum `ratePerDetik` tugas dijalankan per detik.
// Menunggu semua tugas selesai, lalu mengembalikan semua hasilnya.
// Urutan hasil tidak harus sama dengan urutan tugas.
//
// Contoh:
//
//	// 6 tugas, max 2 per detik → ~3 detik total
//	hasil := JalankanTerbatas([]int{1, 2, 3, 4, 5, 6}, func(n int) int {
//	    return n * 10
//	}, 2)
//	sort.Ints(hasil) // → [10, 20, 30, 40, 50, 60]
//
// Hint:
//   - Gunakan time.NewTicker dengan interval 1 detik / ratePerDetik.
//   - Untuk setiap tugas: tunggu satu tick dulu, lalu jalankan fn(tugas[i]).
//   - Kumpulkan semua hasil (urutan boleh acak).
//   - Ingat: kamu boleh jalankan fn() di goroutine atau sequential setelah tick.
func JalankanTerbatas(tugas []int, fn func(int) int, ratePerDetik int) []int {
	// TODO: implementasi di sini
	return nil
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🔧 HELPER (sudah diimplementasikan, jangan diubah)
// ═══════════════════════════════════════════════════════════════════════════════

// _ mencegah error "imported and not used" sebelum implementasi selesai.
var (
	_ = context.Background
	_ = time.Second
	_ sync.WaitGroup
)
