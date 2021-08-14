package response

import "FiberBoot/model/example"

type CustomerResponse struct {
	Customer example.Customer `json:"customer"`
}
