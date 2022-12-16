package main

import (
	"backend-hagowagonetka/internal/app"
	"backend-hagowagonetka/internal/config"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("app is run!")

	config.Load(".")
	app.Launch()
}
