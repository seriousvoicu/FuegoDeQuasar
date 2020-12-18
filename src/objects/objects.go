package objects

import (
	"core"
	"encoding/json"
	"io"
	"math"
	"mathext"
	"sort"
)

//-------------------

type Satellitie struct {
	Name     string   `json:"name"`
	Distance float64  `json:"distance"`
	Message  []string `json:"message"`
	Pos      *core.Vector2
}

type SatellitieCluster struct {
	Satellities []Satellitie `json:"satellities"`
}

//--------------------------

func (element SatellitieCluster) GetAt(index int) Satellitie {
	return element.Satellities[index]
}

func (element SatellitieCluster) Count() int {
	return len(element.Satellities)
}

func (element SatellitieCluster) IsEmpty() bool {
	return (len(element.Satellities) <= 0)
}

//Devuelvo la ubicacion del emisor
func (cluster SatellitieCluster) GetLocation() *core.Vector2 {

	//Agarro los dos primeros satelites y calculo la interseccion
	//La idea es que dos satelites pueden coincidir en cero, 1 o 2 puntos
	intersectionOne, intersectionTwo := mathext.GetIntersection(cluster.GetAt(0), cluster.GetAt(1))

	var intersectsA = true
	var intersectsB = true

	var distance float64

	//Para los siguientes satelites es una cuestion de calcular las distancias entre cada uno y estos puntos
	if cluster.Count() > 2 {
		for i := 2; i < cluster.Count(); i++ {
			if intersectsA {
				distance = mathext.GetDistance(intersectionOne, cluster.GetAt(i).Pos)

				if math.Abs(distance-cluster.GetAt(i).Distance) > 0.005 {
					intersectsA = false
				}
			}

			if intersectsB {
				distance = mathext.GetDistance(intersectionTwo, cluster.GetAt(i).Pos)
				if math.Abs(distance-cluster.GetAt(i).Distance) > 0.005 {
					intersectsB = false
				}
			}

			if intersectsA == false && intersectsB == false {
				return core.GetEmptyVector()
			}
		}
	}

	if intersectsA {
		return intersectionOne
	} else if intersectsB {
		return intersectionTwo
	} else {
		return core.GetEmptyVector()
	}
}

//implementaciones de mathext.Circle
func (element Satellitie) Radious() float64 {
	return float64(element.Distance)
}

func (element Satellitie) Position() *core.Vector2 {
	return element.Pos
}

//--------------------------
func (cluster SatellitieCluster) GetMessage() []string {

	//Para evitar perder tiempo, doy por hecho que el desfasaje es solo al principio del mensaje
	//Solucion: Recorro de atras para adelante xD

	//El desfasaje implica que hay elementos de mas, por lo tanto el mas chicos de los mensajes tiene que representar el tamaño correcto
	//Si todos tienen desfasaje, va a dar error, porque va haber una hilera de elementos vacios, lo que nos deja un elemento vacio sin poder determinar que es
	//A su vez, ninguno de los elementos mas largos deben tener un valor valido en el area de "desfasaje", porque esto implica que hay perdida de informacion

	sort.Slice(cluster.Satellities, func(i, j int) bool {
		return len(cluster.GetAt(i).Message) < len(cluster.GetAt(j).Message)
	})

	//Recorro elemento por elemento de cada uno de los arreglos, levantando el valor que le corresponde
	var message = MergeMessages(cluster.GetAt(0).Message, cluster.GetAt(1).Message)

	for i := 2; i < cluster.Count(); i++ {
		message = MergeMessages(message, cluster.GetAt(i).Message)
	}

	return message
}

//Mergea dos mensajes y devuelve el mensaje mergeado
func MergeMessages(messageOne []string, messageTwo []string) []string {
	var lastElement = 0
	var biggestSize = 0
	var oneBiggest = true

	//Agarro el elemento mas chico
	if len(messageOne) > len(messageTwo) {
		lastElement = len(messageTwo) - 1
		biggestSize = len(messageOne)
		oneBiggest = true
	} else {
		lastElement = len(messageOne) - 1
		biggestSize = len(messageTwo)
		oneBiggest = false
	}

	var newMessage []string = make([]string, lastElement+1)

	//Recorro de atras para delante todos los elementos
	for i := 1; i <= biggestSize; i++ {

		if i <= len(newMessage) {
			if (messageOne[len(messageOne)-i] != messageTwo[len(messageTwo)-i] && messageOne[len(messageOne)-i] != "" && messageTwo[len(messageTwo)-i] != "") || (messageOne[len(messageOne)-i] == "" && messageTwo[len(messageTwo)-i] == "") {
				return nil //No hay coincidencia
			}

			if messageOne[len(messageOne)-i] == "" {
				newMessage[len(newMessage)-i] = messageTwo[len(messageTwo)-i]
			} else {
				newMessage[len(newMessage)-i] = messageOne[len(messageOne)-i]
			}

			lastElement = i
		} else {
			//Valido que todos los elementos en la zona de "desfase" sean elementos vacíos
			if oneBiggest {
				if messageOne[len(messageOne)-i] != "" {
					return nil
				}
			} else {
				if messageTwo[len(messageTwo)-i] != "" {
					return nil
				}
			}
		}

	}

	return newMessage[:]
}

//--------------------------

//Esto lo pongo acá, porque no quiero exponer el objeto concreto
func InstantiateFromJson(jsonInput io.Reader) SatellitieCluster {

	var cluster SatellitieCluster

	err := json.NewDecoder(jsonInput).Decode(&cluster)
	if err != nil {
		return cluster
	}

	return cluster

}
