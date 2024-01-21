package network

type ResponseStatus int

const (
	Success ResponseStatus = iota
	InvalidRequest
	Unauthorized
	DataNotFound
	InternalServerError
	GeneralError
)

func (r ResponseStatus) GetDefaultResponseKey() string {
	return [...]string{"SUCCESS", "INVALID_REQUEST", "UNAUTHORIZED", "DATA_NOT_FOUND", "INTERNAL_SERVER_ERROR", "GENERAL_ERROR"}[r]
}

func (r ResponseStatus) GetDefaultResponseMsg() string {
	return [...]string{"Success", "Invalid Request", "Unauthorized", "Data Not Found", "Internal Server Error", "General Error"}[r]
}
