package store

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type User struct {
	Id         int
	Username   string
	password   string
	ProfilePic string
}

type UIUser struct {
	Id         int
	Username   string `json:"username"`
	Indx       string `json:"indx"`
	LastScore  string `json:"lastScore"`
	ProfilePic string `json:"profilePic"`
}

type UserStore interface {
	CreateUser(ctx context.Context, username, password string) (*User, error)
	GetUserById(ctx context.Context, id int) (*User, error)
	GetUIUserById(ctx context.Context, id int) (*UIUser, error)
	GetUserByUsername(ctx context.Context, id string) (*User, error)
	DeleteUser(ctx context.Context, id int) error
}

func (pg PostgresStore) CreateUser(ctx context.Context, username, password string) (*User, error) {
	username = strings.ToLower(username)
	username = strings.Trim(username, " ")

	query := `INSERT INTO users(username, password)
	VALUES($1, $2)
	RETURNING id;`

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Hashed pwd %s", string(bytes))

	var id int
	err = pg.db.QueryRowContext(ctx, query, username, string(bytes)).Scan(&id)

	user := &User{
		Id:       id,
		Username: username,
	}

	return user, err
}

func (pg PostgresStore) GetUIUserById(ctx context.Context, id int) (*UIUser, error) {
	u, err := pg.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	indx, err := pg.GetIndxByUserId(ctx, u.Id)
	if err != nil {
		return nil, err
	}
	lastRound, err := pg.GetLastRoundScoreByUserId(ctx, u.Id)
	if err != nil {
		return nil, err
	}

	c := cases.Title(language.Und)
	return &UIUser{
		Id:         u.Id,
		Username:   c.String(u.Username),
		Indx:       fmt.Sprintf("%.1f", indx),
		LastScore: strconv.Itoa(lastRound),
		ProfilePic: u.ProfilePic,
	}, nil
}

func (pg PostgresStore) GetUserById(ctx context.Context, id int) (*User, error) {
	query := `SELECT id, username, password, profile_pic
	FROM users WHERE id = $1 AND is_deleted = false;`

	user := new(User)
	err := pg.db.QueryRowContext(ctx, query, id).Scan(&user.Id, &user.Username, &user.password, &user.ProfilePic)

	return user, err
}

func (pg PostgresStore) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	query := `SELECT id, username, password, profile_pic
	FROM users WHERE username = $1 AND is_deleted = false;`

	user := new(User)
	err := pg.db.QueryRowContext(ctx, query, username).Scan(&user.Id, &user.Username, &user.password, &user.ProfilePic)

	return user, err
}

func (pg PostgresStore) DeleteUser(ctx context.Context, id int) error {
	query := `UPDATE users SET is_deleted = true WHERE id = $1;`

	_, err := pg.db.ExecContext(ctx, query, id)

	return err
}

func (u User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
	return err == nil
}

func (u UIUser) ScoreCardName() string {
	c := cases.Title(language.Und)
	if len(u.Username) > 4 {
		return c.String(u.Username[:4]) + "..."
	}
	return c.String(u.Username)
}
