package lib

import (
	"database/sql"
	"errors"
	"strconv"

	"fmt"
	"log"
	"onboarding/models"
	_ "onboarding/models"

	_ "github.com/go-sql-driver/mysql"
)

var DBUser string = "dbuser"
var DBPassword string = "dbpw"
var DBHost string = "docker"
var DBPort string = "3306"
var DBName string = "testdb"

func Initdb() {

	db, err := sql.Open("mysql", DBUser+":"+DBPassword+"@tcp("+DBHost+":"+DBPort+")/"+DBName)

	if err != nil {
		log.Fatal(4, "ERROR: DB Connection: "+err.Error())
	} else {
		//res, err := db.Query("select * from bla;")

		res, err := db.Query(`
		CREATE TABLE IF NOT EXISTS Firmen (
			Id INT auto_increment,
			Name VARCHAR(255),
			Enabled INT,			
			primary key (id)
		);
		`)
		if err != nil {
			//panic(err)
			//log.Fatal(err)
			fmt.Println("an error occourred: %S\n", err.Error())

		} else {
			fmt.Printf("create table: %s\n", res.Err())
		}

	}
	defer db.Close()
}

func InsertFirmen() {
	db, err := sql.Open("mysql", DBUser+":"+DBPassword+"@tcp("+DBHost+":"+DBPort+")/"+DBName)

	res2, err := db.Query(`
	insert into Firmen (			
		Name,
		Enabled) values (
			"TestFirma", "1"
		);
	`)

	CheckErr(err)
	fmt.Printf("insert into: %s\n", res2.Err())
}

func GetFirmen() (models.Firmen, error) {
	fmt.Printf("GetFirmen:\n")

	firmen := models.Firmen{}

	db, err := sql.Open("mysql", DBUser+":"+DBPassword+"@tcp("+DBHost+":"+DBPort+")/"+DBName)

	if err != nil {
		fmt.Printf("ERROR: DB Connection: %S \n" + err.Error())
		return firmen, err
		//		log.Fatal(4, "ERROR: DB Connection: "+err.Error())

	} else {

		rows, err := db.Query("select * from Firmen;")
		CheckErr(err)
		//fmt.Printf("DB ok: %s\n", rows)
		for rows.Next() {
			itm := models.Firma{}
			err = rows.Scan(&itm.Id, &itm.Name, &itm.Enabled)
			fmt.Printf("\nGetFirmen -- Got: %s", itm)
			firmen.Firmen = append(firmen.Firmen, itm)

		}
	}
	defer db.Close()
	if len(firmen.Firmen) == 0 {
		return firmen, errors.New(`{"id":1,"error":"Emtpy sQL Result"}`)
	}
	return firmen, nil
}

func DelFirma(id int) error {
	fmt.Printf("DelFirma got to delete Id: %d\n", id)
	db, err := sql.Open("mysql", DBUser+":"+DBPassword+"@tcp("+DBHost+":"+DBPort+")/"+DBName)
	if err != nil {
		fmt.Printf("ERROR: DB Connection: %S \n" + err.Error())
		return err
		//		log.Fatal(4, "ERROR: DB Connection: "+err.Error())

	} else {

		_, err := db.Query("delete from Firmen where Id = " + strconv.Itoa(id) + ";")
		if err != nil {
			fmt.Printf("db del error: %s\n", err)
			return err
		}
	}
	defer db.Close()
	return err
}
