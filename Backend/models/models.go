package models

type Users struct {
	Id_user  int `json:"id_user" gorm:"primaryKey;autoIncrement:true"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Is_activ string //nanti outputnya di concert untuk jadi int
}

type Admins struct {
	Id_admin int `json:"id_admin" gorm:"primaryKey;autoIncrement:true"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Courses struct {
	Id_course       int `json:"id_course" gorm:"primaryKey;autoIncrement:true"`
	Nama_course     string `json:"nama_course" form:"nama_course"`
	Foto_course     string `json:"foto_course" form:"foto_course"`
	Deskipsi_course string	`json:"deskipsi_course" form:"deskipsi_course"`
	Id_Categoti        int `json:"id_categoti" form:"id_categoti"`
	Harga           float32 `json:"harga" form:"harga"`
	Is_active       string `json:"is_active" form:"is_active"`
}

type Categoris struct {
	Id_categori int 
	Nama_categori string
	Deskripsi_categori string
}