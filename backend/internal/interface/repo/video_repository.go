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

func (r *videoRepository) SearchVideosByString(searchString string, sortColumn string, sortOrder string, page int) (*[]domain.Video, error) {
	const pageSize = 50
	var videos []domain.Video
	var err error

	if searchString == "" {
		// if no search string, return all videos
		err = r.db.Order(sortColumn + " " + sortOrder).Limit(pageSize).Offset((page - 1) * pageSize).Find(&videos).Error
	} else {
		// search by string
		err = r.db.Where("MATCH (title, description) AGAINST (? IN BOOLEAN MODE)", searchString).Order(sortColumn + " " + sortOrder).Limit(pageSize).Offset((page - 1) * pageSize).Find(&videos).Error
	}

	if err != nil {
		config.Logger.Error(err.Error())
		return nil, err
	}
	return &videos, nil
}
