package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()
	//fmt.Println(db.CreateTable(models.UserSchema, "users"))
	//db.Ping()
	//models.CreateUser("Juanita", "123", "juana.nieto@gmail.com")
	//fmt.Println(user)
	//db.TruncateTable("users")
	/*users := models.ListUsers()
	fmt.Println(users)*/
	models.Delete(1)
	users := models.ListUsers()
	fmt.Println(users)
	/*user := models.FindById(4)
	user.Username = "Camila"
	user.Save()
	fmt.Println(user)*/
	db.Close()
}
