package apperrors

type Error struct {
	ErrorCode string  `json:"error_code"`
	ErrorMsg  string  `json:"error_msg"`
	FieldName *string `json:"field_name"`
}

const (
	InternalServerError          = "internal_server_error"
	IntegerGreaterThanZero       = "integer_greater_than_zero"
	PaginationParamsNotExtracted = "pagination_params_not_extracted"
)

var errorMessages = map[string]string{
	InternalServerError:          "Internal server error - Unable to handle request",
	IntegerGreaterThanZero:       "The field is not integer greater than zero",
	PaginationParamsNotExtracted: "Pagination parameters not extracted",
}

// GetErrorDetails returns the Error struct for a given error code
func GetErrorDetails(code string, fieldName *string, msg *string, param *string) Error {
	var errMsg string
	if msg == nil {
		var exists bool
		errMsg, exists = errorMessages[code]
		if !exists {
			errMsg = "Unknown error"
		}
	} else {
		errMsg = *msg
	}

	return Error{
		ErrorCode: code,
		ErrorMsg:  errMsg,
		FieldName: fieldName,
	}
}
