package main

import (
	"os"

	cli "github.com/jawher/mow.cli"
	log "github.com/sirupsen/logrus"
)

var gitHash = "overridden at compile time"

const (
	appName = "go-todo-service"
	appDesc = "The gRPC Todo service example in go"
)

func main() {
	app := cli.App(appName, appDesc)

	logger := log.WithField("git_hash", gitHash)

	app.Action = func() {
		logger.Println("app started")

		logger.Println("bye:)")
	}

	if err := app.Run(os.Args); err != nil {
		logger.WithError(err).Fatalln("app run")
	}
}
