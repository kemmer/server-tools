package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"server-tools/functions"
)

const ServerPort = 7845

func helloWorld(c *fiber.Ctx) error {
	log.Println("request incoming: helloWorld()")

	return c.SendString(functions.HelloWorld())
}

func uptime(c *fiber.Ctx) error {
	log.Println("request incoming: uptime()")

	return c.SendString(fmt.Sprintf("uptime: %q\n", functions.Uptime()))
}

func envs(c *fiber.Ctx) error {
	log.Println("request incoming: envs()")

	return c.SendString(functions.Envs())
}

func runningOs(c *fiber.Ctx) error {
	log.Println("request incoming: runningOs()")

	return c.SendString(functions.RunningOs())
}

func backupInfo(c *fiber.Ctx) error {
	log.Println("request incoming: backupInfo()")

	return c.SendString(functions.BackupInfo())
}

func main() {
	app := fiber.New()

	app.Get("/", helloWorld)
	app.Get("/uptime", uptime)
	app.Get("/envs", envs)
	app.Get("/os", runningOs)
	app.Get("/backup-info", backupInfo)

	err := app.Listen(fmt.Sprintf(":%d", ServerPort))
	if err != nil {
		log.Fatalln("could not initialize server on port 7845")
		return
	}
}
