package main

import (
	"github.com/falk-dev/programe-como-uma-garota/tree/main/descomplicando-o-go/aula-04/crud/users"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := gorm.Open(sqlite.Open("meubanco.db"), &gorm.Config{})
	checkErr(err)

	db.AutoMigrate(&users.User{})

	id := users.Create("Mychelle")
	users.Update(id, "Falk")
	users.List()
	users.Delete(id)
	users.List()
}
