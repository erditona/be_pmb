package module

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aiteung/atdb"
	"github.com/erditona/be_pmb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "tes_db",
}

//connection
func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

// insert function

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertPendaftaran(kdpendaftar int, biodata model.Camaba, asalsekolah model.DaftarSekolah, jurusan model.Jurusan, jalur string, alulbi string, aljurusan string) (InsertedID interface{}) {
	var pendaftaran model.Pendaftaran
	pendaftaran.KDPendaftar = kdpendaftar
	pendaftaran.Biodata = biodata
	pendaftaran.AsalSekolah = asalsekolah
	pendaftaran.Jurusan = jurusan
	pendaftaran.Jalur = jalur
	pendaftaran.AlUlbi = alulbi
	pendaftaran.AlJurusan = aljurusan
	pendaftaran.CreatedAt = primitive.NewDateTimeFromTime(time.Now().UTC())
	return InsertOneDoc("db_pmb", "pendaftaran_maba", pendaftaran)
}

func InsertDaftarCamaba(ktp int, nama string, phone_number string, alamat string) (InsertedID interface{}) {
	var daftarCamaba model.Camaba
	daftarCamaba.Ktp = ktp
	daftarCamaba.Nama = nama
	daftarCamaba.Phone_number = phone_number
	daftarCamaba.Address = alamat
	return InsertOneDoc("db_pmb", "daftar_camaba", daftarCamaba)
}

func InsertDaftarSekolah(kodesklh int, nama string, phone_number string, alamat string) (InsertedID interface{}) {
	var daftarSekolah model.DaftarSekolah
	daftarSekolah.KDSekolah = kodesklh
	daftarSekolah.Nama = nama
	daftarSekolah.Phone_number = phone_number
	daftarSekolah.Address = alamat
	return InsertOneDoc("db_pmb", "daftar_sekolah", daftarSekolah)
}

func InsertDaftarJurusan(kodejurusan string, nama string, jenjang string) (InsertedID interface{}) {
	var daftarJurusan model.Jurusan
	daftarJurusan.KDJurusan = kodejurusan
	daftarJurusan.Nama = nama
	daftarJurusan.Jenjang = jenjang
	return InsertOneDoc("db_pmb", "daftar_jurusan", daftarJurusan)
}

// getfunction

func GetPendaftaranFromKTP(ktp int) (pendaftaran model.Pendaftaran) {
	Pendaftaran := MongoConnect("db_pmb").Collection("pendaftaran_maba")
	filter := bson.M{"biodata.ktp": ktp}
	fmt.Print("ktp");
	err := Pendaftaran.FindOne(context.TODO(), filter).Decode(&pendaftaran)
	if err != nil {
		fmt.Printf("getPendaftaranFromKTP: %v\n", err)
	}
	return pendaftaran
}

func GetCamabaFromPhoneNumber(phone_number string) (camaba model.Camaba) {
	Camaba := MongoConnect("db_pmb").Collection("daftar_camaba")
	filter := bson.M{"phone_number": phone_number}
	err := Camaba.FindOne(context.TODO(), filter).Decode(&camaba)
	if err != nil {
		fmt.Printf("getCamabaFromPhoneNumber: %v\n", err)
	}
	return camaba
}

func GetDaftarSekolahFromKDSekolah(kdsekolah int) (dfsekolah model.DaftarSekolah) {
	Dfsekolah := MongoConnect("db_pmb").Collection("daftar_sekolah")
	filter := bson.M{"kdsekolah": kdsekolah}
	err := Dfsekolah.FindOne(context.TODO(), filter).Decode(&dfsekolah)
	if err != nil {
		fmt.Printf("getCamabaFromPhoneNumber: %v\n", err)
	}
	return dfsekolah
}

func GetJurusanFromKDJurusan(kdjurusan string) (dfjurusan model.Jurusan) {
	Dfjurusan := MongoConnect("db_pmb").Collection("daftar_jurusan")
	filter := bson.M{"kdjurusan": kdjurusan}
	err := Dfjurusan.FindOne(context.TODO(), filter).Decode(&dfjurusan)
	if err != nil {
		fmt.Printf("getJurusanFromKDJurusan: %v\n", err)
	}
	return dfjurusan
}





