package app

func LoadRoute(app *App) {
	api := app.FiberApp.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/", app.UserHandler.Create)

	post := v1.Group("/posts")
	post.Post("", app.PostHandler.CreatePost)
	post.Get("/:id", app.PostHandler.GetPostByID)

}
