package middleware


import (
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
	"cars/models"
	"github.com/gorilla/mux"
)



var carsDB []models.Car
var counter = 0

func init(){
	fmt.Println("Init fake DB")

	carsDB =append(carsDB, models.Car{ID: strconv.Itoa(counter),Model: "BMW X5", Price: "100000"})
	counter++
	
	carsDB =append(carsDB, models.Car{ID: strconv.Itoa(counter),Model: "Mercedes", Price: "100000"})
	counter++
	
	carsDB =append(carsDB, models.Car{ID: strconv.Itoa(counter),Model: "Porshe", Price: "1800000"})
	counter++
	}		

func ReadAllcars(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control_allow-Origin", "*")
	payload :=readAllcars()
	json.NewEncoder(w).Encode(payload)
}

func readAllcars() []models.Car{
	var cars []models.Car
	for _, element:= range carsDB {
		cars =append(cars, element)
	}
	return cars
}


func ReadCar(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	payload := readcar(params["id"])
	json.NewEncoder(w).Encode(payload)
}

func readcar(id string) models.Car {
	var car models.Car
	for _, element:= range carsDB {
		if element.ID ==id{
			car=element
			return car
		}
	}
	return car
}


func CreateCar(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Origin", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "Content-Type")

	var car models.Car 
	_ =json.NewDecoder(r.Body).Decode(&car)

	payload := createcar(car)
	json.NewEncoder(w).Encode(payload)
}

func createcar(car models.Car) models.Car{
	car.ID =strconv.Itoa(counter)
	counter++
	carsDB =append(carsDB, car)
	return car
}





func UpdateCar(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Origin", "PUT")
	w.Header().Set("Access-Control-Allow-Origin", "Content-Type")

	var car models.Car 
	_ =json.NewDecoder(r.Body).Decode(&car)

	params := mux.Vars(r)
	payload := updateCar(params["id"],car)
	json.NewEncoder(w).Encode(payload)
}

func updateCar(id string, car models.Car) models.Car{
	for index, element := range carsDB{
		if element.ID ==id{
			element.Model =car.Model
			element.Price =car.Price
			carsDB[index] =element
			return element
		
		}
	}
	return car

}




func DeleteCar(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Origin", "PUt")
	w.Header().Set("Access-Control-Allow-Origin", "Content-Type")

	
	params := mux.Vars(r)
	payload := deleteCar(params["id"])
	json.NewEncoder(w).Encode(payload)
}

func deleteCar(id string) models.Car {
	var car models.Car
	for index, element := range carsDB {
		if element.ID == id {
			carsDB = append(carsDB[:index], carsDB[index+1:]...)
			car = element
			break
		}
	}
	return car
}