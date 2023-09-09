package model

import (
	"crypto/md5"
	"encoding/hex"
	"golang-basic/config/rest_err"
)

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

// constructor
func NewUserDomain(email string, passord string, name string, age int8) UserDomainInterface {
	return &UserDomain{email, passord, name, age}
}

type UserDomainInterface interface {
	Create() *rest_err.RestErr
	Update(string) *rest_err.RestErr
	FindUser(string) (*UserDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
