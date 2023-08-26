package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/AthanatiusC/pizza-hub/helper"
	"github.com/AthanatiusC/pizza-hub/helper/logger"
	"github.com/AthanatiusC/pizza-hub/model"
	"github.com/AthanatiusC/pizza-hub/repository"
)

type Service struct {
	Repository *repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	ctx := context.WithValue(context.Background(), model.RequestIDKey, "Repository")
	rand.Seed(time.Now().UnixNano()) // initialize global pseudo random generator
	logger.InfoContext(ctx, "Service initialized")
	return &Service{
		Repository: repository,
	}
}

func (s *Service) GetMenus(ctx context.Context) ([]model.Menu, error) {
	return s.Repository.GetMenus(), nil
}

func (s *Service) AddChef(ctx context.Context, request model.Chef) (model.Chef, error) {
	id := helper.GenerateRandomUniqueID()
	if ctx.Value("request_id") != nil {
		id = ctx.Value("request_id").(string)
	}
	logger.InfoContext(ctx, "Chef added!", request)
	return s.Repository.AddChefs(model.Chef{
		ID:   id,
		Name: request.Name,
	}), nil
}

func (s *Service) AddOrder(ctx context.Context, request model.OrderRequest) (model.Order, error) {
	menu := s.Repository.GetMenuByID(request.MenuID)
	if menu.ID == 0 {
		return model.Order{}, fmt.Errorf("menu_id is required")
	}

	id := helper.GenerateRandomUniqueID()
	if ctx.Value("request_id") != nil {
		id = ctx.Value("request_id").(string)
	}

	response := s.Repository.PlaceOrder(model.Order{
		ID:        id,
		MenuID:    request.MenuID,
		Status:    model.OrderStatusPending,
		CreatedAt: time.Now(),
	})
	logger.InfoContext(ctx, "Order placed!", response)
	return response, nil
}

func (s *Service) GetOrder(ctx context.Context, id string) (model.Order, error) {
	order := s.Repository.GetOrderByID(id)
	if order.ID == "" {
		return model.Order{}, fmt.Errorf("order not found")
	}

	return order, nil
}
