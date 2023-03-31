package pmb1214031

import (
	"fmt"
	"testing"

	"github.com/erditona/be_pmb/model"
	"github.com/erditona/be_pmb/module"
)

// test Insert

func TestInsertPendaftaran(t *testing.T) {
	kdpendaftar := 2
	biodata := model.Camaba{
		Ktp : 3201323211222,
		Nama : "Nausha Adam",
		Phone_number : "085718177810",
		Address : "Parongpong, Kab. Bandung Barat",
	}
	asalsekolah := model.DaftarSekolah{
		KDSekolah : 01,
		Nama : "SMK Negeri 1 Cirebon",
		Phone_number : "085718172053",
		Address : "Jl.Perjuangan, Kota Cirebon",
	}
	jurusan := model.Jurusan{
		KDJurusan : "D4TI",
		Nama : "SMK Negeri 1 Cirebon",
		Jenjang : "Diploma 4",
	}
	jalur := "Rapot"
	alulbi := "Universitas Internasional"
	aljurusan := "Sedang trand"
	hasil:=module.InsertPendaftaran(module.MongoConn,"pendaftaran_maba",kdpendaftar, biodata, asalsekolah, jurusan, jalur, alulbi, aljurusan)
	fmt.Println(hasil)
}

func TestInsertDaftarCamaba(t *testing.T) {
	ktp := 232312312123
	nama := "Adam"
	phone_number := "085725722483"
	alamat := "Kota Bandung"
	hasil:=module.InsertDaftarCamaba(module.MongoConn,"daftar_camaba", ktp, nama, phone_number, alamat)
	fmt.Println(hasil)
}

func TestInsertDaftarSekolah(t *testing.T) {
	kodesklh := 7
	nama := "SMA Negeri 9 Bandung"
	phone_number := "085725720202"
	alamat := "Kota Bandung"
	hasil:=module.InsertDaftarSekolah(module.MongoConn,"daftar_sekolah",kodesklh, nama, phone_number, alamat)
	fmt.Println(hasil)
}

func TestInsertDaftarJurusan(t *testing.T) {
	kodejurusan := "D4TI"
	nama := "Teknik Informatika"
	jenjang := "Diploma 4"
	hasil:=module.InsertDaftarJurusan(module.MongoConn,"daftar_jurusan",kodejurusan, nama, jenjang)
	fmt.Println(hasil)
}

// test getFunctionBy

func TestGetPendaftaranFromKTP(t *testing.T) {
	ktp := 320132321321
	pendaftar:=module.GetPendaftaranFromKTP(ktp,module.MongoConn, "pendaftaran_maba")
	fmt.Println(pendaftar)
}

func TestGetCamabaFromPhoneNumber(t *testing.T) {
	phonenumber := "085725722483"
	camaba:=module.GetCamabaFromPhoneNumber(phonenumber,module.MongoConn, "daftar_camaba")
	fmt.Println(camaba)
}

func TestGetDaftarSekolahFromKDSekolah(t *testing.T) {
	kdsekolah := 4
	daftar_sekolah:=module.GetDaftarSekolahFromKDSekolah(kdsekolah,module.MongoConn, "daftar_sekolah")
	fmt.Println(daftar_sekolah)
}

func TestGetJurusanFromKDJurusan(t *testing.T) {
	kdjurusan := "D3TI"
	daftar_jurusan:=module.GetJurusanFromKDJurusan(kdjurusan,module.MongoConn, "daftar_jurusan")
	fmt.Println(daftar_jurusan)
}

//test getFunctionAll

func TestGetAllPendaftaran(t *testing.T) {
	data := module.GetAllPendaftaran(module.MongoConn, "pendaftaran_maba")
	fmt.Println(data)
}

func TestGetAllJurusan(t *testing.T) {
	jurusan := module.GetAllJurusan(module.MongoConn, "daftar_jurusan")
	fmt.Println(jurusan)
}

func TestGetAllSekolah(t *testing.T) {
	sekolah := module.GetAllSekolah(module.MongoConn, "daftar_sekolah")
	fmt.Println(sekolah)
}

func TestGetAllCamaba(t *testing.T) {
	sekolah := module.GetAllCamaba(module.MongoConn, "daftar_camaba")
	fmt.Println(sekolah)
}
