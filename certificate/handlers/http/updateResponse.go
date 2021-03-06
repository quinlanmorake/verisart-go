package http

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	businessTypes "github.com/quinlanmorake/verisart-go/types"	
)

type updateResponse struct {
	Error coreTypes.Result `json:"error"`
	Certificate businessTypes.Certificate `json:"certificate"`
}