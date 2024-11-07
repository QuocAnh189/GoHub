package main

import (
	httpServer "gohub/internal/server/http"

	"github.com/QuocAnh189/GoBin/logger"
	"github.com/QuocAnh189/GoBin/validation"

	"gohub/configs"
	"gohub/database"
	"gohub/database/migrations"
)

//	@title          EventHub (GoHub) Swagger API
//	@version		OAS 3.0
//	@description	Swagger API for EventHub.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Tran Phuoc Anh Quoc
//	@contact.email	anhquoctpdev@gmail.com

//	@license.name	MIT
//	@license.url	https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

func main() {
    cfg := configs.LoadConfig()

    logger.Initialize(cfg.Environment)

    db, err := database.NewDatabase(cfg.DatabaseURI)
    if err != nil {
        logger.Fatal("Cannot connect to database", err)
    }


    err = migrations.AutoMigrate(db)
    if err != nil {
        logger.Fatal("Cannot migrate database", err)
    }

    validator := validation.New()

    httpServer := httpServer.NewServer(validator, db)

    if err := httpServer.Run(); err != nil {
        logger.Fatal("HTTP server failed to start", err)
    }
}
