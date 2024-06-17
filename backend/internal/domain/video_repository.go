package domain

type VideoRepository interface {
	SearchVideosByString(searchString string, sortColumn string, sortOrder string, page int) (*[]Video, error)
}
