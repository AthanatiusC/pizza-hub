package repository

import (
	"testing"

	"github.com/AthanatiusC/pizza-hub/helper"
	"github.com/AthanatiusC/pizza-hub/model"
)

func TestGetMenus(t *testing.T) {
	repository := NewRepository()
	response := repository.GetMenus()
	if len(response) <= 0 {
		t.Errorf("Should return at least 2, got %d", len(response))
	}
}

func TestGetMenuByID(t *testing.T) {
	repository := NewRepository()
	response := repository.GetMenuByID(1)
	if response.ID != 1 {
		t.Errorf("Expected exist, got %d", response.ID)
	}
}

func TestAddChef(t *testing.T) {
	repository := NewRepository()
	request := model.Chef{
		ID:   helper.GenerateRandomUniqueID(),
		Name: "Bjorn Ironside",
	}
	response := repository.AddChefs(request)
	if response.Name != request.Name {
		t.Errorf("Expected %s, got %s", request.Name, response.Name)
	}
	if response.ID != request.ID {
		t.Errorf("Expected %s, got %s", request.ID, response.ID)
	}
}

func TestGetChefs(t *testing.T) {
	repository := NewRepository()
	response := repository.GetChefs()
	if len(response) <= 0 {
		t.Errorf("Should return at least 1, got %d", len(response))
	}
}

func TestPlaceOrder(t *testing.T) {
	repository := NewRepository()
	for i := 0; i < 100; i++ {
		response := repository.PlaceOrder(model.Order{ID: helper.GenerateRandomUniqueID(), MenuID: 1, Status: model.OrderStatusPending})
		if response.ID == "" {
			t.Errorf("Expected success, got empty id")
		}
	}
}

func TestGetOrderByID(t *testing.T) {
	repository := NewRepository()
	order := repository.PlaceOrder(model.Order{ID: helper.GenerateRandomUniqueID(), MenuID: 1, Status: model.OrderStatusPending})
	response := repository.GetOrderByID(order.ID)
	if response.ID != order.ID {
		t.Errorf("Expected %s, got %s", order.ID, response.ID)
	}
}

func TestGetOrders(t *testing.T) {
	repository := NewRepository()
	response := repository.GetOrders()
	if len(response) > 0 {
		t.Errorf("Should return 0, got %d", len(response))
	}
}
