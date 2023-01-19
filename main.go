package main

import (
	"log"
	"starco/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/starco?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = "Holy"
	userInput.Email = "holyshit@gmail.com"
	userInput.Occupation = "developer"
	userInput.Password = "holy123"

	userService.RegisterUser(userInput)

}
