package stripe

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/balance"
)

func getStripeKey() (string, error) {

	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	return os.Getenv("STRIPE_PK_TEST"), nil
}

var stripeKey string
var err error

func InitStripe() {
	stripeKey, err = getStripeKey()
	if err != nil {
		panic(err)
	}
	stripe.Key = stripeKey
	GetBalance()
}

func GetBalance() (*stripe.Balance, error) {

	params := &stripe.BalanceParams{}
	result, err := balance.Get(params)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Current Availible Stripe Balance: $%d ", result.Available[0].Amount)
	return result, nil
}
