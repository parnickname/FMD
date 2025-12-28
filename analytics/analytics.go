package analytics

import (
	"database/sql"
	"fmt"
)

// RevenueByService показывает общую выручку по услугам
func RevenueByService(db *sql.DB) error {
	query := `
		SELECT s.name, SUM(s.price) as total_revenue
		FROM orders o
		JOIN services s ON o.service_id = s.service_id
		WHERE o.status = 'completed'
		GROUP BY s.service_id
		ORDER BY total_revenue DESC
	`

	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("\n=== ВЫРУЧКА ПО УСЛУГАМ ===")
	fmt.Printf("%-30s | %s\n", "Услуга", "Выручка (руб.)")
	fmt.Println("----------------------------------------")

	for rows.Next() {
		var name string
		var revenue float64
		if err := rows.Scan(&name, &revenue); err != nil {
			return err
		}
		fmt.Printf("%-30s | %.2f\n", name, revenue)
	}

	return nil
}

// OrdersByMaster показывает количество заказов по мастерам
func OrdersByMaster(db *sql.DB) error {
	query := `
		SELECT m.name, COUNT(o.order_id) as order_count
		FROM masters m
		LEFT JOIN orders o ON m.master_id = o.master_id
		GROUP BY m.master_id
		ORDER BY order_count DESC
	`

	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("\n=== ЗАКАЗЫ ПО МАСТЕРАМ ===")
	fmt.Printf("%-25s | %s\n", "Мастер", "Кол-во заказов")
	fmt.Println("----------------------------------------")

	for rows.Next() {
		var name string
		var count int
		if err := rows.Scan(&name, &count); err != nil {
			return err
		}
		fmt.Printf("%-25s | %d\n", name, count)
	}

	return nil
}

// AverageCheck показывает средний чек заказа
func AverageCheck(db *sql.DB) error {
	query := `
		SELECT AVG(s.price) as average_check
		FROM orders o
		JOIN services s ON o.service_id = s.service_id
	`

	var avgCheck float64
	err := db.QueryRow(query).Scan(&avgCheck)
	if err != nil {
		return err
	}

	fmt.Println("\n=== СРЕДНИЙ ЧЕК ===")
	fmt.Printf("Средний чек заказа: %.2f руб.\n", avgCheck)

	return nil
}

// CarsByBrand показывает количество автомобилей по маркам
func CarsByBrand(db *sql.DB) error {
	query := `
		SELECT brand, COUNT(*) as car_count
		FROM cars
		GROUP BY brand
		ORDER BY car_count DESC
	`

	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("\n=== АВТОМОБИЛИ ПО МАРКАМ ===")
	fmt.Printf("%-20s | %s\n", "Марка", "Количество")
	fmt.Println("----------------------------------------")

	for rows.Next() {
		var brand string
		var count int
		if err := rows.Scan(&brand, &count); err != nil {
			return err
		}
		fmt.Printf("%-20s | %d\n", brand, count)
	}

	return nil
}

// OrdersByStatus показывает количество заказов по статусам
func OrdersByStatus(db *sql.DB) error {
	query := `
		SELECT status, COUNT(*) as count
		FROM orders
		GROUP BY status
		ORDER BY count DESC
	`

	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("\n=== ЗАКАЗЫ ПО СТАТУСАМ ===")
	fmt.Printf("%-15s | %s\n", "Статус", "Количество")
	fmt.Println("----------------------------------------")

	statusNames := map[string]string{
		"new":         "Новый",
		"in_progress": "В работе",
		"completed":   "Завершён",
	}

	for rows.Next() {
		var status string
		var count int
		if err := rows.Scan(&status, &count); err != nil {
			return err
		}
		displayStatus := statusNames[status]
		if displayStatus == "" {
			displayStatus = status
		}
		fmt.Printf("%-15s | %d\n", displayStatus, count)
	}

	return nil
}

// TotalRevenue показывает общую выручку
func TotalRevenue(db *sql.DB) error {
	query := `
		SELECT SUM(s.price) as total
		FROM orders o
		JOIN services s ON o.service_id = s.service_id
		WHERE o.status = 'completed'
	`

	var total sql.NullFloat64
	err := db.QueryRow(query).Scan(&total)
	if err != nil {
		return err
	}

	fmt.Println("\n=== ОБЩАЯ ВЫРУЧКА ===")
	if total.Valid {
		fmt.Printf("Общая выручка (завершённые заказы): %.2f руб.\n", total.Float64)
	} else {
		fmt.Println("Завершённых заказов нет")
	}

	return nil
}

// ShowAllAnalytics выводит всю аналитику
func ShowAllAnalytics(db *sql.DB) error {
	if err := TotalRevenue(db); err != nil {
		return err
	}
	if err := RevenueByService(db); err != nil {
		return err
	}
	if err := OrdersByMaster(db); err != nil {
		return err
	}
	if err := AverageCheck(db); err != nil {
		return err
	}
	if err := CarsByBrand(db); err != nil {
		return err
	}
	if err := OrdersByStatus(db); err != nil {
		return err
	}
	return nil
}
