package response

import (
	"encoding/json"
	"log"
	"time"
)

/*TransactionStatusResponse struct with contain skeleton json to statusResponse*/
type TransactionStatusResponse struct {
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

/*GetTransactionStatusResponse fascade with return struct TransactionCommitResponse*/
func GetTransactionStatusResponse(response []byte) (TransactionStatusResponse, error) {
	var tsr TransactionStatusResponse

	err := json.Unmarshal(response, &tsr)

	if err != nil {
		log.Printf("Error in GetTransactionStatusResponse  %v", err)

		return tsr, err
	}

	return tsr, nil
}
