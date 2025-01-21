package main

import "go_social_app/internal/app"

func main() {
	application := app.Initialize()

	app.LoadRoute(application)

	application.FiberApp.Listen(":8080")
}
