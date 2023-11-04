package service

import (
	"go.uber.org/mock/gomock"
	"golang-API/config/rest_err"
	"golang-API/model"
	"golang-API/test/mocks"
	"testing"
)

func TestUserDomainService_FindUserByEmailService(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()
	email := "test@test.com"

	repository := mocks.NewMockUserRepository(control)
	service := NewUserDomainService(repository)

	//repository.EXPECT().FindUserByID(gomock.Any())
	repository.EXPECT().FindUserByEmail(email).Return(
		model.NewUserDomain("teste@teste.com", "123", "teste", 18), nil)

	user, err := service.FindUserByEmailService(email)

	if err != nil {
		t.FailNow()
		return
	}

	if user.GetEmail() != "teste@teste.com" {
		t.FailNow()
		return
	}
}
func TestUserDomainService_FindUserByEmailService_when_id_not_exists(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()
	email := "test@test.com"

	repository := mocks.NewMockUserRepository(control)
	service := NewUserDomainService(repository)

	//repository.EXPECT().FindUserByID(gomock.Any())
	repository.EXPECT().FindUserByEmail(email).Return(nil, rest_err.NewNotFoundError("user not found"))

	user, err := service.FindUserByEmailService(email)

	if err == nil {
		t.FailNow()
		return
	}

	if user != nil {
		t.FailNow()
		return
	}
	if err.Message != "user not found" {
		t.FailNow()
		return
	}
}
