package verificadores

import (
	"strings"
)

func FechaNAcimiento(f string) string {

	var month string
	var day string
	var fechaRfc string

	fechaArray := strings.Split(f, "-")

	//ultimas dos cifras del año
	yearLast2 := fechaArray[0][len(fechaArray[0])-2:]

	// de mes palabra string a numero int
	switch strings.ToUpper(fechaArray[1]) {
	case "ENERO", "1", "01" :
		month = "01"
	case "FEBRERO", "2", "02" :
		month = "02"
	case "MARZO", "3", "03" :
		month = "03"
	case "ABRIL", "4", "04":
		month = "04"
	case "MAYO", "5", "05" :
		month = "05"
	case "JUNIO", "6", "06" :
		month = "06"
	case "JULIO",   "7", "07" :
		month = "07"
	case "AGOSTO", "8", "08" :
		month = "08"
	case "SEPTIEMBRE", "9", "09" :
		month = "09"
	case "OCTUBRE", "10" :
		month = "10"
	case "NOVIEMBRE", "11" :
		month = "11"
	case "DICIEMBRE", "12" :
		month = "12"
	}

	if len(string(fechaArray[2])) == 1 {
		day += "0" + string(fechaArray[2])
	} else {
		day = string(fechaArray[2])
	}

	fechaRfc = yearLast2 + month + day

	return fechaRfc
}