package db

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var (
	DB   neo4j.Driver
	Port = "8082"
)

func InitDB() (neo4j.Driver, error) {
	uri := "neo4j+s://3091c5e3.databases.neo4j.io"
	username := "neo4j"
	password := "QEhW3bajhHfMd5gMG0Ho9oX_OTYx3QUNwvMHkcmcRtw"

	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, err
	}

	DB = driver

	return DB, nil
}
