package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "gohub/docs"
	"log"
	"net/http"
	"time"

	authHttp "gohub/domains/auth/port/http"
	categoryHttp "gohub/domains/categories/port/http"
	commandHttp "gohub/domains/commands/port/http"
	conversationHttp "gohub/domains/conversations/port/http"
	couponHttp "gohub/domains/coupons/port/http"
	eventHttp "gohub/domains/events/port/http"
	expenseHttp "gohub/domains/expense/port/http"
	functionHttp "gohub/domains/functions/port/http"
	permissionHttp "gohub/domains/permissions/port/http"
	reviewHttp "gohub/domains/reviews/port/http"
	routeHttp "gohub/domains/roles/port/http"
	statisticHttp "gohub/domains/statistic/port/http"
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
	cfg       *configs.Config
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

	s.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "access-control-allow-origin", "access-control-allow-headers"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
	routesV1 := s.engine.Group("/api/v1")
	authHttp.Routes(routesV1, s.db, s.validator)
	userHttp.Routes(routesV1, s.db, s.validator)
	reviewHttp.Routes(routesV1, s.db, s.validator)
	conversationHttp.Routes(routesV1, s.db, s.validator)
	categoryHttp.Routes(routesV1, s.db, s.validator)
	eventHttp.Routes(routesV1, s.db, s.validator)
	routeHttp.Routes(routesV1, s.db, s.validator)
	functionHttp.Routes(routesV1, s.db, s.validator)
	commandHttp.Routes(routesV1, s.db, s.validator)
	permissionHttp.Routes(routesV1, s.db, s.validator)
	couponHttp.Routes(routesV1, s.db, s.validator)
	expenseHttp.Routes(routesV1, s.db, s.validator)
	statisticHttp.Routes(routesV1, s.db, s.validator)
	return nil
}
