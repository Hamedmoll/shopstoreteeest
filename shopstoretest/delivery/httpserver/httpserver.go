package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"shopstoretest/service/authorizationservice"
	"shopstoretest/service/authservice"
	"shopstoretest/service/categoryservice"
	"shopstoretest/service/productservice"
	"shopstoretest/service/userservice"
)

type Server struct {
	userService          userservice.Service
	authorizationService authorizationservice.Service
	categoryService      categoryservice.Service
	authService          authservice.Service
	productService       productservice.Service
}

func (s Server) Serve() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userGroup := e.Group("/users")
	//adminGroup := e.Group("/admins")
	//categoryGroup := e.Group("/category")

	e.GET("/health-check", s.heathCheck)

	userGroup.POST("/register", s.userRegister)
	userGroup.POST("/login", s.userLogin)
	userGroup.GET("/profile", s.userProfile)

	e.POST("/categories/add", s.addCategory)
	e.POST("/products/add", s.addProduct)

	//adminGroup.POST("/login", s.adminLogin)

	//categoryGroup.POST("/add", s.addCategory)

	e.Logger.Fatal(e.Start(":5555"))
}

func New(userService userservice.Service, authService authservice.Service,
	authorizationService authorizationservice.Service, categoryService categoryservice.Service,
	productService productservice.Service) Server {
	srv := Server{
		userService:          userService,
		authorizationService: authorizationService,
		categoryService:      categoryService,
		authService:          authService,
		productService:       productService,
	}

	return srv
}
