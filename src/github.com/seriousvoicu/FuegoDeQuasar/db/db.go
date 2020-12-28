package db

import (
	"github.com/seriousvoicu/FuegoDeQuasar/exestate"
	"github.com/seriousvoicu/FuegoDeQuasar/vectors"
)

//Model
type SatellitiesRow struct {
	Id         int
	Name       string
	Position_x float64
	Position_y float64
}

type SatellitieService struct {
	Satrepo   SatellitiesRepoInterface
	currState *exestate.State
}

func (this *SatellitieService) /*StateHandler.*/ RegisterState(state *exestate.State) {
	this.currState = state
}

func (this *SatellitieService) /*StateHandler.*/ GetState(consume bool) *exestate.State {
	state := this.currState

	if consume {
		this.currState = nil
	}
	return state
}

func (this *SatellitieService) GetSatellitiePosition(name string) *vectors.Vector2 {
	if this.Satrepo == nil {
		this.RegisterState(exestate.ControlledError("Repositorio no definido (db.db.GetSatellitiePosition)"))
	}

	row := this.Satrepo.GetWithName(name)

	if exestate.OnError(this.Satrepo) {
		this.RegisterState(this.Satrepo.GetState(true))
		return nil
	}

	if row == nil {
		this.RegisterState(exestate.ControlledError("No se pudo obtener fila (db.db.GetSatellitiePosition)"))
		return nil
	}

	vector := vectors.Vector2{X: row.Position_x, Y: row.Position_y, Empty: false}

	return &vector
}

func (this *SatellitieService) GetWithName(name string) *SatellitiesRow {
	row := this.Satrepo.GetWithName(name)

	if exestate.OnError(this.Satrepo) {
		this.RegisterState(this.Satrepo.GetState(true))
		return nil
	}

	return row
}

func (this *SatellitieService) GetAllSatellities() *[]SatellitiesRow {
	satarray := this.Satrepo.GetAllSatellities()

	if exestate.OnError(this.Satrepo) {
		this.RegisterState(this.Satrepo.GetState(true))
		return nil
	}

	return satarray
}

/*
func (this DataBaseManager) Close() {
	if exestate.OnError(this) {
		return
	}

	err := this.dtb.Close()

	if err != nil {
		this.RegisterState(exestate.UncontrolledError("Error interno", err))
	}
}

func Open() (*DataBaseManager, error) {
	sqlDb, err := sql.Open("sqlite3", "db/mainDB.db")

	if err != nil {
		return &DataBaseManager{dtb: nil}, err
	}

	dbm := DataBaseManager{dtb: sqlDb}

	return &dbm, nil
}*/
