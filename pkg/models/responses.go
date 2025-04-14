package models

type ResponseError struct {
	Error string
}

func NewResponseError(err error) *ResponseError {
	return &ResponseError{
		Error: err.Error(),
	}
}

type ResponseData struct {
	Data interface{}
}

func NewResponseData(data interface{}) *ResponseData {
	return &ResponseData{
		Data: data,
	}
}
