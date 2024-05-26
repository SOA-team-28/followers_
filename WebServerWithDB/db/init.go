package db

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var (
	DB   neo4j.Driver
	Port = "8082"
)

func InitDB() (neo4j.Driver, error) {
	uri := "neo4j+s://42362ffe.databases.neo4j.io"
	username := "neo4j"
	password := "DiAWDt3spN6nBL1ckT5_zfiSNOy-BOnu0Cad-l8V_CA"

	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, err
	}

	DB = driver

	return DB, nil
}
