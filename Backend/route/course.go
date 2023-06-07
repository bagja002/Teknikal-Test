package route

import (
	"learning/database"
	"learning/models"

	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"

	"log"

	"github.com/gofiber/fiber/v2"
)

// akses dari Role admin yaitu 1 dan Akses Untuk Use yaitu 2
const roleAdmin = 1
const roleuser = 2

func CreateCourse(c *fiber.Ctx) error {
	//validasi aut
	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(float64)

	//Cek apakah role tersebut merupakan role_admin
	if role != 1 {
		return c.JSON(fiber.Map{
			"pesan": "role tidak sesuai masukan role dengan benar ",
		})
	}
	//cek apakah admin terdaftar
	var admin models.Admins
	database.DB.Where("id_admin = ?", id_admin).Find(&admin)
	if admin.Id_admin == 0 {
		return c.JSON(fiber.Map{
			"pesan": "tidak di temukan admin",
		})
	}

	var course models.Courses
	if err := c.BodyParser(&course); err != nil {
		return c.JSON(fiber.Map{
			"pesan": "harap masukan data dengan benaar",
		})
	}
	CLOUDINARY_URL := "cloudinary://831922423943952:qAo73YkFEmTUXaeAM5QuOXx-eM4@dwr2ibrek"
	ctx := context.Background()
	fileHeader, _ := c.FormFile("foto_course")
	file, _ := fileHeader.Open()
	log.Println(fileHeader.Filename)

	cldServis, _ := cloudinary.NewFromURL(CLOUDINARY_URL)
	resp, _ := cldServis.Upload.Upload(ctx, file, uploader.UploadParams{})
	log.Println(resp.SecureURL)

	course.Foto_course = resp.SecureURL
	if course.Nama_course == "" {
		return c.JSON(fiber.Map{
			"pesan": "nama Course Tidak boleh Kosong",
		})
	}
	if course.Deskipsi_course == "" {
		return c.JSON(fiber.Map{
			"pesan": "Deskripsi Course Tidak boleh Kosong",
		})
	}
	if course.Id_Categoti == 0 {
		return c.JSON(fiber.Map{
			"pesan": "Categori Course Tidak boleh Kosong",
		})
	}
	if course.Is_active == "" {
		course.Is_active = "1"
	}
	database.DB.Create(&course)
	return c.JSON(fiber.Map{
		"data": course,
	})
	//cek apakah nama course kosong

}

func UpdateCourse(c *fiber.Ctx) error {
	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(float64)

	//Cek apakah role tersebut merupakan role_admin
	if role != 1 {
		return c.JSON(fiber.Map{
			"pesan": "role tidak sesuai masukan role dengan benar ",
		})
	}
	//cek apakah admin terdaftar
	var admin models.Admins
	database.DB.Where("id_admin = ?", id_admin).Find(&admin)
	if admin.Id_admin == 0 {
		return c.JSON(fiber.Map{
			"pesan": "tidak di temukan admin",
		})
	}
	//update id_course berdasarkan parameter
	var data map[string]string
	c.BodyParser(&data)
	var course models.Courses
	id_course := c.Params("id_course")
	//pengecekan apakah
	database.DB.Where("id_course = ?", id_course).Select("id_course").Find(&course)

	if course.Id_course == 0 {
		return c.JSON(fiber.Map{
			"pesam": "Maaf Course tidak di temukan",
		})
	}
	//data yang dapat di ubah yaitu nama_course, deskipsi Course, foto course
	update := models.Courses{
		Nama_course:     data["nama_course"],
		Deskipsi_course: data["deskripsi_course"],
		Is_active:       "1",
	}

	database.DB.Model(&course).Where("id_course = ?", id_course).Updates(update)
	return c.JSON(fiber.Map{
		"massage": "Data berhasil di ubah, mohon di cek kembali ya",
	})

}

// sukses
func GetOneCourse(c *fiber.Ctx) error {
	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(float64)

	//Cek apakah role tersebut merupakan role_admin
	if role != 2 {
		return c.JSON(fiber.Map{
			"pesan": "role tidak sesuai masukan role dengan benar ",
		})
	}
	//cek apakah admin terdaftar
	var user models.Users
	database.DB.Where("id_user = ?", id_admin).Find(&user)
	if user.Id_user == 0 {
		return c.JSON(fiber.Map{
			"pesan": "tidak di temukan admin",
		})
	}

	//amil id_course ari parameter

	id_course := c.Params("id_course")

	var course models.Courses

	database.DB.Where("id_course = ?", id_course).Find(&course)
	if course.Id_course == 0 {
		return c.JSON(fiber.Map{
			"pesan": "Course tidak tersedia",
		})
	} else {
		return c.JSON(fiber.Map{
			"status": "Success",
			"data":   course})
	}
}

// sukses
func GetAllCourses(c *fiber.Ctx) error {
	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(float64)

	//Cek apakah role tersebut merupakan role_admin
	if role != 2 {
		return c.JSON(fiber.Map{
			"pesan": "role tidak sesuai masukan role dengan benar ",
		})
	}
	//cek apakah admin terdaftar
	var user models.Users
	database.DB.Where("id_user = ?", id_admin).Find(&user)
	if user.Id_user == 0 {
		return c.JSON(fiber.Map{
			"pesan": "tidak di temukan user",
		})
	}

	var course []models.Courses
	sortby := c.Query("sortby")
	if sortby == "lowest_price" {
		database.DB.Where("harga > 0 AND is_active = 1").Order("harga ASC").Find(&course)
	} else if sortby == "highest_price" {
		database.DB.Where("harga > 0 AND is_active = 1").Order("harga DESC").Find(&course)
	} else if sortby == "free_price" {
		database.DB.Where("harga = 0 AND is_active = 1").Order("harga DESC").Find(&course)
	} else {
		database.DB.Where("is_active = 1").Find(&course)
	}

	return c.JSON(fiber.Map{
		"data":  course,
		"count": len(course),
	})

}

// sukses
func GetCategoriCourses(c *fiber.Ctx) error {

	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(float64)

	//Cek apakah role tersebut merupakan role_admin
	if role != 2 {
		return c.JSON(fiber.Map{
			"pesan": "role tidak sesuai masukan role dengan benar ",
		})
	}
	//cek apakah admin terdaftar
	var user models.Users
	database.DB.Where("id_user = ?", id_admin).Find(&user)
	if user.Id_user == 0 {
		return c.JSON(fiber.Map{
			"pesan": "tidak di temukan admin",
		})
	}

	//autentikasi jwt apakah sudah sesuai Untuk User
	id_categori := c.Params("id_categori")

	var categori models.Categoris

	database.DB.Where("id_categori = ?", id_categori).Find(&categori)
	if categori.Id_categori == 0 {
		return c.JSON(fiber.Map{
			"pesan": "Categori Course tidak tersedia",
		})
	}

	return c.JSON(fiber.Map{
		"pesan": categori,
	})
}

// sukses
func GetAllCategory(c *fiber.Ctx) error {

	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(float64)

	//Cek apakah role tersebut merupakan role_admin
	if role != 2 {
		return c.JSON(fiber.Map{
			"pesan": "role tidak sesuai masukan role dengan benar ",
		})
	}
	//cek apakah admin terdaftar
	var user models.Users
	database.DB.Where("id_user = ?", id_admin).Find(&user)
	if user.Id_user == 0 {
		return c.JSON(fiber.Map{
			"pesan": "tidak di temukan admin",
		})
	}

	var categori []models.Categoris

	result1 := database.DB.Find(&categori)
	if result1.Error != nil {
		return c.JSON(fiber.Map{
			"error": result1.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": categori,
	})
}

func DeleteCourse(c *fiber.Ctx) error {

	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(float64)

	//Cek apakah role tersebut merupakan role_admin
	if role != 1 {
		return c.JSON(fiber.Map{
			"pesan": "role tidak sesuai masukan role dengan benar ",
		})
	}
	//cek apakah admin terdaftar
	var admin models.Admins
	database.DB.Where("id_admin = ?", id_admin).Find(&admin)
	if admin.Id_admin == 0 {
		return c.JSON(fiber.Map{
			"pesan": "tidak di temukan admin",
		})
	}

	var data map[string]string
	c.BodyParser(&data)
	var course models.Courses
	id_course := c.Params("id_course")
	//pengecekan apakah
	database.DB.Where("id_course = ?", id_course).Select("id_course").Find(&course)

	if course.Id_course == 0 {
		return c.JSON(fiber.Map{
			"pesam": "Maaf Course tidak di temukan",
		})
	}
	//data yang dapat di ubah yaitu nama_course, deskipsi Course, foto course

	//jika 1 maka aktif dan jika 0 tidak aktif
	update := models.Courses{
		Is_active: data["is_active"],
	}

	database.DB.Model(&course).Where("id_course = ?", id_course).Updates(update)
	return c.JSON(fiber.Map{
		"massage": "Data berhasil di ubah",
	})
}

func SearchCourse(c *fiber.Ctx) error {
	// akses untuk admin dan user

	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(float64)

	//Cek apakah role tersebut merupakan role_admin
	if role != 2 {
		return c.JSON(fiber.Map{
			"pesan": "role tidak sesuai masukan role dengan benar ",
		})
	}
	//cek apakah admin terdaftar
	var user models.Users
	database.DB.Where("id_user = ?", id_admin).Find(&user)
	if user.Id_user == 0 {
		return c.JSON(fiber.Map{
			"pesan": "tidak di temukan admin",
		})
	}

	searchQuery := c.Query("search")

	var courses []models.Courses
	database.DB.Where("nama_course LIKE ?", "%"+searchQuery+"%").Find(&courses)

	return c.JSON(fiber.Map{
		"pesam": courses,
	})
}

/*
func UploadPhoto(c *fiber.Ctx) error {

	var course models.Course
	_ = c.BodyParser(&course)
	CLOUDINARY_URL:= "cloudinary://831922423943952:qAo73YkFEmTUXaeAM5QuOXx-eM4@dwr2ibrek"
	ctx := context.Background()
	fileHeader,_:=c.FormFile("foto_course")
	file,_ := fileHeader.Open()
	log.Println(fileHeader.Filename)

	cldServis,_:= cloudinary.NewFromURL(CLOUDINARY_URL)
	resp,_:=cldServis.Upload.Upload(ctx , file , uploader.UploadParams{})
	log.Println(resp.SecureURL)

	course.Foto_course = resp.SecureURL
	database.DB.Create(&course)
	return  c.JSON(fiber.Map{
		"data": course,
	})
}
**/
