package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"onboarding/lib"
	_ "onboarding/lib"
	"onboarding/models"
	"time"

	"github.com/gorilla/mux"

	// "database/sql"

	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
)

func test() models.Firma {
	testFirma := models.Firma{
		Id:      1,
		Name:    "LSI",
		Enabled: 1,
	}
	return testFirma
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	(*w).Header().Set("Access-Control-Allow-Headers", allowedHeaders)
}

func createFirma(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	reqBody, err := ioutil.ReadAll(r.Body)
	fmt.Printf("createFirma got: %s\n", reqBody)
	if err != nil {
		fmt.Printf("createFirma error: %s \n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		m := models.Firma{}
		err := json.Unmarshal(reqBody, &m)
		if err != nil {
			fmt.Printf("createFirma Error on json.unmarshal: %s", err)
			//http.Error(w, err.Error(), http.StatusInternalServerError)

		} else {
			//fmt.Printf("del-ID: %d\n", m.Id)
			lib.InsertFirmen(m)
		}

	}

}

func delFirma(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Printf("delFirma error: %s \n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

	} else {
		fmt.Println("got body: ", string(reqBody))
		m := models.Firma{}
		err := json.Unmarshal(reqBody, &m)
		if err != nil {
			fmt.Printf("delFirma Error on json.unmarshal: %s", err)
			//http.Error(w, err.Error(), http.StatusInternalServerError)

		} else {
			fmt.Printf("del-ID: %d\n", m.Id)
			lib.DelFirma(m.Id)
		}

	}

	//fmt.Printf("delFirma got: %s \n", r.Body)
}

func main() {

	//testFirma := test()
	//testFirmen := models.Firmen{}

	//testFirmen.Firmen = append(testFirmen.Firmen, testFirma)
	//testFirmen.Firmen = append(testFirmen.Firmen, testFirma)

	lib.Initdb()
	//testInsertDB(&testFirma)
	//testReadDB()

	// test - übernahme von rückgabewert von function
	err, test := lib.GetFirmen()
	fmt.Printf("got back: %s\n", test, err)
	//-------------

	router := mux.NewRouter()

	router.HandleFunc("/api/delFirma", delFirma)
	router.HandleFunc("/api/createFirma", createFirma)
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {

		//Allow CORS here By * or specific origin
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// an example API handler

		//json.NewEncoder(w).Encode(testFirmen)
		res, err2 := lib.GetFirmen()

		if err2 != nil {
			fmt.Printf("Error on api/health: %s", err2.Error())
			//json.NewEncoder(w).Encode(`{"errors":[` + err2.Error() + `]}`)

			txt := models.ErrorCustom{}
			txt.ErrorText = "empty sql response"
			json.NewEncoder(w).Encode(txt)
		} else {
			fmt.Printf("api/health returning to client: %s", res)
			json.NewEncoder(w).Encode(res)

		}

	})

	// Choose the folder to serve
	staticDir := "/frontend/dist"

	// Create the route
	router.
		PathPrefix("/").
		Handler(http.StripPrefix("/", http.FileServer(http.Dir("."+staticDir))))

	srv := &http.Server{
		Handler: router,
		Addr:    ":8081",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
