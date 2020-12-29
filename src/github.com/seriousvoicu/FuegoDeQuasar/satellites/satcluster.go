package satellites

import (
	"math"
	"sort"

	"github.com/seriousvoicu/FuegoDeQuasar/arrays"
	"github.com/seriousvoicu/FuegoDeQuasar/exestate"
	"github.com/seriousvoicu/FuegoDeQuasar/geometry"
	"github.com/seriousvoicu/FuegoDeQuasar/vectors"
)

type satcluster struct {
	Satellites []satellite `json:"satellites"`

	currState *exestate.State
}

func (this *satcluster) /*StateHandler.*/ RegisterState(state *exestate.State) {
	this.currState = state
}

func (this *satcluster) /*StateHandler.*/ GetState(consume bool) *exestate.State {
	state := this.currState

	if consume {
		this.currState = nil
	}
	return state
}
func (this *satcluster) getAt(index int) *satellite {
	if exestate.OnError(this) {
		return nil
	}

	return &this.Satellites[index]
}

func (this *satcluster) count() int {
	if exestate.OnError(this) {
		return 0
	}

	return len(this.Satellites)
}

func (this *satcluster) isEmpty() bool {
	if exestate.OnError(this) {
		return true
	}

	return (len(this.Satellites) <= 0)
}

func (this *satcluster) setDistances(distances []float32) {
	if exestate.OnError(this) {
		return
	}

	for i := 0; i < this.count(); i++ {
		this.getAt(i).Distance = float64(distances[i])
	}
}

func (this *satcluster) setMessages(messages [][]string) {
	if exestate.OnError(this) {
		return
	}

	for i := 0; i < this.count(); i++ {
		this.getAt(i).Message = messages[i]
	}
}

//Devuelvo la ubicacion del emisor
func (this *satcluster) getLocation() vectors.Vector2 {
	if exestate.OnError(this) {
		return vectors.GetEmptyVector2()
	}

	//Interseccion entre dos satelites
	pointA, pointB, state := geometry.GetCirclesIntersections(this.getAt(0), this.getAt(1))

	if !state.IsOk() {
		this.RegisterState(state)
		return vectors.GetEmptyVector2()
	}

	if pointA.IsEmpty() && pointB.IsEmpty() {
		this.RegisterState(exestate.ControlledError("No se pudo triengular, no hay interseccion (1) (satellities.satcluster.getLocation)"))
		return vectors.GetEmptyVector2()
	}

	var intersectsA = (pointA.IsEmpty() == false)
	var intersectsB = (pointB.IsEmpty() == false)

	//	equals, _ := vectors.Equals(pointA, pointB)

	//Para los restantes satelites se verifican las distancias a los puntos de la interseccion
	if this.count() > 2 {
		for i := 2; i < this.count(); i++ {

			intersectsA = intersectsA && !(math.Abs(pointA.DistanceTo(this.getAt(i).Pos)-this.getAt(i).Distance) > 0.005)
			intersectsB = intersectsB && !(math.Abs(pointB.DistanceTo(this.getAt(i).Pos)-this.getAt(i).Distance) > 0.005)

			if intersectsA == false && intersectsB == false {
				this.RegisterState(exestate.ControlledError("No se pudo triangular, no hay interseccion (2) (satellities.satcluster.getLocation)"))
				return vectors.GetEmptyVector2()
			}
		}
	}

	if intersectsA && intersectsB {
		this.RegisterState(exestate.ControlledError("No se pudo triangular, mas de un punto (satellities.satcluster.getLocation)"))
		return vectors.GetEmptyVector2()
	}

	if intersectsA {
		pointA.Round()
		return pointA
	} else if intersectsB {
		pointB.Round()
		return pointB
	} else {
		this.RegisterState(exestate.ControlledError("No se pudo triangular"))
		return vectors.GetEmptyVector2()
	}
}

func (this *satcluster) getMessage() string {
	if exestate.OnError(this) {
		return ""
	}

	//Ordeno el arreglo de menor a mayor (el menor de los arreglos tiene que ser el del tamaño original del mensaje, caso contrario hay perdida de informacion)
	sort.Slice(this.Satellites, func(i, j int) bool {
		return len(this.getAt(i).Message) < len(this.getAt(j).Message)
	})

	//Mergeo el primer y segundo arreglo, siendo el primero el mas chico de todos
	message, errorType := arrays.MergeStringArrays(this.getAt(0).Message, this.getAt(1).Message)

	if !errorType.IsOk() {
		this.RegisterState(exestate.ControlledError("No se pudo determinar el mensaje (satellites.satcluster.getMessage)"))
		return ""
	}

	//Mergeo cada arreglo mergeado con el siguiente
	for i := 2; i < this.count(); i++ {
		message, errorType = arrays.MergeStringArrays(message, this.getAt(i).Message)
		if !errorType.IsOk() {
			this.RegisterState(exestate.ControlledError("No se pudo determinar el mensaje (satellites.satcluster.getMessage)"))
			return ""
		}
	}

	//Verifico si hay elementos vacios en el mensaje final
	for i := 0; i < len(message); i++ {
		if message[i] == "" {
			this.RegisterState(exestate.ControlledError("No se pudo determinar el mensaje (satellites.satcluster.getMessage)"))
			return ""
		}
	}

	//Paso el arreglo a un string
	finalMsg, _ := arrays.StringArrayToString(message)
	if finalMsg == "" {
		this.RegisterState(exestate.ControlledError("No se pudo determinar el mensaje (satellites.satcluster.getMessage)"))
		return ""
	}

	return finalMsg
}
