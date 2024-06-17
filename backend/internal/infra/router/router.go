package router

import (
	"nicovideo-api/internal/interface/handler"
	"nicovideo-api/internal/interface/repo"
	"nicovideo-api/internal/usecase"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// CORS
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "OPTIONS"},
		MaxAge:          12 * time.Hour,
	}))

	videoRepo := repo.NewVideoRepository(db)
	videoUsecase := usecase.NewVideoUsecase(videoRepo)
	videoHandler := handler.NewVideoHandler(videoUsecase)

	r.GET("/v1/search", videoHandler.SearchVideos)

	return r
}
