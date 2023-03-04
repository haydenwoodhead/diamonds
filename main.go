package main

import "golang.org/x/exp/slog"

func main() {
	logger := slog.Default()

	logger.Info("starting diamonds")
}
