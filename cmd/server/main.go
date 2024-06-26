// cmd/server/main.go
package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/joshua468/access-bank-exam-system/config"
	"github.com/joshua468/access-bank-exam-system/internal/handlers"
	"github.com/joshua468/access-bank-exam-system/middleware"
	"github.com/joshua468/access-bank-exam-system/models"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
}

// @title Access Bank Exam System API
// @version 1.0
// @description This is an exam system API for Access Bank
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	err = cfg.DB.AutoMigrate(&models.User{}, &models.Exam{}, &models.Question{}, &models.Result{})
	if err != nil {
		logrus.Fatal(err)
	}

	r := gin.Default()

	// Apply rate limiting middleware first
	r.Use(middleware.RateLimitMiddleware(1*time.Minute, 10))

	// Apply logging middleware
	r.Use(middleware.LoggingMiddleware())

	// Apply error handling middleware
	r.Use(middleware.ErrorHandlingMiddleware())

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/register", func(c *gin.Context) { handlers.RegisterUser(c, cfg.DB) })
	r.POST("/login", func(c *gin.Context) { handlers.LoginUser(c, cfg.DB) })

	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.POST("/exams", func(c *gin.Context) { handlers.CreateExam(c, cfg.DB) })
		authorized.GET("/exams", func(c *gin.Context) { handlers.ListExams(c, cfg.DB) })
	}

	r.Run()
}
