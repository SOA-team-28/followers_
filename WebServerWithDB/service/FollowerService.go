package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"

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

func (us *FollowerService) GetById(id int) (model.Follower, error) {
	follower, err := us.repo.GetById(id)
	if err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("Follower service retrieved:", follower)
	}
	return follower, err
}

func (us *FollowerService) UpdateUser(existingUserID int, newFollowerID int) error {
	err := us.repo.UpdateUser(existingUserID, newFollowerID)
	if err != nil {
		// Ovde možete dodati dodatnu obradu greške ako je potrebno
		return err
	}
	return nil
}
