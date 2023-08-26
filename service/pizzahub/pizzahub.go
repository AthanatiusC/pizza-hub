package pizzahub

import (
	"context"
	"sync"
	"time"

	"github.com/AthanatiusC/pizza-hub/helper/logger"
	"github.com/AthanatiusC/pizza-hub/model"
	"github.com/AthanatiusC/pizza-hub/repository"
)

type BakeryService struct {
	repository *repository.Repository
}

func NewBakeryService(repository *repository.Repository) *BakeryService {
	return &BakeryService{
		repository: repository,
	}
}
func (b *BakeryService) StartPizzaFactory() {
	ctx := context.WithValue(context.Background(), model.RequestIDKey, "Pizza Factory")
	logger.InfoContext(ctx, "Pizza factory started!")
	for {
		orders := b.repository.GetOrders()
		if len(orders) > 0 {
			chefs := b.repository.GetChefs()
			if len(chefs) == 0 {
				time.Sleep(5 * time.Second)
				continue
			}

			var wg sync.WaitGroup
			wg.Add(len(chefs))
			for _, chef := range chefs {
				go func(chef model.Chef) {
					defer wg.Done()
					if len(orders) > 0 {
						order := orders[0]
						orders = orders[1:]
						menu := b.repository.GetMenuByID(order.MenuID)
						logger.InfoContext(ctx, chef.Name, "is cooking order id", order.ID)
						time.Sleep(menu.Duration) // simulate cooking time
						order.Status = model.OrderStatusDone
						b.repository.PlaceOrder(order)
						logger.InfoContext(ctx, chef.Name, "has finished cooking order id", order.ID, ". Order took", time.Since(order.CreatedAt).Seconds(), "second to complete")
					}
				}(chef)
			}
			wg.Wait()
		}
	}
}
