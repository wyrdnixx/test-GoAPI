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

type Result struct {
	Msg   string `json:"msg"`
	Value bool   `json:"value"`
}
type Uuid struct {
	Id string
}

func checkUserCookie(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")

	//res := Result{}
	c := models.Cookie{}
	var checkUUID Uuid

	reqBody, err := ioutil.ReadAll(r.Body)
	fmt.Printf("checkUserCookie got: %s\n", string(reqBody))
	if err != nil {
		fmt.Printf("error checking cookie: %v", err.Error())
		//res.Msg = "cookieValid"
		//res.Value = false

		c.Id = "-"
		c.Status = err.Error()

		//json.NewEncoder(w).Encode(res)
	} else {
		//err := json.NewDecoder(r.Body).Decode(&checkUUID)
		err := json.Unmarshal(reqBody, &checkUUID)
		if err != nil {
			fmt.Println("Json decoder checkUUID error: ", err.Error())
			//res.Msg = "cookieValid"
			//res.Value = false

			c.Status = err.Error()

		} else {
			fmt.Println("checkUserUUID: ", checkUUID.Id)
			cTemp, err := lib.CheckCookie(checkUUID.Id)
			c = cTemp // copy result to returned cookie
			if err != nil {
				fmt.Println("checkCookie error: ", err.Error())

				c.Status = string(err.Error())
				fmt.Println("checkCookie error string : ", c.Status)
			}

		}

	}

	//	c.Id = "test"
	//	c.Status = "teststatis"
	fmt.Println("checkCookie vor response string : ", c)
	response, err := json.Marshal(c)

	if err != nil {
		fmt.Println("checkUserCookie : Error convert response to json: ", err.Error())
	}
	//fmt.Println("response: ", response)

	//time.Sleep(5 * time.Second)
	fmt.Println("sending answer: ", string(response))
	// w.WriteHeader(http.StatusAccepted)
	w.Write(response)

	//json.NewEncoder(w).Encode(res)
}

/*
func checkUserUUID(c string) bool {
	if c == "de070071-0b1a-45a5-84d0-cc89d631a960" {
		return true
	} else {
		return false
	}
}
*/

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

type event struct {
	ID        string `json:"ID"`
	Type      string `json:"Type"`
	EventData []byte `json:"EventData"`
}
type newAuthRequest struct {
	Username string `njson:"Username"`
	DeviceID string `njson:"DeviceID"`
}

func jsonapi(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	fmt.Printf("jsonApi Got: :%v\n", string(reqBody))
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	if newEvent.Type == "ExampleCreateRequest" {
		var na newAuthRequest

		// Decode the json object
		var j map[string]interface{}
		err2 := json.Unmarshal([]byte(reqBody), &j)
		if err2 != nil {
			panic(err)
		}
		parents, noErr := j["EventData"].(map[string]interface{})
		if noErr != true {
			fmt.Printf("Wrong object\n")
		} else {
			fmt.Printf("-> %v\n", parents["Username"])
			fmt.Printf("-> %v\n", parents["DeviceID"])
			na.Username = parents["Username"].(string)
			na.DeviceID = parents["DeviceID"].(string)
			fmt.Printf("type: %v - %v \n", na.Username, na.DeviceID)
		}

	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func main() {

	// HL7 Listener
	hl7server := tcpserver.NewServer(&tcpserver.Config{
		Host: "0.0.0.0",
		Port: "6661",
	})
	go hl7server.Run()

	//Database
	lib.Initdb()

	// WEbAPI Router
	router := mux.NewRouter()

	router.HandleFunc("/api/jsonapi", jsonapi)
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
