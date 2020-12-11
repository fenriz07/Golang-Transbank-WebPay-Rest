package transaction

import (
	"log"

	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/client"
)

const createTransactionEndpoint = "rswebpaytransaction/api/webpay/v1.0/transactions"
const commitTransactionEndpoint = "rswebpaytransaction/api/webpay/v1.0/transactions"
const refundTransactionEndpoint = "rswebpaytransaction/api/webpay/v1.0/transactions/$TOKEN$/refunds"
const getTransactionStatusEndpoint = "rswebpaytransaction/api/webpay/v1.0/transactions/$TOKEN$/refunds"
const captureEndpoint = "rswebpaytransaction/api/webpay/v1.0/transactions/$TOKEN$/capture"

/*Create create transaction*/
func Create(buyOrder string, sessionID string, amount int, returnURL string) (interface{}, error) {

	body := map[string]interface{}{
		"buy_order":  buyOrder,
		"session_id": sessionID,
		"amount":     amount,
		"return_url": returnURL,
	}

	httpClient := client.GetInstance()

	resp, err := httpClient.Post(createTransactionEndpoint, body)

	if err != nil {
		log.Printf("Transaction fail in create method.  \n%v\n", err)
	}

	return resp, err
}
