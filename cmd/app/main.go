package main

import (
	"backend-hagowagonetka/internal/app"
	"backend-hagowagonetka/internal/config"
)

func main() {
	config.Load(".")
	app.Launch()
}
