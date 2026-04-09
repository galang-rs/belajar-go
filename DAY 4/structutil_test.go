package belajar

import (
	"reflect"
	"testing"
)

// =============================================================
// 1. TABLE-DRIVEN TEST - NewStudent
// =============================================================

func TestNewStudent(t *testing.T) {
	tests := []struct {
		name     string
		sName    string
		age      int
		grade    float64
		expected Student
	}{
		{"normal", "Andi", 17, 85.5, Student{Name: "Andi", Age: 17, Grade: 85.5}},
		{"grade nol", "Budi", 16, 0, Student{Name: "Budi", Age: 16, Grade: 0}},
		{"grade sempurna", "Cici", 18, 100, Student{Name: "Cici", Age: 18, Grade: 100}},
		{"nama kosong", "", 15, 50, Student{Name: "", Age: 15, Grade: 50}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewStudent(tt.sName, tt.age, tt.grade)
			if result != tt.expected {
				t.Errorf("NewStudent(%q, %d, %v) = %v; want %v", tt.sName, tt.age, tt.grade, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 2. TABLE-DRIVEN TEST - IsPass
// =============================================================

func TestIsPass(t *testing.T) {
	tests := []struct {
		name     string
		student  Student
		expected bool
	}{
		{"lulus tinggi", Student{Grade: 90}, true},
		{"tepat batas", Student{Grade: 70}, true},
		{"tidak lulus", Student{Grade: 69.9}, false},
		{"grade nol", Student{Grade: 0}, false},
		{"grade sempurna", Student{Grade: 100}, true},
		{"sedikit di bawah", Student{Grade: 69.99}, false},
		{"sedikit di atas", Student{Grade: 70.01}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.student.IsPass()
			if result != tt.expected {
				t.Errorf("Student{Grade: %v}.IsPass() = %v; want %v", tt.student.Grade, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 3. TABLE-DRIVEN TEST - Info
// =============================================================

func TestInfo(t *testing.T) {
	tests := []struct {
		name     string
		student  Student
		expected string
	}{
		{"normal", Student{Name: "Andi", Age: 17, Grade: 85.5}, "Andi (17 tahun) - Nilai: 85.50"},
		{"grade bulat", Student{Name: "Budi", Age: 16, Grade: 100}, "Budi (16 tahun) - Nilai: 100.00"},
		{"grade nol", Student{Name: "Cici", Age: 15, Grade: 0}, "Cici (15 tahun) - Nilai: 0.00"},
		{"grade desimal panjang", Student{Name: "Dedi", Age: 18, Grade: 77.777}, "Dedi (18 tahun) - Nilai: 77.78"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.student.Info()
			if result != tt.expected {
				t.Errorf("Student{%q, %d, %v}.Info() = %q; want %q", tt.student.Name, tt.student.Age, tt.student.Grade, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 4. TABLE-DRIVEN TEST - AverageGrade
// =============================================================

func TestAverageGrade(t *testing.T) {
	tests := []struct {
		name     string
		students []Student
		expected float64
	}{
		{"slice kosong", []Student{}, 0},
		{"satu siswa", []Student{{Grade: 80}}, 80},
		{"dua siswa", []Student{{Grade: 80}, {Grade: 90}}, 85},
		{"tiga siswa", []Student{{Grade: 70}, {Grade: 80}, {Grade: 90}}, 80},
		{"semua sama", []Student{{Grade: 75}, {Grade: 75}, {Grade: 75}}, 75},
		{"ada nol", []Student{{Grade: 0}, {Grade: 100}}, 50},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AverageGrade(tt.students)
			if result != tt.expected {
				t.Errorf("AverageGrade(...) = %v; want %v", result, tt.expected)
			}
		})
	}
}

// =============================================================
// 5. TABLE-DRIVEN TEST - TopStudent
// =============================================================

func TestTopStudent(t *testing.T) {
	tests := []struct {
		name      string
		students  []Student
		expected  Student
		expectErr bool
	}{
		{"slice kosong", []Student{}, Student{}, true},
		{"satu siswa", []Student{{Name: "A", Grade: 80}}, Student{Name: "A", Grade: 80}, false},
		{"dua siswa", []Student{{Name: "A", Grade: 80}, {Name: "B", Grade: 90}}, Student{Name: "B", Grade: 90}, false},
		{"grade sama ambil pertama", []Student{{Name: "A", Grade: 90}, {Name: "B", Grade: 90}}, Student{Name: "A", Grade: 90}, false},
		{"banyak siswa", []Student{
			{Name: "A", Grade: 70},
			{Name: "B", Grade: 95},
			{Name: "C", Grade: 88},
			{Name: "D", Grade: 60},
		}, Student{Name: "B", Grade: 95}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := TopStudent(tt.students)

			if tt.expectErr {
				if !err {
					t.Error("TopStudent() expected error for empty slice, got none")
				}
				return
			}

			if err {
				t.Errorf("TopStudent() unexpected error")
				return
			}

			if result.Name != tt.expected.Name || result.Grade != tt.expected.Grade {
				t.Errorf("TopStudent(...) = %v; want %v", result, tt.expected)
			}
		})
	}
}

// =============================================================
// 6. TABLE-DRIVEN TEST - FilterByMinGrade
// =============================================================

func TestFilterByMinGrade(t *testing.T) {
	allStudents := []Student{
		{Name: "A", Age: 17, Grade: 90},
		{Name: "B", Age: 16, Grade: 60},
		{Name: "C", Age: 18, Grade: 75},
		{Name: "D", Age: 17, Grade: 50},
		{Name: "E", Age: 16, Grade: 85},
	}

	tests := []struct {
		name     string
		students []Student
		minGrade float64
		expected []Student
	}{
		{"semua lolos", allStudents, 0, allStudents},
		{"tidak ada yang lolos", allStudents, 100, []Student{}},
		{"filter 70", allStudents, 70, []Student{
			{Name: "A", Age: 17, Grade: 90},
			{Name: "C", Age: 18, Grade: 75},
			{Name: "E", Age: 16, Grade: 85},
		}},
		{"filter tepat batas", allStudents, 75, []Student{
			{Name: "A", Age: 17, Grade: 90},
			{Name: "C", Age: 18, Grade: 75},
			{Name: "E", Age: 16, Grade: 85},
		}},
		{"slice kosong", []Student{}, 70, []Student{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FilterByMinGrade(tt.students, tt.minGrade)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("FilterByMinGrade(..., %v) = %v; want %v", tt.minGrade, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 7. TABLE-DRIVEN TEST - SortByGrade
// =============================================================

func TestSortByGrade(t *testing.T) {
	tests := []struct {
		name     string
		students []Student
		expected []Student
	}{
		{"slice kosong", []Student{}, []Student{}},
		{"satu siswa", []Student{{Name: "A", Grade: 80}}, []Student{{Name: "A", Grade: 80}}},
		{"sudah terurut descending", []Student{
			{Name: "A", Grade: 90},
			{Name: "B", Grade: 80},
		}, []Student{
			{Name: "A", Grade: 90},
			{Name: "B", Grade: 80},
		}},
		{"perlu diurutkan", []Student{
			{Name: "A", Grade: 70},
			{Name: "B", Grade: 90},
			{Name: "C", Grade: 80},
		}, []Student{
			{Name: "B", Grade: 90},
			{Name: "C", Grade: 80},
			{Name: "A", Grade: 70},
		}},
		{"grade sama urutan stabil", []Student{
			{Name: "A", Grade: 80},
			{Name: "B", Grade: 80},
			{Name: "C", Grade: 90},
		}, []Student{
			{Name: "C", Grade: 90},
			{Name: "A", Grade: 80},
			{Name: "B", Grade: 80},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SortByGrade(tt.students)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SortByGrade(...) = %v; want %v", result, tt.expected)
			}
		})
	}

	// Test bahwa slice asli tidak berubah
	t.Run("tidak mengubah slice asli", func(t *testing.T) {
		original := []Student{
			{Name: "A", Grade: 70},
			{Name: "B", Grade: 90},
		}
		originalCopy := make([]Student, len(original))
		copy(originalCopy, original)

		SortByGrade(original)

		if !reflect.DeepEqual(original, originalCopy) {
			t.Errorf("SortByGrade() mengubah slice asli: got %v; want %v", original, originalCopy)
		}
	})
}

// =============================================================
// 8. TABLE-DRIVEN TEST - CountByPassFail
// =============================================================

func TestCountByPassFail(t *testing.T) {
	tests := []struct {
		name       string
		students   []Student
		expectPass int
		expectFail int
	}{
		{"slice kosong", []Student{}, 0, 0},
		{"semua lulus", []Student{{Grade: 80}, {Grade: 90}, {Grade: 70}}, 3, 0},
		{"semua gagal", []Student{{Grade: 50}, {Grade: 60}, {Grade: 69}}, 0, 3},
		{"campuran", []Student{{Grade: 80}, {Grade: 60}, {Grade: 70}, {Grade: 50}}, 2, 2},
		{"tepat batas", []Student{{Grade: 70}, {Grade: 69.99}}, 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pass, fail := CountByPassFail(tt.students)
			if pass != tt.expectPass || fail != tt.expectFail {
				t.Errorf("CountByPassFail(...) = (%d, %d); want (%d, %d)", pass, fail, tt.expectPass, tt.expectFail)
			}
		})
	}
}

// =============================================================
// 9. TABLE-DRIVEN TEST - UniqueNames
// =============================================================

func TestUniqueNames(t *testing.T) {
	tests := []struct {
		name     string
		students []Student
		expected []string
	}{
		{"slice kosong", []Student{}, []string{}},
		{"semua unik", []Student{{Name: "A"}, {Name: "B"}, {Name: "C"}}, []string{"A", "B", "C"}},
		{"ada duplikat", []Student{{Name: "Andi"}, {Name: "Budi"}, {Name: "Andi"}, {Name: "Cici"}}, []string{"Andi", "Budi", "Cici"}},
		{"semua sama", []Student{{Name: "A"}, {Name: "A"}, {Name: "A"}}, []string{"A"}},
		{"satu siswa", []Student{{Name: "Solo"}}, []string{"Solo"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := UniqueNames(tt.students)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("UniqueNames(...) = %v; want %v", result, tt.expected)
			}
		})
	}
}

// =============================================================
// 10. TABLE-DRIVEN TEST - NamesByGradeRange
// =============================================================

func TestNamesByGradeRange(t *testing.T) {
	students := []Student{
		{Name: "Andi", Grade: 90},
		{Name: "Budi", Grade: 60},
		{Name: "Cici", Grade: 75},
		{Name: "Dedi", Grade: 80},
		{Name: "Eka", Grade: 50},
	}

	tests := []struct {
		name     string
		students []Student
		min      float64
		max      float64
		expected []string
	}{
		{"range lebar", students, 0, 100, []string{"Andi", "Budi", "Cici", "Dedi", "Eka"}},
		{"range sempit", students, 70, 80, []string{"Cici", "Dedi"}},
		{"tepat batas", students, 75, 75, []string{"Cici"}},
		{"tidak ada yang cocok", students, 95, 100, []string{}},
		{"slice kosong", []Student{}, 0, 100, []string{}},
		{"range bawah", students, 50, 60, []string{"Budi", "Eka"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NamesByGradeRange(tt.students, tt.min, tt.max)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("NamesByGradeRange(..., %v, %v) = %v; want %v", tt.min, tt.max, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 11. PARALLEL TEST - AverageGrade
// =============================================================

func TestAverageGrade_Parallel(t *testing.T) {
	tests := []struct {
		name     string
		students []Student
		expected float64
	}{
		{"case1", []Student{{Grade: 100}, {Grade: 100}}, 100},
		{"case2", []Student{{Grade: 0}, {Grade: 50}, {Grade: 100}}, 50},
		{"case3", []Student{{Grade: 75}}, 75},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := AverageGrade(tt.students)
			if result != tt.expected {
				t.Errorf("AverageGrade(...) = %v; want %v", result, tt.expected)
			}
		})
	}
}

// =============================================================
// 12. BENCHMARK TEST
// =============================================================

func BenchmarkSortByGrade(b *testing.B) {
	students := []Student{
		{Name: "A", Grade: 70},
		{Name: "B", Grade: 95},
		{Name: "C", Grade: 88},
		{Name: "D", Grade: 60},
		{Name: "E", Grade: 92},
		{Name: "F", Grade: 45},
		{Name: "G", Grade: 78},
		{Name: "H", Grade: 83},
	}
	for i := 0; i < b.N; i++ {
		SortByGrade(students)
	}
}

func BenchmarkFilterByMinGrade(b *testing.B) {
	students := []Student{
		{Name: "A", Grade: 70},
		{Name: "B", Grade: 95},
		{Name: "C", Grade: 88},
		{Name: "D", Grade: 60},
		{Name: "E", Grade: 92},
	}
	for i := 0; i < b.N; i++ {
		FilterByMinGrade(students, 75)
	}
}
