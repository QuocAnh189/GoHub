package main

import (
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	socketioServer "gohub/internal/libs/websocket"
	httpServer "gohub/internal/server/http"
	"log"
	"sync"

	"gohub/internal/libs/logger"
	"gohub/internal/libs/validation"

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

	// Initialize HTTP server
	httpSvr := httpServer.NewServer(validator, db)

	// Initialize Socket.IO server
	socketSvr, err := socketioServer.NewServer()
	if err != nil {
		logger.Fatal("Cannot initialize Socket.IO server", err)
	}

	// Run both servers in separate goroutines
	var wg sync.WaitGroup
	wg.Add(2)

	// Run HTTP server
	go func() {
		defer wg.Done()
		if err := httpSvr.Run(); err != nil {
			logger.Fatal("Running HTTP server error:", err)
		}
	}()

	// Run Socket.IO server on port 9000
	go func() {
		defer wg.Done()
		if err := socketSvr.Run(cfg.SocketPort); err != nil {
			log.Fatalf("Socket.IO server error: %v", err)
		}
	}()

	wg.Wait()
}
