package aeropuerto

type Conexion struct {
	origen  string
	destino string
}

type FechaYCodigo struct {
	fecha  string
	codigo string
}

// Aeropuerto modela a un aeropuerto, que contiene la información de todos los vuelos que pasan por él.
type Aeropuerto interface {

	//Dado un vuelo lo agrega al sistema del aeropuerto.
	Agregar_vuelo(Vuelo)

	//Dado el código de un vuelo devuelve toda la información relacionada con él o un error si no se encuentra en el sistema.
	Info_vuelo(string) (string, error)

	//Dado un valor k entero devuelve los k vuelos más prioritarios que haya en el sistema.
	Mostrar_prioridad(int)

	//Dado un origen, un destino y una fecha devuelve el vuelo, si es que hay uno, cuya fecha de salida sea la más cercana a ella y que tenga esa conexión.
	Siguiente_vuelo(string, string, string) Vuelo

	//Dado un valor k entero y un rango de fechas (desde, hasta), devuelve los primeros k vuelos, si es que los hay, que estén en ese rango. El orden es ascendente en fechas.
	Ver_tablero_ASC(int, string, string)

	//Dado un valor k entero y un rango de fechas (desde, hasta), devuelve los primeros k vuelos, si es que los hay, que estén en ese rango. El orden es descendente en fechas.
	Ver_tablero_DESC(int, string, string)

	//Dado un rango de fechas (desde, hasta) borra del sistema todos los vuelos que se encuentren en ese rango y devuelve la información asociada a cada uno de ellos.
	Borrar(string, string)
}
