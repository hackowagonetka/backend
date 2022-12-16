package main

import (
	"backend-hagowagonetka/internal/app"
	"backend-hagowagonetka/internal/config"

	"github.com/sirupsen/logrus"
)

func init() {
	config.Load(".")
}

func main() {
	logrus.Info("app is run!")
	app.Launch()
}
