package route

import (

	"learning/database"
	"learning/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)



// registrasi user
func RegisterUser(c *fiber.Ctx) error {

	var data map[string]string
	//var dataint map[string]int

	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"pesan": "harap masukan data dengan benaar",
		})
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	user := models.Users{
		Nama:     data["nama"],
		Email:    data["email"],
		Password: string(password),
		Is_activ: "1",
	}

	if user.Email == "" {
		return c.JSON(fiber.Map{
			"pesan": "Email Tidak Boleh kosong ",
		})
	}
	if user.Password == "" {
		return c.JSON(fiber.Map{
			"pesan": "Password tidak boleh kosong ",
		})
	}
	var email models.Users
	database.DB.Where("email = ?", user.Email).Select("email").Find(&email)
	if user.Email == email.Email {
		return c.JSON(fiber.Map{
			"status":  "invalid",
			"Massage": "Email  Telah Digunakan Harap Masukan Username yang lain",
		})
	} else {
		database.DB.Create(&user)
		return c.JSON(fiber.Map{
			"status":  "Success",
			"massage": "Selamat datang dan selamat belajar",
		})
	}
}

func LoginUser(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var users models.Users

	database.DB.Where("email = ?", data["email"]).First(&users)

	if users.Id_user == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"pesan": "User tidak di temukan",
		})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect password!",
		})

	}else {
		//Di bikin JWT Authentication buat admin sma seperi user
		claims := jwt.MapClaims{
			"name":     users.Nama,
			"id_admin": users.Id_user,
			"role":     roleuser,
			"exp":      time.Now().Add(time.Hour * 72).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{
			"token": t,
		})

	}

}
