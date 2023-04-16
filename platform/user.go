package platform

type User struct {
	ApiKey string
}

func NewUserData(apiKey string) *User {
	var user User = User{
		ApiKey: apiKey,
	}
	return &user
}
