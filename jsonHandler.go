package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func rintisReqHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("New Request To Rintis")

	// Get msg from body
	body, _ := ioutil.ReadAll(r.Body)

	var request rintisRequest

	// Unmarshal body to struct
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("Error unmarshal JSON: %s", err.Error())
		return
	}

	response := addDataToResponse(request)

	fmt.Println(response)

	rand.Seed(time.Now().UnixNano())
	nice := rand.Intn(3000)
	time.Sleep(time.Duration(nice) * time.Millisecond)

	responseFormatter(w, response, 200)
}

func responseFormatter(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func addDataToResponse(data rintisRequest) (response rintisResponse) {
	response.Pan = data.Pan
	response.ProcessingCode = data.ProcessingCode
	response.TotalAmount = data.TotalAmount
	response.TransmissionDateTime = data.TransmissionDateTime
	response.Stan = data.Stan
	response.LocalTransactionTime = data.LocalTransactionTime
	response.LocalTransactionDate = data.LocalTransactionDate
	response.SettlementDate = data.CaptureDate[:4]
	response.CaptureDate = data.CaptureDate
	response.AcquirerID = data.AcquirerID
	response.Track2Data = data.Track2Data
	response.Refnum = data.Refnum
	response.AuthIdResponse = "123456"
	response.ResponseCode = "12"
	response.TerminalID = data.TerminalID
	response.AdditionalResponseData = "1234567890123456"
	response.Currency = data.Currency
	response.TerminalData = data.TerminalData
	response.ReceivingInstitutionIDCode = "1234567890123456"
	response.AccountTo = data.AccountTo
	response.TokenData = data.TokenData

	return response
}
