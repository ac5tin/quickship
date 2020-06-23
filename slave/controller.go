package slave

import (
	"log"
	"os"
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
	recordid := c.Params("id") // record id
	if err := worker.Run(cmdreq.Command, recordid); err != nil {
		log.Println("Failed to execute command")
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

func clone(c *fiber.Ctx) {
	var req cloneReq
	if err := c.BodyParser(&req); err != nil {
		log.Println("Failed to parse body")
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  "Failed to parse body",
		})
		return
	}
	recordid := c.Params("id") // record id
	if err := worker.Clone(req.Repo, req.Branch, recordid, os.Getenv("GITHUB_TOKEN")); err != nil {
		log.Println("Failed to clone")
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  "Failed to clone",
		})
		return
	}
	// all done
	c.Status(200).JSON(fiber.Map{
		"result": "success",
	})
	return
}

func rm(c *fiber.Ctx) {
	recordid := c.Params("id") // record id
	if err := os.RemoveAll(recordid); err != nil {
		log.Println("Failed to remove")
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  "Failed to remove",
		})
		return
	}
	// all done
	c.Status(200).JSON(fiber.Map{
		"result": "success",
	})
	return
}
