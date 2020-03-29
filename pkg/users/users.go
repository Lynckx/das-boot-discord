package users

type User struct {
	Name string
	Id   string
}

func (u User) GetName() string {
	return u.Name
}
func (u User) GetId() string {
	return u.Id
}

func NewUser(id string, name string) User {
	var newUser User
	newUser.Name = name
	newUser.Id = id
	return newUser
}
