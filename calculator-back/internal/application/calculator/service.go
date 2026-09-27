package calculator

import (
	domaincalculator "github.com/Carlosam7/calculator/internal/domain/services/calculator"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Evaluate(expression string) (float64, error) {
	return domaincalculator.Evaluate(expression)
}
