package main

import (
	"bwastartup/auth"
	"bwastartup/handler"
	"bwastartup/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// dbDriver := "mysql"
	// dbUser := "root"
	// dbPass := ""
	// dbName := "bwastartup"
	// db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	// if err != nil {
	// 	panic(err.Error())
	// }

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()

	api := router.Group("api/v1/")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.GET("/users", userHandler.UsersAll)
	api.POST("/avatar", userHandler.UploadAvatar)

	router.Run()

	// userInput := user.RegisterUserInput{}

	// userInput.Name = "Rafly Fahrezi"
	// userInput.Email = "AM@gmail.com"
	// userInput.Occupation = "anak punk"
	// userInput.Password = "123456"

	// userService.RegisterUser(userInput)

	// user := user.User{
	// 	Name: "Test Simpan",
	// }

	// userRepository.Save(user)

	// fmt.Println("Koneksi Berhasil")

	// var users []user.User
	// //length := len(users)

	// //fmt.Println(length)

	// db.Find(&users)

	// //length = len(users)

	// //fmt.Println(length)

	// for _, user := range users {
	// 	fmt.Println(user.Name)
	// 	fmt.Println(user.Email)
	// 	fmt.Println("=================")
	// }

	// router := gin.Default()
	// router.GET("/hendler", hendler)
	// router.Run()

}

// func hendler(c *gin.Context) {

// 	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	//fmt.Println("Koneksi Berhasil")

// 	var users []user.User
// 	//length := len(users)

// 	//fmt.Println(length)

// 	db.Find(&users)

// 	c.JSON(http.StatusOK, users)

// }
