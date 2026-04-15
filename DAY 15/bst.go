package belajar

// ==================== DAY 15: BINARY SEARCH TREE (BST) ====================
// Topik: Struktur data Binary Search Tree - insert, search, dan traversal.
// Konsep penting: BST memiliki aturan: node kiri < parent < node kanan.

// BSTNode merepresentasikan satu node dalam Binary Search Tree.
type BSTNode struct {
	Value int
	Left  *BSTNode
	Right *BSTNode
}

// BST merepresentasikan Binary Search Tree.
type BST struct {
	Root *BSTNode
}

// NewBST membuat BST baru yang kosong.
// Contoh: tree := NewBST() -> BST kosong
func NewBST() *BST {
	bst := &BST{}
	return bst
}

// Insert menambahkan nilai ke BST.
// Jika nilai sudah ada, abaikan (tidak ada duplikat).
// Aturan: nilai lebih kecil ke kiri, lebih besar ke kanan.
// Contoh:
//
//	tree := NewBST()
//	tree.Insert(5)
//	tree.Insert(3)
//	tree.Insert(7)
//	// Menghasilkan:
//	//     5
//	//    / \
//	//   3   7
//
//	tree.Insert(5) // diabaikan, sudah ada
func (b *BST) Insert(val int) {
	// TODO: implementasi di sini
	if b.Root == nil {
		b.Root = &BSTNode{Value: val}
		return
	}

	if val < b.Root.Value {
		if b.Root.Left == nil {
			b.Root.Left = &BSTNode{Value: val}
		} else {
			b.Insert(val)
		}
	} else {
		if b.Root.Right == nil {
			b.Root.Right = &BSTNode{Value: val}
		} else {
			b.Insert(val)
		}
	}
}

// Search mengecek apakah nilai ada di dalam BST.
// Contoh:
//
//	tree berisi {5, 3, 7}
//	tree.Search(3) -> true
//	tree.Search(10) -> false
//	tree.Search(5) -> true
//
//	BST kosong:
//	tree.Search(1) -> false
func (b *BST) Search(val int) bool {
	// TODO: implementasi di sini
	return false
}

// InOrder mengembalikan slice berisi nilai BST dalam urutan in-order traversal (kiri, root, kanan).
// Hasil in-order traversal pada BST yang valid selalu terurut ascending.
// Contoh:
//
//	tree berisi {5, 3, 7, 1, 4}
//	tree.InOrder() -> []int{1, 3, 4, 5, 7}
//
//	BST kosong:
//	tree.InOrder() -> []int{}
//
// Hint: gunakan rekursi - kunjungi kiri dulu, lalu root, lalu kanan.
func (b *BST) InOrder() []int {
	// TODO: implementasi di sini
	return nil
}

// aku sama sekali tidak paham konsep ini dikarena baru pertama kali mendapatkan test seperti ini dan konsep ini
