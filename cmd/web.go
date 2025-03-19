package cmd

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
	"github.com/spf13/cobra"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Run as web server",
	Long:  "The app will run as a web server.",
	Run:   runWeb,
}

func runWeb(cmd *cobra.Command, args []string) {
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

func init() {
	rootCmd.AddCommand(webCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// webCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// webCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
