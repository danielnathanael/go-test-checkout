// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Cart struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	ItemID    string `json:"item_id"`
	Quantity  int    `json:"quantity"`
	UpdatedAt string `json:"updated_at"`
}

type Item struct {
	Sku       string  `json:"sku"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt string  `json:"deleted_at"`
}

type NewItem struct {
	Sku      string  `json:"sku"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Order struct {
	ID            string `json:"id"`
	UserID        string `json:"user_id"`
	CreatedAt     string `json:"created_at"`
	PaidAt        string `json:"paid_at"`
	PaymentMethod string `json:"payment_method"`
}

type OrderDetail struct {
	OrderID  string  `json:"order_id"`
	ItemID   string  `json:"item_id"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type User struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Carts    []*Cart `json:"carts"`
}
