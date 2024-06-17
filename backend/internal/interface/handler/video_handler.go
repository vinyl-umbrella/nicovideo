package handler

import (
	"net/http"
	"nicovideo-api/config"
	"nicovideo-api/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VideoHandler struct {
	videoUsecase usecase.VideoUsecase
}

func NewVideoHandler(u usecase.VideoUsecase) *VideoHandler {
	return &VideoHandler{videoUsecase: u}
}

func (h *VideoHandler) SearchVideos(c *gin.Context) {
	// GET /api/v1/search?q=keyword&order=ViewCountAsc&page=1

	// validate q
	searchString := c.Query("q")

	// validate order
	order := c.Query("order")
	sortColumName := ""
	sortOrder := ""

	switch order {
	case "":
		sortColumName = "vid"
		sortOrder = "asc"
	case "ViewCountAsc":
		sortColumName = "view_count"
		sortOrder = "asc"
	case "ViewCountDesc":
		sortColumName = "view_count"
		sortOrder = "desc"
	case "RegisteredAtAsc":
		sortColumName = "vid"
		sortOrder = "asc"
	case "RegisteredAtDesc":
		sortColumName = "vid"
		sortOrder = "desc"
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order"})
		return
	}

	// validate page
	page := c.Query("page")
	// default page is 1
	if page == "" {
		page = "1"
	}
	// string to int
	pageNum, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page"})
		config.Logger.Error(err.Error())
		return
	}

	video, err := h.videoUsecase.SearchVideos(searchString, sortColumName, sortOrder, pageNum)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		config.Logger.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, video)
}
