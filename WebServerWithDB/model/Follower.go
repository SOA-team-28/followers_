package model

type Follower struct {
	Id         int
	Followers  []int
	Followable []int
	Followed   []int
}
