package userservice

import (
	"fmt"
	"shopstoretest/entity"
	"shopstoretest/param"
)

func (s Service) Register(req param.UserRegisterRequest) (param.UserRegisterResponse, error) {
	unique, uErr := s.Repository.IsPhoneNumberUnique(req.PhoneNumber)

	if uErr != nil {

		return param.UserRegisterResponse{}, fmt.Errorf("unexpected error %w", uErr)
	}

	if !unique {

		return param.UserRegisterResponse{}, fmt.Errorf("your phone number registered already")
	}

	newUser := entity.User{
		ID:          0,
		Role:        entity.UserRole,
		Name:        req.Name,
		Password:    req.Password,
		Credit:      0,
		PhoneNumber: req.PhoneNumber,
	}

	createdUser, cErr := s.Repository.Register(newUser)

	if cErr != nil {

		return param.UserRegisterResponse{}, fmt.Errorf("unexpected error %w", cErr)
	}

	userInfo := param.UserInfo{
		ID:          createdUser.ID,
		Name:        createdUser.Name,
		PhoneNumber: createdUser.PhoneNumber,
		Credit:      0,
	}

	res := param.UserRegisterResponse{
		UserInfo: userInfo,
	}

	return res, nil
}
