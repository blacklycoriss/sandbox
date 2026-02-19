package user

import (
	"context"
	"encoding/json"

	"github.com/blacklycoriss/sandbox/internal/db"
	"github.com/blacklycoriss/sandbox/internal/models"
)

func RegisterUser(username string) (int, error) {
	var id int
	err := db.Pool.QueryRow(context.Background(), "INSERT INTO users (username) VALUES ($1) RETURNING id", username).Scan(&id)
	return id, err
}
func AddToCart(userID, productID, quantity int) error {
	_, err := db.Pool.Exec(context.Background(), "INSERT INTO carts (user_id, product_id, quantity) VALUES ($1, $2, $3) ON CONFLICT (user_id, product_id) DO UPDATE SET quantity = carts.quantity + $3", userID, productID, quantity)
	if err == nil {
		// Invalidate cache
		db.Redis.Del(context.Background(), "cart:"+string(userID))
	}
	return err
}
func GetCart(userID int) ([]models.CartItem, error) {
	// Check cache
	val, err := db.Redis.Get(context.Background(), "cart:"+string(userID)).Result()
	if err == nil {
		// Cache hit, parse val to []CartItem
		var items []models.CartItem
		// Assuming val is a JSON string representing []CartItem
		if err := json.Unmarshal([]byte(val), &items); err == nil {
			return items, nil
		}
	}
	rows, err := db.Pool.Query(context.Background(), "SELECT product_id, quantity FROM carts WHERE user_id=$1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.CartItem
	for rows.Next() {
		var item models.CartItem
		if err := rows.Scan(&item.ProductID, &item.Quantity); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	jsonBytes, _ := json.Marshal(items)
	db.Redis.Set(context.Background(), "cart:"+string(userID), string(jsonBytes), 0)
	return items, nil
}
