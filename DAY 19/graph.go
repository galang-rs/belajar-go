package belajar

// ==================== DAY 19: GRAPH DASAR ====================
// Topik: Representasi graph menggunakan adjacency list, BFS dan DFS traversal.
// Konsep penting: graph adalah kumpulan node (vertex) yang terhubung oleh edge.

// Graph merepresentasikan unweighted directed graph menggunakan adjacency list.
type Graph struct {
	adjacency map[int][]int
}

// NewGraph membuat Graph baru yang kosong.
// Contoh: g := NewGraph() -> graph kosong
func NewGraph() *Graph {
	// TODO: implementasi di sini
	return nil
}

// AddEdge menambahkan edge dari node src ke node dst (directed).
// Jika node belum ada, otomatis ditambahkan.
// Contoh:
//
//	g := NewGraph()
//	g.AddEdge(1, 2)
//	g.AddEdge(1, 3)
//	g.AddEdge(2, 4)
//	// Graph: 1 -> [2, 3], 2 -> [4]
func (g *Graph) AddEdge(src, dst int) {
	// TODO: implementasi di sini
}

// BFS melakukan Breadth-First Search dari node start.
// Mengembalikan slice node dalam urutan kunjungan BFS.
// Jika start tidak ada di graph, kembalikan slice kosong.
// Jika ada beberapa tetangga, kunjungi sesuai urutan mereka ditambahkan via AddEdge.
// Contoh:
//
//	Graph: 1->[2,3], 2->[4], 3->[4]
//	g.BFS(1) -> []int{1, 2, 3, 4}
//
//	Graph kosong:
//	g.BFS(1) -> []int{}
//
// Hint: gunakan Queue (slice sebagai queue). Tandai node yang sudah dikunjungi.
func (g *Graph) BFS(start int) []int {
	// TODO: implementasi di sini
	return nil
}

// DFS melakukan Depth-First Search dari node start.
// Mengembalikan slice node dalam urutan kunjungan DFS.
// Jika start tidak ada di graph, kembalikan slice kosong.
// Jika ada beberapa tetangga, kunjungi sesuai urutan mereka ditambahkan via AddEdge.
// Contoh:
//
//	Graph: 1->[2,3], 2->[4], 3->[4]
//	g.DFS(1) -> []int{1, 2, 4, 3}
//
//	Graph kosong:
//	g.DFS(1) -> []int{}
//
// Hint: gunakan Stack (slice sebagai stack) atau rekursi. Tandai node yang sudah dikunjungi.
func (g *Graph) DFS(start int) []int {
	// TODO: implementasi di sini
	return nil
}
