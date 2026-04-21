package models

import "api-todo-aula2/db"

func Insert(todo Todo) (id int64, err error) {
	//abrir uma conexão com o banco
	conn, err := db.OpenConnection()
	if err != nil {
		return // sera tratado no handler o retorno do erro
	}
	defer conn.Close()

	sql := `INSERT INTO todos (title, description, completed) VALUES ($1, $2, $3) RETURNING id`
	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Completed).Scan(&id)

	// sera tratado no handler o retorno do erro
	return
}
