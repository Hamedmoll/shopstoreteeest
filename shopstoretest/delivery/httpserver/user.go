package httpserver

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shopstoretest/param"
)

func (s Server) userRegister(c echo.Context) error {
	req := param.UserRegisterRequest{}

	bErr := c.Bind(&req)
	if bErr != nil {
		//fmt.Println("\n\n\n here \n\n\n")
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	res, err := s.userService.Register(req)
	if err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}

func (s Server) userLogin(c echo.Context) error {
	req := param.UserLoginRequest{}

	bErr := c.Bind(&req)
	if bErr != nil {

		return echo.NewHTTPError(http.StatusBadRequest)
	}

	res, lErr := s.userService.Login(req)
	if lErr != nil {

		return echo.NewHTTPError(http.StatusBadRequest, lErr.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (s Server) userProfile(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")

	req := param.UserProfileRequest{Token: token}
	res, pErr := s.userService.Profile(req)

	if pErr != nil {

		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	return c.JSON(http.StatusOK, res)
}
