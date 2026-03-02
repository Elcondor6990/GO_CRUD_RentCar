package main

import (

    "github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
    app := fiber.New()

    app.Use(logger.New(logger.Config{
        Format: "${time} ${status} ${latency} ${ip} ${method} ${path}\n",
        TimeFormat: "02/01 15:04:05",
        TimeZone: "Europe/Rome",
    }))

    app.Listen(":3000")
}