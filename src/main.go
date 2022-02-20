package main

import (
	"github.com/charichu/textsapi/app"
	"os"
)

func main() {
	os.Setenv("LOG_LEVEL", "info")

	app.StartApplication()
}
