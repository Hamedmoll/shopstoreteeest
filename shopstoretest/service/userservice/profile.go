package userservice

import (
	"fmt"
	"shopstoretest/param"
)

func (s Service) Profile(req param.UserProfileRequest) (param.UserProfileResponse, error) {
	barerToken := req.Token

	claim, pErr := s.authService.ParseToken(barerToken)
	if pErr != nil {

		return param.UserProfileResponse{}, fmt.Errorf("cant parse token %w", pErr)
	}

	user, gErr := s.Repository.GetUserByID(claim.ID)
	if gErr != nil {

		return param.UserProfileResponse{}, fmt.Errorf("unexpected error %w", gErr)
	}

	userInfo := param.UserInfo{
		ID:          user.ID,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Credit:      user.Credit,
	}

	return param.UserProfileResponse{UserInfo: userInfo}, nil
}
