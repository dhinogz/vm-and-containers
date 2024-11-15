package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		fmt.Println("Error opening /etc/os-release:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var distro string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			distro = strings.TrimPrefix(line, "PRETTY_NAME=")
			distro = strings.Trim(distro, `"`)
			break
		}
	}

	res := fmt.Sprintf("running on %s", distro)

	slog.Info(res, "equipo", 2, "integrantes", []string{"David", "Sasha", "Ingrid", "Paulina", "Pablo"})
}
