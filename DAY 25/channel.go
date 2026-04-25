package belajar

// ==================== DAY 25: CHANNEL & URUTAN — DARI NOL ====================
//
// 🎯 FOKUS HARI INI:
//   Memahami SATU masalah yang paling sering bikin bingung:
//   "Bagaimana menjaga urutan hasil saat goroutine berjalan paralel?"
//
//   Jalankan test:
//     cd "DAY 25"
//     go test ./... -v
//
// ═══════════════════════════════════════════════════════════════════════════════
// 🧠 KONSEP INTI — BACA INI DULU SEBELUM MULAI NGODING
// ═══════════════════════════════════════════════════════════════════════════════
//
// CHANNEL itu seperti PIPA AIR:
//
//   pengirim  →  [~~~~pipa~~~~]  →  penerima
//
//   ch := make(chan int)     // buat pipa
//   ch <- 42                // kirim (pengirim BLOKIR sampai ada yang baca)
//   v  := <-ch              // terima (penerima BLOKIR sampai ada yang kirim)
//
// BUFFERED CHANNEL = pipa dengan tangki:
//
//   ch := make(chan int, 3)  // tangki muat 3
//   ch <- 1  // langsung masuk tangki, tidak blokir
//   ch <- 2  // langsung masuk tangki, tidak blokir
//   ch <- 3  // langsung masuk tangki, tidak blokir
//   ch <- 4  // BLOKIR! tangki penuh
//
// ═══════════════════════════════════════════════════════════════════════════════
// 🚨 MASALAH URUTAN — INI YANG SERING BIKIN BINGUNG
// ═══════════════════════════════════════════════════════════════════════════════
//
// Bayangkan kamu mau hitung kuadrat [1, 2, 3] secara paralel:
//
//   CARA SALAH (urutan tidak terjaga):
//   ────────────────────────────────
//   ch := make(chan int, 3)
//   for _, v := range []int{1, 2, 3} {
//       go func(n int) { ch <- n * n }(v)
//   }
//   // Hasilnya bisa: [4, 1, 9] atau [9, 4, 1] atau apapun!
//   // Karena goroutine mana yang selesai duluan tidak bisa diprediksi.
//
//   CARA BENAR (urutan terjaga):
//   ────────────────────────────
//   channels := make([]chan int, 3)   // satu channel PER item
//   for i, v := range []int{1, 2, 3} {
//       channels[i] = make(chan int, 1)
//       go func(ch chan int, n int) { ch <- n * n }(channels[i], v)
//   }
//   // Baca BERURUTAN: channels[0], channels[1], channels[2]
//   hasil[0] = <-channels[0]  // pasti hasil dari input[0]
//   hasil[1] = <-channels[1]  // pasti hasil dari input[1]
//   hasil[2] = <-channels[2]  // pasti hasil dari input[2]
//   // Hasilnya: [1, 4, 9] — PASTI urut!
//
// Kunci: goroutine KE-I kirim ke channel KE-I,
//        kita baca channel KE-0, KE-1, KE-2 secara berurutan.
//
// ═══════════════════════════════════════════════════════════════════════════════

import (
	"sync"
	"time"
)

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 1: CHANNEL DASAR — KIRIM DAN TERIMA
// ═══════════════════════════════════════════════════════════════════════════════
//
// Sebelum urusan paralel, pastikan kamu paham channel paling basic.
//

// KirimSatu mengirim nilai `n` ke channel dan mengembalikan channel tersebut.
// Channel sudah ditutup setelah nilai dikirim.
//
// Contoh:
//
//	ch := KirimSatu(42)
//	fmt.Println(<-ch) // → 42
//
// Hint:
//   - Buat channel BUFFERED kapasitas 1 agar tidak blokir.
//   - Kirim `n` ke channel.
//   - Tutup channel (close).
//   - Return channel.
func KirimSatu(n int) <-chan int {
	// TODO: implementasi di sini
	ch := make(chan int)

	ch <- n
	close(ch)

	return ch
}

// KirimBanyak mengirim semua nilai dari `data` ke channel dan menutupnya.
// Pengiriman harus SEQUENTIAL (satu per satu, berurutan), bukan paralel.
//
// Contoh:
//
//	ch := KirimBanyak(1, 2, 3, 4, 5)
//	for v := range ch {
//	    fmt.Println(v) // → 1, 2, 3, 4, 5 (pasti urut)
//	}
//
// Hint:
//   - Buat channel unbuffered (atau buffered).
//   - Jalankan goroutine yang kirim data satu per satu, lalu close.
//   - Return channel.
func KirimBanyak(data ...int) <-chan int {
	// TODO: implementasi di sini
	ch := make(chan int, len(data))

	for _, v := range data {
		ch <- v
	}

	close(ch)
	return ch
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 2: GOROUTINE + CHANNEL SEDERHANA
// ═══════════════════════════════════════════════════════════════════════════════
//
// Jalankan satu goroutine, ambil hasilnya via channel.
//

// HitungDiBackground menjalankan `fn()` di goroutine terpisah
// dan mengembalikan channel yang akan berisi hasilnya.
//
// Contoh:
//
//	ch := HitungDiBackground(func() int {
//	    time.Sleep(10 * time.Millisecond)
//	    return 99
//	})
//	hasil := <-ch // tunggu goroutine selesai
//	// hasil == 99
//
// Hint:
//   - Buat buffered channel kapasitas 1.
//   - Jalankan goroutine: panggil fn(), kirim hasilnya ke channel.
//   - Return channel (tanpa menunggu goroutine selesai).
func HitungDiBackground(fn func() int) <-chan int {
	// TODO: implementasi di sini
	return nil
}

// JalankanN menjalankan `fn` sebanyak `n` kali secara BERSAMAAN (paralel).
// Mengembalikan slice semua hasil.
// URUTAN TIDAK PERLU DIJAGA — boleh acak.
//
// Contoh:
//
//	counter := 0
//	var mu sync.Mutex
//	hasil := JalankanN(func() int {
//	    mu.Lock(); defer mu.Unlock()
//	    counter++
//	    return counter
//	}, 5)
//	sort.Ints(hasil)
//	// hasil = [1, 2, 3, 4, 5]
//
// Hint:
//   - Buat buffered channel kapasitas n.
//   - Jalankan n goroutine, masing-masing kirim fn() ke channel.
//   - Baca n hasil dari channel, masukkan ke slice.
//   - Return slice.
func JalankanN(fn func() int, n int) []int {
	// TODO: implementasi di sini
	return nil
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 3: MENJAGA URUTAN — INTI DAY 25
// ═══════════════════════════════════════════════════════════════════════════════
//
// Ini bagian terpenting. Baca penjelasan di atas (bagian 🚨) sebelum mulai.
//

// TransformasiUrut menerapkan fn ke setiap elemen input secara PARALEL,
// tapi mengembalikan hasil dengan URUTAN YANG SAMA seperti input.
//
// Contoh:
//
//	hasil := TransformasiUrut([]int{3, 1, 4, 1, 5}, func(n int) int {
//	    return n * 2
//	})
//	// hasil = [6, 2, 8, 2, 10] ← urutan PASTI sama dengan input
//
// Hint:
//   - Buat []chan int, satu channel per elemen. Tiap channel buffered kapasitas 1.
//   - Untuk setiap i, jalankan goroutine yang kirim fn(input[i]) ke channels[i].
//   - Baca hasil secara berurutan: hasil[0] = <-channels[0], dst.
//   - Return slice hasil.
func TransformasiUrut(input []int, fn func(int) int) []int {
	// TODO: implementasi di sini
	return nil
}

// TransformasiAcak menerapkan fn ke setiap elemen input secara PARALEL.
// Urutan hasil TIDAK DIJAMIN — boleh berbeda dari urutan input.
//
// Contoh:
//
//	hasil := TransformasiAcak([]int{1, 2, 3}, func(n int) int { return n * n })
//	sort.Ints(hasil)
//	// hasil = [1, 4, 9] (setelah diurutkan)
//
// Hint:
//   - Cukup satu buffered channel kapasitas len(input).
//   - Semua goroutine kirim ke channel yang SAMA.
//   - Baca len(input) kali dari channel.
//   - Bandingkan dengan TransformasiUrut: bedanya di sini kita pakai 1 channel,
//     di TransformasiUrut kita pakai N channel.
func TransformasiAcak(input []int, fn func(int) int) []int {
	// TODO: implementasi di sini
	return nil
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 4: PIPELINE SEDERHANA
// ═══════════════════════════════════════════════════════════════════════════════
//
// Pipeline = rangkaian tahap yang dihubungkan channel.
// Data mengalir dari satu tahap ke tahap berikutnya seperti air di pipa.
//
//   [Sumber] → ch1 → [Tahap1] → ch2 → [Tahap2] → hasil
//
// Setiap tahap berjalan di goroutine sendiri.
// Saat channel input ditutup, tahap tersebut selesai dan tutup channel outputnya.
//

// Gandakan membaca setiap nilai dari `masuk` dan mengirim nilai * 2 ke output.
// Channel output ditutup saat channel input ditutup.
//
// Contoh:
//
//	src := KirimBanyak(1, 2, 3)
//	doubled := Gandakan(src)
//	for v := range doubled {
//	    fmt.Println(v) // → 2, 4, 6
//	}
//
// Hint:
//   - Buat channel output (unbuffered).
//   - Jalankan goroutine: for v := range masuk → kirim v*2 ke output.
//   - Setelah loop selesai (masuk ditutup), close output.
//   - Return output.
func Gandakan(masuk <-chan int) <-chan int {
	// TODO: implementasi di sini
	return nil
}

// Tambah membaca setiap nilai dari `masuk` dan mengirim nilai + n ke output.
// Channel output ditutup saat channel input ditutup.
//
// Contoh:
//
//	src    := KirimBanyak(1, 2, 3)
//	added  := Tambah(src, 10)
//	for v := range added { fmt.Println(v) }
//	// → 11, 12, 13
func Tambah(masuk <-chan int, n int) <-chan int {
	// TODO: implementasi di sini
	return nil
}

// PipelineLengkap menyambung: KirimBanyak → Gandakan → Tambah menjadi satu pipeline.
// Mengembalikan channel akhir yang mengalirkan hasil.
//
// Contoh:
//
//	out := PipelineLengkap([]int{1, 2, 3}, 10)
//	for v := range out {
//	    fmt.Println(v) // → 12, 14, 16  (1*2+10, 2*2+10, 3*2+10)
//	}
//
// Hint:
//   - Tidak perlu goroutine baru. Cukup sambungkan fungsi yang sudah ada.
//   - tahap1 := KirimBanyak(data...)
//   - tahap2 := Gandakan(tahap1)
//   - tahap3 := Tambah(tahap2, tambahan)
//   - return tahap3
func PipelineLengkap(data []int, tambahan int) <-chan int {
	// TODO: implementasi di sini
	return nil
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 5: FAN-OUT DAN FAN-IN SEDERHANA
// ═══════════════════════════════════════════════════════════════════════════════
//
// FAN-OUT  = satu channel → dibagi ke N goroutine
// FAN-IN   = N channel → digabung ke satu channel
//
//   [Sumber] → [worker1] ──┐
//            ↘ [worker2] ──┤──> [gabungan] → hasil
//            ↘ [worker3] ──┘
//

// GabungDua menggabungkan dua channel menjadi satu.
// Nilai dari ch1 DAN ch2 diteruskan ke channel output.
// Output ditutup saat KEDUA channel input ditutup.
//
// Contoh:
//
//	ch1 := KirimBanyak(1, 3, 5)
//	ch2 := KirimBanyak(2, 4, 6)
//	gabungan := GabungDua(ch1, ch2)
//	var hasil []int
//	for v := range gabungan { hasil = append(hasil, v) }
//	sort.Ints(hasil)
//	// hasil = [1, 2, 3, 4, 5, 6]
//
// Hint:
//   - Buat channel output.
//   - Buat WaitGroup dengan Add(2).
//   - Untuk setiap channel input, jalankan goroutine yang drain isinya ke output.
//   - Setiap goroutine Done() saat channel inputnya habis.
//   - Goroutine terpisah: tunggu WaitGroup lalu close output.
func GabungDua(ch1, ch2 <-chan int) <-chan int {
	// TODO: implementasi di sini
	return nil
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 6: SELECT DASAR
// ═══════════════════════════════════════════════════════════════════════════════
//
// SELECT seperti switch tapi untuk channel.
// Select menunggu dan mengeksekusi CASE PERTAMA yang siap.
//
//   select {
//   case v := <-ch1:   // jika ch1 punya nilai → ambil
//   case v := <-ch2:   // jika ch2 punya nilai → ambil
//   default:           // jika tidak ada yang siap → langsung ini (non-blocking)
//   }
//
// ⚠️  PENTING: Select memilih ACAK jika lebih dari satu case siap bersamaan.
//

// AmbilTercepat menjalankan fn1 dan fn2 secara paralel di goroutine masing-masing.
// Mengembalikan hasil dari yang SELESAI LEBIH DULU.
// Hasil satunya diabaikan.
//
// Contoh:
//
//	hasil := AmbilTercepat(
//	    func() int { time.Sleep(50*time.Millisecond); return 1 },
//	    func() int { time.Sleep(10*time.Millisecond); return 2 },
//	)
//	// hasil == 2  (fn2 selesai duluan)
//
// Hint:
//   - Buat dua buffered channel (kapasitas 1).
//   - Jalankan fn1 di goroutine → kirim ke ch1.
//   - Jalankan fn2 di goroutine → kirim ke ch2.
//   - Gunakan select: case v := <-ch1 → return v, case v := <-ch2 → return v.
func AmbilTercepat(fn1, fn2 func() int) int {
	// TODO: implementasi di sini
	return 0
}

// CekAtauDefault memeriksa apakah channel `ch` punya nilai yang bisa langsung diambil.
// Jika ya → return (nilai, true).
// Jika tidak → return (0, false) TANPA BLOKIR.
//
// Contoh:
//
//	ch := make(chan int, 1)
//	ch <- 99
//
//	v, ada := CekAtauDefault(ch)
//	// v=99, ada=true
//
//	v, ada = CekAtauDefault(ch) // channel kosong sekarang
//	// v=0, ada=false
//
// Hint:
//   - Gunakan select dengan default:
//     select {
//     case v := <-ch: return v, true
//     default: return 0, false
//     }
func CekAtauDefault(ch <-chan int) (int, bool) {
	// TODO: implementasi di sini
	return 0, false
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 7: DONE CHANNEL — CARA STOP GOROUTINE
// ═══════════════════════════════════════════════════════════════════════════════
//
// Masalah: bagaimana menghentikan goroutine yang sedang berjalan?
//
// Solusi: "done channel" — channel khusus sebagai sinyal berhenti.
//
//   done := make(chan struct{})
//
//   go func() {
//       for {
//           select {
//           case <-done:       // ada sinyal berhenti
//               return         // goroutine berhenti
//           case v := <-data: // ada data baru
//               proses(v)
//           }
//       }
//   }()
//
//   close(done)  // kirim sinyal berhenti ke goroutine
//
// Mengapa close(done) dan bukan done <- struct{}{}?
//   → close() membangunkan SEMUA goroutine yang menunggu channel tersebut.
//   → kirim nilai hanya membangunkan SATU goroutine.
//

// GeneratorDenganDone menghasilkan angka naik (1, 2, 3, ...) secara terus-menerus
// ke channel output. Berhenti saat channel `done` ditutup.
// Channel output ditutup saat berhenti.
//
// Contoh:
//
//	done := make(chan struct{})
//	ch := GeneratorDenganDone(done)
//
//	fmt.Println(<-ch) // → 1
//	fmt.Println(<-ch) // → 2
//	fmt.Println(<-ch) // → 3
//	close(done)       // stop!
//	// ch akan ditutup, range ch akan selesai
//
// Hint:
//   - Buat channel output.
//   - Jalankan goroutine dengan counter mulai dari 1.
//   - Loop: select antara kirim counter ke out, ATAU done ditutup.
//   - Saat done ditutup: close(out) dan return.
func GeneratorDenganDone(done <-chan struct{}) <-chan int {
	// TODO: implementasi di sini
	return nil
}

// AmbilN mengambil tepat `n` nilai pertama dari channel `src`,
// lalu mengembalikan slice berisi n nilai tersebut.
// Berguna dikombinasikan dengan generator tak terbatas.
//
// Contoh:
//
//	done := make(chan struct{})
//	defer close(done)
//	gen := GeneratorDenganDone(done)
//	hasil := AmbilN(gen, 5)
//	// hasil = [1, 2, 3, 4, 5]
//
// Hint:
//   - Buat slice kapasitas n.
//   - Loop n kali: ambil dari src, append ke slice.
//   - Return slice.
func AmbilN(src <-chan int, n int) []int {
	// TODO: implementasi di sini
	return nil
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BONUS: TANTANGAN — ORDERED WORKER POOL
// ═══════════════════════════════════════════════════════════════════════════════
//
// Ini gabungan semua yang sudah dipelajari hari ini:
// Worker pool TAPI hasilnya tetap urut sesuai input.
//
// Bayangkan kamu punya 1000 URL yang mau di-fetch, pakai 4 worker paralel,
// dan kamu mau hasilnya urut sesuai daftar URL awal.
//

// ProsesUrut memproses setiap elemen `input` dengan fungsi `fn` menggunakan
// `jumlahWorker` goroutine secara paralel, dan mengembalikan hasil dengan
// URUTAN YANG SAMA seperti input.
//
// Contoh:
//
//	hasil := ProsesUrut([]int{5, 3, 1, 4, 2}, func(n int) int {
//	    time.Sleep(time.Duration(n) * time.Millisecond) // waktu bervariasi
//	    return n * 10
//	}, 3)
//	// hasil = [50, 30, 10, 40, 20] ← urutan PASTI sama dengan input
//
// Hint:
//   - Gunakan teknik dari TransformasiUrut: satu channel per elemen.
//   - Tapi kali ini, buat `jumlahWorker` goroutine worker.
//   - Buat jobs channel berisi struct{idx int; val int}.
//   - Tiap worker ambil job, proses fn(job.val), kirim ke channels[job.idx].
//   - Baca hasil berurutan dari channels[0], channels[1], dst.
func ProsesUrut(input []int, fn func(int) int, jumlahWorker int) []int {
	// TODO: implementasi di sini (bonus — tidak ada test untuk ini)
	return nil
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🔧 HELPER — sudah diimplementasikan, jangan diubah
// ═══════════════════════════════════════════════════════════════════════════════

// _ mencegah error "imported and not used"
var (
	_ sync.WaitGroup
	_ = time.Second
)
