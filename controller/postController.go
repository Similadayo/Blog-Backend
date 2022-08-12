package controller

import (
	"fmt"
	"math"
	"strconv"

	"github.com/Similadayo/myBlog/database"
	"github.com/Similadayo/myBlog/models"
	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	var blogPosts models.Posts
	if err := c.BodyParser(&blogPosts); err != nil {
		fmt.Println("Unable to parse body")
	}

	if err := database.DB.Create(&blogPosts).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid payload",
		})
	} else {
		return c.JSON(fiber.Map{
			"message": "Your post was sent",
		})
	}
}

func GetPosts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page - 1) * limit
	var total int64
	var getPosts []models.Posts
	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getPosts)
	database.DB.Model(&models.Posts{}).Count(&total)
	return c.JSON(fiber.Map{
		"data": getPosts,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
}
