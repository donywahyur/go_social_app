package app

import (
	_ "go_social_app/docs"

	"github.com/gofiber/swagger"
)

// swagger handler
func LoadRoute(app *App) {
	api := app.FiberApp.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/swagger/*", swagger.HandlerDefault) // default

	v1.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://localhost:8080/swagger/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))

	authentication := v1.Group("/authentication")
	authentication.Post("/register", app.UserHandler.RegisterUser)

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
