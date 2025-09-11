package main

import (
	"fmt"
	"log"
	"os"

	"golang-starter-kit/config"
	dbpkg "golang-starter-kit/database"
	"golang-starter-kit/internal/controller"
	"golang-starter-kit/internal/repository"
	"golang-starter-kit/internal/routes"
	"golang-starter-kit/internal/service"

	_ "golang-starter-kit/docs" // This is required for swag to find your docs

	"github.com/gin-gonic/gin"
)

// @title           Golang Starter Kit API
// @version         1.0
// @description     A starter kit for building REST APIs with Golang, Gin, and GORM
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func usage() {
	fmt.Println("Usage:")
	fmt.Println("  go run main.go                   # runs server (default)")
	fmt.Println("  go run main.go server            # runs server")
	fmt.Println("  go run main.go migrate           # runs migrations only")
	fmt.Println("  go run main.go seed              # runs seeders only")
	fmt.Println("  go run main.go migrate seed      # runs migrations & seed")
}

func runServer() error {
	cfg := config.LoadConfig()

	db, err := config.ConnectDatabase(cfg)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Wire repo, service, controller
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, cfg.JWT.Secret)
	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(userService)

	// Setup Gin
	router := gin.Default()
	routes.SetupRoutes(router, userController, authController, cfg.JWT.Secret)

	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("Starting server at %s", addr)
	return router.Run(addr)
}

func runSeed() error {
	cfg := config.LoadConfig()

	db, err := config.ConnectDatabase(cfg)
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	// By default run seeders (without running migrations) to support `go run main.go seed`.
	if err := dbpkg.SeedOnly(db); err != nil {
		return fmt.Errorf("seeding failed: %w", err)
	}

	log.Println("seeder finished successfully")
	return nil
}

func runMigrate() error {
	cfg := config.LoadConfig()

	db, err := config.ConnectDatabase(cfg)
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	if err := dbpkg.MigrateOnly(db); err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	return nil
}

func main() {
	if len(os.Args) == 1 {
		// default to server
		if err := runServer(); err != nil {
			log.Fatalf("server exited with: %v", err)
		}
		return
	}

	switch os.Args[1] {
	case "server":
		if err := runServer(); err != nil {
			log.Fatalf("server exited with: %v", err)
		}
	case "migrate":
		if len(os.Args) >= 3 && os.Args[2] == "seed" {
			if err := runSeed(); err != nil {
				log.Fatalf("seeding failed: %v", err)
			}
		} else {
			if err := runMigrate(); err != nil {
				log.Fatalf("migration failed: %v", err)
			}
		}
	case "seed":
		if err := runSeed(); err != nil {
			log.Fatalf("seeding failed: %v", err)
		}
	case "help", "-h", "--help":
		usage()
	default:
		usage()
	}
}
