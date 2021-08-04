package servicios

import (
	JsonForm "encoding/json"
	"log"
	"net/http"
	"strings"

	entrada "validadorRFC-go/structs"
	valida "validadorRFC-go/verificadores"
)

// servicio que se expone en la handleFunc
func Serviciorfc(w http.ResponseWriter, r *http.Request){

	// se inicializa la estructura y se comienza con la lectura de la informacion de entrada
	var contadorOk int
	var contadorNok int
	var inputInfo entrada.Personas
	decoder := JsonForm.NewDecoder(r.Body)
	decoder.Decode(&inputInfo)

	// se inicializa estrcutura de informacion de salida 
	salida := make([]entrada.ResPersonas, 0)

	// lectura de los elementos de entrada
	for i := 0; i < len(inputInfo.Personas); i++ {

		var rfcCorto string
		personaProcesar := inputInfo.Personas

		estructuraOk := valida.ValidaIntegridadInfo(personaProcesar[i])

		if estructuraOk {
			nombreCompleto := strings.ToUpper(personaProcesar[i].ApellidoPaterno + " " + personaProcesar[i].ApellidoMaterno + " " + personaProcesar[i].Nombres)
		    fechaNacimiento := personaProcesar[i].FechaNacimiento
		    log.Println("Nombre                  << ", nombreCompleto)
		    log.Println("Fecha Nacimiento        << ", fechaNacimiento)

			// primera parte del RFC
			parteLetrasRFC := valida.CalcularRFC(personaProcesar[i])

			// fecha de nacimiento RFC
			fechaNacRfc := valida.FechaNAcimiento(personaProcesar[i].FechaNacimiento)
			
			rfcCorto += parteLetrasRFC + fechaNacRfc

			// obtener homonimia
		    var rfc12posc string
		    homonimiaRfc := valida.HomonimiaRfc(nombreCompleto, rfcCorto)
	    	rfc12posc += rfcCorto + homonimiaRfc

		    // obtener digito verificador
		    digitoVerfRfc := valida.DigitoVerificador(rfc12posc)

            // armado del RFC para salida
		    var rfcHomoclave string
		    rfcHomoclave += rfc12posc + digitoVerfRfc

	    	log.Println("Homoclave               << ", rfcHomoclave)

	    	// mueve y guardando datos en la salida
	    	rfcCreado := entrada.ResPersonas {
	    		NombrePersona: nombreCompleto,
	    		RFC: rfcHomoclave,
	    	}
	    	salida = append(salida, rfcCreado)
	    	contadorOk ++

		} else {
			log.Println("Persona con error  << ", strings.ToUpper(personaProcesar[i].ApellidoPaterno + " " + personaProcesar[i].ApellidoMaterno + " " + personaProcesar[i].Nombres),)
			// mueve y guardando datos en la salida
	    	rfcCreado := entrada.ResPersonas {
	    		NombrePersona: strings.ToUpper(personaProcesar[i].ApellidoPaterno + " " + personaProcesar[i].ApellidoMaterno + " " + personaProcesar[i].Nombres),
	    		RFC: "error datos entrada",
	    	}
	    	salida = append(salida, rfcCreado)
			contadorNok ++
		}
	}
	
	// se da formato a salida como Json
	w.Header().Set("Content-Type", "application/json")
	// envia mensaje 200 en la salida
	w.WriteHeader(http.StatusOK)

	// este parrafo convierte la estructura en JSON y la envia a la salida
	JsonForm.NewEncoder(w).Encode(salida)
}
