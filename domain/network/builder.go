package network

type ApiResponse[T any] struct {
	ResponseKey     string `json:"response_key"`
	ResponseMessage string `json:"response_message"`
	Status          bool   `json:"status"`
	Data            T      `json:"data"`
}

func Null() interface{} {
	return nil
}

func BuildSuccessResponse[T any](responseStatus ResponseStatus, data T) ApiResponse[T] {
	return buildResponse(responseStatus.GetDefaultResponseKey(), responseStatus.GetDefaultResponseMsg(), true, data)
}

func BuildErrorResponse(responseStatus ResponseStatus) ApiResponse[interface{}] {
	return buildResponse(responseStatus.GetDefaultResponseKey(), responseStatus.GetDefaultResponseMsg(), false, Null())
}

func BuildCustomResponse(responseStatus ResponseStatus, msg string, status bool) ApiResponse[interface{}] {
	return buildResponse(responseStatus.GetDefaultResponseKey(), msg, status, Null())
}

func buildResponse[T any](statusKey string, message string, status bool, data T) ApiResponse[T] {
	return ApiResponse[T]{
		ResponseKey:     statusKey,
		ResponseMessage: message,
		Status:          status,
		Data:            data,
	}
}
