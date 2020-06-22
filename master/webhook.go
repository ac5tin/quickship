package master

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/gofiber/fiber"
)

func webhook(c *fiber.Ctx) {
	var hookdata map[string]interface{}
	if err := c.BodyParser(&hookdata); err != nil {
		log.Println("Failed to parse body")
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  "Failed to parse body",
		})
		return
	}

	b, err := json.Marshal(hookdata)
	if err != nil {
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  "Failed to parse body",
		})
		return
	}

	ref, err := jsonparser.GetString(b, "ref")
	if err != nil {
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  "Failed to parse body",
		})
		return
	}

	repo, err := jsonparser.GetString(b, "repository", "name")
	if err != nil {
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  "Failed to parse body",
		})
		return
	}

	splits := strings.Split(ref, "/")
	branch := splits[len(splits)-1]

	// all done, now return data
	// return data back to client
	c.Status(200).JSON(fiber.Map{
		"result": "success",
	})
	return
}
