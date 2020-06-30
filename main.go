package main

import (
	"flag"
	"fmt"
	"log"
	"quickship/master"
	"quickship/slave"
	"quickship/store"

	"github.com/gofiber/compression"
	"github.com/gofiber/fiber"
	"github.com/gofiber/recover"
	_ "github.com/joho/godotenv/autoload"
)

var (
	addr = flag.Int("addr", 8000, "TCP Address to listen to")
	ms   = flag.String("ms", "http://localhost", "URL endpoint of master node server")
	path = flag.String("path", "./qs.json", "Path to qs.json which stores deployment info")
	srv  = flag.Bool("s", false, "Server mode")
	name = flag.String("n", "", "Name of deployment")
	up   = flag.String("up", "", "Path to file to be added to deployment list")
	down = flag.String("down", "", "ID of deployment record to be removed")
	list = flag.Bool("l", false, "List deployment records")
)

func main() {
	flag.Parse()
	if *srv {
		server()
		return
	}
	// Command line mode
	cmd()

}

func server() {

	app := fiber.New()
	// middleware
	app.Use(compression.New())

	// store
	s := store.Init(*path)
	app.Use(func(c *fiber.Ctx) {
		c.Locals("store", s)
		c.Next()
	})

	// ==== API ROUTES =====
	app.Get("/ping", func(c *fiber.Ctx) { c.Status(200).Send("pong") })

	masterapi := app.Group("/api/master")
	master.Routes(masterapi)

	slaveapi := app.Group("/api/slave")
	slave.Routes(slaveapi)

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
