package models

//User model for saving user
type User struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string
	Password  string
	Active    bool
}
