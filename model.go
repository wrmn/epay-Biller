package main

type rintisResponse struct {
	Pan                        string `json:"pan"`
	ProcessingCode             string `json:"processingCode"`
	TotalAmount                int    `json:"totalAmount"`
	TransmissionDateTime       string `json:"transmissionDateTime"`
	Stan                       string `json:"stan"`
	LocalTransactionTime       string `json:"localTransactionTime"`
	LocalTransactionDate       string `json:"localTransactionDate"`
	SettlementDate             string `json:"settlementDate"`
	CaptureDate                string `json:"captureDate"`
	AcquirerID                 string `json:"acquirerId"`
	Track2Data                 string `json:"track2Data"`
	Refnum                     string `json:"refnum"`
	AuthIdResponse             string `json:"authorizationIdentificationNumber"`
	ResponseCode               string `json:"responseCode"`
	TerminalID                 string `json:"terminalId"`
	AdditionalResponseData     string `json:"additionalResponseData"`
	Currency                   string `json:"currency"`
	TerminalData               string `json:"terminalData"`
	ReceivingInstitutionIDCode string `json:"receivingInstitutionIDCode"`
	AccountTo                  string `json:"accountTo"`
	TokenData                  string `json:"tokenData"`
}

type rintisRequest struct {
	Pan                  string `json:"pan"`
	ProcessingCode       string `json:"processingCode"`
	TotalAmount          int    `json:"totalAmount"`
	TransmissionDateTime string `json:"transmissionDateTime"`
	Stan                 string `json:"stan"`
	LocalTransactionTime string `json:"localTransactionTime"`
	LocalTransactionDate string `json:"localTransactionDate"`
	CaptureDate          string `json:"captureDate"`
	AcquirerID           string `json:"acquirerId"`
	Track2Data           string `json:"track2Data"`
	Refnum               string `json:"refnum"`
	TerminalID           string `json:"terminalId"`
	CardAcceptorData     string `json:"cardAcceptorData"`
	AdditionalData       string `json:"additionalData"`
	Currency             string `json:"currency"`
	PIN                  string `json:"personalIdentificationNumber"`
	TerminalData         string `json:"terminalData"`
	AccountTo            string `json:"accountTo"`
	TokenData            string `json:"tokenData"`
}
