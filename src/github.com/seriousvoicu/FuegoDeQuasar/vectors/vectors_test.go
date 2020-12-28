package vectors

import (
	"math"
	"strconv"
	"testing"

	"github.com/seriousvoicu/FuegoDeQuasar/exestate"
)

func TestDistanceTo(t *testing.T) {
	var testTable = []struct {
		name    string
		posA    Vector2
		posB    Vector2
		spected float64
		errMsg  string
	}{
		{"TestA", Vector2{X: 0, Y: 0}, Vector2{X: 20, Y: 15}, 25, ""},
		{"Testb", Vector2{X: 20.5552, Y: 20.12345}, Vector2{X: 20.5525, Y: 20.12360}, 0.002704, ""},
		{"TestC", Vector2{X: -20.5552, Y: 20.12345}, Vector2{X: 20.5525, Y: -20.12360}, 57.529714, ""},
		{"TestD", GetEmptyVector2(), Vector2{X: 20.5525, Y: 20.12360}, -1, "Es un vector vacio (vectors.Vector2.DistanceTo)"},
		{"TestE", Vector2{X: 20.5525, Y: 20.12360}, GetEmptyVector2(), -1, "Es un vector vacio (vectors.Vector2.DistanceTo)"},
		{"TestF", GetEmptyVector2(), GetEmptyVector2(), -1, "Es un vector vacio (vectors.Vector2.DistanceTo)"},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {

			result := tt.posA.DistanceTo(&tt.posB)

			if exestate.OnError(&tt.posA) {
				if tt.spected != -1 {
					t.Errorf(tt.posA.GetState(true).UserError)
					return
				} else if tt.spected == -1 && tt.errMsg != tt.posA.GetState(false).UserError {
					t.Errorf("got %q, want %q", tt.posA.GetState(true).UserError, tt.errMsg)
				}

			} else if math.Abs(result-tt.spected) > 0.000001 {
				t.Errorf("got %q, want %q", strconv.FormatFloat(result, 'f', -1, 64), strconv.FormatFloat(tt.spected, 'f', -1, 64))
			}

		})
	}
}

func TestToString(t *testing.T) {
	var testTable = []struct {
		name    string
		posA    Vector2
		spected string
		errMsg  string
	}{
		{"TestA", Vector2{X: 20, Y: 20}, "X: 20 Y: 20", ""},
		{"Testb", Vector2{X: 20.5552, Y: 20.12345}, "X: 20.5552 Y: 20.12345", ""},
		{"TestC", GetEmptyVector2(), "error", "Es un vector vacio (vectors.Vector2.ToString)"},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {

			result := tt.posA.ToString()

			if exestate.OnError(&tt.posA) {
				if tt.spected != "error" {
					t.Errorf(tt.posA.GetState(true).UserError)
					return
				} else if tt.spected == "error" && tt.errMsg != tt.posA.GetState(false).UserError {
					t.Errorf("got %q, want %q", tt.posA.GetState(true).UserError, tt.errMsg)
				}

			} else if result != tt.spected {
				t.Errorf("got %q, want %q", result, tt.spected)
			}

		})
	}
}

func TestEquals(t *testing.T) {
	var testTable = []struct {
		name    string
		posA    Vector2
		posB    Vector2
		spected string
		errMsg  string
	}{
		{"TestA", Vector2{X: 20, Y: 20}, Vector2{X: 20, Y: 20}, "true", ""},
		{"TestB", Vector2{X: 20, Y: 20}, Vector2{X: 20, Y: 25}, "false", ""},
		{"TestC", Vector2{X: 20.002, Y: 20.0005}, Vector2{X: 20.002, Y: 20.0005}, "true", ""},
		{"TestD", Vector2{X: 20, Y: 20.0104}, Vector2{X: 20, Y: 20.0005}, "false", ""},
		{"TestE", GetEmptyVector2(), GetEmptyVector2(), "error", "Almenos uno de los vectores se encuentra vacio (vectors.Vector2.Equals)"},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {

			equal, state := Equals(tt.posA, tt.posB)

			if !state.IsOk() {
				if tt.spected != "error" {
					t.Errorf(state.UserError)
					return
				} else if tt.spected == "error" && tt.errMsg != state.UserError {
					t.Errorf("got %q, want %q", state.UserError, tt.errMsg)
				}

			} else if strconv.FormatBool(equal) != tt.spected {
				t.Errorf("got %q, want %q", strconv.FormatBool(equal), tt.spected)
			}

		})
	}
}
