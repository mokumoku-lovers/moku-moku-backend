package users

type User struct {
	Email      string `json:"email"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Biography  string `json:"biography"`
	Birthday   string `json:"birthday"`
	Password   string `json:"-"`
	ProfilePic string `json:"profile_picture"`
	Points     int32  `json:"points"`
}
