package server

import (
	"fmt"
	"log"
	"net/http"

	_ "gohub/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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
	routes := s.engine.Group("/api/v1")
	{
		routes.GET("/ping", func(c *gin.Context) {
			response.JSON(c, http.StatusOK, gin.H{"message": "Love you change"})
		})
	}
	return nil
}