package usecase

import (
	"github.com/Dionizio8/goexpert-clean-arch/internal/entity"
)

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(OrderRepository entity.OrderRepositoryInterface) ListOrderUseCase {
	return ListOrderUseCase{OrderRepository: OrderRepository}
}

func (l *ListOrderUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := l.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var ordersDTO []OrderOutputDTO
	for _, order := range orders {
		orderDTO := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}

		ordersDTO = append(ordersDTO, orderDTO)
	}

	return ordersDTO, nil
}
