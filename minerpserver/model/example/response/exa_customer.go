package response

import "github.com/JoeyFrancisTribbiani/minerpserver/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
