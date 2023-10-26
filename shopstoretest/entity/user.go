package entity

type User struct {
	ID          uint
	Role        Role
	Name        string
	Password    string
	Credit      uint
	PhoneNumber string
}
