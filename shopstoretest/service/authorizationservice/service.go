package authorizationservice

import (
	"fmt"
	"shopstoretest/cfg"
	"shopstoretest/entity"
	"shopstoretest/repository/mysql"
)

type Repository interface {
	GetUserPermissionTitles(userID uint) ([]entity.PermissionTitle, error)
}

type Service struct {
	Repository Repository
}

func New(cfg cfg.Cfg) Service {
	myRepo := mysql.New(cfg.DataBaseCfg)
	newService := Service{Repository: myRepo}

	return newService
}

func (s Service) CheckAccess(userID uint, permission entity.PermissionTitle) (bool, error) {
	permissionTitles, err := s.Repository.GetUserPermissionTitles(userID)
	if err != nil {

		return false, fmt.Errorf("unexpected error %w", err)
	}

	for _, per := range permissionTitles {
		if per == permission {

			return true, nil
		}
	}

	return false, nil
}
