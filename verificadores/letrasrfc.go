package verificadores

import (
	"math"
	"regexp"
	"strings"
	entrada "validadorRFC-go/structs"
)

func filtrar(i int) float64 {
    if i != -1 {
        return float64(i)
    }
    return float64(45000)
}
func primerLetra(cadena string) string {
    //primera letra apellido paterno
    primerraLetraAP := strings.ToUpper(cadena)
    firstCharAP := primerraLetraAP[0]
    return string(firstCharAP)
}
func primerVocalAP(apellido string) string {
    //primera vocal apellido paterno
    primerraVocalAP := apellido[1:len(apellido)-1]
    a := filtrar(strings.Index(primerraVocalAP, "A"))
    e := filtrar(strings.Index(primerraVocalAP, "E"))
    i := filtrar(strings.Index(primerraVocalAP, "I"))
    o := filtrar(strings.Index(primerraVocalAP, "O"))
    u := filtrar(strings.Index(primerraVocalAP, "U"))
    min := math.Min(a, math.Min(e, math.Min(i, math.Min(o, u))))
    ind := int(min)
    vocal := string(primerraVocalAP[ind])
    return vocal
}
//Func. para validar que el nombre compuesto no contenga María o José como principal.
func validar(persona entrada.Persona) entrada.Persona {
    match, _ := regexp.Compile("^(MA|MA.|MARIA|JOSE)\\s+")
    persona = regla8y10(persona)
    persona.Nombres = match.ReplaceAllString(persona.Nombres, "")
    return persona
}
//Func. para susutituir caracteres especiales como (' ´ `) y eliminar articulos, preposiciones, etc.
func regla8y10(persona entrada.Persona) entrada.Persona {
    matchR8, _ := regexp.Compile("\\b(((L(?:AS?|OS)))|((?:DEL?))|(LOS)|(Y))[ ]?\\b")
    matchR10, _ := regexp.Compile("[^\\w\\s]")
    persona.Nombres = matchR10.ReplaceAllLiteralString(matchR8.ReplaceAllLiteralString(persona.Nombres, ""), "")
    persona.ApellidoPaterno = matchR10.ReplaceAllLiteralString(matchR8.ReplaceAllLiteralString(persona.ApellidoPaterno, ""), "")
    persona.ApellidoMaterno = matchR10.ReplaceAllLiteralString(matchR8.ReplaceAllLiteralString(persona.ApellidoMaterno, ""), "")
    return persona
}
//Func. para censurar palabras inconvenientes
func regla9(expresion string) string {
    matchR9, _ := regexp.Compile("BUE[IY]|C(?:A[CGK][AO]|O(?:GE|J[AEIO])|ULO)|FETO|GUEY|JOTO|K(?:A(?:[CG][AO]|KA)|O(?:GE|JO)|ULO)|M(?:AM[EO]|E(?:A[RS]|ON)|ION|OCO|ULA)|P(?:E(?:D[AO]|NE)|UT[AO])|QULO|R(?:ATA|UIN)")
    if matchR9.MatchString(expresion) {
        return expresion[0:len(expresion)-1] + "X"
    }
    return expresion
}
func formatearDatos(persona entrada.Persona) entrada.Persona {
    //Elimina acentos y convierte a mayúsculas
    var remplazarTildes = strings.NewReplacer("Á", "A", "É", "E", "Í", "I", "Ó", "O", "Ú", "U")
    persona.Nombres = remplazarTildes.Replace(strings.ToUpper(persona.Nombres))
    persona.ApellidoPaterno = remplazarTildes.Replace(strings.ToUpper(persona.ApellidoPaterno))
    persona.ApellidoMaterno = remplazarTildes.Replace(strings.ToUpper(persona.ApellidoMaterno))
    return persona
}
func CalcularRFC(persona entrada.Persona) string {
    persona = formatearDatos(persona)
    var RFC = ""
    //Regla 6,8,10
    persona = validar(persona)
    //Regla 7
    if persona.ApellidoMaterno != "" && persona.ApellidoPaterno == "" {
        RFC += primerLetra(persona.ApellidoMaterno) + primerVocalAP(persona.ApellidoMaterno)
        RFC += persona.Nombres[0:2]
    } else if persona.ApellidoMaterno == "" && persona.ApellidoPaterno != "" {
        RFC += primerLetra(persona.ApellidoPaterno) + primerVocalAP(persona.ApellidoPaterno)
        RFC += persona.Nombres[0:2]
    } else {
        RFC += primerLetra(persona.ApellidoPaterno) + primerVocalAP(persona.ApellidoPaterno) + primerLetra(persona.ApellidoMaterno)
        RFC += primerLetra(persona.Nombres)
    }
    //Regla 9
    return regla9(RFC)
}
