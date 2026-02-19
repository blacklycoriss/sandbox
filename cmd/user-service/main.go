package userservice

import (
	"log"
	"net/http"

	"github.com/blacklycoriss/sandbox/internal/config"
	"github.com/blacklycoriss/sandbox/internal/db"
	"github.com/blacklycoriss/sandbox/internal/messaging"
	"github.com/blacklycoriss/sandbox/services/user"

	"github.com/gorilla/mux"
)

func main() {
	//Init DBs
	errdb := db.InitPostgres(config.GetDBURL())
	if errdb != nil {
		log.Fatalf("Failed to initialize Postgres: %v", errdb)
	}
	db.InitRedis(config.GetRedisAddr())

	//Init RabbitMQ
	rmq_url := config.GetRabbitMQURL()

	rmq, err := messaging.New(rmq_url)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}

	err = rmq.Connect(rmq_url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	//Init Handlers
	r := mux.NewRouter()
	r.HandleFunc("/register", user.RegisterUserHandler).Methods("POST")
	r.HandleFunc("/users", user.GetUsersHandler).Methods("GET")
	r.HandleFunc("/user/{user_id}", user.GetUserHandler).Methods("GET")
	r.HandleFunc("/carts", user.GetCartsHandler).Methods("GET")
	r.HandleFunc("/carts/{cart_id}", user.GetCartHandler).Methods("GET")
	//r.HandleFunc("/carts/add/{user_id}", user.AddToCartHandler).Methods("POST")

	http.ListenAndServe(":8001", r)
}
