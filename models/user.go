package models

type User struct {
	Id       int
	Username string
	Password string
	Power    int
}

func (u *User) TableName() string {
	return "mv_user"
}
