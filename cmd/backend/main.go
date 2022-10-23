package main

import (
	"github.com/hoffme/ddd-backend/internal/apps"
	"github.com/hoffme/ddd-backend/internal/configs"
	"github.com/hoffme/ddd-backend/internal/contexts"
	"github.com/hoffme/ddd-backend/internal/shared/bus"
)

func main() {
	config := configs.LoadConfigFromJSON("./config.json")

	buses := bus.Build()

	contexts.Setup(config.Contexts, buses)
	apps.Setup(config.Apps, buses)
}
