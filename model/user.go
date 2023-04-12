package model

const (
	UserGenderMale = iota
	UserGenderFemale
)

type User struct {
	Id     int64
	Name   string
	Gender int
}

func (u *User) HasUserInfo() string {
	return u.Name
}
