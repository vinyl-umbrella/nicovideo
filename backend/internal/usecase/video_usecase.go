package usecase

import "nicovideo-api/internal/domain"

type VideoUsecase interface {
	SearchVideos(serchString string, sortColumn string, sortOrder string, page int) (*[]domain.Video, error)
}
type videoUsecase struct {
	videoRepository domain.VideoRepository
}

func NewVideoUsecase(repo domain.VideoRepository) VideoUsecase {
	return &videoUsecase{videoRepository: repo}
}

func (v *videoUsecase) SearchVideos(searchString string, sortColumn string, sortOrder string, page int) (*[]domain.Video, error) {
	return v.videoRepository.SearchVideosByString(searchString, sortColumn, sortOrder, page)
}
