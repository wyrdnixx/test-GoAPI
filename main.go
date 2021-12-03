package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"onboarding/lib"
	_ "onboarding/lib"
	"onboarding/models"
	"onboarding/tcpserver"
	_ "onboarding/tcpserver"
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

func checkUserCookie(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	var res struct {
		msg   string
		value bool
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	fmt.Printf("checkUserCookie got: %s\n", string(reqBody))
	if err != nil {
		fmt.Printf("error checking cookie: ", err.Error())
		res.msg = "cookieValid"
		res.value = false
		json.NewEncoder(w).Encode(res)
	} else {
		fmt.Printf("checking cookie: %s\n", string(reqBody))

		res.msg = "cookieValid"
		res.value = true

		response, _ := json.Marshal(res)
		fmt.Println("response: ", response)

		//time.Sleep(5 * time.Second)
		fmt.Println("sending answer")
		w.Header().Set("Contend-Type", "application/json")
		w.Write(response)

		//json.NewEncoder(w).Encode(res)

	}
}

func createFirma(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	reqBody, err := ioutil.ReadAll(r.Body)
	fmt.Printf("createFirma got: %s\n", string(reqBody))
	if err != nil {
		fmt.Printf("createFirma error: %s \n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		m := models.NewCompanyMessage{}
		err := json.Unmarshal(reqBody, &m)
		if err != nil {
			fmt.Printf("createFirma Error on json.unmarshal: %s\n", err)
			//http.Error(w, err.Error(), http.StatusInternalServerError)

		} else {
			//fmt.Printf("Vor insert: %s\n", m.NewCompany.Name)
			//fmt.Printf("Vor insert: %s\n", strconv.Itoa(m.NewCompany.Enabled))

			lib.InsertFirmen(m.NewCompany)
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

	// HL7 Listener
	hl7server := tcpserver.NewServer(&tcpserver.Config{
		Host: "0.0.0.0",
		Port: "666",
	})
	go hl7server.Run()

	//Database
	lib.Initdb()

	// WEbAPI Router
	router := mux.NewRouter()

	router.HandleFunc("/api/delFirma", delFirma)
	router.HandleFunc("/api/createFirma", createFirma)
	router.HandleFunc("/api/checkUserCookie", checkUserCookie)
	router.HandleFunc("/api/getFirmen", func(w http.ResponseWriter, r *http.Request) {

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
			//fmt.Printf("api/health returning to client: %s\n", res.Firmen)
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
