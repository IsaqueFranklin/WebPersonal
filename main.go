package main

import (
  //"fmt"
  //"log" 
  //"html/template"
  "github.com/gofiber/template/html/v2"
  "github.com/gofiber/fiber/v2"
  //"github.com/gofiber/template/html/v2"
)

func main() {
    app := fiber.New(fiber.Config{
      Views: html.New("./views", ".html"),
    })

    app.Static("/", "./public", fiber.Static{
      Compress: true,
    })

    app.Get("/", func(c *fiber.Ctx) error {
        return c.Render("index", fiber.Map{})
    })

    app.Get("/about", func(c *fiber.Ctx) error {
      return c.Render("about", fiber.Map{})
    })

    app.Listen(":4000")
}
//bloating go text.
