package cmd

import (
	"log"
	"net/http"
	"time"

	"github.com/realjv3/gotasks/interfaces/rest"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	log.Println("Starting HTTP server...")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Second * 10))

	authHandler := rest.NewAuthHandler(app.authService, app.userService)
	authHandler.RegisterAuthRoutes(r)

	userHandler := rest.NewUserHandler(app.userService)
	userHandler.RegisterUserRoutes(r)

	taskHandler := rest.NewTaskHandler(app.taskService)
	taskHandler.RegisterRoutes(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func init() {
	rootCmd.AddCommand(webCmd)
}
