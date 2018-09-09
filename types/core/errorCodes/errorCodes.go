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
	USER_ID_IS_INVALID ErrorCode = iota + 1
	EMAIL_IS_INVALID
	NO_USER_WITH_THE_EMAIL_EXISTS

	TRANSFER_OBJECT_HAS_NO_CREATED_AT_VALUE

	INVALID_HTTP_METHOD
	INVALID_ROUTE

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

	ERROR_GENERATING_UUID
	// 11 - 20
)
