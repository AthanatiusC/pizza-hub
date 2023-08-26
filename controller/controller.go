package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/AthanatiusC/pizza-hub/helper"
	"github.com/AthanatiusC/pizza-hub/helper/logger"
	"github.com/AthanatiusC/pizza-hub/helper/response"
	"github.com/AthanatiusC/pizza-hub/model"
	"github.com/AthanatiusC/pizza-hub/service"
)

type Controller struct {
	Service *service.Service
}

func NewController(service *service.Service) *Controller {
	ctx := context.WithValue(context.Background(), model.RequestIDKey, "Controller")
	logger.InfoContext(ctx, "Controller initialized")
	return &Controller{
		Service: service,
	}
}

func (c *Controller) GetMenus(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), model.RequestIDKey, helper.GenerateRandomUniqueID())
	c.Service.GetMenus(ctx)
}

func (c *Controller) AddChef(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), model.RequestIDKey, helper.GenerateRandomUniqueID())
	var request model.Chef
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Error(ctx, w, http.StatusBadRequest, err.Error())
	}

	c.Service.AddChef(ctx, request)
	response.Success(ctx, w, request)
}

func (c *Controller) AddOrder(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), model.RequestIDKey, helper.GenerateRandomUniqueID())
	var request model.OrderRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Error(r.Context(), w, http.StatusBadRequest, err.Error())
	}

	result, err := c.Service.AddOrder(ctx, request)
	if err != nil {
		response.Error(ctx, w, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(ctx, w, result)
}

func (c *Controller) GetOrder(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), model.RequestIDKey, helper.GenerateRandomUniqueID())
	id := r.URL.Query().Get("id")

	result, err := c.Service.GetOrder(r.Context(), id)
	if err != nil {
		response.Error(r.Context(), w, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(ctx, w, result)
}
