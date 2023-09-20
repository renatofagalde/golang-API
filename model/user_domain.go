package model

// unica responsabilidade, ter o valor ou alterar o que pode ser alterado dentro do domain
// valor real do objeto
type userDomain struct {
	id       string
	email    string
	password string
	name     string
	age      int8
}

func (ud *userDomain) AtribuirID(id string) {
	ud.id = id
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
