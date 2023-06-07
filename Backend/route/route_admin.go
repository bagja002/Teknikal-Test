package route

import (
	
	"learning/database"
	"learning/models"

	"time"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)
//akses dari Role admin yaitu 1 dan Akses Untuk Use yaitu 2 
func CreateAdmin(c *fiber.Ctx)error{

	var data map[string]string
	//var dataint map[string]int
	
	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"pesan": "harap masukan data dengan benaar",
		})
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	var admin models.Admins
	admins := models.Admins{
		Nama: data["nama"],
		Email: data["email"],
		Password: string(password),
	}
	if admins.Email == "" {
		return c.JSON(fiber.Map{
			"pesan": "Email Tidak Boleh kosong ",
		})
	}
	if admins.Password == "" {
		return c.JSON(fiber.Map{
			"pesan": "Password tidak boleh kosong ",
		})
	}
	
	database.DB.Where("email = ?", admins.Email).Select("email").Find(&admin)
	if admins.Email == admin.Email {
		return c.JSON(fiber.Map{
			"status":  "invalid",
			"Massage": "Email  Telah Digunakan Harap Masukan Username yang lain",
		})
	} else {
		database.DB.Create(&admins)
		return c.JSON(fiber.Map{
			"status":  "Success",
			"massage": "Admin Telah di buat",
		})
	}


}





func AuthAdmin(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.JSON(fiber.Map{
			"pesan": err.Error(),
		})
	}

	var admin models.Admins
	database.DB.Where("email = ?", data["email"]).First(&admin)
	if admin.Id_admin == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"pesan": "Admin tidak di temukan",
		})
	}
	
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect password!",
			})}
	//Di bikin JWT Authentication buat admin sma seperi user
	claims := jwt.MapClaims{
		"name":     admin.Nama,
		"id_admin": admin.Id_admin,
		"role":     roleAdmin,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"pesan": t,
	})
}

// menghapus Softdelete User
func DeletUser(c *fiber.Ctx) error {

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
	var user models.Users
	id_user := c.Params("id_user")
	//pengecekan apakah
	database.DB.Where("id_user = ?", id_user).Select("id_user").Find(&user)

	if user.Id_user == 0 {
		return c.JSON(fiber.Map{
			"pesam": "Maaf User tidak di temukan",
		})
	}

	//jika 1 maka aktif dan jika 0 tidak aktif
	update := models.Users{
		Is_activ: data["is_active"],
	}

	database.DB.Model(&user).Where("id_user = ?", id_user).Updates(update)
	return c.JSON(fiber.Map{
		"massage": "Data Telah di apus",
	})
}



func Validasi1(c *fiber.Ctx) error {
	// Get token from Authorization header

	// Extract data from claims
	name, _ := c.Locals("name").(string)
	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(string)
	if role != "1" {
		return c.JSON(fiber.Map{
			"pesan": "role tidak sesuai masukan role dengan benar ",
		})
	}

	// Return data
	return c.JSON(fiber.Map{
		"name":     name,
		"id_admin": id_admin,
		"role":     role,
	})

}
func GetStatistik(c *fiber.Ctx) error {
	//Validasi JWT Authentication
	
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

	type Statsitik struct {
		Total_user int
		Total_course int
		Total_frecourse int
	}



	//Total User 
	var user []models.Users
	result1 := database.DB.Where("is_activ = 1").Find(&user)
	if result1.Error != nil {
		return c.JSON(fiber.Map{
			"error": result1.Error.Error(),
		})
	}


	//Total Course 
	var course []models.Courses
	result2 := database.DB.Where("is_active = 1").Find(&course)
	if result2.Error != nil {
		return c.JSON(fiber.Map{
			"error": result1.Error.Error(),
		})
	}

	//total Free Course
	newCourse := make([]models.Courses, len(course))
	copy(newCourse, course)
	result3 := database.DB.Where("harga = 0").Find(&newCourse)
	if result3.Error != nil {
		return c.JSON(fiber.Map{
			"error": result1.Error.Error(),
		})
	}

	
	hasil := Statsitik{
		Total_user: len(user),
		Total_course: len(course),
		Total_frecourse: len(newCourse),
	}	
	return c.JSON(fiber.Map{
		"status":"Sucses",
		"data": hasil,
	})
}

func ExtractTokenData() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get Authorization header value
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		// Check if Authorization header value is in the correct format
		if len(authHeader) < 8 || authHeader[:7] != "Bearer " {
			return c.JSON(fiber.Map{
				"status":"Invalid",
			})
		}

		// Extract token string
		tokenString := authHeader[7:]

		// Parse token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verify signing method and secret key
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})

		if err != nil {
			// Token not valid, handle error
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Get claims from token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			// Claims not valid, handle error
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Extract data from claims and add it to the context
		c.Locals("name", claims["name"].(string))
		c.Locals("id_admin", int(claims["id_admin"].(float64)))
		c.Locals("role", claims["role"].(float64))

		// Continue with next middleware/handler
		return c.Next()
	}
}
