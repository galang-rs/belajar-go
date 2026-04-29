package belajar

import (
	"context"
	"errors"
	"sort"
	"sync"
	"testing"
	"time"
)

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 1: WITHCANCEL — PEMBATALAN MANUAL
// ═══════════════════════════════════════════════════════════════════════════════

func TestBuatProsesLatar_Berjalan(t *testing.T) {
	jalan := make(chan struct{})
	batal := BuatProsesLatar(func(ctx context.Context) {
		close(jalan)
		<-ctx.Done()
	})
	defer batal()

	select {
	case <-jalan:
		t.Log("✅ BuatProsesLatar: goroutine berjalan")
	case <-time.After(500 * time.Millisecond):
		t.Fatal("❌ goroutine tidak pernah berjalan")
	}
}

func TestBuatProsesLatar_Batal(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	done := make(chan error)

	go func() {
		done <- TungguSampaiBatal(ctx)
	}()

	time.Sleep(100 * time.Millisecond)
	cancel()

	select {
	case err := <-done:
		if err == nil {
			t.Fatal("❌ expected error, got nil")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("❌ function tidak return")
	}

}

func TestTungguSampaiBatal(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	err := TungguSampaiBatal(ctx)
	if err == nil {
		t.Fatal("❌ TungguSampaiBatal harus mengembalikan error non-nil")
	}
	if !errors.Is(err, context.Canceled) {
		t.Errorf("❌ err = %v, harusnya context.Canceled", err)
	}
	t.Logf("✅ TungguSampaiBatal: %v", err)
}

func TestJenisKesalahan(t *testing.T) {
	tests := []struct {
		err   error
		harap string
	}{
		{context.Canceled, "dibatalkan"},
		{context.DeadlineExceeded, "kedaluwarsa"},
		{errors.New("error lain"), "tidak diketahui"},
	}
	for _, tc := range tests {
		hasil := JenisKesalahan(tc.err)
		if hasil != tc.harap {
			t.Errorf("❌ JenisKesalahan(%v) = %q, harusnya %q", tc.err, hasil, tc.harap)
		}
	}
	t.Log("✅ JenisKesalahan: semua kasus benar")
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 2: WITHTIMEOUT — BATAS WAKTU OTOMATIS
// ═══════════════════════════════════════════════════════════════════════════════

func TestJalankanDenganTimeout_Selesai(t *testing.T) {
	err := JalankanDenganTimeout(200*time.Millisecond, func(ctx context.Context) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(50 * time.Millisecond):
			return nil
		}
	})
	if err != nil {
		t.Errorf("❌ harusnya nil, dapat %v", err)
	}
	t.Log("✅ JalankanDenganTimeout: selesai tepat waktu")
}

func TestJalankanDenganTimeout_Habis(t *testing.T) {
	err := JalankanDenganTimeout(50*time.Millisecond, func(ctx context.Context) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(200 * time.Millisecond):
			return nil
		}
	})
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("❌ harusnya DeadlineExceeded, dapat %v", err)
	}
	t.Log("✅ JalankanDenganTimeout: timeout terdeteksi")
}

func TestJalankanDenganTimeoutParent_ParentBatal(t *testing.T) {
	parent, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(30 * time.Millisecond)
		cancel() // batalkan parent lebih awal
	}()

	err := JalankanDenganTimeoutParent(parent, 500*time.Millisecond, func(ctx context.Context) error {
		<-ctx.Done()
		return ctx.Err()
	})
	if !errors.Is(err, context.Canceled) {
		t.Errorf("❌ harusnya Canceled (parent dibatal), dapat %v", err)
	}
	t.Log("✅ JalankanDenganTimeoutParent: mengikuti cancellation parent")
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 3: WITHDEADLINE — BATAS WAKTU ABSOLUT
// ═══════════════════════════════════════════════════════════════════════════════

func TestBuatContextDeadline_TidakNil(t *testing.T) {
	kapan := time.Now().Add(100 * time.Millisecond)
	ctx, cancel := BuatContextDeadline(kapan)
	defer cancel()

	if ctx == nil {
		t.Fatal("❌ context tidak boleh nil")
	}
	dl, ada := ctx.Deadline()
	if !ada {
		t.Fatal("❌ context harus punya deadline")
	}
	if !dl.Equal(kapan) {
		t.Errorf("❌ deadline = %v, harusnya %v", dl, kapan)
	}
	t.Log("✅ BuatContextDeadline: deadline tersimpan")
}

func TestBuatContextDeadline_KadaluarsaOtomatis(t *testing.T) {
	kapan := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := BuatContextDeadline(kapan)
	defer cancel()

	<-ctx.Done()
	if !errors.Is(ctx.Err(), context.DeadlineExceeded) {
		t.Errorf("❌ ctx.Err() = %v, harusnya DeadlineExceeded", ctx.Err())
	}
	t.Log("✅ BuatContextDeadline: kedaluarsa otomatis")
}

func TestSisaWaktu_AdaDeadline(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	sisa, ada := SisaWaktu(ctx)
	if !ada {
		t.Fatal("❌ SisaWaktu harus ada = true")
	}
	if sisa <= 0 || sisa > 200*time.Millisecond {
		t.Errorf("❌ SisaWaktu = %v, harusnya antara 0 dan 200ms", sisa)
	}
	t.Logf("✅ SisaWaktu: %v", sisa)
}

func TestSisaWaktu_TanpaDeadline(t *testing.T) {
	ctx := context.Background()
	sisa, ada := SisaWaktu(ctx)
	if ada {
		t.Errorf("❌ Background context tidak punya deadline, ada = %v sisa = %v", ada, sisa)
	}
	if sisa != 0 {
		t.Errorf("❌ sisa harusnya 0, dapat %v", sisa)
	}
	t.Log("✅ SisaWaktu: tanpa deadline = false")
}

func TestSisaWaktu_SudahLewat(t *testing.T) {
	kapan := time.Now().Add(-1 * time.Second) // deadline sudah lewat
	ctx, cancel := context.WithDeadline(context.Background(), kapan)
	defer cancel()

	sisa, ada := SisaWaktu(ctx)
	if !ada {
		t.Fatal("❌ harus ada = true meskipun sudah lewat")
	}
	if sisa != 0 {
		t.Errorf("❌ sisa harusnya 0 (tidak negatif), dapat %v", sisa)
	}
	t.Log("✅ SisaWaktu: sudah lewat → 0")
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 4: WITHVALUE — NILAI DALAM CONTEXT
// ═══════════════════════════════════════════════════════════════════════════════

func TestSimpanAmbilNilai(t *testing.T) {
	ctx := SimpanNilai(context.Background(), KunciUserID, "user-42")
	val, ok := AmbilNilai(ctx, KunciUserID)
	if !ok || val != "user-42" {
		t.Errorf("❌ AmbilNilai = %q %v, harusnya user-42 true", val, ok)
	}
	t.Logf("✅ SimpanNilai/AmbilNilai: %q", val)
}

func TestAmbilNilai_TidakAda(t *testing.T) {
	ctx := context.Background()
	val, ok := AmbilNilai(ctx, KunciUserID)
	if ok || val != "" {
		t.Errorf("❌ harusnya \"\", false — dapat %q %v", val, ok)
	}
	t.Log("✅ AmbilNilai key tidak ada: false")
}

func TestBuatContextRequest(t *testing.T) {
	ctx := BuatContextRequest("user-99", "trace-xyz", "admin")

	uid, ok1 := AmbilNilai(ctx, KunciUserID)
	tid, ok2 := AmbilNilai(ctx, KunciTraceID)
	per, ok3 := AmbilNilai(ctx, KunciPeran)

	if !ok1 || uid != "user-99" {
		t.Errorf("❌ UserID = %q %v", uid, ok1)
	}
	if !ok2 || tid != "trace-xyz" {
		t.Errorf("❌ TraceID = %q %v", tid, ok2)
	}
	if !ok3 || per != "admin" {
		t.Errorf("❌ Peran = %q %v", per, ok3)
	}
	t.Logf("✅ BuatContextRequest: uid=%s tid=%s peran=%s", uid, tid, per)
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 5: PROPAGASI — HIERARKI CONTEXT
// ═══════════════════════════════════════════════════════════════════════════════

func TestHitungAnakYangBatal(t *testing.T) {
	n := 5
	batal, ch := HitungAnakYangBatal(n)

	batal() // batalkan parent → semua child ikut batal

	hasil := make([]int, 0, n)
	timeout := time.After(500 * time.Millisecond)
	for i := 0; i < n; i++ {
		select {
		case idx := <-ch:
			hasil = append(hasil, idx)
		case <-timeout:
			t.Fatalf("❌ timeout: hanya %d dari %d child yang batal", len(hasil), n)
		}
	}

	if len(hasil) != n {
		t.Errorf("❌ jumlah batal = %d, harusnya %d", len(hasil), n)
	}
	sort.Ints(hasil)
	for i, v := range hasil {
		if v != i {
			t.Errorf("❌ indeks[%d] = %d, harusnya %d", i, v, i)
		}
	}
	t.Logf("✅ HitungAnakYangBatal: %d child batal semua saat parent dibatal", n)
}

func TestPropagasiNilai(t *testing.T) {
	dariAnak, dariParent := PropagasiNilai()

	if dariAnak != "user-99" {
		t.Errorf("❌ child tidak bisa baca nilai parent: dapat %q", dariAnak)
	}
	if dariParent != "" {
		t.Errorf("❌ parent seharusnya tidak bisa baca nilai child: dapat %q", dariParent)
	}
	t.Logf("✅ PropagasiNilai: child→parent=%q parent→child=%q", dariAnak, dariParent)
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 6: WORKER POOL DENGAN CONTEXT
// ═══════════════════════════════════════════════════════════════════════════════

func TestWorker_Jalankan(t *testing.T) {
	ctx := context.Background()
	masukan := make(chan int, 3)
	keluaran := make(chan int, 3)
	masukan <- 2
	masukan <- 4
	masukan <- 6
	close(masukan)

	w := &Worker{ID: 1}
	w.Jalankan(ctx, masukan, keluaran, func(n int) int { return n * 10 })

	close(keluaran)
	hasil := []int{}
	for v := range keluaran {
		hasil = append(hasil, v)
	}
	sort.Ints(hasil)

	harap := []int{20, 40, 60}
	for i, v := range harap {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ Worker.Jalankan: %v", hasil)
}

func TestWorker_BerhentiSaatCtxBatal(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	masukan := make(chan int) // unbuffered, tidak ada pengirim
	keluaran := make(chan int, 1)

	selesai := make(chan struct{})
	w := &Worker{ID: 0}
	go func() {
		w.Jalankan(ctx, masukan, keluaran, func(n int) int { return n })
		close(selesai)
	}()

	cancel()
	select {
	case <-selesai:
		t.Log("✅ Worker berhenti saat ctx dibatal")
	case <-time.After(300 * time.Millisecond):
		t.Fatal("❌ Worker tidak berhenti saat ctx dibatal")
	}
}

func TestPoolWorker(t *testing.T) {
	ctx := context.Background()
	masukan := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		masukan <- i
	}
	close(masukan)

	hasil := PoolWorker(ctx, 3, masukan, func(n int) int { return n * n })

	got := []int{}
	for v := range hasil {
		got = append(got, v)
	}
	sort.Ints(got)

	harap := []int{1, 4, 9, 16, 25}
	if len(got) != len(harap) {
		t.Fatalf("❌ panjang %d, harusnya %d", len(got), len(harap))
	}
	for i, v := range harap {
		if got[i] != v {
			t.Errorf("❌ got[%d] = %d, harusnya %d", i, got[i], v)
		}
	}
	t.Logf("✅ PoolWorker: %v", got)
}

func TestPoolWorker_BerhentiSaatBatal(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	masukan := make(chan int) // tidak ada data → worker idle

	hasil := PoolWorker(ctx, 3, masukan, func(n int) int { return n })

	cancel()
	// channel hasil harus tertutup dalam waktu singkat
	timeout := time.After(300 * time.Millisecond)
	for {
		select {
		case _, ok := <-hasil:
			if !ok {
				t.Log("✅ PoolWorker: channel ditutup saat ctx dibatal")
				return
			}
		case <-timeout:
			t.Fatal("❌ PoolWorker tidak menutup channel setelah ctx dibatal")
		}
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 7: PIPELINE DENGAN CONTEXT
// ═══════════════════════════════════════════════════════════════════════════════

func TestBangkitkan(t *testing.T) {
	ctx := context.Background()
	ch := Bangkitkan(ctx, 10, 20, 30)

	hasil := []int{}
	for v := range ch {
		hasil = append(hasil, v)
	}
	if len(hasil) != 3 || hasil[0] != 10 || hasil[1] != 20 || hasil[2] != 30 {
		t.Errorf("❌ Bangkitkan = %v, harusnya [10 20 30]", hasil)
	}
	t.Logf("✅ Bangkitkan: %v", hasil)
}

func TestBangkitkan_BatalDiTengah(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ch := Bangkitkan(ctx, 1, 2, 3, 4, 5)

	<-ch // baca satu
	cancel()

	// drain channel tanpa block selamanya
	timeout := time.After(300 * time.Millisecond)
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				t.Log("✅ Bangkitkan: channel ditutup saat ctx dibatal")
				return
			}
		case <-timeout:
			t.Fatal("❌ Bangkitkan tidak menutup channel setelah ctx dibatal")
		}
	}
}

func TestTransformasi(t *testing.T) {
	ctx := context.Background()
	gen := Bangkitkan(ctx, 1, 2, 3, 4, 5)
	kuadrat := Transformasi(ctx, gen, func(n int) int { return n * n })

	hasil := []int{}
	for v := range kuadrat {
		hasil = append(hasil, v)
	}
	harap := []int{1, 4, 9, 16, 25}
	if len(hasil) != len(harap) {
		t.Fatalf("❌ panjang %d, harusnya %d", len(hasil), len(harap))
	}
	for i, v := range harap {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ Transformasi: %v", hasil)
}

func TestTransformasi_DuaTahap(t *testing.T) {
	ctx := context.Background()
	gen := Bangkitkan(ctx, 1, 2, 3)
	tambah := Transformasi(ctx, gen, func(n int) int { return n + 10 })
	kuadrat := Transformasi(ctx, tambah, func(n int) int { return n * n })

	hasil := []int{}
	for v := range kuadrat {
		hasil = append(hasil, v)
	}
	harap := []int{121, 144, 169}
	for i, v := range harap {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ Transformasi dua tahap: %v", hasil)
}

func TestGabungChannel(t *testing.T) {
	ctx := context.Background()
	ch1 := Bangkitkan(ctx, 1, 2)
	ch2 := Bangkitkan(ctx, 3, 4)
	ch3 := Bangkitkan(ctx, 5, 6)

	gabung := GabungChannel(ctx, ch1, ch2, ch3)

	hasil := []int{}
	for v := range gabung {
		hasil = append(hasil, v)
	}
	sort.Ints(hasil)

	if len(hasil) != 6 {
		t.Fatalf("❌ panjang %d, harusnya 6", len(hasil))
	}
	for i, v := range []int{1, 2, 3, 4, 5, 6} {
		if hasil[i] != v {
			t.Errorf("❌ hasil[%d] = %d, harusnya %d", i, hasil[i], v)
		}
	}
	t.Logf("✅ GabungChannel: %v", hasil)
}

// ═══════════════════════════════════════════════════════════════════════════════
// 🧪 TEST BAGIAN 8: RETRY DENGAN CONTEXT
// ═══════════════════════════════════════════════════════════════════════════════

func TestCobaUlang_Berhasil(t *testing.T) {
	ctx := context.Background()
	n := 0
	err := CobaUlang(ctx, 3, 5*time.Millisecond, func() error {
		n++
		if n < 3 {
			return errors.New("belum berhasil")
		}
		return nil
	})
	if err != nil {
		t.Errorf("❌ harusnya nil, dapat %v", err)
	}
	if n != 3 {
		t.Errorf("❌ percobaan = %d, harusnya 3", n)
	}
	t.Logf("✅ CobaUlang: berhasil di percobaan ke-%d", n)
}

func TestCobaUlang_SemuaGagal(t *testing.T) {
	ctx := context.Background()
	err := CobaUlang(ctx, 3, 5*time.Millisecond, func() error {
		return errors.New("selalu gagal")
	})
	if !errors.Is(err, ErrMaksPercobaan) {
		t.Errorf("❌ harusnya ErrMaksPercobaan, dapat %v", err)
	}
	t.Log("✅ CobaUlang: ErrMaksPercobaan saat semua gagal")
}

func TestCobaUlang_CtxBatal(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var mu sync.Mutex
	n := 0

	go func() {
		time.Sleep(20 * time.Millisecond)
		cancel()
	}()

	err := CobaUlang(ctx, 100, 15*time.Millisecond, func() error {
		mu.Lock()
		n++
		mu.Unlock()
		return errors.New("gagal")
	})

	if !errors.Is(err, context.Canceled) {
		t.Errorf("❌ harusnya context.Canceled, dapat %v", err)
	}
	if n >= 100 {
		t.Errorf("❌ context tidak menghentikan retry (n=%d)", n)
	}
	t.Logf("✅ CobaUlang: berhenti saat ctx batal (percobaan=%d)", n)
}

func TestCobaUlang_LangsungBerhasil(t *testing.T) {
	ctx := context.Background()
	err := CobaUlang(ctx, 5, 5*time.Millisecond, func() error {
		return nil // langsung berhasil
	})
	if err != nil {
		t.Errorf("❌ harusnya nil, dapat %v", err)
	}
	t.Log("✅ CobaUlang: langsung berhasil di percobaan pertama")
}

func TestCobaUlangDenganBackoff_Berhasil(t *testing.T) {
	ctx := context.Background()
	n := 0
	err := CobaUlangDenganBackoff(ctx, 4, 5*time.Millisecond, 20*time.Millisecond, func() error {
		n++
		if n < 3 {
			return errors.New("gagal")
		}
		return nil
	})
	if err != nil {
		t.Errorf("❌ harusnya nil, dapat %v", err)
	}
	t.Logf("✅ CobaUlangDenganBackoff: berhasil di percobaan ke-%d", n)
}

func TestCobaUlangDenganBackoff_BatasJeda(t *testing.T) {
	ctx := context.Background()
	mulai := time.Now()
	// 3 percobaan, jeda 5ms, maks 10ms → total jeda ~15ms (5+10), bukan 5+10+20
	CobaUlangDenganBackoff(ctx, 3, 5*time.Millisecond, 10*time.Millisecond, func() error {
		return errors.New("gagal")
	})
	durasi := time.Since(mulai)
	// total jeda: 5ms + 10ms = 15ms (percobaan 3 tidak ada jeda)
	if durasi > 100*time.Millisecond {
		t.Errorf("❌ durasi terlalu lama: %v (harusnya <100ms)", durasi)
	}
	t.Logf("✅ CobaUlangDenganBackoff: durasi = %v", durasi)
}
