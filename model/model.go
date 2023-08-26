package model

import "time"

var LetterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Context string

func (c Context) RequestID() string {
	return string(c)
}

const (
	OrderStatusPending = "pending"
	OrderStatusCooking = "cooking"
	OrderStatusDone    = "done"
)

var (
	RequestIDKey = Context("request_id")
)

type OrderRequest struct {
	MenuID int `json:"menu_id"`
}

type Menu struct {
	ID       int           `json:"id"`
	Name     string        `json:"name"`
	Duration time.Duration `json:"duration"`
}

type Chef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Order struct {
	ID        string    `json:"id"`
	MenuID    int       `json:"menu_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type Response struct {
	StatusCode int         `json:"-"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
}
