package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var element []elements

type elements struct {
	AtomicMass              string `json:"atomicMass"`              //0
	AtomicNumber            string `json:"atomicNumber"`            //1
	AtomicRadius            string `json:"atomicRadius"`            //2
	BoilingPoint            string `json:"boilingPoint"`            //3
	BondingType             string `json:"bondingType"`             //4
	CpkHexColor             string `json:"cpkHexColor"`             //5
	Density                 string `json:"density"`                 //6
	ElectronAffinity        string `json:"electronAffinity"`        //7
	Electronegativity       string `json:"electronegativity"`       //8
	ElectronicConfiguration string `json:"electronicConfiguration"` //9
	GroupBlock              string `json:"groupBlock"`              //10
	IonRadius               string `json:"ionRadius"`               //11
	IonizationEnergy        string `json:"ionizationEnergy"`        //12
	MeltingPoint            string `json:"meltingPoint"`            //13
	Name                    string `json:"name"`                    //14
	OxidationStates         string `json:"oxidationStates"`         //15
	StandardState           string `json:"standardState"`           //16
	Symbol                  string `json:"symbol"`                  //17
	VanDelWaalsRadius       string `json:"vanDelWaalsRadius"`       //18
	YearDiscovered          string `json:"yearDiscovered"`          //19
}

func allElements(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(element)
}

func atomicNumber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	atomicNumber := vars["atomicNumber"]
	flag := false
	var finalData interface{}

	for _, data := range element {
		atomicNumber = strings.ToLower(atomicNumber)
		DataAtomicNumber := strings.ToLower(data.AtomicNumber)
		if strings.Compare(atomicNumber, DataAtomicNumber) == 0 {
			finalData = data
			flag = true
			break
		}
	}

	if flag {
		_ = json.NewEncoder(w).Encode(finalData)
	} else {
		err := make(map[string]interface{})
		err["status"] = false
		err["message"] = "Not Found"
		_ = json.NewEncoder(w).Encode(err)
	}
}

func atomicName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	atomicName := vars["atomicName"]
	flag := false
	var finalData interface{}

	for _, data := range element {
		atomicName = strings.ToLower(atomicName)
		DataAtomicName := strings.ToLower(data.Name)
		if strings.Compare(atomicName, DataAtomicName) == 0 {
			finalData = data
			flag = true
			break
		}
	}

	if flag {
		_ = json.NewEncoder(w).Encode(finalData)
	} else {
		err := make(map[string]interface{})
		err["status"] = false
		err["message"] = "Not Found"
		_ = json.NewEncoder(w).Encode(err)
	}
}

func atomicSymbol(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	atomicSymbol := vars["atomicSymbol"]
	flag := false
	var finalData interface{}

	for _, data := range element {
		atomicSymbol = strings.ToLower(atomicSymbol)
		DataAtomicSymbol := strings.ToLower(data.Symbol)
		if strings.Compare(atomicSymbol, DataAtomicSymbol) == 0 {
			finalData = data
			flag = true
			break
		}
	}

	if flag {
		_ = json.NewEncoder(w).Encode(finalData)
	} else {
		err := make(map[string]interface{})
		err["status"] = false
		err["message"] = "Not Found"
		_ = json.NewEncoder(w).Encode(err)
	}
}

func atomicBonding(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	atomicBonding := vars["atomicBonding"]
	flag := false
	var finalData []interface{}

	for _, data := range element {
		atomicBonding = strings.ToLower(atomicBonding)
		DataAtomicBonding := strings.ToLower(data.BondingType)
		if strings.Compare(atomicBonding, DataAtomicBonding) == 0 {
			finalData = append(finalData, data)
			flag = true
		}
	}

	if flag {
		_ = json.NewEncoder(w).Encode(finalData)
	} else {
		err := make(map[string]interface{})
		err["status"] = false
		err["message"] = "Not Found"
		_ = json.NewEncoder(w).Encode(err)
	}
}

func atomicGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	atomicGroup := vars["atomicGroup"]
	flag := false
	var finalData []interface{}

	for _, data := range element {
		atomicGroup = strings.ToLower(atomicGroup)
		DataAtomicGroup := strings.ToLower(data.GroupBlock)
		if strings.Compare(atomicGroup, DataAtomicGroup) == 0 {
			finalData = append(finalData, data)
			flag = true
		}
	}

	if flag {
		_ = json.NewEncoder(w).Encode(finalData)
	} else {
		err := make(map[string]interface{})
		err["status"] = false
		err["message"] = "Not Found"
		_ = json.NewEncoder(w).Encode(err)
	}
}

func atomicState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	atomicState := vars["atomicState"]
	flag := false
	var finalData []interface{}

	for _, data := range element {
		atomicState = strings.ToLower(atomicState)
		DataAtomicState := strings.ToLower(data.StandardState)
		if strings.Compare(atomicState, DataAtomicState) == 0 {
			finalData = append(finalData, data)
			flag = true
		}
	}

	if flag {
		_ = json.NewEncoder(w).Encode(finalData)
	} else {
		err := make(map[string]interface{})
		err["status"] = false
		err["message"] = "Not Found"
		_ = json.NewEncoder(w).Encode(err)
	}
}

func main() {

	// Reading JSON file
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error reading file")
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	_ = json.Unmarshal(byteValue, &element)

	// Routers
	router := mux.NewRouter()
	router.HandleFunc("/", allElements).Methods("GET")
	router.HandleFunc("/atomicNumber/{atomicNumber}", atomicNumber).Methods("GET")
	router.HandleFunc("/atomicName/{atomicName}", atomicName).Methods("GET")
	router.HandleFunc("/atomicSymbol/{atomicSymbol}", atomicSymbol).Methods("GET")
	router.HandleFunc("/atomicBonding/{atomicBonding}", atomicBonding).Methods("GET")
	router.HandleFunc("/atomicGroup/{atomicGroup}", atomicGroup).Methods("GET")
	router.HandleFunc("/atomicState/{atomicState}", atomicState).Methods("GET")

	//CORS Headers
	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET"})

	//Starting Server
	port := ":" + os.Getenv("PORT")
	err = http.ListenAndServe(port, handlers.CORS(originsOk, headersOk, methodsOk)(router))
	if err != nil {
		log.Fatal(err)
	}
}
