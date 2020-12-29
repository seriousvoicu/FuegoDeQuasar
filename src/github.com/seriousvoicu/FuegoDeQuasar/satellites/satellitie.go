package satellites

import (
	"strconv"

	"github.com/seriousvoicu/FuegoDeQuasar/arrays"
	"github.com/seriousvoicu/FuegoDeQuasar/exestate"
	"github.com/seriousvoicu/FuegoDeQuasar/vectors"
)

type satellite struct {
	Name     string   `json:"name"`
	Distance float64  `json:"distance"`
	Message  []string `json:"message"`
	Pos      *vectors.Vector2
}

func (this satellite) /*Circle.*/ Radious() float64 {
	return this.Distance
}

func (this satellite) /*Circle.*/ Position() *vectors.Vector2 {
	return this.Pos
}

func (this *satellite) ToString() string {
	str := ""
	str += "Name: " + this.Name + " - "
	str += "Distancia: " + strconv.FormatFloat(this.Distance, 'f', -1, 64) + " - "

	result, state := arrays.StringArrayToString(this.Message)

	if state.IsOk() {
		str += "Mensaje: " + result + " - "
	} else {
		str += "Mensaje: - "
	}

	if this.Pos != nil {
		result = this.Pos.ToString()

		if !exestate.OnError(this.Pos) {
			str += "Pos: " + result + " - "
		} else {
			str += "Pos: - "
		}
	} else {
		str += "Pos: - "
	}

	return str
}
