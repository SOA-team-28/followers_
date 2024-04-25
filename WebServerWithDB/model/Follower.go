package model

type Follower struct {
	Id         int64
	Followers  []int
	Followable []int
	Followed   []int
}
