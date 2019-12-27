package main


import (
	"fmt"
	"log"
	"net/http"
	"cars/router"
)



func main(){
	fmt.Printf("hello world ")
	log.Fatal(http.ListenAndServe(":8000", router.Router()))

}
