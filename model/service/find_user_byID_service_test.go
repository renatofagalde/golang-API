package service

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"golang-API/config/rest_err"
	"golang-API/model"
	"golang-API/test/mocks"
	"testing"
)

func TestUserDomainService_FindUserByIDService(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	repo := mocks.NewMockUserRepository(control)
	service := NewUserDomainService(repo)

	t.Run("when_exists_an_user_return_success", func(t *testing.T) {

		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain("teste@teste.com", "123", "teste", 18)

		//repo.EXPECT().FindUserByID(gomock.Any())
		repo.EXPECT().FindUserByID(id).Return(userDomain, nil)

		userDomainReturn, err := service.FindUserByIDService(id)
		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainReturn.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainReturn.GetAge(), userDomain.GetAge())

	})

	t.Run("when_does_not_exists_an_user_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		repo.EXPECT().FindUserByID(id).Return(nil, rest_err.NewNotFoundError("user not found"))
		userDomainReturn, err := service.FindUserByIDService(id)

		assert.Nil(t, userDomainReturn)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "user not found")
	})
}
