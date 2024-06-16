package domain

type VideoRepository interface {
	SearchVideosByString(serchString string, sortColumn string, sortOrder string, page int) (*[]Video, error)
}
