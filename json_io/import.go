package json_io

import (
	"database/sql"
	"encoding/json"
	"os"

	"autoservice/models"
)

func ImportFromJSON(db *sql.DB, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var data models.ExportData
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return err
	}

	// Импортируем клиентов
	for _, c := range data.Clients {
		_, err := db.Exec(`
			INSERT OR IGNORE INTO clients (client_id, name, phone, email)
			VALUES (?, ?, ?, ?)`,
			c.ClientID, c.Name, c.Phone, c.Email)
		if err != nil {
			return err
		}
	}

	// Импортируем автомобили
	for _, c := range data.Cars {
		_, err := db.Exec(`
			INSERT OR IGNORE INTO cars (car_id, brand, model, year, plate, client_id)
			VALUES (?, ?, ?, ?, ?, ?)`,
			c.CarID, c.Brand, c.Model, c.Year, c.Plate, c.ClientID)
		if err != nil {
			return err
		}
	}

	// Импортируем мастеров
	for _, m := range data.Masters {
		_, err := db.Exec(`
			INSERT OR IGNORE INTO masters (master_id, name, specialization, phone)
			VALUES (?, ?, ?, ?)`,
			m.MasterID, m.Name, m.Specialization, m.Phone)
		if err != nil {
			return err
		}
	}

	// Импортируем услуги
	for _, s := range data.Services {
		_, err := db.Exec(`
			INSERT OR IGNORE INTO services (service_id, name, price, duration)
			VALUES (?, ?, ?, ?)`,
			s.ServiceID, s.Name, s.Price, s.Duration)
		if err != nil {
			return err
		}
	}

	// Импортируем заказы
	for _, o := range data.Orders {
		_, err := db.Exec(`
			INSERT OR IGNORE INTO orders (order_id, car_id, master_id, service_id, order_date, status)
			VALUES (?, ?, ?, ?, ?, ?)`,
			o.OrderID, o.CarID, o.MasterID, o.ServiceID, o.OrderDate, o.Status)
		if err != nil {
			return err
		}
	}

	return nil
}
