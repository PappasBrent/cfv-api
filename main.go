package main

import (
	"cfv-api/config"
	"fmt"
)

func main() {
	if app, err := config.SetupApp(); err != nil {
		fmt.Println(err)
	} else {
		app.Run()
	}
}
