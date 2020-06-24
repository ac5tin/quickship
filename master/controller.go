package master

import (
	"encoding/json"
	"log"
	"quickship/deploy"
	"quickship/store"
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

	splits := strings.Split(ref, "/")
	branch := splits[len(splits)-1]

	uid := c.Params("id") // record id
	s := c.Locals("store").(*store.Store)
	d := s.GetRecordDeploy(uid)
	if d.Branch == branch {
		go deploy.Record(d, uid)
	}

	// all done, now return data
	// return data back to client
	c.Status(200).JSON(fiber.Map{
		"result": "success",
	})
	return
}

func listAll(c *fiber.Ctx) {
	s := c.Locals("store").(*store.Store)
	list := s.GetList()
	// all done, now return data
	// return data back to client
	c.Status(200).JSON(fiber.Map{
		"result": "success",
		"data":   list,
	})
	return
}

func addNewRec(c *fiber.Ctx) {
	s := c.Locals("store").(*store.Store)

	var recdata store.Record
	if err := c.BodyParser(&recdata); err != nil {
		log.Println("Failed to parse body")
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  "Failed to parse body",
		})
		return
	}
	uid, err := AddRecord(recdata.Deploy, recdata.Name, s)
	if err != nil {
		log.Println("Failed to add record")
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  "Failed to add record",
		})
		return
	}

	go deploy.NewRecord(recdata.Deploy, uid)

	// all done, now return data
	// return data back to client
	c.Status(200).JSON(fiber.Map{
		"result": "success",
	})
	return
}

func rmRec(c *fiber.Ctx) {
	s := c.Locals("store").(*store.Store)
	uid := c.Params("id") // record id
	if !s.Exist(uid) {
		// does not exist
		log.Println("ID does not exist")
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  "ID does not exist",
		})
		return
	}
	dp := s.GetRecordDeploy(uid) // deploy record
	go deploy.DelRecord(dp, uid)
	if err := rmRecord(uid, s); err != nil {
		log.Println("Failed to remove record")
		log.Println(err.Error())
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  "Failed to remove record",
		})
		return
	}

	// all done, now return data
	// return data back to client
	c.Status(200).JSON(fiber.Map{
		"result": "success",
	})
	return
}

func getRec(c *fiber.Ctx) {
	s := c.Locals("store").(*store.Store)
	uid := c.Params("id") // record id
	if !s.Exist(uid) {
		// does not exist
		log.Println("ID does not exist")
		c.Status(400).JSON(fiber.Map{
			"result": "error",
			"error":  "ID does not exist",
		})
		return
	}
	recordata := s.GetRecord(uid)
	// all done, now return data
	// return data back to client
	c.Status(200).JSON(fiber.Map{
		"result": "success",
		"data":   recordata,
	})
	return
}
