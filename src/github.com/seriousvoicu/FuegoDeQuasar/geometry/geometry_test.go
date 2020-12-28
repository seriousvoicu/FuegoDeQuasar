package geometry

import (
	"strconv"
	"testing"

	"github.com/seriousvoicu/FuegoDeQuasar/exestate"
	"github.com/seriousvoicu/FuegoDeQuasar/vectors"
)

type testCircle struct {
	pos *vectors.Vector2
	rad float64

	currState *exestate.State
}

func (this *testCircle) Radious() float64 {
	return this.rad
}
func (this *testCircle) Position() *vectors.Vector2 {
	return this.pos
}

func (this *testCircle) /*StateHandler.*/ RegisterState(state *exestate.State) {
	this.currState = state
}

func (this *testCircle) /*StateHandler.*/ GetState(consume bool) *exestate.State {
	state := this.currState

	if consume {
		this.currState = nil
	}
	return state
}

//func GetCirclesIntersections(circleA Circle, circleB Circle) (*vectors.Vector2, *vectors.Vector2, *exestate.State) {
func TestGetCirclesIntersections(t *testing.T) {
	var testTable = []struct {
		name    string
		circleA *testCircle
		circleB *testCircle
		pointA  vectors.Vector2
		pointB  vectors.Vector2
		errMsg  string
	}{
		{"Muy separados", &testCircle{pos: &vectors.Vector2{X: 5, Y: 5}, rad: 5}, &testCircle{pos: &vectors.Vector2{X: -5, Y: -5}, rad: 5}, vectors.GetEmptyVector2(), vectors.GetEmptyVector2(), "No existe interseccion (geometry.IsCirclesIntersectionPossible)"},
		{"Interseccion", &testCircle{pos: &vectors.Vector2{X: 5, Y: 5}, rad: 10}, &testCircle{pos: &vectors.Vector2{X: -5, Y: -5}, rad: 5}, vectors.Vector2{X: -0.22140543, Y: -3.52859457}, vectors.Vector2{X: -3.52859457, Y: -0.22140543}, ""},
		{"Misma posicion distinto radio", &testCircle{pos: &vectors.Vector2{X: 5, Y: 5}, rad: 10}, &testCircle{pos: &vectors.Vector2{X: 5, Y: 5}, rad: 5}, vectors.GetEmptyVector2(), vectors.GetEmptyVector2(), "No existe interseccion (geometry.IsCirclesIntersectionPossible)"},
		{"Misma posicion y radio", &testCircle{pos: &vectors.Vector2{X: 5, Y: 5}, rad: 5}, &testCircle{pos: &vectors.Vector2{X: 5, Y: 5}, rad: 5}, vectors.GetEmptyVector2(), vectors.GetEmptyVector2(), "Infinitas intersecciones: Los circulos se encuentran superpuestos con igual radio (geometry.IsCirclesIntersectionPossible)"},
		{"Radio negativo", &testCircle{pos: &vectors.Vector2{X: 5, Y: 5}, rad: -1}, &testCircle{pos: &vectors.Vector2{X: 5, Y: 5}, rad: 5}, vectors.GetEmptyVector2(), vectors.GetEmptyVector2(), "Se detecto radio negativo (geometry.IsCirclesIntersectionPossible)"},
		{"Radio 0", &testCircle{pos: &vectors.Vector2{X: 5, Y: 5}, rad: 0}, &testCircle{pos: &vectors.Vector2{X: 5, Y: 5}, rad: 5}, vectors.GetEmptyVector2(), vectors.GetEmptyVector2(), "Los circulos no poseen radio (geometry.IsCirclesIntersectionPossible)"},
		{"Interseccion 2", &testCircle{pos: &vectors.Vector2{X: 46, Y: -250}, rad: 100}, &testCircle{pos: &vectors.Vector2{X: -100, Y: -50}, rad: 320}, vectors.Vector2{X: 145.88502500, Y: -254.79393175}, vectors.Vector2{X: 20.12932689, Y: -346.59559137}, ""},
		{"Interseccion un punto", &testCircle{pos: &vectors.Vector2{X: 5, Y: 0}, rad: 5}, &testCircle{pos: &vectors.Vector2{X: -5, Y: 0}, rad: 5}, vectors.Vector2{X: 0, Y: 0}, vectors.GetEmptyVector2(), ""},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {

			resultA, resultB, state := GetCirclesIntersections(tt.circleA, tt.circleB)

			if !state.IsOk() {
				if tt.errMsg != state.UserError {
					t.Errorf(state.UserError)
					return
				}
			} else {

				equalAA, _ := vectors.Equals(resultA, tt.pointA)

				/*	if !state.IsOk() {
					t.Errorf(state.UserError)
					return
				}*/

				equalBB, _ := vectors.Equals(resultB, tt.pointB)

				/*if !state.IsOk() {
					t.Errorf(state.UserError)
					return
				}*/

				equalAB, _ := vectors.Equals(resultA, tt.pointB)

				/*	if !state.IsOk() {
					t.Errorf(state.UserError)
					return
				}*/

				equalBA, _ := vectors.Equals(resultB, tt.pointA)

				/*if !state.IsOk() {
					t.Errorf(state.UserError)
					return
				}*/

				if !equalAA && !equalBB && !equalAB && !equalBA {
					t.Errorf("got %q, want %q", resultA.ToString()+" - "+resultB.ToString(), tt.pointA.ToString()+" - "+tt.pointB.ToString())
				}

			}

		})
	}
}

//func IsCirclesIntersectionPossible(circleA Circle, circleB Circle) (bool, *exestate.State) {
func TestIsCirclesIntersectionPossible(t *testing.T) {
	var testTable = []struct {
		name    string
		circleA *testCircle
		circleB *testCircle
		spected string
		errMsg  string
	}{
		{"Muy separados", &testCircle{pos: &vectors.Vector2{X: 5, Y: 5}, rad: 5}, &testCircle{pos: &vectors.Vector2{X: -5, Y: -5}, rad: 5}, "false", ""},
		{"Interseccion", &testCircle{pos: &vectors.Vector2{X: 5, Y: 5}, rad: 10}, &testCircle{pos: &vectors.Vector2{X: -5, Y: -5}, rad: 5}, "true", ""},
		{"Misma posicion distinto radio", &testCircle{pos: &vectors.Vector2{X: 5, Y: 5}, rad: 10}, &testCircle{pos: &vectors.Vector2{X: 5, Y: 5}, rad: 5}, "false", ""},
		{"Misma posicion y radio", &testCircle{pos: &vectors.Vector2{X: 5, Y: 5}, rad: 5}, &testCircle{pos: &vectors.Vector2{X: 5, Y: 5}, rad: 5}, "error", "Infinitas intersecciones: Los circulos se encuentran superpuestos con igual radio (geometry.IsCirclesIntersectionPossible)"},
		{"Radio negativo", &testCircle{pos: &vectors.Vector2{X: 5, Y: 5}, rad: -1}, &testCircle{pos: &vectors.Vector2{X: 5, Y: 5}, rad: 5}, "error", "Se detecto radio negativo (geometry.IsCirclesIntersectionPossible)"},
		{"Radio 0", &testCircle{pos: &vectors.Vector2{X: 5, Y: 5}, rad: 0}, &testCircle{pos: &vectors.Vector2{X: 5, Y: 5}, rad: 5}, "error", "Los circulos no poseen radio (geometry.IsCirclesIntersectionPossible)"},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {

			result, state := IsCirclesIntersectionPossible(tt.circleA, tt.circleB)

			if !state.IsOk() {
				if tt.spected != "error" {
					t.Errorf(state.UserError)
					return
				} else if tt.spected == "error" && tt.errMsg != state.UserError {
					t.Errorf("got %q, want %q", state.UserError, tt.errMsg)
				}

			} else if strconv.FormatBool(result) != tt.spected {
				t.Errorf("got %q, want %q", strconv.FormatBool(result), tt.spected)
			}

		})
	}
}
