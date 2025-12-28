package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"autoservice/analytics"
	"autoservice/db"
	"autoservice/json_io"
)

func main() {
	conn, err := db.Connect()
	if err != nil {
		fmt.Println("Ошибка подключения к базе данных:", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println()
		fmt.Println("==============================")
		fmt.Println(" СИСТЕМА УПРАВЛЕНИЯ АВТОСЕРВИСОМ")
		fmt.Println("==============================")
		fmt.Println("1 — Создать и заполнить БД")
		fmt.Println("2 — Показать всех клиентов")
		fmt.Println("3 — Показать все автомобили")
		fmt.Println("4 — Показать всех мастеров")
		fmt.Println("5 — Показать все услуги")
		fmt.Println("6 — Показать все заказы")
		fmt.Println("7 — Показать аналитику")
		fmt.Println("8 — Экспорт в JSON")
		fmt.Println("9 — Импорт из JSON")
		fmt.Println("0 — Выход")
		fmt.Println()
		fmt.Print("Выберите пункт меню: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			if err := db.InitDatabase(conn); err != nil {
				fmt.Println("Ошибка создания таблиц:", err)
				continue
			}
			if err := db.SeedDatabase(conn); err != nil {
				fmt.Println("Ошибка заполнения данными:", err)
				continue
			}
			fmt.Println("База данных успешно создана и заполнена тестовыми данными!")

		case "2":
			showClients(conn)

		case "3":
			showCars(conn)

		case "4":
			showMasters(conn)

		case "5":
			showServices(conn)

		case "6":
			showOrders(conn)

		case "7":
			if err := analytics.ShowAllAnalytics(conn); err != nil {
				fmt.Println("Ошибка получения аналитики:", err)
			}

		case "8":
			filename := "export_data.json"
			if err := json_io.ExportToJSON(conn, filename); err != nil {
				fmt.Println("Ошибка экспорта:", err)
				continue
			}
			fmt.Printf("Данные успешно экспортированы в файл %s\n", filename)

		case "9":
			filename := "export_data.json"
			if err := json_io.ImportFromJSON(conn, filename); err != nil {
				fmt.Println("Ошибка импорта:", err)
				continue
			}
			fmt.Printf("Данные успешно импортированы из файла %s\n", filename)

		case "0":
			fmt.Println("До свидания!")
			return

		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}

func showClients(conn *sql.DB) {
	rows, err := conn.Query("SELECT client_id, name, phone, email FROM clients")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\n=== КЛИЕНТЫ ===")
	fmt.Printf("%-4s | %-30s | %-18s | %s\n", "ID", "ФИО", "Телефон", "Email")
	fmt.Println("------------------------------------------------------------------------")

	for rows.Next() {
		var id int
		var name, phone, email string
		if err := rows.Scan(&id, &name, &phone, &email); err != nil {
			fmt.Println("Ошибка чтения:", err)
			continue
		}
		fmt.Printf("%-4d | %-30s | %-18s | %s\n", id, name, phone, email)
	}
}

func showCars(conn *sql.DB) {
	rows, err := conn.Query(`
		SELECT c.car_id, c.brand, c.model, c.year, c.plate, cl.name
		FROM cars c
		JOIN clients cl ON c.client_id = cl.client_id
	`)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\n=== АВТОМОБИЛИ ===")
	fmt.Printf("%-4s | %-12s | %-10s | %-6s | %-10s | %s\n", "ID", "Марка", "Модель", "Год", "Номер", "Владелец")
	fmt.Println("--------------------------------------------------------------------------------")

	for rows.Next() {
		var id, year int
		var brand, model, plate, owner string
		if err := rows.Scan(&id, &brand, &model, &year, &plate, &owner); err != nil {
			fmt.Println("Ошибка чтения:", err)
			continue
		}
		fmt.Printf("%-4d | %-12s | %-10s | %-6d | %-10s | %s\n", id, brand, model, year, plate, owner)
	}
}

func showMasters(conn *sql.DB) {
	rows, err := conn.Query("SELECT master_id, name, specialization, phone FROM masters")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\n=== МАСТЕРА ===")
	fmt.Printf("%-4s | %-20s | %-20s | %s\n", "ID", "Имя", "Специализация", "Телефон")
	fmt.Println("------------------------------------------------------------------------")

	for rows.Next() {
		var id int
		var name, spec, phone string
		if err := rows.Scan(&id, &name, &spec, &phone); err != nil {
			fmt.Println("Ошибка чтения:", err)
			continue
		}
		fmt.Printf("%-4d | %-20s | %-20s | %s\n", id, name, spec, phone)
	}
}

func showServices(conn *sql.DB) {
	rows, err := conn.Query("SELECT service_id, name, price, duration FROM services")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\n=== УСЛУГИ ===")
	fmt.Printf("%-4s | %-30s | %-12s | %s\n", "ID", "Наименование", "Цена (руб.)", "Время (ч.)")
	fmt.Println("------------------------------------------------------------------------")

	for rows.Next() {
		var id int
		var name string
		var price, duration float64
		if err := rows.Scan(&id, &name, &price, &duration); err != nil {
			fmt.Println("Ошибка чтения:", err)
			continue
		}
		fmt.Printf("%-4d | %-30s | %-12.2f | %.1f\n", id, name, price, duration)
	}
}

func showOrders(conn *sql.DB) {
	rows, err := conn.Query(`
		SELECT o.order_id, c.brand || ' ' || c.model AS car,
		       m.name AS master, s.name AS service,
		       o.order_date, o.status
		FROM orders o
		JOIN cars c ON o.car_id = c.car_id
		JOIN masters m ON o.master_id = m.master_id
		JOIN services s ON o.service_id = s.service_id
		ORDER BY o.order_date DESC
	`)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	defer rows.Close()

	statusNames := map[string]string{
		"new":         "Новый",
		"in_progress": "В работе",
		"completed":   "Завершён",
	}

	fmt.Println("\n=== ЗАКАЗЫ ===")
	fmt.Printf("%-4s | %-18s | %-18s | %-25s | %-12s | %s\n", "ID", "Автомобиль", "Мастер", "Услуга", "Дата", "Статус")
	fmt.Println("------------------------------------------------------------------------------------------------------")

	for rows.Next() {
		var id int
		var car, master, service, date, status string
		if err := rows.Scan(&id, &car, &master, &service, &date, &status); err != nil {
			fmt.Println("Ошибка чтения:", err)
			continue
		}
		displayStatus := statusNames[status]
		if displayStatus == "" {
			displayStatus = status
		}
		fmt.Printf("%-4d | %-18s | %-18s | %-25s | %-12s | %s\n", id, car, master, service, date, displayStatus)
	}
}
