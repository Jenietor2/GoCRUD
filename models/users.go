package models

type User struct {
	Id       int
	Username string
	Password string
	Email    string
}

const UserSchema string = `Create table users(
	id int(6) unsigned auto_increment primary key,
	username varchar(30) not null,
	password varchar(100) not null,
	email varchar(50),
	create_data timestamp default current_timestamp
)`
