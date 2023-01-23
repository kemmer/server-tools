package functions

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Uptime() string {
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

	return fmt.Sprintf("uptime: %q\n", out.String())
}

func Envs() string {
	vars := os.Environ()

	return strings.Join(vars, "\n")
}

func RunningOs() string {
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

	return strings.Join(osAttr, "\n")
}
