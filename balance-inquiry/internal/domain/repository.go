package domain

type PointRepository interface {
	GetPointByUser(id string) (*[]Point, error)
}
