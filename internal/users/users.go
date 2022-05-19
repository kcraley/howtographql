package users

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"

	database "github.com/kcraley/howtographql/internal/pkg/db/migrations/mysql"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

// Create creates the given user by inserting a new entry into the Users table.
func (user *User) Create() {
	stmt, err := database.Db.Prepare("INSERT INTO Users(Username, Password) VALUES(?,?)")
	if err != nil {
		log.Fatalf("failed preparing query to create new user: %v", err)
	}
	log.Printf("prepared statement: %s", stmt)

	hashedPass, err := HashPassword(user.Password)
	if err != nil {
		log.Fatalf("failed hashing the user's password: %v", err)
	}
	_, err = stmt.Exec(user.Username, hashedPass)
	if err != nil {
		log.Fatalf("failed creating user entry in the database: %v", err)
	}
}

// HashPassword hashes a given password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPassword hash compares a raw password with it's hashed values.
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GetUserIdByUsername checks if a user exists in the database given a username.
func GetUserIdByUsername(username string) (int, error) {
	stmt, err := database.Db.Prepare("select ID from Users WHERE Username = ?")
	if err != nil {
		log.Fatalf("failed preparing query to get user by id: %v", err)
	}
	row := stmt.QueryRow(username)

	var Id int
	err = row.Scan(&Id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Printf("unable to find user: %v", err)
		}
		return 0, err
	}
	return Id, nil
}
