package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

type userDomain struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int8   `json:"age"`
}

// constructor
func NewUserDomain(email string, passord string, name string, age int8) UserDomainInterface {
	return &userDomain{
		Email:    email,
		Password: passord,
		Name:     name,
		Age:      age,
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

func (ud *userDomain) ToJSON() (string, error) {
	b, err := json.Marshal(ud)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (ud *userDomain) AtribuirID(id string) {
	ud.ID = id
}
func (ud *userDomain) GetEmail() string {
	return ud.Email
}
func (ud *userDomain) GetName() string {
	return ud.Name
}
func (ud *userDomain) GetAge() int8 {
	return ud.Age
}
func (ud *userDomain) GetPassword() string {
	return ud.Password
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
