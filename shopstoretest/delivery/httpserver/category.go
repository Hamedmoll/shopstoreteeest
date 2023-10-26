package httpserver

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shopstoretest/entity"
	"shopstoretest/param"
)

func (s Server) addCategory(c echo.Context) error {
	req := param.AddCategoryRequest{}
	bErr := c.Bind(&req)
	if bErr != nil {

		return echo.NewHTTPError(http.StatusBadRequest)
	}

	BearerTokenString := c.Request().Header.Get("authorization")
	claim, pErr := s.authService.ParseToken(BearerTokenString)
	if pErr != nil {

		return echo.NewHTTPError(http.StatusBadRequest)
	}

	haveAccess, hErr := s.authorizationService.CheckAccess(claim.ID, entity.AddCategoryPermission)
	if hErr != nil {

		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if !haveAccess {

		return echo.NewHTTPError(http.StatusBadRequest)
	}

	res, err := s.categoryService.AddCategory(req)
	if err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}
