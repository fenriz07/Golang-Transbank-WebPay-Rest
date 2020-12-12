package transaction

import (
	"fmt"
	"log"

	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/client"
	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/transaction/response"
)

const createTransactionEndpoint = "rswebpaytransaction/api/webpay/v1.0/transactions"
const commitTransactionEndpoint = "rswebpaytransaction/api/webpay/v1.0/transactions"
const getTransactionStatusEndpoint = "rswebpaytransaction/api/webpay/v1.0/transactions/%s"
const refundTransactionEndpoint = "rswebpaytransaction/api/webpay/v1.0/transactions/%s/refunds"

/*Create It allows to initialize a transaction in Webpay.
In response to the invocation, a token is generated that uniquely represents a transaction.*/
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

/*Commit It allows you to confirm and obtain the result of the transaction once Webpay
has resolved your financial authorization.*/
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

/*GetStatus This operation allows you to obtain the status of the transaction at any time.
Under normal conditions, it is probably not required to execute.
But in the event of an unexpected error,
it allows knowing the status and taking the corresponding actions.*/
func GetStatus(token string) (response.TransactionStatusResponse, error) {

	httpClient := client.GetInstance()

	endpoint := fmt.Sprintf(getTransactionStatusEndpoint, token)

	resp, err := httpClient.Get(endpoint)

	if err != nil {
		log.Printf("Transaction fail in GetStaus method.  \n%v\n", err)
	}

	transactionStatusResponse, err := response.GetTransactionStatusResponse(resp.Body())

	return transactionStatusResponse, err
}

/*Refund It allows you to request from Webpay the cancellation of a transaction made previously
and that is currently in force.*/
func Refund(token string, amount int) (response.TransactionRefundResponse, error) {
	httpClient := client.GetInstance()

	endpoint := fmt.Sprintf(refundTransactionEndpoint, token)

	body := map[string]interface{}{
		"amount": amount,
	}

	resp, err := httpClient.Post(endpoint, body)

	if err != nil {
		log.Printf("Transaction fail in commit method.  \n%v\n", err)
	}

	transactionRefundResponse, err := response.GetTransactionRefundResponse(resp.Body())

	if err != nil {
		return transactionRefundResponse, err
	}

	return transactionRefundResponse, err
}
