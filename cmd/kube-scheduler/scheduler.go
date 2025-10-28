package main

import (
	"k6s/cmd/kube-scheduler/app"
	"log/slog"
)

func main() {
	command := app.NewSchedulerCommand()

	if err := command.Execute(); err != nil {
		slog.Error("error while executing command", "error", err)
		return
	}
}
