package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"bartekpacia/bit-festival-2023/algo"

	"github.com/gorilla/mux"
)

type ReqData struct {
	Ampacity       float64            `json:"ampacity"`
	MaxPower       float64            `json:"maxPower"`
	VeinsUnderLoad int                `json:"veinsUnderLoad"`
	Placements     algo.CableLocation `json:"placements"`
	Temperature    float64            `json:"temperature"`
}

type RespData struct {
	CableType      string             `json:"cableType"`
	VeinsNumber    int                `json:"veinsNumber"`
	CrossSection   float64            `json:"crossSection"`
	InsulationType string             `json:"insulationType"`
	Placements     algo.CableLocation `json:"placements"`
	Temperature    float64            `json:"temperature"`
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	var req ReqData
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("%+v\n", req)

	var resp RespData

	I_ost, err := algo.CalcTemp(req.Ampacity, req.MaxPower, req.Temperature, req.VeinsUnderLoad)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("I_ost:", I_ost)

	przekroj, err := algo.MatchCrossection(I_ost, req.VeinsUnderLoad, req.Placements)
	if err != nil {
		log.Fatalln("failed to match crossection:", err)
	}

	resp.CrossSection = przekroj
	resp.CableType = "YDY"
	resp.VeinsNumber = req.VeinsUnderLoad
	resp.InsulationType = "PVC"
	resp.Placements = req.Placements
	resp.Temperature = req.Temperature

	jData, err := json.Marshal(resp)
	if err != nil {
		// handle error
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(jData)
}

func main() {
	algo.LoadPath = "algo/"
	err := algo.Init()
	if err != nil {
		log.Fatalln(err)
	}

	I_obl := 18.0 // A
	temp := 27.0  // °C
	P := 0.0
	fmt.Printf("I_obl: %f A, temp: %f °C\n", I_obl, temp)
	I_ost, err := algo.CalcTemp(I_obl, P, temp, 3)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("I_ost:", I_ost)

	przekroj, err := algo.MatchCrossection(I_ost, 3, algo.A1)
	if err != nil {
		log.Fatalln("failed to match crossection:", err)
	}

	fmt.Println("przekroj:", przekroj)

	r := mux.NewRouter()
	r.HandleFunc("/", handlePost).Methods("POST")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	srv.ListenAndServe()
}
