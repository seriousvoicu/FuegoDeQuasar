package mathext

import (
	"core"
	"math"
)

type Circle interface {
	Radious() float64
	Position() *core.Vector2
}

//Retorno los puntos de interseccion. Puede haber 2, 1 o ninguno.
func GetIntersection(circleA Circle, circleB Circle) (pointA *core.Vector2, pointB *core.Vector2) {

	var d = GetDistance(circleA.Position(), circleB.Position()) // math.Hypot(posTwo.x-posOne.x, posTwo.y-posOne.y) //Distancia entre ambos puntos

	if d <= circleA.Radious()+circleB.Radious() && d >= math.Abs(circleB.Radious()-circleA.Radious()) {

		var ex = (circleB.Position().X - circleA.Position().X) / d
		var ey = (circleB.Position().Y - circleA.Position().Y) / d

		var x = (circleA.Radious()*circleA.Radious() - circleB.Radious()*circleB.Radious() + d*d) / (2 * d)
		var y = math.Sqrt(circleA.Radious()*circleA.Radious() - x*x)

		var P1, P2 *core.Vector2
		P1 = &core.Vector2{X: RoundZero(circleA.Position().X + x*ex - y*ey), Y: RoundZero(circleA.Position().Y + x*ey + y*ex)}
		P2 = &core.Vector2{X: RoundZero(circleA.Position().X + x*ex + y*ey), Y: RoundZero(circleA.Position().Y + x*ey - y*ex)}

		return P1, P2
	}

	return nil, nil

	/*var d = GetDistance(posOne, posTwo) // math.Hypot(posTwo.x-posOne.x, posTwo.y-posOne.y) //Distancia entre ambos puntos

	if d <= radOne+radTwo && d >= math.Abs(radTwo-radOne) {

		var ex = (posTwo.X - posOne.X) / d
		var ey = (posTwo.Y - posOne.Y) / d

		var x = (radOne*radOne - radTwo*radTwo + d*d) / (2 * d)
		var y = math.Sqrt(radOne*radOne - x*x)

		var P1, P2 *objects.Vector2
		P1 = &objects.Vector2{X: CapZero(posOne.X + x*ex - y*ey), Y: CapZero(posOne.Y + x*ey + y*ex)}
		P2 = &objects.Vector2{X: CapZero(posOne.X + x*ex + y*ey), Y: CapZero(posOne.Y + x*ey - y*ex)}

		return P1, P2
	}

	return nil, nil*/
}

func GetDistance(posA *core.Vector2, posB *core.Vector2) float64 {
	return math.Hypot(posB.X-posA.X, posB.Y-posA.Y)
}

func RoundZero(v float64) float64 {
	if v <= 0.0001 {
		v = 0
	}

	return v
}
