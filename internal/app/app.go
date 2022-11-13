package app

import (
	"blog/internal/httpserver"
)

type App struct {
}

func (m *App) Run() {

	httpserver.Run()

}

func (m *App) Stop() {

}
