package cim

import "github.com/eliothedeman/authorize/auth"

type DeleteCustomerPaymentProfileRequest struct {
	*auth.MerchantAuth       `json:"merchantAuthentication,omitempty"`
	CustomerProfileId        string `json:"customerProfileId"`
	CustomerPaymentProfileId string `json:"customerPaymentProfileId"`
}

type DeleteCustomerPaymentProfileResponse struct {
}

func (c *DeleteCustomerPaymentProfileRequest) ResponseStruct() interface{} {
	return &DeleteCustomerPaymentProfileResponse{}
}

func (c *DeleteCustomerPaymentProfileRequest) SetAuth(a *auth.MerchantAuth) {
	c.MerchantAuth = a
}

func (c *DeleteCustomerPaymentProfileRequest) Method() string {
	return "deleteCustomerPaymentProfileRequest"
}
