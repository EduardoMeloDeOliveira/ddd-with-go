package main

import (
	"log"
	http2 "net/http"
	"time"

	"github.com/EduardoMeloDeOliveira/ddd-with-go/project/internal/domain/entity"
	"github.com/EduardoMeloDeOliveira/ddd-with-go/project/internal/domain/service"
	"github.com/EduardoMeloDeOliveira/ddd-with-go/project/internal/infrastructure/repository_impl"
	"github.com/EduardoMeloDeOliveira/ddd-with-go/project/internal/interface/http"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dsn := "host=localhost user=root password=1033 dbname=goDB sslmode=disable port=5432 TimeZone=America/Sao_Paulo"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Falha ao conectar ao banco %v", err)
	}

	if err := db.AutoMigrate(&entity.User{}); err != nil {
		log.Fatalf("Erro no AutoMigrate da table user", err)
	}

	userRepo := repository_impl.NewUserRepositoryImpl(db)
	userService := service.NewUserService(userRepo)
	userController := http.NewUserController(userService)

	healthcheckController := http.NewHealthcheckController()

	router := gin.Default()

	userController.RegisterRoutes(router)
	healthcheckController.RegisterRoute(router)

	srv := &http2.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Servidor rodando em http://localhost:8080 🚀")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Erro no servidor: %v", err)
	}

}
