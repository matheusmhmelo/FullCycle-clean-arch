package usecase

import (
	"github.com/matheusmhmelo/FullCycle-clean-arch/internal/entity"
)

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (l *ListOrderUseCase) Execute() ([]OrderOutputDTO, error) {
	results, err := l.OrderRepository.List()
	if err != nil {
		return nil, err
	}

	dto := []OrderOutputDTO{}
	for _, result := range results {
		dto = append(dto, OrderOutputDTO{
			ID:         result.ID,
			Price:      result.Price,
			Tax:        result.Tax,
			FinalPrice: result.FinalPrice,
		})
	}

	return dto, nil
}
