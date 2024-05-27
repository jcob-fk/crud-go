package models

import (
	"crud/config"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Usuario struct {
	Id       int
	Nombre   string
	Email    string
	Password string
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateUser(nombre string, correo string, password string) {
	conexion := config.Conexion()
	query, err := conexion.Prepare("INSERT INTO usuario (nombre, email, password, activo) VALUE (?, ?, ?, ?)")
	if err != nil {
		fmt.Print("error al crear usuario")
	}
	securePassword, err := HashPassword(password)
	if err != nil {
		fmt.Print("error al cifrar contraseña")
	}
	query.Exec(nombre, correo, securePassword, 1)
	conexion.Close()
}

func ReadUser(id string) Usuario {
	conexion := config.Conexion()
	row := conexion.QueryRow("SELECT id, nombre, email, password FROM usuario WHERE activo = 1 AND id = ?", id)

	conexion.Close()

	var usuario Usuario
	row.Scan(&usuario.Id, &usuario.Nombre, &usuario.Email, &usuario.Password)

	return usuario
}

func ReadUsers() []Usuario {
	conexion := config.Conexion()
	rows, err := conexion.Query("SELECT id, nombre, email FROM usuario WHERE activo = 1")
	fmt.Println("david el traga pitos")
	if err != nil {
		fmt.Print("error al leer usuarios")
		return nil
	}
	conexion.Close()

	var usuarios []Usuario
	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.Id, &usuario.Nombre, &usuario.Email)
		usuarios = append(usuarios, usuario)
	}

	return usuarios
}

func DeleteUser(id string) {
	conexion := config.Conexion()
	query, err := conexion.Prepare("UPDATE usuario SET activo = 0 WHERE id = ?")
	if err != nil {
		fmt.Print("error al eliminar usuario")

	}

	query.Exec(id)
	conexion.Close()
}

func UpdateUser(id, nombre, email, password string) {
	conexion := config.Conexion()
	query, err := conexion.Prepare("UPDATE usuario SET nombre = ?, email = ?, password = ? WHERE id = ?")
	if err != nil {
		fmt.Print("error al actualizar usuario")
	}
	securePassword, err := HashPassword(password)
	if err != nil {
		fmt.Print("error al cifrar contraseña")
	}
	query.Exec(nombre, email, securePassword, id)
	conexion.Close()
}
