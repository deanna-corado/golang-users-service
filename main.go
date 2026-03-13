package main

import (
	"log"
	"os"
	clients "user-service/client"
	"user-service/config"
	"user-service/controllers"
	"user-service/migrations"
	"user-service/repositories"
	"user-service/routes"
	"user-service/services"

	"github.com/gin-gonic/gin"
	"github.com/go-gormigrate/gormigrate/v2"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
}

func main() {

	m := gormigrate.New(
		config.DB,
		gormigrate.DefaultOptions,
		migrations.GetMigrations(),
	)

	//MIGRATE ALL LISTED MIGRATIONS
	if err := m.Migrate(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migrations applied successfully")

	log.Println("Loaded migrations:", len(migrations.GetMigrations()))

	//ROLLBACK LAST MIGRATION ONLY

	// if err := m.RollbackLast(); err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("Last migration rolled back successfully")

	//ROLLBACK SPECIFIC MIGRATION

	// if err := m.RollbackTo("00001_createUserTable"); err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("Rolled back to specified migration successfully")

	r := gin.Default()

	userRepo := repositories.NewUserRepository(config.DB)

	authService := services.NewAuthService(userRepo) // login/reg + jwt
	userService := services.NewUserService(userRepo) // user crudd]

	moviesClient := clients.NewMoviesClient()

	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(userService)
	movieController := controllers.NewMovieController(moviesClient)

	credRepo := repositories.NewCredentialRepository(config.DB)
	credService := services.NewCredentialService(credRepo)
	credController := controllers.NewCredentialsController(credService)

	routes.RegisterAuthRoutes(r, authController, credController)
	routes.RegisterUserRoutes(r, userController)
	routes.RegisterMovieRoutes(r, movieController)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
