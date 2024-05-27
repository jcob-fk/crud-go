package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Conexion() *sql.DB {
	db, err1 := sql.Open("mysql", os.Getenv("DB_URL"))
	err2 := db.Ping()
	if err1 == nil && err2 == nil {
		fmt.Print("Conexion a DB existosa...")
	} else {
		fmt.Print("Error de conexion")
	}
	return db
}
