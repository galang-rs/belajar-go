package belajar

// ==================== DAY 21: CONCURRENCY (GOROUTINE & CHANNEL) ====================
// Topik: Goroutine, Channel, dan pola-pola concurrency dasar di Go.
// Konsep: goroutine adalah "lightweight thread", channel adalah cara komunikasi antar goroutine.

import (
	"sync"
)

// FanIn menggabungkan beberapa channel int menjadi satu channel output.
// Semua nilai dari setiap channel input akan dikirim ke channel output.
// Channel output ditutup setelah SEMUA channel input sudah ditutup.
// Urutan nilai di output tidak harus sama dengan urutan input (karena concurrent).
// Contoh:
//
//	ch1 := make(chan int) // mengirim: 1, 2
//	ch2 := make(chan int) // mengirim: 3, 4
//	out := FanIn(ch1, ch2)
//	// out akan menerima 1, 2, 3, 4 (urutan bisa berbeda)
//
// Hint: jalankan satu goroutine per channel input, gunakan sync.WaitGroup
// untuk menunggu semua selesai, lalu tutup channel output.
func FanIn(channels ...<-chan int) <-chan int {
	// TODO: implementasi di sini
	out := make(chan int)
	var wg sync.WaitGroup

	// Loop semua channel input
	for _, ch := range channels {
		wg.Add(1)

		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				out <- val // kirim ke output
			}
		}(ch)
	}

	// Tutup channel setelah semua selesai
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// Pipeline membuat pipeline 3 tahap:
//  1. Generator: menghasilkan angka dari slice input ke channel.
//  2. Processor: membaca dari channel, mengalikan setiap angka dengan multiplier, kirim ke channel baru.
//  3. Collector: membaca semua hasil dari channel processor, kumpulkan ke slice.
//
// Fungsi ini mengembalikan slice hasil akhir.
// Urutan output harus sama dengan urutan input (karena pipeline sekuensial per tahap).
// Contoh:
//
//	Pipeline([]int{1, 2, 3}, 10) -> []int{10, 20, 30}
//	Pipeline([]int{5}, 3)        -> []int{15}
//	Pipeline([]int{}, 5)         -> []int{}
//
// Hint: buat 3 fungsi/goroutine terpisah, hubungkan dengan channel.
func Pipeline(input []int, multiplier int) []int {
	// TODO: implementasi di sini
	var wg sync.WaitGroup
	out := make([]int, len(input))
	for k := range input {
		wg.Add(1)
		go func(c int) {
			defer wg.Done()
			out[c] = input[c] * multiplier
		}(k)
	}
	wg.Wait()

	return out
}

// WorkerPool memproses jobs menggunakan sejumlah worker goroutine.
// Setiap job adalah sebuah int, dan hasilnya adalah job * job (kuadrat).
// Fungsi ini mengembalikan map[int]int di mana key = job, value = hasil.
// Contoh:
//
//	WorkerPool([]int{2, 3, 4}, 2) -> map[int]int{2: 4, 3: 9, 4: 16}
//	WorkerPool([]int{5}, 1)       -> map[int]int{5: 25}
//	WorkerPool([]int{}, 3)        -> map[int]int{}
//
// Parameter:
//   - jobs: slice berisi angka-angka yang akan diproses
//   - numWorkers: jumlah goroutine worker yang bekerja secara paralel
//
// Hint: buat channel jobs dan channel results.
// Jalankan numWorkers goroutine yang membaca dari jobs channel.
// Kumpulkan hasil dari results channel.
func WorkerPool(jobs []int, numWorkers int) map[int]int {
	jobCh := make(chan int)
	resultCh := make(chan [2]int) // [job, hasil]

	var wg sync.WaitGroup

	// 🔹 Worker
	worker := func() {
		defer wg.Done()
		for job := range jobCh {
			resultCh <- [2]int{job, job * job}
		}
	}

	// 🔹 Start workers
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker()
	}

	// 🔹 Kirim jobs
	go func() {
		for _, job := range jobs {
			jobCh <- job
		}
		close(jobCh)
	}()

	// 🔹 Tutup resultCh setelah worker selesai
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// 🔹 Kumpulkan hasil
	result := make(map[int]int)
	for res := range resultCh {
		job := res[0]
		val := res[1]
		result[job] = val
	}

	return result
}

// SafeCounter adalah counter yang aman digunakan oleh banyak goroutine secara bersamaan.
// Gunakan sync.Mutex untuk melindungi akses ke data.
type SafeCounter struct {
	// TODO: tambahkan field di sini
}

// NewSafeCounter membuat SafeCounter baru dengan nilai awal 0.
func NewSafeCounter() *SafeCounter {
	// TODO: implementasi di sini
	return nil
}

// Increment menambah counter sebanyak 1. Harus thread-safe.
func (sc *SafeCounter) Increment() {
	// TODO: implementasi di sini
}

// Value mengembalikan nilai counter saat ini. Harus thread-safe.
func (sc *SafeCounter) Value() int {
	// TODO: implementasi di sini
	return 0
}
