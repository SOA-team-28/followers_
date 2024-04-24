package service

import (
	"database-example/model"
	"database-example/repo"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type FollowerService struct {
	repo *repo.FollowerRepository
}

func NewFollowerService(driver neo4j.Driver) *FollowerService {
	repo := repo.NewFollowerRepository(driver)
	return &FollowerService{
		repo: repo,
	}
}

func (us *FollowerService) CreateUser(follower model.Follower) error {
	return us.repo.CreateUser(follower)
}
