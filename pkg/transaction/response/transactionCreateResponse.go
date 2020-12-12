package response

import (
	"encoding/json"
	"log"
)

/*TransactionCreateResponse struct with contain skeleton json to createResponse*/
type TransactionCreateResponse struct {
	Token string `json:"token"`
	URL   string `json:"url"`
}

/*GetTransactionCreateResponse fascade with return struct transactionCreateResponse*/
func GetTransactionCreateResponse(response []byte) (TransactionCreateResponse, error) {
	var tcr TransactionCreateResponse

	err := json.Unmarshal(response, &tcr)

	if err != nil {
		log.Printf("Error en GetTransactionCreateResponse  %v", err)

		return tcr, err
	}

	return tcr, nil
}
