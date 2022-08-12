package routes

import (
	"github.com/Similadayo/myBlog/controller"
	"github.com/Similadayo/myBlog/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Use(middleware.IsAuthenticated)
	app.Post("/api/post", controller.CreatePost)
	app.Get("/api/getposts", controller.GetPosts)
}
