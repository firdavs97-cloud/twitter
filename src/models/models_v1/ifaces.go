package models_v1

import "time"

type UserIface interface {
	Follow(int) error
	Save() error
	Update(string, interface{}) error
	Delete() error
	GetFavouriteTweets() ([]int, error)
	GetFollowers() ([]int, error)
	GetFollowing() ([]int, error)
}

func GetUser(id int) (UserIface, error) { // TODO
	return User{
		ID:       id,
		Username: "TestUser",
	}, nil
}

type TweetIface interface {
	Save() error
	Update(string, interface{}) error
	Delete() error
	AddFavourite(int) error
	GetFavouriteUsers() ([]int, error)
}

func GetTweet(id int) (TweetIface, error) { // TODO
	return Tweet{
		Id:      id,
		Text:    "TestText",
		Created: time.Now().Unix(),
		Author:  1,
		Media:   []int{1, 2},
	}, nil
}
