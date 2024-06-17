package service

import (
	"context"

	"github.com/matheusmhmelo/FullCycle-clean-arch/internal/infra/grpc/pb"
	"github.com/matheusmhmelo/FullCycle-clean-arch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrderUseCase   usecase.ListOrderUseCase
}

func NewOrderService(create usecase.CreateOrderUseCase, list usecase.ListOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: create,
		ListOrderUseCase:   list,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrder(ctx context.Context, in *pb.ListOrderRequest) (*pb.ListOrderResponse, error) {
	output, err := s.ListOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}

	orders := []*pb.OrderContent{}
	for _, order := range output {
		orders = append(orders, &pb.OrderContent{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		})
	}
	return &pb.ListOrderResponse{
		Orders: orders,
	}, nil
}
