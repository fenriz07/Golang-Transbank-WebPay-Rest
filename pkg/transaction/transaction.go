package transaction

import (
	"fmt"
	"log"

	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/client"
	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/transaction/response"
)

const createTransactionEndpoint = "rswebpaytransaction/api/webpay/v1.0/transactions"
const commitTransactionEndpoint = "rswebpaytransaction/api/webpay/v1.0/transactions"
const refundTransactionEndpoint = "rswebpaytransaction/api/webpay/v1.0/transactions/$TOKEN$/refunds"
const getTransactionStatusEndpoint = "rswebpaytransaction/api/webpay/v1.0/transactions/%s/refunds"
const captureEndpoint = "rswebpaytransaction/api/webpay/v1.0/transactions/$TOKEN$/capture"

/*Create create transaction*/
func Create(buyOrder string, sessionID string, amount int, returnURL string) (response.TransactionCreateResponse, error) {

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

	transactionCreateResponse, err := response.GetTransactionCreateResponse(resp.Body())

	return transactionCreateResponse, err
}

/*Commit check the status of a transaction*/
func Commit(token string) (response.TransactionCommitResponse, error) {
	httpClient := client.GetInstance()

	endpoint := fmt.Sprintf("%s/%s", commitTransactionEndpoint, token)

	resp, err := httpClient.Put(endpoint)

	if err != nil {
		log.Printf("Transaction fail in commit method.  \n%v\n", err)
	}

	transactionCommitResponse, err := response.GetTransactionCommitResponse(resp.Body())

	return transactionCommitResponse, err
}

func getStatus(token string) {

	httpClient := client.GetInstance()

	endpoint := fmt.Sprintf(getTransactionStatusEndpoint, token)

	resp, err := httpClient.Get(endpoint)

	if err != nil {
		log.Printf("Transaction fail in commit method.  \n%v\n", err)
	}

	fmt.Println(resp)

}
