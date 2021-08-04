package verificadores

import (
	"log"
	"regexp"
	"strings"
	entrada "validadorRFC-go/structs"
)

var AñoValido = regexp.MustCompile(`^\d{4}$`)
var MesValidoLetra = regexp.MustCompile(`([A-Za-z ])+$`)
var MesValidoNum = regexp.MustCompile(`^(0[1-9])|([1-9])|(1[02])$`)
var DiaValido = regexp.MustCompile(`^(0[1-9])|([1-9])|(1[0-9])|(2[0-9])|(3[01])$`)

var LetraValida = regexp.MustCompile("^[aA-zZ& `’àèìòùáéíóúÁÉÍÓÚÀÈÌÒÙ]+$")
var ReplaceAcentos = strings.NewReplacer("Á", "A", "É", "E", "Í", "I", "Ó", "O", "Ú", "U", "À", "A", "È", "E", "Ì", "I", "Ò", "O", "Ü", "U", "Ù", "U", "’", "", ".", "", ";", "", ":", "")

func EliminaEspacioDmas(s string) string {	
    return strings.Join(strings.Fields(s), " ")
}

func ValidaIntegridadInfo (p entrada.Persona) bool {

	var noApellido = false

	if fechaOK := validFechaNacimiento(p.FechaNacimiento); !fechaOK{
		log.Println("Fecha de nacimiento no valida -> ", p.FechaNacimiento)
		return false
	}

	if palabraOk := LetraValida.MatchString(p.Nombres); !palabraOk{
		log.Println("Nombre con caracteres no permitidos ", p.Nombres)
		return false
	} else {
		EliminaEspacioDmas(p.Nombres)
		// reemplaza acentos del nombre de entrada
	    ReplaceAcentos.Replace(p.Nombres)
	}

	if palabraOk := LetraValida.MatchString(p.ApellidoPaterno); !palabraOk{
		if len(p.ApellidoPaterno) == 0 {
			noApellido = true
		} else {
			log.Println("Apellido con caracteres no permitidos ", p.ApellidoPaterno)
			return false
		}
	} else {
		EliminaEspacioDmas(p.ApellidoPaterno)
		// reemplaza acentos del nombre de entrada
	    ReplaceAcentos.Replace(p.ApellidoPaterno)
	}

	if palabraOk := LetraValida.MatchString(p.ApellidoMaterno); !palabraOk{
		if len(p.ApellidoMaterno) == 0 && noApellido {
			log.Println("No hay apellidos de la persona, minimo informar uno ")
			return false
		} else if noApellido{
				log.Println("Apellido con caracteres no permitidos ", p.ApellidoMaterno)
				return false
			}
		} else {
			EliminaEspacioDmas(p.ApellidoMaterno)
			// reemplaza acentos del nombre de entrada
			ReplaceAcentos.Replace(p.ApellidoMaterno)
	}

	return true

}

func validFechaNacimiento (f string) bool {
	var varAux = ""
	fechaAux := make([]string, 0)

	for key, letter := range f {
		if string(letter) == "-" || string(letter) == "/" || string(letter) == "." {
			fechaAux = append(fechaAux, string(varAux))
			varAux = " "
		} else {
			varAux += string(letter)
		}

		if (len(f)-1) == key {
			fechaAux = append(fechaAux, string(varAux))
		}
	}

	if fechaOK := AñoValido.MatchString(fechaAux[0]); !fechaOK{
		log.Println("Año de nacimiento no valido-> ", fechaAux[0])
		return false
	}
	if fechaOK := MesValidoLetra.MatchString(fechaAux[1]); !fechaOK{
		if fechaOK := MesValidoNum.MatchString(fechaAux[1]); !fechaOK{
			log.Println("Mes de nacimiento no valido-> ", fechaAux[1])
			return false
		}
	}
	if fechaOK := DiaValido.MatchString(fechaAux[2]); !fechaOK{
		log.Println("Dia de nacimiento no valido-> ", fechaAux[2])
		return false
	}
	return true
}