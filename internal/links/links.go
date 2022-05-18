package links

import (
	"log"

	database "github.com/kcraley/howtographql/internal/pkg/db/migrations/mysql"
	"github.com/kcraley/howtographql/internal/users"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func (link Link) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address) VALUES(?,?)")
	if err != nil {
		log.Fatalf("failed preparing query: %v", err)
	}
	res, err := stmt.Exec(link.Title, link.Address)
	if err != nil {
		log.Fatalf("failed executing query: %v", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatalf("failed getting last insert id: %v", err)
	}
	log.Print("row inserted")
	return id
}
