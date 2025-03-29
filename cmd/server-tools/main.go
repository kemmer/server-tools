package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"server-tools/internal/functions"
	"strconv"
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

func stressTest(c *fiber.Ctx) error {
	log.Println("request incoming: stressTest()")

	p := c.AllParams()
	memorySizeGB, _ := p["memorySizeGB"]
	timeSeconds, _ := p["timeSeconds"]

	memorySizeGBInt, _ := strconv.Atoi(memorySizeGB)
	timeSecondsInt, _ := strconv.Atoi(timeSeconds)

	return c.SendString(functions.StressTest(memorySizeGBInt, timeSecondsInt))
}

func initializeApi() *fiber.App {
	app := fiber.New()

	app.Get("/", helloWorld)
	app.Get("/uptime", uptime)
	app.Get("/envs", envs)
	app.Get("/os", runningOs)
	app.Get("/backup-info", backupInfo)
	app.Get("/stress-test/:memorySizeGB/:timeSeconds", stressTest)

	return app
}

func main() {
	app := initializeApi()

	err := app.Listen(fmt.Sprintf(":%d", ServerPort))
	if err != nil {
		log.Fatalf(fmt.Sprintf("could not initialize server on port %d", ServerPort))
	}
}
