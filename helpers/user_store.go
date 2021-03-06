package helpers

import (
	"database/sql"
	"fmt"
)

type Store interface {
	CreateUser(user *User) error
	GetUsers() ([]*FullUser, error)
	UpdateUser(user *UserWithoutTime) error
	DeleteUser(user *UserId) error
}

type DBstore struct {
	Db *sql.DB
}

func (store *DBstore) CreateUser(user *User) error {
	_, err := store.Db.Query("INSERT INTO users(first_name, last_name, age) VALUES($1,$2,$3)", user.FirstName, user.LastName, user.Age)
	return err
}

func (store *DBstore) UpdateUser(user *UserWithoutTime) error {
	_, err := store.Db.Query("update users set first_name = $2, last_name=$3, age=$4 where id=$1", user.Id, user.FirstName, user.LastName, user.Age)
	return err
}

func (store *DBstore) DeleteUser(user *UserId) error {
	_, err := store.Db.Query("delete from users where id = $1", user.Id)
	return err
}

func (store *DBstore) GetUsers() ([]*FullUser, error) {

	rows, err := store.Db.Query("SELECT id, first_name,last_name, age, creation_date from users")

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {

		if err := rows.Close(); err != nil {

			fmt.Println("err on closing rows:", err)
		}

	}(rows)

	var users []*FullUser

	for rows.Next() {
		user := &FullUser{}
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Age, &user.CreationDate); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil

}

var StoreToDB Store

func InitStore(s Store) {
	StoreToDB = s

}
