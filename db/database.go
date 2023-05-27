package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//username:password@tcp(localhost:3306)/database

const url = "root:Admin123*@tcp(localhost:3306)/goweb_db"

var db *sql.DB

func Connect() {
	connection, err := sql.Open("mysql", url)

	if err != nil {
		panic(err)
	}
	fmt.Println("Conexion exitosa")
	db = connection
}

func Close() {
	db.Close()
}

// Verificar la conexión
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func existsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := Query(sql)

	if err != nil {
		fmt.Println("Error:", err)
	}

	return rows.Next()
}

// crea tabla
func CreateTable(schema string, nameTable string) string {
	if !existsTable(nameTable) {
		_, err := Exec(schema)

		if err != nil {
			fmt.Println(err)
		}
		return "Creación exitosa"
	}

	return "No fue porsible crear la tabla"
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)

	if err != nil {
		fmt.Println(err)
	}

	return result, err
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {

	rows, err := db.Query(query, args...)

	if err != nil {
		fmt.Println(err)
	}

	return rows, err
}

func TruncateTable(tableName string){

	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)
}
