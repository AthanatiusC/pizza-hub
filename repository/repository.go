package repository

import (
	"sort"
	"sync"
	"time"

	"github.com/AthanatiusC/pizza-hub/helper"
	"github.com/AthanatiusC/pizza-hub/model"
)

type Repository struct {
	menus  []model.Menu
	chefs  []model.Chef
	orders map[string]model.Order
	Lock   sync.Mutex
}

func NewRepository() *Repository {
	menus := []model.Menu{
		{
			ID:       1,
			Name:     "Pizza Cheese",
			Duration: time.Duration(3) * time.Second,
		},
		{
			ID:       2,
			Name:     "Pizza BBQ",
			Duration: time.Duration(5) * time.Second,
		},
	}

	chefs := []model.Chef{
		{
			ID:   helper.GenerateRandomUniqueID(),
			Name: "Bjorn Ironside",
		},
	}

	return &Repository{
		menus:  menus,
		chefs:  chefs,
		orders: make(map[string]model.Order),
	}
}

func (r *Repository) GetMenus() []model.Menu {
	return r.menus
}

func (r *Repository) GetMenuByID(id int) model.Menu {
	for _, menu := range r.menus {
		if menu.ID == id {
			return menu
		}
	}
	return model.Menu{}
}

func (r *Repository) AddChefs(request model.Chef) model.Chef {
	r.chefs = append(r.chefs, request)
	return request
}

func (r *Repository) GetChefs() []model.Chef {
	return r.chefs
}

func (r *Repository) PlaceOrder(request model.Order) model.Order {
	r.Lock.Lock()
	defer r.Lock.Unlock()
	r.orders[request.ID] = request
	return request
}

func (r *Repository) GetOrderByID(id string) model.Order {
	r.Lock.Lock()
	defer r.Lock.Unlock()
	return r.orders[id]
}

func (r *Repository) GetOrders() []model.Order {
	r.Lock.Lock()
	defer r.Lock.Unlock()
	var orders []model.Order
	for _, order := range r.orders {
		if order.Status != model.OrderStatusDone {
			orders = append(orders, order)
		}
	}

	sort.Slice(orders, func(i, j int) bool {
		return orders[i].CreatedAt.Before(orders[j].CreatedAt)
	})
	return orders
}
