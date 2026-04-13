package belajar

import (
	"reflect"
	"testing"
)

// =============================================================
// 1. TABLE-DRIVEN TEST - BFS
// =============================================================

func TestGraph_BFS(t *testing.T) {
	tests := []struct {
		name     string
		edges    [][2]int
		start    int
		expected []int
	}{
		{
			"graph standar",
			[][2]int{{1, 2}, {1, 3}, {2, 4}, {3, 4}},
			1,
			[]int{1, 2, 3, 4},
		},
		{
			"linear",
			[][2]int{{1, 2}, {2, 3}, {3, 4}},
			1,
			[]int{1, 2, 3, 4},
		},
		{
			"satu node",
			[][2]int{{1, 2}},
			1,
			[]int{1, 2},
		},
		{
			"start tidak ada",
			[][2]int{{1, 2}},
			99,
			[]int{},
		},
		{
			"graph kosong",
			[][2]int{},
			1,
			[]int{},
		},
		{
			"bercabang banyak",
			[][2]int{{1, 2}, {1, 3}, {1, 4}, {2, 5}, {3, 5}},
			1,
			[]int{1, 2, 3, 4, 5},
		},
		{
			"start di tengah",
			[][2]int{{1, 2}, {2, 3}, {3, 4}},
			2,
			[]int{2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph()
			for _, e := range tt.edges {
				g.AddEdge(e[0], e[1])
			}
			result := g.BFS(tt.start)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BFS(%d) = %v; want %v", tt.start, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 2. TABLE-DRIVEN TEST - DFS
// =============================================================

func TestGraph_DFS(t *testing.T) {
	tests := []struct {
		name     string
		edges    [][2]int
		start    int
		expected []int
	}{
		{
			"graph standar",
			[][2]int{{1, 2}, {1, 3}, {2, 4}, {3, 4}},
			1,
			[]int{1, 2, 4, 3},
		},
		{
			"linear",
			[][2]int{{1, 2}, {2, 3}, {3, 4}},
			1,
			[]int{1, 2, 3, 4},
		},
		{
			"satu edge",
			[][2]int{{1, 2}},
			1,
			[]int{1, 2},
		},
		{
			"start tidak ada",
			[][2]int{{1, 2}},
			99,
			[]int{},
		},
		{
			"graph kosong",
			[][2]int{},
			1,
			[]int{},
		},
		{
			"bercabang",
			[][2]int{{1, 2}, {1, 3}, {2, 4}},
			1,
			[]int{1, 2, 4, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph()
			for _, e := range tt.edges {
				g.AddEdge(e[0], e[1])
			}
			result := g.DFS(tt.start)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("DFS(%d) = %v; want %v", tt.start, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 3. TEST - BFS dan DFS mengunjungi semua node yang reachable
// =============================================================

func TestGraph_BFS_DFS_SameNodes(t *testing.T) {
	g := NewGraph()
	edges := [][2]int{{1, 2}, {1, 3}, {2, 4}, {3, 4}, {4, 5}}
	for _, e := range edges {
		g.AddEdge(e[0], e[1])
	}

	bfsResult := g.BFS(1)
	dfsResult := g.DFS(1)

	if len(bfsResult) != len(dfsResult) {
		t.Errorf("BFS mengunjungi %d node, DFS mengunjungi %d node; harus sama",
			len(bfsResult), len(dfsResult))
	}

	// Pastikan keduanya mengunjungi node yang sama (meskipun urutan beda)
	bfsSet := make(map[int]bool)
	dfsSet := make(map[int]bool)
	for _, v := range bfsResult {
		bfsSet[v] = true
	}
	for _, v := range dfsResult {
		dfsSet[v] = true
	}
	if !reflect.DeepEqual(bfsSet, dfsSet) {
		t.Errorf("BFS dan DFS harus mengunjungi node yang sama: BFS=%v, DFS=%v", bfsResult, dfsResult)
	}
}

// =============================================================
// 4. TEST - AddEdge
// =============================================================

func TestGraph_AddEdge(t *testing.T) {
	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)

	result := g.BFS(1)
	if len(result) != 3 {
		t.Errorf("Setelah AddEdge(1,2) dan AddEdge(1,3), BFS(1) harus mengunjungi 3 node, got %d", len(result))
	}
}

// =============================================================
// 5. BENCHMARK
// =============================================================

func BenchmarkGraph_BFS(b *testing.B) {
	g := NewGraph()
	for i := 0; i < 100; i++ {
		g.AddEdge(i, i+1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.BFS(0)
	}
}

func BenchmarkGraph_DFS(b *testing.B) {
	g := NewGraph()
	for i := 0; i < 100; i++ {
		g.AddEdge(i, i+1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.DFS(0)
	}
}
