package dto

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatarUrl"`
	FullName  string `json:"fullName"`
	UserName  string `json:"userName"`
}
