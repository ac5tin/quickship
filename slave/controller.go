package slave

import (
	"log"
	"quickship/worker"

	"github.com/gofiber/fiber"
)

func cmd(c *fiber.Ctx) {
	var cmdreq CmdReq
	if err := c.BodyParser(&cmdreq); err != nil {
		log.Println("Failed to parse body")
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  "Failed to parse body",
		})
		return
	}
	//uid := c.Params("id") // record id
	if err := worker.Run(cmdreq.Command); err != nil {
		log.Println("Failed to parse body")
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  "Failed to execute command",
		})
		return
	}
	// all done
	c.Status(200).JSON(fiber.Map{
		"result": "success",
	})
	return
}
