package belajar

// ==================== DAY 22: ASYNC GOLANG — DARI NOL ====================
//
// 🎯 TUJUAN HARI INI:
//   Memahami apa itu async, kenapa perlu, dan bagaimana Go melakukannya
//   dengan goroutine + channel + select + context.
//
// ═══════════════════════════════════════════════════════════════════════════
// 📖 KONSEP 1: APA ITU SYNC vs ASYNC?
// ═══════════════════════════════════════════════════════════════════════════
//
// SYNC (Sequential / Berurutan):
//   Bayangkan kamu memasak mie, dan kamu duduk diam menunggu air mendidih.
//   Tidak ada yang kamu kerjakan selama menunggu. Baru setelah mendidih,
//   kamu masukkan mie, tunggu lagi, dst.
//
//   Kode:
//     result1 := tugasBerat1()  // ← tunggu selesai
//     result2 := tugasBerat2()  // ← baru mulai ini
//
//   Total waktu = waktu1 + waktu2 + ...
//
// ASYNC (Concurrent / Bersamaan):
//   Kamu nyalakan kompor, lalu SAMBIL menunggu air mendidih, kamu iris bawang.
//   Kamu tidak buang-buang waktu menunggu.
//
//   Kode:
//     go tugasBerat1()  // ← mulai, tidak ditunggu
//     go tugasBerat2()  // ← langsung mulai juga
//     // keduanya jalan BERSAMAAN
//
//   Total waktu = max(waktu1, waktu2) -- jauh lebih singkat!
//
// ═══════════════════════════════════════════════════════════════════════════
// 📖 KONSEP 2: GOROUTINE — "Lightweight Thread" milik Go
// ═══════════════════════════════════════════════════════════════════════════
//
// Thread OS = berat (MB per thread, lambat dibuat)
// Goroutine  = ringan (~2KB awal, sangat cepat dibuat, bisa jutaan sekaligus)
//
// Cara membuat goroutine:
//   go namaFungsi()
//   go func() { ... }()
//
// Hanya tambah kata `go` di depan pemanggilan fungsi. Selesai!
//
// ⚠️  MASALAH UTAMA:
//   Kalau main() selesai, SEMUA goroutine ikut mati — bahkan yang belum selesai!
//   Solusi: kamu HARUS menunggu goroutine selesai sebelum main() keluar.
//   Caranya: pakai sync.WaitGroup ATAU channel.
//
// ═══════════════════════════════════════════════════════════════════════════
// 📖 KONSEP 3: CHANNEL — "Pipa" komunikasi antar goroutine
// ═══════════════════════════════════════════════════════════════════════════
//
// Channel = pipa data yang thread-safe untuk komunikasi goroutine.
// Prinsip Go: "Don't communicate by sharing memory; share memory by communicating."
// Artinya: jangan pakai variabel global bersama, gunakan channel!
//
// Membuat channel:
//   ch := make(chan int)       // unbuffered: kapasitas 0
//   ch := make(chan int, 5)    // buffered:  kapasitas 5
//
// Kirim data:
//   ch <- 42      // BLOKIR sampai ada yang siap menerima (unbuffered)
//
// Terima data:
//   nilai := <-ch // BLOKIR sampai ada yang kirim
//
// Tutup channel:
//   close(ch)     // sinyal "tidak ada data lagi yang dikirim"
//                 // ⚠️ hanya pengirim yang boleh menutup channel!
//
// Range channel (baca sampai channel ditutup):
//   for v := range ch { ... }
//
// ═══════════════════════════════════════════════════════════════════════════
// 📖 KONSEP 4: UNBUFFERED vs BUFFERED CHANNEL
// ═══════════════════════════════════════════════════════════════════════════
//
// UNBUFFERED — make(chan T)
//   Pengirim dan penerima harus siap BERSAMAAN (rendezvous/sinkronisasi).
//   Ibarat telepon: dua orang harus angkat di waktu yang sama.
//   → Cocok untuk sinkronisasi ketat antar goroutine.
//
// BUFFERED — make(chan T, N)
//   Pengirim bisa kirim hingga N item tanpa blokir (selama buffer tidak penuh).
//   Ibarat SMS: kirim beberapa pesan, penerima baca nanti.
//   → Cocok untuk menampung hasil dari banyak goroutine.
//
// ═══════════════════════════════════════════════════════════════════════════
// 📖 KONSEP 5: SELECT — Menunggu beberapa channel sekaligus
// ═══════════════════════════════════════════════════════════════════════════
//
// select { } seperti switch, tapi untuk channel operations.
// Ia memilih case yang pertama SIAP. Jika beberapa siap, dipilih acak.
//
//   select {
//   case v := <-ch1:
//       // ch1 punya data
//   case v := <-ch2:
//       // ch2 punya data
//   case <-time.After(1 * time.Second):
//       // tidak ada yang siap dalam 1 detik → timeout
//   default:
//       // tidak ada yang siap sekarang → langsung jalan (non-blocking)
//   }
//
// ═══════════════════════════════════════════════════════════════════════════
// 📖 KONSEP 6: CONTEXT — Kontrol pembatalan & deadline
// ═══════════════════════════════════════════════════════════════════════════
//
// context.Context digunakan untuk:
//   1. Membatalkan goroutine dari luar (cancel)
//   2. Memberi batas waktu (deadline/timeout)
//   3. Meneruskan nilai request-scoped (key-value)
//
//   ctx, cancel := context.WithCancel(context.Background())
//   defer cancel()  // ← SELALU defer cancel untuk cegah goroutine leak!
//
//   go func() {
//       select {
//       case <-ctx.Done():
//           return  // dibatalkan, bersihkan dan keluar
//       case hasil := <-kerjaBerat():
//           kirimHasil(hasil)
//       }
//   }()
//
//   cancel()  // semua goroutine yang dengerin ctx akan berhenti
//
// ═══════════════════════════════════════════════════════════════════════════

import (
	"context"
	"sync"
	"time"
)

// ═══════════════════════════════════════════════════════════════════════════
// 🟢 BAGIAN 1: GOROUTINE + WAITGROUP — Pola paling dasar async Go
// ═══════════════════════════════════════════════════════════════════════════
//
// sync.WaitGroup bekerja seperti "counter menunggu":
//   wg.Add(n)  → tambah counter sebanyak n (sebelum goroutine mulai)
//   wg.Done()  → kurangi counter sebanyak 1 (goroutine selesai)
//   wg.Wait()  → BLOKIR sampai counter = 0 (semua goroutine selesai)

// RunInParallel menjalankan semua fungsi dalam `tasks` secara paralel
// (setiap fungsi dijalankan di goroutine terpisah), lalu menunggu SEMUA
// selesai sebelum fungsi ini return.
//
// Analogi: kamu punya N karyawan. Kamu beri tiap orang 1 tugas sekaligus,
// lalu kamu tunggu sampai semua lapor "selesai".
//
// Contoh:
//
//	hasil := make([]int, 3)
//	RunInParallel(
//	    func() { hasil[0] = 1 },
//	    func() { hasil[1] = 2 },
//	    func() { hasil[2] = 3 },
//	)
//	// setelah return: hasil == [1, 2, 3]
//
// Hint:
//
//	var wg sync.WaitGroup
//	for _, task := range tasks {
//	    wg.Add(1)
//	    go func(t func()) {
//	        defer wg.Done()
//	        t()
//	    }(task)
//	}
//	wg.Wait()
func RunInParallel(tasks ...func()) {
	// TODO: implementasi di sini
	var wg sync.WaitGroup

	for _, task := range tasks {
		wg.Add(1)
		go func(t func()) {
			defer wg.Done()
			t()
		}(task)
	}

	wg.Wait()
}

// ═══════════════════════════════════════════════════════════════════════════
// 🟢 BAGIAN 2: CHANNEL DASAR — "Future/Promise" pattern
// ═══════════════════════════════════════════════════════════════════════════
//
// Pola ini: fungsi langsung return channel, goroutine isi hasilnya nanti.
// Pemanggil bisa kerjakan hal lain dulu, baru ambil hasil ketika butuh.

// AsyncDouble menerima sebuah angka, menghitung n*2 di goroutine terpisah,
// dan mengembalikan hasilnya melalui channel.
//
// Pemanggil mendapat <-chan int (channel read-only). Dia bisa:
//   hasil := <-AsyncDouble(21)  // tunggu sampai goroutine selesai
//
// Atau:
//   ch := AsyncDouble(21)       // mulai async
//   // ... kerjakan hal lain ...
//   hasil := <-ch               // baru ambil hasil ketika butuh
//
// Contoh:
//
//	hasil := <-AsyncDouble(21)  // → 42
//	hasil := <-AsyncDouble(5)   // → 10
//
// Hint:
//   - out := make(chan int, 1)  ← buffered agar goroutine tidak blokir
//   - Goroutine hitung dan kirim: out <- n * 2
//   - return out
func AsyncDouble(n int) <-chan int {
	out := make(chan int, 1) // buffered: goroutine tidak blokir saat kirim

	go func() {
		out <- n * 2
	}()

	return out
}

// ═══════════════════════════════════════════════════════════════════════════
// 🟢 BAGIAN 3: KUMPULKAN HASIL DARI BANYAK GOROUTINE
// ═══════════════════════════════════════════════════════════════════════════

// GatherResults menjalankan semua `producers` secara paralel,
// mengumpulkan semua hasilnya, dan mengembalikan sebagai slice.
// Urutan hasil tidak harus sama dengan urutan producers.
//
// Analogi: kamu kirim 5 kurir ke 5 toko. Siapa cepat dia yang duluan
// balik. Kamu kumpulkan semua belanjaan.
//
// Contoh:
//
//	hasil := GatherResults(
//	    func() int { return 10 },
//	    func() int { return 20 },
//	    func() int { return 30 },
//	)
//	// sort.Ints(hasil) → [10, 20, 30]
//
// Hint:
//   - ch := make(chan int, len(producers))  ← buffered, satu slot per producer
//   - Goroutine kirim hasil ke ch
//   - Kumpulkan len(producers) nilai dari ch ke slice
func GatherResults(producers ...func() int) []int {
	ch := make(chan int, len(producers)) // buffered: semua goroutine bisa kirim tanpa blokir

	for _, p := range producers {
		go func(producer func() int) {
			ch <- producer()
		}(p)
	}

	results := make([]int, len(producers))
	for i := range results {
		results[i] = <-ch
	}

	return results
}

// ═══════════════════════════════════════════════════════════════════════════
// 🟢 BAGIAN 4: SELECT + TIMEOUT — Jangan tunggu selamanya!
// ═══════════════════════════════════════════════════════════════════════════
//
// Ini pola WAJIB untuk HTTP request, query DB, atau operasi apapun
// yang bisa macet. Selalu beri batas waktu!

// WithTimeout menjalankan `work` secara async.
// Jika selesai dalam `timeout` → kembalikan (hasil, true).
// Jika timeout lebih dulu      → kembalikan (0, false).
//
// Contoh:
//
//	// work cepat → berhasil
//	hasil, ok := WithTimeout(func() int {
//	    time.Sleep(10 * time.Millisecond)
//	    return 42
//	}, 1*time.Second)
//	// hasil=42, ok=true
//
//	// work lambat → timeout
//	hasil, ok := WithTimeout(func() int {
//	    time.Sleep(5 * time.Second)
//	    return 99
//	}, 100*time.Millisecond)
//	// hasil=0, ok=false
//
// Hint:
//
//	ch := make(chan int, 1)
//	go func() { ch <- work() }()
//	select {
//	case v := <-ch:
//	    return v, true
//	case <-time.After(timeout):
//	    return 0, false
//	}
func WithTimeout(work func() int, timeout time.Duration) (int, bool) {
	ch := make(chan int, 1)

	go func() {
		ch <- work()
	}()

	select {
	case v := <-ch:
		return v, true
	case <-time.After(timeout):
		return 0, false
	}
}

// ═══════════════════════════════════════════════════════════════════════════
// 🟢 BAGIAN 5: SELECT + DEFAULT — Non-blocking channel read
// ═══════════════════════════════════════════════════════════════════════════
//
// Kadang kamu tidak mau blokir — cukup cek apakah ada data, kalau tidak ada
// langsung lanjut. Gunakan `default` dalam select.

// TryReceive mencoba membaca dari channel tanpa blokir.
// Ada nilai → (nilai, true).
// Tidak ada → (0, false) SEGERA tanpa menunggu.
//
// Contoh:
//
//	ch := make(chan int, 1)
//
//	v, ok := TryReceive(ch)  // kosong → (0, false)
//
//	ch <- 99
//	v, ok = TryReceive(ch)   // ada isi → (99, true)
//
// Hint:
//
//	select {
//	case v := <-ch:
//	    return v, true
//	default:
//	    return 0, false
//	}
func TryReceive(ch <-chan int) (int, bool) {
	select {
	case v := <-ch:
		return v, true
	default:
		return 0, false
	}
}

// ═══════════════════════════════════════════════════════════════════════════
// 🟢 BAGIAN 6: CONTEXT — Batalkan goroutine dari luar
// ═══════════════════════════════════════════════════════════════════════════
//
// context.WithCancel → buat ctx + cancel function
// ctx.Done()         → channel yang ditutup saat cancel() dipanggil
//
// POLA STANDAR di Go:
//   ctx, cancel := context.WithCancel(context.Background())
//   defer cancel()  // ← WAJIB, hindari goroutine leak
//
// Goroutine yang baik selalu "mendengarkan" ctx.Done() dan berhenti jika sinyal datang.

// RunWithCancel menjalankan `work` di goroutine terpisah.
// Mengembalikan:
//   - resultCh: channel untuk menerima hasil
//   - cancel:   fungsi untuk membatalkan work
//
// Contoh:
//
//	// Skenario 1: work selesai normal
//	resultCh, cancel := RunWithCancel(func(ctx context.Context) int {
//	    time.Sleep(10 * time.Millisecond)
//	    return 42
//	})
//	defer cancel()
//	fmt.Println(<-resultCh)  // 42
//
//	// Skenario 2: dibatalkan sebelum selesai
//	resultCh, cancel := RunWithCancel(func(ctx context.Context) int {
//	    select {
//	    case <-ctx.Done():
//	        return -1             // patuhi pembatalan
//	    case <-time.After(1 * time.Hour):
//	        return 999
//	    }
//	})
//	cancel()                   // batalkan segera
//	fmt.Println(<-resultCh)   // -1
//
// Hint:
//   - ctx, cancel := context.WithCancel(context.Background())
//   - ch := make(chan int, 1)
//   - go func() { ch <- work(ctx) }()
//   - return ch, cancel
func RunWithCancel(work func(ctx context.Context) int) (<-chan int, func()) {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int, 1)

	go func() {
		ch <- work(ctx)
	}()

	return ch, cancel
}

// ═══════════════════════════════════════════════════════════════════════════
// 🟢 BAGIAN 7: SEMAPHORE — Batasi goroutine concurrent dengan channel
// ═══════════════════════════════════════════════════════════════════════════
//
// Masalah: 1000 tasks → launch 1000 goroutine → kehabisan resource.
// Solusi: Semaphore — batasi max N goroutine yang aktif sekaligus.
//
// Trik Go: gunakan BUFFERED CHANNEL sebagai semaphore!
//   sem := make(chan struct{}, N)
//   sem <- struct{}{}   // acquire: masuk antrian, blokir jika penuh
//   <-sem               // release: keluar antrian, beri slot untuk yang lain

// BoundedParallel menjalankan semua `tasks`, tapi maksimum `maxConcurrent`
// goroutine yang berjalan bersamaan.
//
// Analogi: antrian kasir. 1000 pembeli, hanya 10 kasir buka.
// Pembeli masuk satu per satu saat kasir kosong.
//
// Contoh:
//
//	var mu sync.Mutex
//	active := 0
//	maxSeen := 0
//	tasks := make([]func(), 20)
//	for i := range tasks {
//	    tasks[i] = func() {
//	        mu.Lock()
//	        active++
//	        if active > maxSeen { maxSeen = active }
//	        mu.Unlock()
//	        time.Sleep(10 * time.Millisecond)
//	        mu.Lock()
//	        active--
//	        mu.Unlock()
//	    }
//	}
//	BoundedParallel(tasks, 5)
//	// maxSeen <= 5  ← tidak pernah lebih dari 5 goroutine aktif!
//
// Hint:
//
//	sem := make(chan struct{}, maxConcurrent)
//	var wg sync.WaitGroup
//	for _, task := range tasks {
//	    wg.Add(1)
//	    sem <- struct{}{}              // acquire slot
//	    go func(t func()) {
//	        defer wg.Done()
//	        defer func() { <-sem }()  // release slot saat selesai
//	        t()
//	    }(task)
//	}
//	wg.Wait()
func BoundedParallel(tasks []func(), maxConcurrent int) {
	sem := make(chan struct{}, maxConcurrent)
	var wg sync.WaitGroup

	for _, task := range tasks {
		wg.Add(1)
		sem <- struct{}{} // acquire: tunggu slot kosong

		go func(t func()) {
			defer wg.Done()
			defer func() { <-sem }() // release: kembalikan slot
			t()
		}(task)
	}

	wg.Wait()
}

// ═══════════════════════════════════════════════════════════════════════════
// 🟢 BAGIAN 8: KUMPULKAN ERROR DARI GOROUTINE
// ═══════════════════════════════════════════════════════════════════════════

// RunAndCollectErrors menjalankan semua `tasks` secara paralel.
// Tiap task bisa return nil (berhasil) atau error (gagal).
// Mengembalikan slice semua error non-nil.
//
// Analogi: kirim 5 perwakilan ke 5 kota. Setelah semua balik, kumpulkan
// laporan masalah. Kalau semua aman, kamu dapat slice kosong.
//
// Contoh:
//
//	import "errors"
//	errs := RunAndCollectErrors(
//	    func() error { return nil },
//	    func() error { return errors.New("gagal A") },
//	    func() error { return nil },
//	    func() error { return errors.New("gagal B") },
//	)
//	// len(errs) == 2
//
// Hint:
//   - ch := make(chan error, len(tasks))  ← buffered
//   - Goroutine kirim hasil (termasuk nil) ke ch
//   - Kumpulkan len(tasks) nilai, filter hanya non-nil
func RunAndCollectErrors(tasks ...func() error) []error {
	ch := make(chan error, len(tasks))

	for _, task := range tasks {
		go func(t func() error) {
			ch <- t()
		}(task)
	}

	var errs []error
	for range tasks {
		if err := <-ch; err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

// ═══════════════════════════════════════════════════════════════════════════
// 🟢 BAGIAN 9: DONE CHANNEL PATTERN — Hentikan goroutine dari luar
// ═══════════════════════════════════════════════════════════════════════════
//
// Pola ini: goroutine berjalan selamanya, kamu beri sinyal berhenti via channel.
// Sinyal: close(stop)  → menutup channel → semua <-stop langsung return

// StreamNumbers menghasilkan angka 1, 2, 3, ... terus-menerus
// sampai channel `stop` ditutup.
//
// Analogi: kran air yang terus mengalir. close(stop) = tutup kran.
//
// Contoh:
//
//	stop := make(chan struct{})
//	numCh := StreamNumbers(stop)
//
//	fmt.Println(<-numCh)  // 1
//	fmt.Println(<-numCh)  // 2
//	fmt.Println(<-numCh)  // 3
//
//	close(stop)           // hentikan stream
//	// numCh akan ditutup setelahnya
//
// Hint:
//
//	out := make(chan int)
//	go func() {
//	    n := 1
//	    for {
//	        select {
//	        case <-stop:
//	            close(out)
//	            return
//	        case out <- n:
//	            n++
//	        }
//	    }
//	}()
//	return out
func StreamNumbers(stop <-chan struct{}) <-chan int {
	out := make(chan int)

	go func() {
		n := 1
		for {
			select {
			case <-stop:
				close(out)
				return
			case out <- n:
				n++
			}
		}
	}()

	return out
}

// ═══════════════════════════════════════════════════════════════════════════
// 🟢 BAGIAN 10: SYNC.ONCE — Jalankan tepat sekali
// ═══════════════════════════════════════════════════════════════════════════
//
// sync.Once memastikan suatu fungsi hanya berjalan SEKALI,
// meskipun dipanggil dari banyak goroutine bersamaan.
// Cocok untuk: inisialisasi singleton, setup database, lazy init.

// SingletonInit menjalankan `init` tepat satu kali menggunakan `once`.
// Meskipun dipanggil 100 goroutine bersamaan, `init` hanya dieksekusi sekali.
//
// Contoh:
//
//	count := 0
//	inc := func() { count++ }
//
//	var once sync.Once
//	var wg sync.WaitGroup
//	for i := 0; i < 100; i++ {
//	    wg.Add(1)
//	    go func() {
//	        defer wg.Done()
//	        SingletonInit(&once, inc)
//	    }()
//	}
//	wg.Wait()
//	// count == 1  ← hanya dieksekusi sekali!
//
// Hint: once.Do(init)
func SingletonInit(once *sync.Once, init func()) {
	once.Do(init)
}

// ═══════════════════════════════════════════════════════════════════════════
// 🟢 BAGIAN 11: CHANNEL DIRECTION — Arah aliran data di channel
// ═══════════════════════════════════════════════════════════════════════════
//
// ❓ PERTANYAAN: Apa bedanya `ch <- nilai` dan `nilai := <-ch` ?
//
// Jawaban simpel:
//   ch <- nilai   →  KIRIM nilai ke dalam channel  (tanda panah ke KANAN ch)
//   nilai := <-ch →  TERIMA nilai dari channel     (tanda panah ke KIRI)
//
// Bayangkan channel sebagai PIPA AIR:
//   ch <- nilai   → kamu TUANG air ke pipa  (masuk)
//   nilai := <-ch → kamu AMBIL air dari pipa (keluar)
//
// ───────────────────────────────────────────────────────────────────────────
// 📌 1. KIRIM ke channel:  ch <- nilai
// ───────────────────────────────────────────────────────────────────────────
//
//   ch := make(chan int)
//   ch <- 42          // ← kirim angka 42 ke dalam channel ch
//                     //   BLOKIR sampai ada goroutine yang siap MENERIMA
//
//   ch <- "halo"      // ← kirim string (jika chan string)
//   ch <- struct{}{}  // ← kirim sinyal kosong (pola umum untuk stop signal)
//
// ───────────────────────────────────────────────────────────────────────────
// 📌 2. TERIMA dari channel:  nilai := <-ch
// ───────────────────────────────────────────────────────────────────────────
//
//   nilai := <-ch             // ← terima dan simpan ke variabel baru
//   nilai, ok := <-ch        // ← ok=false jika channel sudah ditutup
//   <-ch                      // ← terima tapi BUANG nilainya (hanya tunggu sinyal)
//   for v := range ch { }    // ← terima terus sampai channel ditutup
//
// ───────────────────────────────────────────────────────────────────────────
// 📌 3. DIRECTIONAL CHANNEL — Channel dengan arah terbatas di tipe fungsi
// ───────────────────────────────────────────────────────────────────────────
//
//   chan int      → bisa kirim DAN terima (bi-directional)
//   chan<- int    → hanya bisa KIRIM     (send-only)
//   <-chan int    → hanya bisa TERIMA    (receive-only)
//
//   Kenapa pakai directional? → Mencegah bug! Compiler Go akan ERROR
//   jika kamu mencoba kirim ke receive-only channel, atau sebaliknya.
//
//   Contoh:
//     func pengirim(ch chan<- int) {   // hanya bisa kirim
//         ch <- 99                     // ✅ OK
//         nilai := <-ch               // ❌ COMPILE ERROR!
//     }
//
//     func penerima(ch <-chan int) {   // hanya bisa terima
//         nilai := <-ch               // ✅ OK
//         ch <- 99                    // ❌ COMPILE ERROR!
//     }
//
// ═══════════════════════════════════════════════════════════════════════════

// SendOnly adalah contoh fungsi yang hanya MENGIRIM ke channel.
// Parameter `ch chan<- int` berarti fungsi ini HANYA boleh kirim ke ch,
// tidak boleh baca. Compiler akan error jika coba baca dari sini.
//
// Contoh penggunaan:
//
//	ch := make(chan int, 3)
//	SendOnly(ch, 10, 20, 30)
//	// sekarang ch berisi: [10, 20, 30]
//	fmt.Println(<-ch) // → 10
//	fmt.Println(<-ch) // → 20
//	fmt.Println(<-ch) // → 30
func SendOnly(ch chan<- int, values ...int) {
	for _, v := range values {
		ch <- v // ← kirim setiap nilai ke channel
	}
}

// ReceiveOnly adalah contoh fungsi yang hanya MENERIMA dari channel.
// Parameter `ch <-chan int` berarti fungsi ini HANYA boleh terima dari ch.
// Mengumpulkan `n` nilai dari channel dan return sebagai slice.
//
// Contoh penggunaan:
//
//	ch := make(chan int, 3)
//	ch <- 100
//	ch <- 200
//	ch <- 300
//	hasil := ReceiveOnly(ch, 3)
//	// hasil → [100, 200, 300]
func ReceiveOnly(ch <-chan int, n int) []int {
	hasil := make([]int, n)
	for i := 0; i < n; i++ {
		hasil[i] = <-ch // ← terima nilai dari channel, simpan ke slice
	}
	return hasil
}

// ───────────────────────────────────────────────────────────────────────────
// 🔗 PIPELINE PATTERN — Sambung channel seperti rantai pipa
// ───────────────────────────────────────────────────────────────────────────
//
// Pipeline = serangkaian goroutine yang dihubungkan channel.
// Output satu goroutine → menjadi input goroutine berikutnya.
//
//   goroutine A --[ch1]--> goroutine B --[ch2]--> goroutine C
//
// Analogi: conveyor belt pabrik. Setiap stasiun proses 1 hal,
// lalu lempar ke stasiun berikutnya.
//
//   [Generator] --chan--> [Doubler] --chan--> [Adder] --chan--> hasil
//
// ───────────────────────────────────────────────────────────────────────────

// Generate menghasilkan angka-angka dari `nums` ke dalam channel.
// Ini adalah TAHAP PERTAMA pipeline (sumber data).
//
// Contoh:
//
//	ch := Generate(1, 2, 3, 4, 5)
//	// ch akan mengalirkan: 1, 2, 3, 4, 5 lalu channel ditutup
func Generate(nums ...int) <-chan int {
	out := make(chan int, len(nums)) // buffered agar tidak blokir
	go func() {
		for _, n := range nums {
			out <- n // ← kirim tiap angka ke channel output
		}
		close(out) // ← tutup channel jika semua sudah dikirim
	}()
	return out // ← kembalikan receive-only channel ke pemanggil
}

// Double mengalikan tiap nilai dari channel input dengan 2,
// dan mengalirkan hasilnya ke channel output baru.
// Ini adalah TAHAP TENGAH pipeline (transformasi data).
//
// Contoh:
//
//	src := Generate(1, 2, 3)     // menghasilkan: 1, 2, 3
//	doubled := Double(src)        // menghasilkan: 2, 4, 6
//	for v := range doubled {
//	    fmt.Println(v)            // → 2, 4, 6
//	}
func Double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in { // ← range channel: baca sampai in ditutup
			out <- v * 2   //   kirim hasil ke channel output
		}
		close(out) // ← tutup output jika input sudah habis
	}()
	return out
}

// AddN menambahkan `n` ke tiap nilai dari channel input.
// Ini contoh lain tahap transformasi pipeline.
//
// Contoh pipeline lengkap:
//
//	// Tahap 1: Generate angka 1..5
//	src := Generate(1, 2, 3, 4, 5)
//
//	// Tahap 2: Kalikan 2  → 2, 4, 6, 8, 10
//	doubled := Double(src)
//
//	// Tahap 3: Tambah 10  → 12, 14, 16, 18, 20
//	result := AddN(doubled, 10)
//
//	for v := range result {
//	    fmt.Println(v) // → 12, 14, 16, 18, 20
//	}
func AddN(in <-chan int, n int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in { // ← baca dari channel input
			out <- v + n    //   kirim ke channel output
		}
		close(out)
	}()
	return out
}

// ───────────────────────────────────────────────────────────────────────────
// 🔀 MERGE/FAN-IN PATTERN — Gabungkan beberapa channel jadi satu
// ───────────────────────────────────────────────────────────────────────────
//
//   channel A ──┐
//               ├──> [merge] ──> satu channel gabungan
//   channel B ──┘
//
// Berguna ketika kamu punya beberapa sumber data dan ingin
// proses semua hasilnya dari satu tempat.
//
// ───────────────────────────────────────────────────────────────────────────

// Merge menggabungkan DUA channel menjadi SATU channel.
// Nilai dari ch1 dan ch2 dialirkan ke channel output bersamaan.
// Output channel ditutup setelah kedua input ditutup.
//
// Contoh:
//
//	angkaGanjil  := Generate(1, 3, 5)    // channel pertama
//	angkaGenap   := Generate(2, 4, 6)    // channel kedua
//	semua        := Merge(angkaGanjil, angkaGenap)
//
//	var hasil []int
//	for v := range semua {
//	    hasil = append(hasil, v)
//	}
//	sort.Ints(hasil)
//	// hasil → [1, 2, 3, 4, 5, 6]
func Merge(ch1, ch2 <-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	// Fungsi helper: forward semua nilai dari `src` ke `out`
	forward := func(src <-chan int) {
		defer wg.Done()
		for v := range src { // ← baca semua nilai dari src
			out <- v          //   forwad ke out
		}
	}

	wg.Add(2)
	go forward(ch1) // ← goroutine 1: drain channel pertama
	go forward(ch2) // ← goroutine 2: drain channel kedua

	// Goroutine penjaga: tutup out saat kedua sumber selesai
	go func() {
		wg.Wait()  // ← tunggu kedua goroutine forward selesai
		close(out) // ← baru tutup channel output
	}()

	return out
}

// ═══════════════════════════════════════════════════════════════════════════
// 📝 RINGKASAN OPERATOR CHANNEL
// ═══════════════════════════════════════════════════════════════════════════
//
//  OPERATOR        ARTI                           BLOKIR?
//  ─────────────── ────────────────────────────── ────────────────────────
//  ch <- val       Kirim val ke channel            Ya (unbuffered/penuh)
//  val := <-ch     Terima dari channel             Ya (kosong)
//  val, ok := <-ch Terima + cek channel terbuka   Ya (kosong)
//  <-ch            Terima dan buang (tunggu sinyal) Ya (kosong)
//  close(ch)       Tutup channel                   Tidak
//  range ch        Baca sampai channel ditutup     Ya per iterasi
//
//  TIPE               BOLEH KIRIM   BOLEH TERIMA
//  ─────────────────  ───────────   ────────────
//  chan int           ✅            ✅
//  chan<- int         ✅            ❌ (compile error)
//  <-chan int         ❌ (error)    ✅
//
//  CATATAN: `<<` BUKAN operator channel!
//    << adalah BITSHIFT KIRI:  1 << 3 = 8  (geser bit ke kiri 3 posisi)
//    Operator channel Go hanya: <-
// ═══════════════════════════════════════════════════════════════════════════
