package belajar

import (
	"sort"
	"sync"
	"testing"
	"time"
)

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 1: COUNTER DENGAN MUTEX
// ═══════════════════════════════════════════════════════════════════════════════

func TestCounter_Tambah(t *testing.T) {
	c := &Counter{}
	c.Tambah(5)
	c.Tambah(3)
	if c.Nilai() != 8 {
		t.Errorf("❌ Nilai() = %d, harusnya 8", c.Nilai())
	}
	t.Logf("✅ Counter.Tambah: %d", c.Nilai())
}

func TestCounter_Kurang(t *testing.T) {
	c := &Counter{}
	c.Tambah(10)
	c.Kurang(4)
	if c.Nilai() != 6 {
		t.Errorf("❌ Nilai() = %d, harusnya 6", c.Nilai())
	}
	t.Logf("✅ Counter.Kurang: %d", c.Nilai())
}

func TestCounter_Reset(t *testing.T) {
	c := &Counter{}
	c.Tambah(42)
	c.Reset()
	if c.Nilai() != 0 {
		t.Errorf("❌ setelah Reset Nilai() = %d, harusnya 0", c.Nilai())
	}
	t.Log("✅ Counter.Reset: 0")
}

func TestCounter_Paralel(t *testing.T) {
	c := &Counter{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Tambah(1)
		}()
	}
	wg.Wait()
	if c.Nilai() != 1000 {
		t.Errorf("❌ Nilai() = %d, harusnya 1000 (race condition?)", c.Nilai())
	}
	t.Logf("✅ Counter paralel: %d", c.Nilai())
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 2: CACHE DENGAN RWMUTEX
// ═══════════════════════════════════════════════════════════════════════════════

func TestCache_SetGet(t *testing.T) {
	c := &Cache{}
	c.Set("nama", "Gopher")
	v, ok := c.Get("nama")
	if !ok || v != "Gopher" {
		t.Errorf("❌ Get(nama) = %q %v, harusnya Gopher true", v, ok)
	}
	t.Logf("✅ Cache.Set/Get: %q", v)
}

func TestCache_GetTidakAda(t *testing.T) {
	c := &Cache{}
	v, ok := c.Get("tidak_ada")
	if ok || v != "" {
		t.Errorf("❌ Get(tidak_ada) = %q %v, harusnya \"\" false", v, ok)
	}
	t.Log("✅ Cache.Get tidak ada: false")
}

func TestCache_Hapus(t *testing.T) {
	c := &Cache{}
	c.Set("kunci", "nilai")
	c.Hapus("kunci")
	_, ok := c.Get("kunci")
	if ok {
		t.Error("❌ setelah Hapus, Get harusnya false")
	}
	t.Log("✅ Cache.Hapus berhasil")
}

func TestCache_Panjang(t *testing.T) {
	c := &Cache{}
	c.Set("a", "1")
	c.Set("b", "2")
	c.Set("c", "3")
	if c.Panjang() != 3 {
		t.Errorf("❌ Panjang() = %d, harusnya 3", c.Panjang())
	}
	c.Hapus("a")
	if c.Panjang() != 2 {
		t.Errorf("❌ setelah hapus Panjang() = %d, harusnya 2", c.Panjang())
	}
	t.Logf("✅ Cache.Panjang: %d", c.Panjang())
}

func TestCache_Paralel(t *testing.T) {
	c := &Cache{}
	var wg sync.WaitGroup
	// 50 penulis + 50 pembaca bersamaan
	for i := 0; i < 50; i++ {
		wg.Add(2)
		i := i
		go func() {
			defer wg.Done()
			key := string(rune('a' + i%26))
			c.Set(key, "val")
		}()
		go func() {
			defer wg.Done()
			key := string(rune('a' + i%26))
			c.Get(key)
		}()
	}
	wg.Wait()
	t.Log("✅ Cache paralel selesai tanpa race")
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 3: SYNC.ONCE
// ═══════════════════════════════════════════════════════════════════════════════

func TestBuatSingleton(t *testing.T) {
	s1 := BuatSingleton()
	s2 := BuatSingleton()
	if s1 == nil {
		t.Fatal("❌ BuatSingleton() tidak boleh nil")
	}
	if s1 != s2 {
		t.Error("❌ s1 != s2, harusnya pointer yang sama (singleton)")
	}
	t.Log("✅ BuatSingleton: s1 == s2")
}

func TestBuatSingleton_Paralel(t *testing.T) {
	// Reset tidak bisa dilakukan pada var paket, tapi pastikan tidak race
	var wg sync.WaitGroup
	results := make([]*Singleton, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			results[i] = BuatSingleton()
		}()
	}
	wg.Wait()
	for i, s := range results {
		if s == nil {
			t.Errorf("❌ results[%d] nil", i)
		}
		if s != results[0] {
			t.Errorf("❌ results[%d] berbeda pointer dari results[0]", i)
		}
	}
	t.Log("✅ BuatSingleton paralel: semua pointer sama")
}

func TestHitungPanggilan(t *testing.T) {
	var once sync.Once
	hitung := 0
	fn := func() { hitung++ }

	HitungPanggilan(&once, fn)
	HitungPanggilan(&once, fn)
	HitungPanggilan(&once, fn)

	if hitung != 1 {
		t.Errorf("❌ fn dipanggil %d kali, harusnya 1", hitung)
	}
	t.Logf("✅ HitungPanggilan: fn hanya dipanggil %d kali", hitung)
}

func TestHitungPanggilan_Paralel(t *testing.T) {
	var once sync.Once
	hitung := 0
	var mu sync.Mutex
	fn := func() {
		mu.Lock()
		hitung++
		mu.Unlock()
	}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			HitungPanggilan(&once, fn)
		}()
	}
	wg.Wait()

	if hitung != 1 {
		t.Errorf("❌ fn dipanggil %d kali secara paralel, harusnya 1", hitung)
	}
	t.Log("✅ HitungPanggilan paralel: fn hanya dipanggil 1 kali")
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 4: KOUNTER ATOMIK
// ═══════════════════════════════════════════════════════════════════════════════

func TestKounterAtomik_Tambah(t *testing.T) {
	k := &KounterAtomik{}
	k.Tambah(5)
	k.Tambah(3)
	if k.Nilai() != 8 {
		t.Errorf("❌ Nilai() = %d, harusnya 8", k.Nilai())
	}
	t.Logf("✅ KounterAtomik.Tambah: %d", k.Nilai())
}

func TestKounterAtomik_Reset(t *testing.T) {
	k := &KounterAtomik{}
	k.Tambah(99)
	k.Reset()
	if k.Nilai() != 0 {
		t.Errorf("❌ setelah Reset Nilai() = %d, harusnya 0", k.Nilai())
	}
	t.Log("✅ KounterAtomik.Reset: 0")
}

func TestKounterAtomik_TukarJikaSama(t *testing.T) {
	k := &KounterAtomik{}
	k.Tambah(10)

	ok := k.TukarJikaSama(10, 20)
	if !ok {
		t.Error("❌ TukarJikaSama(10,20) harusnya true")
	}
	if k.Nilai() != 20 {
		t.Errorf("❌ setelah tukar Nilai() = %d, harusnya 20", k.Nilai())
	}

	ok = k.TukarJikaSama(10, 30)
	if ok {
		t.Error("❌ TukarJikaSama(10,30) harusnya false (nilai sudah 20)")
	}
	if k.Nilai() != 20 {
		t.Errorf("❌ nilai berubah jadi %d padahal tukar gagal", k.Nilai())
	}
	t.Log("✅ KounterAtomik.TukarJikaSama OK")
}

func TestKounterAtomik_Paralel(t *testing.T) {
	k := &KounterAtomik{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			k.Tambah(1)
		}()
	}
	wg.Wait()
	if k.Nilai() != 1000 {
		t.Errorf("❌ Nilai() = %d, harusnya 1000", k.Nilai())
	}
	t.Logf("✅ KounterAtomik paralel: %d", k.Nilai())
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 5: PETA AMAN
// ═══════════════════════════════════════════════════════════════════════════════

func TestPetaAman_SimpanAmbil(t *testing.T) {
	p := &PetaAman{}
	p.Simpan("a", 1)
	p.Simpan("b", 2)

	v, ok := p.Ambil("a")
	if !ok || v != 1 {
		t.Errorf("❌ Ambil(a) = %d %v, harusnya 1 true", v, ok)
	}
	t.Logf("✅ PetaAman.Ambil: %d", v)
}

func TestPetaAman_AmbilTidakAda(t *testing.T) {
	p := &PetaAman{}
	v, ok := p.Ambil("x")
	if ok || v != 0 {
		t.Errorf("❌ Ambil tidak ada = %d %v, harusnya 0 false", v, ok)
	}
	t.Log("✅ PetaAman.Ambil tidak ada: false")
}

func TestPetaAman_Hapus(t *testing.T) {
	p := &PetaAman{}
	p.Simpan("k", 99)
	p.Hapus("k")
	_, ok := p.Ambil("k")
	if ok {
		t.Error("❌ setelah Hapus, Ambil harusnya false")
	}
	t.Log("✅ PetaAman.Hapus OK")
}

func TestPetaAman_Jumlah(t *testing.T) {
	p := &PetaAman{}
	p.Simpan("x", 1)
	p.Simpan("y", 2)
	p.Simpan("z", 3)
	if p.Jumlah() != 3 {
		t.Errorf("❌ Jumlah() = %d, harusnya 3", p.Jumlah())
	}
	p.Hapus("x")
	if p.Jumlah() != 2 {
		t.Errorf("❌ setelah hapus Jumlah() = %d, harusnya 2", p.Jumlah())
	}
	t.Logf("✅ PetaAman.Jumlah: %d", p.Jumlah())
}

func TestPetaAman_Paralel(t *testing.T) {
	p := &PetaAman{}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			key := string(rune('a' + i%26))
			p.Simpan(key, i)
			p.Ambil(key)
		}()
	}
	wg.Wait()
	t.Log("✅ PetaAman paralel selesai tanpa race")
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 6: SEMAPHORE
// ═══════════════════════════════════════════════════════════════════════════════

func TestSemaphore_BuatDanAcquireRelease(t *testing.T) {
	sem := BuatSemaphore(3)
	if sem == nil {
		t.Fatal("❌ BuatSemaphore tidak boleh nil")
	}

	sem.Acquire()
	sem.Acquire()
	if sem.SlotTerpakai() != 2 {
		t.Errorf("❌ SlotTerpakai = %d, harusnya 2", sem.SlotTerpakai())
	}
	sem.Release()
	if sem.SlotTerpakai() != 1 {
		t.Errorf("❌ setelah Release SlotTerpakai = %d, harusnya 1", sem.SlotTerpakai())
	}
	t.Logf("✅ Semaphore Acquire/Release: slot=%d", sem.SlotTerpakai())
}

func TestSemaphore_BatasiKonkurensi(t *testing.T) {
	maks := 3
	sem := BuatSemaphore(maks)
	var (
		mu      sync.Mutex
		aktif   int
		puncak  int
		wg      sync.WaitGroup
	)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem.Acquire()
			defer sem.Release()

			mu.Lock()
			aktif++
			if aktif > puncak {
				puncak = aktif
			}
			mu.Unlock()

			time.Sleep(20 * time.Millisecond)

			mu.Lock()
			aktif--
			mu.Unlock()
		}()
	}

	wg.Wait()
	if puncak > maks {
		t.Errorf("❌ puncak goroutine aktif = %d, harusnya <= %d", puncak, maks)
	}
	t.Logf("✅ Semaphore membatasi ke %d (puncak=%d)", maks, puncak)
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 7: TAMBAH PARALEL
// ═══════════════════════════════════════════════════════════════════════════════

func TestTambahParalel(t *testing.T) {
	hasil := TambahParalel(5, func(i int) int { return i * i })
	if hasil == nil {
		t.Fatal("❌ hasil tidak boleh nil")
	}
	if len(hasil) != 5 {
		t.Fatalf("❌ panjang %d, harusnya 5", len(hasil))
	}
	sort.Ints(hasil)
	harusnya := []int{0, 1, 4, 9, 16}
	for i, v := range harusnya {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ TambahParalel: %v", hasil)
}

func TestTambahParalel_BenarParalel(t *testing.T) {
	mulai := time.Now()
	// 5 fungsi masing-masing sleep 50ms — paralel seharusnya ~50ms total
	TambahParalel(5, func(i int) int {
		time.Sleep(50 * time.Millisecond)
		return i
	})
	durasi := time.Since(mulai)
	if durasi > 150*time.Millisecond {
		t.Errorf("❌ durasi %v, harusnya paralel (~50ms) bukan sequential (~250ms)", durasi)
	}
	t.Logf("✅ TambahParalel berjalan paralel: %v", durasi)
}

func TestTambahParalel_RaceCondition(t *testing.T) {
	// Jalankan dengan go test -race untuk mendeteksi race condition
	hasil := TambahParalel(100, func(i int) int { return i })
	if len(hasil) != 100 {
		t.Errorf("❌ panjang %d, harusnya 100", len(hasil))
	}
	t.Logf("✅ TambahParalel 100 goroutine: %d hasil", len(hasil))
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 8: KOLEKSI AMAN
// ═══════════════════════════════════════════════════════════════════════════════

func TestKoleksiAman_TambahAmbil(t *testing.T) {
	k := BuatKoleksi(5)
	if k == nil {
		t.Fatal("❌ BuatKoleksi tidak boleh nil")
	}

	ok := k.Tambah(10)
	if !ok {
		t.Error("❌ Tambah(10) harusnya true")
	}
	ok = k.Tambah(20)
	if !ok {
		t.Error("❌ Tambah(20) harusnya true")
	}

	v, ok := k.Ambil()
	if !ok || v != 10 {
		t.Errorf("❌ Ambil() = %d %v, harusnya 10 true (FIFO)", v, ok)
	}
	v, ok = k.Ambil()
	if !ok || v != 20 {
		t.Errorf("❌ Ambil() = %d %v, harusnya 20 true", v, ok)
	}
	t.Log("✅ KoleksiAman FIFO OK")
}

func TestKoleksiAman_Penuh(t *testing.T) {
	k := BuatKoleksi(2)
	k.Tambah(1)
	k.Tambah(2)
	ok := k.Tambah(3)
	if ok {
		t.Error("❌ Tambah saat penuh harusnya false")
	}
	if k.Ukuran() != 2 {
		t.Errorf("❌ Ukuran() = %d, harusnya 2", k.Ukuran())
	}
	t.Log("✅ KoleksiAman menolak saat penuh")
}

func TestKoleksiAman_AmbilDariKosong(t *testing.T) {
	k := BuatKoleksi(3)
	v, ok := k.Ambil()
	if ok || v != 0 {
		t.Errorf("❌ Ambil dari kosong = %d %v, harusnya 0 false", v, ok)
	}
	t.Log("✅ KoleksiAman.Ambil dari kosong: false")
}

func TestKoleksiAman_Ukuran(t *testing.T) {
	k := BuatKoleksi(5)
	if k.Ukuran() != 0 {
		t.Errorf("❌ Ukuran awal = %d, harusnya 0", k.Ukuran())
	}
	k.Tambah(1)
	k.Tambah(2)
	if k.Ukuran() != 2 {
		t.Errorf("❌ Ukuran() = %d, harusnya 2", k.Ukuran())
	}
	k.Ambil()
	if k.Ukuran() != 1 {
		t.Errorf("❌ setelah Ambil Ukuran() = %d, harusnya 1", k.Ukuran())
	}
	t.Logf("✅ KoleksiAman.Ukuran: %d", k.Ukuran())
}

func TestKoleksiAman_Snapshot(t *testing.T) {
	k := BuatKoleksi(5)
	k.Tambah(1)
	k.Tambah(2)
	k.Tambah(3)

	snap := k.Snapshot()
	if len(snap) != 3 {
		t.Fatalf("❌ Snapshot panjang %d, harusnya 3", len(snap))
	}
	// Ubah snapshot — tidak boleh mempengaruhi koleksi asli
	snap[0] = 999
	if k.Ukuran() != 3 {
		t.Error("❌ mengubah Snapshot mempengaruhi koleksi asli")
	}
	// Verifikasi isi asli lewat Ambil
	v, _ := k.Ambil()
	if v != 1 {
		t.Errorf("❌ elemen pertama = %d, harusnya 1 (snapshot bukan salinan?)", v)
	}
	t.Logf("✅ KoleksiAman.Snapshot: %v (salinan aman)", snap)
}

func TestKoleksiAman_Paralel(t *testing.T) {
	k := BuatKoleksi(200)
	var wg sync.WaitGroup

	// 100 goroutine tambah
	for i := 0; i < 100; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			k.Tambah(i)
		}()
	}

	// 50 goroutine ambil
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			k.Ambil()
		}()
	}

	wg.Wait()
	t.Logf("✅ KoleksiAman paralel selesai, ukuran akhir: %d", k.Ukuran())
}
