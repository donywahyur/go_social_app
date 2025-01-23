package app

func LoadRoute(app *App) {
	api := app.FiberApp.Group("/api")

	v1 := api.Group("/v1")

	v1.Get("/", app.UserHandler.Create)
}
