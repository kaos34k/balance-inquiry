package usecase

import (
	"balance-inquiry/internal/domain"
)

type PointUsecase struct {
	pointRepository domain.PointRepository
}

func NewPointUsecase(pointRepository domain.PointRepository) *PointUsecase {
	return &PointUsecase{
		pointRepository: pointRepository,
	}
}

func (u *PointUsecase) GetPointByUser(id string) (*[]domain.Point, error) {
	return u.pointRepository.GetPointByUser(id)
}
