package response

import (
	"encoding/json"
	"log"
	"time"
)

/*TransactionCommitResponse struct with contain skeleton json to commitResponse*/
type TransactionCommitResponse struct {
	Vci        string `json:"vci"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	BuyOrder   string `json:"buy_order"`
	SessionID  string `json:"session_id"`
	CardDetail struct {
		CardNumber string `json:"card_number"`
	} `json:"card_detail"`
	AccountingDate     string    `json:"accounting_date"`
	TransactionDate    time.Time `json:"transaction_date"`
	AuthorizationCode  string    `json:"authorization_code"`
	PaymentTypeCode    string    `json:"payment_type_code"`
	ResponseCode       int       `json:"response_code"`
	InstallmentsNumber int       `json:"installments_number"`
}

/*GetTransactionCommitResponse fascade with return struct TransactionCommitResponse*/
func GetTransactionCommitResponse(response []byte) (TransactionCommitResponse, error) {
	var tcr TransactionCommitResponse

	err := json.Unmarshal(response, &tcr)

	if err != nil {
		log.Printf("Error en GetTransactionCommitResponse  %v", err)

		return tcr, err
	}

	return tcr, nil
}
