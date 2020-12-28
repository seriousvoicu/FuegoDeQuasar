package db

import (
	"testing"

	"github.com/seriousvoicu/FuegoDeQuasar/exestate"
	"github.com/seriousvoicu/FuegoDeQuasar/vectors"
)

func TestGetWithName(t *testing.T) {
	var getTests = []struct {
		name    string
		spected int
	}{
		{"Kenobi", 1},
		{"Skywalker", 2},
		{"Sato", 3},
	}
	repo := SatellitiesRepo{}
	service := SatellitieService{Satrepo: &repo}

	for _, tt := range getTests {
		t.Run(tt.name, func(t *testing.T) {
			s := service.GetWithName(tt.name)

			if exestate.OnError(&service) {
				t.Errorf(service.GetState(true).UserError)
				return
			}

			if s.Name != tt.name || s.Id != tt.spected {
				t.Errorf("got %q, want %q", tt.spected, s.Id)
			}

		})
	}
}

func TestGetSatellitiePosition(t *testing.T) {
	var getPosTests = []struct {
		name    string
		spected vectors.Vector2
	}{
		{"Kenobi", vectors.Vector2{X: -500, Y: -200}},
		{"Skywalker", vectors.Vector2{X: 100, Y: -100}},
		{"Sato", vectors.Vector2{X: 500, Y: 100}},
	}

	repo := SatellitiesRepo{}
	service := SatellitieService{Satrepo: &repo}

	for _, tt := range getPosTests {
		t.Run(tt.name, func(t *testing.T) {
			s := service.GetSatellitiePosition(tt.name)

			if exestate.OnError(&service) {
				t.Errorf(service.GetState(true).UserError)
				return
			}

			eq, state := vectors.Equals(*s, tt.spected)

			if !state.IsOk() {
				t.Errorf(state.UserError)
				return
			}

			if !eq {
				strSpected := tt.spected.ToString()
				strGot := s.ToString()

				t.Errorf("got %q, want %q", strGot, strSpected)
			}

		})
	}

}

func TestGetAllSatellities(t *testing.T) {
	spected := 3
	repo := SatellitiesRepo{}
	service := SatellitieService{Satrepo: &repo}

	s := service.GetAllSatellities()

	if exestate.OnError(&service) {
		t.Errorf(service.GetState(true).UserError)
		return
	}

	if len(*s) != spected {
		t.Errorf("got %q, want %q", len(*s), spected)
	}
}
