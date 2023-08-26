package service

import (
	"context"
	"testing"

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
