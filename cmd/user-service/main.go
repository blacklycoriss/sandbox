package userservice

import (
	"net/http"

	"github.com/blacklycoriss/sandbox/internal/config"
	"github.com/blacklycoriss/sandbox/internal/db"
	"github.com/blacklycoriss/sandbox/internal/messaging"
	"github.com/blacklycoriss/sandbox/services/user"

	"github.com/gorilla/mux"
)

func main() {
	db.InitPostgres(config.GetDBURL())
	db.InitRedis(config.GetRedisAddr())
	messaging.InitRabbitMQ(config.GetRabbitMQURL()) // If publishing
	r := mux.NewRouter()
	r.HandleFunc("/register", user.RegisterHandler).Methods("POST")
	r.HandleFunc("/cart/{user_id}", user.GetCartHandler).Methods("GET")
	r.HandleFunc("/cart/add/{user_id}", user.AddToCartHandler).Methods("POST") // Note: Adjusted endpoint for param
	http.ListenAndServe(":8001", r)
}
