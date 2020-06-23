package master

import "github.com/gofiber/fiber"

// Routes - define /api/master
func Routes(router *fiber.Group) {
	router.Post("/webhook/:id", webhook)
	router.Get("/list/all", listAll)
	router.Post("/record/add", addNewRec)
	router.Delete("/record/:id", rmRec)
}
