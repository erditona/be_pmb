package pmb1214031

import (
	"errors"
	"fmt"
	"testing"

	"github.com/erditona/be_pmb/model"
	"github.com/erditona/be_pmb/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// test Insert

// func TestInsertPendaftaran(t *testing.T) {
// 	kdpendaftar := 3
// 	biodata := model.Camaba{
// 		Ktp : 3201323211222,
// 		Nama : "Dinan Adam",
// 		Phone_number : "085718177810",
// 		Address : "Parongpong, Kab. Bandung Barat",
// 	}
// 	asalsekolah := model.DaftarSekolah{
// 		KDSekolah : 01,
// 		Nama : "SMK Negeri 1 Cirebon",
// 		Phone_number : "085718172053",
// 		Address : "Jl.Perjuangan, Kota Cirebon",
// 	}
// 	jurusan := model.Jurusan{
// 		KDJurusan : "D4TI",
// 		Nama : "SMK Negeri 1 Cirebon",
// 		Jenjang : "Diploma 4",
// 	}
// 	jalur := "Rapot"
// 	alulbi := "Universitas Internasional"
// 	aljurusan := "Sedang trand"
// 	hasil:=module.InsertPendaftaran(module.MongoConn,"pendaftaran_maba",kdpendaftar, biodata, asalsekolah, jurusan, jalur, alulbi, aljurusan)
// 	fmt.Println(hasil)
// }

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


//GetAllFromId
//PendaftaranFromKDPendaftar
func TestGetPendaftaranFromKDPendaftar(t *testing.T) {
	kdpendaftar := 19062301
	biodata, err := module.GetPendaftaranFromKDPendaftar(kdpendaftar, module.MongoConn, "pendaftaran_maba")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			t.Fatalf("no data found for KDPendaftar %d", kdpendaftar)
		}
		t.Fatalf("error retrieving data for KDPendaftar %d: %v", kdpendaftar, err)
	}
	fmt.Println(biodata)
}

//PendaftaranFromID
func TestGetPendaftaranFromID(t *testing.T) {
	id := "642632f1ba550201c9bc41ed"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	biodata, err := module.GetPendaftaranFromID(objectID, module.MongoConn, "pendaftaran_maba")
	if err != nil {
		t.Fatalf("error calling GetPresensiFromID: %v", err)
	}
	fmt.Println(biodata)
}

//JurusanFromID
func TestGetJurusanFromID(t *testing.T) {
	id := "642bf18142f046a5b82bcca0"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	biodata, err := module.GetJurusanFromID(objectID, module.MongoConn, "daftar_jurusan")
	if err != nil {
		t.Fatalf("error calling GetPresensiFromID: %v", err)
	}
	fmt.Println(biodata)
}

//SekolahFromID
func TestGetSekolahFromID(t *testing.T) {
	id := "648149fee1c66741118b6c92"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	biodata, err := module.GetSekolahFromID(objectID, module.MongoConn, "daftar_sekolah")
	if err != nil {
		t.Fatalf("error calling GetPresensiFromID: %v", err)
	}
	fmt.Println(biodata)
}

//CamabaFromID
func TestGetCamabaFromID(t *testing.T) {
	id := "642bf8e042f046a5b82bccbf"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	biodata, err := module.GetCamabaFromID(objectID, module.MongoConn, "daftar_camaba")
	if err != nil {
		t.Fatalf("error calling GetPresensiFromID: %v", err)
	}
	fmt.Println(biodata)
}

//InsertV2
//Pendaftaran

func TestInsertPendaftaran(t *testing.T) {
	kdpendaftar := 3
	biodata := model.Camaba{
		Ktp:           3201323211222,
		Nama:          "TestInsertBaruBanget",
		Phone_number:  "085718177810",
		Address:       "Parongpong, Kab. Bandung Barat",
	}
	asalsekolah := model.DaftarSekolah{
		KDSekolah:     01,
		Nama:          "SMK Negeri 1 Cirebon",
		Phone_number:  "085718172053",
		Address:       "Jl.Perjuangan, Kota Cirebon",
	}
	jurusan := model.Jurusan{
		KDJurusan:     "D4TI",
		Nama:          "SMK Negeri 1 Cirebon",
		Jenjang:       "Diploma 4",
	}
	jalur := "Rapot"
	alulbi := "Universitas Internasional"
	aljurusan := "Sedang trand"
	insertedID, err := module.InsertPendaftaran(module.MongoConn, "pendaftaran_maba", kdpendaftar, biodata, asalsekolah, jurusan, jalur, alulbi, aljurusan)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan ID %s\n", insertedID.Hex())
}

// func TestInsertPendaftaran(t *testing.T) {
// 	kdpendaftar := 3
// 	biodata := model.Camaba{
// 		Ktp : 3201323211222,
// 		Nama : "TestInsert",
// 		Phone_number : "085718177810",
// 		Address : "Parongpong, Kab. Bandung Barat",
// 	}
// 	asalsekolah := model.DaftarSekolah{
// 		KDSekolah : 01,
// 		Nama : "SMK Negeri 1 Cirebon",
// 		Phone_number : "085718172053",
// 		Address : "Jl.Perjuangan, Kota Cirebon",
// 	}
// 	jurusan := model.Jurusan{
// 		KDJurusan : "D4TI",
// 		Nama : "SMK Negeri 1 Cirebon",
// 		Jenjang : "Diploma 4",
// 	}
// 	jalur := "Rapot"
// 	alulbi := "Universitas Internasional"
// 	aljurusan := "Sedang trand"
// 	insertedID, err := module.InsertPendaftaran(module.MongoConn,"pendaftaran_maba",kdpendaftar, biodata, asalsekolah, jurusan, jalur, alulbi, aljurusan)
// 	if err != nil {
// 		t.Errorf("Error inserting data: %v", err)
// 	}
// 	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
// }

//Camaba
func TestInsertCamaba(t *testing.T) {
	ktp := 232312312150
	nama := "NamaTest"
	phone_number := "085725722450"
	alamat := "Kota Test"
	insertedID, err := module.InsertSekolah(module.MongoConn,"daftar_camaba",ktp, nama, phone_number, alamat)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}
func TestInsertSekolah(t *testing.T) {
	kodesklh := 50
	nama := "sekolahTes"
	phone_number := "085718173250"
	alamat := "Kab Bandung"
	insertedID, err := module.InsertSekolah(module.MongoConn,"daftar_sekolah",kodesklh, nama, phone_number, alamat)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}
//jurusan
func TestInsertJurusan(t *testing.T) {
	kodejurusan := "D4Test3"
	nama := "D4Test3"
	jenjang := "Diploma 4"
	insertedID, err := module.InsertJurusan(module.MongoConn, "daftar_jurusan",kodejurusan, nama, jenjang)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

//update-delete

//Pendaftaran
func TestDeletePendaftaranByID(t *testing.T) {
	id := "64814992a8bfb03d29a6cb3b" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeletePendaftaranByID(objectID, module.MongoConn, "pendaftaran_maba")
	if err != nil {
		t.Fatalf("error calling DeleteSekolahByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetPendaftaranFromID
	_, err = module.GetPendaftaranFromID(objectID, module.MongoConn, "pendaftaran_maba")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}

//Sekolah
func TestDeleteSekolahByID(t *testing.T) {
	id := "6482fe79fb7f825ba14da7fd" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeleteSekolahByID(objectID, module.MongoConn, "daftar_sekolah")
	if err != nil {
		t.Fatalf("error calling DeleteSekolahByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetJurusanFromID
	_, err = module.GetSekolahFromID(objectID, module.MongoConn, "daftar_sekolah")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}

//Jurusan
func TestDeleteJurusanByID(t *testing.T) {
	id := "648bcd3e6a2b200a59c4aad4" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeleteJurusanByID(objectID, module.MongoConn, "daftar_jurusan")
	if err != nil {
		t.Fatalf("error calling DeleteJurusanByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetJurusanFromID
	_, err = module.GetJurusanFromID(objectID, module.MongoConn, "daftar_jurusan")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}



