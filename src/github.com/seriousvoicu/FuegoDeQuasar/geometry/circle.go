package geometry

import (
	"math"

	"github.com/seriousvoicu/FuegoDeQuasar/exestate"
	"github.com/seriousvoicu/FuegoDeQuasar/vectors"
)

type Circle interface {
	Radious() float64
	Position() *vectors.Vector2
}

//Retorno los puntos de interseccion. Puede haber 2, 1 o ninguno.
func GetCirclesIntersections(circleA Circle, circleB Circle) (vectors.Vector2, vectors.Vector2, *exestate.State) {

	//Antes que nada evaluo si puede haber interseccion
	possible, state := IsCirclesIntersectionPossible(circleA, circleB)

	if !state.IsOk() {
		return vectors.GetEmptyVector2(), vectors.GetEmptyVector2(), state
	} else if !possible {
		return vectors.GetEmptyVector2(), vectors.GetEmptyVector2(), exestate.ControlledError("No existe interseccion (geometry.IsCirclesIntersectionPossible)")
	}

	//Calculo la interseccion de los circulos

	var d = circleA.Position().DistanceTo(circleB.Position()) // math.Hypot(posTwo.x-posOne.x, posTwo.y-posOne.y) //Distancia entre ambos puntos

	if d <= circleA.Radious()+circleB.Radious() && d >= math.Abs(circleB.Radious()-circleA.Radious()) {

		var ex = (circleB.Position().X - circleA.Position().X) / d
		var ey = (circleB.Position().Y - circleA.Position().Y) / d

		var x = (circleA.Radious()*circleA.Radious() - circleB.Radious()*circleB.Radious() + d*d) / (2 * d)
		var y = math.Sqrt(circleA.Radious()*circleA.Radious() - x*x)

		var P1, P2 vectors.Vector2
		P1 = vectors.Vector2{X: RoundZero(circleA.Position().X + x*ex - y*ey), Y: RoundZero(circleA.Position().Y + x*ey + y*ex)}
		P2 = vectors.Vector2{X: RoundZero(circleA.Position().X + x*ex + y*ey), Y: RoundZero(circleA.Position().Y + x*ey - y*ex)}

		if P1.X == P2.X && P1.Y == P2.Y {
			P2 = vectors.GetEmptyVector2()
		}

		return P1, P2, exestate.Ok()
	}

	return vectors.GetEmptyVector2(), vectors.GetEmptyVector2(), exestate.Ok()
}

//El valor menor que 0.0001 lo redondeo a cero
func RoundZero(v float64) float64 {
	if math.Abs(v) <= 0.0001 {
		v = 0
	}

	return v
}

//Validaciones iniciales para determinar si dos ciculos efectivamente pueden coindidir
func IsCirclesIntersectionPossible(circleA Circle, circleB Circle) (bool, *exestate.State) {
	if circleA.Radious() == 0 || circleB.Radious() == 0 {
		return false, exestate.ControlledError("Los circulos no poseen radio (geometry.IsCirclesIntersectionPossible)")
	} else if circleA.Radious() < 0 || circleB.Radious() < 0 {
		return false, exestate.ControlledError("Se detecto radio negativo (geometry.IsCirclesIntersectionPossible)")
	} else if circleA.Position().X == circleB.Position().X && circleA.Position().Y == circleB.Position().Y {
		if circleA.Radious() != circleB.Radious() {
			return false, exestate.Ok() // exestate.ControlledError("Sin interseccion: Los circulos se encuentran superpuestos con distinto radio")
		} else {
			return false, exestate.ControlledError("Infinitas intersecciones: Los circulos se encuentran superpuestos con igual radio (geometry.IsCirclesIntersectionPossible)")
		}
	} else if circleA.Position().DistanceTo(circleB.Position()) > (circleA.Radious() + circleB.Radious()) {
		return false, exestate.Ok() //exestate.ControlledError("Sin interseccion: Los circulos se encuentran muy separados entre si")
	}

	return true, exestate.Ok()
}
