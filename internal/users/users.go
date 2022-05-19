package users

import (
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
