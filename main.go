package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const ServerPort = 7845

func helloWorld(c *fiber.Ctx) error {
	log.Println("request incoming: helloWorld()")

	randomNumber := rand.Int() % 10000001
	return c.SendString(fmt.Sprintf("Hello, world! 👋 (ID: %d)", randomNumber))
}

func uptime(c *fiber.Ctx) error {
	log.Println("request incoming: uptime()")

	cmd := exec.Command("uptime")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if errors.Is(err, exec.ErrDot) {
		log.Fatalln("executable found but missing complete path so it won't run")
	}
	if err != nil {
		log.Fatalln(err)
	}

	return c.SendString(fmt.Sprintf("uptime: %q\n", out.String()))
}

func envs(c *fiber.Ctx) error {
	log.Println("request incoming: envs()")

	vars := os.Environ()

	return c.SendString(strings.Join(vars, "\n"))
}

func runningOs(c *fiber.Ctx) error {
	log.Println("request incoming: runningOs()")

	var osName string
	switch runtime.GOOS {
	case "darwin":
		osName = "MacOS"
	default:
		osName = runtime.GOOS
	}

	var osAttr []string
	osAttr = append(osAttr, fmt.Sprintf("OS: %s", osName))
	osAttr = append(osAttr, fmt.Sprintf("arch: %s", runtime.GOARCH))

	return c.SendString(strings.Join(osAttr, "\n"))
}

func backupInfo(c *fiber.Ctx) error {
	log.Println("request incoming: backupInfo()")

	if _, err := os.Stat("./backup"); errors.Is(err, fs.ErrNotExist) {
		err = os.Mkdir("backup", 0755)
		if err != nil && !errors.Is(err, fs.ErrExist) {
			log.Println(err)
			return c.SendString("cannot create folder './backup'")
		}

		return c.SendString("latest backup: never")
	}

	d, err := os.Open("./backup")
	if err != nil {
		log.Println(err)
		return c.SendString("cannot open folder './backup'")
	}

	files, err := d.ReadDir(0)
	if err != nil {
		log.Println(err)
		return c.SendString("cannot list files from folder './backup'")
	}

	var backupList []string
	backupList = append(backupList, "backup list:")

	for _, f := range files {
		if !f.IsDir() {
			backupList = append(backupList, f.Name())
		}
	}

	if len(backupList) == 1 {
		return c.SendString("latest backup: never")
	}

	return c.SendString(strings.Join(backupList, "\n"))
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
