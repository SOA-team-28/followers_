package repo

import (
	"database-example/model"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type FollowerRepository struct {
	driver neo4j.Driver
}

func NewFollowerRepository(driver neo4j.Driver) *FollowerRepository {
	return &FollowerRepository{
		driver: driver,
	}
}

func (ur *FollowerRepository) CreateUser(follower model.Follower) error {
	session := ur.driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		_, err := transaction.Run(
			"CREATE (f:Follower {id: $id, followers: $followers, followable: $followable, followed: $followed})",
			map[string]interface{}{
				"id":         follower.Id,
				"followers":  follower.Followers,
				"followable": follower.Followable,
				"followed":   follower.Followed,
			},
		)
		return nil, err
	})

	return err
}
