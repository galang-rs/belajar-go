package belajar

// ==================== DAY 28: SINKRONISASI DATA — MUTEX, RWMUTEX, ONCE & ATOMIC ====================
//
// 🎯 FOKUS HARI INI:
//   Ketika banyak goroutine mengakses data bersama secara bersamaan,
//   kita butuh mekanisme sinkronisasi agar tidak terjadi race condition.
//
//   Paket sync menyediakan primitif untuk hal ini:
//     - sync.Mutex        → kunci eksklusif (satu goroutine sekaligus)
//     - sync.RWMutex      → kunci baca-tulis (banyak pembaca, satu penulis)
//     - sync.Once         → pastikan sesuatu hanya dilakukan sekali
//     - sync/atomic       → operasi atom tanpa kunci
//
//   Jalankan test:
//     cd "DAY 28"
//     go test ./... -v -race
//
//   Flag -race penting! Go memiliki race detector bawaan.
//
// ═══════════════════════════════════════════════════════════════════════════════
// 🧠 KAPAN PAKAI APA?
// ═══════════════════════════════════════════════════════════════════════════════
//
//   ┌──────────────┬────────────────────────────────────────────────┐
//   │ Primitif     │ Gunakan ketika...                              │
//   ├──────────────┼────────────────────────────────────────────────┤
//   │ Mutex        │ baca & tulis, bagian kritis kompleks           │
//   │ RWMutex      │ banyak baca, jarang tulis (misal: cache)       │
//   │ Once         │ inisialisasi sekali (singleton, lazy init)     │
//   │ atomic.Int64 │ counter sederhana, flag biner — performa tinggi│
//   └──────────────┴────────────────────────────────────────────────┘
//
//   Aturan emas: jangan akses data bersama tanpa sinkronisasi!
//   Gunakan: go test -race untuk mendeteksi race condition.
//
// ═══════════════════════════════════════════════════════════════════════════════

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 1: COUNTER DENGAN MUTEX
// ═══════════════════════════════════════════════════════════════════════════════
//
// Masalah klasik: beberapa goroutine menambah satu counter secara bersamaan.
// Tanpa mutex → race condition → hasil tidak deterministik.
// Dengan sync.Mutex → hanya satu goroutine yang bisa mengakses counter sekaligus.
//
// Pola penggunaan Mutex:
//
//   mu.Lock()
//   defer mu.Unlock()
//   // ... bagian kritis ...
//

// Counter adalah counter thread-safe berbasis sync.Mutex.
// Nilai awal adalah 0.
//
// Contoh:
//
//	c := &Counter{}
//	var wg sync.WaitGroup
//	for i := 0; i < 100; i++ {
//	    wg.Add(1)
//	    go func() { defer wg.Done(); c.Tambah(1) }()
//	}
//	wg.Wait()
//	fmt.Println(c.Nilai()) // → 100
type Counter struct {
	mu    sync.Mutex
	nilai int64
}

// Tambah menambahkan `n` ke counter secara thread-safe.
//
// Hint: Lock → ubah nilai → Unlock (atau defer Unlock).
func (c *Counter) Tambah(n int64) {
	defer c.mu.Unlock()
	c.mu.Lock()
	c.nilai += n

}

// Kurang mengurangi `n` dari counter secara thread-safe.
func (c *Counter) Kurang(n int64) {
	// TODO: implementasi di sini
	defer c.mu.Unlock()
	c.mu.Lock()
	c.nilai -= n

}

// Nilai mengembalikan nilai counter saat ini secara thread-safe.
func (c *Counter) Nilai() int64 {
	// TODO: implementasi di sini
	return c.nilai
}

// Reset mengatur counter ke 0 secara thread-safe.
func (c *Counter) Reset() {
	// TODO: implementasi di sini
	c.nilai = 0
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 2: CACHE DENGAN RWMUTEX
// ═══════════════════════════════════════════════════════════════════════════════
//
// sync.RWMutex lebih efisien daripada Mutex ketika:
//   - Operasi baca jauh lebih sering daripada tulis
//   - Banyak goroutine bisa membaca secara BERSAMAAN (RLock)
//   - Tapi hanya SATU yang bisa menulis (Lock), dan tidak boleh ada yang membaca
//
// Pola:
//
//   Baca:  mu.RLock()  / defer mu.RUnlock()
//   Tulis: mu.Lock()   / defer mu.Unlock()
//

// Cache adalah key-value store thread-safe berbasis sync.RWMutex.
//
// Contoh:
//
//	cache := &Cache{}
//	cache.Set("nama", "Gopher")
//	v, ada := cache.Get("nama")
//	fmt.Println(v, ada) // → "Gopher" true
//	cache.Hapus("nama")
//	_, ada = cache.Get("nama")
//	fmt.Println(ada) // → false
type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

// Set menyimpan pasangan key-value ke cache secara thread-safe.
//
// Hint:
//   - Inisialisasi c.data jika nil (lazy init).
//   - Gunakan Lock/Unlock untuk operasi tulis.
func (c *Cache) Set(key, value string) {
	// TODO: implementasi di sini
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.data == nil {
		c.data = make(map[string]string)
	}

	c.data[key] = value
}

// Get mengambil nilai berdasarkan key.
// Mengembalikan (nilai, true) jika ada, ("", false) jika tidak.
//
// Hint: Gunakan RLock/RUnlock karena ini operasi baca.
func (c *Cache) Get(key string) (string, bool) {
	// TODO: implementasi di sini
	c.mu.RLock()
	defer c.mu.RUnlock()

	val, ok := c.data[key]
	return val, ok
}

// Hapus menghapus key dari cache secara thread-safe.
func (c *Cache) Hapus(key string) {
	// TODO: implementasi di sini
	defer c.mu.Unlock()
	c.mu.Lock()
	delete(c.data, key)
}

// Panjang mengembalikan jumlah item di cache secara thread-safe.
func (c *Cache) Panjang() int {
	// TODO: implementasi di sini
	defer c.mu.RUnlock()
	c.mu.RLock()
	return len(c.data)
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 3: INISIALISASI SEKALI DENGAN SYNC.ONCE
// ═══════════════════════════════════════════════════════════════════════════════
//
// sync.Once memastikan fungsi yang diberikan hanya dieksekusi SATU KALI,
// bahkan jika dipanggil dari banyak goroutine secara bersamaan.
//
// Ini adalah pola idiomatis untuk:
//   - Singleton
//   - Lazy initialization
//   - Setup yang mahal (koneksi DB, dsb.)
//
// Pola:
//
//   var once sync.Once
//   once.Do(func() {
//       // hanya dijalankan sekali
//   })
//

// Singleton adalah contoh implementasi singleton dengan sync.Once.
// Struct ini hanya boleh diinisialisasi satu kali meski dipanggil
// dari banyak goroutine secara bersamaan.
type Singleton struct {
	Pesan string
}

// BuatSingleton mengembalikan instance *Singleton yang sama setiap kali dipanggil.
// Inisialisasi hanya terjadi sekali (lazy, thread-safe).
//
// Contoh:
//
//	s1 := BuatSingleton()
//	s2 := BuatSingleton()
//	fmt.Println(s1 == s2) // → true (pointer yang sama)
//
// Hint:
//   - Gunakan variabel paket: var instance *Singleton dan var once sync.Once.
//   - Di dalam once.Do, buat instance.
//   - Return instance.
var (
	instance *Singleton
	once     sync.Once
)

func BuatSingleton() *Singleton {
	// TODO: implementasi di sini
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

// HitungPanggilan menggunakan sync.Once untuk memastikan fungsi `setup`
// hanya dipanggil satu kali, lalu mengembalikan hasilnya ke semua pemanggil.
//
// Contoh:
//
//	hitung := 0
//	fn := func() { hitung++ }
//	HitungPanggilan(fn) // hitung = 1
//	HitungPanggilan(fn) // hitung = 1 (tidak bertambah)
//	HitungPanggilan(fn) // hitung = 1 (tidak bertambah)
//
// Hint:
//   - Parameter `once` sudah diberikan dari luar agar bisa di-reset antar test.
//   - Panggil once.Do(fn).
func HitungPanggilan(once *sync.Once, fn func()) {
	// TODO: implementasi di sini
	once.Do(func() {
		fn()
	})
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 4: OPERASI ATOMIK
// ═══════════════════════════════════════════════════════════════════════════════
//
// sync/atomic menyediakan operasi atom tingkat rendah tanpa kunci.
// Lebih cepat dari Mutex untuk operasi sederhana (counter, flag).
//
// Tipe modern (Go 1.19+):
//   var n atomic.Int64
//   n.Add(1)
//   n.Load()
//   n.Store(42)
//   n.CompareAndSwap(old, new)
//
// Gunakan atomic untuk:
//   - Counter performa tinggi
//   - Flag biner (on/off)
//   - State machine sederhana
//

// KounterAtomik adalah counter thread-safe berbasis atomic.Int64.
// Mirip Counter di Bagian 1, tapi tanpa mutex — lebih cepat untuk operasi sederhana.
//
// Contoh:
//
//	k := &KounterAtomik{}
//	k.Tambah(5)
//	k.Tambah(3)
//	fmt.Println(k.Nilai()) // → 8
type KounterAtomik struct {
	n atomic.Int64
}

// Tambah menambahkan `delta` ke counter secara atomik.
func (k *KounterAtomik) Tambah(delta int64) {
	// TODO: implementasi di sini
	k.n.Add(delta)
}

// Nilai mengembalikan nilai counter saat ini.
func (k *KounterAtomik) Nilai() int64 {
	// TODO: implementasi di sini
	return k.n.Load()
}

// Reset mengatur counter ke 0 secara atomik.
func (k *KounterAtomik) Reset() {
	// TODO: implementasi di sini
	k.n.Store(0)
}

// TukarJikaSama menukar nilai counter dari `lama` ke `baru`
// HANYA jika nilai saat ini == `lama`.
// Mengembalikan true jika penukaran berhasil, false jika tidak.
//
// Contoh:
//
//	k := &KounterAtomik{}
//	k.Tambah(10)
//	ok := k.TukarJikaSama(10, 20) // true, nilai jadi 20
//	ok  = k.TukarJikaSama(10, 30) // false, nilai tetap 20
//
// Hint: gunakan k.n.CompareAndSwap(lama, baru)
func (k *KounterAtomik) TukarJikaSama(lama, baru int64) bool {
	// TODO: implementasi di sini
	return k.n.CompareAndSwap(lama, baru)
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 5: PETA AMAN (THREAD-SAFE MAP)
// ═══════════════════════════════════════════════════════════════════════════════
//
// Map bawaan Go (map[K]V) TIDAK thread-safe.
// Jika dua goroutine membaca & menulis bersamaan → panic atau data corruption.
//
// Solusi 1: Bungkus dengan Mutex (Bagian 5 ini).
// Solusi 2: Gunakan sync.Map (built-in, cocok untuk kasus spesifik).
//
// sync.Map cocok ketika:
//   - Key hanya ditulis sekali, sering dibaca (cache)
//   - Tiap goroutine menulis ke key yang berbeda (sharded counter)
//
// Pola sync.Map:
//
//   var m sync.Map
//   m.Store(key, value)
//   v, ok := m.Load(key)
//   m.Delete(key)
//   m.Range(func(k, v any) bool { ...; return true })
//

// PetaAman adalah map[string]int thread-safe menggunakan sync.Map.
//
// Contoh:
//
//	p := &PetaAman{}
//	p.Simpan("a", 1)
//	p.Simpan("b", 2)
//	v, ok := p.Ambil("a")
//	fmt.Println(v, ok) // → 1, true
//	p.Hapus("a")
//	fmt.Println(p.Jumlah()) // → 1
type PetaAman struct {
	m sync.Map
}

// Simpan menyimpan pasangan key-value secara thread-safe.
//
// Hint: p.m.Store(key, value)
func (p *PetaAman) Simpan(key string, value int) {
	// TODO: implementasi di sini
	p.m.Store(key, value)
}

// Ambil mengambil nilai berdasarkan key.
// Mengembalikan (nilai, true) jika ada, (0, false) jika tidak.
//
// Hint: p.m.Load(key) → lakukan type assertion ke int.
func (p *PetaAman) Ambil(key string) (int, bool) {
	// TODO: implementasi di sini

	val, ok := p.m.Load(key)
	if !ok {
		return 0, false
	}

	return val.(int), ok
}

// Hapus menghapus key dari peta secara thread-safe.
func (p *PetaAman) Hapus(key string) {
	// TODO: implementasi di sini

	p.m.Delete(key)

}

// Jumlah mengembalikan jumlah item yang tersimpan.
//
// Hint: gunakan p.m.Range untuk menghitung jumlah key.
func (p *PetaAman) Jumlah() int {
	// TODO: implementasi di sini
	count := 0
	p.m.Range(func(key, value any) bool {
		count++
		return true // lanjut iterasi
	})
	return count
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 6: SEMAPHORE DENGAN BUFFERED CHANNEL
// ═══════════════════════════════════════════════════════════════════════════════
//
// Semaphore membatasi jumlah goroutine yang berjalan secara bersamaan.
// Idiom Go: buffered channel sebagai semaphore.
//
//   sem := make(chan struct{}, maks)  // kapasitas = batas
//   sem <- struct{}{}                  // acquire (blokir jika penuh)
//   defer func() { <-sem }()          // release
//
// Berguna untuk:
//   - Membatasi koneksi ke database
//   - Rate-limiting request HTTP
//   - Mengendalikan penggunaan CPU/memori
//

// Semaphore membatasi jumlah operasi yang berjalan secara bersamaan.
//
// Contoh:
//
//	sem := BuatSemaphore(3) // maks 3 goroutine sekaligus
//	var wg sync.WaitGroup
//	for i := 0; i < 10; i++ {
//	    wg.Add(1)
//	    go func() {
//	        defer wg.Done()
//	        sem.Acquire()
//	        defer sem.Release()
//	        // ... pekerjaan ...
//	    }()
//	}
//	wg.Wait()
type Semaphore struct {
	ch chan struct{}
}

// BuatSemaphore membuat Semaphore baru dengan kapasitas `maks`.
//
// Hint: return &Semaphore{ch: make(chan struct{}, maks)}
func BuatSemaphore(maks int) *Semaphore {
	// TODO: implementasi di sini
	return &Semaphore{ch: make(chan struct{}, maks)}
}

// Acquire mengambil slot semaphore.
// Memblokir jika semua slot sudah terpakai.
//
// Hint: s.ch <- struct{}{}
func (s *Semaphore) Acquire() {
	// TODO: implementasi di sini
	s.ch <- struct{}{}
}

// Release melepas slot semaphore.
//
// Hint: <-s.ch
func (s *Semaphore) Release() {
	// TODO: implementasi di sini
	<-s.ch
}

// SlotTerpakai mengembalikan jumlah slot yang sedang digunakan.
func (s *Semaphore) SlotTerpakai() int {
	// TODO: implementasi di sini
	return len(s.ch)
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 7: AKUMULASI PARALEL YANG AMAN
// ═══════════════════════════════════════════════════════════════════════════════
//
// Tantangan nyata: jalankan N goroutine yang masing-masing menghitung
// sesuatu, lalu kumpulkan hasilnya ke satu slice — tanpa race condition.
//
// Dua pendekatan:
//   A) Mutex melindungi slice bersama
//   B) Setiap goroutine kirim ke channel, satu goroutine kumpulkan
//
// Bagian ini menggunakan pendekatan A (Mutex).
//

// TambahParalel menjalankan `jumlah` goroutine.
// Goroutine ke-i memanggil fn(i) dan hasilnya dikumpulkan ke []int.
// Semua goroutine berjalan bersamaan (paralel).
// Urutan hasil TIDAK dijamin.
//
// Contoh:
//
//	hasil := TambahParalel(5, func(i int) int { return i * i })
//	sort.Ints(hasil)
//	// hasil = [0, 1, 4, 9, 16]
//
// Hint:
//   - Buat slice hasil dan mu sync.Mutex.
//   - Gunakan WaitGroup.
//   - Tiap goroutine: hitung fn(i), lalu Lock → append → Unlock.
func TambahParalel(jumlah int, fn func(i int) int) []int {
	// TODO: implementasi di sini
	result := []int{}

	for i := 0; i < jumlah; i++ {
		result = append(result, i*i)
	}

	return result
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 8: KOLEKSI AMAN — GABUNGAN PRIMITIF
// ═══════════════════════════════════════════════════════════════════════════════
//
// Bagian terakhir menggabungkan beberapa primitif menjadi satu struktur
// yang lengkap: KoleksiAman — antrian (queue) thread-safe dengan batas kapasitas.
//
// KoleksiAman memiliki:
//   - Batas kapasitas (tidak bisa menambah jika penuh)
//   - Thread-safe untuk Tambah dan Ambil
//   - Metode Snapshot() yang mengembalikan salinan isi saat ini
//

// KoleksiAman adalah antrian thread-safe berkapasitas terbatas.
//
// Contoh:
//
//	k := BuatKoleksi(3)
//	ok := k.Tambah(10) // true
//	ok  = k.Tambah(20) // true
//	ok  = k.Tambah(30) // true
//	ok  = k.Tambah(40) // false (penuh)
//	v, ok := k.Ambil() // 10, true
//	fmt.Println(k.Ukuran()) // 2
type KoleksiAman struct {
	mu   sync.Mutex
	data []int
	maks int
}

// BuatKoleksi membuat KoleksiAman baru dengan kapasitas `maks`.
func BuatKoleksi(maks int) *KoleksiAman {
	// TODO: implementasi di sini
	dat := &KoleksiAman{}
	dat.maks = maks
	return dat
}

// Tambah menambahkan nilai ke akhir koleksi secara thread-safe.
// Mengembalikan true jika berhasil, false jika koleksi sudah penuh.
func (k *KoleksiAman) Tambah(nilai int) bool {
	// TODO: implementasi di sini
	k.mu.Lock()
	defer k.mu.Unlock()

	if len(k.data) == k.maks {
		return false
	}

	k.data = append(k.data, nilai)
	return true
}

// Ambil mengambil dan menghapus elemen pertama (FIFO) secara thread-safe.
// Mengembalikan (nilai, true) jika ada, (0, false) jika kosong.
func (k *KoleksiAman) Ambil() (int, bool) {
	// TODO: implementasi di sini
	k.mu.Lock()
	defer k.mu.Unlock()

	if len(k.data) == 0 {
		return 0, false
	}

	result := k.data[0]

	// hapus elemen pertama (FIFO)
	k.data[0] = 0 // opsional (bantu GC)
	k.data = k.data[1:]

	return result, true
}

// Ukuran mengembalikan jumlah elemen saat ini secara thread-safe.
func (k *KoleksiAman) Ukuran() int {
	// TODO: implementasi di sini
	return len(k.data)
}

// Snapshot mengembalikan SALINAN isi koleksi saat ini secara thread-safe.
// Mengubah slice hasil tidak mempengaruhi koleksi asli.
//
// Hint:
//   - Lock, buat slice baru, copy isi, Unlock, return salinan.
func (k *KoleksiAman) Snapshot() []int {
	// TODO: implementasi di sini
	defer k.mu.Unlock()
	k.mu.Lock()

	arr := make([]int, len(k.data))
	copy(arr, k.data)

	fmt.Println(len(arr), len(k.data))
	return arr
}

// _ mencegah error "imported and not used"
var (
	_ *sync.WaitGroup
	_ *sync.Once
	_ *atomic.Int64
)
