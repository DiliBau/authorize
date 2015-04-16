package cim

import "github.com/eliothedeman/authorize/auth"

type GetCustomerProfileRequest struct {
	*auth.MerchantAuth `json:"merchantAuthentication,omitempty"`
	RefId              string `json:"refId,omitempty"`
	CustomerProfileId  string `json:"customerProfileId"`
}

type GetCustomerProfileResponse struct {
}

func (c *GetCustomerProfileRequest) ResponseStruct() interface{} {
	return 0
}

func (c *GetCustomerProfileRequest) SetAuth(a *auth.MerchantAuth) {
	c.MerchantAuth = a
}

func (c *GetCustomerProfileRequest) Method() string {
	return "getCustomerProfileRequest"
}
