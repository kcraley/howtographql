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
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address, UserID) VALUES(?,?, ?)")
	if err != nil {
		log.Fatalf("failed preparing query: %v", err)
	}
	res, err := stmt.Exec(link.Title, link.Address, link.User.ID)
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

func GetAll() []Link {
	stmt, err := database.Db.Prepare("select L.id, L.title, L.address, L.UserID, U.Username from Links L inner join Users U on L.UserID = U.ID")
	if err != nil {
		log.Fatalf("failed preparing query to get all links: %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatalf("failed preforming query to get all links: %v", err)
	}
	defer rows.Close()

	var links []Link
	var username string
	var id string
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address, &id, &username)
		if err != nil {
			log.Fatalf("failed scanning row: %v", err)
		}
		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return links
}
