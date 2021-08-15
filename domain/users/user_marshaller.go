package users

import "encoding/json"

type PublicUser struct {
	Id          int64  `json:"id"`
	DisplayName string `json:"display_name"`
	Biography   string `json:"biography"`
	ProfilePic  string `json:"profile_picture"`
	Points      int32  `json:"points"`
	DateCreated string `json:"date_created"`
}

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
	if isPublic {
		userJson, _ := json.Marshal(user)
		var publicUser PublicUser
		json.Unmarshal(userJson, &publicUser)
		return publicUser
	}
	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}
