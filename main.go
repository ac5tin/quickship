package main

import (
	"flag"
	"fmt"
	"log"
	"quickship/deploy"
	"quickship/master"
	"quickship/slave"
	"quickship/store"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	_ "github.com/joho/godotenv/autoload"
)

var (
	port    = flag.Int("p", 7291, "Port to listen to")
	ms      = flag.String("ms", "http://localhost", "URL endpoint of master node server")
	path    = flag.String("path", "./qs.json", "Path to qs.json which stores deployment info")
	srv     = flag.Bool("s", false, "Server mode")
	list    = flag.Bool("l", false, "List deployment records")
	rid     = flag.String("id", "", "Record ID")
	name    = flag.String("n", "", "Name of deployment")
	up      = flag.String("up", "", "Path to file to be added to deployment list")
	down    = flag.String("down", "", "ID of deployment record to be removed")
	addnode = flag.String("addnode", "", "Add a single node to a record")
	delnode = flag.String("delnode", "", "Delete a single node")
	info    = flag.Bool("i", false, "Display info of a record")
	rd      = flag.Bool("rd", false, "Re-deploy a record")
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
	app.Use(middleware.Compress())

	// store
	s := store.Init(*path)
	go deploy.KeepStoreAlive(s)
	app.Use(func(c *fiber.Ctx) {
		c.Locals("store", s)
		c.Next()
	})

	// ==== API ROUTES =====
	app.Get("/ping", func(c *fiber.Ctx) { c.Status(200).Send("pong") })

	masterapi := app.Group("/api/master")
	master.Routes(&masterapi)

	slaveapi := app.Group("/api/slave")
	slave.Routes(&slaveapi)

	// ===== ERROR RECOVER =====
	app.Use(middleware.Recover())

	// start server
	log.Println(fmt.Sprintf("Listening on PORT %d", *port))
	if err := app.Listen(*port); err != nil {
		log.Fatal(err.Error())
	}
}
