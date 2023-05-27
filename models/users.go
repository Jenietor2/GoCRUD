package models

import "gomysql/db"

type User struct {
	Id       int64
	Username string
	Password string
	Email    string
}

type Users []User

const UserSchema string = `Create table users(
	id int(6) unsigned auto_increment primary key,
	username varchar(30) not null,
	password varchar(100) not null,
	email varchar(50),
	create_data timestamp default current_timestamp
)`

func newUser(username, password, email string) *User {
	user := &User{
		Username: username,
		Password: password,
		Email:    email,
	}

	return user
}

func CreateUser(username, password, email string) *User {
	user := newUser(username, password, email)
	user.insert()
	return user
}

func (user *User) insert() {
	sql := "insert users set username=?, password=?, email=?"
	result, _ := db.Exec(sql, user.Username, user.Password, user.Email)
	user.Id, _ = result.LastInsertId()
}

func ListUsers() Users {
	sql := "select id, username, password, email from users"
	users := Users{}
	rows, _ := db.Query(sql)

	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}

	return users
}

func FindById(id int) *User {
	user := newUser("", "", "")
	sql := "select id, username, password, email from users where id=?"
	row, _ := db.Query(sql, id)

	for row.Next() {
		row.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	}
	return user
}

func (user *User) update() {
	sql := "update users set username=?, password=?, email=? where id=?"

	db.Exec(sql, user.Username, user.Password, user.Email, user.Id)
}

func (user *User) Save() {
	if user.Id == 0 {
		user.insert()
	} else {
		user.update()
	}
}

func Delete(id int) {

	sql := "delete from users where id=?"
	db.Exec(sql, id)
}
