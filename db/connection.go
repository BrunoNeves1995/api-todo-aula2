package db

import (
	"api-todo-aula2/configs"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

/*
Abri uma conexão com o banco de dados
*/
func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB() // pegando as configurações do banco

	stringConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database,
	)

	//Abrindo a string de conexão com o banco
	conn, err := sql.Open("postgres", stringConnection)
	if err != nil {
		panic("OpenConnection -> Erro ao abrir a conexão com o banco de dados" + err.Error())
	}

	//Testando a conexão do banco para ver se esta acesssivel
	err = conn.Ping()
	if err != nil {
		panic("OpenConnection -> Erro ao testar a conexão com o banco de dados" + err.Error())
	}

	return conn, err

}
