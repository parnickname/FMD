package db

import "database/sql"

func InitDatabase(db *sql.DB) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS clients (
			client_id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			phone TEXT,
			email TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS cars (
			car_id INTEGER PRIMARY KEY AUTOINCREMENT,
			brand TEXT NOT NULL,
			model TEXT NOT NULL,
			year INTEGER,
			plate TEXT UNIQUE,
			client_id INTEGER REFERENCES clients(client_id)
		)`,
		`CREATE TABLE IF NOT EXISTS masters (
			master_id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			specialization TEXT,
			phone TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS services (
			service_id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			price REAL,
			duration REAL
		)`,
		`CREATE TABLE IF NOT EXISTS orders (
			order_id INTEGER PRIMARY KEY AUTOINCREMENT,
			car_id INTEGER REFERENCES cars(car_id),
			master_id INTEGER REFERENCES masters(master_id),
			service_id INTEGER REFERENCES services(service_id),
			order_date DATE DEFAULT CURRENT_DATE,
			status TEXT DEFAULT 'new'
		)`,
	}

	for _, q := range queries {
		if _, err := db.Exec(q); err != nil {
			return err
		}
	}

	return nil
}

func SeedDatabase(db *sql.DB) error {
	// Проверяем, есть ли уже данные
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM clients").Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return nil // Данные уже есть
	}

	// Добавляем тестовых клиентов
	clients := []string{
		`INSERT INTO clients (name, phone, email) VALUES ('Иванов Иван Иванович', '+7-999-123-45-67', 'ivanov@mail.ru')`,
		`INSERT INTO clients (name, phone, email) VALUES ('Петров Пётр Петрович', '+7-999-234-56-78', 'petrov@mail.ru')`,
		`INSERT INTO clients (name, phone, email) VALUES ('Сидоров Сидор Сидорович', '+7-999-345-67-89', 'sidorov@mail.ru')`,
		`INSERT INTO clients (name, phone, email) VALUES ('Козлов Андрей Викторович', '+7-999-456-78-90', 'kozlov@mail.ru')`,
		`INSERT INTO clients (name, phone, email) VALUES ('Новиков Дмитрий Сергеевич', '+7-999-567-89-01', 'novikov@mail.ru')`,
	}

	// Добавляем автомобили
	cars := []string{
		`INSERT INTO cars (brand, model, year, plate, client_id) VALUES ('Toyota', 'Camry', 2020, 'А123БВ77', 1)`,
		`INSERT INTO cars (brand, model, year, plate, client_id) VALUES ('BMW', 'X5', 2019, 'В456ГД99', 2)`,
		`INSERT INTO cars (brand, model, year, plate, client_id) VALUES ('Mercedes', 'E200', 2021, 'Е789ЖЗ77', 3)`,
		`INSERT INTO cars (brand, model, year, plate, client_id) VALUES ('Audi', 'A4', 2018, 'К012МН50', 4)`,
		`INSERT INTO cars (brand, model, year, plate, client_id) VALUES ('Volkswagen', 'Polo', 2022, 'О345ПР77', 5)`,
		`INSERT INTO cars (brand, model, year, plate, client_id) VALUES ('Toyota', 'RAV4', 2021, 'С678ТУ99', 1)`,
	}

	// Добавляем мастеров
	masters := []string{
		`INSERT INTO masters (name, specialization, phone) VALUES ('Кузнецов Алексей', 'Двигатели', '+7-999-111-11-11')`,
		`INSERT INTO masters (name, specialization, phone) VALUES ('Смирнов Виктор', 'Ходовая часть', '+7-999-222-22-22')`,
		`INSERT INTO masters (name, specialization, phone) VALUES ('Попов Михаил', 'Электрика', '+7-999-333-33-33')`,
		`INSERT INTO masters (name, specialization, phone) VALUES ('Волков Сергей', 'Кузовной ремонт', '+7-999-444-44-44')`,
	}

	// Добавляем услуги
	services := []string{
		`INSERT INTO services (name, price, duration) VALUES ('Замена масла', 2500.0, 1.0)`,
		`INSERT INTO services (name, price, duration) VALUES ('Диагностика двигателя', 3000.0, 2.0)`,
		`INSERT INTO services (name, price, duration) VALUES ('Замена тормозных колодок', 4500.0, 1.5)`,
		`INSERT INTO services (name, price, duration) VALUES ('Развал-схождение', 3500.0, 1.0)`,
		`INSERT INTO services (name, price, duration) VALUES ('Замена свечей зажигания', 2000.0, 0.5)`,
		`INSERT INTO services (name, price, duration) VALUES ('Компьютерная диагностика', 1500.0, 0.5)`,
		`INSERT INTO services (name, price, duration) VALUES ('Замена ремня ГРМ', 8000.0, 4.0)`,
		`INSERT INTO services (name, price, duration) VALUES ('Покраска элемента', 15000.0, 8.0)`,
	}

	// Добавляем заказы
	orders := []string{
		`INSERT INTO orders (car_id, master_id, service_id, order_date, status) VALUES (1, 1, 1, '2025-12-01', 'completed')`,
		`INSERT INTO orders (car_id, master_id, service_id, order_date, status) VALUES (2, 2, 3, '2025-12-05', 'completed')`,
		`INSERT INTO orders (car_id, master_id, service_id, order_date, status) VALUES (3, 1, 2, '2025-12-10', 'completed')`,
		`INSERT INTO orders (car_id, master_id, service_id, order_date, status) VALUES (4, 3, 6, '2025-12-15', 'completed')`,
		`INSERT INTO orders (car_id, master_id, service_id, order_date, status) VALUES (5, 2, 4, '2025-12-18', 'in_progress')`,
		`INSERT INTO orders (car_id, master_id, service_id, order_date, status) VALUES (6, 1, 7, '2025-12-20', 'in_progress')`,
		`INSERT INTO orders (car_id, master_id, service_id, order_date, status) VALUES (1, 4, 8, '2025-12-22', 'new')`,
		`INSERT INTO orders (car_id, master_id, service_id, order_date, status) VALUES (2, 1, 5, '2025-12-25', 'new')`,
	}

	allQueries := append(clients, cars...)
	allQueries = append(allQueries, masters...)
	allQueries = append(allQueries, services...)
	allQueries = append(allQueries, orders...)

	for _, q := range allQueries {
		if _, err := db.Exec(q); err != nil {
			return err
		}
	}

	return nil
}
