package app

import (
	"go_project/internal/config"
	"go_project/internal/handler"
	"go_project/internal/middleware"
	"go_project/internal/repository"
	"go_project/internal/service"
	"go_project/pkg/database"
	"go_project/pkg/logger"

	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
	cfg    *config.Config
}

func (a *App) Router() *gin.Engine {
	return a.router
}

func New(cfg *config.Config) (*App, error) {
	db, err := database.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	logger.Init(cfg.LogLevel)

	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	orderRepo := repository.NewOrderRepository(db)

	userService := service.NewUserService(userRepo)
	productService := service.NewProductService(productRepo)
	orderService := service.NewOrderService(orderRepo, productRepo)
	authService := service.NewAuthService(userRepo, cfg.JWTSecret)

	userHandler := handler.NewUserHandler(userService)
	productHandler := handler.NewProductHandler(productService)
	orderHandler := handler.NewOrderHandler(orderService)

	router := gin.New()
	router.Use(
		middleware.LoggingMiddlewareMain(),
		middleware.ErrorMiddleware(),
		gin.Recovery(),
	)
	// 添加一個簡單的根路由，用於測試服務是否正常運行
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the e-commerce API"})
	})

	v1 := router.Group("/api/v1")
	{
		v1.POST("/register", userHandler.Register)
		v1.POST("/login", userHandler.Login)

		auth := v1.Group("/")
		auth.Use(middleware.AuthMiddlewareMain(authService))
		{
			auth.POST("/products", productHandler.ListProducts)
			auth.POST("/createProducts", productHandler.CreateProduct)
			auth.GET("/products/:id", productHandler.GetProduct)

			auth.POST("/orders", orderHandler.CreateOrder)
			auth.GET("/orders/:id", orderHandler.GetOrder)
		}
	}

	return &App{router: router, cfg: cfg}, nil
}

func (a *App) Run() error {
	return a.router.Run(a.cfg.ServerAddress)
}
