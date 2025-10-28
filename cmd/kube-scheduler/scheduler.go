package main

import (
	"k6s/cmd/kube-scheduler/app"
	"log/slog"
)

func main() {
	cmd := app.NewSchedulerCommand()

	_ = cmd.Flags()
	if err := cmd.Execute(); err != nil {
		slog.Error("error while executing command", "error", err)
		return
	}
}
