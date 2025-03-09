package customer

import (
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/paymentmethod"
)

var CustomerBillingDetails stripe.PaymentMethodBillingDetailsParams

func SetBillingDetails(name string, email string, phone string, city string, country string, ln1 string, ln2 string, zip string, state string) {
	var address stripe.AddressParams
	address.City = stripe.String(city)
	address.Country = stripe.String(country)
	address.Line1 = stripe.String(ln1)
	address.Line2 = stripe.String(ln2)
	address.PostalCode = stripe.String(zip)
	address.State = stripe.String(state)
	CustomerBillingDetails = stripe.PaymentMethodBillingDetailsParams{
		Name:    stripe.String(name),
		Email:   stripe.String(email),
		Phone:   stripe.String(phone),
		Address: &address,
	}
}

func SetPaymentMethodUSBANK(acct string, route string) (*stripe.PaymentMethod, error) {

	params := &stripe.PaymentMethodParams{
		Type: stripe.String(string(stripe.PaymentMethodTypeUSBankAccount)),
		USBankAccount: &stripe.PaymentMethodUSBankAccountParams{
			AccountHolderType: stripe.String(string(stripe.PaymentMethodUSBankAccountAccountHolderTypeIndividual)),
			AccountNumber:     stripe.String(acct),
			RoutingNumber:     stripe.String(route),
		},
		BillingDetails: &stripe.PaymentMethodBillingDetailsParams{
			Name:    stripe.String(*CustomerBillingDetails.Name),
			Email:   stripe.String(*CustomerBillingDetails.Email),
			Phone:   stripe.String(*CustomerBillingDetails.Phone),
			Address: CustomerBillingDetails.Address,
		},
	}
	result, err := paymentmethod.New(params)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func SetPaymentMethodCard(num string, month int64, year int64, cvc string) (*stripe.PaymentMethod, error) {

	params := &stripe.PaymentMethodParams{
		Type: stripe.String(string(stripe.PaymentMethodTypeCard)),
		Card: &stripe.PaymentMethodCardParams{
			Number:   stripe.String(num),
			ExpMonth: stripe.Int64(month),
			ExpYear:  stripe.Int64(year),
			CVC:      stripe.String(cvc),
		},
		BillingDetails: &stripe.PaymentMethodBillingDetailsParams{
			Name:    stripe.String(*CustomerBillingDetails.Name),
			Email:   stripe.String(*CustomerBillingDetails.Email),
			Phone:   stripe.String(*CustomerBillingDetails.Phone),
			Address: CustomerBillingDetails.Address,
		},
	}
	result, err := paymentmethod.New(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
