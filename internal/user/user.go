package user

type Credentials struct {
	Login    string
	Password string
}

type User struct {
	Id          int64
	Name        string
	LastName    string
	Age         int8
	Credentials Credentials
}
