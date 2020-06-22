package main

import (
	"flag"
	"fmt"
	"log"
	"quickship/master"

	"github.com/gofiber/compression"
	"github.com/gofiber/fiber"
	"github.com/gofiber/recover"
)

var (
	addr = flag.Int("addr", 8000, "TCP Address to listen to")
)

func main() {
	flag.Parse()

	server()
}

func server() {

	app := fiber.New()
	// middleware
	app.Use(compression.New())
	// ==== API ROUTES =====
	app.Get("/ping", func(c *fiber.Ctx) { c.Status(200).Send("pong") })

	masterapi := app.Group("/api/master")
	master.Routes(masterapi)

	// ===== ERROR RECOVER =====
	cfg := recover.Config{
		Handler: func(c *fiber.Ctx, err error) {
			c.SendString(err.Error())
			c.SendStatus(500)
		},
	}
	app.Use(recover.New(cfg))

	// start server
	log.Println(fmt.Sprintf("Listening on PORT %d", *addr))
	if err := app.Listen(*addr); err != nil {
		log.Fatal(err.Error())
	}
}
