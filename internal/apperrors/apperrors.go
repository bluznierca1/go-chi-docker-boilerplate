package apperrors

type Error struct {
	ErrorCode string  `json:"error_code"`
	ErrorMsg  string  `json:"error_msg"`
	FieldName *string `json:"field_name"`
}

const (
	ErrInternalServerError = "internal_server_error"
)
