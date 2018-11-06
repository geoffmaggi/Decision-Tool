package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/astaxie/beego/config"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

// InitDatabase initalizes the mysql database
// builds the schema of the table
// forigen key restriction is not handled in here but they're
// handled in each objects Save and Destroy methods
func InitDatabase(conf config.Configer) *gorp.DbMap {
	dbsrc := fmt.Sprintf("%s:%s@/%s",
		conf.String("database::user"),
		conf.String("database::password"),
		conf.String("database::name"))

	db, err := sql.Open("mysql", dbsrc)
	if err != nil {
		log.Fatalf("Unable to connect to mysql : %#v\n", err)
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}

	dbmap.AddTableWithName(Person{}, "person").SetKeys(true, "person_id")
	dbmap.AddTableWithName(Decision{}, "decision").SetKeys(true, "decision_id")
	dbmap.AddTableWithName(Ballot{}, "ballot").SetKeys(true, "ballot_id")
	dbmap.AddTableWithName(Alternative{}, "alternative").SetKeys(true, "alternative_id")
	dbmap.AddTableWithName(Criterion{}, "criterion").SetKeys(true, "criterion_id")
	dbmap.AddTableWithName(Vote{}, "vote")
	dbmap.AddTableWithName(Rating{}, "rating")

	if err = dbmap.CreateTablesIfNotExists(); err != nil {
		log.Fatalln(err)
	}

	// Always create an admin account
	_, err = dbmap.Exec("DELETE FROM person WHERE person_id='0'")
	if err != nil {
		log.Fatalln(err)
	}

	hashed := HashPassword(conf.String("admin::password"))
	_, err = dbmap.Exec("INSERT INTO person VALUES('-1',?,?,?,?)",
		conf.String("admin::email"),
		hashed,
		conf.String("admin::name_first"),
		conf.String("admin::name_last"))
	if err != nil {
		log.Fatalln(err)
	}

	_, err = dbmap.Exec("UPDATE person set person_id='0' WHERE person_id='-1'")
	if err != nil {
		log.Fatalln(err)
	}

	return dbmap
}
