package main

import (
	"errors"
	"klik_test/auth"
	"klik_test/handler"
	"klik_test/middleware"
	"klik_test/transaction"
	"klik_test/user"
	"os"

	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	secret := os.Getenv("SECRET_KEY")
	auth.Secret = []byte(secret)

	dsn := "root:r00tp4ss@tcp(127.0.0.1:3306)/klik_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	if secret == "" {
		log.Fatal(errors.New("No Secret Key Found"))
	}

	userRepository := user.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	userService := user.NewService(userRepository)
	authService := auth.NewService()
	transactionService := transaction.NewService(transactionRepository)

	userHandler := handler.NewUserHandler(userService, authService)
	transactionHandler := handler.NewTransactionHandler(transactionService, authService)

	router := gin.Default()
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte(secret))
	store.Options(sessions.Options{MaxAge: 3600})
	router.Use(sessions.Sessions("mysession", store))

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("login", userHandler.Login)
	api.POST("/email_checker", userHandler.AvailabilityEmail)

	api.Use(middleware.Authentication())
	{
		api.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Everything is ok",
			})
		})

		api.POST("/transaction", transactionHandler.CreateTransaction)
	}

	router.Run()
}
