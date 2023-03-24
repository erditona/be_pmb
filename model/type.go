package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pendaftaran struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	KDPendaftar  int                `bson:"kdpendaftar,omitempty" json:"kdpendaftar,omitempty"`
	Biodata      Camaba             `bson:"biodata,omitempty" json:"biodata,omitempty"`
	AsalSekolah  DaftarSekolah      `bson:"asalsekolah,omitempty" json:"asalsekolah,omitempty"`
	Jurusan  	 Jurusan  	        `bson:"jurusan,omitempty" json:"jurusan,omitempty"`
	Jalur  	     string  	        `bson:"jalur,omitempty" json:"jalur,omitempty"`
	AlUlbi  	 string    	    	`bson:"alulbi,omitempty" json:"alulbi,omitempty"`
	AlJurusan    string  	        `bson:"aljurusan,omitempty" json:"aljurusan,omitempty"`
	CreatedAt     primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
}

type Camaba struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Ktp          int                `bson:"ktp,omitempty" json:"ktp,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Address      string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
}

type DaftarSekolah struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	KDSekolah    int                `bson:"kdsekolah,omitempty" json:"kdsekolah,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Address      string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
}

type Jurusan struct {
	ID           primitive.ObjectID	`bson:"_id,omitempty" json:"_id,omitempty"`
	KDJurusan    string             `bson:"kdjurusan,omitempty" json:"kdjurusan,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Jenjang 	 string             `bson:"jenjang,omitempty" json:"jenjang,omitempty"`
}