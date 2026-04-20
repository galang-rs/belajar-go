package belajar

import (
	"sort"
	"sync"
	"testing"
)

// =============================================================
// 1. TABLE-DRIVEN TEST - FanIn
// =============================================================

func TestFanIn(t *testing.T) {
	tests := []struct {
		name     string
		inputs   [][]int
		expected []int // sorted untuk perbandingan (urutan tidak dijamin)
	}{
		{
			"dua channel",
			[][]int{{1, 2}, {3, 4}},
			[]int{1, 2, 3, 4},
		},
		{
			"tiga channel",
			[][]int{{10}, {20}, {30}},
			[]int{10, 20, 30},
		},
		{
			"satu channel",
			[][]int{{5, 6, 7}},
			[]int{5, 6, 7},
		},
		{
			"channel kosong dan berisi",
			[][]int{{}, {1, 2}},
			[]int{1, 2},
		},
		{
			"semua channel kosong",
			[][]int{{}, {}},
			[]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Buat channel untuk setiap input
			channels := make([]<-chan int, len(tt.inputs))
			for i, vals := range tt.inputs {
				ch := make(chan int)
				channels[i] = ch
				go func(data []int) {
					for _, v := range data {
						ch <- v
					}
					close(ch)
				}(vals)
			}

			out := FanIn(channels...)
			if out == nil {
				t.Fatal("FanIn() mengembalikan nil, seharusnya channel")
			}

			var result []int
			for v := range out {
				result = append(result, v)
			}

			// Sort keduanya untuk perbandingan (urutan tidak dijamin)
			sort.Ints(result)
			sort.Ints(tt.expected)

			if len(result) != len(tt.expected) {
				t.Errorf("FanIn() menghasilkan %d item, want %d item", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("FanIn() result[%d] = %d; want %d", i, result[i], tt.expected[i])
				}
			}
		})
	}
}

// =============================================================
// 2. TEST - FanIn channel harus ditutup
// =============================================================

func TestFanIn_ChannelClosed(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
		close(ch1)
	}()
	go func() {
		ch2 <- 2
		close(ch2)
	}()

	out := FanIn(ch1, ch2)
	if out == nil {
		t.Fatal("FanIn() mengembalikan nil")
	}

	count := 0
	for range out {
		count++
	}
	// Jika kita sampai sini, berarti channel sudah ditutup dengan benar
	if count != 2 {
		t.Errorf("FanIn() menghasilkan %d item, want 2", count)
	}
}

// =============================================================
// 3. TABLE-DRIVEN TEST - Pipeline
// =============================================================

func TestPipeline(t *testing.T) {
	tests := []struct {
		name       string
		input      []int
		multiplier int
		expected   []int
	}{
		{
			"kalikan 10",
			[]int{1, 2, 3},
			10,
			[]int{10, 20, 30},
		},
		{
			"kalikan 3",
			[]int{5},
			3,
			[]int{15},
		},
		{
			"input kosong",
			[]int{},
			5,
			[]int{},
		},
		{
			"kalikan 1 (identitas)",
			[]int{7, 8, 9},
			1,
			[]int{7, 8, 9},
		},
		{
			"kalikan 0",
			[]int{1, 2, 3},
			0,
			[]int{0, 0, 0},
		},
		{
			"angka negatif",
			[]int{-2, -1, 0, 1, 2},
			2,
			[]int{-4, -2, 0, 2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Pipeline(tt.input, tt.multiplier)

			if len(result) != len(tt.expected) {
				t.Errorf("Pipeline(%v, %d) menghasilkan %d item, want %d",
					tt.input, tt.multiplier, len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Pipeline(%v, %d)[%d] = %d; want %d",
						tt.input, tt.multiplier, i, result[i], tt.expected[i])
				}
			}
		})
	}
}

// =============================================================
// 4. TABLE-DRIVEN TEST - WorkerPool
// =============================================================

func TestWorkerPool(t *testing.T) {
	tests := []struct {
		name       string
		jobs       []int
		numWorkers int
		expected   map[int]int
	}{
		{
			"3 jobs, 2 workers",
			[]int{2, 3, 4},
			2,
			map[int]int{2: 4, 3: 9, 4: 16},
		},
		{
			"1 job, 1 worker",
			[]int{5},
			1,
			map[int]int{5: 25},
		},
		{
			"jobs kosong",
			[]int{},
			3,
			map[int]int{},
		},
		{
			"5 jobs, 3 workers",
			[]int{1, 2, 3, 4, 5},
			3,
			map[int]int{1: 1, 2: 4, 3: 9, 4: 16, 5: 25},
		},
		{
			"angka negatif",
			[]int{-3, -2},
			2,
			map[int]int{-3: 9, -2: 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WorkerPool(tt.jobs, tt.numWorkers)

			if len(result) != len(tt.expected) {
				t.Errorf("WorkerPool() menghasilkan %d hasil, want %d", len(result), len(tt.expected))
				return
			}
			for k, v := range tt.expected {
				if result[k] != v {
					t.Errorf("WorkerPool() result[%d] = %d; want %d", k, result[k], v)
				}
			}
		})
	}
}

// =============================================================
// 5. TEST - SafeCounter
// =============================================================

func TestSafeCounter_Basic(t *testing.T) {
	sc := NewSafeCounter()
	if sc == nil {
		t.Fatal("NewSafeCounter() mengembalikan nil")
	}

	if v := sc.Value(); v != 0 {
		t.Errorf("NewSafeCounter().Value() = %d; want 0", v)
	}

	sc.Increment()
	sc.Increment()
	sc.Increment()

	if v := sc.Value(); v != 3 {
		t.Errorf("Setelah 3x Increment(), Value() = %d; want 3", v)
	}
}

// =============================================================
// 6. TEST - SafeCounter concurrent safety
// =============================================================

func TestSafeCounter_Concurrent(t *testing.T) {
	sc := NewSafeCounter()
	if sc == nil {
		t.Fatal("NewSafeCounter() mengembalikan nil")
	}

	numGoroutines := 100
	incrementsPerGoroutine := 100

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				sc.Increment()
			}
		}()
	}

	wg.Wait()

	expected := numGoroutines * incrementsPerGoroutine
	if v := sc.Value(); v != expected {
		t.Errorf("Setelah %d concurrent increments, Value() = %d; want %d", expected, v, expected)
	}
}

// =============================================================
// 7. BENCHMARK
// =============================================================

func BenchmarkPipeline(b *testing.B) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = i
	}
	for i := 0; i < b.N; i++ {
		Pipeline(input, 2)
	}
}

func BenchmarkWorkerPool(b *testing.B) {
	jobs := make([]int, 100)
	for i := range jobs {
		jobs[i] = i
	}
	for i := 0; i < b.N; i++ {
		WorkerPool(jobs, 4)
	}
}

func BenchmarkSafeCounter(b *testing.B) {
	sc := NewSafeCounter()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sc.Increment()
		}
	})
}
