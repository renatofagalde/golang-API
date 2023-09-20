package model

import "github.com/goccy/go-json"

func (ud *userDomain) ToJSON() (string, error) {
	b, err := json.Marshal(ud)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
