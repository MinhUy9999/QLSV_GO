package main

import (
	"fmt"
	"strings"
)

// Khai báo struct SinhVien
type SinhVien struct {
	ID       string
	Ten      string
	Tuoi     int
	GioiTinh string
	Lop      string
}

// Khai báo map dùng để lưu trữ sinh viên
var sinhVienMap = make(map[string]SinhVien)

// Hàm thêm sinh viên
func themSinhVien(sv SinhVien) {
	sinhVienMap[sv.ID] = sv
	fmt.Println("Đã thêm sinh viên:", sv.Ten)
}

// Hàm tìm kiếm sinh viên theo ID
func timSinhVien(id string) (SinhVien, bool) {
	sv, ok := sinhVienMap[id]
	return sv, ok
}

// Hàm sửa thông tin sinh viên
func suaSinhVien(id string, sv SinhVien) {
	_, exists := sinhVienMap[id]
	if exists {
		sinhVienMap[id] = sv
		fmt.Println("Đã sửa thông tin sinh viên:", sv.Ten)
	} else {
		fmt.Println("Không tìm thấy sinh viên với ID:", id)
	}
}

// Hàm xóa sinh viên
func xoaSinhVien(id string) {
	_, exists := sinhVienMap[id]
	if exists {
		delete(sinhVienMap, id)
		fmt.Println("Đã xóa sinh viên với ID:", id)
	} else {
		fmt.Println("Không tìm thấy sinh viên với ID:", id)
	}
}

// Hàm xuất ra danh sách sinh viên theo giới tính
func xuatSinhVienTheoGioiTinh(gioiTinh string) {
	gioiTinh = strings.ToLower(strings.TrimSpace(gioiTinh)) // Chuẩn hóa chuỗi nhập vào

	if gioiTinh != "nam" && gioiTinh != "nu" {
		fmt.Println("Giới tính không hợp lệ. Vui lòng nhập 'Nam' hoặc 'Nu'.")
		return
	}

	fmt.Println("Danh sách sinh viên giới tính", strings.Title(gioiTinh)+":")
	for _, sv := range sinhVienMap {
		if strings.ToLower(sv.GioiTinh) == gioiTinh {
			fmt.Println(sv)
		}
	}
}

func main() {
	// Thêm một số sinh viên
	themSinhVien(SinhVien{ID: "SV001", Ten: "Trương Phạm Minh Uy", Tuoi: 26, GioiTinh: "Nam", Lop: "CNTT1"})
	themSinhVien(SinhVien{ID: "SV002", Ten: "Tiêu Viêm", Tuoi: 21, GioiTinh: "Nam", Lop: "CNTT1"})
	themSinhVien(SinhVien{ID: "SV003", Ten: "Trần Thị Ngọc Thi", Tuoi: 25, GioiTinh: "Nu", Lop: "CNTT2"})

	// Tìm kiếm sinh viên
	sv, found := timSinhVien("SV002")
	if found {
		fmt.Println("Tìm thấy sinh viên:", sv)
	} else {
		fmt.Println("Không tìm thấy sinh viên với ID: SV002")
	}

	// Sửa sinh viên
	suaSinhVien("SV002", SinhVien{ID: "SV002", Ten: "Tiêu Viêm (Updated)", Tuoi: 22, GioiTinh: "Nam", Lop: "CNTT1"})

	// Xóa sinh viên
	xoaSinhVien("SV001")

	// Xuất danh sách sinh viên theo giới tính
	xuatSinhVienTheoGioiTinh("Nam")
	xuatSinhVienTheoGioiTinh("Nu")
}
