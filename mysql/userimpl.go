package mysql

import "github.com/imjowend/patron-dao/models"

type UserImplMysql struct{}

func (dao UserImplMysql) Create(u *models.User) error {
	query := "INSERT INTO users (first_name, last_name, email) VALUES (?, ?, ?)"
	db := get()
	defer db.Close()

	statement, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	result, err := statement.Exec(u.FirstName, u.LastName, u.Email)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = int(id)
	return nil
}

func (dao UserImplMysql) GetAll() ([]models.User, error) {
	query := "SELECT * FROM users"
	users := []models.User{}
	db := get()
	defer db.Close()

	statement, err := db.Prepare(query)
	if err != nil {
		return users, err
	}
	defer statement.Close()
	rows, err := statement.Query()
	if err != nil {
		return users, err
	}
	for rows.Next() {
		var row models.User
		err := rows.Scan(&row.ID, &row.FirstName, &row.LastName, &row.Email)
		if err != nil {
			return users, err
		}
		users = append(users, row)
	}
	return users, nil
}
