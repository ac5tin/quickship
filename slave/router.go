package slave

import "github.com/gofiber/fiber"

// Routes - define /api/slave
func Routes(router *fiber.Group) {
	router.Post("/cmd/:id", cmd)
	router.Post("/clone/:id", clone)
	router.Post("/pull/:id", pull)
	router.Delete("/delete/:id", rm)
	router.Post("/ping", ping)
}
