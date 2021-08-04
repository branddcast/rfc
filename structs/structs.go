package structs

// estructura de entrada , POST
type Personas struct {
	Personas []Persona `json:"Personas"`
  }
  
  type Persona struct {
	Nombres          string `json:"Nombres"`
	ApellidoPaterno  string `json:"ApellidoPaterno"`
	ApellidoMaterno  string `json:"ApellidoMaterno"`
	FechaNacimiento  string `json:"FechaNacimiento"`
  }

// estructura de salida del servicio
type ResPersonas struct {
	NombrePersona    string `json:"NombrePersona"`
	RFC              string `json:"RFC"`
}