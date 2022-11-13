package main

import (
	"blog/internal/app"
	"blog/log"
)

func main() {

	log.Info(" ====================================== ")
	log.Info(" =========== service start ============ ")
	log.Info(" ====================================== ")

	app := &app.App{}
	app.Run()

	log.Info(" ====================================== ")
	log.Info(" ============ service end ============= ")
	log.Info(" ====================================== ")

	return
}
