package model

import (
	"crypto/md5"
	"encoding/hex"
)

type userDomain struct {
	email    string
	password string
	name     string
	age      int8
}

// constructor
func NewUserDomain(email string, passord string, name string, age int8) UserDomainInterface {
	return &userDomain{email, passord, name, age}
}

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string

	EncryptPassword()
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}
func (ud *userDomain) GetName() string {
	return ud.name
}
func (ud *userDomain) GetAge() int8 {
	return ud.age
}
func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}
