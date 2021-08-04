package verificadores

import (
	"strconv"
	"strings"
)

var Anexo1 = map[string]string{" ":"00","A":"11","B":"12","O":"26","0":"00","C":"13","P":"27","1":"01","D":"14","Q":"28","2":"02","E":"15","R":"29","3":"03","F":"16","S":"32","4":"04","G":"17","T":"33","5":"05","H":"18","U":"34","6":"06","I":"19","V":"35","7":"07","J":"21","W":"36","8":"08","K":"22","X":"37","9":"09","L":"23","Y":"38","Ñ":"10","M":"24","Z":"39","N":"25","&":"40"}
var Anexo2 = map[int]string{0 :"1",17:"I",1 :"2",18:"J",2 :"3",19:"K",3 :"4",20:"L",4 :"5",21:"M",5 :"6",22:"N",6 :"7",23:"P",7 :"8",24:"Q",8 :"9",25:"R",9 :"A",26:"S",10:"B",27:"T",11:"C",28:"U",12:"D",29:"V",13:"E",30:"W",14:"F",31:"X",15:"G",32:"Y",16:"H",33:"Z"}
var Anexo3 = map[string]string{"0":"00","1":"01","2":"02","3":"03","4":"04","5":"05","6":"06","7":"07","8":"08","9":"09","A":"10","B":"11","C":"12","D":"13","E":"14","F":"15","G":"16","H":"17","I":"18","J":"19","K":"20","L":"21","M":"22","N":"23","Ñ":"38","&":"24","O":"25","P":"26","Q":"27","R":"28","S":"29","T":"30","U":"31","V":"32","W":"33","X":"34","Y":"35","Z":"36"," ":"37"}

func HomonimiaRfc (in string, rf string) string {

	// Se agrega un cero al valor de la primera letra para uniformar el criterio de los números a tomar de dos en dos
	var cadena = "0"
	var keyCount int
	var homonimiaVar string
	var sumatoria int
	parNumeros := make([]string, 0)

	// separa la cadena de entrada por letras
	separaIn := strings.Split(in, "")

	// Se asignaran valores a las letras del nombre o razón social de acuerdo a la tabla del Anexo1
	for key, value := range separaIn {
		keyCount += key
		for keyMap, valueMap := range Anexo1{
			if keyMap == value {
				cadena += valueMap
			}
		}
	}

	for i, j := 2, 0; i <= len(cadena); i, j = i+1, j+1 {
		parNumeros = append(parNumeros, cadena[j:i])
    }

    // Se efectuaran las multiplicaciones de los números tomados de dos en dos para la posición de la pareja	
	for i, j := 1, 0; i < len(cadena); i, j = i+1, j+1 {
		num0, _ := strconv.Atoi(parNumeros[j])
		num1, _ := strconv.Atoi(string(cadena[i]))

		multiResul :=  (num0 * num1)
		sumatoria += multiResul
    }

	// Se suma el resultado de las multiplicaciones y del resultado obtenido, se tomaran las tres últimas cifras y estas se dividen entre el factor 34,
	aux := strconv.Itoa(sumatoria)
	ultimos3Digitos := aux[len(aux)-3:]

	resultadoSumatoria, _ := strconv.Atoi(ultimos3Digitos)

	cociente := resultadoSumatoria / 34
	residuo := resultadoSumatoria % 34

	// Con el cociente y el residuo se consulta la tabla del Anexo II y se asigna la homonimia.
	for key, value := range Anexo2 {
		if key == cociente {
			homonimiaVar += value
		}
	}
	for key, value := range Anexo2 {
		if key == residuo {
			homonimiaVar += value
		}
	}
	return homonimiaVar
}

func DigitoVerificador (in string) string {

	var cadena string
	var digitoVerificador string
	var keyCount int
	sumatoria := 0
	multiplica := 0

	separaIn := strings.Split(in, "")

	// Se asignaran los valores del Anexo III a las letras y números del registro federal de contribuyentes formado a 12 posiciones
	for key, value := range separaIn {
		keyCount += key
		for keyMap, valueMap := range Anexo3{
			if keyMap == value {
				cadena += valueMap
			}
		}
	}

	// la siguiente fórmula:
    // (Vi * (Pi + 1)) + (Vi * (Pi + 1)) + ..............+ (Vi * (Pi + 1)) MOD 11
    // Vi Valor asociado al carácter de acuerdo a la tabla del Anexo III.
    // Pi Posición que ocupa el i-esimo carácter tomando de derecha a izquierda es decir P toma los valores de 1 a 12.

	for i, j, pi := 0, 2, 12; i < len(cadena); i, j, pi = i+2, j+2, pi-1 {
		vi, _ := strconv.Atoi(string(cadena[i:j]))
		multiplica = (vi) * (pi + 1)
		sumatoria += multiplica
    }

	// El resultado de la suma se divide entre el factor 11.
	residuo := sumatoria % 11

	/*Si el residuo es igual a cero, este será el valor que se le asignara al dígito verificador.
      Si el residuo es mayor a cero se restara este al factor 11: 11-3 =8
      Si el residuo es igual a 10 el dígito verificador será “A”.
      Por lo tanto “8“ es el dígito verificador de este ejemplo: GODE561231GR8.*/

	if residuo == 0 {
		digitoVerificador = strconv.Itoa(residuo)
	} else {
		if residuo > 0 {
			if residuo == 10 {
				digitoVerificador = "A"
			} else {
				residuo -= 11
				if residuo < 0 {
					residuo = residuo * -1
				}
				digitoVerificador = strconv.Itoa(residuo)
			}
		}
	}
	return digitoVerificador
}
