package satellities

import (
	"io"
	"log"
	"strconv"
	"testing"

	db "github.com/seriousvoicu/FuegoDeQuasar/db"
	"github.com/seriousvoicu/FuegoDeQuasar/exestate"
	vectors "github.com/seriousvoicu/FuegoDeQuasar/vectors"
)

type testIoReader struct {
	jsonTxt string

	data []byte

	readIndex int64
	testID    int
}

func (this *testIoReader) Read(b []byte) (n int, err error) {
	if this.data == nil {
		//var satArray []satellitie
		log.Printf("test %s", strconv.Itoa(this.testID))
		if this.testID == 0 {
			/*satArray := []satellitie{
				satellitie{Name: "test1", Pos: &vectors.Vector2{X: 5, Y: 5}},
				satellitie{Name: "test2", Pos: &vectors.Vector2{X: 10, Y: 12}},
				satellitie{Name: "test3", Pos: &vectors.Vector2{X: 16, Y: 18}},
			}

			satClis := satcluster{Satellities: satArray}

			bta, _ := json.Marshal(satClis)

			this.data = bta*/
			//log.Printf("test %s", this.data)
			this.data = []byte(this.jsonTxt)
			/*	`{"satellities":
				[
				   {"name":"kenobi","distance":100.0,"message":["este","","","mensaje",""]},
				   {"name":"skywalker","distance":115.5,"message":["","es","","","secreto"]},
				   {"name":"sato","distance":142.7,"message":["este","","un","",""]}
				]}`,
			)*/
		} else {
			//bta := []byte(`{"name":"prueba"}`)
		}

	}

	if this.readIndex >= int64(len(this.data)) {
		err = io.EOF
		return
	}

	n = copy(b, this.data[this.readIndex:])
	this.readIndex += int64(n)
	return
	/*

		this.currIdx += 1

		return this.currIdx - 1, nil*/
}

type testSatellitiesRepo struct {
	currState *exestate.State

	testID int
}

func (this *testSatellitiesRepo) /*StateHandler.*/ RegisterState(state *exestate.State) {
	this.currState = state
}

func (this *testSatellitiesRepo) /*StateHandler.*/ GetState(consume bool) *exestate.State {
	state := this.currState

	if consume {
		this.currState = nil
	}
	return state
}

func (this *testSatellitiesRepo) SatellitiesCount() *int {
	return nil
}

func (this *testSatellitiesRepo) GetAllSatellities() *[]db.SatellitiesRow {
	if this.testID == 0 {
		return &[]db.SatellitiesRow{
			db.SatellitiesRow{
				Id:         0,
				Name:       "test1",
				Position_x: 5,
				Position_y: 5,
			},
			db.SatellitiesRow{
				Id:         0,
				Name:       "test2",
				Position_x: 10,
				Position_y: 12,
			},
			db.SatellitiesRow{
				Id:         0,
				Name:       "test3",
				Position_x: 16,
				Position_y: 18,
			},
		}
	}

	return nil
}

func (this *testSatellitiesRepo) GetWithName(name string) *db.SatellitiesRow {
	if this.testID == 0 {
		return &db.SatellitiesRow{
			Id:         0,
			Name:       "test1",
			Position_x: 5,
			Position_y: 5,
		}
	}

	return nil
}

func satArrayToString(arrayA []satellitie) string {

	str := ""

	for ii := 0; ii < len(arrayA); ii++ {
		str += " { " + arrayA[ii].ToString() + " } "
	}
	str += " - "

	return str
}

func satArrayEquals(arrayA []satellitie, arrayB []satellitie) bool {

	if len(arrayA) != len(arrayB) {
		return false
	}

	for ii := 0; ii < len(arrayA); ii++ {

		if !satEqual(&arrayA[ii], &arrayB[ii]) {
			return false
		}
	}

	return true
}

func satEqual(satA *satellitie, satB *satellitie) bool {

	if satA == nil || satB == nil {
		return false
	}

	posEqual, _ := vectors.Equals(*satA.Pos, *satB.Pos)

	if satA.Name == satB.Name && posEqual {
		return true
	}

	return false
}

/*TESTINGS AQUI ABAJO*/

//func (this *Satmanager) InstantiateClusterFromDB(dbRepo db.SatellitiesRepoInterface) {
/*func TestInstantiateClusterFromDB(t *testing.T) {
	var testTable = []struct {
		name     string
		testCode int
		spected  []satellitie
		errMsg   string
	}{
		{
			"TestA satelites normal",
			0,
			[]satellitie{
				satellitie{Name: "test1", Pos: &vectors.Vector2{X: 5, Y: 5}},
				satellitie{Name: "test2", Pos: &vectors.Vector2{X: 10, Y: 12}},
				satellitie{Name: "test3", Pos: &vectors.Vector2{X: 16, Y: 18}},
			},
			"",
		},
	}

	manager := Satmanager{}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {

			manager.InstantiateClusterFromDB(&testSatellitiesRepo{})

			if exestate.OnError(&manager) {
				t.Errorf(manager.GetState(true).UserError)
				return
			}

			if !satArrayEquals(tt.spected, manager.cluster.Satellities) {
				t.Errorf("got %q, want %q", satArrayToString(manager.cluster.Satellities), satArrayToString(tt.spected))
			}

		})
	}

}*/

//func (this *Satmanager) InstantiateClusterFromJson(jsonInput io.Reader, dbRepo db.SatellitiesRepoInterface) {
/*func TestInstantiateClusterFromJson(t *testing.T) {
	var testTable = []struct {
		name     string
		testCode int
		spected  []satellitie
		errMsg   string
	}{
		{
			"TestA satelites normal",
			0,
			[]satellitie{
				satellitie{Name: "test1", Distance: 5},
				satellitie{Name: "test2", Distance: 5},
				satellitie{Name: "test3", Distance: 5},
			},
			"",
		},
	}

	manager := Satmanager{}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {

			manager.InstantiateClusterFromJson(&testIoReader{testID: tt.testCode, readIndex: 0}, &testSatellitiesRepo{})

			if exestate.OnError(&manager) {
				t.Errorf(manager.GetState(true).UserError)
				return
			}

			if !satArrayEquals(tt.spected, manager.cluster.Satellities) {
				t.Errorf("got %q, want %q", satArrayToString(manager.cluster.Satellities), satArrayToString(tt.spected))
			}

		})
	}

}*/

//func (this *Satmanager) SetClusterMessages(messages [][]string) {

//func (this *Satmanager) SetClusterDistances(distances []float32) {

//func (this *Satmanager) GetMessage() *string {
func TestGetMessage(t *testing.T) {
	var testTable = []struct {
		name     string
		managers Satmanager
		spected  string
		errMsg   string
	}{
		{
			"Test ok",
			Satmanager{cluster: &satcluster{Satellities: []satellitie{
				satellitie{Message: []string{"a", "", ""}},
				satellitie{Message: []string{"", "b", ""}},
				satellitie{Message: []string{"", "", "c"}},
			}}},
			"a b c",
			"",
		},
		{
			"Celda sin poder determinar",
			Satmanager{cluster: &satcluster{Satellities: []satellitie{
				satellitie{Message: []string{"", "a", "", ""}},
				satellitie{Message: []string{"", "", "b", ""}},
				satellitie{Message: []string{"", "", "", "c"}},
			}}},
			"",
			"No se pudo determinar el mensaje (satellities.satcluster.getMessage)",
		},
		{
			"Celda sin poder determinar",
			Satmanager{cluster: &satcluster{Satellities: []satellitie{
				satellitie{Message: []string{"", "a", " ", ""}},
				satellitie{Message: []string{"", " ", "b", ""}},
				satellitie{Message: []string{"", "", "", "c"}},
			}}},
			"",
			"No se pudo determinar el mensaje (satellities.satcluster.getMessage)",
		},
		{
			"Celda sin poder determinar",
			Satmanager{cluster: &satcluster{Satellities: []satellitie{
				satellitie{Message: []string{"", "a", " ", ""}},
				satellitie{Message: []string{"", " ", "b", ""}},
				satellitie{Message: nil},
			}}},
			"",
			"No se pudo determinar el mensaje (satellities.satcluster.getMessage)",
		},
		{
			"Celda sin poder determinar",
			Satmanager{cluster: &satcluster{Satellities: []satellitie{
				satellitie{Message: []string{"", "a", " ", "/"}},
				satellitie{Message: []string{"", "-", "b", ""}},
				satellitie{Message: []string{"", " ", "b", ""}},
			}}},
			"",
			"No se pudo determinar el mensaje (satellities.satcluster.getMessage)",
		},
		{
			"Celda sin poder determinar",
			Satmanager{cluster: &satcluster{Satellities: []satellitie{
				satellitie{Message: []string{"a", "", "", ""}},
				satellitie{Message: []string{"", "b", ""}},
				satellitie{Message: []string{"", "", "c"}},
			}}},
			"",
			"No se pudo determinar el mensaje (satellities.satcluster.getMessage)",
		},
		{
			"Celda sin poder determinar",
			Satmanager{},
			"",
			"Cluster no inicializado (satellities.Satmanager.GetMessage)",
		},
	}

	var str string

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {

			str = tt.managers.GetMessage()

			if exestate.OnError(&tt.managers) {
				if tt.managers.GetState(false).UserError != tt.errMsg {
					t.Errorf("got %q, want %q", tt.managers.GetState(true).UserError, tt.errMsg)
					return
				}
				return
			}

			if str != tt.spected {
				t.Errorf("got %q, want %q", str, tt.spected)
			}

		})
	}
}

//func (this *Satmanager) GetLocation() *vectors.Vector2 {
func TestGetLocation(t *testing.T) {
	var testTable = []struct {
		name     string
		managers Satmanager
		spected  vectors.Vector2
		errMsg   string
	}{
		{
			"Test superpuestos",
			Satmanager{cluster: &satcluster{Satellities: []satellitie{
				satellitie{Pos: &vectors.Vector2{X: 5, Y: 5}, Distance: 5},
				satellitie{Pos: &vectors.Vector2{X: 5, Y: 5}, Distance: 5},
				satellitie{Pos: &vectors.Vector2{X: 5, Y: 5}, Distance: 5},
			}}},
			vectors.GetEmptyVector2(),
			"Infinitas intersecciones: Los circulos se encuentran superpuestos con igual radio (geometry.IsCirclesIntersectionPossible)",
		},
		{
			"Test ok",
			Satmanager{cluster: &satcluster{Satellities: []satellitie{
				satellitie{Pos: &vectors.Vector2{X: 5, Y: 0}, Distance: 5},
				satellitie{Pos: &vectors.Vector2{X: -5, Y: 0}, Distance: 5},
				satellitie{Pos: &vectors.Vector2{X: 0, Y: -5}, Distance: 5},
			}}},
			vectors.Vector2{X: 0, Y: 0},
			"",
		},
	}

	var pos vectors.Vector2

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {

			pos = tt.managers.GetLocation()

			if exestate.OnError(&tt.managers) {
				if tt.errMsg == "" || tt.managers.GetState(false).UserError != tt.errMsg {
					t.Errorf("got %q, want %q", tt.managers.GetState(true).UserError, tt.errMsg)
					return
				}

				return
			}

			equals, _ := vectors.Equals(pos, tt.spected)

			if !equals {
				t.Errorf("got %q, want %q", pos.ToString(), tt.spected.ToString())
			}

		})
	}
}
