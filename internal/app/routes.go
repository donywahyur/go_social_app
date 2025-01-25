package app

func LoadRoute(app *App) {
	api := app.FiberApp.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/", app.UserHandler.Create)

	posts := v1.Group("/posts")
	posts.Post("", app.PostHandler.CreatePost)
	posts.Get("/:id", app.PostHandler.GetPostByID)
	posts.Put("/:id", app.PostHandler.UpdatePost)
	posts.Post("/:id/comments", app.PostHandler.CreateComment)

}
