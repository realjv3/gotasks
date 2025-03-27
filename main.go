package main

import "github.com/realjv3/gotasks/cmd"

func main() {
	cmd.RootCmd.AddCommand(cmd.WebCmd)
	cmd.RootCmd.AddCommand(cmd.TaskCmd)
	cmd.TaskCmd.AddCommand(cmd.TaskAddCmd, cmd.TaskGetCmd, cmd.TaskCompleteCmd)

	cmd.Execute()
}
