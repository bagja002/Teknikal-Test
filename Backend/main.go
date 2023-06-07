package main

import (
	//"learning/database"
	"log"
	//get ini
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"learning/database"
	"learning/route"
	"github.com/gofiber/fiber/v2/middleware/recover"

)

func main(){
	//instalasi App Fiber dan Middleware
	database.Connect()
	app := fiber.New()
	app.Use(
		cors.New(cors.Config{
		AllowCredentials: true,
		},
	), 
	logger.New(),
	recover.New())
	//goup route
	admin := app.Group("/admin")
	user := app.Group("/users")
	
	
	//user route
	user.Post("/register", route.RegisterUser)
	user.Post("/login", route.LoginUser)
	user.Get("/getcourse/:id_course", route.ExtractTokenData(),route.GetOneCourse)
	user.Get("/getcourses",route.ExtractTokenData() ,route.SearchCourse)
	user.Get("/getAllcourses",route.ExtractTokenData(), route.GetAllCourses)
	
	user.Get("/GetCategoryCourse/:id_categori",route.ExtractTokenData() ,route.GetCategoriCourses)
	user.Get("/getAllcategori",route.ExtractTokenData(),route.GetAllCategory)


	//admin route
	admin.Post("/courses", route.ExtractTokenData(),route.CreateCourse)
	admin.Post("/register", route.CreateAdmin)
	admin.Post("/loginadmin", route.AuthAdmin)
	admin.Delete("/deleteuser/:id_user", route.ExtractTokenData(),route.DeletUser)
	admin.Put("/course/:id_course", route.ExtractTokenData(),route.UpdateCourse)
	admin.Delete("/course/:id_course", route.ExtractTokenData(),route.DeleteCourse)
	admin.Get("/getstatistik", route.ExtractTokenData() ,route.GetStatistik)






	//listening port
	log.Fatal(app.Listen(":9000"))
}