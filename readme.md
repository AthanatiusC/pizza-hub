# Pizza HUB
Pizza Menu:
Pizza Cheese
Duration to process: 3 second
Pizza BBQ
Duration to process: 5 second

Because we run PizzaHub in Silicon Valley, we want our PizzaHub scalable and can process many orders simultaneously. 1 chef only can take a single order at one time. so we can recruit many chef to help us process the order.

API requirement:

POST /chef
add new chef

GET /menus
list of menus

POST /orders
add new orders

Build system in API, **No Dependencies except for HTTP Library.**
Criteria:
Performance
Scalability
Clean Code
Test

No dependencies mean:
- No framework
- No ORM
- No library
- No package manager
- No build tools
- No testing library
- No database

Run the program by executing the following command:
```
go run cmd/main.go
```
