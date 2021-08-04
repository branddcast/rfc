package main

import (
	"log"
	"net/http"

	servicio "validadorRFC-go/servicios"
)

// funcion principal del programa
// funcion handle que se encarga de desplegar el API se indica puerto y direccion que se expone
func main() {

	log.Println("<<<<<< - Server listening in localhost:8080 - >>>>>>")
	
    http.HandleFunc("/serviciorfc", servicio.Serviciorfc)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
