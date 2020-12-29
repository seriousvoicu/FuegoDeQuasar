package db

import (
	"database/sql"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/seriousvoicu/FuegoDeQuasar/exestate"
)

type SatellitiesRepoInterface interface {
	GetAllSatellities() *[]SatellitiesRow
	SatellitiesCount() int
	GetWithName(name string) *SatellitiesRow
	exestate.StateHandler
}

type SatellitiesRepo struct {
	currState *exestate.State
}

func (this *SatellitiesRepo) /*StateHandler.*/ RegisterState(state *exestate.State) {
	this.currState = state
}

func (this *SatellitiesRepo) /*StateHandler.*/ GetState(consume bool) *exestate.State {
	state := this.currState

	if consume {
		this.currState = nil
	}
	return state
}

func (this *SatellitiesRepo) SatellitiesCount() int {
	if exestate.OnError(this) {
		return -1
	}

	var count int

	//La ruta para andar en el google app engine tiene que ser la siguiente: db/mainDB.db
	sqlDb, err := sql.Open("sqlite3", "mainDB.db")

	if err != nil {
		this.RegisterState(exestate.UncontrolledError("No se pudo abrir la base", err))
		return -1
	}

	defer sqlDb.Close()

	err = sqlDb.QueryRow("SELECT COUNT(*) FROM Satellities").Scan(&count)

	if err != nil {
		this.RegisterState(exestate.UncontrolledError("No se pudo leer de la base", err))
		return -1
	}

	return count
}

func (this *SatellitiesRepo) GetAllSatellities() *[]SatellitiesRow {
	if exestate.OnError(this) {
		return nil
	}

	count := this.SatellitiesCount()

	if exestate.OnError(this) {
		return nil
	}

	//La ruta para andar en el google app engine tiene que ser la siguiente: db/mainDB.db
	sqlDb, err := sql.Open("sqlite3", "mainDB.db")

	if err != nil {
		this.RegisterState(exestate.UncontrolledError("No se pudo abrir la base", err))
		return nil
	}

	defer sqlDb.Close()

	rows, err := sqlDb.Query("SELECT * FROM Satellities order by id")

	if err != nil {
		this.RegisterState(exestate.UncontrolledError("No se pudo leer de la base", err))
		return nil
	}

	allRows := make([]SatellitiesRow, count)

	var id int
	var name string
	var x float64
	var y float64
	var index int = 0
	for rows.Next() {
		rows.Scan(&id, &name, &x, &y)

		allRows[index] = SatellitiesRow{id, name, x, y}
		index++
	}

	return &allRows
}

func (this *SatellitiesRepo) GetWithName(name string) *SatellitiesRow {
	if exestate.OnError(this) {
		return nil
	}

	//La ruta para andar en el google app engine tiene que ser la siguiente: db/mainDB.db
	sqlDb, err := sql.Open("sqlite3", "mainDB.db")

	if err != nil {
		this.RegisterState(exestate.UncontrolledError("No se pudo abrir la base (db.db.GetWithName)", err))
		return nil
	}

	defer sqlDb.Close()

	var id int
	var rtname string
	var x float64
	var y float64
	err = sqlDb.QueryRow("SELECT id, name, position_x, position_y FROM Satellities where upper(name)=?", strings.ToUpper(name)).Scan(&id, &rtname, &x, &y)

	if err != nil {
		this.RegisterState(exestate.UncontrolledError("No se pudo leer de la base (db.db.GetWithName)", err))
		return nil
	}

	if &x == nil || &y == nil {
		this.RegisterState(exestate.ControlledError("No se pudieron obtener los datos (db.db.GetWithName)"))
		return nil
	}

	row := SatellitiesRow{id, rtname, x, y}

	return &row
}
