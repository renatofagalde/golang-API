package model

import "golang-basic/config/rest_err"

// constructor
func NewUserDomain(email string, passord string, name string, age int8) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: passord,
		name:     name,
		age:      age,
	}
}

func NewUseUpdaterDomain(name string, age int8) UserDomainInterface {
	return &userDomain{
		name: name,
		age:  age,
	}
}

func NewUserLoginDomain(email, password string) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password}
}

type UserDomainInterface interface {
	GetID() string

	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string

	ToJSON() (string, error)
	EncryptPassword()

	AtribuirID(string)

	GenerateToken() (string, *rest_err.RestErr)
}
