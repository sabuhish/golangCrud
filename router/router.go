package router


import (
	"github.com/gorilla/mux"
	"cars/middleware"

)

func Router() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/api/cars",middleware.ReadAllcars).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/cars/{id}",middleware.ReadCar).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/cars/",middleware.CreateCar).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/cars/{id}",middleware.UpdateCar).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/cars/{id}",middleware.DeleteCar).Methods("DELETE", "OPTIONS")



	return router
}
