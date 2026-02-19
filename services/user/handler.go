package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var req struct{ Username string }
	json.NewDecoder(r.Body).Decode(&req)
	id, err := RegisterUser(req.Username)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	//Get all users
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	//Get user by ID
}

//	func AddToCartHandler(w http.ResponseWriter, r *http.Request) {
//		var req struct{ ProductID, Quantity int }
//		json.NewDecoder(r.Body).Decode(&req)
//		userID, _ := strconv.Atoi(mux.Vars(r)["user_id"])
//		err := AddToCart(userID, req.ProductID, req.Quantity)
//		if err != nil {
//			http.Error(w, err.Error(), 500)
//			return
//		}
//		w.WriteHeader(200)
//	}

func GetCartsHandler(w http.ResponseWriter, r *http.Request) {
	//Get all carts
}

func GetCartHandler(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(mux.Vars(r)["user_id"])
	cart, err := GetCart(userID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(cart)
}
