package models_v1

type Tweet struct {
	Id      int    `json:"id" mysql:"id"`
	Text    string `json:"text" mysql:"text" validate:"min=14,max=280"`
	Created int64  `json:"created" mysql:"created"`
	Author  int    `json:"author" mysql:"author" validate:"required"`
	Media   []int  `json:"media" mysql:"media" validate:"min=0,max=10"`
}

type AddFavouriteRequest struct {
	IdUser  int `json:"idUser" validate:"required"`
	IdTweet int `json:"idTweet" validate:"required"`
}

func (t Tweet) GetFavouriteUsers() ([]int, error) {
	return nil, nil
}

func (t Tweet) AddFavourite(i int) error {
	return nil
}

func (t Tweet) Save() error {
	return nil
}

func (t Tweet) Update(col string, val interface{}) error {
	return nil
}

func (t Tweet) Delete() error {
	return nil
}
