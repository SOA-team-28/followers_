package model

type Follower struct {
	Id         int64
	Followers  []int
	Followable []int
	Followed   []int
}

func (f *Follower) ConvertToIntSlice(int32Slice []int32) []int {
	intSlice := make([]int, len(int32Slice))
	for i, v := range int32Slice {
		intSlice[i] = int(v)
	}
	return intSlice
}
