package belajar

// ==================== DAY 26: CHANNEL SEBAGAI RETURN VALUE ====================
//
// 🎯 FOKUS HARI INI:
//   Semua fungsi return <-chan int (atau varian channel lainnya).
//   Kamu tidak perlu manual convert channel → int → slice lagi.
//   Hasilnya bisa langsung di-compose seperti LEGO.
//
//   Jalankan test:
//     cd "DAY 26"
//     go test ./... -v
//
// ═══════════════════════════════════════════════════════════════════════════════
// 🧠 POLA BARU YANG AKAN KAMU PELAJARI
// ═══════════════════════════════════════════════════════════════════════════════
//
// Day 25 kamu sudah paham: goroutine kirim ke channel.
// Day 26 naik level: fungsi RETURN channel sehingga bisa di-chain langsung.
//
//   POLA LAMA (Day 25 style):
//   ─────────────────────────
//   ch := HitungDiBackground(fn)
//   hasil := <-ch                   // ambil dulu jadi int
//   kuadrat := hasil * hasil        // baru bisa diproses
//
//   POLA BARU (Day 26 style):
//   ─────────────────────────
//   hasil := KuadratAsync(5)        // return <-chan int langsung
//   // tidak perlu convert! langsung compose:
//   ganda  := Kali(hasil, 2)        // <-chan int → <-chan int
//   dibagi := Filter(ganda, func(n int) bool { return n > 10 })
//   // baca di ujung aja:
//   for v := range dibagi { fmt.Println(v) }
//
// Ini disebut "channel pipeline composition" — setiap fungsi menerima
// dan mengembalikan channel, seperti UNIX pipe: cmd1 | cmd2 | cmd3
//
// ═══════════════════════════════════════════════════════════════════════════════

import (
	"sync"
	"time"
)

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 1: ASYNC YANG RETURN CHANNEL
// ═══════════════════════════════════════════════════════════════════════════════
//
// Fungsi-fungsi ini menjalankan komputasi di goroutine dan langsung
// return channel — kamu bisa langsung pipe ke fungsi berikutnya.
//

// KuadratAsync menghitung n*n di goroutine terpisah.
// Mengembalikan channel yang akan berisi hasilnya.
// Fungsi ini langsung return TANPA menunggu perhitungan selesai.
//
// Contoh:
//
//	ch := KuadratAsync(7)
//	fmt.Println(<-ch) // → 49  (hasil datang dari goroutine)
//
// Hint:
//   - Buat buffered channel kapasitas 1.
//   - Jalankan goroutine: ch <- n * n, lalu close(ch).
//   - Return channel SEBELUM goroutine selesai.
func KuadratAsync(n int) <-chan int {
	// TODO: implementasi di sini

	ch := make(chan int)

	go func() {
		defer close(ch)

		ch <- n * n
	}()

	return ch

}

// JumlahAsync menghitung a+b di goroutine terpisah.
// Mengembalikan channel yang akan berisi hasilnya.
//
// Contoh:
//
//	ch := JumlahAsync(3, 4)
//	fmt.Println(<-ch) // → 7
func JumlahAsync(a, b int) <-chan int {
	// TODO: implementasi di sini

	ch := make(chan int)

	go func() {
		defer close(ch)

		ch <- a + b
	}()

	return ch
}

// FactorialAsync menghitung n! (faktorial) di goroutine terpisah.
// n! = 1 × 2 × 3 × ... × n
// Catatan: 0! = 1
//
// Contoh:
//
//	ch := FactorialAsync(5)
//	fmt.Println(<-ch) // → 120  (5! = 1×2×3×4×5)
//
// Hint:
//   - Hitung faktorial dengan loop biasa di dalam goroutine.
func FactorialAsync(n int) <-chan int {
	// TODO: implementasi di sini

	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)
	var count int

	go func() {
		if n == 0 {
			count = 1
			wg.Done()
			return
		}
		go func() {
			defer wg.Done()
			i := FactorialAsync(n - 1)
			count = n * <-i
		}()
	}()

	go func() {
		wg.Wait()
		ch <- count
		close(ch)
	}()

	return ch
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 2: TRANSFORM CHANNEL → CHANNEL
// ═══════════════════════════════════════════════════════════════════════════════
//
// Fungsi-fungsi ini menerima <-chan int dan mengembalikan <-chan int.
// Ini adalah blok penyusun pipeline yang bisa disambung tanpa batas.
//
//   src → [Filter] → [Kali] → [Ambil] → hasil
//

// Filter meneruskan nilai dari src ke output HANYA jika fn(nilai) == true.
// Nilai yang tidak lolos fn dibuang. Channel output ditutup saat src ditutup.
//
// Contoh:
//
//	src := GenRange(1, 10)                              // 1,2,3,...,10
//	genap := Filter(src, func(n int) bool { return n%2 == 0 })
//	for v := range genap { fmt.Println(v) }             // → 2,4,6,8,10
//
// Hint:
//   - Buat channel output (unbuffered).
//   - Goroutine: for v := range src → if fn(v) { out <- v }.
//   - defer close(out) di awal goroutine.
func Filter(src <-chan int, fn func(int) bool) <-chan int {
	// TODO: implementasi di sini
	ch := make(chan int)
	var wg sync.WaitGroup

	for v := range src {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if fn(v) {
				ch <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch

}

// Kali menerima setiap nilai dari src dan mengirim nilai × faktor ke output.
// Channel output ditutup saat src ditutup.
//
// Contoh:
//
//	src := GenRange(1, 5)     // 1,2,3,4,5
//	tiga := Kali(src, 3)
//	for v := range tiga { fmt.Println(v) } // → 3,6,9,12,15
func Kali(src <-chan int, faktor int) <-chan int {
	// TODO: implementasi di sini
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for v := range src {
			ch <- v * faktor
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

// Kurangi menerima setiap nilai dari src dan mengirim nilai - n ke output.
// Channel output ditutup saat src ditutup.
//
// Contoh:
//
//	src := GenRange(5, 9)   // 5,6,7,8,9
//	out := Kurangi(src, 3)
//	for v := range out { fmt.Println(v) } // → 2,3,4,5,6
func Kurangi(src <-chan int, n int) <-chan int {
	// TODO: implementasi di sini
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for v := range src {
			ch <- v - n
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

// Ambil mengambil TEPAT n nilai pertama dari src dan menutup output.
// Sisa nilai di src DIABAIKAN (tidak dibaca).
// Return channel (bukan []int!) agar bisa di-pipe terus.
//
// Contoh:
//
//	// Tanpa ini kamu harus convert ke slice dulu, dengan ini langsung pipe!
//	src := GenRange(1, 100)   // 1,2,3,...,100
//	lima := Ambil(src, 5)
//	for v := range lima { fmt.Println(v) } // → 1,2,3,4,5
//
// Hint:
//   - Buat channel output buffered kapasitas n.
//   - Goroutine: loop n kali, ambil dari src, kirim ke out, lalu close(out).
//   - Setelah close, goroutine boleh selesai (sisa src tidak perlu di-drain).
func Ambil(src <-chan int, n int) <-chan int {
	// TODO: implementasi di sini
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			rslt, ok := <-src
			if ok {
				ch <- rslt
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
// 📦 BAGIAN 3: GENERATOR — CHANNEL TANPA INPUT EKSTERNAL
// ═══════════════════════════════════════════════════════════════════════════════
//
// Generator menghasilkan nilai ke channel tanpa perlu input slice.
// Berguna sebagai sumber di awal pipeline.
//

// GenRange menghasilkan angka dari `dari` sampai `sampai` (inklusif) ke channel.
// Channel ditutup setelah semua angka dikirim.
//
// Contoh:
//
//	for v := range GenRange(3, 6) {
//	    fmt.Println(v) // → 3, 4, 5, 6
//	}
//
// Hint:
//   - Buat goroutine yang loop dari `dari` sampai `sampai` (i <= sampai).
//   - Kirim tiap i ke channel. close setelah loop.
func GenRange(dari, sampai int) <-chan int {
	// TODO: implementasi di sini
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := dari; i <= sampai; i++ {
			ch <- i
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

// GenUlangi menghasilkan nilai `v` sebanyak `n` kali ke channel.
// Channel ditutup setelah semua nilai dikirim.
//
// Contoh:
//
//	for val := range GenUlangi(7, 4) {
//	    fmt.Println(val) // → 7, 7, 7, 7
//	}
func GenUlangi(v, n int) <-chan int {
	// TODO: implementasi di sini
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			ch <- v
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 4: FAN-OUT → []<-chan int
// ═══════════════════════════════════════════════════════════════════════════════
//
// Fan-out = satu channel dibagi ke N channel.
// Setiap nilai dari src dikirim ke SATU dari N channel output (round-robin).
// Hasilnya adalah slice of channels — masing-masing bisa diproses mandiri.
//
//            ┌→ channels[0] → worker0
//   src ─────┼→ channels[1] → worker1
//            └→ channels[2] → worker2
//

// SebarKe membagi nilai dari src ke n channel output secara ROUND-ROBIN.
// channels[0] dapat nilai ke-1, channels[1] dapat nilai ke-2, dst.
// Saat src habis, semua channel output ditutup.
//
// Contoh:
//
//	src := GenRange(1, 6)        // 1,2,3,4,5,6
//	chs := SebarKe(src, 3)
//	// channels[0] ← 1, 4
//	// channels[1] ← 2, 5
//	// channels[2] ← 3, 6
//
// Hint:
//   - Buat []chan int ukuran n, masing-masing buffered.
//   - Goroutine: for v := range src → kirim ke channels[i%n] → i++.
//   - Setelah src habis, close semua channel.
//   - Return sebagai []<-chan int.
func SebarKe(src <-chan int, n int) []<-chan int {
	// TODO: implementasi di sini
	panic("belum diimplementasi")
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 5: FAN-IN — GABUNGKAN N CHANNEL
// ═══════════════════════════════════════════════════════════════════════════════
//
// Fan-in adalah kebalikan fan-out: N channel → 1 channel.
// Ini generalisasi dari GabungDua (Day 25) untuk sembarang jumlah channel.
//

// GabungSemua menggabungkan semua channel dalam `channels` menjadi satu.
// Nilai dari channel manapun diteruskan ke output.
// Output ditutup saat SEMUA channel input ditutup.
//
// Contoh:
//
//	ch1 := GenRange(1, 3)    // 1,2,3
//	ch2 := GenRange(4, 6)    // 4,5,6
//	ch3 := GenRange(7, 9)    // 7,8,9
//	gabung := GabungSemua(ch1, ch2, ch3)
//	var hasil []int
//	for v := range gabung { hasil = append(hasil, v) }
//	sort.Ints(hasil)
//	// hasil = [1,2,3,4,5,6,7,8,9]
//
// Hint:
//   - Buat WaitGroup dengan Add(len(channels)).
//   - Untuk tiap channel, goroutine drain isinya ke output, lalu Done().
//   - Goroutine terpisah: tunggu WaitGroup lalu close output.
func GabungSemua(channels ...<-chan int) <-chan int {
	// TODO: implementasi di sini
	panic("belum diimplementasi")
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 6: TIMEOUT — BERHENTI OTOMATIS
// ═══════════════════════════════════════════════════════════════════════════════
//
// Masalah umum: channel yang lambat / infinite generator.
// Solusi: bungkus dengan timeout — teruskan nilai selama belum timeout,
// lalu tutup channel output.
//

// DenganTimeout meneruskan nilai dari src ke output selama belum melebihi dur.
// Saat timeout tercapai, channel output ditutup (nilai yang belum sampai dibuang).
//
// Contoh:
//
//	done := make(chan struct{})
//	defer close(done)
//	gen := GeneratorDenganDone(done)            // 1,2,3,... (dari Day 25)
//	terbatas := DenganTimeout(gen, 50*time.Millisecond)
//	for v := range terbatas { fmt.Println(v) } // → entah berapa, tergantung kecepatan
//
// Hint:
//   - Buat output channel dan timer := time.After(dur).
//   - Goroutine dengan for-select:
//     case <-timer → close(out), return
//     case v, ok := <-src → if !ok { close(out); return }; out <- v
func DenganTimeout(src <-chan int, dur time.Duration) <-chan int {
	// TODO: implementasi di sini
	panic("belum diimplementasi")
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 7: BATCH — KUMPULKAN SEBELUM PROSES
// ═══════════════════════════════════════════════════════════════════════════════
//
// Batch berguna saat kamu ingin proses data per-kelompok, bukan satu-satu.
// Misalnya: kirim ke database tiap 100 item, bukan tiap 1 item.
//
//   src: 1,2,3,4,5,6,7  (size=3)
//   out: [1,2,3], [4,5,6], [7]   ← batch terakhir boleh kurang dari size
//

// Batch mengumpulkan nilai dari src menjadi kelompok berukuran `size`.
// Setiap kelompok dikirim sebagai []int ke channel output.
// Kelompok terakhir boleh lebih kecil dari size jika src habis.
// Channel output ditutup saat src ditutup.
//
// Contoh:
//
//	src := GenRange(1, 7)
//	batched := Batch(src, 3)
//	for grup := range batched {
//	    fmt.Println(grup) // → [1 2 3], [4 5 6], [7]
//	}
//
// Hint:
//   - Goroutine dengan loop:
//   - Kumpulkan nilai dari src ke slice sementara sampai len==size atau src tutup.
//   - Kalau slice tidak kosong, kirim ke out.
//   - Kalau src tutup dan slice kosong, close(out) dan return.
func Batch(src <-chan int, size int) <-chan []int {
	// TODO: implementasi di sini
	panic("belum diimplementasi")
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 8: ZIP — PASANGKAN DUA CHANNEL
// ═══════════════════════════════════════════════════════════════════════════════
//
// Zip menggabungkan dua channel nilai-per-nilai menjadi pasangan.
//
//   a: 1, 2, 3
//   b: 10, 20, 30
//   zip: [1,10], [2,20], [3,30]
//

// Zip membaca satu nilai dari `a` dan satu dari `b` secara bersamaan,
// lalu mengirim pasangan [2]int{valA, valB} ke output.
// Berhenti saat salah satu channel ditutup.
// Channel output ditutup setelah berhenti.
//
// Contoh:
//
//	a := GenRange(1, 3)    // 1,2,3
//	b := GenRange(10, 12)  // 10,11,12
//	for pair := range Zip(a, b) {
//	    fmt.Println(pair) // → [1 10], [2 11], [3 12]
//	}
//
// Hint:
//   - Goroutine dengan for-loop.
//   - Di tiap iterasi: baca va dari a (cek ok), baca vb dari b (cek ok).
//   - Jika salah satu ok==false, close(out) dan return.
//   - Kirim [2]int{va, vb} ke out.
func Zip(a, b <-chan int) <-chan [2]int {
	// TODO: implementasi di sini
	panic("belum diimplementasi")
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 9: COMPOSE PIPELINE LENGKAP
// ═══════════════════════════════════════════════════════════════════════════════
//
// Ini adalah "showcase" dari semua yang sudah dipelajari hari ini.
// Semua fungsi disambung menjadi satu pipeline tanpa satu pun konversi manual.
//

// HitungGenapKuadrat menghasilkan kuadrat dari bilangan genap
// dalam range [dari, sampai], diproses secara async.
// Return: <-chan int yang berisi hasil, urut dari terkecil.
//
// Langkah pipeline:
//  1. GenRange(dari, sampai)          → hasilkan semua angka
//  2. Filter(_, genap)                → hanya angka genap
//  3. Kali(_, diriSendiri)            → kuadratkan  ← MASALAH: Kali pakai faktor tetap!
//     Solusi: gunakan TransformasiChan (implementasikan di bawah) atau Kali yang kreatif
//  4. Return channel akhir
//
// Contoh:
//
//	out := HitungGenapKuadrat(1, 6)
//	for v := range out { fmt.Println(v) } // → 4, 16, 36  (2²,4²,6²)
//
// Hint:
//   - Gunakan TransformasiChan (implementasikan dulu) untuk n*n.
//   - Atau kreatif: pakai Kali dua kali tidak bisa karena src sudah habis.
//     Gunakan TransformasiChan dengan fn = func(n int) int { return n*n }.
func HitungGenapKuadrat(dari, sampai int) <-chan int {
	// TODO: implementasi di sini (butuh TransformasiChan di bawah)
	panic("belum diimplementasi")
}

// TransformasiChan menerapkan fn ke setiap nilai dari src dan mengirim hasilnya.
// Ini adalah versi "channel-native" dari TransformasiUrut (Day 25).
// PERBEDAAN dengan Day 25: di sini return <-chan int, bukan []int!
//
// Contoh:
//
//	src := GenRange(1, 5)
//	kuadrat := TransformasiChan(src, func(n int) int { return n * n })
//	for v := range kuadrat { fmt.Println(v) } // → 1, 4, 9, 16, 25
func TransformasiChan(src <-chan int, fn func(int) int) <-chan int {
	// TODO: implementasi di sini
	panic("belum diimplementasi")
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🔧 HELPER — sudah diimplementasikan, jangan diubah
// ═══════════════════════════════════════════════════════════════════════════════

// GeneratorDenganDone menghasilkan angka naik (1, 2, 3, ...) tanpa henti.
// Berhenti dan menutup channel output saat `done` ditutup.
// Dipakai oleh TestDenganTimeout — JANGAN diubah.
func GeneratorDenganDone(done <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		i := 1
		for {
			select {
			case <-done:
				return
			case ch <- i:
				i++
			}
		}
	}()
	return ch
}

// _ mencegah error "imported and not used"
var (
	_ sync.WaitGroup
	_ = time.Second
)
