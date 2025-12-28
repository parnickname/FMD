package json_io

import (
	"database/sql"
	"encoding/json"
	"os"

	"autoservice/models"
)

func ExportToJSON(db *sql.DB, filename string) error {
	data := models.ExportData{}

	// Загружаем клиентов
	rows, err := db.Query("SELECT client_id, name, phone, email FROM clients")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var c models.Client
		if err := rows.Scan(&c.ClientID, &c.Name, &c.Phone, &c.Email); err != nil {
			return err
		}
		data.Clients = append(data.Clients, c)
	}

	// Загружаем автомобили
	rows, err = db.Query("SELECT car_id, brand, model, year, plate, client_id FROM cars")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var c models.Car
		if err := rows.Scan(&c.CarID, &c.Brand, &c.Model, &c.Year, &c.Plate, &c.ClientID); err != nil {
			return err
		}
		data.Cars = append(data.Cars, c)
	}

	// Загружаем мастеров
	rows, err = db.Query("SELECT master_id, name, specialization, phone FROM masters")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var m models.Master
		if err := rows.Scan(&m.MasterID, &m.Name, &m.Specialization, &m.Phone); err != nil {
			return err
		}
		data.Masters = append(data.Masters, m)
	}

	// Загружаем услуги
	rows, err = db.Query("SELECT service_id, name, price, duration FROM services")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var s models.Service
		if err := rows.Scan(&s.ServiceID, &s.Name, &s.Price, &s.Duration); err != nil {
			return err
		}
		data.Services = append(data.Services, s)
	}

	// Загружаем заказы
	rows, err = db.Query("SELECT order_id, car_id, master_id, service_id, order_date, status FROM orders")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var o models.Order
		if err := rows.Scan(&o.OrderID, &o.CarID, &o.MasterID, &o.ServiceID, &o.OrderDate, &o.Status); err != nil {
			return err
		}
		data.Orders = append(data.Orders, o)
	}

	// Записываем в файл
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}
