package userservice

import (
	"shopstoretest/cfg"
	"shopstoretest/entity"
	"shopstoretest/repository/mysql"
	"shopstoretest/service/authservice"
)

type Service struct {
	Repository  Repository
	authService AuthService
}

type AuthService interface {
	CreateAccessToken(id uint, role entity.Role) (string, error)
	CreateRefreshToken(id uint, role entity.Role) (string, error)
	ParseToken(barerToken string) (authservice.Claim, error)
}

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(user entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, error)
	GetUserByID(id uint) (entity.User, error)
}

func New(cfg cfg.Cfg) Service {
	myRepo := mysql.New(cfg.DataBaseCfg)
	authSrv := authservice.New(cfg)

	return Service{
		Repository:  myRepo,
		authService: authSrv,
	}
}
