package handler

import (
	"gin_news/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	
	router.MaxMultipartMemory = 8 << 20

	router.LoadHTMLGlob("templates/**/*")
	router.Static("/static","./static")
	auth := router.Group("/auth")
	{
		auth.GET("/sign-up", h.signUp)
		auth.GET("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity) 
	{
		api.GET("/news", h.getAllNews)
		api.GET("/:id", h.getNewsById)
		changeNews := api.Group("/", h.checkAccess) 
		{
			changeNews.POST("/addnews", h.createNews)
			changeNews.GET("/addnews", h.createNews)
			changeNews.DELETE("/del_news", h.delete_news)
			changeNews.PUT("/update_news", h.update_news)
		}

	}

	return router
}
