package app

import (
	"balance-inquiry/internal/domain"
	"balance-inquiry/internal/usecase"
)

type MyApp struct {
	pointUsecase usecase.PointUsecase
}

func NewMyApp(pointUsecase usecase.PointUsecase) *MyApp {
	return &MyApp{
		pointUsecase: pointUsecase,
	}
}

func (a *MyApp) HandleRequest(id string) (*[]domain.Point, error) {
	points, err := a.pointUsecase.GetPointByUser(id)
	if err != nil {
		return nil, err
	}

	return points, nil
}
