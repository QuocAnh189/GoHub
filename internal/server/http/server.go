package server

import (
	"fmt"
	"log"
	"net/http"

	_ "gohub/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	authHttp "gohub/domains/auth/port/http"
	categoryHttp "gohub/domains/categories/port/http"
	commandHttp "gohub/domains/commands/port/http"
	conversationHttp "gohub/domains/conversations/port/http"
	eventHttp "gohub/domains/events/port/http"
	functionHttp "gohub/domains/functions/port/http"
	reviewHttp "gohub/domains/reviews/port/http"
	routeHttp "gohub/domains/roles/port/http"
	userHttp "gohub/domains/users/port/http"

	"github.com/QuocAnh189/GoBin/logger"
	"github.com/QuocAnh189/GoBin/validation"
	"github.com/gin-gonic/gin"

	"gohub/configs"
	"gohub/database"
	"gohub/pkg/response"
)

type Server struct {
	engine    *gin.Engine
	cfg       *configs.Schema
	validator validation.Validation
	db        database.IDatabase
}

func NewServer(validator validation.Validation, db database.IDatabase) *Server {
	return &Server{
		engine:    gin.Default(),
		cfg:       configs.GetConfig(),
		validator: validator,
		db:        db,
	}
}

func (s Server) Run() error {
	_ = s.engine.SetTrustedProxies(nil)
	if s.cfg.Environment == configs.ProductionEnv {
		gin.SetMode(gin.ReleaseMode)
	}

	if err := s.MapRoutes(); err != nil {
		log.Fatalf("MapRoutes Error: %v", err)
	}

	s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.engine.GET("/", func(c *gin.Context) {
		response.JSON(c, http.StatusOK, gin.H{"message": "Welcome to GoHub API"})
	})

	// Start http server
	logger.Info("HTTP server is listening on PORT: ", s.cfg.HttpPort)
	if err := s.engine.Run(fmt.Sprintf(":%d", s.cfg.HttpPort)); err != nil {
		log.Fatalf("Running HTTP server: %v", err)
	}

	return nil
}

func (s Server) GetEngine() *gin.Engine {
	return s.engine
}

func (s Server) MapRoutes() error {
	routes_v1 := s.engine.Group("/api/v1")
	authHttp.Routes(routes_v1, s.db, s.validator)
	userHttp.Routes(routes_v1, s.db, s.validator)
	reviewHttp.Routes(routes_v1, s.db, s.validator)
	conversationHttp.Routes(routes_v1, s.db, s.validator)
	categoryHttp.Routes(routes_v1, s.db, s.validator)
	eventHttp.Routes(routes_v1, s.db, s.validator)
	routeHttp.Routes(routes_v1, s.db, s.validator)
	functionHttp.Routes(routes_v1, s.db, s.validator)
	commandHttp.Routes(routes_v1, s.db, s.validator)
	return nil
}