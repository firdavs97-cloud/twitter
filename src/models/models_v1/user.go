package models_v1

type User struct {
	ID       int    `json:"id"       mysql:"id"`
	Username string `json:"username" mysql:"id" validate:"required,min=6,max=30"`
	Created  int64  `json:"created" mysql:"created"`
}

type FollowUserRequest struct {
	IdFollower  int `json:"idFollower" validate:"required"`
	IdFollowing int `json:"idFollowing" validate:"required"`
}

func (u User) GetFavouriteTweets() ([]int, error) {
	return nil, nil
}

func (u User) GetFollowers() ([]int, error) {
	return nil, nil
}

func (u User) GetFollowing() ([]int, error) {
	return nil, nil
}

func (u User) Save() error {
	return nil
}

func (u User) Update(s string, i interface{}) error {
	return nil
}

func (u User) Delete() error {
	return nil
}

func (u User) Follow(int) error {
	return nil
}
