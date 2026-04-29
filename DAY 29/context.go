package belajar

// ==================== DAY 29: CONTEXT — PEMBATALAN, TIMEOUT & NILAI ====================
//
// 🎯 FOKUS HARI INI:
//   Package `context` adalah cara idiomatis Go untuk:
//     - Membatalkan operasi yang berjalan (cancellation)
//     - Membatasi waktu eksekusi (timeout / deadline)
//     - Meneruskan nilai request-scoped antar goroutine
//
//   Empat fungsi pembuat context turunan:
//     - context.WithCancel   → batalkan secara manual
//     - context.WithTimeout  → batalkan otomatis setelah durasi
//     - context.WithDeadline → batalkan otomatis pada waktu absolut
//     - context.WithValue    → sisipkan nilai ke dalam context
//
//   Jalankan test:
//     cd "DAY 29"
//     go test ./... -v -race
//
// ═══════════════════════════════════════════════════════════════════════════════
// 🧠 KAPAN PAKAI APA?
// ═══════════════════════════════════════════════════════════════════════════════
//
//   ┌──────────────────────┬──────────────────────────────────────────────────┐
//   │ Fungsi               │ Gunakan ketika...                                │
//   ├──────────────────────┼──────────────────────────────────────────────────┤
//   │ WithCancel           │ ingin membatalkan manual (tombol stop)           │
//   │ WithTimeout          │ operasi harus selesai dalam N detik              │
//   │ WithDeadline         │ harus selesai sebelum waktu X (absolut)          │
//   │ WithValue            │ kirim metadata (user ID, trace ID, dll.)         │
//   └──────────────────────┴──────────────────────────────────────────────────┘
//
//   Aturan emas:
//     1. Selalu `defer cancel()` setelah WithCancel/Timeout/Deadline!
//     2. Context adalah parameter PERTAMA fungsi (ctx context.Context).
//     3. Jangan simpan context dalam struct — lewatkan sebagai argumen.
//     4. Gunakan tipe khusus (bukan string mentah) sebagai key WithValue.
//
// ═══════════════════════════════════════════════════════════════════════════════

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 1: WITHCANCEL — PEMBATALAN MANUAL
// ═══════════════════════════════════════════════════════════════════════════════
//
// Ketika cancel() dipanggil:
//   - ctx.Done() tertutup → goroutine yang select <-ctx.Done() bangun
//   - ctx.Err() == context.Canceled
//   - Semua child context ikut dibatalkan
//
// Pola dasar goroutine yang bisa dibatalkan:
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//	go func() {
//	    for {
//	        select {
//	        case <-ctx.Done():
//	            return // berhenti dengan bersih
//	        default:
//	            // lakukan pekerjaan
//	        }
//	    }
//	}()
//

// BuatProsesLatar menjalankan fn di goroutine terpisah dengan context yang bisa dibatalkan.
// fn harus menghormati ctx — berhenti saat ctx.Done() tertutup.
// Mengembalikan fungsi batal untuk menghentikan proses dari luar.
//
// Contoh:
//
//	batal := BuatProsesLatar(func(ctx context.Context) {
//	    for {
//	        select {
//	        case <-ctx.Done():
//	            return
//	        default:
//	            time.Sleep(10 * time.Millisecond)
//	        }
//	    }
//	})
//	time.Sleep(100 * time.Millisecond)
//	batal()
//
// Hint: context.WithCancel → go fn(ctx) → return cancel
func BuatProsesLatar(fn func(ctx context.Context)) context.CancelFunc {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fn(ctx)
			}
		}
	}()

	return cancel
}

// TungguSampaiBatal memblokir hingga ctx dibatalkan.
// Mengembalikan ctx.Err() yang selalu non-nil saat fungsi ini return.
//
// Hint: <-ctx.Done() lalu return ctx.Err()
func TungguSampaiBatal(ctx context.Context) error {
	<-ctx.Done()

	return ctx.Err()
}

// JenisKesalahan mengembalikan penjelasan singkat dari error context:
//   - context.Canceled         → "dibatalkan"
//   - context.DeadlineExceeded → "kedaluwarsa"
//   - selainnya                → "tidak diketahui"
//
// Hint: errors.Is(err, context.Canceled) / errors.Is(err, context.DeadlineExceeded)
func JenisKesalahan(err error) string {
	if errors.Is(err, context.Canceled) {
		return "dibatalkan"
	} else if errors.Is(err, context.DeadlineExceeded) {
		return "kedaluwarsa"
	} else {
		return "tidak diketahui"
	}
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 2: WITHTIMEOUT — BATAS WAKTU OTOMATIS
// ═══════════════════════════════════════════════════════════════════════════════
//
// context.WithTimeout(parent, d) otomatis membatalkan context setelah durasi d.
// Setelah timeout: ctx.Err() == context.DeadlineExceeded.
//
// ⚠️ Penting: selalu defer cancel() meski timeout terjadi lebih dulu
//    untuk mencegah goroutine leak!
//

// JalankanDenganTimeout menjalankan fn dengan batas waktu `batas`.
// Jika fn tidak selesai dalam `batas`, ctx di dalam fn menjadi Done
// dan mengembalikan context.DeadlineExceeded.
// Jika fn selesai tepat waktu, kembalikan error dari fn (bisa nil).
//
// Contoh:
//
//	err := JalankanDenganTimeout(100*time.Millisecond, func(ctx context.Context) error {
//	    select {
//	    case <-ctx.Done():
//	        return ctx.Err()
//	    case <-time.After(50 * time.Millisecond):
//	        return nil
//	    }
//	})
//	// err == nil
//
// Hint: context.WithTimeout → defer cancel() → return fn(ctx)
func JalankanDenganTimeout(batas time.Duration, fn func(ctx context.Context) error) error {
	ctx1 := context.Background()
	ctx2, _ := context.WithTimeout(ctx1, batas)

	return fn(ctx2)
}

// JalankanDenganTimeoutParent seperti JalankanDenganTimeout, tapi menerima
// parent context dari luar sehingga timeout bersarang mengikuti parent.
//
// Hint: context.WithTimeout(parent, batas)
func JalankanDenganTimeoutParent(parent context.Context, batas time.Duration, fn func(ctx context.Context) error) error {
	ctx2, _ := context.WithTimeout(parent, batas)

	return fn(ctx2)
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 3: WITHDEADLINE — BATAS WAKTU ABSOLUT
// ═══════════════════════════════════════════════════════════════════════════════
//
// Perbedaan Deadline vs Timeout:
//   - Timeout: "berikan saya 5 detik dari SEKARANG"
//   - Deadline: "harus selesai sebelum jam 15:00:00" (waktu absolut)
//
// ctx.Deadline() mengembalikan (time.Time, bool)
// bool == false jika context tidak memiliki deadline.
//

// BuatContextDeadline membuat context yang kedaluwarsa pada waktu absolut `kapan`.
// Mengembalikan context dan fungsi batal.
//
// Hint: context.WithDeadline(context.Background(), kapan)
func BuatContextDeadline(kapan time.Time) (context.Context, context.CancelFunc) {
	ctx1 := context.Background()
	ctx, ok := context.WithDeadline(ctx1, kapan)
	return ctx, ok
}

// SisaWaktu mengembalikan sisa waktu sebelum ctx kedaluwarsa.
// Mengembalikan (durasi, true) jika ada deadline, (0, false) jika tidak ada.
// Durasi tidak boleh negatif — kembalikan 0 jika sudah lewat.
//
// Hint: ctx.Deadline() → time.Until(deadline)
func SisaWaktu(ctx context.Context) (time.Duration, bool) {
	val, ok := ctx.Deadline()
	kapan := time.Now().Add(100 * time.Second) // deadline sudah lewat
	if time.Until(time.Now())-time.Until(val) > 0 && time.Until(time.Now())-time.Until(val) < time.Until(kapan) {
		fmt.Println("masuk sini ")
		return time.Until(time.Now()) - time.Until(val), true
	} else if time.Until(val) < 0 {
		return 0, false
	}

	return time.Until(val), ok
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 4: WITHVALUE — NILAI DALAM CONTEXT
// ═══════════════════════════════════════════════════════════════════════════════
//
// ⚠️ KRITIS: Selalu gunakan tipe khusus sebagai key, BUKAN string biasa!
//
//   ❌ SALAH:  context.WithValue(ctx, "user_id", "123")
//   ✅ BENAR:  context.WithValue(ctx, KunciUserID, "123")
//
// Alasan: mencegah tabrakan key antar paket yang berbeda.
// WithValue hanya untuk data request-scoped. BUKAN untuk parameter fungsi.
//

// KunciCtx adalah tipe khusus untuk key context.WithValue.
type KunciCtx string

// Konstanta key yang digunakan untuk context request.
const (
	KunciUserID  KunciCtx = "user_id"
	KunciTraceID KunciCtx = "trace_id"
	KunciPeran   KunciCtx = "peran"
)

// SimpanNilai menyimpan pasangan key-value ke dalam context baru.
// Mengembalikan context turunan yang mengandung nilai tersebut.
//
// Hint: context.WithValue(parent, key, value)
func SimpanNilai(parent context.Context, key KunciCtx, value string) context.Context {
	ctx := context.WithValue(parent, key, value)

	return ctx
}

// AmbilNilai mengambil nilai dari context berdasarkan key.
// Mengembalikan ("", false) jika key tidak ada atau tipe tidak cocok.
//
// Hint: ctx.Value(key) → type assertion ke string
func AmbilNilai(ctx context.Context, key KunciCtx) (string, bool) {
	val := ctx.Value(key)
	s, ok := val.(string)
	if !ok {
		return "", false
	}
	if s == "" {
		return "", false
	}
	return s, true
}

// BuatContextRequest membuat context berisi metadata request (userID, traceID, peran).
// Simulasi middleware HTTP yang menyuntikkan info ke context sebelum handler.
//
// Contoh:
//
//	ctx := BuatContextRequest("user-42", "trace-abc", "admin")
//	uid, _ := AmbilNilai(ctx, KunciUserID)  // "user-42"
//	tid, _ := AmbilNilai(ctx, KunciTraceID) // "trace-abc"
//	per, _ := AmbilNilai(ctx, KunciPeran)   // "admin"
func BuatContextRequest(userID, traceID, peran string) context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, KunciUserID, userID)
	ctx = context.WithValue(ctx, KunciTraceID, traceID)
	ctx = context.WithValue(ctx, KunciPeran, peran)
	return ctx
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 5: PROPAGASI — HIERARKI CONTEXT
// ═══════════════════════════════════════════════════════════════════════════════
//
// Context membentuk pohon hierarki: parent → child → grandchild.
//
//   ✅ Membatalkan PARENT → semua descendant ikut batal
//   ❌ Membatalkan CHILD  → parent & saudara TIDAK terpengaruh
//   ✅ Nilai parent bisa dibaca child (mewarisi ke bawah)
//   ❌ Nilai child TIDAK bisa dibaca parent (tidak mewarisi ke atas)
//

// HitungAnakYangBatal membuat `jumlah` child context dari satu parent.
// Setiap child menjalankan goroutine yang menunggu ctx.Done() lalu kirim indeksnya ke channel.
//
// Mengembalikan:
//   - batalParent: fungsi yang membatalkan parent (sekaligus semua child)
//   - selesai:     buffered channel yang menerima indeks tiap child yang batal
//
// Hint:
//   - Buat parent dengan WithCancel
//   - Loop i: buat child dari parent, goroutine <-child.Done() → ch <- i
func HitungAnakYangBatal(jumlah int) (batalParent context.CancelFunc, selesai <-chan int) {
	ch := make(chan int)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		defer close(ch)
		for i := 0; i < jumlah; i++ {
			ch <- i
		}
	}()

	return cancel, ch
}

// PropagasiNilai menunjukkan arah propagasi nilai dalam hierarki context.
// Parent menyimpan KunciUserID; child menyimpan KunciTraceID.
//
// Mengembalikan:
//   - nilaiDariAnak:   nilai KunciUserID yang dibaca oleh child (harus ada)
//   - nilaiDariParent: nilai KunciTraceID yang dibaca oleh parent (harus kosong)
func PropagasiNilai() (nilaiDariAnak, nilaiDariParent string) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, KunciUserID, "")
	ctx = context.WithValue(ctx, KunciTraceID, "")
	return
} // ga nyambung aku ini

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 6: WORKER POOL DENGAN CONTEXT
// ═══════════════════════════════════════════════════════════════════════════════
//
// Alur worker pool:
//   masukan → [Worker 0] ↘
//   masukan → [Worker 1] → keluaran
//   masukan → [Worker 2] ↗
//
// Semua worker berhenti saat ctx dibatalkan atau channel masukan ditutup.
//

// Worker adalah unit pemrosesan tunggal yang context-aware.
type Worker struct {
	ID int
}

// Jalankan memulai loop pemrosesan worker.
// Membaca dari `masukan`, memanggil `proses`, mengirim hasil ke `keluaran`.
// Berhenti saat ctx.Done() tertutup ATAU channel masukan ditutup (ok == false).
//
// Pola idiomatis dua cabang select:
//
//	for {
//	    select {
//	    case <-ctx.Done():
//	        return
//	    case job, ok := <-masukan:
//	        if !ok { return }
//	        keluaran <- proses(job)
//	    }
//	}
func (w *Worker) Jalankan(ctx context.Context, masukan <-chan int, keluaran chan<- int, proses func(int) int) {
	go func() {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-masukan:
			if !ok {
				return
			}
			keluaran <- proses(job)
		}
	}()
}

// PoolWorker membuat pool `jumlah` Worker yang bekerja secara paralel.
// Semua worker membaca dari channel masukan yang sama (fan-out).
// Semua hasil dikirim ke channel keluaran bersama (fan-in).
// Channel keluaran ditutup setelah semua worker selesai.
//
// Contoh:
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//	masukan := make(chan int, 5)
//	for i := 1; i <= 5; i++ { masukan <- i }
//	close(masukan)
//	hasil := PoolWorker(ctx, 3, masukan, func(n int) int { return n * 2 })
//	for v := range hasil { fmt.Println(v) }
//
// Hint: WaitGroup + goroutine per worker + goroutine tutup channel setelah semua done
func PoolWorker(ctx context.Context, jumlah int, masukan <-chan int, proses func(int) int) <-chan int {

	keluaran := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < jumlah; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case val, ok := <-masukan:
					if !ok {
						return
					}
					hasil := proses(val)

					select {
					case <-ctx.Done():
						return
					case keluaran <- hasil:
					}
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(keluaran)
	}()

	return keluaran
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 7: PIPELINE DENGAN CONTEXT
// ═══════════════════════════════════════════════════════════════════════════════
//
// Pipeline = rantai goroutine yang saling terhubung via channel:
//
//   Bangkitkan → Transformasi → Transformasi → ... → hasil akhir
//
// Dengan context, seluruh pipeline bisa dihentikan sekaligus dari satu titik.
//

// Bangkitkan adalah generator: mengirim nums ke channel satu per satu.
// Channel ditutup setelah semua nums terkirim atau ctx dibatalkan.
//
// Hint:
//
//	go func() {
//	    defer close(ch)
//	    for _, n := range nums {
//	        select {
//	        case <-ctx.Done(): return
//	        case ch <- n:
//	        }
//	    }
//	}()
func Bangkitkan(ctx context.Context, nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, n := range nums {
			select {
			case <-ctx.Done():
				return
			case ch <- n:
			}
		}
	}()
	return ch
}

// Transformasi menerapkan fn ke setiap nilai dari channel masukan.
// Channel keluaran ditutup setelah masukan habis atau ctx dibatalkan.
//
// Contoh pipeline dua tahap:
//
//	gen     := Bangkitkan(ctx, 1, 2, 3)
//	tambah  := Transformasi(ctx, gen, func(n int) int { return n + 10 })
//	kuadrat := Transformasi(ctx, tambah, func(n int) int { return n * n })
//	// kuadrat: 121, 144, 169
func Transformasi(ctx context.Context, masukan <-chan int, fn func(int) int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for v := range masukan {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- fn(v)
			}
		}
	}()
	return ch
}

// GabungChannel menggabungkan beberapa channel menjadi satu (fan-in).
// Semua nilai dari `channels` diteruskan ke channel hasil.
// Channel hasil ditutup setelah semua masukan ditutup atau ctx dibatalkan.
func GabungChannel(ctx context.Context, channels ...<-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		select {
		case <-ctx.Done():
			return
		default:
			for _, v := range channels {
				for j := range v {
					select {
					case <-ctx.Done():
						return
					default:
						ch <- j
					}
				}
			}
		}
	}()

	return ch
}

// ═══════════════════════════════════════════════════════════════════════════════
// 📦 BAGIAN 8: RETRY DENGAN CONTEXT
// ═══════════════════════════════════════════════════════════════════════════════
//
// Retry adalah pola umum untuk operasi yang bisa gagal sementara.
// Dengan context, retry bisa dihentikan dari luar tanpa menunggu semua percobaan.
//
// Urutan prioritas:
//   1. fn() sukses (nil)  → langsung return nil
//   2. ctx batal           → return ctx.Err()
//   3. semua gagal         → return ErrMaksPercobaan
//

// ErrMaksPercobaan adalah error yang dikembalikan saat semua percobaan gagal.
var ErrMaksPercobaan = errors.New("semua percobaan gagal")

// CobaUlang mencoba memanggil fn hingga maksPercobaan kali.
//
// Berhenti lebih awal jika:
//   - fn() mengembalikan nil (berhasil)
//   - ctx dibatalkan antar jeda (return ctx.Err())
//
// Antar percobaan, tunggu `jeda` sambil tetap menghormati ctx.
// Jika semua percobaan habis → return ErrMaksPercobaan.
//
// Contoh:
//
//	n := 0
//	err := CobaUlang(ctx, 3, 5*time.Millisecond, func() error {
//	    n++
//	    if n < 3 { return errors.New("gagal") }
//	    return nil
//	})
//	// err == nil, n == 3
//
// Hint:
//
//	for i := 0; i < maksPercobaan; i++ {
//	    if fn() == nil { return nil }
//	    select {
//	    case <-ctx.Done(): return ctx.Err()
//	    case <-time.After(jeda):
//	    }
//	}
//	return ErrMaksPercobaan
func CobaUlang(ctx context.Context, maksPercobaan int, jeda time.Duration, fn func() error) error {
	for i := 0; i < maksPercobaan; i++ {
		if fn() == nil {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(jeda):
		}
	}
	return nil
}

// CobaUlangDenganBackoff seperti CobaUlang, tapi jeda bertumbuh eksponensial:
//
//	jeda pertama = jeda, kedua = 2×jeda, ketiga = 4×jeda, dst.
//
// Maksimum jeda dibatasi `jedaMaks` agar tidak terlalu lama.
// Backoff eksponensial mengurangi beban server saat terjadi gangguan massal.
//
// Hint: tunda *= 2; if tunda > jedaMaks { tunda = jedaMaks }
func CobaUlangDenganBackoff(ctx context.Context, maksPercobaan int, jeda, jedaMaks time.Duration, fn func() error) error {
	for i := 1; i <= maksPercobaan; i++ {
		if fn() == nil {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			delay := jeda * time.Duration(i)
			<-time.After(delay)
		}
	}
	return nil
}

// ─── pastikan import tidak dilaporkan unused ───────────────────────────────────
var (
	_ *sync.WaitGroup
	_ = time.Second
	_ = errors.New
)
