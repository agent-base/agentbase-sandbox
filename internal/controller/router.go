package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/agent-base/agentbase-sandbox/internal/middleware"
	"github.com/agent-base/agentbase-sandbox/internal/static"
	"net/http"
)

func Setup(Router *gin.Engine) {
	PublicGroup := Router.Group("")
	PrivateGroup := Router.Group("/v1/sandbox/")

	PrivateGroup.Use(middleware.Auth())

	{
		// health check
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	InitRunRouter(PrivateGroup)
	InitDependencyRouter(PrivateGroup)
}

func InitDependencyRouter(Router *gin.RouterGroup) {
	dependencyRouter := Router.Group("dependencies")
	{
		dependencyRouter.GET("", GetDependencies)
		dependencyRouter.POST("update", UpdateDependencies)
		dependencyRouter.GET("refresh", RefreshDependencies)
	}
}

func InitRunRouter(Router *gin.RouterGroup) {
	runRouter := Router.Group("")
	{
		runRouter.POST(
			"run",
			middleware.MaxRequest(static.GetAgentBaseSandboxGlobalConfigurations().MaxRequests),
			middleware.MaxWorker(static.GetAgentBaseSandboxGlobalConfigurations().MaxWorkers),
			RunSandboxController,
		)
	}
}
