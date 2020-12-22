package links

import (
	database "github.com/JHaig343/graphql-server/internal/pkg/db/migrations/mysql"
	"github.com/JHaig343/graphql-server/internal/users"
	"log"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

// Save saves a new link (entered via GraphQL) into the MySQL DB
func (link Link) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title, Address) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(link.Title, link.Address)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Row Inserted!")
	return id
}

// GetAll retreives all links from the MySQL DB
func GetAll() []Link {
	stmt, err := database.Db.Prepare("SELECT id, address, title from Links")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var links []Link
	//
	for rows.Next() {
		var link Link
		// Scan matching values from the retrieved ro into the link struct
		err := rows.Scan(&link.ID, &link.Title, &link.Address)
		if err != nil {
			log.Fatal(err)
		}
		links = append(links, link)
	}
	// Something went wrong getting db rows
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return links
}
