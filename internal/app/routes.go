package app

func LoadRoute(app *App) {
	api := app.FiberApp.Group("/api")
	v1 := api.Group("/v1")

	users := v1.Group("/users")
	users.Get("/feeds", app.UserHandler.GetUserFeed)
	users.Get("/:id", app.UserHandler.GetUserByID)
	users.Get("/:id/follow", app.UserHandler.FollowUser)
	users.Get("/:id/unfollow", app.UserHandler.UnfollowUser)

	posts := v1.Group("/posts")
	posts.Post("", app.PostHandler.CreatePost)
	posts.Get("/:id", app.PostHandler.GetPostByID)
	posts.Put("/:id", app.PostHandler.UpdatePost)
	posts.Post("/:id/comments", app.PostHandler.CreateComment)

}
