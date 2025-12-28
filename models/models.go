package models

type Client struct {
	ClientID int    `json:"client_id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

type Car struct {
	CarID    int    `json:"car_id"`
	Brand    string `json:"brand"`
	Model    string `json:"model"`
	Year     int    `json:"year"`
	Plate    string `json:"plate"`
	ClientID int    `json:"client_id"`
}

type Master struct {
	MasterID       int    `json:"master_id"`
	Name           string `json:"name"`
	Specialization string `json:"specialization"`
	Phone          string `json:"phone"`
}

type Service struct {
	ServiceID int     `json:"service_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Duration  float64 `json:"duration"`
}

type Order struct {
	OrderID   int    `json:"order_id"`
	CarID     int    `json:"car_id"`
	MasterID  int    `json:"master_id"`
	ServiceID int    `json:"service_id"`
	OrderDate string `json:"order_date"`
	Status    string `json:"status"`
}

type ExportData struct {
	Clients  []Client  `json:"clients"`
	Cars     []Car     `json:"cars"`
	Masters  []Master  `json:"masters"`
	Services []Service `json:"services"`
	Orders   []Order   `json:"orders"`
}
