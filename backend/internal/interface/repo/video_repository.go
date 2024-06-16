package repo

import (
	"nicovideo-api/config"
	"nicovideo-api/internal/domain"

	"gorm.io/gorm"
)

type videoRepository struct {
	db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) domain.VideoRepository {
	return &videoRepository{db}
}

func (r *videoRepository) SearchVideosByString(serchString string, sortColumn string, sortOrder string, page int) (*[]domain.Video, error) {
	var videos []domain.Video
	err := r.db.Where("match (title) against (?)", serchString).Order(sortColumn + " " + sortOrder).Limit(20).Offset((page - 1) * 20).Find(&videos).Error

	if err != nil {
		config.Logger.Error(err.Error())
		return nil, err
	}
	return &videos, nil
}
