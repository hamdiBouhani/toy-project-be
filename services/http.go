package services

import (
	"flag"
	"fmt"
	"time"
	"toy-project-be/common/context"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type HttpService struct {
	context.DefaultService
	router       *gin.Engine
	openAccessed *gin.RouterGroup
	port         string
	corsConfig   cors.Config
}

const HTTP_SERVICE = "http_base"

//Id for the HTTP service
func (svc HttpService) Id() string {
	return HTTP_SERVICE
}

func (svc *HttpService) Configure(ctx *context.Context) error {

	var port = flag.String("port", "8080", "Http port to serve application on")
	flag.Parse()

	svc.port = fmt.Sprintf(":%s", *port)

	svc.corsConfig = cors.Config{
		AllowOrigins: []string{"*"},
		//AllowOrigins:     []string{"http://localhost:8666", "http://localhost:8081"}, //Allow from everywhere for now
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	return svc.DefaultService.Configure(ctx)
}

func (svc *HttpService) Start() error {

	err := svc.registerRoutes()
	if err != nil {
		return err
	}

	return svc.router.Run(svc.port) //Blocks
}

func (svc *HttpService) registerRoutes() error {

	svc.router = gin.Default()
	svc.router.Use(cors.New(svc.corsConfig))

	//Define openAccessed route.
	svc.openAccessed = svc.router.Group("/")
	svc.openAccessed.GET("/ping", svc.ping) //PING/PONG

	return nil
}

func (svc *HttpService) ping(c *gin.Context) {
	c.JSON(200, "pong")
}
