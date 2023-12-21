package router

import (
	"ms-briapi-log-stream/config"
	"ms-briapi-log-stream/controllers"
	"ms-briapi-log-stream/middlewares"
	"ms-briapi-log-stream/repo"
	"ms-briapi-log-stream/usecase"
	"ms-briapi-log-stream/usecase/error"
	"time"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.New()

	db := config.ConnectDB()
	v1 := router.Group("v1")
	{
		newRoute := v1.Group("ms-briapi-log-stream")
		errorsUsecase := error.NewErrorHandlerUsecase()

		logStreamRepo := repo.CreateLogStreamRepoImpl(db)
		logStreamUsecase := usecase.CreateLogStreamUsecaseImpl(logStreamRepo)

		middlewares.NewErrorHandler(v1, errorsUsecase)
		router.Use(middlewares.TimeoutMiddleware(50 * time.Second))
		controllers.CreateLogStreamController(newRoute, logStreamUsecase, errorsUsecase)
	}
	return router
}
