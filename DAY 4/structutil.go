package belajar

// Student merepresentasikan data seorang siswa.
type Student struct {
	Name  string
	Age   int
	Grade float64
}

// NewStudent membuat instance Student baru.
// Contoh: NewStudent("Andi", 17, 85.5) -> Student{Name: "Andi", Age: 17, Grade: 85.5}
func NewStudent(name string, age int, grade float64) Student {
	// TODO: implementasi di sini
	return Student{}
}

// IsPass mengecek apakah siswa lulus (Grade >= 70.0).
// Contoh: Student{Grade: 85}.IsPass() -> true
//
//	Student{Grade: 65}.IsPass() -> false
//	Student{Grade: 70}.IsPass() -> true (batas tepat = lulus)
func (s Student) IsPass() bool {
	// TODO: implementasi di sini
	return false
}

// Info mengembalikan string informasi siswa dalam format:
// "Nama (Umur tahun) - Nilai: Grade"
// Grade diformat 2 angka di belakang koma.
// Contoh: Student{Name: "Andi", Age: 17, Grade: 85.5}.Info() -> "Andi (17 tahun) - Nilai: 85.50"
//
//	Student{Name: "Budi", Age: 16, Grade: 100.0}.Info() -> "Budi (16 tahun) - Nilai: 100.00"
func (s Student) Info() string {
	// TODO: implementasi di sini
	// Hint: gunakan fmt.Sprintf
	return ""
}

// AverageGrade menghitung rata-rata Grade dari slice Student.
// Kembalikan 0 jika slice kosong.
// Contoh: AverageGrade([]Student{{Grade: 80}, {Grade: 90}}) -> 85.0
//
//	AverageGrade([]Student{}) -> 0
func AverageGrade(students []Student) float64 {
	// TODO: implementasi di sini
	return 0
}

// TopStudent mengembalikan Student dengan Grade tertinggi.
// Jika ada yang sama, kembalikan yang pertama ditemukan.
// Kembalikan Student kosong dan true (error) jika slice kosong.
// Contoh: TopStudent([]Student{{Name: "A", Grade: 90}, {Name: "B", Grade: 80}}) -> Student{Name: "A", Grade: 90}, false
//
//	TopStudent([]Student{}) -> Student{}, true
func TopStudent(students []Student) (Student, bool) {
	// TODO: implementasi di sini
	return Student{}, true
}

// FilterByMinGrade mengembalikan slice Student yang Grade-nya >= minGrade.
// Contoh: FilterByMinGrade([]Student{{Name: "A", Grade: 90}, {Name: "B", Grade: 60}}, 70) -> []Student{{Name: "A", Grade: 90}}
//
//	FilterByMinGrade([]Student{}, 70) -> []Student{}
func FilterByMinGrade(students []Student, minGrade float64) []Student {
	// TODO: implementasi di sini
	return nil
}

// SortByGrade mengurutkan slice Student berdasarkan Grade secara descending (tertinggi dulu).
// Mengembalikan slice baru (TIDAK mengubah slice asli).
// Jika Grade sama, urutan berdasarkan kemunculan asli (stable sort).
// Contoh: SortByGrade([]Student{{Name: "A", Grade: 70}, {Name: "B", Grade: 90}})
//
//	-> []Student{{Name: "B", Grade: 90}, {Name: "A", Grade: 70}}
func SortByGrade(students []Student) []Student {
	// TODO: implementasi di sini
	// Hint: gunakan sort.SliceStable
	return nil
}

// CountByPassFail menghitung jumlah siswa yang lulus dan tidak lulus (Grade >= 70 = lulus).
// Return pertama = jumlah lulus, return kedua = jumlah tidak lulus.
// Contoh: CountByPassFail([]Student{{Grade: 80}, {Grade: 60}, {Grade: 70}}) -> 2, 1
//
//	CountByPassFail([]Student{}) -> 0, 0
func CountByPassFail(students []Student) (int, int) {
	// TODO: implementasi di sini
	return 0, 0
}

// UniqueNames mengembalikan daftar nama unik dari slice Student.
// Urutan berdasarkan kemunculan pertama.
// Contoh: UniqueNames([]Student{{Name: "Andi"}, {Name: "Budi"}, {Name: "Andi"}}) -> []string{"Andi", "Budi"}
//
//	UniqueNames([]Student{}) -> []string{}
func UniqueNames(students []Student) []string {
	// TODO: implementasi di sini
	return nil
}

// NamesByGradeRange mengembalikan nama-nama siswa yang Grade-nya berada di antara min dan max (inklusif).
// Contoh: NamesByGradeRange([]Student{{Name: "A", Grade: 80}, {Name: "B", Grade: 60}}, 70, 90) -> []string{"A"}
//
//	NamesByGradeRange([]Student{{Name: "A", Grade: 70}}, 70, 70) -> []string{"A"}
func NamesByGradeRange(students []Student, min, max float64) []string {
	// TODO: implementasi di sini
	return nil
}
