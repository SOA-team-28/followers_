package repo

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type FollowerRepository struct {
	session neo4j.Session
}

func NewFollowerRepository(session neo4j.Session) *FollowerRepository {
	return &FollowerRepository{
		session: session,
	}
}

func (ur *FollowerRepository) CreateUser(follower Follower) error {
	_, err := ur.session.Run(
		"CREATE (u: {id: $id, username: $username, email: $email})",
		map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	)
	return err
}
