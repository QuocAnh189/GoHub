package main

import (
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	httpServer "gohub/internal/server/http"

	"github.com/QuocAnh189/GoBin/logger"
	"github.com/QuocAnh189/GoBin/validation"

	"gohub/configs"
	"gohub/database"
	// "gohub/database/migrations"
)

//	@title          EventHub (GoHub) Swagger API
//	@version		OAS 3.0
//	@description	Swagger API for EventHub.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Tran Phuoc Anh Quoc
//	@contact.email	anhquoctpdev@gmail.com

//	@license.name	MIT
//	@license.url	https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
const (
	key    = "randomString"
	MaxAge = 86400 * 30
	IsProd = false
)

func main() {
	cfg := configs.LoadConfig(".")

	logger.Initialize(cfg.Environment)

	db, err := database.NewDatabase(cfg.DatabaseURI)
	if err != nil {
		logger.Fatal("Cannot connect to database", err)
	}

	//err = migrations.AutoMigrate(db)
	//if err != nil {
	//	logger.Fatal("Cannot migrate database", err)
	//}

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProd

	gothic.Store = store

	goth.UseProviders(
		google.New(
			cfg.GoogleClientID,
			cfg.GoogleClientSecret,
			"http://localhost:8888/api/v1/auth/external-auth-callback",
		),
	)

	validator := validation.New()

	//go func() {
	//	httpSvr := httpServer.NewServer(validator, db)
	//	if err = httpSvr.Run(); err != nil {
	//		logger.Fatal(err)
	//	}
	//}()
	//
	//grpcSvr := grpcServer.NewServer(validator, db)
	//if err = grpcSvr.Run(); err != nil {
	//	logger.Fatal(err)
	//}

	httpSvr := httpServer.NewServer(validator, db)
	if err = httpSvr.Run(); err != nil {
		logger.Fatal(err)
	}
}
