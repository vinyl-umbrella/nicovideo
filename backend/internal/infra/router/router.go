package router

import (
	"nicovideo-api/internal/interface/handler"
	"nicovideo-api/internal/interface/repo"
	"nicovideo-api/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	videoRepo := repo.NewVideoRepository(db)
	videoUsecase := usecase.NewVideoUsecase(videoRepo)
	videoHandler := handler.NewVideoHandler(videoUsecase)

	r.GET("/api/v1/search", videoHandler.SearchVideos)

	return r
}
