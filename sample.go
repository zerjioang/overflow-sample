package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	uuid "overflow-uint8/go.uuid"
)

func main() {
	//inject a integer overflow
	// Range: 0 through 255.
	var overflowedFinalValue uint8 = 12
	var a uint8 = 6+overflowedFinalValue
	var b uint8 = 250
	c := a+b
	fmt.Println(c)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	u1 := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %s\n", u1)

	app.Listen(":3000")
}
