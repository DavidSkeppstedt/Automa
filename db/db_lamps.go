package db

import (
	"github.com/DavidSkeppstedt/Automa/model"
)

func FetchLamps() (result []model.Lamp, err error) {
	rows, err := Db.Query("SELECT * FROM lamps")
	defer rows.Close()
	for rows.Next() {
		var tmp model.Lamp
		rows.Scan(&tmp.Id, &tmp.Name, &tmp.Zone, &tmp.Lamp)

		result = append(result, tmp)
	}
	return result, err
}

func LampExists(lamp int) (bool, error) {
	rows, err := Db.Query("SELECT EXISTS (SELECT lamp FROM lamps WHERE lamp = $1 LIMIT 1)", lamp)
	defer rows.Close()
	rows.Next()
	var exist bool
	rows.Scan(&exist)
	return exist, err
}

func GetLamp(lamp int) (result model.Lamp, err error) {
	rows, err := Db.Query("SELECT * FROM lamps WHERE lamp = $1", lamp)
	defer rows.Close()
	rows.Next()
	var tmp model.Lamp
	rows.Scan(&tmp.Id, &tmp.Name, &tmp.Zone, &tmp.Lamp)
	return tmp, err
}

func AddLamp(aLamp model.Lamp) (err error) {
	_, err = Db.Query("INSERT INTO lamps (name,zone,lamp) VALUES($1,$2,$3)", aLamp.Name, aLamp.Zone, aLamp.Lamp)
	return
}
