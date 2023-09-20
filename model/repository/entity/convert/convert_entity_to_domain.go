package convert

import (
	"golang-basic/model"
	"golang-basic/model/repository/entity"
)

func ConvertEntityToDomain(userEntity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(userEntity.Email, userEntity.Password, userEntity.Name, userEntity.Age)
	domain.AtribuirID(userEntity.ID.Hex())
	return domain
}
