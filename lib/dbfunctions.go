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
			fmt.Printf("Init-DB completed: %s\n", res.Err())
		}

	}
	defer db.Close()
}

func InsertFirmen(fi models.NewCompany) {
	db, err := sql.Open("mysql", DBUser+":"+DBPassword+"@tcp("+DBHost+":"+DBPort+")/"+DBName)

	fmt.Printf("insertFirma got Name %s\n", fi.Name)
	fmt.Printf("insertFirma got Enabled %s\n", strconv.Itoa(fi.Enabled))
	fmt.Println(fi.Name)
	res2, err := db.Query(`
	insert into Firmen (			
		Name,
		Enabled) values (
			"` + fi.Name + `" , " ` + strconv.Itoa(fi.Enabled) + `"
		);
	`)

	CheckErr(err)
	fmt.Printf("insert into: %s\n", res2.Err())
}

func GetFirmen() (models.Firmen, error) {
	//fmt.Printf("GetFirmen:\n")

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
			//fmt.Printf("\nGetFirmen -- Got: %s", itm)
			firmen.Firmen = append(firmen.Firmen, itm)

		}
	}
	defer db.Close()
	if len(firmen.Firmen) == 0 {
		return firmen, errors.New(`{"id":1,"error":"Emtpy sQL Result"}`)
	}
	return firmen, nil
}

func CheckCookie(id string) (models.Cookie, error) {

	db, err := sql.Open("mysql", DBUser+":"+DBPassword+"@tcp("+DBHost+":"+DBPort+")/"+DBName)
	var c models.Cookie

	fmt.Println("Testing cookieId: ", id)
	if err != nil {
		fmt.Printf("ERROR: DB Connection: %S \n" + err.Error())
		return c, err
		//		log.Fatal(4, "ERROR: DB Connection: "+err.Error())

	} else {

		err := db.QueryRow("select * from cookies where Id = \""+id+"\";").Scan(&c.Id, &c.Status)
		if err != nil && err != sql.ErrNoRows {
			fmt.Printf("ERROR: DB Select: %S \n" + err.Error())

			return c, err
		} else if err == sql.ErrNoRows {
			fmt.Println("Requested cookieId not found")
			return c, errors.New("CookieNotFound")
		} else {
			fmt.Println("checkCookie SQL Result: ", c.Id, c.Status)

			return c, nil
		}
	}
	//defer db.Close()
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
