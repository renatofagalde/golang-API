package model

import "golang-basic/config/rest_err"

func (ud *UserDomain) FindUser(string) (*UserDomain, *rest_err.RestErr) {
	return ud, nil
}
