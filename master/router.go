package master

import "github.com/gofiber/fiber"

// Routes - define /api/master
func Routes(router *fiber.Group) {
	router.Post("/webhook/:id", webhook)
}
