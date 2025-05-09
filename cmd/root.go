package cmd

import (
	"database/sql"
	"log"

	"github.com/realjv3/gotasks/domain"
	"github.com/realjv3/gotasks/interfaces/storage"
	"github.com/realjv3/gotasks/services"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var app *App

type App struct {
	userService domain.UserService
	authService domain.AuthService
	taskService domain.TaskService
}

func initApp() {
	db, err := sql.Open("sqlite3", "tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	userRepo, err := storage.NewUserRepo(db)
	if err != nil {
		log.Fatal(err)
	}

	taskRepo, err := storage.NewTaskRepo(db)
	if err != nil {
		log.Fatal(err)
	}

	app = &App{
		userService: services.NewUserService(userRepo),
		authService: services.NewAuthService(userRepo),
		taskService: services.NewTaskService(taskRepo),
	}

}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gotasks",
	Short: "A task manager",
	Long:  "A task manager to help you get stuff done.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	initApp()

	err := RootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
