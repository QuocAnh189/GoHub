package main

import (
	httpServer "gohub/internal/server/http"

	"github.com/QuocAnh189/GoBin/logger"
	"github.com/QuocAnh189/GoBin/validation"

	"gohub/configs"
	"gohub/database"
)

func main() {
    cfg := configs.LoadConfig()

    logger.Initialize(cfg.Environment)

    db, err := database.NewDatabase(cfg.DatabaseURI)
    if err != nil {
        logger.Fatal("Cannot connect to database", err)
    }

    validator := validation.New()

    httpServer := httpServer.NewServer(validator, db)

    if err := httpServer.Run(); err != nil {
        logger.Fatal("HTTP server failed to start", err)
    }
}
