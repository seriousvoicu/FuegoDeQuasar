package satellities

import (
	"encoding/json"
	"io"

	db "github.com/seriousvoicu/FuegoDeQuasar/db"
	"github.com/seriousvoicu/FuegoDeQuasar/exestate"
	"github.com/seriousvoicu/FuegoDeQuasar/vectors"
)

type Satmanager struct {
	cluster *satcluster

	currState *exestate.State
}

func (this *Satmanager) /*StateHandler.*/ RegisterState(state *exestate.State) {
	this.currState = state
}

func (this *Satmanager) /*StateHandler.*/ GetState(consume bool) *exestate.State {
	state := this.currState

	if consume {
		this.currState = nil
	}
	return state
}

func (this *Satmanager) SetClusterDistances(distances []float32) {
	if exestate.OnError(this) {
		return
	}

	if this.cluster == nil {
		this.RegisterState(exestate.ControlledError("El cluster no esta inicializado"))
	}

	this.cluster.setDistances(distances)

	if exestate.OnError(this.cluster) {
		this.RegisterState(this.cluster.GetState(true))
	}
}

func (this *Satmanager) SetClusterMessages(messages [][]string) {
	if exestate.OnError(this) {
		return
	}

	if this.cluster == nil {
		this.RegisterState(exestate.ControlledError("El cluster no esta inicializado"))
	}

	this.cluster.setMessages(messages)

	if exestate.OnError(this.cluster) {
		this.RegisterState(this.cluster.GetState(true))
	}
}

func (this *Satmanager) InstantiateClusterFromJson(jsonInput io.Reader, dbRepo db.SatellitiesRepoInterface) {
	if exestate.OnError(this) {
		return
	}

	var cluster satcluster

	err := json.NewDecoder(jsonInput).Decode(&cluster)
	if err != nil {
		this.RegisterState(exestate.UncontrolledError("Error en el procesamiento del json", err))
		return
	}
	service := db.SatellitieService{Satrepo: dbRepo}
	for i := 0; i < cluster.count(); i++ {
		pos := service.GetSatellitiePosition(cluster.getAt(i).Name)

		if exestate.OnError(&service) {
			this.RegisterState(service.GetState(true))
			return
		}

		cluster.getAt(i).Pos = pos
	}

	this.cluster = &cluster

}

func (this *Satmanager) InstantiateClusterFromDB(dbRepo db.SatellitiesRepoInterface) {
	if exestate.OnError(this) {
		return
	}

	service := db.SatellitieService{Satrepo: dbRepo}

	rows := service.GetAllSatellities()
	if exestate.OnError(&service) {
		this.RegisterState(service.GetState(true))
		return
	}

	if rows == nil || len(*rows) <= 0 {
		this.RegisterState(exestate.ControlledError("No se pudieron obtener satelites (satellities.satmanager.InstantiateClusterFromDB)"))
		return
	}

	var cluster satcluster
	cluster.Satellities = make([]satellitie, len(*rows)) //[count]Satellitie //Init array

	for i := 0; i < len(*rows); i++ {
		cluster.Satellities[i] = satellitie{Name: (*rows)[i].Name, Pos: &vectors.Vector2{X: float64((*rows)[i].Position_x), Y: float64((*rows)[i].Position_y)}}
	}

	this.cluster = &cluster
}

func (this *Satmanager) GetMessage() string {
	if this.cluster == nil {
		this.RegisterState(exestate.ControlledError("Cluster no inicializado (satellities.Satmanager.GetMessage)"))
		return ""
	}

	if exestate.OnError(this) {
		return ""
	}

	msg := this.cluster.getMessage()

	if exestate.OnError(this.cluster) {
		this.RegisterState(this.cluster.GetState(true))
		return ""
	}

	return msg
}

func (this *Satmanager) GetLocation() vectors.Vector2 {

	if this.cluster == nil {
		this.RegisterState(exestate.ControlledError("Cluster no inicializado (satellities.Satmanager.GetMessage)"))
		return vectors.GetEmptyVector2()
	}

	if exestate.OnError(this) {
		return vectors.GetEmptyVector2()
	}

	pos := this.cluster.getLocation()

	if exestate.OnError(this.cluster) {
		this.RegisterState(this.cluster.GetState(true))
		return vectors.GetEmptyVector2()
	}

	return pos
}
