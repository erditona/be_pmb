package module

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/aiteung/atdb"
	"github.com/erditona/be_pmb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "tes_db_pmb",
}

var MongoConn = atdb.MongoConnect(MongoInfo)


// insert function

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

// func InsertPendaftaran(db *mongo.Database, col string, kdpendaftar int, biodata model.Camaba, asalsekolah model.DaftarSekolah, jurusan model.Jurusan, jalur string, alulbi string, aljurusan string) (InsertedID interface{}) {
// 	var pendaftaran model.Pendaftaran
// 	pendaftaran.KDPendaftar = kdpendaftar
// 	pendaftaran.Biodata = biodata
// 	pendaftaran.AsalSekolah = asalsekolah
// 	pendaftaran.Jurusan = jurusan
// 	pendaftaran.Jalur = jalur
// 	pendaftaran.AlUlbi = alulbi
// 	pendaftaran.AlJurusan = aljurusan
// 	pendaftaran.CreatedAt = primitive.NewDateTimeFromTime(time.Now().UTC())
// 	return InsertOneDoc(db, col, pendaftaran)
// }
func InsertPendaftaran(db *mongo.Database, col string, kdpendaftar int, biodata model.Camaba, asalsekolah model.DaftarSekolah, jurusan model.Jurusan, jalur string, alulbi string, aljurusan string) (insertedID primitive.ObjectID, err error) {
	pendaftaran := bson.M{
		"kdpendaftar": kdpendaftar,
		"biodata":     biodata,
		"asalsekolah": asalsekolah,
		"jurusan":     jurusan,
		"jalur":       jalur,
		"alulbi":      alulbi,
		"aljurusan":   aljurusan,
		"created_at":  primitive.NewDateTimeFromTime(time.Now().UTC()),
	}
	result, err := db.Collection(col).InsertOne(context.Background(), pendaftaran)
	if err != nil {
		fmt.Printf("InsertPendaftaran: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func InsertDaftarCamaba(db *mongo.Database,col string,ktp int, nama string, phone_number string, alamat string) (InsertedID interface{}) {
	var daftarCamaba model.Camaba
	daftarCamaba.Ktp = ktp
	daftarCamaba.Nama = nama
	daftarCamaba.Phone_number = phone_number
	daftarCamaba.Address = alamat
	return InsertOneDoc(db, col, daftarCamaba)
}
func InsertCamaba(db *mongo.Database,col string,ktp int, nama string, phone_number string, alamat string) (insertedID primitive.ObjectID, err error) {
	camaba := bson.M{
		"ktp":          ktp,
		"nama":         nama,
		"phone_number": phone_number,
		"alamat":       alamat,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), camaba)
	if err != nil {
		fmt.Printf("InsertCamaba: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func InsertDaftarSekolah(db *mongo.Database,col string,kodesklh int, nama string, phone_number string, alamat string) (InsertedID interface{}) {
	var daftarSekolah model.DaftarSekolah
	daftarSekolah.KDSekolah = kodesklh
	daftarSekolah.Nama = nama
	daftarSekolah.Phone_number = phone_number
	daftarSekolah.Address = alamat
	return InsertOneDoc(db, col, daftarSekolah)
}
func InsertSekolah(db *mongo.Database,col string,kodesklh int, nama string, phone_number string, alamat string) (insertedID primitive.ObjectID, err error) {
	sekolah := bson.M{
		"kdsekolah":    kodesklh,
		"nama":         nama,
		"phone_number": phone_number,
		"alamat":       alamat,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), sekolah)
	if err != nil {
		fmt.Printf("InsertSekolah: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func InsertDaftarJurusan(db *mongo.Database,col string,kodejurusan string, nama string, jenjang string) (InsertedID interface{}) {
	var daftarJurusan model.Jurusan
	daftarJurusan.KDJurusan = kodejurusan
	daftarJurusan.Nama = nama
	daftarJurusan.Jenjang = jenjang
	return InsertOneDoc(db, col, daftarJurusan)
}
func InsertJurusan(db *mongo.Database,col string,kodejurusan string, nama string, jenjang string) (insertedID primitive.ObjectID, err error) {
	jurusan := bson.M{
		"kdjurusan": kodejurusan,
		"nama":      nama,
		"jenjang":   jenjang,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), jurusan)
	if err != nil {
		fmt.Printf("InsertJurusan: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

// getfunction

func GetPendaftaranFromKTP(ktp int, db *mongo.Database, col string) (pendaftaran model.Pendaftaran) {
	Pendaftaran := db.Collection(col)
	filter := bson.M{"biodata.ktp": ktp}
	fmt.Print("ktp");
	err := Pendaftaran.FindOne(context.TODO(), filter).Decode(&pendaftaran)
	if err != nil {
		fmt.Printf("getPendaftaranFromKTP: %v\n", err)
	}
	return pendaftaran
}

func GetCamabaFromPhoneNumber(phone_number string, db *mongo.Database, col string) (camaba model.Camaba) {
	Camaba := db.Collection(col)
	filter := bson.M{"phone_number": phone_number}
	err := Camaba.FindOne(context.TODO(), filter).Decode(&camaba)
	if err != nil {
		fmt.Printf("getCamabaFromPhoneNumber: %v\n", err)
	}
	return camaba
}

func GetDaftarSekolahFromKDSekolah(kdsekolah int, db *mongo.Database, col string) (dfsekolah model.DaftarSekolah) {
	Dfsekolah := db.Collection(col)
	filter := bson.M{"kdsekolah": kdsekolah}
	err := Dfsekolah.FindOne(context.TODO(), filter).Decode(&dfsekolah)
	if err != nil { 
		fmt.Printf("getCamabaFromPhoneNumber: %v\n", err)
	}
	return dfsekolah
}

func GetJurusanFromKDJurusan(kdjurusan string, db *mongo.Database, col string) (dfjurusan model.Jurusan) {
	Dfjurusan := db.Collection(col)
	filter := bson.M{"kdjurusan": kdjurusan}
	err := Dfjurusan.FindOne(context.TODO(), filter).Decode(&dfjurusan)
	if err != nil {
		fmt.Printf("getJurusanFromKDJurusan: %v\n", err)
	}
	return dfjurusan
}

//GetFunctionAll

func GetAllPendaftaran(db *mongo.Database, col string) (pendaftaran []model.Pendaftaran) {
	data_pendaftaran := db.Collection(col)
	filter := bson.M{}
	cursor, err := data_pendaftaran.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &pendaftaran)
	if err != nil {
		fmt.Println(err)
	}
	return pendaftaran
}

func GetAllJurusan(db *mongo.Database, col string) (jurusan []model.Jurusan) {
	data_jurusan := db.Collection(col)
	filter := bson.D{}
	// var results []jurusan
	cur, err := data_jurusan.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetAllJurusan: %v\n", err)
	}
	err = cur.All(context.TODO(), &jurusan)
	if err != nil {
		fmt.Println(err)
	}
	return jurusan
}

func GetAllSekolah(db *mongo.Database, col string) (sekolah []model.DaftarSekolah) {
	daftar_sekolah := db.Collection(col)
	filter := bson.D{}
	// var results []DaftarSekolah
	cur, err := daftar_sekolah.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetAllSekolah: %v\n", err)
	}
	err = cur.All(context.TODO(), &sekolah)
	if err != nil {
		fmt.Println(err)
	}
	return sekolah
}

func GetAllCamaba(db *mongo.Database, col string) (camaba []model.Camaba) {
	daftar_camaba := db.Collection(col)
	filter := bson.D{}
	// var results []DaftarCamaba
	cur, err := daftar_camaba.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetAllCamaba: %v\n", err)
	}
	err = cur.All(context.TODO(), &camaba)
	if err != nil {
		fmt.Println(err)
	}
	return camaba
}


//GetAllFromId
//PendaftaranFromId
func GetPendaftaranFromID(_id primitive.ObjectID, db *mongo.Database, col string) (staf model.Pendaftaran, errs error) {
	pendaftar := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := pendaftar.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return staf, fmt.Errorf("no data found for ID %s", _id)
		}
		return staf, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return staf, nil
}

//JurusanFromId
func GetJurusanFromID(_id primitive.ObjectID, db *mongo.Database, col string) (staf model.Jurusan, errs error) {
	jurusan := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := jurusan.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return staf, fmt.Errorf("no data found for ID %s", _id)
		}
		return staf, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return staf, nil
}

//SekolahFromID
func GetSekolahFromID(_id primitive.ObjectID, db *mongo.Database, col string) (staf model.DaftarSekolah, errs error) {
	sekolah := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := sekolah.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return staf, fmt.Errorf("no data found for ID %s", _id)
		}
		return staf, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return staf, nil
}

//camabaFromID
func GetCamabaFromID(_id primitive.ObjectID, db *mongo.Database, col string) (staf model.Camaba, errs error) {
	camaba := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := camaba.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return staf, fmt.Errorf("no data found for ID %s", _id)
		}
		return staf, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return staf, nil
}


//Update-Delete

//Jurusan
func UpdateJurusan(db *mongo.Database,col string,id primitive.ObjectID, kodejurusan string, nama string, jenjang string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"kdjurusan": kodejurusan,
			"nama":      nama,
			"jenjang":   jenjang,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateJurusan: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

func DeleteJurusanByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	jurusan := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := jurusan.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}





