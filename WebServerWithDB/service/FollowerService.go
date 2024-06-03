package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
	"log"

	saga "database-example/saga"

	events "database-example/saga/check_login"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type FollowerService struct {
	repo              *repo.FollowerRepository
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewFollowerService(driver neo4j.Driver, publisher saga.Publisher, subscriber saga.Subscriber) (*FollowerService, error) {
	repo := repo.NewFollowerRepository(driver)
	u := &FollowerService{
		repo:              repo,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	log.Println("subsrciber u handleru:", u.commandSubscriber)
	err := u.commandSubscriber.Subscribe(u.CheckLoginAvailability)

	if err != nil {
		log.Println("Error subscribing to commands:", err)
		return nil, err
	}
	return u, nil
}

func (us *FollowerService) CreateFollower(follower model.Follower) error {
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

func (us *FollowerService) UpdateFollower(existingUserID int, newFollowerID int) error {
	err := us.repo.UpdateUser(existingUserID, newFollowerID)
	if err != nil {
		// Ovde možete dodati dodatnu obradu greške ako je potrebno
		return err
	}
	return nil
}

func (service *FollowerService) FindFollower(id int) (*model.Follower, error) {
	user, err := service.repo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &user, nil
}

func (us *FollowerService) DeleteFollower(id int) error {
	err := us.repo.DeleteUser(id)
	if err != nil {
		// Ovde možete dodati dodatnu obradu greške ako je potrebno
		return err
	}
	return nil
}
func (us *FollowerService) CheckLoginAvailability(command *events.LoginCommand) (model.Follower, error) {

	fmt.Println("Usao u checklogin:", command)
	reply := &events.LoginReply{}
	follower, err := us.repo.GetById(command.Id)
	if follower.ReportNumber > 3 {
		reply.Type = events.CannotLogin
	} else {
		reply.Type = events.CanLogin
	}
	if err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("Follower service retrieved:", follower)
	}
	return follower, err
}
