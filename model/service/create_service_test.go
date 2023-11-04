package service

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang-API/config/rest_err"
	"golang-API/model"
	"golang-API/test/mocks"
	"testing"
)

func TestUserDomainService_CreateUserServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_user_already_exists_returns_error", func(t *testing.T) {

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 50)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(userDomain, nil)

		user, err := service.CreateService(userDomain)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "email já existente")
	})

	t.Run("when_user_is_not_registered_returns_error", func(t *testing.T) {

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 50)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(nil, nil)

		repository.EXPECT().CreateUser(userDomain).Return(
			nil, rest_err.NewInternalServerError("error trying to create user"))

		user, err := service.CreateService(userDomain)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to create user")
	})

	t.Run("when_user_is_not_registered_returns_success", func(t *testing.T) {

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 50)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(nil, nil)

		repository.EXPECT().CreateUser(userDomain).Return(userDomain, nil)

		user, err := service.CreateService(userDomain)

		assert.Nil(t, err)
		assert.EqualValues(t, user.GetName(), userDomain.GetName())
		assert.EqualValues(t, user.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, user.GetAge(), userDomain.GetAge())
		assert.EqualValues(t, user.GetID(), userDomain.GetID())
		assert.EqualValues(t, user.GetPassword(), userDomain.GetPassword())
	})
}
