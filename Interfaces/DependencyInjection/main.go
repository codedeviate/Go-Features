package main

import "fmt"

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	fmt.Println("Log:", message)
}

type App struct {
	logger Logger
}

func (a *App) Run() {
	a.logger.Log("Application started")
}

func main() {
	logger := ConsoleLogger{}
	app := App{logger: logger}
	app.Run()
}
