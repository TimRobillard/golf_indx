package store

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type User struct {
	Id       int
	Username string
	password string
}

type UserStorage interface {
	CreateUser(string, string) (*User, error)
	GetUserById(int) (*User, error)
	GetUserByUsername(string) (*User, error)
	DeleteUser(int) error
}

func (pg PostgresStore) CreateUser(username, password string) (*User, error) {
	username = strings.ToLower(username)
	username = strings.Trim(username, " ")

	query := `INSERT INTO users(username, password)
	VALUE($1, $2)
	RETURNING id;`

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}

	var id int
	err = pg.db.QueryRow(query, username, string(bytes)).Scan(&id)

	user := &User{
		Id:       id,
		Username: username,
	}

	return user, nil
}

func (pg PostgresStore) GetUserById(id int) (*User, error) {
	query := `SELECT id, username, password
	FROM users WHERE id = $1 AND is_deleted = false;`

	user := new(User)
	err := pg.db.QueryRow(query, id).Scan(&user.Id, &user.Username, &user.password)

	return user, err
}

func (pg PostgresStore) GetUserByUsername(username string) (*User, error) {
	query := `SELECT id, username, password
	FROM users WHERE username = $1 AND is_deleted = false;`

	user := new(User)
	err := pg.db.QueryRow(query, username).Scan(&user.Id, &user.Username, &user.password)

	return user, err
}

func (pg PostgresStore) DeleteUser(id int) error {
	query := `UPDATE users SET is_deleted = true WHERE id = $1;`

	_, err := pg.db.Exec(query, id)

	return err
}

func (u User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
	return err == nil
}

func (u User) ScoreCardName() string {
	c := cases.Title(language.Und)
	if len(u.Username) > 4 {
		return c.String(u.Username[:4]) + "..."
	}
	return c.String(u.Username)
}
