package models

import "api-todo-aula2/db"

func Update(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE todos SET title=$1, description=$2, completed=$3 WHERE id=$4`, todo.Title, todo.Description, todo.Completed, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
