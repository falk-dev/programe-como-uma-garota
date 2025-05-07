package main

import (
	"fmt"

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
	// db - instância da conexão com o banco de dados, que permite a manipulação de dados
	// err - possível erro de conexão
	// gorm.Open() - Open() é uma função da biblioteca gorm para abrir uma conexão com o banco de dados
	// essa função espera dois argumentos: dialector (qual é o banco e como acessar) e a configuração do gorm
	// sqlite.Open("meubanco.db") - informa que o banco de dados a ser usado é o SQLite e o nome do banco é 'meubanco.db'
	// &gorm.Config{} - struct de configuração opcional, podendo configurar coisas como: desabilitar logs, configuras callbacks etc.
	db, err := gorm.Open(sqlite.Open("meubanco.db"), &gorm.Config{})
	checkErr(err)

	// forma automática de manter o banco sincronizado com a struct que está em 'users/model.go'
	// que define a estrutura do meu banco de dados
	// o automigrate() não apaga colunas antigas nem dados antigos caso a estrutura do banco de dados mude,
	// apenas preenche com 'null' caso não seja preenchido
	// mas adiciona novas colunas
	db.AutoMigrate(&users.User{})

	id := users.Create(db, "Mychelle")

	users.Update(db, id, "Falk")

	listUsers := users.List(db)
	// 'range' permite que um array seja percorrido
	// quando invocado, retorna dois valores: índice e valor
	// ao usar '_', estamos ignorando o índice
	for _, user := range listUsers {
		fmt.Println(user.Name)
	}

	users.Delete(db, id)

	listUsers = users.List(db)
	for _, user := range listUsers {
		fmt.Println(user.Name)
	}
}
