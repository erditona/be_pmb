package pmb1214031

import (
	"fmt"
	"testing"

	"github.com/erditona/be_pmb/model"
	"github.com/erditona/be_pmb/module"
)

func TestInsertPendaftaran(t *testing.T) {
	kdpendaftar := 1
	biodata := model.Camaba{
		Ktp : 320132321321,
		Nama : "Dito Adam",
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
	hasil:=module.InsertPendaftaran(kdpendaftar, biodata, asalsekolah, jurusan, jalur, alulbi, aljurusan)
	fmt.Println(hasil)
}

func TestInsertDaftarCamaba(t *testing.T) {
	ktp := 232312312312
	nama := "Dito"
	phone_number := "085725722483"
	alamat := "Kota Bandung"
	hasil:=module.InsertDaftarCamaba(ktp, nama, phone_number, alamat)
	fmt.Println(hasil)
}

func TestInsertDaftarSekolah(t *testing.T) {
	kodesklh := 4
	nama := "SMA Negeri 1 Bandung"
	phone_number := "085725722483"
	alamat := "Kota Bandung"
	hasil:=module.InsertDaftarSekolah(kodesklh, nama, phone_number, alamat)
	fmt.Println(hasil)
}

func TestInsertDaftarJurusan(t *testing.T) {
	kodejurusan := "D3TI"
	nama := "Teknik Informatika"
	jenjang := "Diploma 3"
	hasil:=module.InsertDaftarJurusan(kodejurusan, nama, jenjang)
	fmt.Println(hasil)
}

// test getFunction

func TestGetPendaftaranFromKTP(t *testing.T) {
	ktp := 320132321321
	pendaftar:=module.GetPendaftaranFromKTP(ktp)
	fmt.Println(pendaftar)
}

func TestGetCamabaFromPhoneNumber(t *testing.T) {
	phonenumber := "085725722483"
	camaba:=module.GetCamabaFromPhoneNumber(phonenumber)
	fmt.Println(camaba)
}

func TestGetDaftarSekolahFromKDSekolah(t *testing.T) {
	kdsekolah := 1
	daftar_sekolah:=module.GetDaftarSekolahFromKDSekolah(kdsekolah)
	fmt.Println(daftar_sekolah)
}

func TestGetJurusanFromKDJurusan(t *testing.T) {
	kdjurusan := "D3TI"
	daftar_jurusan:=module.GetJurusanFromKDJurusan(kdjurusan)
	fmt.Println(daftar_jurusan)
}



