package http

import (
	"net/http"

	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"

	httpHelpers "github.com/quinlanmorake/verisart-go/helpers/http"
	parameterNames "github.com/quinlanmorake/verisart-go/helpers/http/parameterNames"
	
	user "github.com/quinlanmorake/verisart-go/user"
)

/*
 This generates a token that can be used to make any of the authenticated http requests
 It should be set as the value of the "Authorization" HTTP header
*/
func GenerateToken(w http.ResponseWriter, r *http.Request) {
	response := generateTokenResponse{}
	defer func() {
		httpHelpers.WriteResponse(w, response)
	}()

	if userIdAsRequestParameter, getUserIdResult := httpHelpers.GetParameterFromRequest(r, parameterNames.USER_ID, errorCodes.USER_ID_IS_INVALID); getUserIdResult.IsNotOk() {
		response.Error = getUserIdResult		
	} else {
		response.Jwt, response.Error = user.GenerateToken(coreTypes.UserId(userIdAsRequestParameter))
	}
}
