package db

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var (
	DB   neo4j.Driver
	Port = "8082"
)

func InitDB() (neo4j.Driver, error) {
	uri := "neo4j+s://d26caf33.databases.neo4j.io:7"
	username := "neo4j"
	password := "BVCanBUL0zeBkauJBqd34ik0GmJwWvERXBPz-1Q_t18"

	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, err
	}

	DB = driver

	return DB, nil
}
