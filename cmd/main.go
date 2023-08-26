package main

import (
	"net/http"

	"github.com/AthanatiusC/pizza-hub/controller"
	"github.com/AthanatiusC/pizza-hub/repository"
	"github.com/AthanatiusC/pizza-hub/service"
	"github.com/AthanatiusC/pizza-hub/service/pizzahub"
)

func Handle(handler http.HandlerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
			return
		}
		handler(w, r)
	}
}

func main() {
	repository := repository.NewRepository()
	service := service.NewService(repository)
	controller := controller.NewController(service)

	http.HandleFunc("/menus", Handle(controller.GetMenus, http.MethodGet))
	http.HandleFunc("/chef", Handle(controller.AddChef, http.MethodPost))
	http.HandleFunc("/order", Handle(controller.AddOrder, http.MethodPost))
	http.HandleFunc("/order/info", Handle(controller.GetOrder, http.MethodGet))

	pizzaService := pizzahub.NewBakeryService(repository)
	go pizzaService.StartPizzaFactory()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
