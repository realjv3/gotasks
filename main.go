package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/realjv3/gotasks/interfaces/rest"
	"github.com/realjv3/gotasks/interfaces/storage"
	"github.com/realjv3/gotasks/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.Println("Initializing database...")

	db, err := sql.Open("sqlite3", "tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Initializing services...")

	userRepo, err := storage.NewUserRepo(db)
	if err != nil {
		log.Fatal(err)
	}

	taskRepo, err := storage.NewTaskRepo(db)
	if err != nil {
		log.Fatal(err)
	}

	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userRepo)
	taskService := services.NewTaskService(taskRepo)

	log.Println("Starting HTTP server...")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Second * 10))

	authHandler := rest.NewAuthHandler(authService, userService)
	authHandler.RegisterAuthRoutes(r)

	userHandler := rest.NewUserHandler(userService)
	userHandler.RegisterUserRoutes(r)

	taskHandler := rest.NewTaskHandler(taskService)
	taskHandler.RegisterRoutes(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
