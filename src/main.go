package main

import (
	"github.com/charichu/textsapi/src/app"
	"os"
)

func main() {
	os.Setenv("LOG_LEVEL", "info")

	app.StartApplication()
}
