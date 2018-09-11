package errorCodes

/*
 We put these in a subpackage so as to make it clearer to use them in the code; that is, the intended use is
  a := errorCodes.USER_EMAIL_IS_INVALID

 We are adding in the number such that a particular error code is easy to find when looking through the code.
 Consumers of the api would be looking at the docs for this information
*/

// This is used in unit tests when we need to return a failure but don't care about the error
const UNIT_TEST_GENERIC_FAILURE = -1

const (
	// 1 - 10
	INVALID_ROUTE	ErrorCode = iota + 1
	INVALID_HTTP_METHOD
	INVALID_REQUEST_BODY 

	ERROR_GENERATING_UUID
	
	USER_ID_IS_INVALID 
	EMAIL_IS_INVALID
	NO_USER_WITH_THE_EMAIL_EXISTS
	NO_USER_WITH_THE_GIVEN_ID_EXISTS

	JWT_ERROR_PARSING_THE_PRIVATE_KEY
	JWT_ERROR_PARSING_THE_PUBLIC_KEY
	JWT_TOKEN_HAS_INVALID_ISSUER
	JWT_TOKEN_HAS_EXPIRED
	JWT_TOKEN_SIGNING_ERROR
	JWT_COULD_NOT_DECODE_TOKEN
	JWT_COULD_NOT_UNMARSHAL_TOKEN
	JWT_AUTHORIZATION_HEADER_WAS_NOT_SET

	INVALID_DATABASE_CONFIG
	DATABASE_INIT_ERROR
	DATABASE_NOT_INITIALIZED

	DATABASE_ADD_DID_NOT_RETURN_AN_ID
	DATABASE_ADD_OBJECT_ALREADY_HAS_AN_ID
	DATABASE_ADD_FAILED

	DATABASE_LOAD_FAILED

	DATABASE_UPDATE_NO_ID_WAS_PROVIDED
	DATABASE_UPDATE_FAILED

	DATABASE_DELETE_NO_ID_WAS_PROVIDED
	DATABASE_DELETE_FAILED

	CONFIG_INIT_COULD_NOT_GET_CURRENT_EXE_PATH
	CONFIG_INIT_SYMBOL_LINK_ERROR
	CONFIG_INIT_COULD_NOT_FIND_CONFIG_FILE
	CONFIG_INIT_COULD_NOT_READ_FILE
	CONFIG_INIT_COULD_NOT_UNMARSHAL_CONFIG_FILE

	ERROR_UNMARSHALING_USER_RECORD_FROM_DATABASE

	CERTIFICATE_MISSING_TITLE
	CERTIFICATE_HAS_INVALID_YEAR

	ERROR_UNMARSHALING_CERTIFICATE_RECORD_FROM_DATABASE
	NO_CERTIFICATE_WITH_THE_GIVEN_ID_COULD_BE_FOUND
	INSUFFICIENT_PERMISSIONS_TO_MODIFY_THE_REQUESTED_CERTIFICATE
	CERTIFICATE_ID_IS_INVALID
	
	TRANSFER_OBJECT_HAS_NO_CREATED_AT_VALUE	
	CERTIFICATE_ALREADY_HAS_PENDING_TRANSFER
	ERROR_UNMARSHALING_TRANSFER_RECORD_FROM_DATABASE
	NO_TRANSFER_FOR_THE_GIVEN_CERTIFICATE_ID_COULD_BE_FOUND
	UNABLE_TO_TRANSFER_TO_YOURSELF
	INSUFFICIENT_PERMISSION_TO_ACCEPT_THE_TRANSFER
	
	// 11 - 20
)
