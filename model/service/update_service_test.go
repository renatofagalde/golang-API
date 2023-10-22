package service

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"golang-basic/config/rest_err"
	"golang-basic/model"
	"golang-basic/test/mocks"
	"testing"
)

func TestUserDomainService_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_sending_a_valid_user_and_userId_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 50)

		repository.EXPECT().UpdateUser(id, userDomain).Return(nil)

		err := service.UpdateService(id, userDomain)

		assert.Nil(t, err)
	})

	t.Run("when_sending_a_invalid_user_and_userId_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 50)

		repository.EXPECT().UpdateUser(id, userDomain).Return(
			rest_err.NewInternalServerError("error trying to update user"))

		err := service.UpdateService(id, userDomain)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to update user")
	})
}
