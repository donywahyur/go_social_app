package app

import (
	_ "go_social_app/docs"
	"go_social_app/internal/helpers"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

// swagger handler
func LoadRoute(app *App) {
	app.FiberApp.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Content-Length, Accept, X-Requested-With, Authorization, X-Forwarded-For",
	}))
	app.FiberApp.Use(recover.New())

	app.FiberApp.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        10,
		Expiration: 30 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.JSON(helpers.ResponseApi(fiber.StatusTooManyRequests, "Too many requests", nil))
		},
	}))

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
	authentication.Get("/activate/:token", app.UserHandler.ActivationUser)
	authentication.Post("/login", app.UserHandler.LoginUser)

	users := v1.Group("/users", app.Middlewares.CheckAuth)
	users.Get("/feeds", app.UserHandler.GetUserFeed)
	users.Get("/:id", app.UserHandler.GetUserByID)
	users.Get("/:id/follow", app.UserHandler.FollowUser)
	users.Get("/:id/unfollow", app.UserHandler.UnfollowUser)

	posts := v1.Group("/posts", app.Middlewares.CheckAuth)
	posts.Post("", app.PostHandler.CreatePost)
	posts.Get("/:id", app.PostHandler.GetPostByID)
	posts.Put("/:id", app.Middlewares.CheckRolePrecendence(2), app.PostHandler.UpdatePost)
	posts.Delete("/:id", app.Middlewares.CheckRolePrecendence(3), app.PostHandler.DeletePost)
	posts.Post("/:id/comments", app.PostHandler.CreateComment)

}
