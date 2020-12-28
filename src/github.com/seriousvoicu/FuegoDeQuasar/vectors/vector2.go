package vectors

import (
	"math"
	"strconv"

	"github.com/seriousvoicu/FuegoDeQuasar/exestate"
)

type Vector2 struct {
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Empty bool

	currState *exestate.State
}

func (this *Vector2) /*StateHandler.*/ RegisterState(state *exestate.State) {
	this.currState = state
}

func (this *Vector2) /*StateHandler.*/ GetState(consume bool) *exestate.State {

	state := this.currState

	if consume {
		this.currState = nil
	}
	return state
}

func (this *Vector2) IsEmpty() bool {
	if exestate.OnError(this) {
		return true
	}

	return this.Empty
}

func (this *Vector2) Round() {
	if exestate.OnError(this) {
		return
	}

	this.X = math.Round(this.X)
	this.Y = math.Round(this.Y)
}

func (this *Vector2) DistanceTo(posB *Vector2) float64 {
	if exestate.OnError(this) {
		return 0
	}

	if this.IsEmpty() || posB.IsEmpty() {
		this.RegisterState(exestate.ControlledError("Es un vector vacio (vectors.Vector2.DistanceTo)"))
		return -1
	}

	return math.Hypot(posB.X-this.X, posB.Y-this.Y)
}

func (this *Vector2) ToString() string {
	if exestate.OnError(this) {
		return "X: - Y: -"
	}

	if this.IsEmpty() {
		this.RegisterState(exestate.ControlledError("Es un vector vacio (vectors.Vector2.ToString)"))
		return "X: - Y: -"
	}

	x := strconv.FormatFloat(this.X, 'f', -1, 32)

	y := strconv.FormatFloat(this.Y, 'f', -1, 32)

	result := "X: " + x + " Y: " + y

	return result
}

func Equals(vectorA Vector2, vectorB Vector2) (bool, *exestate.State) {
	if vectorA.IsEmpty() || vectorB.IsEmpty() {
		return false, exestate.ControlledError("Almenos uno de los vectores se encuentra vacio (vectors.Vector2.Equals)")
	}

	/*if vectorA.X == nil || vectorA.Y == nil || vectorB.X == nil || vectorB.Y == nil {
		return false, exestate.ControlledError("Almenos uno de los vectores se sin una de sus cordenadas definidas (vectors.Vector2.Equals)")
	}*/

	if math.Abs(vectorA.X-vectorB.X) < 0.001 && math.Abs(vectorA.Y-vectorB.Y) < 0.001 {
		return true, exestate.Ok()
	}

	return false, exestate.Ok()
}

func GetEmptyVector2() Vector2 {
	return Vector2{0, 0, true, exestate.Ok()}
}
