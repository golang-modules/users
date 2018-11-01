package main

import (
	"log"
	"users/models"
	"users/mysql"
)

func main() {

	user := models.User{}
	user.ID = 1
	user.FirstName = "FirstName Update"
	user.LastName = "LastName Update"
	user.Email = "first.updated@email.com"
	user.Password = "password"

	uc := mysql.New()
	// res, err := uc.Create(user)

	// log.Println(res)
	// log.Println(err)

	//res, err := uc.GetAll()

	//res, err := uc.Get(1)
	//res, err := uc.UpdatePassword(1, "newpasswored")
	res, err := uc.Delete(1)
	log.Println(res)
	log.Println(err)
}
