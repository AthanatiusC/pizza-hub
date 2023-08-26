package service

import (
	"context"
	"testing"

	"github.com/AthanatiusC/pizza-hub/model"
	"github.com/AthanatiusC/pizza-hub/repository"
)

func TestGetMenus(t *testing.T) {
	repository := repository.NewRepository()
	service := NewService(repository)

	response, err := service.GetMenus(context.TODO())
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if response == nil {
		t.Errorf("Expected not nil, got nil")
	}
}

func TestAddChef(t *testing.T) {
	repository := repository.NewRepository()
	service := NewService(repository)

	response, err := service.AddChef(context.TODO(), model.Chef{
		Name: "Test Chef",
	})
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if response.ID == "" {
		t.Errorf("Expected id not nil, got nil")
	}
}

func TestAddOrder(t *testing.T) {
	repository := repository.NewRepository()
	service := NewService(repository)

	response, err := service.AddOrder(context.TODO(), model.OrderRequest{
		MenuID: 1,
	})
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if response.ID == "" {
		t.Errorf("Expected id not nil, got nil")
	}
}

func TestGetOrder(t *testing.T) {
	repository := repository.NewRepository()
	service := NewService(repository)

	order, err := service.AddOrder(context.TODO(), model.OrderRequest{
		MenuID: 1,
	})
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	response, err := service.GetOrder(context.TODO(), order.ID)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if response.ID != order.ID {
		t.Errorf("Expected id %s, got %s", order.ID, response.ID)
	}
}
