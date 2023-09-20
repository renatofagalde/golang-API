package model

// constructor
func NewUserDomain(email string, passord string, name string, age int8) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: passord,
		name:     name,
		age:      age,
	}
}

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string

	ToJSON() (string, error)
	EncryptPassword()

	AtribuirID(string)
}
