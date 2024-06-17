package database

import (
	"database/sql"

	"github.com/matheusmhmelo/FullCycle-clean-arch/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) List() ([]*entity.Order, error) {
	rows, err := r.Db.Query("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []*entity.Order{}
	for rows.Next() {
		var res entity.Order
		if err := rows.Scan(&res.ID, &res.Price, &res.Tax, &res.FinalPrice); err != nil {
			return results, err
		}
		results = append(results, &res)
	}

	return results, nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
