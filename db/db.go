package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"strings"
)

var Db *sql.DB

//init - Init metod run before main.
func init() {
	var err error
	/*
	*credentials.txt must have the following structure to work:
	*host_ip
	*port
	*user
	*password
	*dbname
	*sslmode*/
	credentials, ferr := ioutil.ReadFile("config/dbconfig.config")
	Check(ferr)
	config := strings.Split(string(credentials), "\n")

	dbConfig := "host=" + config[0] +
		" port=" + config[1] +
		" user=" + config[2] +
		" password=" + config[3] +
		" dbname=" + config[4] +
		" sslmode=" + config[5]

	Db, err = sql.Open("postgres", dbConfig)
	Check(err)
}

func Check(e error) bool {
	if e != nil {
		log.Fatal(e)
		return false
	}
	return true
}
