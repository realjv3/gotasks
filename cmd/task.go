package cmd

import (
	"fmt"
	"log/slog"
	"strconv"

	"github.com/realjv3/gotasks/domain"

	"github.com/spf13/cobra"
)

// TaskCmd represents the task command
var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage tasks",
	Long:  "Create and finish tasks",
}

var TaskAddCmd = &cobra.Command{
	Use:   "add {title} {description}",
	Short: "Create task",
	Long:  "This will create a tasks with the given title and description",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		task, err := app.taskService.CreateTask(&domain.Task{
			Title:       args[0],
			Description: args[1],
		})
		if err != nil {
			slog.Error("error creating task", err.Error())
			return
		}

		slog.Info(fmt.Sprintf("task created: %#v", task))
	},
}

var TaskGetCmd = &cobra.Command{
	Use:   "get {userID}",
	Short: "Get a user's tasks",
	Long:  "This will get all of a user's tasks",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		userID, err := strconv.Atoi(args[0])
		if err != nil {
			slog.Error("error converting user ID to int", err.Error())
			return
		}

		tasks, err := app.taskService.GetTasksByUser(userID)
		if err != nil {
			slog.Error("error creating task", err.Error())
			return
		}

		for _, task := range tasks {
			slog.Info(fmt.Sprintf("%#v", task))
		}
	},
}

var TaskCompleteCmd = &cobra.Command{
	Use:   "complete {taskID}",
	Short: "Complete a task",
	Long:  "This will mark a task a completed",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			slog.Error("error converting task ID to int", err.Error())
			return
		}

		task, err := app.taskService.FinishTask(taskID)
		if err != nil {
			slog.Error("error completing task", err.Error())
			return
		}

		slog.Info(fmt.Sprintf("%#v", task))
	},
}
