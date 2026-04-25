package belajar

// ==================== DAY 24: CHANNEL LANJUTAN — LATIHAN MANDIRI ====================
//
// 🎯 TUJUAN HARI INI:
//   Melanjutkan pemahaman channel dengan pola-pola yang lebih kompleks dan
//   sering dipakai di dunia nyata. Implementasikan setiap fungsi, lalu
//   jalankan test untuk memvalidasi hasilmu:
//
//     cd "DAY 24"
//     go test ./... -v
//
// ═══════════════════════════════════════════════════════════════════════════════
// 📖 REVIEW KILAT: POLA CHANNEL YANG SUDAH DIKETAHUI
// ═══════════════════════════════════════════════════════════════════════════════
//
//  DAY 22 ✅  goroutine + WaitGroup, Future/Promise, timeout, semaphore
//  DAY 23 ✅  fan-out, fan-in, pipeline, worker pool, done channel
//
// ═══════════════════════════════════════════════════════════════════════════════
// 📖 POLA BARU HARI INI
// ═══════════════════════════════════════════════════════════════════════════════
//
//  1. ORDERED FAN-OUT  — fan-out tapi hasil tetap urut seperti input
//  2. BATCH PROCESSOR  — kumpulkan N item lalu proses sekaligus
//  3. RETRY PATTERN    — coba ulang otomatis jika fungsi gagal
//  4. RATE LIMITER     — batasi berapa operasi per satuan waktu
//  5. BROADCAST        — satu pengirim, banyak penerima (pub-sub sederhana)
//  6. OR-DONE          — gabungkan banyak done channel jadi satu
//  7. TIMED GENERATOR  — hasilkan nilai berkala setiap interval waktu
//
// ═══════════════════════════════════════════════════════════════════════════════

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 1: ORDERED FAN-OUT
// ═══════════════════════════════════════════════════════════════════════════════
//
// Masalah: KuadratAsync di DAY 23 menghasilkan nilai yang URUTANNYA ACAK
//          karena goroutine mana yang selesai duluan tidak bisa ditebak.
//
// Solusi: alih-alih satu channel bersama, buat SATU channel per item.
//         Goroutine ke-i kirim hasilnya ke channel ke-i.
//         Kita baca hasilnya secara berurutan: channel[0], channel[1], dst.
//         → Tiap goroutine masih jalan paralel, tapi urutan terjaga!

// MapOrdered menjalankan `fn(nums[i])` untuk setiap nums[i] secara PARALEL,
// dan mengembalikan slice hasil dengan URUTAN yang SAMA seperti input.
//
// Berbeda dengan KuadratAsync (DAY 23) yang urutannya acak —
// fungsi ini menjamin hasil[0] = fn(nums[0]), hasil[1] = fn(nums[1]), dst.
//
// Contoh:
//
//	hasil := MapOrdered(func(n int) int { return n * n }, 1, 2, 3, 4, 5)
//	// hasil → [1, 4, 9, 16, 25]  ← urutan pasti seperti input
//
// Hint:
//   - Buat slice of channel, satu channel per elemen input.
//   - Tiap goroutine kirim hasilnya ke channel di indeks yang sama.
//   - Kumpulkan hasilnya dengan membaca channel satu per satu secara berurutan.
func MapOrdered(fn func(int) int, nums ...int) []int {
	// TODO: implementasi di sini
	result := []int{}
	var wg sync.WaitGroup
	wg.Add(len(nums))

	for _, v := range nums {
		go func() {
			defer wg.Done()
			result = append(result, fn(v))
		}()
	}

	go func() {
		wg.Wait()
	}()

	return result
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 2: BATCH PROCESSOR
// ═══════════════════════════════════════════════════════════════════════════════
//
// Pola ini berguna ketika proses per-item mahal, tapi proses per-batch lebih efisien.
// Contoh nyata: batch insert ke database, mengirim notifikasi dalam kelompok.
//
// Cara kerja:
//   Kumpulkan item dari channel input sampai:
//     (a) batch sudah penuh sebanyak `ukuran` item → kirim batch
//     (b) channel input ditutup dan masih ada sisa item → kirim sisa

// BatchChannel membaca nilai dari `masuk` dan mengelompokkannya menjadi
// batch berukuran maksimum `ukuran`. Setiap batch dikirim ke channel output.
// Channel output ditutup setelah channel input ditutup dan semua batch terkirim.
//
// Contoh (ukuran=3, masuk = 1,2,3,4,5):
//
//	masuk := make(chan int, 5)
//	for _, v := range []int{1, 2, 3, 4, 5} { masuk <- v }
//	close(masuk)
//
//	batchCh := BatchChannel(masuk, 3)
//	fmt.Println(<-batchCh) // → [1 2 3]
//	fmt.Println(<-batchCh) // → [4 5]
//
// Hint:
//   - Gunakan goroutine + `for v := range masuk` untuk membaca input.
//   - Kumpulkan item ke slice sementara (batch).
//   - Saat batch mencapai `ukuran`, kirim ke output dan reset batch.
//   - Jangan lupa kirim sisa batch setelah loop selesai.
func BatchChannel(masuk <-chan int, ukuran int) <-chan []int {
	// TODO: implementasi di sini
	var wg sync.WaitGroup
	wg.Add(int(math.Ceil(float64(len(masuk)) / float64(ukuran))))
	ch := make(chan []int)

	for i := 0; i < int(math.Ceil(float64(len(masuk))/float64(ukuran))); i++ {
		go func() {
			defer wg.Done()
			arr := []int{}
			for j := 0; j < ukuran; j++ {
				data := <-masuk
				if data != 0 {
					arr = append(arr, data)
				}
			}
			fmt.Println(arr)
			ch <- arr
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 3: RETRY PATTERN
// ═══════════════════════════════════════════════════════════════════════════════
//
// Di dunia nyata, operasi bisa gagal sementara (jaringan putus, API limit, dsb).
// Retry = coba lagi otomatis hingga berhasil atau batas percobaan habis.
//
// Skema:
//   percobaan 1 → gagal → tunggu delay → percobaan 2 → berhasil ✅

// CobaUlang menjalankan `fn` secara berulang sampai berhasil (error == nil)
// atau sudah mencoba sebanyak `maksPercobaan` kali.
// Setiap percobaan yang gagal ditunggu selama `delay` sebelum mencoba lagi.
// Mengembalikan hasil terakhir dan error terakhir.
//
// Contoh:
//
//	count := 0
//	hasil, err := CobaUlang(func() (int, error) {
//	    count++
//	    if count < 3 {
//	        return 0, errors.New("gagal")
//	    }
//	    return 42, nil
//	}, 5, 0)
//	// hasil=42, err=nil, count=3
//
// Hint:
//   - Gunakan loop sebanyak `maksPercobaan` kali.
//   - Jika fn() berhasil (err == nil), langsung return.
//   - Jika gagal dan masih ada percobaan tersisa, tunggu `delay`.
//   - Return nilai dan error dari percobaan terakhir.
func CobaUlang(fn func() (int, error), maksPercobaan int, delay time.Duration) (int, error) {
	// TODO: implementasi di sini
	for i := 0; i < maksPercobaan; i++ {
		count, err := fn()
		if err != nil {
			time.Sleep(delay)
		} else {
			return count, err
		}

	}

	return 0, fmt.Errorf("gagal")
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 4: RATE LIMITER
// ═══════════════════════════════════════════════════════════════════════════════
//
// Rate Limiter = batasi berapa banyak operasi yang boleh berjalan per waktu.
// Contoh nyata: API hanya boleh dipanggil 5x per detik.
//
// Trik Go: gunakan time.Ticker → setiap `interval`, kirim sinyal ke channel.
// Pemanggil harus "ambil token dulu" sebelum boleh menjalankan operasi.
//
// Ilustrasi:
//   token tersedia:  [---][---][---]     (satu tiap interval)
//   worker ambil:      ↓    ↓    ↓
//   hasil:           op1  op2  op3       (ter-throttle, tidak sekaligus)

// BuatRateLimiter mengembalikan channel yang mengirimkan satu token (struct{})
// setiap `interval`. Pemanggil harus membaca dari channel ini sebelum
// menjalankan operasi untuk memastikan rate terjaga.
// Channel berjalan terus sampai `selesai` ditutup.
//
// Contoh:
//
//	selesai := make(chan struct{})
//	limiter := BuatRateLimiter(100*time.Millisecond, selesai)
//
//	for i := 0; i < 3; i++ {
//	    <-limiter       // tunggu token
//	    fmt.Println("operasi", i+1)
//	}
//	close(selesai)
//
// Hint:
//   - Buat goroutine yang loop selamanya sampai `selesai` ditutup.
//   - Gunakan time.NewTicker(interval) untuk mendapat sinyal berkala.
//   - Setiap tick → kirim struct{}{} ke channel output.
//   - Jangan lupa ticker.Stop() untuk menghindari goroutine leak.
func BuatRateLimiter(interval time.Duration, selesai <-chan struct{}) <-chan struct{} {
	// TODO: implementasi di sini
	var wg sync.WaitGroup
	ch := make(chan struct{})

	go func() {
		ticker := time.NewTicker(interval)
		defer wg.Done()
		defer close(ch)
		wg.Add(1)

		for {
			select {
			case <-ticker.C:
				ch <- struct{}{}
			case <-selesai:
				return
			}
		}
	}()

	go func() {
		wg.Wait()
	}()
	return ch
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 5: BROADCAST (Pub-Sub Sederhana)
// ═══════════════════════════════════════════════════════════════════════════════
//
// Pola: SATU pengirim, BANYAK penerima.
// Setiap pesan yang dikirim harus diterima oleh SEMUA subscriber.
//
// Di DAY 23, SebarKeSemua sudah mirip ini — tapi Broadcaster di sini
// lebih fleksibel: subscriber bisa ditambah dinamis sebelum siaran dimulai.

// Broadcaster adalah struktur untuk mengelola broadcast ke banyak subscriber.
type Broadcaster struct {
	mu          sync.Mutex
	subscribers []chan int
}

// Subscribe mendaftarkan subscriber baru dan mengembalikan channel-nya.
// Setiap nilai yang di-broadcast akan dikirim ke channel ini.
// `buffer` menentukan kapasitas channel subscriber.
//
// Contoh:
//
//	b := &Broadcaster{}
//	ch1 := b.Subscribe(10)
//	ch2 := b.Subscribe(10)
//	// sekarang b punya 2 subscriber
//
// Hint:
//   - Buat channel baru dengan kapasitas `buffer`.
//   - Tambahkan ke daftar subscriber (jangan lupa lock mutex).
//   - Return channel tersebut sebagai receive-only.
func (b *Broadcaster) Subscribe(buffer int) <-chan int {
	// TODO: implementasi di sini
	ch := make(chan int, buffer)

	b.mu.Lock()
	b.subscribers = append(b.subscribers, ch)
	b.mu.Unlock()

	return ch
}

// Broadcast mengirimkan `nilai` ke SEMUA subscriber yang terdaftar.
//
// Contoh:
//
//	b := &Broadcaster{}
//	ch1 := b.Subscribe(1)
//	ch2 := b.Subscribe(1)
//	b.Broadcast(99)
//	fmt.Println(<-ch1) // → 99
//	fmt.Println(<-ch2) // → 99
//
// Hint:
//   - Lock mutex untuk akses aman ke daftar subscriber.
//   - Iterasi setiap subscriber dan kirim nilai ke channel-nya.
func (b *Broadcaster) Broadcast(nilai int) {
	// TODO: implementasi di sini
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, ch := range b.subscribers {
		ch <- nilai
	}
}

// Close menutup semua channel subscriber.
// Dipanggil oleh pengirim saat tidak ada data lagi yang akan di-broadcast.
//
// Hint:
//   - Lock mutex, lalu close setiap channel di daftar subscriber.
func (b *Broadcaster) Close() {
	// TODO: implementasi di sini
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, ch := range b.subscribers {
		close(ch)
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 6: OR-DONE CHANNEL
// ═══════════════════════════════════════════════════════════════════════════════
//
// Masalah: kamu punya N goroutine dengan N channel done terpisah.
//          Kamu ingin tahu saat SALAH SATU dari mereka selesai.
//
// Solusi: OrDone — gabungkan banyak done channel menjadi satu.
//         Begitu SATU channel ditutup, channel output ikut ditutup.
//
// Ilustrasi:
//   done1 ──┐
//   done2 ──┤──> orDone  (ditutup saat done1 ATAU done2 ATAU done3 ditutup)
//   done3 ──┘

// OrDone mengembalikan sebuah channel yang ditutup ketika SALAH SATU
// dari channel `dones` ditutup.
//
// Contoh:
//
//	d1 := make(chan struct{})
//	d2 := make(chan struct{})
//	d3 := make(chan struct{})
//	combined := OrDone(d1, d2, d3)
//
//	close(d2)      // tutup salah satu
//	<-combined     // langsung lolos karena d2 sudah tutup
//
// Hint:
//   - Kasus dasar: 0 channel → return channel yang tidak pernah tutup.
//     1 channel → return channel itu langsung.
//   - Kasus umum: spawn goroutine yang tunggu dones[0] ATAU OrDone(dones[1:]...).
//   - Gunakan select untuk menunggu yang pertama selesai, lalu close output.
//   - Pendekatan rekursif membuat ini sangat ringkas.
func OrDone(dones ...<-chan struct{}) <-chan struct{} {
	// TODO: implementasi di sini
	ch := make(chan struct{})

	go func() {
		defer close(ch)

		switch len(dones) {
		case 0:
			return
		case 1:
			<-dones[0]
		default:
			select {
			case <-dones[0]:
			case <-dones[1]:
			case <-OrDone(dones[2:]...):
			}
		}
	}()

	return ch
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 7: TIMED GENERATOR
// ═══════════════════════════════════════════════════════════════════════════════
//
// Berbeda dengan GeneratorAngka (DAY 23) yang mengirim data secepat mungkin,
// TimedGenerator mengirim nilai secara BERKALA setiap `interval` waktu.
//
// Kegunaan nyata:
//   - Health-check ping setiap 30 detik
//   - Polling status setiap 1 menit
//   - Progress update setiap 100ms

// TimedGenerator menghasilkan nilai dari `fn()` setiap `interval` waktu
// dan mengirimkannya ke channel output. Berhenti saat `selesai` ditutup.
// Channel output ditutup saat generator berhenti.
//
// Contoh:
//
//	selesai := make(chan struct{})
//	counter := 0
//	ch := TimedGenerator(func() int {
//	    counter++
//	    return counter
//	}, 50*time.Millisecond, selesai)
//
//	fmt.Println(<-ch) // → 1  (setelah ~50ms)
//	fmt.Println(<-ch) // → 2  (setelah ~100ms)
//	close(selesai)
//
// Hint:
//   - Kombinasikan pola GeneratorAngka (DAY 23) dengan time.NewTicker.
//   - Setiap tick → panggil fn() dan kirim hasilnya ke output.
//   - Jika `selesai` ditutup → hentikan generator dan tutup output.
//   - Hati-hati: saat mengirim ke output, `selesai` mungkin juga sudah tutup
//     → gunakan select lagi saat pengiriman untuk menghindari goroutine leak.
func TimedGenerator(fn func() int, interval time.Duration, selesai <-chan struct{}) <-chan int {
	// TODO: implementasi di sini
	ch := make(chan int)

	go func() {
		defer close(ch)

		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-selesai:
				return
			case <-ticker.C:
				ch <- fn()
			}
		}
	}()
	return ch
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 8: PIPELINE DENGAN FILTER
// ═══════════════════════════════════════════════════════════════════════════════
//
// Melengkapi pipeline DAY 23 (Sumber → Kalikan → Tambahkan)
// dengan tahap FILTER: hanya loloskan nilai yang memenuhi kondisi tertentu.

// Filter membaca tiap nilai dari `masuk`, dan hanya meneruskan nilai
// yang membuat `predikat(nilai)` mengembalikan true.
// Channel output ditutup saat channel input ditutup.
//
// Contoh:
//
//	src    := Sumber(1, 2, 3, 4, 5, 6)
//	genap  := Filter(src, func(n int) bool { return n%2 == 0 })
//	for v := range genap { fmt.Println(v) }
//	// → 2, 4, 6
//
// Hint:
//   - Pola sama seperti Kalikan/Tambahkan di DAY 23 — goroutine + range masuk.
//   - Bedanya: kirim ke output HANYA JIKA predikat(v) == true.
func Filter(masuk <-chan int, predikat func(int) bool) <-chan int {
	// TODO: implementasi di sini
	ch := make(chan int)
	var wg sync.WaitGroup
	go func() {
		for v := range masuk {
			wg.Add(1)
			go func(val int) {
				defer wg.Done()
				if predikat(val) {
					ch <- val
				}
			}(v)
		}
		wg.Wait()
		close(ch)
	}()

	return ch
}

// Reduce membaca semua nilai dari `masuk` dan menggabungkannya menjadi
// satu nilai menggunakan fungsi `akumulator(sejauhIni, nilaiBaru)`.
// `awal` adalah nilai awal akumulator. Fungsi ini BLOKIR sampai channel ditutup.
//
// Contoh (sum):
//
//	src := Sumber(1, 2, 3, 4, 5)
//	total := Reduce(src, func(acc, v int) int { return acc + v }, 0)
//	// total → 15
//
// Contoh (max):
//
//	src := Sumber(3, 1, 4, 1, 5, 9, 2, 6)
//	maks := Reduce(src, func(acc, v int) int {
//	    if v > acc { return v }
//	    return acc
//	}, 0)
//	// maks → 9
//
// Hint:
//   - Baca semua nilai dari `masuk` dengan range.
//   - Setiap nilai → update akumulator: awal = akumulator(awal, v).
//   - Return nilai akumulator setelah loop selesai.
func Reduce(masuk <-chan int, akumulator func(int, int) int, awal int) int {
	// TODO: implementasi di sini
	var wg sync.WaitGroup
	var mu sync.Mutex

	for v := range masuk {
		wg.Add(1)

		go func(val int) {
			defer wg.Done()

			mu.Lock()
			awal = akumulator(awal, val)
			mu.Unlock()
		}(v)
	}

	wg.Wait()

	return awal
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 9: CHANNEL DENGAN NILAI TERAKHIR (Latest Value)
// ═══════════════════════════════════════════════════════════════════════════════
//
// Masalah: producer sangat cepat, consumer lambat.
//          Consumer tidak perlu semua nilai — cukup nilai TERBARU.
//
// Analogi: live feed harga saham. Kamu hanya peduli harga SEKARANG,
//          bukan semua pergerakan harga yang terlewat.
//
// Teknik: goroutine "relay" yang selalu simpan nilai terbaru,
//         dan kirim ke consumer saat consumer siap.

// Terbaru membungkus channel `masuk` dengan channel baru (buffered kapasitas 1)
// yang hanya meneruskan nilai TERBARU.
// Jika ada nilai baru datang sebelum consumer sempat baca, nilai lama dibuang.
// Channel output ditutup saat channel input ditutup.
//
// Contoh:
//
//	masuk := make(chan int, 10)
//	terbaru := Terbaru(masuk)
//
//	masuk <- 1
//	masuk <- 2
//	masuk <- 3
//	close(masuk)
//
//	// Consumer lambat — dapat nilai valid (tidak deadlock)
//	fmt.Println(<-terbaru) // bisa 1, 2, atau 3 — yang terpenting tidak deadlock
//
// Hint:
//   - Gunakan buffered channel kapasitas 1 sebagai output.
//   - Goroutine relay: untuk setiap nilai dari masuk, coba kosongkan buffer
//     output dulu (non-blocking) sebelum menaruh nilai baru.
//   - Cara mengosongkan: select { case <-out: default: } → buang nilai lama jika ada.
func Terbaru(masuk <-chan int) <-chan int {
	// TODO: implementasi di sini
	ch := make(chan int)

	for v := range masuk {

		go func() {
			ch <- v
		}()
	}

	return ch
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🔧 HELPER (sudah diimplementasikan, boleh dipakai di kode kamu)
// ═══════════════════════════════════════════════════════════════════════════════

// Sumber membuat channel yang mengalirkan angka-angka dari `data`,
// lalu menutup channel setelah semua terkirim.
// (Sama seperti DAY 23 — disertakan agar package berdiri sendiri)
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

// buatChannelD24 adalah helper untuk membuat channel yang sudah berisi data.
func buatChannelD24(data ...int) <-chan int {
	ch := make(chan int, len(data))
	for _, v := range data {
		ch <- v
	}
	close(ch)
	return ch
}

// _ dipakai agar import tidak error sebelum implementasi.
var _ = time.Second
var _ sync.Mutex
