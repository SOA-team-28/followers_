package repo

import (
	"database-example/model"
	"fmt"

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
func (ur *FollowerRepository) GetById(id int) (model.Follower, error) {
	var user model.Follower

	session := ur.driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	defer session.Close()

	_, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (f:Follower {id: $id}) RETURN f.id, f.followers, f.followable, f.followed",
			map[string]interface{}{"id": id},
		)
		if err != nil {
			return nil, err
		}

		if result.Next() {
			record := result.Record()
			fmt.Println("record")
			fmt.Println(record)
			id := record.GetByIndex(0).(int64)
			fmt.Println("id")
			fmt.Println(id)
			followers := convertToIntSlice(record.GetByIndex(1).([]interface{}))
			followable := convertToIntSlice(record.GetByIndex(2).([]interface{}))
			followed := convertToIntSlice(record.GetByIndex(3).([]interface{}))

			follower := model.Follower{
				Id:         id,
				Followers:  followers,
				Followable: followable,
				Followed:   followed,
			}
			user.Id = id
			user.Followable = followable
			user.Followed = followed
			user.Followers = followers
			fmt.Println("follower")
			fmt.Println(follower)
			return follower, nil
		} else {
			return model.Follower{}, fmt.Errorf("Follower with ID %d not found", id)
		}

	})

	if err != nil {
		return model.Follower{}, err
	}
	fmt.Println("USER")
	fmt.Println(user)
	return user, nil
}

// Funkcija za konvertovanje interfejsa u listu intova
func convertToIntSlice(interfaces []interface{}) []int {
	var ints []int
	for _, v := range interfaces {
		if intValue, ok := v.(int64); ok {
			ints = append(ints, int(intValue))
		}
	}
	return ints
}
