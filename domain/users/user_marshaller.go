package users

import "encoding/json"

//TODO: PublicUser type

type PrivateUser struct {
	Id          int64  `json:"id"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Biography   string `json:"biography"`
	Birthday    string `json:"birthday"`
	ProfilePic  string `json:"profile_picture"`
	Points      int32  `json:"points"`
	DateCreated string `json:"date_created"`
}

func (user *User) Marshall(isPublic bool) interface{} {
	//TODO: PublicUser marshalling
	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}
