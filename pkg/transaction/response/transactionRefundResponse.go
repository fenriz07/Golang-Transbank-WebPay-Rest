package response

import (
	"encoding/json"
	"log"
	"time"
)

/*TransactionRefundResponse struct with contain skeleton json to refundResponse*/
type TransactionRefundResponse struct {
	Type              string    `json:"type"`
	AuthorizationCode string    `json:"authorization_code"`
	AuthorizationDate time.Time `json:"authorization_date"`
	NullifiedAmount   float64   `json:"nullified_amount"`
	Balance           float64   `json:"balance"`
	ResponseCode      int       `json:"response_code"`
}

/*GetTransactionRefundResponse fascade with return struct TransactionCommitResponse*/
func GetTransactionRefundResponse(response []byte) (TransactionRefundResponse, error) {
	var trr TransactionRefundResponse

	err := json.Unmarshal(response, &trr)

	if err != nil {
		log.Printf("Error in GetTransactionRefundResponse  %v", err)

		return trr, err
	}

	return trr, nil
}
