package main

import (
	_ "github.com/lib/pq"

	"github.com/scmbr/vk-gamejam/backend/internal/app"
)

const configsDir = "configs"

func main() {
	app.Run(configsDir)
}
