package main

import "go_social_app/internal/app"

// @title Fiber Go Social API
// @version 1.0
// @description API for go social
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1
//
//

func main() {
	application := app.Initialize()

	app.LoadRoute(application)

	application.FiberApp.Listen(":8080")
}
