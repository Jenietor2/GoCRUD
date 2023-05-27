package main

import (
	"gomysql/db"
)

func main() {
	db.Connect()
	//fmt.Println(db.CreateTable(models.UserSchema, "users"))
	//db.Ping()
	db.TruncateTable("users")
	db.Close()
}
